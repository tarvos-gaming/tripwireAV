package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"flag"

	"github.com/tarvos-gaming/tripwireAV/engine"
)

func main() {
	pathPtr := flag.String("file", "", "path to jar to scan")
	flag.Parse()
	println("tripwireAV\nCopyright © 2021 Tarvos Gaming")
	e := engine.New()
	var matches []engine.Matches
	if *pathPtr != "" {
		if filepath.Ext(*pathPtr) == ".jar" {
			matches = append(matches, e.ScanJAR((*pathPtr)))
		} else {
			log.Fatal("File is not a jar.")
		}
	} else {
		mcDir := getMinecraftPath()
		err := filepath.Walk(mcDir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if filepath.Ext(path) == ".jar" {
					scanMatches := e.ScanJAR(path)
					matches = append(matches, scanMatches)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Println("Matches:")
	matched := false
	for _, file := range matches {
		if len(file.Matches) > 0 {
			matched = true
			fmt.Printf("%s:\n", file.Name)
			for _, match := range file.Matches {
				fmt.Printf("-> [%s] - %s\n", match.Rule, match.Metas[0].Value)
			}
		}
	}
	if !matched {
		fmt.Println("None.")
	}
}
