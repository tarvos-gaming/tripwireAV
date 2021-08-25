package rules

import (
	"log"
	"regexp"
)

func TAV0003(contents []byte) RuleResponse {
	ruleName := "TAV0003"
	desc := "Discord Session Stealer"
	matched := false
	exprs := []string{
		"discord/Local Storage/leveldb",
		"discordptb/Local Storage/leveldb",
		"discordcanary/Local Storage/leveldb",
		"Local Storage/leveldb",
		"Local Storage",
		"leveldb",
	}
	for _, expr := range exprs {
		match, err := regexp.Match(expr, contents)
		if err != nil {
			log.Fatal(err)
		}
		if match {
			matched = true
		}
	}
	return RuleResponse{
		Name:        ruleName,
		Description: desc,
		Match:       matched,
	}
}
