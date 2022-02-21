package utilities

import (
	"OLC2_Project1/server/interpreter/AST"
	"OLC2_Project1/server/parser"
)

type TreeShapeListener struct {
	*parser.BaseProjectParserListener
	Tree AST.Tree
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (l *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	l.Tree = ctx.GetTree()
}
