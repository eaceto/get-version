package main

import (
	"get-module/version"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"strings"
)

func main() {

	var path string

	app := &cli.App{
		Name: "get-version",
		Usage: "It obtains the version of your application, using the source code, agnostically to the technology that is being used.",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "path",
				Value: ".",
				Usage: "root path of the source code to analyze",
				Destination: &path,
			},
		},
		Action: func(c *cli.Context) error {
			path = strings.TrimSpace(path)

			version := version.GetVersion{
				RootPath: path,
				Logger: log.Default(),
			}

			var err error
			if _, err = version.Analyze();  err != nil {

			}

			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}