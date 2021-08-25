package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	progressbar "github.com/schollz/progressbar/v3"
	"github.com/tarvos-gaming/tripwireAV/scanner"
	"github.com/tarvos-gaming/tripwireAV/scanner/rules"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No path to jar given.")
		return
	}
	println("tripwireAV\nCopyright © 2021 Tarvos Gaming")
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileReader, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	zipReader, err := zip.NewReader(bytes.NewReader(fileReader), int64(len(fileReader)))
	if err != nil {
		log.Fatal(err)
	}
	bar := progressbar.NewOptions(int(len(zipReader.File)),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionShowCount(),
		progressbar.OptionSetDescription(""),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	matches := []rules.RuleResponse{}
	for _, file := range zipReader.File {
		bar.Add(1)
		bar.Describe(fmt.Sprintf("Scanning %s", file.Name))
		if !file.FileInfo().IsDir() {
			contents, err := deflate(file)
			if err != nil {
				log.Fatal(err)
			}
			scannerMatches := scanner.Scan(contents)
			matches = append(matches, scannerMatches...)
		}
	}
	if len(matches) == 0 {
		fmt.Println("\nNo matches found.")
		return
	}
	fmt.Println("\nMatches:")
	for _, match := range matches {
		fmt.Printf("[%s] - %s\n", match.Name, match.Description)
	}
}

func deflate(file *zip.File) ([]byte, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
