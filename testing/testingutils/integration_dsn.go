// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package testingutils

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

func CreateTestDatabase(dsn, databaseName string) (testDBDsn string, err error) {
	dbCreator, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return "", fmt.Errorf("could not open connection to postgres in CreateTestDatabase: %s", err)
	}

	// We drop the database with the given name if it exists, to ensure we are using a clean database:
	// f.e. when database migrations have been edited during development, this is applied on the new database
	_, err = dbCreator.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", databaseName))
	if err != nil {
		return "", fmt.Errorf("could not drop database '%s': %s", databaseName, err)
	}

	_, err = dbCreator.Exec(fmt.Sprintf("CREATE DATABASE %s;", databaseName))
	if err != nil {
		return "", fmt.Errorf("could not create database '%s': %s", databaseName, err)
	}

	// Parse DNS, edit the path to the new database
	u, err := url.Parse(dsn)
	if err != nil {
		return "", fmt.Errorf("could not parse dsn '%s': %s", dsn, err)
	}

	// The u.Path needs to have a `/` prefix
	u.Path = fmt.Sprintf("/%s", databaseName)

	// We don't open the database connection here, because txdb can be used instead of postgres as driver
	testDBDsn = u.String()

	return testDBDsn, nil
}

func AddQueryParamToAddress(address, key, value string) string {
	u, _ := url.Parse(address)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add(key, value)
	u.RawQuery = q.Encode()

	return u.String()
}
