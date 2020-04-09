package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "nullz",
		Version: "0.1.5",
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
	checkNullzError(err)
}

func execute(path string) {
	processInput(path)
	setupOutputDir()
	walkFileTree()
	replaceInputWithOutput()
}

func processInput(path string) {
	// Check if input exists
	_, err := os.Stat(path)
	checkNullzError(err)

	// Set abs path to input and output dirs in env
	initEnvVars(path)
}

func initEnvVars(path string) {
	cwd, err := os.Getwd()
	checkNullzError(err)
	inputFileName := filepath.Base(path)

	os.Setenv("ABS_INPUT", filepath.Join(cwd, inputFileName))
	os.Setenv("ABS_OUTPUT", filepath.Join(cwd, "nullz-models"))
}

func setupOutputDir() {
	path := os.Getenv("ABS_OUTPUT")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func walkFileTree() {
	path := os.Getenv("ABS_INPUT")
	err := filepath.Walk(path, convert)
	checkNullzError(err)
}

// Remove the input dir and rename output dir with name of input
func replaceInputWithOutput() {
	err := os.RemoveAll(os.Getenv("ABS_INPUT"))
	checkNullzError(err)

	err = os.Rename(os.Getenv("ABS_OUTPUT"), os.Getenv("ABS_INPUT"))
	checkNullzError(err)
}

func convert(path string, info os.FileInfo, err error) error {
	if isDirectory(path) {
		return nil
	}

	input, err := os.Open(path)
	checkNullzError(err)
	defer input.Close()

	output, err := os.Create(filepath.Join(os.Getenv("ABS_OUTPUT"), filepath.Base(path)))
	checkNullzError(err)
	defer output.Close()

	writer := bufio.NewWriter(output)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := replaceWithNullz(scanner.Text())
		writer.WriteString(line + "\n")
		writer.Flush()
	}

	checkNullzError(scanner.Err())
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
	checkNullzError(err)
	return fileInfo.IsDir()
}

func checkNullzError(err error) {
	if err != nil {
		log.Fatal("nullz error: ", err)
	}
}
