package cli

import (
	"context"
	"io/ioutil"

	"github.com/bonedaddy/go-defi/bclient"
	"github.com/bonedaddy/go-defi/config"
	"github.com/bonedaddy/go-defi/txmatch"
	"github.com/urfave/cli/v2"
)

func txMatchCommand() *cli.Command {
	return &cli.Command{
		Name:    "txmatch",
		Aliases: []string{"txm"},
		Usage:   "allows filtering for transactions that match predefined conditions",
		Action: func(c *cli.Context) error {
			ctx, cancel := context.WithCancel(c.Context)
			defer cancel()
			cfg, err := config.LoadConfig(c.String("config.path"))
			if err != nil {
				return err
			}
			client, err := cfg.EthClient(ctx)
			if err != nil {
				return err
			}
			abiBytes, err := ioutil.ReadFile(c.String("abi.file"))
			if err != nil {
				return err
			}
			bc, err := bclient.NewClient(ctx, client)
			if err != nil {
				return err
			}
			matcher, err := txmatch.NewMatcher(
				bc,
				string(abiBytes),
				[]string{c.String("method")},
				[]string{c.String("contract.address")},
			)
			if err != nil {
				return err
			}
			return matcher.Match(c.Uint64("start.block"), c.Uint64("end.block"))
		},
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:  "start.block",
				Value: 0,
				Usage: "start of the block range to query",
			},
			&cli.Uint64Flag{
				Name:  "end.block",
				Value: 0,
				Usage: "end of the block range to query",
			},
			&cli.StringFlag{
				Name:  "contract.address",
				Value: "",
				Usage: "contract address to filter transactions against",
			},
			&cli.StringFlag{
				Name:  "method",
				Usage: "method to filter transactions against",
			},
			&cli.StringFlag{
				Name:  "abi.file",
				Usage: "file containing the abi definition",
			},
		},
	}
}
