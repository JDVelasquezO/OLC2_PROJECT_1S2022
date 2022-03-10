package Access

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/SymbolTable/Environment"
	"encoding/json"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type InstanceArray struct {
	Type       SymbolTable.DataType
	Dimensions *arrayList.List
}

//func NewInstanceArray(dtype SymbolTable.DataType, dimensions *arrayList.List) InstanceArray {
//	return InstanceArray{
//		Type:       dtype,
//		Dimensions: dimensions,
//	}
//}

func (i InstanceArray) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
	listIntDimensions := i.GetIntDimensions(table)
	values := i.AddValues(listIntDimensions)

	data, err := json.MarshalIndent(values, "", "  ")
	if err != nil {
		panic(err)
	}

	stringQuery := string(data)
	fmt.Printf("data %v", stringQuery)
	object := Environment.NewArray("", SymbolTable.INTEGER, values, listIntDimensions)
	objectVal := SymbolTable.ReturnType{
		Value: object,
		Type:  SymbolTable.INTEGER,
	}

	return SymbolTable.ReturnType{
		Value: objectVal,
		Type:  SymbolTable.ARRAY,
	}
}

func (i InstanceArray) GetIntDimensions(table SymbolTable.SymbolTable) *arrayList.List {
	listIntDimensions := arrayList.New()

	for x := 0; x < i.Dimensions.Len(); x++ {
		dim := i.Dimensions.GetValue(x)

		valRet := dim.(Abstract.Expression).GetValue(table)

		if valRet.Type != SymbolTable.INTEGER {
			return nil
		}
		listIntDimensions.Add(valRet.Value)
	}

	return listIntDimensions
}

func (i InstanceArray) AddValues(list *arrayList.List) []interface{} {
	newList := list.Clone()
	lengthPiv := newList.GetValue(0).(int)
	newList.RemoveAtIndex(0)

	s := make([]interface{}, lengthPiv)

	if newList.Len() > 0 {
		for x := 0; x < lengthPiv; x++ {
			s[x] = i.AddValues(newList)
		}
	} else {
		for x := 0; x < lengthPiv; x++ {
			s[x] = -1
		}
	}

	return s
}
