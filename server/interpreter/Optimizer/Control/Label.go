package Control

type Label struct {
	Id  string
	Row int
	Col int
}

func NewLabel(id string, row int, col int) Label {
	return Label{
		Id:  id,
		Row: row,
		Col: col,
	}
}

func (l Label) Execute() interface{} {
	return l.Id + ":"
}
