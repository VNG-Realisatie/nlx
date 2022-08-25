// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package transactionlog

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/txlog-api/api"
)

type APITransactionLogger struct {
	logger     *zap.Logger
	direction  api.CreateRecordRequest_Direction
	client     api.TXLogClient
	connection *grpc.ClientConn
	cancelFunc context.CancelFunc
}

type NewAPITransactionLoggerArgs struct {
	Logger       *zap.Logger
	Direction    Direction
	APIAddress   string
	InternalCert *common_tls.CertificateBundle
}

func NewAPITransactionLogger(args *NewAPITransactionLoggerArgs) (TransactionLogger, error) {
	var direction api.CreateRecordRequest_Direction

	switch args.Direction {
	case DirectionIn:
		direction = api.CreateRecordRequest_IN

	case DirectionOut:
		direction = api.CreateRecordRequest_OUT

	default:
		return nil, errors.New("invalid direction value")
	}

	dialCredentials := credentials.NewTLS(args.InternalCert.TLSConfig())
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(dialCredentials),
	}

	var grpcTimeout = 5 * time.Second

	grpcCtx, cancel := context.WithTimeout(context.Background(), grpcTimeout)

	txlogConn, err := grpc.DialContext(grpcCtx, args.APIAddress, dialOptions...)
	if err != nil {
		cancel()
		return nil, err
	}

	txlogClient := api.NewTXLogClient(txlogConn)

	result := &APITransactionLogger{
		logger:     args.Logger,
		direction:  direction,
		client:     txlogClient,
		connection: txlogConn,
		cancelFunc: cancel,
	}

	return result, nil
}

func (txl *APITransactionLogger) AddRecord(rec *Record) error {
	dataJSON, err := json.Marshal(rec.Data)
	if err != nil {
		return errors.Wrap(err, "failed to convert data to json")
	}

	dataSubjects := make([]*api.CreateRecordRequest_DataSubject, len(rec.DataSubjects))

	i := 0

	for key, value := range rec.DataSubjects {
		dataSubjects[i] = &api.CreateRecordRequest_DataSubject{
			Key:   key,
			Value: value,
		}

		i++
	}

	_, err = txl.client.CreateRecord(context.Background(), &api.CreateRecordRequest{
		SourceOrganization: rec.SrcOrganization,
		DestOrganization:   rec.DestOrganization,
		ServiceName:        rec.ServiceName,
		TransactionID:      rec.TransactionID,
		Delegator:          rec.Delegator,
		OrderReference:     rec.OrderReference,
		Data:               string(dataJSON),
		Direction:          txl.direction,
		DataSubjects:       dataSubjects,
	})
	if err != nil {
		return err
	}

	return nil
}

func (txl *APITransactionLogger) Close() error {
	txl.cancelFunc()
	return txl.connection.Close()
}
