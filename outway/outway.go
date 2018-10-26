// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"hash/crc64"
	"reflect"
	"sync"
	"time"

	"go.nlx.io/nlx/common/process"

	"github.com/jmoiron/sqlx"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.nlx.io/nlx/common/orgtls"
	"go.nlx.io/nlx/common/transactionlog"
	"go.nlx.io/nlx/directory/directoryapi"
)

// Outway handles requests from inside the organization
type Outway struct {
	wg               *sync.WaitGroup
	organizationName string // the organization running this outway

	tlsOptions orgtls.TLSOptions
	tlsRoots   *x509.CertPool

	logger *zap.Logger

	txlogger transactionlog.TransactionLogger

	directoryClient directoryapi.DirectoryClient

	requestFlake *sonyflake.Sonyflake
	ecmaTable    *crc64.Table

	servicesLock sync.RWMutex
	services     map[string]HTTPService // services mapped by <organizationName>.<serviceName>
}

// NewOutway creates a new Outway and sets it up to handle requests.
func NewOutway(process *process.Process, logger *zap.Logger, logdb *sqlx.DB, tlsOptions orgtls.TLSOptions, directoryAddress string) (*Outway, error) {
	// load certs and get organization name from cert
	roots, orgCert, err := orgtls.Load(tlsOptions)
	if err != nil {
		logger.Fatal("failed to load tls certs", zap.Error(err))
	}
	if len(orgCert.Subject.Organization) != 1 {
		return nil, errors.New("cannot obtain organization name from self cert")
	}
	organizationName := orgCert.Subject.Organization[0]

	o := &Outway{
		wg:               &sync.WaitGroup{},
		logger:           logger.With(zap.String("outway-organization-name", organizationName)),
		organizationName: organizationName,

		tlsOptions: tlsOptions,
		tlsRoots:   roots,

		requestFlake: sonyflake.NewSonyflake(sonyflake.Settings{}),
		ecmaTable:    crc64.MakeTable(crc64.ECMA),
	}

	// setup transactionlog
	if logdb == nil {
		o.txlogger = transactionlog.NewDiscardTransactionLogger()
	} else {
		o.txlogger, err = transactionlog.NewPostgresTransactionLogger(logger, logdb, transactionlog.DirectionOut)
		if err != nil {
			return nil, errors.Wrap(err, "failed to setup transactionlog")
		}
	}

	orgKeypair, err := tls.LoadX509KeyPair(tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read tls keypair")
	}
	directoryDialCredentials := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{orgKeypair},
		RootCAs:      roots,
	})
	directoryDialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(directoryDialCredentials),
	}
	directoryConnCtx, directoryConnCtxCancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer directoryConnCtxCancel()
	directoryConn, err := grpc.DialContext(directoryConnCtx, directoryAddress, directoryDialOptions...)
	if err != nil {
		logger.Fatal("failed to setup connection to directory service", zap.Error(err))
	}
	o.directoryClient = directoryapi.NewDirectoryClient(directoryConn)
	err = o.updateServiceList(process)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update internal service directory")
	}

	go o.keepServiceListUpToDate(process)
	return o, nil
}

func (o *Outway) keepServiceListUpToDate(process *process.Process) {
	o.wg.Add(1)
	defer o.wg.Done()

	expBackOff := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Factor: 2,
		Max:    20 * time.Second,
	}

	const baseInterval = 30 * time.Second
	interval := baseInterval
	for {
		select {
		case <-process.ShutdownComplete:
			return
		case <-time.After(interval):
			err := o.updateServiceList(process)
			if err != nil {
				o.logger.Warn("failed to update the service list", zap.Error(err))
				interval = expBackOff.Duration() // Changing interval on retry
			} else {
				interval = baseInterval // Resetting interval to base on success update
				expBackOff.Reset()
			}
		}
	}
}

func (o *Outway) updateServiceList(process *process.Process) error {
	services := make(map[string]HTTPService)
	resp, err := o.directoryClient.ListServices(context.Background(), &directoryapi.ListServicesRequest{})
	if err != nil {
		return errors.Wrap(err, "failed to fetch services from directory")
	}
	o.servicesLock.Lock()
	defer o.servicesLock.Unlock()
	shutDown := make(chan struct{})
	process.CloseGracefully(func() error {
		close(shutDown)
		return nil
	})
	for _, serviceToImplement := range resp.Services {
		select {
		case <-shutDown:
			// On app shutdown we have no need to update services.
			// So we need to wait until started updated is finished and exit
			return nil
		default:
			// Need default to not to block
		}

		service, exists := o.services[serviceToImplement.OrganizationName+"."+serviceToImplement.ServiceName]
		if !exists || !reflect.DeepEqual(service.GetInwayAddresses(), serviceToImplement.InwayAddresses) {
			// create the service
			rrlbService, err := NewRoundRobinLoadBalancedHTTPService(o.logger, o.tlsRoots, o.tlsOptions.OrgCertFile, o.tlsOptions.OrgKeyFile, serviceToImplement.OrganizationName, serviceToImplement.ServiceName, serviceToImplement.InwayAddresses)
			if err != nil {
				if err == errNoInwaysAvailable {
					o.logger.Info("service exists but there are no inwayaddresses available", zap.String("service-organization-name", serviceToImplement.OrganizationName), zap.String("service-name", serviceToImplement.ServiceName))
					continue
				}
				o.logger.Error("failed to create new service", zap.String("service-organization-name", serviceToImplement.OrganizationName), zap.String("service-name", serviceToImplement.ServiceName), zap.Error(err))
				continue
			}
			service = rrlbService

		}
		services[service.FullName()] = service
	}

	o.services = services
	return nil
}
