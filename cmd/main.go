package main

import (
	"context"
	"os"

	"github.com/cloudslit/deca/internal"
	"github.com/cloudslit/deca/pkg/logger"
	"github.com/urfave/cli/v2"
)

var VERSION = "0.0.1"

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "ca"
	app.Version = VERSION
	app.Usage = "CA PKI"
	app.Commands = []*cli.Command{
		newTlsCmd(ctx),
		newOcspCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

func newTlsCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "tls",
		Usage: "Run TLS server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "App configuration file(.json,.yaml,.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return internal.Run(ctx,
				internal.SetConfigFile(c.String("conf")),
				internal.SetVersion(VERSION),
				internal.SetAppType(c.Command.Name),
			)
		},
	}
}

func newOcspCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "ocsp",
		Usage: "Run OCSP server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "App configuration file(.json,.yaml,.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return internal.Run(ctx,
				internal.SetConfigFile(c.String("conf")),
				internal.SetVersion(VERSION),
				internal.SetAppType(c.Command.Name),
			)
		},
	}
}
