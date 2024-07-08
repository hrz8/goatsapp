package server

import (
	"context"
	"os"

	"github.com/hrz8/goatsapp/config"
	"github.com/hrz8/goatsapp/internal/dbrepo"
	"github.com/hrz8/goatsapp/pkg/whatsapp"
	"github.com/hrz8/gofx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
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

func NewWaCli(lc fx.Lifecycle, cfg *config.Config, pool *pgxpool.Pool, ev whatsapp.EventHandler) *whatsapp.Client {
	db := stdlib.OpenDBFromPool(pool)
	cli := whatsapp.NewClient(
		whatsapp.WithDB(db),
		whatsapp.WithOsInfo(cfg.AppName, cfg.AppVersionNumber),
		whatsapp.WithEventHandler(ev),
	)
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			err := cli.Upgrade()
			if err != nil {
				return err
			}
			cli.Restore()
			return nil
		},
		OnStop: func(_ context.Context) error {
			cli.Backup()
			return nil
		},
	})
	return cli
}
