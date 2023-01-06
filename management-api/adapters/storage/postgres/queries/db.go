// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

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
	if q.countReceivedOutgoingAccessRequestsForOutwayStmt, err = db.PrepareContext(ctx, countReceivedOutgoingAccessRequestsForOutway); err != nil {
		return nil, fmt.Errorf("error preparing query CountReceivedOutgoingAccessRequestsForOutway: %w", err)
	}
	if q.createAccessGrantStmt, err = db.PrepareContext(ctx, createAccessGrant); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccessGrant: %w", err)
	}
	if q.createAccessProofStmt, err = db.PrepareContext(ctx, createAccessProof); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAccessProof: %w", err)
	}
	if q.createAuditLogStmt, err = db.PrepareContext(ctx, createAuditLog); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAuditLog: %w", err)
	}
	if q.createAuditLogServiceStmt, err = db.PrepareContext(ctx, createAuditLogService); err != nil {
		return nil, fmt.Errorf("error preparing query CreateAuditLogService: %w", err)
	}
	if q.createIncomingAccessRequestStmt, err = db.PrepareContext(ctx, createIncomingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query CreateIncomingAccessRequest: %w", err)
	}
	if q.createOutgoingAccessRequestStmt, err = db.PrepareContext(ctx, createOutgoingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query CreateOutgoingAccessRequest: %w", err)
	}
	if q.createTermsOfServiceStmt, err = db.PrepareContext(ctx, createTermsOfService); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTermsOfService: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.createUserRolesStmt, err = db.PrepareContext(ctx, createUserRoles); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUserRoles: %w", err)
	}
	if q.deleteIncomingAccessRequestStmt, err = db.PrepareContext(ctx, deleteIncomingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteIncomingAccessRequest: %w", err)
	}
	if q.deleteOutgoingAccessRequestStmt, err = db.PrepareContext(ctx, deleteOutgoingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOutgoingAccessRequest: %w", err)
	}
	if q.deleteOutgoingAccessRequestsStmt, err = db.PrepareContext(ctx, deleteOutgoingAccessRequests); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteOutgoingAccessRequests: %w", err)
	}
	if q.doesInwayExistByNameStmt, err = db.PrepareContext(ctx, doesInwayExistByName); err != nil {
		return nil, fmt.Errorf("error preparing query DoesInwayExistByName: %w", err)
	}
	if q.getAccessGrantStmt, err = db.PrepareContext(ctx, getAccessGrant); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccessGrant: %w", err)
	}
	if q.getAccessGrantIDOfIncomingAccessRequestStmt, err = db.PrepareContext(ctx, getAccessGrantIDOfIncomingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccessGrantIDOfIncomingAccessRequest: %w", err)
	}
	if q.getAccessProofStmt, err = db.PrepareContext(ctx, getAccessProof); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccessProof: %w", err)
	}
	if q.getAccessProofByOutgoingAccessRequestStmt, err = db.PrepareContext(ctx, getAccessProofByOutgoingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query GetAccessProofByOutgoingAccessRequest: %w", err)
	}
	if q.getIncomingAccessRequestStmt, err = db.PrepareContext(ctx, getIncomingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query GetIncomingAccessRequest: %w", err)
	}
	if q.getIncomingAccessRequestsByServiceCountStmt, err = db.PrepareContext(ctx, getIncomingAccessRequestsByServiceCount); err != nil {
		return nil, fmt.Errorf("error preparing query GetIncomingAccessRequestsByServiceCount: %w", err)
	}
	if q.getInwayByNameStmt, err = db.PrepareContext(ctx, getInwayByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetInwayByName: %w", err)
	}
	if q.getLatestAccessGrantForServiceStmt, err = db.PrepareContext(ctx, getLatestAccessGrantForService); err != nil {
		return nil, fmt.Errorf("error preparing query GetLatestAccessGrantForService: %w", err)
	}
	if q.getLatestIncomingAccessRequestStmt, err = db.PrepareContext(ctx, getLatestIncomingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query GetLatestIncomingAccessRequest: %w", err)
	}
	if q.getLatestOutgoingAccessRequestStmt, err = db.PrepareContext(ctx, getLatestOutgoingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query GetLatestOutgoingAccessRequest: %w", err)
	}
	if q.getOutgoingAccessRequestStmt, err = db.PrepareContext(ctx, getOutgoingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query GetOutgoingAccessRequest: %w", err)
	}
	if q.getSettingsStmt, err = db.PrepareContext(ctx, getSettings); err != nil {
		return nil, fmt.Errorf("error preparing query GetSettings: %w", err)
	}
	if q.getUserByEmailStmt, err = db.PrepareContext(ctx, getUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByEmail: %w", err)
	}
	if q.listAccessGrantsForServiceStmt, err = db.PrepareContext(ctx, listAccessGrantsForService); err != nil {
		return nil, fmt.Errorf("error preparing query ListAccessGrantsForService: %w", err)
	}
	if q.listAllLatestOutgoingAccessRequestsStmt, err = db.PrepareContext(ctx, listAllLatestOutgoingAccessRequests); err != nil {
		return nil, fmt.Errorf("error preparing query ListAllLatestOutgoingAccessRequests: %w", err)
	}
	if q.listAuditLogServicesStmt, err = db.PrepareContext(ctx, listAuditLogServices); err != nil {
		return nil, fmt.Errorf("error preparing query ListAuditLogServices: %w", err)
	}
	if q.listAuditLogsStmt, err = db.PrepareContext(ctx, listAuditLogs); err != nil {
		return nil, fmt.Errorf("error preparing query ListAuditLogs: %w", err)
	}
	if q.listIncomingAccessRequestsStmt, err = db.PrepareContext(ctx, listIncomingAccessRequests); err != nil {
		return nil, fmt.Errorf("error preparing query ListIncomingAccessRequests: %w", err)
	}
	if q.listInwaysStmt, err = db.PrepareContext(ctx, listInways); err != nil {
		return nil, fmt.Errorf("error preparing query ListInways: %w", err)
	}
	if q.listInwaysForServiceStmt, err = db.PrepareContext(ctx, listInwaysForService); err != nil {
		return nil, fmt.Errorf("error preparing query ListInwaysForService: %w", err)
	}
	if q.listLatestOutgoingAccessRequestsStmt, err = db.PrepareContext(ctx, listLatestOutgoingAccessRequests); err != nil {
		return nil, fmt.Errorf("error preparing query ListLatestOutgoingAccessRequests: %w", err)
	}
	if q.listPermissionsStmt, err = db.PrepareContext(ctx, listPermissions); err != nil {
		return nil, fmt.Errorf("error preparing query ListPermissions: %w", err)
	}
	if q.listPermissionsForRoleStmt, err = db.PrepareContext(ctx, listPermissionsForRole); err != nil {
		return nil, fmt.Errorf("error preparing query ListPermissionsForRole: %w", err)
	}
	if q.listRolesForUserStmt, err = db.PrepareContext(ctx, listRolesForUser); err != nil {
		return nil, fmt.Errorf("error preparing query ListRolesForUser: %w", err)
	}
	if q.listServicesStmt, err = db.PrepareContext(ctx, listServices); err != nil {
		return nil, fmt.Errorf("error preparing query ListServices: %w", err)
	}
	if q.listServicesForInwayStmt, err = db.PrepareContext(ctx, listServicesForInway); err != nil {
		return nil, fmt.Errorf("error preparing query ListServicesForInway: %w", err)
	}
	if q.listTermsOfServiceStmt, err = db.PrepareContext(ctx, listTermsOfService); err != nil {
		return nil, fmt.Errorf("error preparing query ListTermsOfService: %w", err)
	}
	if q.removeInwayByNameStmt, err = db.PrepareContext(ctx, removeInwayByName); err != nil {
		return nil, fmt.Errorf("error preparing query RemoveInwayByName: %w", err)
	}
	if q.revokeAccessGrantStmt, err = db.PrepareContext(ctx, revokeAccessGrant); err != nil {
		return nil, fmt.Errorf("error preparing query RevokeAccessGrant: %w", err)
	}
	if q.revokeAccessProofStmt, err = db.PrepareContext(ctx, revokeAccessProof); err != nil {
		return nil, fmt.Errorf("error preparing query RevokeAccessProof: %w", err)
	}
	if q.setAuditLogAsSucceededStmt, err = db.PrepareContext(ctx, setAuditLogAsSucceeded); err != nil {
		return nil, fmt.Errorf("error preparing query SetAuditLogAsSucceeded: %w", err)
	}
	if q.terminateAccessGrantStmt, err = db.PrepareContext(ctx, terminateAccessGrant); err != nil {
		return nil, fmt.Errorf("error preparing query TerminateAccessGrant: %w", err)
	}
	if q.terminateAccessProofStmt, err = db.PrepareContext(ctx, terminateAccessProof); err != nil {
		return nil, fmt.Errorf("error preparing query TerminateAccessProof: %w", err)
	}
	if q.updateIncomingAccessRequestStmt, err = db.PrepareContext(ctx, updateIncomingAccessRequest); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateIncomingAccessRequest: %w", err)
	}
	if q.updateInwayStmt, err = db.PrepareContext(ctx, updateInway); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateInway: %w", err)
	}
	if q.updateOutgoingAccessRequestStateStmt, err = db.PrepareContext(ctx, updateOutgoingAccessRequestState); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateOutgoingAccessRequestState: %w", err)
	}
	if q.updateSettingsStmt, err = db.PrepareContext(ctx, updateSettings); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateSettings: %w", err)
	}
	if q.upsertInwayStmt, err = db.PrepareContext(ctx, upsertInway); err != nil {
		return nil, fmt.Errorf("error preparing query UpsertInway: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countReceivedOutgoingAccessRequestsForOutwayStmt != nil {
		if cerr := q.countReceivedOutgoingAccessRequestsForOutwayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countReceivedOutgoingAccessRequestsForOutwayStmt: %w", cerr)
		}
	}
	if q.createAccessGrantStmt != nil {
		if cerr := q.createAccessGrantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccessGrantStmt: %w", cerr)
		}
	}
	if q.createAccessProofStmt != nil {
		if cerr := q.createAccessProofStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAccessProofStmt: %w", cerr)
		}
	}
	if q.createAuditLogStmt != nil {
		if cerr := q.createAuditLogStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAuditLogStmt: %w", cerr)
		}
	}
	if q.createAuditLogServiceStmt != nil {
		if cerr := q.createAuditLogServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createAuditLogServiceStmt: %w", cerr)
		}
	}
	if q.createIncomingAccessRequestStmt != nil {
		if cerr := q.createIncomingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createIncomingAccessRequestStmt: %w", cerr)
		}
	}
	if q.createOutgoingAccessRequestStmt != nil {
		if cerr := q.createOutgoingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createOutgoingAccessRequestStmt: %w", cerr)
		}
	}
	if q.createTermsOfServiceStmt != nil {
		if cerr := q.createTermsOfServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTermsOfServiceStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.createUserRolesStmt != nil {
		if cerr := q.createUserRolesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserRolesStmt: %w", cerr)
		}
	}
	if q.deleteIncomingAccessRequestStmt != nil {
		if cerr := q.deleteIncomingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteIncomingAccessRequestStmt: %w", cerr)
		}
	}
	if q.deleteOutgoingAccessRequestStmt != nil {
		if cerr := q.deleteOutgoingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOutgoingAccessRequestStmt: %w", cerr)
		}
	}
	if q.deleteOutgoingAccessRequestsStmt != nil {
		if cerr := q.deleteOutgoingAccessRequestsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteOutgoingAccessRequestsStmt: %w", cerr)
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
	if q.getAccessGrantIDOfIncomingAccessRequestStmt != nil {
		if cerr := q.getAccessGrantIDOfIncomingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccessGrantIDOfIncomingAccessRequestStmt: %w", cerr)
		}
	}
	if q.getAccessProofStmt != nil {
		if cerr := q.getAccessProofStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccessProofStmt: %w", cerr)
		}
	}
	if q.getAccessProofByOutgoingAccessRequestStmt != nil {
		if cerr := q.getAccessProofByOutgoingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getAccessProofByOutgoingAccessRequestStmt: %w", cerr)
		}
	}
	if q.getIncomingAccessRequestStmt != nil {
		if cerr := q.getIncomingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getIncomingAccessRequestStmt: %w", cerr)
		}
	}
	if q.getIncomingAccessRequestsByServiceCountStmt != nil {
		if cerr := q.getIncomingAccessRequestsByServiceCountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getIncomingAccessRequestsByServiceCountStmt: %w", cerr)
		}
	}
	if q.getInwayByNameStmt != nil {
		if cerr := q.getInwayByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInwayByNameStmt: %w", cerr)
		}
	}
	if q.getLatestAccessGrantForServiceStmt != nil {
		if cerr := q.getLatestAccessGrantForServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLatestAccessGrantForServiceStmt: %w", cerr)
		}
	}
	if q.getLatestIncomingAccessRequestStmt != nil {
		if cerr := q.getLatestIncomingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLatestIncomingAccessRequestStmt: %w", cerr)
		}
	}
	if q.getLatestOutgoingAccessRequestStmt != nil {
		if cerr := q.getLatestOutgoingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getLatestOutgoingAccessRequestStmt: %w", cerr)
		}
	}
	if q.getOutgoingAccessRequestStmt != nil {
		if cerr := q.getOutgoingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getOutgoingAccessRequestStmt: %w", cerr)
		}
	}
	if q.getSettingsStmt != nil {
		if cerr := q.getSettingsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getSettingsStmt: %w", cerr)
		}
	}
	if q.getUserByEmailStmt != nil {
		if cerr := q.getUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByEmailStmt: %w", cerr)
		}
	}
	if q.listAccessGrantsForServiceStmt != nil {
		if cerr := q.listAccessGrantsForServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAccessGrantsForServiceStmt: %w", cerr)
		}
	}
	if q.listAllLatestOutgoingAccessRequestsStmt != nil {
		if cerr := q.listAllLatestOutgoingAccessRequestsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAllLatestOutgoingAccessRequestsStmt: %w", cerr)
		}
	}
	if q.listAuditLogServicesStmt != nil {
		if cerr := q.listAuditLogServicesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAuditLogServicesStmt: %w", cerr)
		}
	}
	if q.listAuditLogsStmt != nil {
		if cerr := q.listAuditLogsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listAuditLogsStmt: %w", cerr)
		}
	}
	if q.listIncomingAccessRequestsStmt != nil {
		if cerr := q.listIncomingAccessRequestsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listIncomingAccessRequestsStmt: %w", cerr)
		}
	}
	if q.listInwaysStmt != nil {
		if cerr := q.listInwaysStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listInwaysStmt: %w", cerr)
		}
	}
	if q.listInwaysForServiceStmt != nil {
		if cerr := q.listInwaysForServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listInwaysForServiceStmt: %w", cerr)
		}
	}
	if q.listLatestOutgoingAccessRequestsStmt != nil {
		if cerr := q.listLatestOutgoingAccessRequestsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listLatestOutgoingAccessRequestsStmt: %w", cerr)
		}
	}
	if q.listPermissionsStmt != nil {
		if cerr := q.listPermissionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPermissionsStmt: %w", cerr)
		}
	}
	if q.listPermissionsForRoleStmt != nil {
		if cerr := q.listPermissionsForRoleStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listPermissionsForRoleStmt: %w", cerr)
		}
	}
	if q.listRolesForUserStmt != nil {
		if cerr := q.listRolesForUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listRolesForUserStmt: %w", cerr)
		}
	}
	if q.listServicesStmt != nil {
		if cerr := q.listServicesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listServicesStmt: %w", cerr)
		}
	}
	if q.listServicesForInwayStmt != nil {
		if cerr := q.listServicesForInwayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listServicesForInwayStmt: %w", cerr)
		}
	}
	if q.listTermsOfServiceStmt != nil {
		if cerr := q.listTermsOfServiceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTermsOfServiceStmt: %w", cerr)
		}
	}
	if q.removeInwayByNameStmt != nil {
		if cerr := q.removeInwayByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing removeInwayByNameStmt: %w", cerr)
		}
	}
	if q.revokeAccessGrantStmt != nil {
		if cerr := q.revokeAccessGrantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing revokeAccessGrantStmt: %w", cerr)
		}
	}
	if q.revokeAccessProofStmt != nil {
		if cerr := q.revokeAccessProofStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing revokeAccessProofStmt: %w", cerr)
		}
	}
	if q.setAuditLogAsSucceededStmt != nil {
		if cerr := q.setAuditLogAsSucceededStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing setAuditLogAsSucceededStmt: %w", cerr)
		}
	}
	if q.terminateAccessGrantStmt != nil {
		if cerr := q.terminateAccessGrantStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing terminateAccessGrantStmt: %w", cerr)
		}
	}
	if q.terminateAccessProofStmt != nil {
		if cerr := q.terminateAccessProofStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing terminateAccessProofStmt: %w", cerr)
		}
	}
	if q.updateIncomingAccessRequestStmt != nil {
		if cerr := q.updateIncomingAccessRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateIncomingAccessRequestStmt: %w", cerr)
		}
	}
	if q.updateInwayStmt != nil {
		if cerr := q.updateInwayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateInwayStmt: %w", cerr)
		}
	}
	if q.updateOutgoingAccessRequestStateStmt != nil {
		if cerr := q.updateOutgoingAccessRequestStateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateOutgoingAccessRequestStateStmt: %w", cerr)
		}
	}
	if q.updateSettingsStmt != nil {
		if cerr := q.updateSettingsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateSettingsStmt: %w", cerr)
		}
	}
	if q.upsertInwayStmt != nil {
		if cerr := q.upsertInwayStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing upsertInwayStmt: %w", cerr)
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
	db                                               DBTX
	tx                                               *sql.Tx
	countReceivedOutgoingAccessRequestsForOutwayStmt *sql.Stmt
	createAccessGrantStmt                            *sql.Stmt
	createAccessProofStmt                            *sql.Stmt
	createAuditLogStmt                               *sql.Stmt
	createAuditLogServiceStmt                        *sql.Stmt
	createIncomingAccessRequestStmt                  *sql.Stmt
	createOutgoingAccessRequestStmt                  *sql.Stmt
	createTermsOfServiceStmt                         *sql.Stmt
	createUserStmt                                   *sql.Stmt
	createUserRolesStmt                              *sql.Stmt
	deleteIncomingAccessRequestStmt                  *sql.Stmt
	deleteOutgoingAccessRequestStmt                  *sql.Stmt
	deleteOutgoingAccessRequestsStmt                 *sql.Stmt
	doesInwayExistByNameStmt                         *sql.Stmt
	getAccessGrantStmt                               *sql.Stmt
	getAccessGrantIDOfIncomingAccessRequestStmt      *sql.Stmt
	getAccessProofStmt                               *sql.Stmt
	getAccessProofByOutgoingAccessRequestStmt        *sql.Stmt
	getIncomingAccessRequestStmt                     *sql.Stmt
	getIncomingAccessRequestsByServiceCountStmt      *sql.Stmt
	getInwayByNameStmt                               *sql.Stmt
	getLatestAccessGrantForServiceStmt               *sql.Stmt
	getLatestIncomingAccessRequestStmt               *sql.Stmt
	getLatestOutgoingAccessRequestStmt               *sql.Stmt
	getOutgoingAccessRequestStmt                     *sql.Stmt
	getSettingsStmt                                  *sql.Stmt
	getUserByEmailStmt                               *sql.Stmt
	listAccessGrantsForServiceStmt                   *sql.Stmt
	listAllLatestOutgoingAccessRequestsStmt          *sql.Stmt
	listAuditLogServicesStmt                         *sql.Stmt
	listAuditLogsStmt                                *sql.Stmt
	listIncomingAccessRequestsStmt                   *sql.Stmt
	listInwaysStmt                                   *sql.Stmt
	listInwaysForServiceStmt                         *sql.Stmt
	listLatestOutgoingAccessRequestsStmt             *sql.Stmt
	listPermissionsStmt                              *sql.Stmt
	listPermissionsForRoleStmt                       *sql.Stmt
	listRolesForUserStmt                             *sql.Stmt
	listServicesStmt                                 *sql.Stmt
	listServicesForInwayStmt                         *sql.Stmt
	listTermsOfServiceStmt                           *sql.Stmt
	removeInwayByNameStmt                            *sql.Stmt
	revokeAccessGrantStmt                            *sql.Stmt
	revokeAccessProofStmt                            *sql.Stmt
	setAuditLogAsSucceededStmt                       *sql.Stmt
	terminateAccessGrantStmt                         *sql.Stmt
	terminateAccessProofStmt                         *sql.Stmt
	updateIncomingAccessRequestStmt                  *sql.Stmt
	updateInwayStmt                                  *sql.Stmt
	updateOutgoingAccessRequestStateStmt             *sql.Stmt
	updateSettingsStmt                               *sql.Stmt
	upsertInwayStmt                                  *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
		tx: tx,
		countReceivedOutgoingAccessRequestsForOutwayStmt: q.countReceivedOutgoingAccessRequestsForOutwayStmt,
		createAccessGrantStmt:                            q.createAccessGrantStmt,
		createAccessProofStmt:                            q.createAccessProofStmt,
		createAuditLogStmt:                               q.createAuditLogStmt,
		createAuditLogServiceStmt:                        q.createAuditLogServiceStmt,
		createIncomingAccessRequestStmt:                  q.createIncomingAccessRequestStmt,
		createOutgoingAccessRequestStmt:                  q.createOutgoingAccessRequestStmt,
		createTermsOfServiceStmt:                         q.createTermsOfServiceStmt,
		createUserStmt:                                   q.createUserStmt,
		createUserRolesStmt:                              q.createUserRolesStmt,
		deleteIncomingAccessRequestStmt:                  q.deleteIncomingAccessRequestStmt,
		deleteOutgoingAccessRequestStmt:                  q.deleteOutgoingAccessRequestStmt,
		deleteOutgoingAccessRequestsStmt:                 q.deleteOutgoingAccessRequestsStmt,
		doesInwayExistByNameStmt:                         q.doesInwayExistByNameStmt,
		getAccessGrantStmt:                               q.getAccessGrantStmt,
		getAccessGrantIDOfIncomingAccessRequestStmt:      q.getAccessGrantIDOfIncomingAccessRequestStmt,
		getAccessProofStmt:                               q.getAccessProofStmt,
		getAccessProofByOutgoingAccessRequestStmt:        q.getAccessProofByOutgoingAccessRequestStmt,
		getIncomingAccessRequestStmt:                     q.getIncomingAccessRequestStmt,
		getIncomingAccessRequestsByServiceCountStmt:      q.getIncomingAccessRequestsByServiceCountStmt,
		getInwayByNameStmt:                               q.getInwayByNameStmt,
		getLatestAccessGrantForServiceStmt:               q.getLatestAccessGrantForServiceStmt,
		getLatestIncomingAccessRequestStmt:               q.getLatestIncomingAccessRequestStmt,
		getLatestOutgoingAccessRequestStmt:               q.getLatestOutgoingAccessRequestStmt,
		getOutgoingAccessRequestStmt:                     q.getOutgoingAccessRequestStmt,
		getSettingsStmt:                                  q.getSettingsStmt,
		getUserByEmailStmt:                               q.getUserByEmailStmt,
		listAccessGrantsForServiceStmt:                   q.listAccessGrantsForServiceStmt,
		listAllLatestOutgoingAccessRequestsStmt:          q.listAllLatestOutgoingAccessRequestsStmt,
		listAuditLogServicesStmt:                         q.listAuditLogServicesStmt,
		listAuditLogsStmt:                                q.listAuditLogsStmt,
		listIncomingAccessRequestsStmt:                   q.listIncomingAccessRequestsStmt,
		listInwaysStmt:                                   q.listInwaysStmt,
		listInwaysForServiceStmt:                         q.listInwaysForServiceStmt,
		listLatestOutgoingAccessRequestsStmt:             q.listLatestOutgoingAccessRequestsStmt,
		listPermissionsStmt:                              q.listPermissionsStmt,
		listPermissionsForRoleStmt:                       q.listPermissionsForRoleStmt,
		listRolesForUserStmt:                             q.listRolesForUserStmt,
		listServicesStmt:                                 q.listServicesStmt,
		listServicesForInwayStmt:                         q.listServicesForInwayStmt,
		listTermsOfServiceStmt:                           q.listTermsOfServiceStmt,
		removeInwayByNameStmt:                            q.removeInwayByNameStmt,
		revokeAccessGrantStmt:                            q.revokeAccessGrantStmt,
		revokeAccessProofStmt:                            q.revokeAccessProofStmt,
		setAuditLogAsSucceededStmt:                       q.setAuditLogAsSucceededStmt,
		terminateAccessGrantStmt:                         q.terminateAccessGrantStmt,
		terminateAccessProofStmt:                         q.terminateAccessProofStmt,
		updateIncomingAccessRequestStmt:                  q.updateIncomingAccessRequestStmt,
		updateInwayStmt:                                  q.updateInwayStmt,
		updateOutgoingAccessRequestStateStmt:             q.updateOutgoingAccessRequestStateStmt,
		updateSettingsStmt:                               q.updateSettingsStmt,
		upsertInwayStmt:                                  q.upsertInwayStmt,
	}
}
