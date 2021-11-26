// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package pgadapter

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"go.nlx.io/nlx/directory-api/domain"
)

func (r *PostgreSQLRepository) ListParticipants(ctx context.Context) ([]*domain.Participant, error) {
	rows, err := r.selectParticipantsStmt.Queryx()
	if err != nil {
		return nil, fmt.Errorf("failed to execute stmtSelectParticipants: %v", err)
	}

	type dbParticipant struct {
		SerialNumber string    `db:"serial_number"`
		Name         string    `db:"name"`
		CreatedAt    time.Time `db:"created_at"`
		Inways       uint      `db:"inways"`
		Outways      uint      `db:"outways"`
		Services     uint      `db:"services"`
	}

	var queryResult []*dbParticipant

	err = sqlx.StructScan(rows, &queryResult)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Participant, len(queryResult))

	for i, o := range queryResult {
		org, err := domain.NewOrganization(o.Name, o.SerialNumber)
		if err != nil {
			return nil, err
		}

		p, err := domain.NewParticipant(&domain.NewParticipantArgs{
			Organization: org,
			Statistics: &domain.NewParticipantStatisticsArgs{
				Inways:   o.Inways,
				Outways:  o.Outways,
				Services: o.Services,
			},
			CreatedAt: o.CreatedAt,
		})
		if err != nil {
			return nil, err
		}

		result[i] = p
	}

	return result, nil
}

func prepareSelectParticipantsStatement(db *sqlx.DB) (*sqlx.Stmt, error) {
	listParticipantsStatement, err := db.Preparex(`
		SELECT
			serial_number,
			name,
			created_at,
			(select count(id) FROM directory.inways as i where i.organization_id = o.id) as inways,
			(select count(id) FROM directory.outways as ow where ow.organization_id = o.id) as outways,
			(select count(id) FROM directory.services as s where s.organization_id = o.id) as services
		FROM directory.organizations as o
	`)
	if err != nil {
		return nil, err
	}

	return listParticipantsStatement, nil
}
