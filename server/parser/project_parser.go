// Code generated from ProjectParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // ProjectParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2_Project1/server/interpreter/AST"
import "OLC2_Project1/server/interpreter/Abstract"
import "OLC2_Project1/server/interpreter/AST/Expression"
import "OLC2_Project1/server/interpreter/AST/Natives"
import "OLC2_Project1/server/interpreter/SymbolTable"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 47, 252,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 3, 3, 7, 3, 39, 10, 3, 12, 3, 14, 3, 42, 11, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 58, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 96, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 7, 6, 107, 10, 6, 12, 6, 14, 6, 110, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 124,
	10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 146, 10, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 156, 10,
	10, 12, 10, 14, 10, 159, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	176, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 5, 13, 187, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 7, 14, 198, 10, 14, 12, 14, 14, 14, 201, 11, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 221, 10, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 233,
	10, 16, 12, 16, 14, 16, 236, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 250, 10, 17, 3,
	17, 2, 6, 10, 18, 26, 30, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 2, 6, 4, 2, 36, 36, 38, 41, 3, 2, 13, 14, 4, 2, 42, 43,
	46, 46, 3, 2, 44, 45, 2, 260, 2, 34, 3, 2, 2, 2, 4, 40, 3, 2, 2, 2, 6,
	57, 3, 2, 2, 2, 8, 95, 3, 2, 2, 2, 10, 97, 3, 2, 2, 2, 12, 123, 3, 2, 2,
	2, 14, 125, 3, 2, 2, 2, 16, 145, 3, 2, 2, 2, 18, 147, 3, 2, 2, 2, 20, 160,
	3, 2, 2, 2, 22, 175, 3, 2, 2, 2, 24, 186, 3, 2, 2, 2, 26, 188, 3, 2, 2,
	2, 28, 202, 3, 2, 2, 2, 30, 220, 3, 2, 2, 2, 32, 249, 3, 2, 2, 2, 34, 35,
	5, 4, 3, 2, 35, 36, 8, 2, 1, 2, 36, 3, 3, 2, 2, 2, 37, 39, 5, 6, 4, 2,
	38, 37, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3,
	2, 2, 2, 41, 43, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 8, 3, 1, 2, 44,
	5, 3, 2, 2, 2, 45, 46, 5, 8, 5, 2, 46, 47, 8, 4, 1, 2, 47, 58, 3, 2, 2,
	2, 48, 49, 5, 12, 7, 2, 49, 50, 8, 4, 1, 2, 50, 58, 3, 2, 2, 2, 51, 52,
	5, 14, 8, 2, 52, 53, 8, 4, 1, 2, 53, 58, 3, 2, 2, 2, 54, 55, 5, 20, 11,
	2, 55, 56, 8, 4, 1, 2, 56, 58, 3, 2, 2, 2, 57, 45, 3, 2, 2, 2, 57, 48,
	3, 2, 2, 2, 57, 51, 3, 2, 2, 2, 57, 54, 3, 2, 2, 2, 58, 7, 3, 2, 2, 2,
	59, 60, 7, 4, 2, 2, 60, 61, 7, 28, 2, 2, 61, 62, 7, 30, 2, 2, 62, 63, 5,
	24, 13, 2, 63, 64, 7, 31, 2, 2, 64, 65, 7, 25, 2, 2, 65, 66, 8, 5, 1, 2,
	66, 96, 3, 2, 2, 2, 67, 68, 7, 4, 2, 2, 68, 69, 7, 28, 2, 2, 69, 70, 7,
	30, 2, 2, 70, 71, 5, 24, 13, 2, 71, 72, 7, 26, 2, 2, 72, 73, 5, 10, 6,
	2, 73, 74, 7, 31, 2, 2, 74, 75, 7, 25, 2, 2, 75, 76, 8, 5, 1, 2, 76, 96,
	3, 2, 2, 2, 77, 78, 7, 3, 2, 2, 78, 79, 7, 28, 2, 2, 79, 80, 7, 30, 2,
	2, 80, 81, 5, 24, 13, 2, 81, 82, 7, 31, 2, 2, 82, 83, 7, 25, 2, 2, 83,
	84, 8, 5, 1, 2, 84, 96, 3, 2, 2, 2, 85, 86, 7, 3, 2, 2, 86, 87, 7, 28,
	2, 2, 87, 88, 7, 30, 2, 2, 88, 89, 5, 24, 13, 2, 89, 90, 7, 26, 2, 2, 90,
	91, 5, 10, 6, 2, 91, 92, 7, 31, 2, 2, 92, 93, 7, 25, 2, 2, 93, 94, 8, 5,
	1, 2, 94, 96, 3, 2, 2, 2, 95, 59, 3, 2, 2, 2, 95, 67, 3, 2, 2, 2, 95, 77,
	3, 2, 2, 2, 95, 85, 3, 2, 2, 2, 96, 9, 3, 2, 2, 2, 97, 98, 8, 6, 1, 2,
	98, 99, 5, 24, 13, 2, 99, 100, 8, 6, 1, 2, 100, 108, 3, 2, 2, 2, 101, 102,
	12, 4, 2, 2, 102, 103, 7, 26, 2, 2, 103, 104, 5, 24, 13, 2, 104, 105, 8,
	6, 1, 2, 105, 107, 3, 2, 2, 2, 106, 101, 3, 2, 2, 2, 107, 110, 3, 2, 2,
	2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 11, 3, 2, 2, 2, 110,
	108, 3, 2, 2, 2, 111, 112, 7, 5, 2, 2, 112, 113, 5, 16, 9, 2, 113, 114,
	7, 23, 2, 2, 114, 115, 5, 24, 13, 2, 115, 116, 7, 25, 2, 2, 116, 117, 8,
	7, 1, 2, 117, 124, 3, 2, 2, 2, 118, 119, 7, 5, 2, 2, 119, 120, 5, 16, 9,
	2, 120, 121, 7, 25, 2, 2, 121, 122, 8, 7, 1, 2, 122, 124, 3, 2, 2, 2, 123,
	111, 3, 2, 2, 2, 123, 118, 3, 2, 2, 2, 124, 13, 3, 2, 2, 2, 125, 126, 5,
	18, 10, 2, 126, 127, 7, 23, 2, 2, 127, 128, 5, 24, 13, 2, 128, 129, 7,
	25, 2, 2, 129, 130, 8, 8, 1, 2, 130, 15, 3, 2, 2, 2, 131, 132, 5, 18, 10,
	2, 132, 133, 8, 9, 1, 2, 133, 146, 3, 2, 2, 2, 134, 135, 5, 18, 10, 2,
	135, 136, 7, 27, 2, 2, 136, 137, 7, 6, 2, 2, 137, 138, 8, 9, 1, 2, 138,
	146, 3, 2, 2, 2, 139, 140, 5, 18, 10, 2, 140, 141, 7, 27, 2, 2, 141, 142,
	7, 29, 2, 2, 142, 143, 7, 10, 2, 2, 143, 144, 8, 9, 1, 2, 144, 146, 3,
	2, 2, 2, 145, 131, 3, 2, 2, 2, 145, 134, 3, 2, 2, 2, 145, 139, 3, 2, 2,
	2, 146, 17, 3, 2, 2, 2, 147, 148, 8, 10, 1, 2, 148, 149, 7, 22, 2, 2, 149,
	150, 8, 10, 1, 2, 150, 157, 3, 2, 2, 2, 151, 152, 12, 4, 2, 2, 152, 153,
	7, 26, 2, 2, 153, 154, 7, 22, 2, 2, 154, 156, 8, 10, 1, 2, 155, 151, 3,
	2, 2, 2, 156, 159, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2,
	2, 158, 19, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160, 161, 7, 11, 2, 2, 161,
	162, 7, 30, 2, 2, 162, 163, 5, 24, 13, 2, 163, 164, 7, 31, 2, 2, 164, 165,
	5, 22, 12, 2, 165, 166, 8, 11, 1, 2, 166, 21, 3, 2, 2, 2, 167, 168, 7,
	34, 2, 2, 168, 169, 5, 4, 3, 2, 169, 170, 7, 35, 2, 2, 170, 171, 8, 12,
	1, 2, 171, 176, 3, 2, 2, 2, 172, 173, 7, 34, 2, 2, 173, 174, 7, 35, 2,
	2, 174, 176, 8, 12, 1, 2, 175, 167, 3, 2, 2, 2, 175, 172, 3, 2, 2, 2, 176,
	23, 3, 2, 2, 2, 177, 178, 5, 26, 14, 2, 178, 179, 8, 13, 1, 2, 179, 187,
	3, 2, 2, 2, 180, 181, 5, 30, 16, 2, 181, 182, 8, 13, 1, 2, 182, 187, 3,
	2, 2, 2, 183, 184, 5, 28, 15, 2, 184, 185, 8, 13, 1, 2, 185, 187, 3, 2,
	2, 2, 186, 177, 3, 2, 2, 2, 186, 180, 3, 2, 2, 2, 186, 183, 3, 2, 2, 2,
	187, 25, 3, 2, 2, 2, 188, 189, 8, 14, 1, 2, 189, 190, 5, 30, 16, 2, 190,
	191, 8, 14, 1, 2, 191, 199, 3, 2, 2, 2, 192, 193, 12, 4, 2, 2, 193, 194,
	9, 2, 2, 2, 194, 195, 5, 26, 14, 5, 195, 196, 8, 14, 1, 2, 196, 198, 3,
	2, 2, 2, 197, 192, 3, 2, 2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2,
	2, 199, 200, 3, 2, 2, 2, 200, 27, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202,
	203, 5, 26, 14, 2, 203, 204, 9, 3, 2, 2, 204, 205, 5, 26, 14, 2, 205, 206,
	8, 15, 1, 2, 206, 29, 3, 2, 2, 2, 207, 208, 8, 16, 1, 2, 208, 209, 7, 45,
	2, 2, 209, 210, 5, 24, 13, 2, 210, 211, 8, 16, 1, 2, 211, 221, 3, 2, 2,
	2, 212, 213, 5, 32, 17, 2, 213, 214, 8, 16, 1, 2, 214, 221, 3, 2, 2, 2,
	215, 216, 7, 30, 2, 2, 216, 217, 5, 24, 13, 2, 217, 218, 7, 31, 2, 2, 218,
	219, 8, 16, 1, 2, 219, 221, 3, 2, 2, 2, 220, 207, 3, 2, 2, 2, 220, 212,
	3, 2, 2, 2, 220, 215, 3, 2, 2, 2, 221, 234, 3, 2, 2, 2, 222, 223, 12, 6,
	2, 2, 223, 224, 9, 4, 2, 2, 224, 225, 5, 30, 16, 7, 225, 226, 8, 16, 1,
	2, 226, 233, 3, 2, 2, 2, 227, 228, 12, 5, 2, 2, 228, 229, 9, 5, 2, 2, 229,
	230, 5, 30, 16, 6, 230, 231, 8, 16, 1, 2, 231, 233, 3, 2, 2, 2, 232, 222,
	3, 2, 2, 2, 232, 227, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2,
	2, 2, 234, 235, 3, 2, 2, 2, 235, 31, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2,
	237, 238, 7, 17, 2, 2, 238, 250, 8, 17, 1, 2, 239, 240, 7, 18, 2, 2, 240,
	250, 8, 17, 1, 2, 241, 242, 7, 20, 2, 2, 242, 250, 8, 17, 1, 2, 243, 244,
	7, 19, 2, 2, 244, 250, 8, 17, 1, 2, 245, 246, 7, 21, 2, 2, 246, 250, 8,
	17, 1, 2, 247, 248, 7, 22, 2, 2, 248, 250, 8, 17, 1, 2, 249, 237, 3, 2,
	2, 2, 249, 239, 3, 2, 2, 2, 249, 241, 3, 2, 2, 2, 249, 243, 3, 2, 2, 2,
	249, 245, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 33, 3, 2, 2, 2, 16, 40,
	57, 95, 108, 123, 145, 157, 175, 186, 199, 220, 232, 234, 249,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'println'", "'print'", "'let'", "'String'", "'Integer'", "'Real'",
	"'Boolean'", "'str'", "'If'", "'entonces'", "'AND'", "'OR'", "", "", "",
	"", "", "", "", "", "'='", "'.'", "';'", "','", "':'", "'!'", "'&'", "'('",
	"')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='", "'>'", "'<'", "'>='",
	"'<='", "'*'", "'/'", "'+'", "'-'", "'%'",
}
var symbolicNames = []string{
	"", "PRINTLN", "PRINT", "DECLARAR", "RSTRING", "RINTEGER", "RREAL", "RBOOLEAN",
	"RSTR", "RIF", "RELSE", "RAND", "ROR", "MULTILINE", "INLINE", "INTEGER",
	"FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON",
	"COMMA", "COLON", "ADMIRATION", "REFERENCE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY",
	"RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL",
	"GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL",
	"DIV", "ADD", "SUB", "MOD", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instructions", "instruction", "print_prod", "listVars", "declaration_prod",
	"assign_prod", "ids_type", "listIds", "conditional_prod", "bloq", "expression",
	"expr_rel", "expr_logic", "expr_arit", "primitive",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ProjectParser struct {
	*antlr.BaseParser
}

func NewProjectParser(input antlr.TokenStream) *ProjectParser {
	this := new(ProjectParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ProjectParser.g4"

	return this
}

type BinaryOperation struct {
	Op1      int
	Operator string
	Op2      int
}

// ProjectParser tokens.
const (
	ProjectParserEOF               = antlr.TokenEOF
	ProjectParserPRINTLN           = 1
	ProjectParserPRINT             = 2
	ProjectParserDECLARAR          = 3
	ProjectParserRSTRING           = 4
	ProjectParserRINTEGER          = 5
	ProjectParserRREAL             = 6
	ProjectParserRBOOLEAN          = 7
	ProjectParserRSTR              = 8
	ProjectParserRIF               = 9
	ProjectParserRELSE             = 10
	ProjectParserRAND              = 11
	ProjectParserROR               = 12
	ProjectParserMULTILINE         = 13
	ProjectParserINLINE            = 14
	ProjectParserINTEGER           = 15
	ProjectParserFLOAT             = 16
	ProjectParserCHAR              = 17
	ProjectParserSTRING            = 18
	ProjectParserBOOLEAN           = 19
	ProjectParserID                = 20
	ProjectParserEQUAL             = 21
	ProjectParserDOT               = 22
	ProjectParserSEMICOLON         = 23
	ProjectParserCOMMA             = 24
	ProjectParserCOLON             = 25
	ProjectParserADMIRATION        = 26
	ProjectParserREFERENCE         = 27
	ProjectParserLEFT_PAR          = 28
	ProjectParserRIGHT_PAR         = 29
	ProjectParserLEFT_KEY          = 30
	ProjectParserRIGHT_KEY         = 31
	ProjectParserLEFT_BRACKET      = 32
	ProjectParserRIGHT_BRACKET     = 33
	ProjectParserEQUEAL_EQUAL      = 34
	ProjectParserNOTEQUAL          = 35
	ProjectParserGREATER_THAN      = 36
	ProjectParserLESS_THAN         = 37
	ProjectParserGREATER_EQUALTHAN = 38
	ProjectParserLESS_EQUEALTHAN   = 39
	ProjectParserMUL               = 40
	ProjectParserDIV               = 41
	ProjectParserADD               = 42
	ProjectParserSUB               = 43
	ProjectParserMOD               = 44
	ProjectParserWHITESPACE        = 45
)

// ProjectParser rules.
const (
	ProjectParserRULE_start            = 0
	ProjectParserRULE_instructions     = 1
	ProjectParserRULE_instruction      = 2
	ProjectParserRULE_print_prod       = 3
	ProjectParserRULE_listVars         = 4
	ProjectParserRULE_declaration_prod = 5
	ProjectParserRULE_assign_prod      = 6
	ProjectParserRULE_ids_type         = 7
	ProjectParserRULE_listIds          = 8
	ProjectParserRULE_conditional_prod = 9
	ProjectParserRULE_bloq             = 10
	ProjectParserRULE_expression       = 11
	ProjectParserRULE_expr_rel         = 12
	ProjectParserRULE_expr_logic       = 13
	ProjectParserRULE_expr_arit        = 14
	ProjectParserRULE_primitive        = 15
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instructions returns the _instructions rule contexts.
	Get_instructions() IInstructionsContext

	// Set_instructions sets the _instructions rule contexts.
	Set_instructions(IInstructionsContext)

	// GetTree returns the tree attribute.
	GetTree() AST.Tree

	// SetTree sets the tree attribute.
	SetTree(AST.Tree)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	tree          AST.Tree
	_instructions IInstructionsContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *StartContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *StartContext) GetTree() AST.Tree { return s.tree }

func (s *StartContext) SetTree(v AST.Tree) { s.tree = v }

func (s *StartContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ProjectParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProjectParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)

		var _x = p.Instructions()

		localctx.(*StartContext)._instructions = _x
	}
	localctx.(*StartContext).tree = AST.NewTree(localctx.(*StartContext).Get_instructions().GetL())

	return localctx
}

