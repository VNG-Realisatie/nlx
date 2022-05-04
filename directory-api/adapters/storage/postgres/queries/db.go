// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

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
	if q.clearOrganizationInwayStmt, err = db.PrepareContext(ctx, clearOrganizationInway); err != nil {
		return nil, fmt.Errorf("error preparing query ClearOrganizationInway: %w", err)
	}
	if q.getInwayStmt, err = db.PrepareContext(ctx, getInway); err != nil {
		return nil, fmt.Errorf("error preparing query GetInway: %w", err)
	}
	if q.getServiceStmt, err = db.PrepareContext(ctx, getService); err != nil {
		return nil, fmt.Errorf("error preparing query GetService: %w", err)
	}
	if q.selectInwayByAddressStmt, err = db.PrepareContext(ctx, selectInwayByAddress); err != nil {
		return nil, fmt.Errorf("error preparing query SelectInwayByAddress: %w", err)
	}
	if q.selectOrganizationInwayAddressStmt, err = db.PrepareContext(ctx, selectOrganizationInwayAddress); err != nil {
		return nil, fmt.Errorf("error preparing query SelectOrganizationInwayAddress: %w", err)
	}
	if q.selectOrganizationInwayManagementAPIProxyAddressStmt, err = db.PrepareContext(ctx, selectOrganizationInwayManagementAPIProxyAddress); err != nil {
		return nil, fmt.Errorf("error preparing query SelectOrganizationInwayManagementAPIProxyAddress: %w", err)
	}
	if q.selectOrganizationsStmt, err = db.PrepareContext(ctx, selectOrganizations); err != nil {
		return nil, fmt.Errorf("error preparing query SelectOrganizations: %w", err)
	}
	if q.selectVersionStatisticsStmt, err = db.PrepareContext(ctx, selectVersionStatistics); err != nil {
		return nil, fmt.Errorf("error preparing query SelectVersionStatistics: %w", err)
	}
	if q.setOrganizationEmailStmt, err = db.PrepareContext(ctx, setOrganizationEmail); err != nil {
		return nil, fmt.Errorf("error preparing query SetOrganizationEmail: %w", err)
	}
	if q.setOrganizationInwayStmt, err = db.PrepareContext(ctx, setOrganizationInway); err != nil {
		return nil, fmt.Errorf("error preparing query SetOrganizationInway: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.clearOrganizationInwayStmt != nil {
		if cerr := q.clearOrganizationInwayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing clearOrganizationInwayStmt: %w", cerr)
		}
	}
	if q.getInwayStmt != nil {
		if cerr := q.getInwayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInwayStmt: %w", cerr)
		}
	}
	if q.getServiceStmt != nil {
		if cerr := q.getServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getServiceStmt: %w", cerr)
		}
	}
	if q.selectInwayByAddressStmt != nil {
		if cerr := q.selectInwayByAddressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectInwayByAddressStmt: %w", cerr)
		}
	}
	if q.selectOrganizationInwayAddressStmt != nil {
		if cerr := q.selectOrganizationInwayAddressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectOrganizationInwayAddressStmt: %w", cerr)
		}
	}
	if q.selectOrganizationInwayManagementAPIProxyAddressStmt != nil {
		if cerr := q.selectOrganizationInwayManagementAPIProxyAddressStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectOrganizationInwayManagementAPIProxyAddressStmt: %w", cerr)
		}
	}
	if q.selectOrganizationsStmt != nil {
		if cerr := q.selectOrganizationsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectOrganizationsStmt: %w", cerr)
		}
	}
	if q.selectVersionStatisticsStmt != nil {
		if cerr := q.selectVersionStatisticsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing selectVersionStatisticsStmt: %w", cerr)
		}
	}
	if q.setOrganizationEmailStmt != nil {
		if cerr := q.setOrganizationEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setOrganizationEmailStmt: %w", cerr)
		}
	}
	if q.setOrganizationInwayStmt != nil {
		if cerr := q.setOrganizationInwayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setOrganizationInwayStmt: %w", cerr)
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
	db                                                   DBTX
	tx                                                   *sql.Tx
	clearOrganizationInwayStmt                           *sql.Stmt
	getInwayStmt                                         *sql.Stmt
	getServiceStmt                                       *sql.Stmt
	selectInwayByAddressStmt                             *sql.Stmt
	selectOrganizationInwayAddressStmt                   *sql.Stmt
	selectOrganizationInwayManagementAPIProxyAddressStmt *sql.Stmt
	selectOrganizationsStmt                              *sql.Stmt
	selectVersionStatisticsStmt                          *sql.Stmt
	setOrganizationEmailStmt                             *sql.Stmt
	setOrganizationInwayStmt                             *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                 tx,
		tx:                                 tx,
		clearOrganizationInwayStmt:         q.clearOrganizationInwayStmt,
		getInwayStmt:                       q.getInwayStmt,
		getServiceStmt:                     q.getServiceStmt,
		selectInwayByAddressStmt:           q.selectInwayByAddressStmt,
		selectOrganizationInwayAddressStmt: q.selectOrganizationInwayAddressStmt,
		selectOrganizationInwayManagementAPIProxyAddressStmt: q.selectOrganizationInwayManagementAPIProxyAddressStmt,
		selectOrganizationsStmt:                              q.selectOrganizationsStmt,
		selectVersionStatisticsStmt:                          q.selectVersionStatisticsStmt,
		setOrganizationEmailStmt:                             q.setOrganizationEmailStmt,
		setOrganizationInwayStmt:                             q.setOrganizationInwayStmt,
	}
}
