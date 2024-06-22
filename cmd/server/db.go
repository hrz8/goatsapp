package server

import (
	"context"
	"os"

	"github.com/hrz8/goatsapp/config"
	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/hrz8/gofx"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func NewDBRepo(db *pgxpool.Pool) *dbrepo.Queries {
	return dbrepo.New(db)
}

func NewDB(lc fx.Lifecycle, logger *gofx.Logger, cfg *config.Config) *pgxpool.Pool {
	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		logger.Error("cannot connect to database", "err", err)
		os.Exit(1)
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			conn.Close()
			return nil
		},
	})

	return conn
}
