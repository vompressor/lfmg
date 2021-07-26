package main

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/vompressor/license_generator"
)

func main() {
	var appName = "lfmg"

	var path string

	var showInfoWithBody = false

	var owner string
	var year string

	genLicense := func(c *cli.Context) error {
		if c.NArg() == 0 {
			println("input license key")
			os.Exit(1)
		} else if c.NArg() != 1 {
			println("too many license key")
			os.Exit(1)
		}
		var license = c.Args().First()

		if license == "README.md" ||
			license == "readme" ||
			license == "r" ||
			license == "read" ||
			license == "readme.md" {
			return license_generator.CreateREADMEmd(path)
		}

		gp := filepath.Join(path, "LICENSE")

		err := license_generator.WriteLicenseBodyToPath(license, gp, year, owner)

		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		println("generated license " + license + " to " + gp)
		println(`must modify year and creator
		e.g. MIT License [3:15] and [3:22]
		     Copyright (c) [year] [fullname]
		                   ^      ^`)

		// Gen license logic

		return nil
	}

	app := &cli.App{
		Version: "1.3.1",

		Name:  appName,
		Usage: "Set \"LICENSE\" your project!",

		Commands: []*cli.Command{

			// Show license list
			{
				Name:      "list",
				Aliases:   []string{"l"},
				Usage:     "show \"LICENSE\" list",
				HideHelp:  true,
				UsageText: appName + " list",
				Action: func(c *cli.Context) error {
					err := license_generator.PrintLicenseList()
					if err != nil {
						println(err.Error())
						os.Exit(1)
					}
					os.Exit(0)
					return nil
				},
			},

			// generate "LICENSE" or "README.md"
			{
				Name:      "generate",
				Aliases:   []string{"gen", "g"},
				Usage:     "generate \"LICENSE\" or \"README.md\"",
				ArgsUsage: "[\"license key\" or \"readme\"]",
				HideHelp:  true,
				Action:    genLicense,

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "path",
						Aliases:     []string{"p"},
						Value:       ".",
						DefaultText: "working directory",
						Usage:       "`PATH` to save \"LICENSE\" or \"README.md\"",
						Destination: &path,
					},
					&cli.StringFlag{
						Name:        "year",
						Aliases:     []string{"y"},
						Value:       "",
						DefaultText: "not set",
						Usage:       "set `year`",
						Destination: &year,
					},
					&cli.StringFlag{
						Name:        "owner",
						Aliases:     []string{"o"},
						Value:       "",
						DefaultText: "not set",
						Usage:       "set `owner`",
						Destination: &owner,
					},
				},
			},

			// Get license info
			{
				Name:      "info",
				Aliases:   []string{"i"},
				Usage:     "get license info",
				ArgsUsage: "[license key]",
				HideHelp:  true,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "body",
						Aliases:     []string{"content", "b", "c"},
						Usage:       "show license content body",
						Destination: &showInfoWithBody,
						DefaultText: "unshow",
					},
				},
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						println("input license key")
						os.Exit(1)
					} else if c.NArg() != 1 {
						println("too many license key")
						os.Exit(1)
					}
					var license = c.Args().First()
					// TODO:: err expt
					license_generator.PrintLicenseInfo(license)

					if showInfoWithBody {
						println("\nbody:\n")
						license_generator.PrintLicenseBody(license)
					}

					return nil
				},
			},
			{
				Name:    "body",
				Aliases: []string{"b"},
				Usage:   "show \"LICENSE\" content",

				ArgsUsage: "[license key]",
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						println("input license key")
						os.Exit(1)
					} else if c.NArg() != 1 {
						println("too many license key")
						os.Exit(1)
					}
					var license = c.Args().First()
					license_generator.PrintLicenseBody(license)

					os.Exit(0)
					return nil
				},
			},
			{
				Name:      "cache",
				Aliases:   []string{"c"},
				Usage:     "cache control",
				UsageText: appName + " cache [command options]",
				ArgsUsage: " ",
				Action: func(c *cli.Context) error {
					println("input arg..")
					os.Exit(0)
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name:            "clear",
						Aliases:         []string{"c"},
						Usage:           "clear cache",
						ArgsUsage:       "[license key]",
						HideHelpCommand: true,

						HideHelp:  true,
						UsageText: appName + " cache [command options]",
						Action: func(c *cli.Context) error {
							license_generator.DelCache()
							println("cache clear..")
							os.Exit(0)
							return nil
						},
					},
				},
			},
		},
	}
	// TODO:: clear cache

	app.EnableBashCompletion = true
	app.Run(os.Args)

}
