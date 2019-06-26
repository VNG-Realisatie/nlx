// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package insightapi

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"go.nlx.io/nlx/insight-api/irma"
)

type InsightLogFetcher interface {
	GetLogRecords(rowsPerPage, page int, dataSubjectsByIrmaAttribute map[string][]string, claims *irma.VerificationResultClaims) (*GetLogRecordsResponse, error)
}

type InsightDatabase struct {
	db                             *sqlx.DB
	logger                         *zap.Logger
	stmtCreateMatchDataSubjects    *sqlx.Stmt
	rawStmtInsertMatchDataSubjects string
	rawStmtFetchLogs               string
	rawStmtGetRowCount             string
}

func NewInsightDatabase(logger *zap.Logger, db *sqlx.DB) (*InsightDatabase, error) {
	i := &InsightDatabase{
		db:     db,
		logger: logger,
	}
	var err error
	i.stmtCreateMatchDataSubjects, err = db.Preparex(`
		CREATE TEMPORARY TABLE matchDataSubjects(
			key varchar(100),
			value varchar(100)
		)
		ON COMMIT DROP
	`)
	if err != nil {
		return nil, err
	}

	i.rawStmtInsertMatchDataSubjects = `INSERT INTO matchDataSubjects (key, value) VALUES ($1, $2)`

	i.rawStmtFetchLogs = `
		WITH matchedRecords AS (
			SELECT DISTINCT record_id
			FROM transactionlog.datasubjects
				INNER JOIN matchDataSubjects
					ON datasubjects.key = matchDataSubjects.key
						AND datasubjects.value = matchDataSubjects.value
		)
		SELECT
			created,
			src_organization,
			dest_organization,
			service_name,
			logrecord_id,
			data AS data_json
		FROM transactionlog.records
			INNER JOIN matchedRecords
				ON records.id = matchedRecords.record_id
		ORDER BY created DESC
		LIMIT CASE WHEN $1>0 THEN $1 ELSE NULL END
		OFFSET $2
	`

	i.rawStmtGetRowCount = `
	WITH matchedRecords AS (
		SELECT DISTINCT record_id
		FROM transactionlog.datasubjects
			INNER JOIN matchDataSubjects
				ON datasubjects.key = matchDataSubjects.key
					AND datasubjects.value = matchDataSubjects.value
	)
	SELECT
		COUNT(*)
	FROM transactionlog.records
		INNER JOIN matchedRecords
			ON records.id = matchedRecords.record_id
	`

	return i, nil
}

func (i *InsightDatabase) GetLogRecords(rowsPerPage, page int, dataSubjectsByIrmaAttribute map[string][]string, claims *irma.VerificationResultClaims) (*GetLogRecordsResponse, error) {
	var tx *sqlx.Tx
	var err error
	for retry := 0; retry < 3; retry++ {
		tx, err = i.db.Beginx()
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}
	defer func() {
		errRollback := tx.Rollback()
		if errRollback == sql.ErrTxDone {
			return // tx was already committed
		}
		if errRollback != nil {
			i.logger.Error("error rolling back transaction", zap.Error(errRollback))
		}
	}()

	err = i.matchDataSubjects(tx, dataSubjectsByIrmaAttribute, claims)
	if err != nil {
		return nil, err
	}

	out, err := i.fetchRecords(tx, page, rowsPerPage)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		i.logger.Warn("failed to commit transaction after succesful select", zap.Error(err))
	}

	return out, nil
}

func (i *InsightDatabase) matchDataSubjects(tx *sqlx.Tx, dataSubjectsByIrmaAttribute map[string][]string, claims *irma.VerificationResultClaims) error {
	// TODO: #345 investigate possible other solutions or optimizations for this crazy exercise
	_, err := tx.Stmtx(i.stmtCreateMatchDataSubjects).Exec()
	if err != nil {
		return fmt.Errorf("error executing stmtCreateMatchDataSubjects: %s", err)
	}

	stmtInsertMatchDataSubjects, err := tx.Preparex(i.rawStmtInsertMatchDataSubjects)
	if err != nil {
		return fmt.Errorf("error perparing stmtInsertMatchDataSubjects: %s", err)
	}
	defer stmtInsertMatchDataSubjects.Close()

	for irmaAttributeKey, irmaAttributeValue := range claims.Attributes {
		for _, dataSubjectKey := range dataSubjectsByIrmaAttribute[irmaAttributeKey] {
			_, err = stmtInsertMatchDataSubjects.Exec(dataSubjectKey, irmaAttributeValue)
			if err != nil {
				return fmt.Errorf("error executing stmtInsertMatchDataSubjects: %s", err)
			}
		}
	}

	return nil
}

func (i *InsightDatabase) fetchRecords(tx *sqlx.Tx, page, rowsPerPage int) (*GetLogRecordsResponse, error) {
	stmtFetchLogs, err := tx.Preparex(i.rawStmtFetchLogs)
	if err != nil {
		return nil, fmt.Errorf("error preparing stmtFetchLogs: %s", err)
	}
	defer stmtFetchLogs.Close()
	res, err := stmtFetchLogs.Queryx(rowsPerPage, page)
	if err != nil {
		return nil, fmt.Errorf("error executing stmtFetchLogs: %s", err)
	}

	var out = &GetLogRecordsResponse{
		Records:     make([]*Record, 0),
		RowsPerPage: rowsPerPage,
		Page:        page,
	}

	for res.Next() {
		rec := &Record{}
		err = res.StructScan(rec)
		if err != nil {
			return nil, fmt.Errorf("error preforming struct scan: %s", err)
		}
		err = rec.DataJSON.Unmarshal(&rec.Record.Data)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling record data: %s", err)
		}
		out.Records = append(out.Records, rec)
	}

	stmtGetRowCount, err := tx.Preparex(i.rawStmtGetRowCount)
	if err != nil {
		return nil, fmt.Errorf("error preparing stmtGetRowCount: %s", err)
	}

	resRowCount := stmtGetRowCount.QueryRowx()
	err = resRowCount.Scan(&out.RowCount)
	if err != nil {
		return nil, fmt.Errorf("error scanning rowcount: %s", err)
	}

	return out, nil
}
