package rules

import (
	"log"
	"regexp"
)

func TAV0002(contents []byte) RuleResponse {
	ruleName := "TAV0002"
	desc := "Discord Webhook Call (data exfil + telemetery)"
	matched := false
	exprs := []string{
		"discord.com/api/webhooks",
		"webhooks",
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
