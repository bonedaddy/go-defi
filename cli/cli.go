package cli

import (
	"github.com/urfave/cli/v2"
)

func New(name, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.Commands = cli.Commands{configCommand(), txMatchCommand(), transactionsCommand()}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config.path",
			Aliases: []string{"cfg"},
			Usage:   "path to the configuration file",
			Value:   "config.yaml",
		},
	}
	return app
}
