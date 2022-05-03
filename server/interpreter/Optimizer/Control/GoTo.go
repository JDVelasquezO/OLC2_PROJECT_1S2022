package Control

type GoTo struct {
	Label string
	Row   int
	Col   int
}

func NewGoTo(label string, row int, col int) GoTo {
	return GoTo{
		Label: label,
		Row:   row,
		Col:   col,
	}
}

func (g GoTo) Execute() interface{} {
	return "goto " + g.Label + ";"
}
