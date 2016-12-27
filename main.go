package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "ansibrest plugin"
	app.Usage = "ansibrest plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "endpoint",
			Usage:  "endpoint",
			EnvVar: "PLUGIN_ENDPOINT",
		},
		cli.StringFlag{
			Name:   "playbook",
			Usage:  "playbook",
			EnvVar: "PLUGIN_PLAYBOOK",
		},
		cli.StringFlag{
			Name:   "inventory",
			Usage:  "inventory",
			EnvVar: "PLUGIN_INVENTORY",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		EndPoint:   c.String("endpoint"),
		PlayBook:   c.String("playbook"),
		Inventory: c.String("inventory"),
	}

	return plugin.Exec()
}
