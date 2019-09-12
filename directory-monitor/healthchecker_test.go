// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

// build intergration

package monitor

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
)

var db *sql.DB

const database string = "TEST"

func TestMain(m *testing.M) {

	var err error

	pool, err := dockertest.NewPool("")

	log.Print("we are setting up main.")

	healthr, err := pool.BuildAndRun("health", "./Dockerfile", []string{"context", ".."})
	healthr.Expire(20)

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("postgres", "11.1", []string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=" + database})
	// make sure database docker is killed on time.
	resource.Expire(60)

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:secret@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), database))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	log.Print("done.")

	// run the tests!
	m.Run()
	// When you're done, kill and remove the container
	err = pool.Purge(resource)
}

func TestSomething(t *testing.T) {
	fmt.Println("we are running a test.")
}

func TestNew(t *testing.T) {
	fmt.Println("we are running another test.")
}
