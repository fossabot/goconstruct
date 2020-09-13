package {{goconstruct::name}}

import (
	"context"
	"flag"
	"log"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func NewCommand(logger *log.Logger) *ffcli.Command {
	var (
		rootFlagSet = flag.NewFlagSet("{{goconstruct::name}}", flag.ExitOnError)
	)

	return &ffcli.Command{
		Name:       "{{goconstruct::name}}",
		ShortUsage: "{{goconstruct::name}} <subcommand>",
		FlagSet:    replicaFlagSet,
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}
}
