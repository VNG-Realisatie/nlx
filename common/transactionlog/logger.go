// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

import (
	"database/sql"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// TransactionLogger abstracts the writing of transactionlogs.
type TransactionLogger interface {
	AddRecord(rec *Record) error
}

// DiscardTransactionLogger discards records it gets
type DiscardTransactionLogger struct{}

// NewDiscardTransactionLogger creates a new TransactionLogger that discards all records.
func NewDiscardTransactionLogger() TransactionLogger {
	return &DiscardTransactionLogger{}
}

// AddRecord implements TransactionLogger.AddRecord and discards any given record.
func (txl *DiscardTransactionLogger) AddRecord(rec *Record) error {
	return nil
}

// PostgresTransactionLogger helps inway and outway to write transaction logs to a data store.
type PostgresTransactionLogger struct {
	zapLogger                *zap.Logger
	logdb                    *sqlx.DB
	stmtInsertTransactionLog *sqlx.NamedStmt
	stmtInsertDataSubject    *sqlx.Stmt
}

// NewPostgresTransactionLogger prepares a new TransactionLogger.
func NewPostgresTransactionLogger(zapLogger *zap.Logger, logdb *sqlx.DB, direction Direction) (TransactionLogger, error) {
	txl := &PostgresTransactionLogger{
		logdb: logdb,
	}

	switch direction {
	case DirectionIn:
	case DirectionOut:
	default:
		return nil, errors.New("invalid direction value")
	}

	var err error
	txl.stmtInsertDataSubject, err = logdb.Preparex(`
		INSERT INTO transactionlog.datasubjects (
			record_id,
			key,
			value
		) VALUES (
			$1,
			$2,
			$3
		)
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtInsertDataSubject")
	}

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
		RETURNING id
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

	dbtx, err := txl.logdb.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to start db transaction for log insertion")
	}
	defer func() {
		// always try a rollback when returning from function
		errRollback := dbtx.Rollback()
		if errRollback == sql.ErrTxDone {
			return // tx was already comitted, all is fine
		}
		if errRollback != nil {
			txl.zapLogger.Error("error rolling back transaction", zap.Error(errRollback))
		}
	}()

	var recordID uint64
	err = dbtx.NamedStmt(txl.stmtInsertTransactionLog).QueryRowx(insertRecord).Scan(&recordID)
	if err != nil {
		return errors.Wrap(err, "failed to insert log into data store")
	}
	for key, value := range rec.DataSubjects {
		_, err = dbtx.Stmtx(txl.stmtInsertDataSubject).Exec(recordID, key, value)
		if err != nil {
			return errors.Wrap(err, "failed to insert datasubject into data store")
		}
	}

	err = dbtx.Commit()
	if err != nil {
		return errors.Wrap(err, "failed to commit database transaction")
	}

	return nil
}
