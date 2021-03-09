package main

import (
	"os"

	"github.com/kris-nova/novaarchive/filesystem"

	"github.com/kris-nova/tambourine"

	"github.com/kris-nova/logger"

	"github.com/urfave/cli"
)

var (
	serviceOptions        = &tambourine.ServiceOptions{}
	kubeConfig     string = ""
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "kubeConfig",
				Value:       "~/.kube/config",
				Usage:       "kube config path",
				Destination: &kubeConfig,
			},
		},
		Action: func(c *cli.Context) error {
			serviceOptions.KubeConfigPath = filesystem.NewPath(kubeConfig)
			logger.BitwiseLevel = logger.LogEverything
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
