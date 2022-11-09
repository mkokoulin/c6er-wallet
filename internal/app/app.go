package app

import (
	"context"
	"os"

	"github.com/mkokoulin/c6er-wallet.git/internal/config"
	"github.com/mkokoulin/c6er-wallet.git/internal/database"
	"github.com/mkokoulin/c6er-wallet.git/internal/handlers"
	"github.com/mkokoulin/c6er-wallet.git/internal/models"
	"github.com/mkokoulin/c6er-wallet.git/internal/router"
	"github.com/mkokoulin/c6er-wallet.git/internal/server"
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type App struct {
	DB *database.Database
	Logger *zerolog.Logger
	Server *server.Server
}

func New(ctx context.Context)(*App, error) {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel)

	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	conn, err := gorm.Open(postgres.Open(cfg.DSN))
	if err != nil {
		return nil, err
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	err = conn.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	db := database.New(conn, &logger)

	h := handlers.New(db, cfg, &logger)

	r := router.New(h, cfg)

	s := server.New(cfg.Address, r)


	return &App{
		DB: db,
		Logger: &logger,
		Server: s,
	}, nil
}