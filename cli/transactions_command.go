package cli

import (
	"context"
	"fmt"

	"github.com/bonedaddy/go-defi/bclient"
	"github.com/bonedaddy/go-defi/config"
	"github.com/urfave/cli/v2"
)

func transactionsCommand() *cli.Command {
	return &cli.Command{
		Name:    "transactions",
		Aliases: []string{"txs"},
		Usage:   "command for workign with transactions",
		Subcommands: cli.Commands{
			&cli.Command{
				Name:    "encode",
				Aliases: []string{"enc"},
				Usage:   "hex encoded transaction binary representation",
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
					bc, err := bclient.NewClient(ctx, client)
					if err != nil {
						return err
					}
					enc, err := bc.EncodeTx(c.String("tx.hash"))
					if err != nil {
						return err
					}
					fmt.Println(enc)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "tx.hash",
				Value: "",
				Usage: "hash of the transaction to work with",
			},
		},
	}
}
