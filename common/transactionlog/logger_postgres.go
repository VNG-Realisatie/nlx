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

type PostgresTransactionLogger struct {
	zapLogger                *zap.Logger
	logdb                    *sqlx.DB
	stmtInsertTransactionLog *sqlx.NamedStmt
	stmtInsertDataSubject    *sqlx.Stmt
	direction                Direction
}

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

	txl.direction = direction

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
			delegator,
			order_reference,
			data
		) VALUES (
			:direction,
			:src_organization,
			:dest_organization,
			:service_name,
			:logrecord_id,
			:delegator,
			:order_reference,
			:data_json
		)
		RETURNING id
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtInsertTransactionLog")
	}

	return txl, nil
}

func (txl *PostgresTransactionLogger) AddRecord(rec *Record) error {
	dataJSON, err := json.Marshal(rec.Data)
	if err != nil {
		return errors.Wrap(err, "failed to convert data to json")
	}

	insertRecord := struct {
		*Record
		Direction string
		DataJSON  types.JSONText
	}{
		Record:    rec,
		Direction: string(txl.direction),
		DataJSON:  types.JSONText(dataJSON),
	}

	dbtx, err := txl.logdb.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to start db transaction for log insertion")
	}

	defer func() {
		// always try a rollback when returning from function
		errRollback := dbtx.Rollback()
		if errRollback == sql.ErrTxDone {
			return // tx was already committed, all is fine
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

func (txl *PostgresTransactionLogger) Close() error {
	return txl.logdb.Close()
}
