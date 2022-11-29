// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

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
	if q.countExpiredSessionsStmt, err = db.PrepareContext(ctx, countExpiredSessions); err != nil {
		return nil, fmt.Errorf("error preparing query CountExpiredSessions: %w", err)
	}
	if q.createSessionStmt, err = db.PrepareContext(ctx, createSession); err != nil {
		return nil, fmt.Errorf("error preparing query CreateSession: %w", err)
	}
	if q.deleteExpiredSessionsStmt, err = db.PrepareContext(ctx, deleteExpiredSessions); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteExpiredSessions: %w", err)
	}
	if q.deleteSessionStmt, err = db.PrepareContext(ctx, deleteSession); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteSession: %w", err)
	}
	if q.getSessionStmt, err = db.PrepareContext(ctx, getSession); err != nil {
		return nil, fmt.Errorf("error preparing query GetSession: %w", err)
	}
	if q.updateSessionStmt, err = db.PrepareContext(ctx, updateSession); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSession: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countExpiredSessionsStmt != nil {
		if cerr := q.countExpiredSessionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countExpiredSessionsStmt: %w", cerr)
		}
	}
	if q.createSessionStmt != nil {
		if cerr := q.createSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createSessionStmt: %w", cerr)
		}
	}
	if q.deleteExpiredSessionsStmt != nil {
		if cerr := q.deleteExpiredSessionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteExpiredSessionsStmt: %w", cerr)
		}
	}
	if q.deleteSessionStmt != nil {
		if cerr := q.deleteSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteSessionStmt: %w", cerr)
		}
	}
	if q.getSessionStmt != nil {
		if cerr := q.getSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSessionStmt: %w", cerr)
		}
	}
	if q.updateSessionStmt != nil {
		if cerr := q.updateSessionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSessionStmt: %w", cerr)
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
	db                        DBTX
	tx                        *sql.Tx
	countExpiredSessionsStmt  *sql.Stmt
	createSessionStmt         *sql.Stmt
	deleteExpiredSessionsStmt *sql.Stmt
	deleteSessionStmt         *sql.Stmt
	getSessionStmt            *sql.Stmt
	updateSessionStmt         *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                        tx,
		tx:                        tx,
		countExpiredSessionsStmt:  q.countExpiredSessionsStmt,
		createSessionStmt:         q.createSessionStmt,
		deleteExpiredSessionsStmt: q.deleteExpiredSessionsStmt,
		deleteSessionStmt:         q.deleteSessionStmt,
		getSessionStmt:            q.getSessionStmt,
		updateSessionStmt:         q.updateSessionStmt,
	}
}