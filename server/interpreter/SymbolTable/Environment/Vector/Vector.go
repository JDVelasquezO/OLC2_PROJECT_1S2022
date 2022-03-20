package Vector

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type Vector struct {
	SymbolTable.Symbol
	Start  *Node
	Length int
}

func NewVector(start *Node) Vector {
	return Vector{
		Start:  start,
		Length: 0,
	}
}

func (v Vector) Push(value Abstract.Expression, table SymbolTable.SymbolTable) Vector {
	node := NewNode(value)
	if v.isEmpty() {
		v.Start = &node
	} else {
		start := v.Start
		for start.Next != nil {
			start = start.Next
		}
		start.Next = &node
	}

	v.Value = append(v.Value.([]interface{}), value.GetValue(table).Value)
	v.Length += 1
	return v
}

func (v Vector) GetValue(list *arrayList.List, table SymbolTable.SymbolTable) interface{} {
	listClone := list.Clone()
	indexPiv := listClone.GetValue(0).(Abstract.Expression).GetValue(table).Value.(int)
	if v.isEmpty() {
		return nil
	} else {
		pointer := v.Start
		counter := 0
		for counter < indexPiv && pointer.Next != nil {
			pointer = pointer.Next
			counter += 1
		}

		if counter != indexPiv {
			return nil
		} else {
			return pointer.Value
		}
	}
}

func (v Vector) isEmpty() bool {
	return v.Start == nil
}

func (v Vector) GetLength() int {
	return v.Length
}

func (v Vector) ChangeValue(newVal interface{}, symbolTable SymbolTable.SymbolTable) {

}
