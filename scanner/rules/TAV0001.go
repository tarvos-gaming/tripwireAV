package rules

import (
	"log"
	"regexp"
)

func TAV0001(contents []byte) RuleResponse {
	ruleName := "TAV0001"
	desc := "Minecraft Account Session Stealer"
	matched := false
	exprs := []string{
		"launcher_accounts",
		"launcher_profiles",
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