// IInstructionsContext is an interface to support dynamic dispatch.
type IInstructionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetE returns the e rule context list.
	GetE() []IInstructionContext

	// SetE sets the e rule context list.
	SetE([]IInstructionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstructionsContext differentiates from other interfaces.
	IsInstructionsContext()
}

type InstructionsContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruction IInstructionContext
	e            []IInstructionContext
}

func NewEmptyInstructionsContext() *InstructionsContext {
	var p = new(InstructionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_instructions
	return p
}

func (*InstructionsContext) IsInstructionsContext() {}

func NewInstructionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsContext {
	var p = new(InstructionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_instructions

	return p
}

func (s *InstructionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionsContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *InstructionsContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *InstructionsContext) GetE() []IInstructionContext { return s.e }

func (s *InstructionsContext) SetE(v []IInstructionContext) { s.e = v }

func (s *InstructionsContext) GetL() *arrayList.List { return s.l }

func (s *InstructionsContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstructionsContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *InstructionsContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *InstructionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterInstructions(s)
	}
}

func (s *InstructionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitInstructions(s)
	}
}

func (p *ProjectParser) Instructions() (localctx IInstructionsContext) {
	localctx = NewInstructionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProjectParserRULE_instructions)

	localctx.(*InstructionsContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProjectParserPRINTLN)|(1<<ProjectParserPRINT)|(1<<ProjectParserDECLARAR)|(1<<ProjectParserRIF)|(1<<ProjectParserID))) != 0 {
		{
			p.SetState(35)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstructionsContext).GetE()
	for _, e := range listInt {
		localctx.(*InstructionsContext).l.Add(e.GetInstr())
	}
	fmt.Print("tipo %T", localctx.(*InstructionsContext).GetE())

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_print_prod returns the _print_prod rule contexts.
	Get_print_prod() IPrint_prodContext

	// Get_declaration_prod returns the _declaration_prod rule contexts.
	Get_declaration_prod() IDeclaration_prodContext

	// Get_assign_prod returns the _assign_prod rule contexts.
	Get_assign_prod() IAssign_prodContext

	// Get_conditional_prod returns the _conditional_prod rule contexts.
	Get_conditional_prod() IConditional_prodContext

	// Set_print_prod sets the _print_prod rule contexts.
	Set_print_prod(IPrint_prodContext)

	// Set_declaration_prod sets the _declaration_prod rule contexts.
	Set_declaration_prod(IDeclaration_prodContext)

	// Set_assign_prod sets the _assign_prod rule contexts.
	Set_assign_prod(IAssign_prodContext)

	// Set_conditional_prod sets the _conditional_prod rule contexts.
	Set_conditional_prod(IConditional_prodContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	instr             Abstract.Instruction
	_print_prod       IPrint_prodContext
	_declaration_prod IDeclaration_prodContext
	_assign_prod      IAssign_prodContext
	_conditional_prod IConditional_prodContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_print_prod() IPrint_prodContext { return s._print_prod }

func (s *InstructionContext) Get_declaration_prod() IDeclaration_prodContext {
	return s._declaration_prod
}

func (s *InstructionContext) Get_assign_prod() IAssign_prodContext { return s._assign_prod }

func (s *InstructionContext) Get_conditional_prod() IConditional_prodContext {
	return s._conditional_prod
}

func (s *InstructionContext) Set_print_prod(v IPrint_prodContext) { s._print_prod = v }

func (s *InstructionContext) Set_declaration_prod(v IDeclaration_prodContext) {
	s._declaration_prod = v
}

func (s *InstructionContext) Set_assign_prod(v IAssign_prodContext) { s._assign_prod = v }

func (s *InstructionContext) Set_conditional_prod(v IConditional_prodContext) {
	s._conditional_prod = v
}

func (s *InstructionContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *InstructionContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *InstructionContext) Print_prod() IPrint_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrint_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrint_prodContext)
}

