package errors

type Error struct {
	Id    int
	Type  string
	Line  int
	Col   int
	Msg   string
	Ambit string
}

var TypeError []Error
var CounterError int
