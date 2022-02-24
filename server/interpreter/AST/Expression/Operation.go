package Expression

import (
	"OLC2_Project1/server/interpreter/Abstract"
	"OLC2_Project1/server/interpreter/SymbolTable"
	"OLC2_Project1/server/interpreter/errors"
	"fmt"
	"math"
	"strconv"
)

var sum = [5][5]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

var multiDiv = [5][5]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

var sub = [5][5]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

var relational = [5][5]SymbolTable.DataType{
	{SymbolTable.INTEGER, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.BOOLEAN, SymbolTable.NULL},
	{SymbolTable.FLOAT, SymbolTable.FLOAT, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.BOOLEAN, SymbolTable.NULL},
	{SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL, SymbolTable.NULL},
}

type Operation struct {
	OpLeft   Abstract.Expression
	Operator string
	OpRight  Abstract.Expression
	Unary    bool
	Row      int
	Col      int
}

func NewOperation(OpLeft Abstract.Expression, Operator string, OpRight Abstract.Expression,
	unary bool, row int, col int) Operation {
	e := Operation{OpLeft, Operator, OpRight, unary, row, col}
	return e
}

func (p Operation) GetValue(symbolTable SymbolTable.SymbolTable) SymbolTable.ReturnType {
	var retLeft SymbolTable.ReturnType
	var retRight SymbolTable.ReturnType

	if p.OpLeft.GetValue(symbolTable).Value == nil {

		row := strconv.Itoa(p.Row)
		col := strconv.Itoa(p.Col)
		errors.CounterError += 1
		msg := "(" + row + ", " + col + ") ID " + p.OpLeft.(Identifier).Id + " no declarado o asignado"
		err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
		errors.TypeError = append(errors.TypeError, err)
		return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
	}

	if p.Unary {
		retLeft = p.OpLeft.GetValue(symbolTable)
	} else {
		if p.OpRight.GetValue(symbolTable).Value == nil {
			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") ID " + p.OpRight.(Identifier).Id + " no asignado"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		retLeft = p.OpLeft.GetValue(symbolTable)
		retRight = p.OpRight.GetValue(symbolTable)
	}

	var priority SymbolTable.DataType

	switch p.Operator {
	case "+":
		priority = sum[retLeft.Type][retRight.Type]

		if priority == SymbolTable.INTEGER {
			fmt.Println(retLeft.Type)
			fmt.Println(retRight.Type)
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) + retRight.Value.(int)}

		} else if priority == SymbolTable.FLOAT {
			op1, _ := strconv.ParseFloat(fmt.Sprintf("%v", retLeft.Value), 64)
			op2, _ := strconv.ParseFloat(fmt.Sprintf("%v", retRight.Value), 64)
			return SymbolTable.ReturnType{Type: priority, Value: op1 + op2}

		} else if priority == SymbolTable.STR {
			r1 := fmt.Sprintf("%v", retLeft.Value)
			r2 := fmt.Sprintf("%v", retRight.Value)

			return SymbolTable.ReturnType{Type: priority, Value: r1 + r2}
		}

	case "-":
		if p.Unary {
			value := retLeft.Value.(int) * -1
			return SymbolTable.ReturnType{Type: retLeft.Type, Value: value}
		}

		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '-' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = sub[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) - retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) - retRight.Value.(float64)}
		}

	case "*":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '*' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = multiDiv[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) * retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) * retRight.Value.(float64)}
		}

	case "/":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '/' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = multiDiv[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) / retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) / retRight.Value.(float64)}
		}

	case "%":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '%' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = multiDiv[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) % retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: math.Mod(retLeft.Value.(float64), retRight.Value.(float64))}
		}

	case "pow", "powf":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación 'pow' no soportada para este tipo de dato"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = relational[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER && p.Operator == "pow" {
			return SymbolTable.ReturnType{Type: priority, Value: math.Pow(float64(retLeft.Value.(int)), float64(retRight.Value.(int)))}
		} else if priority == SymbolTable.FLOAT && p.Operator == "powf" {
			return SymbolTable.ReturnType{Type: priority, Value: math.Pow(retLeft.Value.(float64), retRight.Value.(float64))}
		} else {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: el operador " + p.Operator + " no puede operarse con este tipo de dato "
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

	case ">":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '>' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = relational[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) > retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) > retRight.Value.(float64)}
		}

	case "<":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '<' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = relational[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) < retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) < retRight.Value.(float64)}
		}

	case ">=":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '>=' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = relational[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) >= retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) >= retRight.Value.(float64)}
		}

	case "<=":
		if retLeft.Type == SymbolTable.STR ||
			retLeft.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.BOOLEAN ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '<=' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = relational[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) <= retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) <= retRight.Value.(float64)}
		}

	case "==":
		if retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.NULL {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación '==' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}

		priority = relational[retLeft.Type][retRight.Type]
		if priority == SymbolTable.INTEGER {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(int) == retRight.Value.(int)}
		} else if priority == SymbolTable.FLOAT {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(float64) == retRight.Value.(float64)}
		} else if priority == SymbolTable.BOOLEAN {

			var left bool
			var right bool
			if retLeft.Value == "true" {
				left = true
			} else {
				left = false
			}

			if retRight.Value == "true" {
				right = true
			} else {
				right = false
			}
			return SymbolTable.ReturnType{Type: priority, Value: left == right}

		} else if priority == SymbolTable.STR {
			return SymbolTable.ReturnType{Type: priority, Value: retLeft.Value.(string) == retRight.Value.(string)}
		}

	case "AND":
		if retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.INTEGER ||
			retRight.Type == SymbolTable.INTEGER ||
			retLeft.Type == SymbolTable.FLOAT ||
			retRight.Type == SymbolTable.FLOAT ||
			retLeft.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.STR {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación 'AND' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}
		var left bool
		var right bool
		if retLeft.Value == "true" {
			left = true
		} else {
			left = false
		}

		if retRight.Value == "true" {
			right = true
		} else {
			right = false
		}

		return SymbolTable.ReturnType{Type: priority, Value: left && right}

	case "OR":
		if retLeft.Type == SymbolTable.NULL ||
			retRight.Type == SymbolTable.NULL ||
			retLeft.Type == SymbolTable.INTEGER ||
			retRight.Type == SymbolTable.INTEGER ||
			retLeft.Type == SymbolTable.FLOAT ||
			retRight.Type == SymbolTable.FLOAT ||
			retLeft.Type == SymbolTable.STR ||
			retRight.Type == SymbolTable.STR {

			row := strconv.Itoa(p.Row)
			col := strconv.Itoa(p.Col)
			errors.CounterError += 1
			msg := "(" + row + ", " + col + ") Error: Operación 'OR' no soportada"
			err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
			errors.TypeError = append(errors.TypeError, err)
			return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
		}
		var left bool
		var right bool
		if retLeft.Value == "true" {
			left = true
		} else {
			left = false
		}

		if retRight.Value == "true" {
			right = true
		} else {
			right = false
		}
		return SymbolTable.ReturnType{Type: priority, Value: left || right}

	}

	row := strconv.Itoa(p.Row)
	col := strconv.Itoa(p.Col)
	errors.CounterError += 1
	msg := "(" + row + ", " + col + ") Error: tipos de datos no soportados "
	err := errors.NewError(errors.CounterError, p.Row, p.Col, msg, symbolTable.Name)
	errors.TypeError = append(errors.TypeError, err)
	return SymbolTable.ReturnType{Type: SymbolTable.ERROR, Value: err}
}