func (s *InstructionContext) Declaration_prod() IDeclaration_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaration_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaration_prodContext)
}

func (s *InstructionContext) Assign_prod() IAssign_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_prodContext)
}

func (s *InstructionContext) Conditional_prod() IConditional_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditional_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditional_prodContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *ProjectParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProjectParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserPRINTLN, ProjectParserPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)

			var _x = p.Print_prod()

			localctx.(*InstructionContext)._print_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_print_prod().GetInstr()

	case ProjectParserDECLARAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)

			var _x = p.Declaration_prod()

			localctx.(*InstructionContext)._declaration_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_declaration_prod().GetInstr()

	case ProjectParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)

			var _x = p.Assign_prod()

			localctx.(*InstructionContext)._assign_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_assign_prod().GetInstr()

	case ProjectParserRIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(52)

			var _x = p.Conditional_prod()

			localctx.(*InstructionContext)._conditional_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_conditional_prod().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrint_prodContext is an interface to support dynamic dispatch.
type IPrint_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LEFT_PAR returns the _LEFT_PAR token.
	Get_LEFT_PAR() antlr.Token

	// Set_LEFT_PAR sets the _LEFT_PAR token.
	Set_LEFT_PAR(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// GetOpBefore returns the opBefore rule contexts.
	GetOpBefore() IExpressionContext

	// Get_listVars returns the _listVars rule contexts.
	Get_listVars() IListVarsContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// SetOpBefore sets the opBefore rule contexts.
	SetOpBefore(IExpressionContext)

	// Set_listVars sets the _listVars rule contexts.
	Set_listVars(IListVarsContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsPrint_prodContext differentiates from other interfaces.
	IsPrint_prodContext()
}

type Print_prodContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_LEFT_PAR   antlr.Token
	_expression IExpressionContext
	opBefore    IExpressionContext
	_listVars   IListVarsContext
}

func NewEmptyPrint_prodContext() *Print_prodContext {
	var p = new(Print_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_print_prod
	return p
}

func (*Print_prodContext) IsPrint_prodContext() {}

func NewPrint_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_prodContext {
	var p = new(Print_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_print_prod

	return p
}

func (s *Print_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_prodContext) Get_LEFT_PAR() antlr.Token { return s._LEFT_PAR }

func (s *Print_prodContext) Set_LEFT_PAR(v antlr.Token) { s._LEFT_PAR = v }

func (s *Print_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Print_prodContext) GetOpBefore() IExpressionContext { return s.opBefore }

func (s *Print_prodContext) Get_listVars() IListVarsContext { return s._listVars }

func (s *Print_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Print_prodContext) SetOpBefore(v IExpressionContext) { s.opBefore = v }

func (s *Print_prodContext) Set_listVars(v IListVarsContext) { s._listVars = v }

func (s *Print_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Print_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Print_prodContext) PRINT() antlr.TerminalNode {
	return s.GetToken(ProjectParserPRINT, 0)
}

func (s *Print_prodContext) ADMIRATION() antlr.TerminalNode {
	return s.GetToken(ProjectParserADMIRATION, 0)
}

func (s *Print_prodContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *Print_prodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Print_prodContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_PAR, 0)
}

func (s *Print_prodContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *Print_prodContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *Print_prodContext) ListVars() IListVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListVarsContext)
}

func (s *Print_prodContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(ProjectParserPRINTLN, 0)
}

func (s *Print_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterPrint_prod(s)
	}
}

func (s *Print_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitPrint_prod(s)
	}
}

