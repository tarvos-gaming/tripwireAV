package rules

type RuleResponse struct {
	Name        string
	Description string
	Match       bool
}

func Rules() []func(contents []byte) RuleResponse {
	return []func(contents []byte) RuleResponse{
		TAV0001,
		TAV0002,
		TAV0003,
	}
}
