package project

import (
	"context"
	"errors"
	"log"
)

type execFn func(context.Context, []string) error

type generateConfig struct {
	projectName *string
	projectPath *string
}

func generate(cfg generateConfig, logger *log.Logger) execFn {
	return func(ctx context.Context, _ []string) error {
		logger.Printf("name: %s\n", *cfg.projectName)
		logger.Printf("path: %s\n", *cfg.projectPath)
		return errors.New("not implemented")
	}
}
