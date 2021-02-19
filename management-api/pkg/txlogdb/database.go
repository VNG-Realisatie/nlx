package txlogdb

import (
	"context"
	"time"

	"gorm.io/gorm"

	"go.nlx.io/nlx/common/transactionlog"
)

type Record struct {
	Direction    transactionlog.Direction
	Source       string `gorm:"column:src_organization"`
	Destination  string `gorm:"column:dest_organization"`
	ServiceName  string
	RequestCount int
	CreatedAt    time.Time `gorm:"column:created"`
}

func (*Record) TableName() string {
	return "transactionlog.records"
}

type Filters struct {
	From        time.Time
	To          time.Time
	Source      string
	Destination string
	Direction   transactionlog.Direction
}

type TxlogDatabase interface {
	FilterRecords(ctx context.Context, filters *Filters) ([]Record, error)
}

type TxlogPostgresDatabase struct {
	DB *gorm.DB
}

func (db *TxlogPostgresDatabase) FilterRecords(ctx context.Context, filters *Filters) ([]Record, error) {
	query := db.DB.
		WithContext(ctx).
		Model(Record{})

	if !filters.From.IsZero() {
		query = query.Where("created >= ?", filters.From)
	}

	if !filters.To.IsZero() {
		query = query.Where("created <= ?", filters.To)
	}

	if len(filters.Source) > 0 {
		query = query.Where("src_organization = ?", filters.Source)
	}

	if len(filters.Destination) > 0 {
		query = query.Where("dest_organization = ?", filters.Destination)
	}

	if len(filters.Direction) > 0 {
		query = query.Where("direction = ?", filters.Direction)
	}

	records := []Record{}

	if err := query.
		Select(
			"src_organization",
			"dest_organization",
			"direction",
			"MIN(created)",
			"service_name",
			"COUNT(*) request_count",
		).
		Group("src_organization, dest_organization, service_name, direction, TO_CHAR(created, 'Mon-YYYY')").
		Find(&records).
		Error; err != nil {
		return nil, err
	}

	return records, nil
}
