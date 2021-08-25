package engine

import (
	"io/ioutil"
	"log"
	"os"

	yara "github.com/hillu/go-yara/v4"
)

type Engine struct {
	Compiler *yara.Compiler
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
	return Engine{
		Compiler: c,
	}
}

func (e Engine) Scan(contents []byte) yara.MatchRules {
	r, err := e.Compiler.GetRules()
	if err != nil {
		log.Fatal(err)
	}
	var matches yara.MatchRules
	err = r.ScanMem(contents, 0, 0, &matches)
	if err != nil {
		log.Fatal(err)
	}
	return matches
}
