package testingutils

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

var createIfNotExistsBase = `
DO
$do$
BEGIN
   IF EXISTS (SELECT FROM pg_database WHERE datname = '%s') THEN
      RAISE NOTICE 'Database already exists';
   ELSE
      CREATE EXTENSION IF NOT EXISTS dblink;
      PERFORM dblink_exec('dbname=' || current_database()
                        , 'CREATE DATABASE %s');
   END IF;
END
$do$;`

func CreateTestDatabase(dsn, databaseName string) (testDBDsn string, err error) {
	dbCreator, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return "", fmt.Errorf("could not open connection to postgres in CreateTestDatabase for creating the testing database: %s", err)
	}

	// Create new database with given name if it does not exists
	_, err = dbCreator.Exec(fmt.Sprintf(createIfNotExistsBase, databaseName, databaseName))
	if err != nil {
		return "", err
	}

	// Parse DNS, edit the path to the new database
	u, err := url.Parse(dsn)
	if err != nil {
		return "", err
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
