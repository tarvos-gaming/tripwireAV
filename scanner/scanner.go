package scanner

import "github.com/tarvos-gaming/tripwireAV/scanner/rules"

func Scan(contents []byte) []rules.RuleResponse {
	matches := []rules.RuleResponse{}
	for _, rule := range rules.Rules() {
		resp := rule(contents)
		if resp.Match {
			matches = append(matches, resp)
		}
	}
	return matches
}
