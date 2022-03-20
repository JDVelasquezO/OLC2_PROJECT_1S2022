package DecVectors

import (
	"OLC2_Project1/server/interpreter/Abstract"
	arrayList "github.com/colegno/arraylist"
)

type AssignVector struct {
	Val       Abstract.Expression
	Id        string
	ListIndex *arrayList.List
}

func NewAssignVector(val Abstract.Expression, id string, listIndex *arrayList.List) AssignVector {
	return AssignVector{
		Val:       val,
		Id:        id,
		ListIndex: listIndex,
	}
}

//func (a AssignVector) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
//	if a.Val == nil {
//		row := 0
//		col := 0
//		errors.CounterError += 1
//		msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") No se puede mezclar esos tipos de datos \n"
//		err := errors.NewError(errors.CounterError, row, col, msg, symbolTable.Name)
//		errors.TypeError = append(errors.TypeError, err)
//		interpreter.Console += fmt.Sprintf("%v", err.Msg)
//		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
//	}
//
//	foundSymbol := symbolTable.GetSymbol(a.Id)
//	symbolVector := foundSymbol.(Vector.Vector)
//	symbolVector.ChangeValue(symbolVector.Value, a.Val, symbolTable)
//	return symbolVector
//}