func (p *ProjectParser) Print_prod() (localctx IPrint_prodContext) {
	localctx = NewPrint_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProjectParserRULE_print_prod)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Match(ProjectParserPRINT)
		}
		{
			p.SetState(58)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(59)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(60)

			var _x = p.Expression()

			localctx.(*Print_prodContext)._expression = _x
		}
		{
			p.SetState(61)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(62)
			p.Match(ProjectParserSEMICOLON)
		}
		localctx.(*Print_prodContext).instr = Natives.NewPrint(localctx.(*Print_prodContext).Get_expression().GetP(), false, (func() int {
			if localctx.(*Print_prodContext).Get_LEFT_PAR() == nil {
				return 0
			} else {
				return localctx.(*Print_prodContext).Get_LEFT_PAR().GetLine()
			}
		}()), localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Match(ProjectParserPRINT)
		}
		{
			p.SetState(66)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(67)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(68)

			var _x = p.Expression()

			localctx.(*Print_prodContext).opBefore = _x
		}
		{
			p.SetState(69)
			p.Match(ProjectParserCOMMA)
		}
		{
			p.SetState(70)

			var _x = p.listVars(0)

			localctx.(*Print_prodContext)._listVars = _x
		}
		{
			p.SetState(71)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(72)
			p.Match(ProjectParserSEMICOLON)
		}
		localctx.(*Print_prodContext).instr = Natives.NewPrintWithAfter(localctx.(*Print_prodContext).GetOpBefore().GetP(), localctx.(*Print_prodContext).Get_listVars().GetList(), false, (func() int {
			if localctx.(*Print_prodContext).Get_LEFT_PAR() == nil {
				return 0
			} else {
				return localctx.(*Print_prodContext).Get_LEFT_PAR().GetLine()
			}
		}()), localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)
			p.Match(ProjectParserPRINTLN)
		}
		{
			p.SetState(76)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(77)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(78)

			var _x = p.Expression()

			localctx.(*Print_prodContext)._expression = _x
		}
		{
			p.SetState(79)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(80)
			p.Match(ProjectParserSEMICOLON)
		}
		localctx.(*Print_prodContext).instr = Natives.NewPrint(localctx.(*Print_prodContext).Get_expression().GetP(), true, (func() int {
			if localctx.(*Print_prodContext).Get_LEFT_PAR() == nil {
				return 0
			} else {
				return localctx.(*Print_prodContext).Get_LEFT_PAR().GetLine()
			}
		}()), localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Match(ProjectParserPRINTLN)
		}
		{
			p.SetState(84)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(85)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(86)

			var _x = p.Expression()

			localctx.(*Print_prodContext).opBefore = _x
		}
		{
			p.SetState(87)
			p.Match(ProjectParserCOMMA)
		}
		{
			p.SetState(88)

			var _x = p.listVars(0)

			localctx.(*Print_prodContext)._listVars = _x
		}
		{
			p.SetState(89)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(90)
			p.Match(ProjectParserSEMICOLON)
		}
		localctx.(*Print_prodContext).instr = Natives.NewPrintWithAfter(localctx.(*Print_prodContext).GetOpBefore().GetP(), localctx.(*Print_prodContext).Get_listVars().GetList(), true, (func() int {
			if localctx.(*Print_prodContext).Get_LEFT_PAR() == nil {
				return 0
			} else {
				return localctx.(*Print_prodContext).Get_LEFT_PAR().GetLine()
			}
		}()), localctx.(*Print_prodContext).Get_LEFT_PAR().GetColumn())

	}

	return localctx
}

// IListVarsContext is an interface to support dynamic dispatch.
type IListVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSub returns the sub rule contexts.
	GetSub() IListVarsContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetSub sets the sub rule contexts.
	SetSub(IListVarsContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetList returns the list attribute.
	GetList() *arrayList.List

	// SetList sets the list attribute.
	SetList(*arrayList.List)

	// IsListVarsContext differentiates from other interfaces.
	IsListVarsContext()
}

type ListVarsContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	list        *arrayList.List
	sub         IListVarsContext
	_expression IExpressionContext
}

func NewEmptyListVarsContext() *ListVarsContext {
	var p = new(ListVarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_listVars
	return p
}

func (*ListVarsContext) IsListVarsContext() {}

func NewListVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListVarsContext {
	var p = new(ListVarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_listVars

	return p
}

func (s *ListVarsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListVarsContext) GetSub() IListVarsContext { return s.sub }

func (s *ListVarsContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListVarsContext) SetSub(v IListVarsContext) { s.sub = v }

func (s *ListVarsContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListVarsContext) GetList() *arrayList.List { return s.list }

func (s *ListVarsContext) SetList(v *arrayList.List) { s.list = v }

func (s *ListVarsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListVarsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *ListVarsContext) ListVars() IListVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListVarsContext)
}

func (s *ListVarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListVarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListVarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterListVars(s)
	}
}

func (s *ListVarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitListVars(s)
	}
}

func (p *ProjectParser) ListVars() (localctx IListVarsContext) {
	return p.listVars(0)
}

func (p *ProjectParser) listVars(_p int) (localctx IListVarsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListVarsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListVarsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, ProjectParserRULE_listVars, _p)

	localctx.(*ListVarsContext).list = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)

		var _x = p.Expression()

		localctx.(*ListVarsContext)._expression = _x
	}

	localctx.(*ListVarsContext).list.Add(localctx.(*ListVarsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListVarsContext(p, _parentctx, _parentState)
			localctx.(*ListVarsContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listVars)
			p.SetState(99)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(100)
				p.Match(ProjectParserCOMMA)
			}
			{
				p.SetState(101)

				var _x = p.Expression()

				localctx.(*ListVarsContext)._expression = _x
			}

			localctx.(*ListVarsContext).GetSub().GetList().Add(localctx.(*ListVarsContext).Get_expression().GetP())
			localctx.(*ListVarsContext).list = localctx.(*ListVarsContext).GetSub().GetList()

		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclaration_prodContext is an interface to support dynamic dispatch.
type IDeclaration_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ids_type returns the _ids_type rule contexts.
	Get_ids_type() IIds_typeContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_ids_type sets the _ids_type rule contexts.
	Set_ids_type(IIds_typeContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsDeclaration_prodContext differentiates from other interfaces.
	IsDeclaration_prodContext()
}

type Declaration_prodContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_ids_type   IIds_typeContext
	_expression IExpressionContext
}

func NewEmptyDeclaration_prodContext() *Declaration_prodContext {
	var p = new(Declaration_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_declaration_prod
	return p
}

func (*Declaration_prodContext) IsDeclaration_prodContext() {}

func NewDeclaration_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaration_prodContext {
	var p = new(Declaration_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_declaration_prod

	return p
}

func (s *Declaration_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaration_prodContext) Get_ids_type() IIds_typeContext { return s._ids_type }

func (s *Declaration_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Declaration_prodContext) Set_ids_type(v IIds_typeContext) { s._ids_type = v }

func (s *Declaration_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Declaration_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Declaration_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Declaration_prodContext) DECLARAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserDECLARAR, 0)
}

func (s *Declaration_prodContext) Ids_type() IIds_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIds_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIds_typeContext)
}

func (s *Declaration_prodContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ProjectParserEQUAL, 0)
}

func (s *Declaration_prodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Declaration_prodContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *Declaration_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaration_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaration_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterDeclaration_prod(s)
	}
}

func (s *Declaration_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitDeclaration_prod(s)
	}
}

func (p *ProjectParser) Declaration_prod() (localctx IDeclaration_prodContext) {
	localctx = NewDeclaration_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ProjectParserRULE_declaration_prod)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Match(ProjectParserDECLARAR)
		}
		{
			p.SetState(110)

			var _x = p.Ids_type()

			localctx.(*Declaration_prodContext)._ids_type = _x
		}
		{
			p.SetState(111)
			p.Match(ProjectParserEQUAL)
		}
		{
			p.SetState(112)

			var _x = p.Expression()

			localctx.(*Declaration_prodContext)._expression = _x
		}
		{
			p.SetState(113)
			p.Match(ProjectParserSEMICOLON)
		}

		localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Match(ProjectParserDECLARAR)
		}
		{
			p.SetState(117)

			var _x = p.Ids_type()

			localctx.(*Declaration_prodContext)._ids_type = _x
		}
		{
			p.SetState(118)
			p.Match(ProjectParserSEMICOLON)
		}

		localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList())

	}

	return localctx
}

