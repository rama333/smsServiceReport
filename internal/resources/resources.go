package resources

import (
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"smsServiceReport/internal/config"
)

type R struct {
}

func New(logger *zap.SugaredLogger) (*R, error) {
	logger.Info("open bd")

	config.Config.DB, config.Config.DBerr = sqlx.Open("clickhouse", config.Config.DBURL)
	if config.Config.DBerr != nil {
		return nil, config.Config.DBerr
	}

	logger.Info("ping conn state error->", config.Config.DB.Ping())

	return &R{}, nil
}

func (r *R) Release() error {

	return config.Config.DB.Close()
}
