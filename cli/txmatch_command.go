package cli

import (
	"context"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

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
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt)

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
			logger, err := cfg.ZapLogger()
			if err != nil {
				return err
			}
			matcher, err := txmatch.NewMatcher(
				logger,
				bc,
				string(abiBytes),
				c.StringSlice("methods"),
				[]string{c.String("contract.address")},
			)
			if err != nil {
				return err
			}
			go func() {
				<-ch
				cancel()
			}()
			return matcher.Match(c.String("out.file"), c.Uint64("start.block"), c.Uint64("end.block"))
		},
		Flags: []cli.Flag{
			&cli.Uint64Flag{
				Name:  "start.block",
				Value: 10950651,
				Usage: "start of the block range to query",
			},
			&cli.Uint64Flag{
				Name:  "end.block",
				Value: 12089509,
				Usage: "end of the block range to query",
			},
			&cli.StringFlag{
				Name:  "contract.address",
				Value: "",
				Usage: "contract address to filter transactions against",
			},
			&cli.StringSliceFlag{
				Name:  "methods",
				Usage: "method to filter transactions against",
			},
			&cli.StringFlag{
				Name:  "abi.file",
				Usage: "file containing the abi definition",
				Value: "abifile.txt",
			},
			&cli.StringFlag{
				Name:  "out.file",
				Usage: "file to store matched addresses in",
				Value: "outfile.txt",
			},
		},
	}
}
