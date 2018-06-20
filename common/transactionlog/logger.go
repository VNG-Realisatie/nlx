package transactionlog

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
)

type TransactionLogger interface {
	AddRecord(rec *Record) error
}

// DiscardTransactionLogger discards records it gets
type DiscardTransactionLogger struct{}

func NewDiscardTransactionLogger() TransactionLogger {
	return &DiscardTransactionLogger{}
}

func (txl *DiscardTransactionLogger) AddRecord(rec *Record) error {
	return nil
}

// PostgresTransactionLogger helps inway and outway to write transaction logs to a data store.
type PostgresTransactionLogger struct {
	logdb                    *sqlx.DB
	stmtInsertTransactionLog *sqlx.NamedStmt
}

// NewPostgressTransactionLogger prepares a new TransactionLogger.
func NewPostgresTransactionLogger(logdb *sqlx.DB, direction Direction) (TransactionLogger, error) {
	txl := &PostgresTransactionLogger{
		logdb: logdb,
	}

	var err error
	txl.stmtInsertTransactionLog, err = logdb.PrepareNamed(`
		INSERT INTO transactionlog.records (
			direction,
			src_organization,
			dest_organization,
			service_name,
			logrecord_id,
			data
		) VALUES (
			'` + string(direction) + `'::::transactionlog.direction,
			:src_organization,
			:dest_organization,
			:service_name,
			:logrecord_id,
			:data_json
		)
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtInsertTransactionLog")
	}

	return txl, nil
}

// AddRecord inserts a record into the datastore. Returns an error when failed.
func (txl *PostgresTransactionLogger) AddRecord(rec *Record) error {
	dataJSON, err := json.Marshal(rec.Data)
	if err != nil {
		return errors.Wrap(err, "failed to convert data to json")
	}
	insertRecord := struct {
		*Record
		DataJSON types.JSONText
	}{
		Record:   rec,
		DataJSON: types.JSONText(dataJSON),
	}

	_, err = txl.stmtInsertTransactionLog.Exec(insertRecord)
	if err != nil {
		return errors.Wrap(err, "failed to insert log into data store")
	}
	return nil
}
