package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := &cli.App{
		Name: "nullz",
		Version: "0.1.0",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "parkerduckworth",
				Email: "duckw0rth@protonmail.com",
			},
		},
		Usage: "convert nullable sqlboiler models for integration with msgpak",
		Commands: []*cli.Command{
			{
				Name:  "convert",
				Usage: "convert file provided as [PATH]",
				ArgsUsage: "[PATH]",
				Action: func(c *cli.Context) error {
					if c.NArg() != 1 {
						nullzError(fmt.Errorf("incorrect number of arguments"))
						cli.ShowAppHelpAndExit(c, 1)
					}

					convert(c.Args().Get(0))

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		nullzError(err)
	}
}

func convert(path string) {
	f, err := os.Open(path)
	if err != nil {
		nullzError(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "null.", "nullz.")
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		nullzError(err)
	}
}

func nullzError(err error) {
	log.Fatal("nullz error: ", err)
}
