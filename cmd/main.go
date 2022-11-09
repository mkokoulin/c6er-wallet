package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"github.com/mkokoulin/c6er-wallet.git/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer cancel()

	app, err := app.New(ctx)
	if err != nil {
		log.Println(fmt.Errorf("error in creating app: %w", err))

		return
	}

	stop, err := app.Server.Start(ctx)
	if err != nil {
		log.Println(fmt.Errorf("error in starting server: %w", err))

		return
	}

	<-ctx.Done()

	stop(ctx)

	app.Logger.Info().Msg("application stopped")
}
