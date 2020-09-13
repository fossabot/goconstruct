package project

import (
	"context"
	"flag"
	"log"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func NewCommand(logger *log.Logger) *ffcli.Command {
	var (
		rootFlagSet = flag.NewFlagSet("project", flag.ExitOnError)
		tmplsPath   = rootFlagSet.String("templates-path", "", "Path to templates.")
		tmpls       = rootFlagSet.String("templates", "glue", "A comma-separate list of template names.")
		config      = rootFlagSet.String("config", "config.hcl", "A config file defining values for the required variables for all templates used.")
	)

	templates := strings.Split(*tmpls, ",")

	return &ffcli.Command{
		Name:       "project",
		ShortUsage: "project <subcommand>",
		FlagSet:    rootFlagSet,
		Subcommands: []*ffcli.Command{
			{
				Name:       "generate",
				ShortUsage: "generate [flags]",
				ShortHelp:  "Generates a new project.",
				FlagSet:    rootFlagSet,
				Exec: generate(
					generateConfig{
						templatesPath:      tmplsPath,
						templates:          templates,
						templateConfigFile: config,
					},
					logger),
			},
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}
}
