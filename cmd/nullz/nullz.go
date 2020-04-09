package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "nullz",
		Version: "0.1.6",
		Authors: []*cli.Author{
			{
				Name:  "parkerduckworth",
				Email: "duckw0rth@protonmail.com",
			},
		},
		Usage: "convert nullable sqlboiler models for integration with msgpak",
		Commands: []*cli.Command{
			{
				Name:      "convert",
				Usage:     "convert directory of model files provided as [PATH]",
				ArgsUsage: "[PATH]",
				Action: func(c *cli.Context) error {
					if c.NArg() != 1 {
						cli.ShowAppHelpAndExit(c, 1)
					}

					execute(c.Args().Get(0))

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	terminateIfError(err)
}

func execute(path string) {
	err := processInput(path)
	terminateIfError(err)

	err = setupOutputDir()
	terminateIfError(err)

	err = walkFileTree()
	terminateIfError(err)

	err = replaceInputWithOutput()
	terminateIfError(err)
}

func processInput(path string) error {
	// Check if input exists
	_, err := os.Stat(path)
	if err != nil {
		return errors.Wrap(err, "input does not exist")
	}

	// Set abs path to input and output dirs in env
	err = initEnvVars(path)
	if err != nil {
		return errors.Wrap(err, "unable to init environment variables")
	}

	return nil
}

func initEnvVars(path string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "unable to determine current working directory")
	}

	inputFileName := filepath.Base(path)

	os.Setenv("ABS_INPUT", filepath.Join(cwd, inputFileName))
	os.Setenv("ABS_OUTPUT", filepath.Join(cwd, "nullz-models"))

	return nil
}

func setupOutputDir() error {
	path := os.Getenv("ABS_OUTPUT")
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	} else {
		return errors.Wrap(err, "failed to make output directory")
	}

	return nil
}

func walkFileTree() error {
	path := os.Getenv("ABS_INPUT")
	err := filepath.Walk(path, convert)
	if err != nil {
		return errors.Wrap(err, "failed to traverse input directory")
	}

	return nil
}

func replaceInputWithOutput() error {
	// Remove the input dir and rename output dir with name of input
	err := os.RemoveAll(os.Getenv("ABS_INPUT"))
	if err != nil {
		return errors.Wrap(err, "failed to remove input directory")
	}

	err = os.Rename(os.Getenv("ABS_OUTPUT"), os.Getenv("ABS_INPUT"))
	if err != nil {
		return errors.Wrap(err, "failed to rename output directory")
	}

	return nil
}

func convert(path string, info os.FileInfo, err error) error {
	if isDirectory(path) {
		return nil
	}

	input, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to open input file")
	}
	defer input.Close()

	output, err := os.Create(filepath.Join(os.Getenv("ABS_OUTPUT"), filepath.Base(path)))
	if err != nil {
		return errors.Wrap(err, "failed to create output file")
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := replaceWithNullz(scanner.Text())
		writer.WriteString(line + "\n")
		writer.Flush()
	}

	err = scanner.Err()
	if err != nil {
		return errors.Wrap(err, "error detected in input file contents")
	}

	log.Println("converted", path)
	return nil
}

func replaceWithNullz(line string) string {
	var replacedLine string
	replacedLine = strings.ReplaceAll(line, "null.", "nullz.")
	replacedLine = strings.ReplaceAll(replacedLine, "volatiletech/null", "parkerduckworth/nullz")
	return replacedLine
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	terminateIfError(err)
	return fileInfo.IsDir()
}

func terminateIfError(err error) {
	if err != nil {
		log.Fatal("nullz error: ", err)
	}
}
