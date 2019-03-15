package parser

import "fmt"

type Metadata struct {
	Name        string
	Label       string
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

// Rows returns a []ruleRow which makes iterating rows of inputs and outputs
// easier.
func (r Rules) Rows() []ruleRow {
	max := len(r.Inputs)
	if len(r.Outputs) > max {
		max = len(r.Outputs)
	}

	rows := make([]ruleRow, max, max)

	for i, input := range r.Inputs {
		rows[i].Input = input
		rows[i].InputIndex = fmt.Sprintf("%v", i)
	}

	for i, output := range r.Outputs {
		rows[i].Output = output
		rows[i].OutputIndex = fmt.Sprintf("%v", i)
	}

	return rows
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
