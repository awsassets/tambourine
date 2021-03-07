package main

import (
	"os"

	"github.com/kris-nova/tambourine"

	"github.com/kris-nova/logger"

	"github.com/urfave/cli"
)

var serviceOptions = &tambourine.ServiceOptions{}

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "verbose",
				Value: 0,
				Usage: "verbosity level",
			},
		},
		Action: func(c *cli.Context) error {
			service := tambourine.New(serviceOptions)
			return service.Run()
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}
}
