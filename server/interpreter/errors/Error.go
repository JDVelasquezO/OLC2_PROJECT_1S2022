package errors

import "time"

type Error struct {
	Id    int
	Type  string
	Line  int
	Col   int
	Msg   string
	Ambit string
	Date  string
}

func NewError(id int, line int, col int, msg string, ambit string) Error {
	msgDate := time.Now().Format(time.RFC822)
	return Error{
		Id:    id,
		Type:  "Semantic",
		Line:  line,
		Col:   col,
		Msg:   msg,
		Ambit: ambit,
		Date:  msgDate,
	}
}

var TypeError []Error
var CounterError int
