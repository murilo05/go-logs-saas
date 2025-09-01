package postgres

import (
	"context"
	"errors"
	"fmt"
	"go-log-saas/internal/adapter/config"
	"log"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PG struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	zap.SugaredLogger
}

func Config(config *config.DB) *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	var DATABASE_URL string = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	pgxCfg, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	pgxCfg.MaxConns = defaultMaxConns
	pgxCfg.MinConns = defaultMinConns
	pgxCfg.MaxConnLifetime = defaultMaxConnLifetime
	pgxCfg.MaxConnIdleTime = defaultMaxConnIdleTime
	pgxCfg.HealthCheckPeriod = defaultHealthCheckPeriod
	pgxCfg.ConnConfig.ConnectTimeout = defaultConnectTimeout

	pgxCfg.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	pgxCfg.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	pgxCfg.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return pgxCfg
}

func NewDatabase(ctx context.Context, configDB *config.DB, logger *zap.SugaredLogger) *PG {
	pg, err := NewPostgres(ctx, configDB, logger)
	if err != nil {
		logger.Error("Error initializing database connection: %s", err)
		os.Exit(1)
	}

	logger.Info("Successfully connected to the database", "db", configDB.Connection)

	return pg

}

func NewPostgres(ctx context.Context, configDB *config.DB, logger *zap.SugaredLogger) (*PG, error) {
	pgxCfg := Config(configDB)

	db, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		logger.Errorf("failed to create pgxConfig")
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		logger.Errorf("failed to ping database")
		return nil, err
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	defer db.Close()

	return &PG{
		db,
		&psql,
		*logger,
	}, nil
}

func (db *PG) ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	} else {
		return fmt.Sprintf("unexpected error: %s", err)
	}
}

func (db *PG) Close() {
	db.Pool.Close()
}
