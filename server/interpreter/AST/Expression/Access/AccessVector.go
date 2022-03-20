package Access

import (
	"OLC2_Project1/server/interpreter/SymbolTable"
	arrayList "github.com/colegno/arraylist"
)

type VectorAccess struct {
	Id  string
	Dim *arrayList.List
}

func (v VectorAccess) Execute(symbolTable SymbolTable.SymbolTable) interface{} {
	//TODO implement me
	panic("implement me")
}

func (v VectorAccess) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	//TODO implement me
	panic("implement me")
}

func NewAccessVector(id string, dims *arrayList.List) VectorAccess {
	return VectorAccess{
		Id:  id,
		Dim: dims,
	}
}

//func (v VectorAccess) Execute(table SymbolTable.SymbolTable) interface{} {
//	res := v.GetValue(table)
//	return res
//}

//func (v VectorAccess) GetValue(table SymbolTable.SymbolTable) SymbolTable.ReturnType {
//	exists := table.ExistsSymbol(v.Id)
//
//	if !exists {
//		fmt.Println("No existe el vector")
//		return SymbolTable.ReturnType{
//			Value: nil,
//			Type:  SymbolTable.NULL,
//		}
//	}
//
//	symbol := table.GetSymbol(v.Id)
//	if reflect.TypeOf(symbol) != reflect.TypeOf(Vector.Vector{}) {
//		fmt.Println("No es un vector")
//		return SymbolTable.ReturnType{
//			Value: 0,
//			Type:  SymbolTable.NULL,
//		}
//	}
//
//	objVector := symbol.(Vector.Vector)
//
//	//if objVector.ListIntDim.Len() < v.Dim.Len() {
//	//	fmt.Println("Dimensiones invalidas")
//	//	return SymbolTable.ReturnType{
//	//		Value: 0,
//	//		Type:  SymbolTable.NULL,
//	//	}
//	//}
//
//	//val := objVector.GetValue(0)
//	return SymbolTable.ReturnType{
//		Value: val,
//		Type:  objVector.DataType,
//	}
//}
