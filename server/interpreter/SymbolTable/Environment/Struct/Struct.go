package Struct

import arrayList "github.com/colegno/arraylist"

type Struct struct {
	Id    string
	Items *arrayList.List
}

func NewStruct(id string, items *arrayList.List) *Struct {
	return &Struct{
		Id:    id,
		Items: items,
	}
}
