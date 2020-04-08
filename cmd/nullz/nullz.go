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
	err := setupOutputInfo(path)
	checkNullzError(err)

	walkDirTree(path)
}

func setupOutputInfo(path string) error {
	parentDir, err := getParentDir(path)
	checkNullzError(err)

	os.Setenv("CONVERTED_PATH", parentDir+"/converted-models")

	if err := os.Mkdir(os.Getenv("CONVERTED_PATH"), 0777); err != nil {
		return err
	}

	return nil
}

func getParentDir(path string) (string, error) {
	parentDir := filepath.Dir(strings.TrimSuffix(path, "/"))
	if parentDir == "." {
		abspath, err := os.Getwd()
		if err != nil {
			return "", err
		}
		parentDir = abspath
	}

	return parentDir, nil
}

func walkDirTree(path string) {
	err := filepath.Walk(path, convert)
	checkNullzError(err)
}

func convert(path string, info os.FileInfo, err error) error {
	isDir, err := isDirectory(path)
	checkNullzError(err)
	if isDir {
		return nil
	}

	input, err := os.Open(path)
	checkNullzError(err)
	defer input.Close()

	outputPath := filepath.Join(os.Getenv("CONVERTED_PATH"), filepath.Base(path))
	output, err := os.Create(outputPath)
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

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func checkNullzError(err error) {
	if err != nil {
		log.Fatal("nullz error: ", err)
	}
}