// IAssign_prodContext is an interface to support dynamic dispatch.
type IAssign_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listIds returns the _listIds rule contexts.
	Get_listIds() IListIdsContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_listIds sets the _listIds rule contexts.
	Set_listIds(IListIdsContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsAssign_prodContext differentiates from other interfaces.
	IsAssign_prodContext()
}

type Assign_prodContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_listIds    IListIdsContext
	_expression IExpressionContext
}

func NewEmptyAssign_prodContext() *Assign_prodContext {
	var p = new(Assign_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_assign_prod
	return p
}

func (*Assign_prodContext) IsAssign_prodContext() {}

func NewAssign_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_prodContext {
	var p = new(Assign_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_assign_prod

	return p
}

func (s *Assign_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_prodContext) Get_listIds() IListIdsContext { return s._listIds }

func (s *Assign_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Assign_prodContext) Set_listIds(v IListIdsContext) { s._listIds = v }

func (s *Assign_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Assign_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Assign_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Assign_prodContext) ListIds() IListIdsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListIdsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListIdsContext)
}

func (s *Assign_prodContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ProjectParserEQUAL, 0)
}

func (s *Assign_prodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assign_prodContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *Assign_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterAssign_prod(s)
	}
}

func (s *Assign_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitAssign_prod(s)
	}
}

func (p *ProjectParser) Assign_prod() (localctx IAssign_prodContext) {
	localctx = NewAssign_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ProjectParserRULE_assign_prod)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)

		var _x = p.listIds(0)

		localctx.(*Assign_prodContext)._listIds = _x
	}
	{
		p.SetState(124)
		p.Match(ProjectParserEQUAL)
	}
	{
		p.SetState(125)

		var _x = p.Expression()

		localctx.(*Assign_prodContext)._expression = _x
	}
	{
		p.SetState(126)
		p.Match(ProjectParserSEMICOLON)
	}

	localctx.(*Assign_prodContext).instr = Natives.NewAssign(localctx.(*Assign_prodContext).Get_listIds().GetList(), localctx.(*Assign_prodContext).Get_expression().GetP())

	return localctx
}

// IIds_typeContext is an interface to support dynamic dispatch.
type IIds_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listIds returns the _listIds rule contexts.
	Get_listIds() IListIdsContext

	// Set_listIds sets the _listIds rule contexts.
	Set_listIds(IListIdsContext)

	// GetList returns the list attribute.
	GetList() *arrayList.List

	// SetList sets the list attribute.
	SetList(*arrayList.List)

	// IsIds_typeContext differentiates from other interfaces.
	IsIds_typeContext()
}

type Ids_typeContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	list     *arrayList.List
	_listIds IListIdsContext
}

func NewEmptyIds_typeContext() *Ids_typeContext {
	var p = new(Ids_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_ids_type
	return p
}

func (*Ids_typeContext) IsIds_typeContext() {}

func NewIds_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ids_typeContext {
	var p = new(Ids_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_ids_type

	return p
}

func (s *Ids_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Ids_typeContext) Get_listIds() IListIdsContext { return s._listIds }

func (s *Ids_typeContext) Set_listIds(v IListIdsContext) { s._listIds = v }

func (s *Ids_typeContext) GetList() *arrayList.List { return s.list }

func (s *Ids_typeContext) SetList(v *arrayList.List) { s.list = v }

func (s *Ids_typeContext) ListIds() IListIdsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListIdsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListIdsContext)
}

func (s *Ids_typeContext) COLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOLON, 0)
}

func (s *Ids_typeContext) RSTRING() antlr.TerminalNode {
	return s.GetToken(ProjectParserRSTRING, 0)
}

func (s *Ids_typeContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(ProjectParserREFERENCE, 0)
}

func (s *Ids_typeContext) RSTR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRSTR, 0)
}

func (s *Ids_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ids_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ids_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterIds_type(s)
	}
}

func (s *Ids_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitIds_type(s)
	}
}

func (p *ProjectParser) Ids_type() (localctx IIds_typeContext) {
	localctx = NewIds_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ProjectParserRULE_ids_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		localctx.(*Ids_typeContext).list = localctx.(*Ids_typeContext).Get_listIds().GetList()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		{
			p.SetState(133)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(134)
			p.Match(ProjectParserRSTRING)
		}
		localctx.(*Ids_typeContext).list = localctx.(*Ids_typeContext).Get_listIds().GetList()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		{
			p.SetState(138)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(139)
			p.Match(ProjectParserREFERENCE)
		}
		{
			p.SetState(140)
			p.Match(ProjectParserRSTR)
		}
		localctx.(*Ids_typeContext).list = localctx.(*Ids_typeContext).Get_listIds().GetList()

	}

	return localctx
}

// IListIdsContext is an interface to support dynamic dispatch.
type IListIdsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetSub returns the sub rule contexts.
	GetSub() IListIdsContext

	// SetSub sets the sub rule contexts.
	SetSub(IListIdsContext)

	// GetList returns the list attribute.
	GetList() *arrayList.List

	// SetList sets the list attribute.
	SetList(*arrayList.List)

	// IsListIdsContext differentiates from other interfaces.
	IsListIdsContext()
}

type ListIdsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	list   *arrayList.List
	sub    IListIdsContext
	_ID    antlr.Token
}

func NewEmptyListIdsContext() *ListIdsContext {
	var p = new(ListIdsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_listIds
	return p
}

func (*ListIdsContext) IsListIdsContext() {}

func NewListIdsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIdsContext {
	var p = new(ListIdsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_listIds

	return p
}

func (s *ListIdsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIdsContext) Get_ID() antlr.Token { return s._ID }

func (s *ListIdsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListIdsContext) GetSub() IListIdsContext { return s.sub }

func (s *ListIdsContext) SetSub(v IListIdsContext) { s.sub = v }

func (s *ListIdsContext) GetList() *arrayList.List { return s.list }

func (s *ListIdsContext) SetList(v *arrayList.List) { s.list = v }

func (s *ListIdsContext) ID() antlr.TerminalNode {
	return s.GetToken(ProjectParserID, 0)
}

func (s *ListIdsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *ListIdsContext) ListIds() IListIdsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListIdsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListIdsContext)
}

func (s *ListIdsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIdsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIdsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterListIds(s)
	}
}

func (s *ListIdsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitListIds(s)
	}
}

func (p *ProjectParser) ListIds() (localctx IListIdsContext) {
	return p.listIds(0)
}

func (p *ProjectParser) listIds(_p int) (localctx IListIdsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListIdsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListIdsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, ProjectParserRULE_listIds, _p)

	localctx.(*ListIdsContext).list = arrayList.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)

		var _m = p.Match(ProjectParserID)

		localctx.(*ListIdsContext)._ID = _m
	}
	localctx.(*ListIdsContext).list.Add(Expression.NewIdentifier((func() string {
		if localctx.(*ListIdsContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListIdsContext).Get_ID().GetText()
		}
	}()), (func() int {
		if localctx.(*ListIdsContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListIdsContext).Get_ID().GetLine()
		}
	}()), localctx.(*ListIdsContext).Get_ID().GetColumn()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListIdsContext(p, _parentctx, _parentState)
			localctx.(*ListIdsContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listIds)
			p.SetState(149)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(150)
				p.Match(ProjectParserCOMMA)
			}
			{
				p.SetState(151)

				var _m = p.Match(ProjectParserID)

				localctx.(*ListIdsContext)._ID = _m
			}

			localctx.(*ListIdsContext).GetSub().GetList().Add(Expression.NewIdentifier((func() string {
				if localctx.(*ListIdsContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListIdsContext).Get_ID().GetText()
				}
			}()), (func() int {
				if localctx.(*ListIdsContext).Get_ID() == nil {
					return 0
				} else {
					return localctx.(*ListIdsContext).Get_ID().GetLine()
				}
			}()), localctx.(*ListIdsContext).Get_ID().GetColumn()))
			localctx.(*ListIdsContext).list = localctx.(*ListIdsContext).GetSub().GetList()

		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IConditional_prodContext is an interface to support dynamic dispatch.
type IConditional_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_bloq returns the _bloq rule contexts.
	Get_bloq() IBloqContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_bloq sets the _bloq rule contexts.
	Set_bloq(IBloqContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsConditional_prodContext differentiates from other interfaces.
	IsConditional_prodContext()
}

type Conditional_prodContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_expression IExpressionContext
	_bloq       IBloqContext
}

