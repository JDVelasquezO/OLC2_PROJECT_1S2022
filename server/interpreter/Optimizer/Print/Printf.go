package Print

import (
	"OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
	"fmt"
)

type Printf struct {
	Row      int
	Col      int
	Format   string
	TypeCast string
	Value    AbstractOptimizer.Expression
}

func NewPrintf(format string, typeCast string, value AbstractOptimizer.Expression, row int, col int) Printf {
	return Printf{
		Row:      row,
		Col:      col,
		Format:   format,
		TypeCast: typeCast,
		Value:    value,
	}
}

func (f Printf) Execute() interface{} {
	if len(f.Format) > 3 {
		f.Format = f.Format[1 : len(f.Format)-1]
	}

	return "printf(\"" + f.Format + "\", (" + f.TypeCast + ")" + fmt.Sprintf("%v", f.Value.GetValue()) + ");"
}
