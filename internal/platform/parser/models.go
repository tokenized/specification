package parser

type Metadata struct {
	Name        string
	Label       string
	Type        string
	Description string
	Validation  string
	Rejection   string
}

// Rules are the TX input/output rules for the Transaction.
type Rules struct {
	Fee     int64
	Inputs  []RuleInputOutput
	Outputs []RuleInputOutput
}

// RuleInputOutput can represent a TX input or output.
type RuleInputOutput struct {
	Name     string
	Label    string
	Comments string
}

// ruleRow exists as a presentation helper that wraps a pair of
// RuleInputOutput structs together for the purpose of being able to render
// them together as a single row.
type ruleRow struct {
	InputIndex  string
	Input       RuleInputOutput
	OutputIndex string
	Output      RuleInputOutput
}