func NewEmptyConditional_prodContext() *Conditional_prodContext {
	var p = new(Conditional_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_conditional_prod
	return p
}

func (*Conditional_prodContext) IsConditional_prodContext() {}

func NewConditional_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Conditional_prodContext {
	var p = new(Conditional_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_conditional_prod

	return p
}

func (s *Conditional_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Conditional_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Conditional_prodContext) Get_bloq() IBloqContext { return s._bloq }

func (s *Conditional_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Conditional_prodContext) Set_bloq(v IBloqContext) { s._bloq = v }

func (s *Conditional_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Conditional_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Conditional_prodContext) RIF() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIF, 0)
}

func (s *Conditional_prodContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *Conditional_prodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Conditional_prodContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_PAR, 0)
}

func (s *Conditional_prodContext) Bloq() IBloqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqContext)
}

func (s *Conditional_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Conditional_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Conditional_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterConditional_prod(s)
	}
}

func (s *Conditional_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitConditional_prod(s)
	}
}

func (p *ProjectParser) Conditional_prod() (localctx IConditional_prodContext) {
	localctx = NewConditional_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ProjectParserRULE_conditional_prod)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(ProjectParserRIF)
	}
	{
		p.SetState(159)
		p.Match(ProjectParserLEFT_PAR)
	}
	{
		p.SetState(160)

		var _x = p.Expression()

		localctx.(*Conditional_prodContext)._expression = _x
	}
	{
		p.SetState(161)
		p.Match(ProjectParserRIGHT_PAR)
	}
	{
		p.SetState(162)

		var _x = p.Bloq()

		localctx.(*Conditional_prodContext)._bloq = _x
	}
	localctx.(*Conditional_prodContext).instr = Natives.NewIf(localctx.(*Conditional_prodContext).Get_expression().GetP(), localctx.(*Conditional_prodContext).Get_bloq().GetContent(), nil, nil)

	return localctx
}

// IBloqContext is an interface to support dynamic dispatch.
type IBloqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instructions returns the _instructions rule contexts.
	Get_instructions() IInstructionsContext

	// Set_instructions sets the _instructions rule contexts.
	Set_instructions(IInstructionsContext)

	// GetContent returns the content attribute.
	GetContent() *arrayList.List

	// SetContent sets the content attribute.
	SetContent(*arrayList.List)

	// IsBloqContext differentiates from other interfaces.
	IsBloqContext()
}

type BloqContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	content       *arrayList.List
	_instructions IInstructionsContext
}

func NewEmptyBloqContext() *BloqContext {
	var p = new(BloqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_bloq
	return p
}

func (*BloqContext) IsBloqContext() {}

func NewBloqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqContext {
	var p = new(BloqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_bloq

	return p
}

func (s *BloqContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *BloqContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *BloqContext) GetContent() *arrayList.List { return s.content }

func (s *BloqContext) SetContent(v *arrayList.List) { s.content = v }

func (s *BloqContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_BRACKET, 0)
}

func (s *BloqContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *BloqContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_BRACKET, 0)
}

func (s *BloqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterBloq(s)
	}
}

func (s *BloqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitBloq(s)
	}
}

func (p *ProjectParser) Bloq() (localctx IBloqContext) {
	localctx = NewBloqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ProjectParserRULE_bloq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(ProjectParserLEFT_BRACKET)
		}
		{
			p.SetState(166)

			var _x = p.Instructions()

			localctx.(*BloqContext)._instructions = _x
		}
		{
			p.SetState(167)
			p.Match(ProjectParserRIGHT_BRACKET)
		}
		localctx.(*BloqContext).content = localctx.(*BloqContext).Get_instructions().GetL()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Match(ProjectParserLEFT_BRACKET)
		}
		{
			p.SetState(171)
			p.Match(ProjectParserRIGHT_BRACKET)
		}
		localctx.(*BloqContext).content = arrayList.New()

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr_rel returns the _expr_rel rule contexts.
	Get_expr_rel() IExpr_relContext

	// Get_expr_arit returns the _expr_arit rule contexts.
	Get_expr_arit() IExpr_aritContext

	// Get_expr_logic returns the _expr_logic rule contexts.
	Get_expr_logic() IExpr_logicContext

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// Set_expr_logic sets the _expr_logic rule contexts.
	Set_expr_logic(IExpr_logicContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           Abstract.Expression
	_expr_rel   IExpr_relContext
	_expr_arit  IExpr_aritContext
	_expr_logic IExpr_logicContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Get_expr_rel() IExpr_relContext { return s._expr_rel }

func (s *ExpressionContext) Get_expr_arit() IExpr_aritContext { return s._expr_arit }

func (s *ExpressionContext) Get_expr_logic() IExpr_logicContext { return s._expr_logic }

func (s *ExpressionContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *ExpressionContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *ExpressionContext) Set_expr_logic(v IExpr_logicContext) { s._expr_logic = v }

func (s *ExpressionContext) GetP() Abstract.Expression { return s.p }

func (s *ExpressionContext) SetP(v Abstract.Expression) { s.p = v }

func (s *ExpressionContext) Expr_rel() IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *ExpressionContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *ExpressionContext) Expr_logic() IExpr_logicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_logicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_logicContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ProjectParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ProjectParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)

			var _x = p.expr_rel(0)

			localctx.(*ExpressionContext)._expr_rel = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_rel().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(178)

			var _x = p.expr_arit(0)

			localctx.(*ExpressionContext)._expr_arit = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)

			var _x = p.Expr_logic()

			localctx.(*ExpressionContext)._expr_logic = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_logic().GetP()

	}

	return localctx
}

// IExpr_relContext is an interface to support dynamic dispatch.
type IExpr_relContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpLeft returns the opLeft rule contexts.
	GetOpLeft() IExpr_relContext

	// Get_expr_arit returns the _expr_arit rule contexts.
	Get_expr_arit() IExpr_aritContext

	// GetOpRight returns the opRight rule contexts.
	GetOpRight() IExpr_relContext

	// SetOpLeft sets the opLeft rule contexts.
	SetOpLeft(IExpr_relContext)

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// SetOpRight sets the opRight rule contexts.
	SetOpRight(IExpr_relContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsExpr_relContext differentiates from other interfaces.
	IsExpr_relContext()
}

type Expr_relContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          Abstract.Expression
	opLeft     IExpr_relContext
	_expr_arit IExpr_aritContext
	op         antlr.Token
	opRight    IExpr_relContext
}

func NewEmptyExpr_relContext() *Expr_relContext {
	var p = new(Expr_relContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_expr_rel
	return p
}

func (*Expr_relContext) IsExpr_relContext() {}

func NewExpr_relContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_relContext {
	var p = new(Expr_relContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_expr_rel

	return p
}

func (s *Expr_relContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_relContext) GetOp() antlr.Token { return s.op }

func (s *Expr_relContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_relContext) GetOpLeft() IExpr_relContext { return s.opLeft }

func (s *Expr_relContext) Get_expr_arit() IExpr_aritContext { return s._expr_arit }

func (s *Expr_relContext) GetOpRight() IExpr_relContext { return s.opRight }

func (s *Expr_relContext) SetOpLeft(v IExpr_relContext) { s.opLeft = v }

func (s *Expr_relContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *Expr_relContext) SetOpRight(v IExpr_relContext) { s.opRight = v }

func (s *Expr_relContext) GetP() Abstract.Expression { return s.p }

func (s *Expr_relContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Expr_relContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *Expr_relContext) AllExpr_rel() []IExpr_relContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_relContext)(nil)).Elem())
	var tst = make([]IExpr_relContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_relContext)
		}
	}

	return tst
}

func (s *Expr_relContext) Expr_rel(i int) IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *Expr_relContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(ProjectParserGREATER_THAN, 0)
}

func (s *Expr_relContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(ProjectParserLESS_THAN, 0)
}

func (s *Expr_relContext) GREATER_EQUALTHAN() antlr.TerminalNode {
	return s.GetToken(ProjectParserGREATER_EQUALTHAN, 0)
}

func (s *Expr_relContext) LESS_EQUEALTHAN() antlr.TerminalNode {
	return s.GetToken(ProjectParserLESS_EQUEALTHAN, 0)
}

