package project

import (
	"context"
	"flag"
	"log"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func NewCommand(logger *log.Logger) *ffcli.Command {
	var (
		rootFlagSet = flag.NewFlagSet("project", flag.ExitOnError)
		name        = rootFlagSet.String("name", "", "The new project's name.")
		path        = rootFlagSet.String("path", "", "Where on your filesystem the project should be generated.")
	)

	return &ffcli.Command{
		Name:       "project",
		ShortUsage: "project <subcommand>",
		FlagSet:    rootFlagSet,
		Subcommands: []*ffcli.Command{
			{
				Name:       "generate ",
				ShortUsage: "generate [flags]",
				ShortHelp:  "Generates a new project.",
				FlagSet:    rootFlagSet,
				Exec: generate(
					generateConfig{
						projectName: name,
						projectPath: path,
					},
					logger),
			},
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}
}
