package Function

type Call struct {
	Row int
	Col int
	Id  string
}

func NewCall(id string, row int, col int) Call {
	return Call{
		Row: row,
		Col: col,
		Id:  id,
	}
}

func (c Call) Execute() interface{} {
	return c.Id + "();"
}
