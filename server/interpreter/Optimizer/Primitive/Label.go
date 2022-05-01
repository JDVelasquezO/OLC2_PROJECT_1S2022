package Primitive

type Label struct {
	Name string
}

func NewLabel(name string) Label {
	return Label{
		Name: name,
	}
}
