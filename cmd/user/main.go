package main

import (
	"context"
	config2 "go-social/cmd/user/config"
	"go-social/internal/user/app"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	config, err := config2.NewConfig()
	if err != nil {
		panic(err)
	}
	app, err := app.Init(config)
	err = app.Run()
	if err != nil {
		slog.Info(err.Error())
		return
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		slog.Info("signal.Notify", v)
	case done := <-ctx.Done():
		slog.Info("ctx.Done", done)
	}

}
