package cli

import (
	"github.com/bonedaddy/go-defi/config"
	"github.com/urfave/cli/v2"
)

func configCommand() *cli.Command {
	return &cli.Command{
		Name:    "configuration",
		Aliases: []string{"config", "cfg"},
		Usage:   "configuration management command group",
		Subcommands: cli.Commands{
			&cli.Command{
				Name:    "generate",
				Aliases: []string{"gen"},
				Action: func(c *cli.Context) error {
					return config.NewConfig(c.String("config.path"))
				},
			},
		},
	}
}
