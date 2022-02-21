package utilities

import "github.com/antlr/antlr4/runtime/Go/antlr"

type CustomSyntaxError struct {
	Row, Column int
	Msg         string
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []CustomSyntaxError
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, symbol interface{}, row, col int,
	msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, CustomSyntaxError{
		Row:    row,
		Column: col,
		Msg:    msg,
	})
}