func (s *Expr_relContext) EQUEAL_EQUAL() antlr.TerminalNode {
	return s.GetToken(ProjectParserEQUEAL_EQUAL, 0)
}

func (s *Expr_relContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_relContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_relContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterExpr_rel(s)
	}
}

func (s *Expr_relContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitExpr_rel(s)
	}
}

func (p *ProjectParser) Expr_rel() (localctx IExpr_relContext) {
	return p.expr_rel(0)
}

func (p *ProjectParser) expr_rel(_p int) (localctx IExpr_relContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_relContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_relContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ProjectParserRULE_expr_rel, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)

		var _x = p.expr_arit(0)

		localctx.(*Expr_relContext)._expr_arit = _x
	}
	localctx.(*Expr_relContext).p = localctx.(*Expr_relContext).Get_expr_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_relContext(p, _parentctx, _parentState)
			localctx.(*Expr_relContext).opLeft = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_expr_rel)
			p.SetState(190)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(191)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expr_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(ProjectParserEQUEAL_EQUAL-34))|(1<<(ProjectParserGREATER_THAN-34))|(1<<(ProjectParserLESS_THAN-34))|(1<<(ProjectParserGREATER_EQUALTHAN-34))|(1<<(ProjectParserLESS_EQUEALTHAN-34)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expr_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(192)

				var _x = p.expr_rel(3)

				localctx.(*Expr_relContext).opRight = _x
			}
			localctx.(*Expr_relContext).p = Expression.NewOperation(localctx.(*Expr_relContext).GetOpLeft().GetP(), (func() string {
				if localctx.(*Expr_relContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expr_relContext).GetOp().GetText()
				}
			}()), localctx.(*Expr_relContext).GetOpRight().GetP(), false, (func() int {
				if localctx.(*Expr_relContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*Expr_relContext).GetOp().GetLine()
				}
			}()), localctx.(*Expr_relContext).GetOp().GetColumn())

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr_logicContext is an interface to support dynamic dispatch.
type IExpr_logicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpLeft returns the opLeft rule contexts.
	GetOpLeft() IExpr_relContext

	// GetOpRight returns the opRight rule contexts.
	GetOpRight() IExpr_relContext

	// SetOpLeft sets the opLeft rule contexts.
	SetOpLeft(IExpr_relContext)

	// SetOpRight sets the opRight rule contexts.
	SetOpRight(IExpr_relContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsExpr_logicContext differentiates from other interfaces.
	IsExpr_logicContext()
}

type Expr_logicContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	p       Abstract.Expression
	opLeft  IExpr_relContext
	op      antlr.Token
	opRight IExpr_relContext
}

func NewEmptyExpr_logicContext() *Expr_logicContext {
	var p = new(Expr_logicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_expr_logic
	return p
}

func (*Expr_logicContext) IsExpr_logicContext() {}

func NewExpr_logicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_logicContext {
	var p = new(Expr_logicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_expr_logic

	return p
}

func (s *Expr_logicContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_logicContext) GetOp() antlr.Token { return s.op }

func (s *Expr_logicContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_logicContext) GetOpLeft() IExpr_relContext { return s.opLeft }

func (s *Expr_logicContext) GetOpRight() IExpr_relContext { return s.opRight }

func (s *Expr_logicContext) SetOpLeft(v IExpr_relContext) { s.opLeft = v }

func (s *Expr_logicContext) SetOpRight(v IExpr_relContext) { s.opRight = v }

func (s *Expr_logicContext) GetP() Abstract.Expression { return s.p }

func (s *Expr_logicContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Expr_logicContext) AllExpr_rel() []IExpr_relContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_relContext)(nil)).Elem())
	var tst = make([]IExpr_relContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_relContext)
		}
	}

	return tst
}

func (s *Expr_logicContext) Expr_rel(i int) IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *Expr_logicContext) RAND() antlr.TerminalNode {
	return s.GetToken(ProjectParserRAND, 0)
}

func (s *Expr_logicContext) ROR() antlr.TerminalNode {
	return s.GetToken(ProjectParserROR, 0)
}

func (s *Expr_logicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_logicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_logicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterExpr_logic(s)
	}
}

func (s *Expr_logicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitExpr_logic(s)
	}
}

func (p *ProjectParser) Expr_logic() (localctx IExpr_logicContext) {
	localctx = NewExpr_logicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ProjectParserRULE_expr_logic)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)

		var _x = p.expr_rel(0)

		localctx.(*Expr_logicContext).opLeft = _x
	}
	{
		p.SetState(201)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Expr_logicContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ProjectParserRAND || _la == ProjectParserROR) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Expr_logicContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(202)

		var _x = p.expr_rel(0)

		localctx.(*Expr_logicContext).opRight = _x
	}
	localctx.(*Expr_logicContext).p = Expression.NewOperation(localctx.(*Expr_logicContext).GetOpLeft().GetP(), (func() string {
		if localctx.(*Expr_logicContext).GetOp() == nil {
			return ""
		} else {
			return localctx.(*Expr_logicContext).GetOp().GetText()
		}
	}()), localctx.(*Expr_logicContext).GetOpRight().GetP(), false, (func() int {
		if localctx.(*Expr_logicContext).GetOp() == nil {
			return 0
		} else {
			return localctx.(*Expr_logicContext).GetOp().GetLine()
		}
	}()), localctx.(*Expr_logicContext).GetOp().GetColumn())

	return localctx
}

// IExpr_aritContext is an interface to support dynamic dispatch.
type IExpr_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetOpLeft returns the opLeft rule contexts.
	GetOpLeft() IExpr_aritContext

	// GetOpU returns the opU rule contexts.
	GetOpU() IExpressionContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_primitive returns the _primitive rule contexts.
	Get_primitive() IPrimitiveContext

	// GetOpRight returns the opRight rule contexts.
	GetOpRight() IExpr_aritContext

	// SetOpLeft sets the opLeft rule contexts.
	SetOpLeft(IExpr_aritContext)

	// SetOpU sets the opU rule contexts.
	SetOpU(IExpressionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_primitive sets the _primitive rule contexts.
	Set_primitive(IPrimitiveContext)

	// SetOpRight sets the opRight rule contexts.
	SetOpRight(IExpr_aritContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsExpr_aritContext differentiates from other interfaces.
	IsExpr_aritContext()
}

type Expr_aritContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           Abstract.Expression
	opLeft      IExpr_aritContext
	opU         IExpressionContext
	_expression IExpressionContext
	_primitive  IPrimitiveContext
	op          antlr.Token
	opRight     IExpr_aritContext
}

func NewEmptyExpr_aritContext() *Expr_aritContext {
	var p = new(Expr_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_expr_arit
	return p
}

func (*Expr_aritContext) IsExpr_aritContext() {}

func NewExpr_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_aritContext {
	var p = new(Expr_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_expr_arit

	return p
}

func (s *Expr_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expr_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_aritContext) GetOpLeft() IExpr_aritContext { return s.opLeft }

func (s *Expr_aritContext) GetOpU() IExpressionContext { return s.opU }

func (s *Expr_aritContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_aritContext) Get_primitive() IPrimitiveContext { return s._primitive }

func (s *Expr_aritContext) GetOpRight() IExpr_aritContext { return s.opRight }

func (s *Expr_aritContext) SetOpLeft(v IExpr_aritContext) { s.opLeft = v }

func (s *Expr_aritContext) SetOpU(v IExpressionContext) { s.opU = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_primitive(v IPrimitiveContext) { s._primitive = v }

func (s *Expr_aritContext) SetOpRight(v IExpr_aritContext) { s.opRight = v }

func (s *Expr_aritContext) GetP() Abstract.Expression { return s.p }

func (s *Expr_aritContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Expr_aritContext) SUB() antlr.TerminalNode {
	return s.GetToken(ProjectParserSUB, 0)
}

func (s *Expr_aritContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expr_aritContext) Primitive() IPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *Expr_aritContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *Expr_aritContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_PAR, 0)
}

func (s *Expr_aritContext) AllExpr_arit() []IExpr_aritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem())
	var tst = make([]IExpr_aritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_aritContext)
		}
	}

	return tst
}

func (s *Expr_aritContext) Expr_arit(i int) IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *Expr_aritContext) MUL() antlr.TerminalNode {
	return s.GetToken(ProjectParserMUL, 0)
}

