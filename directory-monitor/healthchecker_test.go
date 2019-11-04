// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

// +build integration

package monitor

import (
	"crypto/tls"
	"database/sql"
	"log"
	"path/filepath"
	"testing"

	"go.nlx.io/nlx/directory-monitor/health"
	"go.nlx.io/nlx/directory-registration-api/registrationservice"

	common_db "go.nlx.io/nlx/common/db"
	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-db/dbversion"

	"github.com/fgrosse/zaptest"
	"github.com/stretchr/testify/assert"
	"go.nlx.io/nlx/common/orgtls"
)

var db *sql.DB

const PostgresDSN string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func TestMain(m *testing.M) {

	var err error

	log.Print("we are setting up main.")

	db, err = sql.Open("postgres", PostgresDSN)
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping to database: %s", err)
	}

	m.Run()
}

func getAvailabilities(t *testing.T) []*availability {
	h := RunningHealthChecker
	newAvailabilities := []*availability{}
	err := h.stmtSelectAvailabilities.Select(&newAvailabilities)
	if err != nil {
		t.Fatal(err)
	}
	return newAvailabilities
}

func TestSetupHealthMonitor(t *testing.T) {
	logger := zaptest.Logger(t)
	proc := process.NewProcess(logger)
	db, _ := InitDatabase(PostgresDSN)
	proc.CloseGracefully(db.Close)
	common_db.WaitForLatestDBVersion(logger, db.DB, dbversion.LatestDirectoryDBVersion)

	tlsOptions := orgtls.TLSOptions{
		NLXRootCert: filepath.Join("..", "testing", "root.crt"),
		OrgCertFile: filepath.Join("..", "testing", "org-nlx-test.crt"),
		OrgKeyFile:  filepath.Join("..", "testing", "org-nlx-test.key"),
	}

	caCertPool, _ := orgtls.LoadRootCert(tlsOptions.NLXRootCert)

	certKeyPair, _ := tls.LoadX509KeyPair(
		tlsOptions.OrgCertFile, tlsOptions.OrgKeyFile)

	// TODO make timers really short so we can check all
	// functionality.

	var err error

	go func() {
		err = RunHealthChecker(
			proc, logger, db,
			PostgresDSN,
			caCertPool,
			&certKeyPair,
			10)
		if err != nil {
			t.Fatal(err)
		}
	}()

	r, _ := registrationservice.NewRegisterInwayHandler(
		db,
		logger,
		caCertPool,
		&certKeyPair,
	)

	r.InsertInway("healttest1org", "tada1", "http://localhost:12345")
	r.InsertInway("healttest2org", "tada2", "http://localhost:12345")
	r.InsertInway("healttest3org", "tada3", "http://localhost:12345")

	newAvailabilities := getAvailabilities(t)
	assert.Equal(t, len(newAvailabilities), 3)

	status := &health.Status{}
	status.Healthy = true
	status.Version = "testversion"

	av := *newAvailabilities[2]
	RunningHealthChecker.updateAvailabilityHealth(av, status.Healthy)
	RunningHealthChecker.updateInwayVersion(av, status.Version)
	newAvailabilities = getAvailabilities(t)
	av = *newAvailabilities[2]
	// TODO check that status is in database
	vs := []string{}

	err = db.Select(&vs, `
		SELECT version from directory.inways
		where version is not null
	`)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(vs), 1)
	assert.Equal(t, vs[0], "testversion")
}
