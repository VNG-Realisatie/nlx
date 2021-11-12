package database

import (
	"context"
)

type InwayService struct {
	InwayID   uint
	ServiceID uint
}

func (i *InwayService) TableName() string {
	return "nlx_management.inways_services"
}

func (db *PostgresConfigDatabase) DeleteServicesConnectedToInway(ctx context.Context, inway *Inway) error {
	return db.DB.
		WithContext(ctx).
		Where(&InwayService{InwayID: inway.ID}).
		Delete(&InwayService{}).Error
}
