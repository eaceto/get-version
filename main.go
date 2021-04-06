package main

import (
	"fmt"
	"get-module/version"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
	"strings"
)

func main() {

	var path string
	verbose := false

	app := &cli.App{
		Name: "get-version",
		Usage: "It obtains the version of your application, using the source code, agnostically to the technology that is being used.",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "path",
				Value: ".",
				Usage: "--path SOURCE_CODE_PATH",
				Required: false,
				Destination: &path,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Required:    false,
				Hidden:      false,
				Value:       false,
				Destination: &verbose,
			},
		},
		Action: func(c *cli.Context) error {
			path = strings.TrimSpace(path)

			logger := log.Default()
			if !verbose {
				logger = nil
			}

			appVersion := version.GetVersion{
				RootPath: path,
				Logger: logger,
			}

			if info, err := appVersion.Analyze();  err == nil {
				fmt.Print(info.Version)
				return nil
			} else {
				return err
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}