package engine

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yara "github.com/hillu/go-yara/v4"
	"github.com/schollz/progressbar/v3"
)

type Engine struct {
	Compiler *yara.Compiler
	Rules    *yara.Rules
}

func New() Engine {
	c, err := yara.NewCompiler()
	if err != nil {
		log.Fatalf("Failed to init YARA parser: %s", err)
	}
	ruleFiles, err := ioutil.ReadDir("./rules")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range ruleFiles {
		if !file.IsDir() {
			f, err := os.Open("./rules/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			c.AddFile(f, "tripwire")
		}
	}
	r, err := c.GetRules()
	if err != nil {
		log.Fatal(err)
	}
	return Engine{
		Compiler: c,
		Rules:    r,
	}
}

func (e Engine) scan(contents []byte) yara.MatchRules {
	var matches yara.MatchRules
	err := e.Rules.ScanMem(contents, 0, 0, &matches)
	if err != nil {
		log.Fatal(err)
	}
	return matches
}

type Matches struct {
	Matches yara.MatchRules
	Name    string
}

func (e Engine) ScanJAR(path string) Matches {
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

	// look for manifest
	// TODO: find better method of detecting game jar
	mc := false
	for _, file := range zipReader.File {
		if file.Name == "META-INF/MOJANGCS.RSA" {
			mc = true
		}
	}

	if mc {
		return Matches{
			yara.MatchRules{},
			path,
		}
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
	var matches yara.MatchRules
	for _, file := range zipReader.File {
		bar.Add(1)
		bar.Describe(fmt.Sprintf("Scanning %s", file.Name))
		if !file.FileInfo().IsDir() {
			contents, err := deflate(file)
			if err != nil {
				log.Fatal(err)
			}
			scannerMatches := e.scan(contents)
			matches = append(matches, scannerMatches...)
		}
	}
	err = bar.Close()
	if err != nil {
		log.Fatal(err)
	}
	return Matches{
		matches,
		path,
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
