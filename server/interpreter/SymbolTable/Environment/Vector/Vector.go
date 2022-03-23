package Vector

import (
	"OLC2_Project1/server/interpreter"
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strconv"
)

type Vector struct {
	SymbolTable.Symbol
	Start           *Node
	Length          int
	IsCapacityFixed bool
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
	if v.IsCapacityFixed {
		if len(v.Value.([]interface{})) > v.Length {
			v.Length += v.Length
		}
	} else {
		v.Length += 1
	}

	return v
}

func (v Vector) Remove(index int, table SymbolTable.SymbolTable) interface{} {
	if !v.isEmpty() {
		if index == 0 {
			first := v.Start
			v.Start = v.Start.Next
			first.Next = nil
		} else if index < v.Length {
			pointer := v.Start
			counter := 0
			for counter < (index - 1) {
				pointer = pointer.Next
				counter += 1
			}

			temp := pointer.Next
			pointer.Next = temp.Next
			temp.Next = nil
		} else {
			goto IndexError
		}

		v.Value = removeIndex(v.Value.([]interface{}), index)
		v.Length -= 1
	}

	return v

IndexError:
	row := v.Row
	col := v.Col
	errors.CounterError += 1
	msg := "(" + strconv.Itoa(row) + ", " + strconv.Itoa(col) + ") Error: Indice excedido: " + strconv.Itoa(index) + " \n"
	err := errors.NewError(errors.CounterError, row, col, msg, table.Name)
	errors.TypeError = append(errors.TypeError, err)
	interpreter.Console += fmt.Sprintf("%v", err.Msg)
	return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err.Msg}
}

func removeIndex(s []interface{}, index int) []interface{} {
	return append(s[:index], s[index+1:]...)
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

func (v Vector) Contains() {

}

func (v Vector) isEmpty() bool {
	return v.Start == nil
}

func (v Vector) GetLength() int {
	return v.Length
}

func (v Vector) ChangeValue(newVal interface{}, symbolTable SymbolTable.SymbolTable) {

}
