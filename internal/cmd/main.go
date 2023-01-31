package cmd

import (
	"context"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-sigint
		cancel()
	}()

	app := &cli.App{
		Name:    "airc",
		Usage:   "Air-based utility for live reloading with config building by throwing env variables",
		Version: "v0.0.7",
		Commands: []*cli.Command{
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build configuration",

				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:        "config",
						Aliases:     []string{"c"},
						Value:       ".air.toml",
						DefaultText: ".air.toml",
					},
				},
				Action: func(context *cli.Context) error {
					return build(ctx, context.Path("config"))
				},
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Usage:   "Run air",

				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:        "config",
						Aliases:     []string{"c"},
						Value:       ".air.toml",
						DefaultText: ".air.toml",
					},
				},
				Action: func(context *cli.Context) error {
					return run(ctx, context.Path("config"))
				},
			},
			{
				Name:    "build-run",
				Aliases: []string{"br"},
				Usage:   "Build config and run air",

				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:        "config",
						Aliases:     []string{"c"},
						Value:       ".air.toml",
						DefaultText: ".air.toml",
					},
				},
				Action: func(context *cli.Context) error {
					configPath := context.Path("config")
					if err := build(ctx, configPath); err != nil {
						return err
					}
					return run(ctx, configPath)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
