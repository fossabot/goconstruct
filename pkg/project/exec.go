package project

import (
	"context"
	"errors"
	"log"
)

type execFn func(context.Context, []string) error

type generateConfig struct {
	templatesPath      *string
	templates          []string
	templateConfigFile *string
}

func generate(cfg generateConfig, logger *log.Logger) execFn {
	return func(ctx context.Context, _ []string) error {
		logger.Printf("templates: %v\n", cfg.templates)
		return errors.New("not implemented")
	}
}
