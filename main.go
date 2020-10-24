package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/mccurdyc/goconstruct/pkg/project"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/rs/zerolog"
)

func main() {
	l := zerolog.New(os.Stdout)
	logger := log.New(l, "goconstruct - ", log.Ldate|log.Ltime|log.LUTC)

	os.Exit(Run(context.Background(), os.Args[1:], logger))
}

func Run(ctx context.Context, args []string, l *log.Logger) int {
	var (
		rootFlagSet = flag.NewFlagSet("goconstruct", flag.ExitOnError)
	)

	root := &ffcli.Command{
		Name:       "goconstruct",
		ShortUsage: "goconstruct <subcommand> [flags]",
		FlagSet:    rootFlagSet,
		Subcommands: []*ffcli.Command{
			project.NewCommand(args[2:], l),
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	err := root.ParseAndRun(context.Background(), os.Args[1:])
	if err != nil {
		l.Fatal(err)
		return 1
	}

	return 0
}
