package main

import (
	"context"

	"github.com/asawo/api/config"
	"github.com/asawo/api/db"
	"github.com/asawo/api/logger"
)

// setupLogger creates a new logger given the log level
func setupLogger(env *config.Env) (*logger.Log, error) {
	log, err := logger.New(env.LogLevel)
	if err != nil {
		return nil, err
	}

	return log, nil
}

// setupDB creates a new DB to store data
func setupDB(ctx context.Context, env *config.Env) (db.DB, error) {
	repo, err := db.New(env)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
