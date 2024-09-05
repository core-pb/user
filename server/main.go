package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/core-pb/authenticate/authenticate/v1/authenticateconnect"
	tc "github.com/core-pb/tag/client"
	"github.com/core-pb/tag/tag/v1"
	"github.com/core-pb/user/user/v1/userconnect"
	"github.com/uptrace/bun"
	"go.x2ox.com/sorbifolia/bunpgd"
	"go.x2ox.com/sorbifolia/crpc"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		closeCh     = make(chan os.Signal, 1)
		server, err = crpc.NewServer(
			crpc.WithHealthAndMetrics(":80", ""),
			crpc.WithCertFromCheck("CERT", "cert", "build/output/cert"),
			crpc.WithCORS(nil),
		)
	)
	if err != nil {
		slog.Error("create server", slog.String("err", err.Error()))
		os.Exit(1)
	}

	initDB(ctx)
	initTagServer(ctx)
	initAuthServer()

	server.Handle(userconnect.NewBaseHandler(base{}))
	server.Handle(userconnect.NewTagHandler(_tag{}))

	if err = server.Run(); err != nil {
		slog.Error("server run", slog.String("err", err.Error()))
		os.Exit(1)
	}

	signal.Notify(closeCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	server.SetReady()
	sig := <-closeCh
	server.SetNoReady("server close")

	slog.Info("server exit signal: signal notify %s", slog.String("sig", sig.String()))

	cancel()

	if err = server.Close(); err != nil {
		slog.Error("server close", slog.String("err", err.Error()))
	}
}

func initDB(ctx context.Context) {
	var (
		_ctx, cancel = context.WithCancelCause(ctx)
		err          error
	)

	if db, err = bunpgd.Open(os.Getenv("DB_DSN"),
		bunpgd.WithMaxOpenConns(256),
		bunpgd.WithMaxIdleConns(8),
		bunpgd.WithConnMaxIdleTime(time.Second*12),
		bun.WithDiscardUnknownColumns(),
		bunpgd.WithSLog(),
		bunpgd.WithCreateTable(_ctx, cancel, &User{}, &UserTag{}, &UserAuth{}),
	); err != nil {
		slog.Error("connect db", slog.String("err", err.Error()))
		os.Exit(1)
	}

	if err = _ctx.Err(); err != nil {
		slog.Error("create table", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

var module *tag.Module

func initTagServer(ctx context.Context) {
	tc.Set(nil, os.Getenv("TAG_ADDR"))

	res, err := tc.Get().RegisterModule(ctx, "user")
	if err != nil {
		slog.Error("register module", slog.String("err", err.Error()))
		os.Exit(1)
	}
	module = res
}

var auth authenticateconnect.BaseClient

func initAuthServer() {
	auth = crpc.Client(authenticateconnect.NewBaseClient, nil, os.Getenv("AUTH_ADDR"))
}
