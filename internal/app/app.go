package app

import (
	"github.com/urfave/cli/v2"

	"github.com/craftcms/nitro/internal/command"
	"github.com/craftcms/nitro/internal/executor"
)

var (
	// Version is the application version that is passed at runtime.
	Version = "1.0.0"
)

func NewApp(e executor.Executor) *cli.App {
	return &cli.App{
		Name:        "nitro",
		UsageText:   "nitro [global options] command [command options] [arguments...]",
		Usage:       "Local Craft CMS on Tap.",
		Version:     Version,
		Description: "Nitro creates virtual machines with Multipass and provides a CLI for common DevOps tasks.",
		Action:      cli.ShowAppHelp,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "machine",
				Aliases:     []string{"m"},
				Value:       "nitro-dev",
				Usage:       "Provide a machine name",
				DefaultText: "nitro-dev",
			},
		},
		Commands: []*cli.Command{
			command.Initialize(),
			command.Bootstrap(),
			command.Add(),
			command.Remove(),
			command.Attach(),
			command.SSH(),
			command.XOn(),
			command.XOff(),
			command.Start(),
			command.Stop(),
			command.Destroy(),
			command.SQL(),
			command.Redis(e),
			command.Update(),
			command.Logs(e),
			command.IP(),
		},
	}
}
