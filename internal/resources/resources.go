package resources

import (
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type R struct {
	Config Config
	Conn   *sqlx.DB
}

type Config struct {
	DiagPort    int    `envconfig:"DIAG_PORT" default:"8081" required:"true"`
	RESTAPIPort int    `envconfig:"PORT" default:"8080" required:"true"`
	DBURL       string `envconfig:"DB_URL" default:"tcp://dockerhost:9000?debug=true" required:"true"`
}

func New(logger *zap.SugaredLogger) (*R, error) {
	conf := Config{}

	logger.Info("starrart...")

	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}

	//conn, err := sql.Open("pgx", conf.DBURL)
	conn, err := sqlx.Open("clickhouse", "tcp://192.168.114.145:9000?debug=true")
	if err != nil {
		return nil, err
	}

	logger.Info("starrart...", conn.Ping())

	//db := reform.NewDB(conn, postgresql.Dialect, reform.NewPrintfLogger(logger.Infof))

	return &R{
		Config: conf,
		Conn:   conn,
	}, nil
}

func (r *R) Release() error {
	return r.Conn.Close()
}