func (s *Expr_aritContext) DIV() antlr.TerminalNode {
	return s.GetToken(ProjectParserDIV, 0)
}

func (s *Expr_aritContext) MOD() antlr.TerminalNode {
	return s.GetToken(ProjectParserMOD, 0)
}

func (s *Expr_aritContext) ADD() antlr.TerminalNode {
	return s.GetToken(ProjectParserADD, 0)
}

func (s *Expr_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterExpr_arit(s)
	}
}

func (s *Expr_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitExpr_arit(s)
	}
}

func (p *ProjectParser) Expr_arit() (localctx IExpr_aritContext) {
	return p.expr_arit(0)
}

func (p *ProjectParser) expr_arit(_p int) (localctx IExpr_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, ProjectParserRULE_expr_arit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserSUB:
		{
			p.SetState(206)
			p.Match(ProjectParserSUB)
		}
		{
			p.SetState(207)

			var _x = p.Expression()

			localctx.(*Expr_aritContext).opU = _x
			localctx.(*Expr_aritContext)._expression = _x
		}
		localctx.(*Expr_aritContext).p = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpU().GetP(), "-", nil, true, localctx.(*Expr_aritContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_aritContext).GetOpU().GetStart().GetColumn())

	case ProjectParserINTEGER, ProjectParserFLOAT, ProjectParserCHAR, ProjectParserSTRING, ProjectParserBOOLEAN, ProjectParserID:
		{
			p.SetState(210)

			var _x = p.Primitive()

			localctx.(*Expr_aritContext)._primitive = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitive().GetP()

	case ProjectParserLEFT_PAR:
		{
			p.SetState(213)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(214)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(215)
			p.Match(ProjectParserRIGHT_PAR)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opLeft = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_expr_arit)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(221)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ProjectParserMUL-40))|(1<<(ProjectParserDIV-40))|(1<<(ProjectParserMOD-40)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(222)

					var _x = p.expr_arit(5)

					localctx.(*Expr_aritContext).opRight = _x
				}
				localctx.(*Expr_aritContext).p = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpRight().GetP(), false, (func() int {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expr_aritContext).GetOp().GetColumn())

			case 2:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opLeft = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_expr_arit)
				p.SetState(225)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(226)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ProjectParserADD || _la == ProjectParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(227)

					var _x = p.expr_arit(4)

					localctx.(*Expr_aritContext).opRight = _x
				}
				localctx.(*Expr_aritContext).p = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpRight().GetP(), false, (func() int {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expr_aritContext).GetOp().GetColumn())

			}

		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitiveContext is an interface to support dynamic dispatch.
type IPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INTEGER returns the _INTEGER token.
	Get_INTEGER() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_CHAR returns the _CHAR token.
	Get_CHAR() antlr.Token

	// Get_BOOLEAN returns the _BOOLEAN token.
	Get_BOOLEAN() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_INTEGER sets the _INTEGER token.
	Set_INTEGER(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsPrimitiveContext differentiates from other interfaces.
	IsPrimitiveContext()
}

type PrimitiveContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        Abstract.Expression
	_INTEGER antlr.Token
	_FLOAT   antlr.Token
	_STRING  antlr.Token
	_CHAR    antlr.Token
	_BOOLEAN antlr.Token
	_ID      antlr.Token
}

func NewEmptyPrimitiveContext() *PrimitiveContext {
	var p = new(PrimitiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_primitive
	return p
}

func (*PrimitiveContext) IsPrimitiveContext() {}

func NewPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveContext {
	var p = new(PrimitiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_primitive

	return p
}

func (s *PrimitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveContext) Get_INTEGER() antlr.Token { return s._INTEGER }

func (s *PrimitiveContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *PrimitiveContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitiveContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *PrimitiveContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *PrimitiveContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitiveContext) Set_INTEGER(v antlr.Token) { s._INTEGER = v }

func (s *PrimitiveContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *PrimitiveContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitiveContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *PrimitiveContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *PrimitiveContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitiveContext) GetP() Abstract.Expression { return s.p }

func (s *PrimitiveContext) SetP(v Abstract.Expression) { s.p = v }

func (s *PrimitiveContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ProjectParserINTEGER, 0)
}

func (s *PrimitiveContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ProjectParserFLOAT, 0)
}

func (s *PrimitiveContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProjectParserSTRING, 0)
}

func (s *PrimitiveContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserCHAR, 0)
}

func (s *PrimitiveContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ProjectParserBOOLEAN, 0)
}

func (s *PrimitiveContext) ID() antlr.TerminalNode {
	return s.GetToken(ProjectParserID, 0)
}

func (s *PrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterPrimitive(s)
	}
}

func (s *PrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitPrimitive(s)
	}
}

func (p *ProjectParser) Primitive() (localctx IPrimitiveContext) {
	localctx = NewPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ProjectParserRULE_primitive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)

			var _m = p.Match(ProjectParserINTEGER)

			localctx.(*PrimitiveContext)._INTEGER = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitiveContext).Get_INTEGER() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_INTEGER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitiveContext).p = Expression.NewPrimitive(num, SymbolTable.INTEGER, (func() int {
			if localctx.(*PrimitiveContext).Get_INTEGER() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_INTEGER().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_INTEGER().GetColumn())

	case ProjectParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(237)

			var _m = p.Match(ProjectParserFLOAT)

			localctx.(*PrimitiveContext)._FLOAT = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitiveContext).Get_FLOAT() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_FLOAT().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitiveContext).p = Expression.NewPrimitive(num, SymbolTable.FLOAT, (func() int {
			if localctx.(*PrimitiveContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_FLOAT().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_FLOAT().GetColumn())

	case ProjectParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)

			var _m = p.Match(ProjectParserSTRING)

			localctx.(*PrimitiveContext)._STRING = _m
		}

		str := (func() string {
			if localctx.(*PrimitiveContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitiveContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*PrimitiveContext).p = Expression.NewPrimitive(str, SymbolTable.STR, (func() int {
			if localctx.(*PrimitiveContext).Get_STRING() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_STRING().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_STRING().GetColumn())

	case ProjectParserCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(241)

			var _m = p.Match(ProjectParserCHAR)

			localctx.(*PrimitiveContext)._CHAR = _m
		}

		chr := (func() string {
			if localctx.(*PrimitiveContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_CHAR().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitiveContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_CHAR().GetText()
			}
		}()))-1]
		localctx.(*PrimitiveContext).p = Expression.NewPrimitive(chr, SymbolTable.CHAR, (func() int {
			if localctx.(*PrimitiveContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_CHAR().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_CHAR().GetColumn())

	case ProjectParserBOOLEAN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(243)

			var _m = p.Match(ProjectParserBOOLEAN)

			localctx.(*PrimitiveContext)._BOOLEAN = _m
		}

		str := (func() string {
			if localctx.(*PrimitiveContext).Get_BOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_BOOLEAN().GetText()
			}
		}())
		fmt.Println(str)
		localctx.(*PrimitiveContext).p = Expression.NewPrimitive(str, SymbolTable.BOOLEAN, (func() int {
			if localctx.(*PrimitiveContext).Get_BOOLEAN() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_BOOLEAN().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())

	case ProjectParserID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(245)

			var _m = p.Match(ProjectParserID)

			localctx.(*PrimitiveContext)._ID = _m
		}
		localctx.(*PrimitiveContext).p = Expression.NewIdentifier((func() string {
			if localctx.(*PrimitiveContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_ID().GetText()
			}
		}()), (func() int {
			if localctx.(*PrimitiveContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_ID().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_ID().GetColumn())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *ProjectParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *ListVarsContext = nil
		if localctx != nil {
			t = localctx.(*ListVarsContext)
		}
		return p.ListVars_Sempred(t, predIndex)

	case 8:
		var t *ListIdsContext = nil
		if localctx != nil {
			t = localctx.(*ListIdsContext)
		}
		return p.ListIds_Sempred(t, predIndex)

	case 12:
		var t *Expr_relContext = nil
		if localctx != nil {
			t = localctx.(*Expr_relContext)
		}
		return p.Expr_rel_Sempred(t, predIndex)

	case 14:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ProjectParser) ListVars_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) ListIds_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) Expr_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
