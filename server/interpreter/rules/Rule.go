package rules

type Rule struct {
	Rule      string
	Original  string
	Optimized string
	File      int
}

func NewRule(rule string, original string,
	optimized string, file int) Rule {

	return Rule{
		Rule:      rule,
		Original:  original,
		Optimized: optimized,
		File:      file,
	}
}

var TypeRule []Rule
var CounterRule int
