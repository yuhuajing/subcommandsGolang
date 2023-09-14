package cmd

import (
	"github.com/urfave/cli/v2"
)

var (
	PowCmd = &cli.Command{
		Name:        "pow",
		Usage:       "pow actions",
		ArgsUsage:   "",
		Description: ``,
		Subcommands: []*cli.Command{
			{
				Name:  "hash_rate",
				Usage: "hash_rate",
				//ArgsUsage: "<keyFile>",
				Action: hashRate,
				Flags: []cli.Flag{&cli.StringFlag{
					Name:  "remotedb",
					Usage: "URL for remote database",
				}, &cli.BoolFlag{
					Name:  "usb",
					Usage: "Enable monitoring and management of USB hardware wallets",
				},
				},
				Description: ``,
			},
		},
	}
)

func hashRate(ctx *cli.Context) (err error) {
	return
}
