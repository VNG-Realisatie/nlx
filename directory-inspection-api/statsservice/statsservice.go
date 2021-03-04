package statsservice

import (
	"context"
	"go.nlx.io/nlx/directory-inspection-api/stats"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StatsService struct {
	logger                      *zap.Logger
	stmtSelectVersionStatistics *sqlx.Stmt
}

func New(logger *zap.Logger, db *sqlx.DB) (*StatsService, error) {
	s := &StatsService{
		logger: logger.With(zap.String("handler", "stats-service")),
	}

	var err error

	// All the outways announcements for the last day (24 hours) are fetched and counted per version,
	// the inways are updated per organization so they have no time constraint.
	s.stmtSelectVersionStatistics, err = db.Preparex(`
		SELECT 'outway' AS type
		,      version
		,      COUNT(*) AS amount
		FROM   directory.outways
		WHERE  announced > now() - interval '1 days'
		GROUP BY version
		UNION
		SELECT 'inway' AS type
		,      version
		,      COUNT(*) AS amount
		FROM   directory.inways
		GROUP BY version
		ORDER BY type, version DESC
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSelectVersionStatistics")
	}

	return s, nil
}
func (s StatsService) ListVersionStatistics(context.Context, *stats.StatsRequest) (*stats.StatsResponse, error) {
	s.logger.Info("rpc request ListVersionStatistics")

	resp := &stats.StatsResponse{}

	err := s.stmtSelectVersionStatistics.Select(&resp.Versions)
	if err != nil {
		s.logger.Error("failed to select stats using stmtSelectVersionStatistics", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	return resp, nil
}
