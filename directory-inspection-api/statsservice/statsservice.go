package statsservice

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"go.nlx.io/nlx/directory-inspection-api/stats"
	"go.uber.org/zap"
)

type StatsService struct {
	logger          *zap.Logger
	stmtSelectStats *sqlx.Stmt
}

func New(logger *zap.Logger, db *sqlx.DB) (*StatsService, error) {
	s := &StatsService{
		logger: logger.With(zap.String("handler", "stats-service")),
	}

	var err error
	s.stmtSelectStats, err = db.Preparex(`
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
		return nil, errors.Wrap(err, "failed to prepare stmtSelectStats")
	}

	return s, nil
}
func (s StatsService) ListStats(context.Context, *stats.StatsRequest) (*stats.StatsResponse, error) {
	s.logger.Info("rpc request ListsStats")
	resp := &stats.StatsResponse{}

	err := s.stmtSelectStats.Select(&resp.Versions)
	if err != nil {
		s.logger.Error("failed to select stats using stmtSelectStats", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	return resp, nil
}
