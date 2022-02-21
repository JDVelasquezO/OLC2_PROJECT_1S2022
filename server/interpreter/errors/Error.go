package errors

type Error struct {
	Id    int
	Type  string
	Line  int
	Col   int
	Msg   string
	Ambit string
}

func NewError(id int, line int, col int, msg string, ambit string) Error {
	return Error{
		Id:    id,
		Type:  "Semantic",
		Line:  line,
		Col:   col,
		Msg:   msg,
		Ambit: ambit,
	}
}

var TypeError []Error
var CounterError int
