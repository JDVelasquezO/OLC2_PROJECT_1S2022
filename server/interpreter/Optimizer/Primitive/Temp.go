package Primitive

type Temp struct {
	Value interface{}
	Temp  string
	Row   int
	Col   int
}

func NewTemp(temp string, row int, col int) Temp {
	return Temp{
		Temp: temp,
		Row:  row,
		Col:  col,
	}
}

func (t Temp) GetValue() interface{} {
	return t.Temp
}
