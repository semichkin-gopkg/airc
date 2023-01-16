package cmd

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Main() {
	app := &cli.App{
		Name:        "airc",
		Usage:       "Air config builder",
		Description: "A tool for generating .toml config for air (https://github.com/cosmtrek/air)",
		Version:     "v0.0.1",
		Commands: []*cli.Command{
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build configuration",

				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:        "output",
						Aliases:     []string{"o"},
						Value:       ".air.toml",
						DefaultText: ".air.toml",
					},
				},
				Action: func(context *cli.Context) error {
					return build(context.Path("output"))
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
