package parser

type Metadata struct {
	Name        string
	Label       string
	Description string
	Validation  string
	Rejection   string
}

type Rules struct {
	Fee    int64
	Inputs []RuleInput
	Output []RuleOutput
}

type RuleInput struct {
	Name        string
	Label       string
	Description string
}

type RuleOutput struct {
	Name  string
	Label string
}
