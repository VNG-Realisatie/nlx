// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createAccessGrantStmt, err = db.PrepareContext(ctx, createAccessGrant); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccessGrant: %w", err)
	}
	if q.doesInwayExistByNameStmt, err = db.PrepareContext(ctx, doesInwayExistByName); err != nil {
		return nil, fmt.Errorf("error preparing query DoesInwayExistByName: %w", err)
	}
	if q.getAccessGrantStmt, err = db.PrepareContext(ctx, getAccessGrant); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccessGrant: %w", err)
	}
	if q.getSettingsStmt, err = db.PrepareContext(ctx, getSettings); err != nil {
		return nil, fmt.Errorf("error preparing query GetSettings: %w", err)
	}
	if q.updateSettingsStmt, err = db.PrepareContext(ctx, updateSettings); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSettings: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createAccessGrantStmt != nil {
		if cerr := q.createAccessGrantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccessGrantStmt: %w", cerr)
		}
	}
	if q.doesInwayExistByNameStmt != nil {
		if cerr := q.doesInwayExistByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing doesInwayExistByNameStmt: %w", cerr)
		}
	}
	if q.getAccessGrantStmt != nil {
		if cerr := q.getAccessGrantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccessGrantStmt: %w", cerr)
		}
	}
	if q.getSettingsStmt != nil {
		if cerr := q.getSettingsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSettingsStmt: %w", cerr)
		}
	}
	if q.updateSettingsStmt != nil {
		if cerr := q.updateSettingsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSettingsStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                       DBTX
	tx                       *sql.Tx
	createAccessGrantStmt    *sql.Stmt
	doesInwayExistByNameStmt *sql.Stmt
	getAccessGrantStmt       *sql.Stmt
	getSettingsStmt          *sql.Stmt
	updateSettingsStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                       tx,
		tx:                       tx,
		createAccessGrantStmt:    q.createAccessGrantStmt,
		doesInwayExistByNameStmt: q.doesInwayExistByNameStmt,
		getAccessGrantStmt:       q.getAccessGrantStmt,
		getSettingsStmt:          q.getSettingsStmt,
		updateSettingsStmt:       q.updateSettingsStmt,
	}
}
