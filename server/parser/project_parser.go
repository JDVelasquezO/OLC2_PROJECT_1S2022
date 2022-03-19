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
import "OLC2_Project1/server/interpreter/SymbolTable/Environment"
import "OLC2_Project1/server/interpreter/AST/ExpressionSpecial"
import "OLC2_Project1/server/interpreter/AST/Natives/DecArrays"
import "OLC2_Project1/server/interpreter/AST/Expression/Access"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 70, 774,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 112, 10, 3, 12, 3, 14, 3, 115, 11, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 154, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 171, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 181, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 196, 10, 8, 12, 8, 14, 8,
	199, 11, 8, 3, 9, 7, 9, 202, 10, 9, 12, 9, 14, 9, 205, 11, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 226, 10, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 236, 10, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 274, 10, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 285, 10, 12, 12, 12, 14,
	12, 288, 11, 12, 3, 13, 3, 13, 5, 13, 292, 10, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 5, 13, 298, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 304, 10,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 312, 10, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 5, 13, 318, 10, 13, 3, 13, 3, 13, 5, 13, 322, 10,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 328, 10, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 5, 13, 334, 10, 13, 3, 13, 3, 13, 5, 13, 338, 10, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 5, 14, 344, 10, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 5, 14, 353, 10, 14, 3, 14, 3, 14, 5, 14, 357, 10, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 373, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 7, 16, 383, 10, 16, 12, 16, 14, 16, 386, 11, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 394, 10, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 416, 10,
	18, 3, 19, 6, 19, 419, 10, 19, 13, 19, 14, 19, 420, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22,
	446, 10, 22, 12, 22, 14, 22, 449, 11, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24,
	465, 10, 24, 12, 24, 14, 24, 468, 11, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 479, 10, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 512, 10, 30, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 7, 31, 523,
	10, 31, 12, 31, 14, 31, 526, 11, 31, 3, 32, 3, 32, 5, 32, 530, 10, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 535, 10, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 556, 10, 33, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 567, 10, 34, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 577, 10, 35, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 589,
	10, 37, 3, 37, 3, 37, 5, 37, 593, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 5, 38, 613, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 7, 41, 632, 10, 41, 12, 41, 14, 41, 635, 11, 41, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 7, 43, 651, 10, 43, 12, 43, 14, 43, 654, 11, 43, 3, 44, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 677,
	10, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 7, 44, 689, 10, 44, 12, 44, 14, 44, 692, 11, 44, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45,
	706, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 5,
	46, 716, 10, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47,
	3, 47, 5, 47, 727, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 5, 49, 748, 10, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 3, 50, 3, 50, 3, 50, 5, 50, 759, 10, 50, 3, 50, 3, 50, 5, 50, 763,
	10, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 5, 50, 772, 10,
	50, 3, 50, 2, 12, 4, 14, 22, 30, 42, 46, 60, 80, 84, 86, 51, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
	82, 84, 86, 88, 90, 92, 94, 96, 98, 2, 7, 3, 2, 40, 41, 3, 2, 54, 59, 3,
	2, 60, 62, 3, 2, 63, 64, 3, 2, 65, 66, 2, 809, 2, 100, 3, 2, 2, 2, 4, 103,
	3, 2, 2, 2, 6, 153, 3, 2, 2, 2, 8, 155, 3, 2, 2, 2, 10, 170, 3, 2, 2, 2,
	12, 180, 3, 2, 2, 2, 14, 182, 3, 2, 2, 2, 16, 203, 3, 2, 2, 2, 18, 235,
	3, 2, 2, 2, 20, 273, 3, 2, 2, 2, 22, 275, 3, 2, 2, 2, 24, 337, 3, 2, 2,
	2, 26, 356, 3, 2, 2, 2, 28, 372, 3, 2, 2, 2, 30, 374, 3, 2, 2, 2, 32, 393,
	3, 2, 2, 2, 34, 415, 3, 2, 2, 2, 36, 418, 3, 2, 2, 2, 38, 424, 3, 2, 2,
	2, 40, 430, 3, 2, 2, 2, 42, 437, 3, 2, 2, 2, 44, 450, 3, 2, 2, 2, 46, 455,
	3, 2, 2, 2, 48, 478, 3, 2, 2, 2, 50, 480, 3, 2, 2, 2, 52, 485, 3, 2, 2,
	2, 54, 489, 3, 2, 2, 2, 56, 496, 3, 2, 2, 2, 58, 511, 3, 2, 2, 2, 60, 513,
	3, 2, 2, 2, 62, 527, 3, 2, 2, 2, 64, 555, 3, 2, 2, 2, 66, 566, 3, 2, 2,
	2, 68, 576, 3, 2, 2, 2, 70, 578, 3, 2, 2, 2, 72, 592, 3, 2, 2, 2, 74, 612,
	3, 2, 2, 2, 76, 614, 3, 2, 2, 2, 78, 619, 3, 2, 2, 2, 80, 623, 3, 2, 2,
	2, 82, 636, 3, 2, 2, 2, 84, 641, 3, 2, 2, 2, 86, 676, 3, 2, 2, 2, 88, 705,
	3, 2, 2, 2, 90, 715, 3, 2, 2, 2, 92, 726, 3, 2, 2, 2, 94, 728, 3, 2, 2,
	2, 96, 747, 3, 2, 2, 2, 98, 771, 3, 2, 2, 2, 100, 101, 5, 4, 3, 2, 101,
	102, 8, 2, 1, 2, 102, 3, 3, 2, 2, 2, 103, 104, 8, 3, 1, 2, 104, 105, 5,
	6, 4, 2, 105, 106, 8, 3, 1, 2, 106, 113, 3, 2, 2, 2, 107, 108, 12, 4, 2,
	2, 108, 109, 5, 6, 4, 2, 109, 110, 8, 3, 1, 2, 110, 112, 3, 2, 2, 2, 111,
	107, 3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114,
	3, 2, 2, 2, 114, 5, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 117, 5, 8, 5,
	2, 117, 118, 8, 4, 1, 2, 118, 154, 3, 2, 2, 2, 119, 120, 7, 29, 2, 2, 120,
	121, 7, 37, 2, 2, 121, 122, 7, 48, 2, 2, 122, 123, 7, 49, 2, 2, 123, 124,
	5, 10, 6, 2, 124, 125, 8, 4, 1, 2, 125, 154, 3, 2, 2, 2, 126, 127, 7, 29,
	2, 2, 127, 128, 7, 37, 2, 2, 128, 129, 7, 48, 2, 2, 129, 130, 7, 49, 2,
	2, 130, 131, 7, 46, 2, 2, 131, 132, 5, 96, 49, 2, 132, 133, 5, 10, 6, 2,
	133, 134, 8, 4, 1, 2, 134, 154, 3, 2, 2, 2, 135, 136, 7, 29, 2, 2, 136,
	137, 7, 37, 2, 2, 137, 138, 7, 48, 2, 2, 138, 139, 5, 14, 8, 2, 139, 140,
	7, 49, 2, 2, 140, 141, 5, 10, 6, 2, 141, 142, 8, 4, 1, 2, 142, 154, 3,
	2, 2, 2, 143, 144, 7, 29, 2, 2, 144, 145, 7, 37, 2, 2, 145, 146, 7, 48,
	2, 2, 146, 147, 5, 14, 8, 2, 147, 148, 7, 49, 2, 2, 148, 149, 7, 46, 2,
	2, 149, 150, 5, 96, 49, 2, 150, 151, 5, 10, 6, 2, 151, 152, 8, 4, 1, 2,
	152, 154, 3, 2, 2, 2, 153, 116, 3, 2, 2, 2, 153, 119, 3, 2, 2, 2, 153,
	126, 3, 2, 2, 2, 153, 135, 3, 2, 2, 2, 153, 143, 3, 2, 2, 2, 154, 7, 3,
	2, 2, 2, 155, 156, 7, 29, 2, 2, 156, 157, 7, 28, 2, 2, 157, 158, 7, 48,
	2, 2, 158, 159, 7, 49, 2, 2, 159, 160, 5, 10, 6, 2, 160, 161, 8, 5, 1,
	2, 161, 9, 3, 2, 2, 2, 162, 163, 7, 50, 2, 2, 163, 164, 5, 16, 9, 2, 164,
	165, 7, 51, 2, 2, 165, 166, 8, 6, 1, 2, 166, 171, 3, 2, 2, 2, 167, 168,
	7, 50, 2, 2, 168, 169, 7, 51, 2, 2, 169, 171, 8, 6, 1, 2, 170, 162, 3,
	2, 2, 2, 170, 167, 3, 2, 2, 2, 171, 11, 3, 2, 2, 2, 172, 173, 7, 50, 2,
	2, 173, 174, 5, 16, 9, 2, 174, 175, 7, 51, 2, 2, 175, 176, 8, 7, 1, 2,
	176, 181, 3, 2, 2, 2, 177, 178, 5, 16, 9, 2, 178, 179, 8, 7, 1, 2, 179,
	181, 3, 2, 2, 2, 180, 172, 3, 2, 2, 2, 180, 177, 3, 2, 2, 2, 181, 13, 3,
	2, 2, 2, 182, 183, 8, 8, 1, 2, 183, 184, 7, 37, 2, 2, 184, 185, 7, 42,
	2, 2, 185, 186, 5, 96, 49, 2, 186, 187, 8, 8, 1, 2, 187, 197, 3, 2, 2,
	2, 188, 189, 12, 4, 2, 2, 189, 190, 7, 41, 2, 2, 190, 191, 7, 37, 2, 2,
	191, 192, 7, 42, 2, 2, 192, 193, 5, 96, 49, 2, 193, 194, 8, 8, 1, 2, 194,
	196, 3, 2, 2, 2, 195, 188, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195,
	3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 15, 3, 2, 2, 2, 199, 197, 3, 2,
	2, 2, 200, 202, 5, 18, 10, 2, 201, 200, 3, 2, 2, 2, 202, 205, 3, 2, 2,
	2, 203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 206, 3, 2, 2, 2, 205,
	203, 3, 2, 2, 2, 206, 207, 8, 9, 1, 2, 207, 17, 3, 2, 2, 2, 208, 209, 5,
	20, 11, 2, 209, 210, 8, 10, 1, 2, 210, 236, 3, 2, 2, 2, 211, 212, 5, 24,
	13, 2, 212, 213, 8, 10, 1, 2, 213, 236, 3, 2, 2, 2, 214, 215, 5, 26, 14,
	2, 215, 216, 8, 10, 1, 2, 216, 236, 3, 2, 2, 2, 217, 218, 5, 32, 17, 2,
	218, 219, 8, 10, 1, 2, 219, 236, 3, 2, 2, 2, 220, 221, 5, 48, 25, 2, 221,
	222, 8, 10, 1, 2, 222, 236, 3, 2, 2, 2, 223, 225, 5, 84, 43, 2, 224, 226,
	7, 41, 2, 2, 225, 224, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 227, 3, 2,
	2, 2, 227, 228, 8, 10, 1, 2, 228, 236, 3, 2, 2, 2, 229, 230, 5, 62, 32,
	2, 230, 231, 8, 10, 1, 2, 231, 236, 3, 2, 2, 2, 232, 233, 5, 66, 34, 2,
	233, 234, 8, 10, 1, 2, 234, 236, 3, 2, 2, 2, 235, 208, 3, 2, 2, 2, 235,
	211, 3, 2, 2, 2, 235, 214, 3, 2, 2, 2, 235, 217, 3, 2, 2, 2, 235, 220,
	3, 2, 2, 2, 235, 223, 3, 2, 2, 2, 235, 229, 3, 2, 2, 2, 235, 232, 3, 2,
	2, 2, 236, 19, 3, 2, 2, 2, 237, 238, 7, 4, 2, 2, 238, 239, 7, 43, 2, 2,
	239, 240, 7, 48, 2, 2, 240, 241, 5, 74, 38, 2, 241, 242, 7, 49, 2, 2, 242,
	243, 9, 2, 2, 2, 243, 244, 8, 11, 1, 2, 244, 274, 3, 2, 2, 2, 245, 246,
	7, 4, 2, 2, 246, 247, 7, 43, 2, 2, 247, 248, 7, 48, 2, 2, 248, 249, 5,
	74, 38, 2, 249, 250, 7, 41, 2, 2, 250, 251, 5, 22, 12, 2, 251, 252, 7,
	49, 2, 2, 252, 253, 9, 2, 2, 2, 253, 254, 8, 11, 1, 2, 254, 274, 3, 2,
	2, 2, 255, 256, 7, 3, 2, 2, 256, 257, 7, 43, 2, 2, 257, 258, 7, 48, 2,
	2, 258, 259, 5, 74, 38, 2, 259, 260, 7, 49, 2, 2, 260, 261, 9, 2, 2, 2,
	261, 262, 8, 11, 1, 2, 262, 274, 3, 2, 2, 2, 263, 264, 7, 3, 2, 2, 264,
	265, 7, 43, 2, 2, 265, 266, 7, 48, 2, 2, 266, 267, 5, 74, 38, 2, 267, 268,
	7, 41, 2, 2, 268, 269, 5, 22, 12, 2, 269, 270, 7, 49, 2, 2, 270, 271, 9,
	2, 2, 2, 271, 272, 8, 11, 1, 2, 272, 274, 3, 2, 2, 2, 273, 237, 3, 2, 2,
	2, 273, 245, 3, 2, 2, 2, 273, 255, 3, 2, 2, 2, 273, 263, 3, 2, 2, 2, 274,
	21, 3, 2, 2, 2, 275, 276, 8, 12, 1, 2, 276, 277, 5, 74, 38, 2, 277, 278,
	8, 12, 1, 2, 278, 286, 3, 2, 2, 2, 279, 280, 12, 4, 2, 2, 280, 281, 7,
	41, 2, 2, 281, 282, 5, 74, 38, 2, 282, 283, 8, 12, 1, 2, 283, 285, 3, 2,
	2, 2, 284, 279, 3, 2, 2, 2, 285, 288, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2,
	286, 287, 3, 2, 2, 2, 287, 23, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 289, 291,
	7, 5, 2, 2, 290, 292, 7, 6, 2, 2, 291, 290, 3, 2, 2, 2, 291, 292, 3, 2,
	2, 2, 292, 293, 3, 2, 2, 2, 293, 294, 5, 28, 15, 2, 294, 295, 7, 38, 2,
	2, 295, 297, 5, 74, 38, 2, 296, 298, 9, 2, 2, 2, 297, 296, 3, 2, 2, 2,
	297, 298, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 300, 8, 13, 1, 2, 300,
	338, 3, 2, 2, 2, 301, 303, 7, 5, 2, 2, 302, 304, 7, 6, 2, 2, 303, 302,
	3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 306, 5, 28,
	15, 2, 306, 307, 7, 42, 2, 2, 307, 308, 5, 96, 49, 2, 308, 309, 7, 38,
	2, 2, 309, 311, 5, 74, 38, 2, 310, 312, 9, 2, 2, 2, 311, 310, 3, 2, 2,
	2, 311, 312, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 314, 8, 13, 1, 2, 314,
	338, 3, 2, 2, 2, 315, 317, 7, 5, 2, 2, 316, 318, 7, 6, 2, 2, 317, 316,
	3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 321, 5, 28,
	15, 2, 320, 322, 9, 2, 2, 2, 321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2,
	322, 323, 3, 2, 2, 2, 323, 324, 8, 13, 1, 2, 324, 338, 3, 2, 2, 2, 325,
	327, 7, 5, 2, 2, 326, 328, 7, 6, 2, 2, 327, 326, 3, 2, 2, 2, 327, 328,
	3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 330, 5, 28, 15, 2, 330, 331, 7,
	42, 2, 2, 331, 333, 5, 96, 49, 2, 332, 334, 9, 2, 2, 2, 333, 332, 3, 2,
	2, 2, 333, 334, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 336, 8, 13, 1, 2,
	336, 338, 3, 2, 2, 2, 337, 289, 3, 2, 2, 2, 337, 301, 3, 2, 2, 2, 337,
	315, 3, 2, 2, 2, 337, 325, 3, 2, 2, 2, 338, 25, 3, 2, 2, 2, 339, 340, 5,
	30, 16, 2, 340, 341, 7, 38, 2, 2, 341, 343, 5, 74, 38, 2, 342, 344, 9,
	2, 2, 2, 343, 342, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 345, 3, 2, 2,
	2, 345, 346, 8, 14, 1, 2, 346, 357, 3, 2, 2, 2, 347, 348, 7, 37, 2, 2,
	348, 349, 5, 80, 41, 2, 349, 350, 7, 38, 2, 2, 350, 352, 5, 74, 38, 2,
	351, 353, 9, 2, 2, 2, 352, 351, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353,
	354, 3, 2, 2, 2, 354, 355, 8, 14, 1, 2, 355, 357, 3, 2, 2, 2, 356, 339,
	3, 2, 2, 2, 356, 347, 3, 2, 2, 2, 357, 27, 3, 2, 2, 2, 358, 359, 5, 30,
	16, 2, 359, 360, 8, 15, 1, 2, 360, 373, 3, 2, 2, 2, 361, 362, 5, 30, 16,
	2, 362, 363, 7, 42, 2, 2, 363, 364, 7, 7, 2, 2, 364, 365, 8, 15, 1, 2,
	365, 373, 3, 2, 2, 2, 366, 367, 5, 30, 16, 2, 367, 368, 7, 42, 2, 2, 368,
	369, 7, 44, 2, 2, 369, 370, 7, 12, 2, 2, 370, 371, 8, 15, 1, 2, 371, 373,
	3, 2, 2, 2, 372, 358, 3, 2, 2, 2, 372, 361, 3, 2, 2, 2, 372, 366, 3, 2,
	2, 2, 373, 29, 3, 2, 2, 2, 374, 375, 8, 16, 1, 2, 375, 376, 7, 37, 2, 2,
	376, 377, 8, 16, 1, 2, 377, 384, 3, 2, 2, 2, 378, 379, 12, 4, 2, 2, 379,
	380, 7, 41, 2, 2, 380, 381, 7, 37, 2, 2, 381, 383, 8, 16, 1, 2, 382, 378,
	3, 2, 2, 2, 383, 386, 3, 2, 2, 2, 384, 382, 3, 2, 2, 2, 384, 385, 3, 2,
	2, 2, 385, 31, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 387, 388, 5, 34, 18, 2,
	388, 389, 8, 17, 1, 2, 389, 394, 3, 2, 2, 2, 390, 391, 5, 40, 21, 2, 391,
	392, 8, 17, 1, 2, 392, 394, 3, 2, 2, 2, 393, 387, 3, 2, 2, 2, 393, 390,
	3, 2, 2, 2, 394, 33, 3, 2, 2, 2, 395, 396, 7, 14, 2, 2, 396, 397, 5, 74,
	38, 2, 397, 398, 5, 10, 6, 2, 398, 399, 8, 18, 1, 2, 399, 416, 3, 2, 2,
	2, 400, 401, 7, 14, 2, 2, 401, 402, 5, 74, 38, 2, 402, 403, 5, 10, 6, 2,
	403, 404, 7, 15, 2, 2, 404, 405, 5, 10, 6, 2, 405, 406, 8, 18, 1, 2, 406,
	416, 3, 2, 2, 2, 407, 408, 7, 14, 2, 2, 408, 409, 5, 74, 38, 2, 409, 410,
	5, 10, 6, 2, 410, 411, 5, 36, 19, 2, 411, 412, 7, 15, 2, 2, 412, 413, 5,
	10, 6, 2, 413, 414, 8, 18, 1, 2, 414, 416, 3, 2, 2, 2, 415, 395, 3, 2,
	2, 2, 415, 400, 3, 2, 2, 2, 415, 407, 3, 2, 2, 2, 416, 35, 3, 2, 2, 2,
	417, 419, 5, 38, 20, 2, 418, 417, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420,
	418, 3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 423,
	8, 19, 1, 2, 423, 37, 3, 2, 2, 2, 424, 425, 7, 15, 2, 2, 425, 426, 7, 14,
	2, 2, 426, 427, 5, 74, 38, 2, 427, 428, 5, 10, 6, 2, 428, 429, 8, 20, 1,
	2, 429, 39, 3, 2, 2, 2, 430, 431, 7, 16, 2, 2, 431, 432, 5, 74, 38, 2,
	432, 433, 7, 50, 2, 2, 433, 434, 5, 42, 22, 2, 434, 435, 7, 51, 2, 2, 435,
	436, 8, 21, 1, 2, 436, 41, 3, 2, 2, 2, 437, 438, 8, 22, 1, 2, 438, 439,
	5, 44, 23, 2, 439, 440, 8, 22, 1, 2, 440, 447, 3, 2, 2, 2, 441, 442, 12,
	4, 2, 2, 442, 443, 5, 44, 23, 2, 443, 444, 8, 22, 1, 2, 444, 446, 3, 2,
	2, 2, 445, 441, 3, 2, 2, 2, 446, 449, 3, 2, 2, 2, 447, 445, 3, 2, 2, 2,
	447, 448, 3, 2, 2, 2, 448, 43, 3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 450, 451,
	5, 46, 24, 2, 451, 452, 7, 68, 2, 2, 452, 453, 5, 12, 7, 2, 453, 454, 8,
	23, 1, 2, 454, 45, 3, 2, 2, 2, 455, 456, 8, 24, 1, 2, 456, 457, 5, 74,
	38, 2, 457, 458, 8, 24, 1, 2, 458, 466, 3, 2, 2, 2, 459, 460, 12, 4, 2,
	2, 460, 461, 7, 67, 2, 2, 461, 462, 5, 74, 38, 2, 462, 463, 8, 24, 1, 2,
	463, 465, 3, 2, 2, 2, 464, 459, 3, 2, 2, 2, 465, 468, 3, 2, 2, 2, 466,
	464, 3, 2, 2, 2, 466, 467, 3, 2, 2, 2, 467, 47, 3, 2, 2, 2, 468, 466, 3,
	2, 2, 2, 469, 470, 5, 50, 26, 2, 470, 471, 8, 25, 1, 2, 471, 479, 3, 2,
	2, 2, 472, 473, 5, 52, 27, 2, 473, 474, 8, 25, 1, 2, 474, 479, 3, 2, 2,
	2, 475, 476, 5, 54, 28, 2, 476, 477, 8, 25, 1, 2, 477, 479, 3, 2, 2, 2,
	478, 469, 3, 2, 2, 2, 478, 472, 3, 2, 2, 2, 478, 475, 3, 2, 2, 2, 479,
	49, 3, 2, 2, 2, 480, 481, 7, 17, 2, 2, 481, 482, 5, 74, 38, 2, 482, 483,
	5, 10, 6, 2, 483, 484, 8, 26, 1, 2, 484, 51, 3, 2, 2, 2, 485, 486, 7, 18,
	2, 2, 486, 487, 5, 10, 6, 2, 487, 488, 8, 27, 1, 2, 488, 53, 3, 2, 2, 2,
	489, 490, 7, 19, 2, 2, 490, 491, 5, 74, 38, 2, 491, 492, 7, 20, 2, 2, 492,
	493, 5, 56, 29, 2, 493, 494, 5, 10, 6, 2, 494, 495, 8, 28, 1, 2, 495, 55,
	3, 2, 2, 2, 496, 497, 5, 74, 38, 2, 497, 498, 7, 47, 2, 2, 498, 499, 5,
	74, 38, 2, 499, 500, 8, 29, 1, 2, 500, 57, 3, 2, 2, 2, 501, 502, 7, 37,
	2, 2, 502, 503, 7, 48, 2, 2, 503, 504, 7, 49, 2, 2, 504, 512, 8, 30, 1,
	2, 505, 506, 7, 37, 2, 2, 506, 507, 7, 48, 2, 2, 507, 508, 5, 60, 31, 2,
	508, 509, 7, 49, 2, 2, 509, 510, 8, 30, 1, 2, 510, 512, 3, 2, 2, 2, 511,
	501, 3, 2, 2, 2, 511, 505, 3, 2, 2, 2, 512, 59, 3, 2, 2, 2, 513, 514, 8,
	31, 1, 2, 514, 515, 5, 74, 38, 2, 515, 516, 8, 31, 1, 2, 516, 524, 3, 2,
	2, 2, 517, 518, 12, 4, 2, 2, 518, 519, 9, 2, 2, 2, 519, 520, 5, 74, 38,
	2, 520, 521, 8, 31, 1, 2, 521, 523, 3, 2, 2, 2, 522, 517, 3, 2, 2, 2, 523,
	526, 3, 2, 2, 2, 524, 522, 3, 2, 2, 2, 524, 525, 3, 2, 2, 2, 525, 61, 3,
	2, 2, 2, 526, 524, 3, 2, 2, 2, 527, 529, 7, 5, 2, 2, 528, 530, 7, 6, 2,
	2, 529, 528, 3, 2, 2, 2, 529, 530, 3, 2, 2, 2, 530, 531, 3, 2, 2, 2, 531,
	534, 7, 37, 2, 2, 532, 533, 7, 42, 2, 2, 533, 535, 5, 64, 33, 2, 534, 532,
	3, 2, 2, 2, 534, 535, 3, 2, 2, 2, 535, 536, 3, 2, 2, 2, 536, 537, 7, 38,
	2, 2, 537, 538, 5, 74, 38, 2, 538, 539, 9, 2, 2, 2, 539, 540, 8, 32, 1,
	2, 540, 63, 3, 2, 2, 2, 541, 542, 7, 52, 2, 2, 542, 543, 5, 64, 33, 2,
	543, 544, 7, 40, 2, 2, 544, 545, 5, 74, 38, 2, 545, 546, 7, 53, 2, 2, 546,
	547, 8, 33, 1, 2, 547, 556, 3, 2, 2, 2, 548, 549, 7, 52, 2, 2, 549, 550,
	5, 96, 49, 2, 550, 551, 7, 40, 2, 2, 551, 552, 5, 74, 38, 2, 552, 553,
	7, 53, 2, 2, 553, 554, 8, 33, 1, 2, 554, 556, 3, 2, 2, 2, 555, 541, 3,
	2, 2, 2, 555, 548, 3, 2, 2, 2, 556, 65, 3, 2, 2, 2, 557, 558, 5, 68, 35,
	2, 558, 559, 8, 34, 1, 2, 559, 567, 3, 2, 2, 2, 560, 561, 5, 70, 36, 2,
	561, 562, 8, 34, 1, 2, 562, 567, 3, 2, 2, 2, 563, 564, 5, 72, 37, 2, 564,
	565, 8, 34, 1, 2, 565, 567, 3, 2, 2, 2, 566, 557, 3, 2, 2, 2, 566, 560,
	3, 2, 2, 2, 566, 563, 3, 2, 2, 2, 567, 67, 3, 2, 2, 2, 568, 569, 7, 21,
	2, 2, 569, 570, 7, 40, 2, 2, 570, 577, 8, 35, 1, 2, 571, 572, 7, 21, 2,
	2, 572, 573, 5, 74, 38, 2, 573, 574, 7, 40, 2, 2, 574, 575, 8, 35, 1, 2,
	575, 577, 3, 2, 2, 2, 576, 568, 3, 2, 2, 2, 576, 571, 3, 2, 2, 2, 577,
	69, 3, 2, 2, 2, 578, 579, 7, 22, 2, 2, 579, 580, 7, 40, 2, 2, 580, 581,
	8, 36, 1, 2, 581, 71, 3, 2, 2, 2, 582, 583, 7, 23, 2, 2, 583, 584, 7, 40,
	2, 2, 584, 593, 8, 37, 1, 2, 585, 586, 7, 23, 2, 2, 586, 588, 5, 74, 38,
	2, 587, 589, 7, 40, 2, 2, 588, 587, 3, 2, 2, 2, 588, 589, 3, 2, 2, 2, 589,
	590, 3, 2, 2, 2, 590, 591, 8, 37, 1, 2, 591, 593, 3, 2, 2, 2, 592, 582,
	3, 2, 2, 2, 592, 585, 3, 2, 2, 2, 593, 73, 3, 2, 2, 2, 594, 595, 5, 88,
	45, 2, 595, 596, 8, 38, 1, 2, 596, 613, 3, 2, 2, 2, 597, 598, 5, 32, 17,
	2, 598, 599, 8, 38, 1, 2, 599, 613, 3, 2, 2, 2, 600, 601, 5, 52, 27, 2,
	601, 602, 8, 38, 1, 2, 602, 613, 3, 2, 2, 2, 603, 604, 5, 84, 43, 2, 604,
	605, 8, 38, 1, 2, 605, 613, 3, 2, 2, 2, 606, 607, 5, 92, 47, 2, 607, 608,
	8, 38, 1, 2, 608, 613, 3, 2, 2, 2, 609, 610, 5, 76, 39, 2, 610, 611, 8,
	38, 1, 2, 611, 613, 3, 2, 2, 2, 612, 594, 3, 2, 2, 2, 612, 597, 3, 2, 2,
	2, 612, 600, 3, 2, 2, 2, 612, 603, 3, 2, 2, 2, 612, 606, 3, 2, 2, 2, 612,
	609, 3, 2, 2, 2, 613, 75, 3, 2, 2, 2, 614, 615, 7, 52, 2, 2, 615, 616,
	5, 60, 31, 2, 616, 617, 7, 53, 2, 2, 617, 618, 8, 39, 1, 2, 618, 77, 3,
	2, 2, 2, 619, 620, 7, 37, 2, 2, 620, 621, 5, 80, 41, 2, 621, 622, 8, 40,
	1, 2, 622, 79, 3, 2, 2, 2, 623, 624, 8, 41, 1, 2, 624, 625, 5, 82, 42,
	2, 625, 626, 8, 41, 1, 2, 626, 633, 3, 2, 2, 2, 627, 628, 12, 4, 2, 2,
	628, 629, 5, 82, 42, 2, 629, 630, 8, 41, 1, 2, 630, 632, 3, 2, 2, 2, 631,
	627, 3, 2, 2, 2, 632, 635, 3, 2, 2, 2, 633, 631, 3, 2, 2, 2, 633, 634,
	3, 2, 2, 2, 634, 81, 3, 2, 2, 2, 635, 633, 3, 2, 2, 2, 636, 637, 7, 52,
	2, 2, 637, 638, 5, 74, 38, 2, 638, 639, 7, 53, 2, 2, 639, 640, 8, 42, 1,
	2, 640, 83, 3, 2, 2, 2, 641, 642, 8, 43, 1, 2, 642, 643, 5, 86, 44, 2,
	643, 644, 8, 43, 1, 2, 644, 652, 3, 2, 2, 2, 645, 646, 12, 4, 2, 2, 646,
	647, 9, 3, 2, 2, 647, 648, 5, 84, 43, 5, 648, 649, 8, 43, 1, 2, 649, 651,
	3, 2, 2, 2, 650, 645, 3, 2, 2, 2, 651, 654, 3, 2, 2, 2, 652, 650, 3, 2,
	2, 2, 652, 653, 3, 2, 2, 2, 653, 85, 3, 2, 2, 2, 654, 652, 3, 2, 2, 2,
	655, 656, 8, 44, 1, 2, 656, 657, 7, 64, 2, 2, 657, 658, 5, 74, 38, 2, 658,
	659, 8, 44, 1, 2, 659, 677, 3, 2, 2, 2, 660, 661, 5, 90, 46, 2, 661, 662,
	7, 48, 2, 2, 662, 663, 5, 86, 44, 2, 663, 664, 7, 41, 2, 2, 664, 665, 5,
	86, 44, 2, 665, 666, 7, 49, 2, 2, 666, 667, 8, 44, 1, 2, 667, 677, 3, 2,
	2, 2, 668, 669, 5, 88, 45, 2, 669, 670, 8, 44, 1, 2, 670, 677, 3, 2, 2,
	2, 671, 672, 7, 48, 2, 2, 672, 673, 5, 74, 38, 2, 673, 674, 7, 49, 2, 2,
	674, 675, 8, 44, 1, 2, 675, 677, 3, 2, 2, 2, 676, 655, 3, 2, 2, 2, 676,
	660, 3, 2, 2, 2, 676, 668, 3, 2, 2, 2, 676, 671, 3, 2, 2, 2, 677, 690,
	3, 2, 2, 2, 678, 679, 12, 6, 2, 2, 679, 680, 9, 4, 2, 2, 680, 681, 5, 86,
	44, 7, 681, 682, 8, 44, 1, 2, 682, 689, 3, 2, 2, 2, 683, 684, 12, 5, 2,
	2, 684, 685, 9, 5, 2, 2, 685, 686, 5, 86, 44, 6, 686, 687, 8, 44, 1, 2,
	687, 689, 3, 2, 2, 2, 688, 678, 3, 2, 2, 2, 688, 683, 3, 2, 2, 2, 689,
	692, 3, 2, 2, 2, 690, 688, 3, 2, 2, 2, 690, 691, 3, 2, 2, 2, 691, 87, 3,
	2, 2, 2, 692, 690, 3, 2, 2, 2, 693, 694, 5, 58, 30, 2, 694, 695, 8, 45,
	1, 2, 695, 706, 3, 2, 2, 2, 696, 697, 5, 98, 50, 2, 697, 698, 8, 45, 1,
	2, 698, 706, 3, 2, 2, 2, 699, 700, 5, 94, 48, 2, 700, 701, 8, 45, 1, 2,
	701, 706, 3, 2, 2, 2, 702, 703, 5, 78, 40, 2, 703, 704, 8, 45, 1, 2, 704,
	706, 3, 2, 2, 2, 705, 693, 3, 2, 2, 2, 705, 696, 3, 2, 2, 2, 705, 699,
	3, 2, 2, 2, 705, 702, 3, 2, 2, 2, 706, 89, 3, 2, 2, 2, 707, 708, 7, 8,
	2, 2, 708, 709, 7, 45, 2, 2, 709, 710, 7, 24, 2, 2, 710, 716, 8, 46, 1,
	2, 711, 712, 7, 9, 2, 2, 712, 713, 7, 45, 2, 2, 713, 714, 7, 25, 2, 2,
	714, 716, 8, 46, 1, 2, 715, 707, 3, 2, 2, 2, 715, 711, 3, 2, 2, 2, 716,
	91, 3, 2, 2, 2, 717, 718, 7, 43, 2, 2, 718, 719, 5, 74, 38, 2, 719, 720,
	8, 47, 1, 2, 720, 727, 3, 2, 2, 2, 721, 722, 5, 84, 43, 2, 722, 723, 9,
	6, 2, 2, 723, 724, 5, 84, 43, 2, 724, 725, 8, 47, 1, 2, 725, 727, 3, 2,
	2, 2, 726, 717, 3, 2, 2, 2, 726, 721, 3, 2, 2, 2, 727, 93, 3, 2, 2, 2,
	728, 729, 7, 48, 2, 2, 729, 730, 5, 74, 38, 2, 730, 731, 7, 13, 2, 2, 731,
	732, 5, 96, 49, 2, 732, 733, 7, 49, 2, 2, 733, 734, 8, 48, 1, 2, 734, 95,
	3, 2, 2, 2, 735, 736, 7, 8, 2, 2, 736, 748, 8, 49, 1, 2, 737, 738, 7, 9,
	2, 2, 738, 748, 8, 49, 1, 2, 739, 740, 7, 12, 2, 2, 740, 748, 8, 49, 1,
	2, 741, 742, 7, 7, 2, 2, 742, 748, 8, 49, 1, 2, 743, 744, 7, 10, 2, 2,
	744, 748, 8, 49, 1, 2, 745, 746, 7, 11, 2, 2, 746, 748, 8, 49, 1, 2, 747,
	735, 3, 2, 2, 2, 747, 737, 3, 2, 2, 2, 747, 739, 3, 2, 2, 2, 747, 741,
	3, 2, 2, 2, 747, 743, 3, 2, 2, 2, 747, 745, 3, 2, 2, 2, 748, 97, 3, 2,
	2, 2, 749, 750, 7, 32, 2, 2, 750, 772, 8, 50, 1, 2, 751, 752, 7, 33, 2,
	2, 752, 772, 8, 50, 1, 2, 753, 762, 7, 35, 2, 2, 754, 755, 7, 39, 2, 2,
	755, 759, 7, 26, 2, 2, 756, 757, 7, 39, 2, 2, 757, 759, 7, 27, 2, 2, 758,
	754, 3, 2, 2, 2, 758, 756, 3, 2, 2, 2, 759, 760, 3, 2, 2, 2, 760, 761,
	7, 48, 2, 2, 761, 763, 7, 49, 2, 2, 762, 758, 3, 2, 2, 2, 762, 763, 3,
	2, 2, 2, 763, 764, 3, 2, 2, 2, 764, 772, 8, 50, 1, 2, 765, 766, 7, 34,
	2, 2, 766, 772, 8, 50, 1, 2, 767, 768, 7, 36, 2, 2, 768, 772, 8, 50, 1,
	2, 769, 770, 7, 37, 2, 2, 770, 772, 8, 50, 1, 2, 771, 749, 3, 2, 2, 2,
	771, 751, 3, 2, 2, 2, 771, 753, 3, 2, 2, 2, 771, 765, 3, 2, 2, 2, 771,
	767, 3, 2, 2, 2, 771, 769, 3, 2, 2, 2, 772, 99, 3, 2, 2, 2, 54, 113, 153,
	170, 180, 197, 203, 225, 235, 273, 286, 291, 297, 303, 311, 317, 321, 327,
	333, 337, 343, 352, 356, 372, 384, 393, 415, 420, 447, 466, 478, 511, 524,
	529, 534, 555, 566, 576, 588, 592, 612, 633, 652, 676, 688, 690, 705, 715,
	726, 747, 758, 762, 771,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", "'f64'",
	"'bool'", "'char'", "'&str'", "'as'", "'if'", "'else'", "'match'", "'while'",
	"'loop'", "'for'", "'in'", "'break'", "'continue'", "'return'", "'pow'",
	"'powf'", "'to_string'", "'to_owned'", "'main'", "'fn'", "", "", "", "",
	"", "", "", "", "'='", "'.'", "';'", "','", "':'", "'!'", "'&'", "'::'",
	"'->'", "'..'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='",
	"'>'", "'<'", "'>='", "'<='", "'*'", "'/'", "'%'", "'+'", "'-'", "'&&'",
	"'||'", "'|'", "'=>'", "'_'",
}
var symbolicNames = []string{
	"", "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL",
	"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE",
	"RLOOP", "RFOR", "RIN", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF",
	"TOSTRING", "TOOWNED", "RMAIN", "RFN", "MULTILINE", "INLINE", "INTEGER",
	"FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON",
	"COMMA", "COLON", "ADMIRATION", "REFERENCE", "HERITAGE", "ARROW", "RANGE",
	"LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET",
	"EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN",
	"LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB", "AND", "OR", "PIPE",
	"EQUAL_ARROW", "UNDERSCORE", "WHITESPACE",
}

var ruleNames = []string{
	"start", "listFuncs", "function", "funcMain", "bloq", "bloq_match", "listParams",
	"instructions", "instruction", "print_prod", "listVars", "declaration_prod",
	"assign_prod", "ids_type", "listIds", "conditional_prod", "if_prod", "list_else_if",
	"else_if", "match_prod", "list_instr_match", "instr_match", "expr_match",
	"bucle_prod", "while_prod", "loop_prod", "forin_prod", "range_prod", "called_func",
	"listExpressions", "dec_arr", "listDim", "transfer_prod", "break_instr",
	"continue_instr", "return_instr", "expression", "arraydata", "access_array",
	"listInArray", "inArray", "expr_rel", "expr_arit", "expr_valor", "pow_op",
	"expr_logic", "expr_cast", "data_type", "primitive",
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
	ProjectParserMUT               = 4
	ProjectParserRSTRING           = 5
	ProjectParserRINTEGER          = 6
	ProjectParserRREAL             = 7
	ProjectParserRBOOLEAN          = 8
	ProjectParserRCHAR             = 9
	ProjectParserRSTR              = 10
	ProjectParserRAS               = 11
	ProjectParserRIF               = 12
	ProjectParserRELSE             = 13
	ProjectParserRMATCH            = 14
	ProjectParserRWHILE            = 15
	ProjectParserRLOOP             = 16
	ProjectParserRFOR              = 17
	ProjectParserRIN               = 18
	ProjectParserRBREAK            = 19
	ProjectParserRCONTINUE         = 20
	ProjectParserRRETURN           = 21
	ProjectParserPOWI              = 22
	ProjectParserPOWF              = 23
	ProjectParserTOSTRING          = 24
	ProjectParserTOOWNED           = 25
	ProjectParserRMAIN             = 26
	ProjectParserRFN               = 27
	ProjectParserMULTILINE         = 28
	ProjectParserINLINE            = 29
	ProjectParserINTEGER           = 30
	ProjectParserFLOAT             = 31
	ProjectParserCHAR              = 32
	ProjectParserSTRING            = 33
	ProjectParserBOOLEAN           = 34
	ProjectParserID                = 35
	ProjectParserEQUAL             = 36
	ProjectParserDOT               = 37
	ProjectParserSEMICOLON         = 38
	ProjectParserCOMMA             = 39
	ProjectParserCOLON             = 40
	ProjectParserADMIRATION        = 41
	ProjectParserREFERENCE         = 42
	ProjectParserHERITAGE          = 43
	ProjectParserARROW             = 44
	ProjectParserRANGE             = 45
	ProjectParserLEFT_PAR          = 46
	ProjectParserRIGHT_PAR         = 47
	ProjectParserLEFT_KEY          = 48
	ProjectParserRIGHT_KEY         = 49
	ProjectParserLEFT_BRACKET      = 50
	ProjectParserRIGHT_BRACKET     = 51
	ProjectParserEQUEAL_EQUAL      = 52
	ProjectParserNOTEQUAL          = 53
	ProjectParserGREATER_THAN      = 54
	ProjectParserLESS_THAN         = 55
	ProjectParserGREATER_EQUALTHAN = 56
	ProjectParserLESS_EQUEALTHAN   = 57
	ProjectParserMUL               = 58
	ProjectParserDIV               = 59
	ProjectParserMOD               = 60
	ProjectParserADD               = 61
	ProjectParserSUB               = 62
	ProjectParserAND               = 63
	ProjectParserOR                = 64
	ProjectParserPIPE              = 65
	ProjectParserEQUAL_ARROW       = 66
	ProjectParserUNDERSCORE        = 67
	ProjectParserWHITESPACE        = 68
)

// ProjectParser rules.
const (
	ProjectParserRULE_start            = 0
	ProjectParserRULE_listFuncs        = 1
	ProjectParserRULE_function         = 2
	ProjectParserRULE_funcMain         = 3
	ProjectParserRULE_bloq             = 4
	ProjectParserRULE_bloq_match       = 5
	ProjectParserRULE_listParams       = 6
	ProjectParserRULE_instructions     = 7
	ProjectParserRULE_instruction      = 8
	ProjectParserRULE_print_prod       = 9
	ProjectParserRULE_listVars         = 10
	ProjectParserRULE_declaration_prod = 11
	ProjectParserRULE_assign_prod      = 12
	ProjectParserRULE_ids_type         = 13
	ProjectParserRULE_listIds          = 14
	ProjectParserRULE_conditional_prod = 15
	ProjectParserRULE_if_prod          = 16
	ProjectParserRULE_list_else_if     = 17
	ProjectParserRULE_else_if          = 18
	ProjectParserRULE_match_prod       = 19
	ProjectParserRULE_list_instr_match = 20
	ProjectParserRULE_instr_match      = 21
	ProjectParserRULE_expr_match       = 22
	ProjectParserRULE_bucle_prod       = 23
	ProjectParserRULE_while_prod       = 24
	ProjectParserRULE_loop_prod        = 25
	ProjectParserRULE_forin_prod       = 26
	ProjectParserRULE_range_prod       = 27
	ProjectParserRULE_called_func      = 28
	ProjectParserRULE_listExpressions  = 29
	ProjectParserRULE_dec_arr          = 30
	ProjectParserRULE_listDim          = 31
	ProjectParserRULE_transfer_prod    = 32
	ProjectParserRULE_break_instr      = 33
	ProjectParserRULE_continue_instr   = 34
	ProjectParserRULE_return_instr     = 35
	ProjectParserRULE_expression       = 36
	ProjectParserRULE_arraydata        = 37
	ProjectParserRULE_access_array     = 38
	ProjectParserRULE_listInArray      = 39
	ProjectParserRULE_inArray          = 40
	ProjectParserRULE_expr_rel         = 41
	ProjectParserRULE_expr_arit        = 42
	ProjectParserRULE_expr_valor       = 43
	ProjectParserRULE_pow_op           = 44
	ProjectParserRULE_expr_logic       = 45
	ProjectParserRULE_expr_cast        = 46
	ProjectParserRULE_data_type        = 47
	ProjectParserRULE_primitive        = 48
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listFuncs returns the _listFuncs rule contexts.
	Get_listFuncs() IListFuncsContext

	// Set_listFuncs sets the _listFuncs rule contexts.
	Set_listFuncs(IListFuncsContext)

	// GetTree returns the tree attribute.
	GetTree() AST.Tree

	// SetTree sets the tree attribute.
	SetTree(AST.Tree)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	tree       AST.Tree
	_listFuncs IListFuncsContext
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

func (s *StartContext) Get_listFuncs() IListFuncsContext { return s._listFuncs }

func (s *StartContext) Set_listFuncs(v IListFuncsContext) { s._listFuncs = v }

func (s *StartContext) GetTree() AST.Tree { return s.tree }

func (s *StartContext) SetTree(v AST.Tree) { s.tree = v }

func (s *StartContext) ListFuncs() IListFuncsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListFuncsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListFuncsContext)
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
		p.SetState(98)

		var _x = p.listFuncs(0)

		localctx.(*StartContext)._listFuncs = _x
	}
	localctx.(*StartContext).tree = AST.NewTree(localctx.(*StartContext).Get_listFuncs().GetL())

	return localctx
}

// IListFuncsContext is an interface to support dynamic dispatch.
type IListFuncsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSubList returns the subList rule contexts.
	GetSubList() IListFuncsContext

	// Get_function returns the _function rule contexts.
	Get_function() IFunctionContext

	// SetSubList sets the subList rule contexts.
	SetSubList(IListFuncsContext)

	// Set_function sets the _function rule contexts.
	Set_function(IFunctionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListFuncsContext differentiates from other interfaces.
	IsListFuncsContext()
}

type ListFuncsContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	l         *arrayList.List
	subList   IListFuncsContext
	_function IFunctionContext
}

func NewEmptyListFuncsContext() *ListFuncsContext {
	var p = new(ListFuncsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_listFuncs
	return p
}

func (*ListFuncsContext) IsListFuncsContext() {}

func NewListFuncsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListFuncsContext {
	var p = new(ListFuncsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_listFuncs

	return p
}

func (s *ListFuncsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListFuncsContext) GetSubList() IListFuncsContext { return s.subList }

func (s *ListFuncsContext) Get_function() IFunctionContext { return s._function }

func (s *ListFuncsContext) SetSubList(v IListFuncsContext) { s.subList = v }

func (s *ListFuncsContext) Set_function(v IFunctionContext) { s._function = v }

func (s *ListFuncsContext) GetL() *arrayList.List { return s.l }

func (s *ListFuncsContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListFuncsContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ListFuncsContext) ListFuncs() IListFuncsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListFuncsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListFuncsContext)
}

func (s *ListFuncsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListFuncsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListFuncsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterListFuncs(s)
	}
}

func (s *ListFuncsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitListFuncs(s)
	}
}

func (p *ProjectParser) ListFuncs() (localctx IListFuncsContext) {
	return p.listFuncs(0)
}

func (p *ProjectParser) listFuncs(_p int) (localctx IListFuncsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListFuncsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListFuncsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ProjectParserRULE_listFuncs, _p)

	localctx.(*ListFuncsContext).l = arrayList.New()

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
		p.SetState(102)

		var _x = p.Function()

		localctx.(*ListFuncsContext)._function = _x
	}
	localctx.(*ListFuncsContext).l.Add(localctx.(*ListFuncsContext).Get_function().GetInstr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListFuncsContext(p, _parentctx, _parentState)
			localctx.(*ListFuncsContext).subList = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listFuncs)
			p.SetState(105)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(106)

				var _x = p.Function()

				localctx.(*ListFuncsContext)._function = _x
			}

			localctx.(*ListFuncsContext).GetSubList().GetL().Add(localctx.(*ListFuncsContext).Get_function().GetInstr())
			localctx.(*ListFuncsContext).l = localctx.(*ListFuncsContext).GetSubList().GetL()

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RFN returns the _RFN token.
	Get_RFN() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_RFN sets the _RFN token.
	Set_RFN(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_funcMain returns the _funcMain rule contexts.
	Get_funcMain() IFuncMainContext

	// Get_bloq returns the _bloq rule contexts.
	Get_bloq() IBloqContext

	// Get_data_type returns the _data_type rule contexts.
	Get_data_type() IData_typeContext

	// Get_listParams returns the _listParams rule contexts.
	Get_listParams() IListParamsContext

	// Set_funcMain sets the _funcMain rule contexts.
	Set_funcMain(IFuncMainContext)

	// Set_bloq sets the _bloq rule contexts.
	Set_bloq(IBloqContext)

	// Set_data_type sets the _data_type rule contexts.
	Set_data_type(IData_typeContext)

	// Set_listParams sets the _listParams rule contexts.
	Set_listParams(IListParamsContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_funcMain   IFuncMainContext
	_RFN        antlr.Token
	_ID         antlr.Token
	_bloq       IBloqContext
	_data_type  IData_typeContext
	_listParams IListParamsContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Get_RFN() antlr.Token { return s._RFN }

func (s *FunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *FunctionContext) Set_RFN(v antlr.Token) { s._RFN = v }

func (s *FunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FunctionContext) Get_funcMain() IFuncMainContext { return s._funcMain }

func (s *FunctionContext) Get_bloq() IBloqContext { return s._bloq }

func (s *FunctionContext) Get_data_type() IData_typeContext { return s._data_type }

func (s *FunctionContext) Get_listParams() IListParamsContext { return s._listParams }

func (s *FunctionContext) Set_funcMain(v IFuncMainContext) { s._funcMain = v }

func (s *FunctionContext) Set_bloq(v IBloqContext) { s._bloq = v }

func (s *FunctionContext) Set_data_type(v IData_typeContext) { s._data_type = v }

func (s *FunctionContext) Set_listParams(v IListParamsContext) { s._listParams = v }

func (s *FunctionContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *FunctionContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *FunctionContext) FuncMain() IFuncMainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncMainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncMainContext)
}

func (s *FunctionContext) RFN() antlr.TerminalNode {
	return s.GetToken(ProjectParserRFN, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(ProjectParserID, 0)
}

func (s *FunctionContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *FunctionContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_PAR, 0)
}

func (s *FunctionContext) Bloq() IBloqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqContext)
}

func (s *FunctionContext) ARROW() antlr.TerminalNode {
	return s.GetToken(ProjectParserARROW, 0)
}

func (s *FunctionContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *FunctionContext) ListParams() IListParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *ProjectParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ProjectParserRULE_function)
	listParams := arrayList.New()

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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)

			var _x = p.FuncMain()

			localctx.(*FunctionContext)._funcMain = _x
		}
		localctx.(*FunctionContext).instr = localctx.(*FunctionContext).Get_funcMain().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)

			var _m = p.Match(ProjectParserRFN)

			localctx.(*FunctionContext)._RFN = _m
		}
		{
			p.SetState(118)

			var _m = p.Match(ProjectParserID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(119)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(120)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(121)

			var _x = p.Bloq()

			localctx.(*FunctionContext)._bloq = _x
		}

		localctx.(*FunctionContext).instr = Environment.NewFunction((func() int {
			if localctx.(*FunctionContext).Get_RFN() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_RFN().GetLine()
			}
		}()), localctx.(*FunctionContext).Get_RFN().GetColumn(), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), listParams, localctx.(*FunctionContext).Get_bloq().GetContent(), "void")

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)

			var _m = p.Match(ProjectParserRFN)

			localctx.(*FunctionContext)._RFN = _m
		}
		{
			p.SetState(125)

			var _m = p.Match(ProjectParserID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(126)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(127)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(128)
			p.Match(ProjectParserARROW)
		}
		{
			p.SetState(129)

			var _x = p.Data_type()

			localctx.(*FunctionContext)._data_type = _x
		}
		{
			p.SetState(130)

			var _x = p.Bloq()

			localctx.(*FunctionContext)._bloq = _x
		}

		localctx.(*FunctionContext).instr = Environment.NewFunction((func() int {
			if localctx.(*FunctionContext).Get_RFN() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_RFN().GetLine()
			}
		}()), localctx.(*FunctionContext).Get_RFN().GetColumn(), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), listParams, localctx.(*FunctionContext).Get_bloq().GetContent(), localctx.(*FunctionContext).Get_data_type().GetData())

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(133)

			var _m = p.Match(ProjectParserRFN)

			localctx.(*FunctionContext)._RFN = _m
		}
		{
			p.SetState(134)

			var _m = p.Match(ProjectParserID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(135)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(136)

			var _x = p.listParams(0)

			localctx.(*FunctionContext)._listParams = _x
		}
		{
			p.SetState(137)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(138)

			var _x = p.Bloq()

			localctx.(*FunctionContext)._bloq = _x
		}

		localctx.(*FunctionContext).instr = Environment.NewFunction((func() int {
			if localctx.(*FunctionContext).Get_RFN() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_RFN().GetLine()
			}
		}()), localctx.(*FunctionContext).Get_RFN().GetColumn(), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParams().GetL(), localctx.(*FunctionContext).Get_bloq().GetContent(), "void")

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(141)

			var _m = p.Match(ProjectParserRFN)

			localctx.(*FunctionContext)._RFN = _m
		}
		{
			p.SetState(142)

			var _m = p.Match(ProjectParserID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(143)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(144)

			var _x = p.listParams(0)

			localctx.(*FunctionContext)._listParams = _x
		}
		{
			p.SetState(145)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(146)
			p.Match(ProjectParserARROW)
		}
		{
			p.SetState(147)

			var _x = p.Data_type()

			localctx.(*FunctionContext)._data_type = _x
		}
		{
			p.SetState(148)

			var _x = p.Bloq()

			localctx.(*FunctionContext)._bloq = _x
		}

		localctx.(*FunctionContext).instr = Environment.NewFunction((func() int {
			if localctx.(*FunctionContext).Get_RFN() == nil {
				return 0
			} else {
				return localctx.(*FunctionContext).Get_RFN().GetLine()
			}
		}()), localctx.(*FunctionContext).Get_RFN().GetColumn(), (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_listParams().GetL(), localctx.(*FunctionContext).Get_bloq().GetContent(), localctx.(*FunctionContext).Get_data_type().GetData())

	}

	return localctx
}

// IFuncMainContext is an interface to support dynamic dispatch.
type IFuncMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RFN returns the _RFN token.
	Get_RFN() antlr.Token

	// Set_RFN sets the _RFN token.
	Set_RFN(antlr.Token)

	// Get_bloq returns the _bloq rule contexts.
	Get_bloq() IBloqContext

	// Set_bloq sets the _bloq rule contexts.
	Set_bloq(IBloqContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsFuncMainContext differentiates from other interfaces.
	IsFuncMainContext()
}

type FuncMainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	instr  Abstract.Instruction
	_RFN   antlr.Token
	_bloq  IBloqContext
}

func NewEmptyFuncMainContext() *FuncMainContext {
	var p = new(FuncMainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_funcMain
	return p
}

func (*FuncMainContext) IsFuncMainContext() {}

func NewFuncMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncMainContext {
	var p = new(FuncMainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_funcMain

	return p
}

func (s *FuncMainContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncMainContext) Get_RFN() antlr.Token { return s._RFN }

func (s *FuncMainContext) Set_RFN(v antlr.Token) { s._RFN = v }

func (s *FuncMainContext) Get_bloq() IBloqContext { return s._bloq }

func (s *FuncMainContext) Set_bloq(v IBloqContext) { s._bloq = v }

func (s *FuncMainContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *FuncMainContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *FuncMainContext) RFN() antlr.TerminalNode {
	return s.GetToken(ProjectParserRFN, 0)
}

func (s *FuncMainContext) RMAIN() antlr.TerminalNode {
	return s.GetToken(ProjectParserRMAIN, 0)
}

func (s *FuncMainContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *FuncMainContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_PAR, 0)
}

func (s *FuncMainContext) Bloq() IBloqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqContext)
}

func (s *FuncMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncMainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncMainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterFuncMain(s)
	}
}

func (s *FuncMainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitFuncMain(s)
	}
}

func (p *ProjectParser) FuncMain() (localctx IFuncMainContext) {
	localctx = NewFuncMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ProjectParserRULE_funcMain)
	listParams := arrayList.New()

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
		p.SetState(153)

		var _m = p.Match(ProjectParserRFN)

		localctx.(*FuncMainContext)._RFN = _m
	}
	{
		p.SetState(154)
		p.Match(ProjectParserRMAIN)
	}
	{
		p.SetState(155)
		p.Match(ProjectParserLEFT_PAR)
	}
	{
		p.SetState(156)
		p.Match(ProjectParserRIGHT_PAR)
	}
	{
		p.SetState(157)

		var _x = p.Bloq()

		localctx.(*FuncMainContext)._bloq = _x
	}
	localctx.(*FuncMainContext).instr = Environment.NewFunction((func() int {
		if localctx.(*FuncMainContext).Get_RFN() == nil {
			return 0
		} else {
			return localctx.(*FuncMainContext).Get_RFN().GetLine()
		}
	}()), localctx.(*FuncMainContext).Get_RFN().GetColumn(), "main", listParams, localctx.(*FuncMainContext).Get_bloq().GetContent(), "void")

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

func (s *BloqContext) LEFT_KEY() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_KEY, 0)
}

func (s *BloqContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *BloqContext) RIGHT_KEY() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_KEY, 0)
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
	p.EnterRule(localctx, 8, ProjectParserRULE_bloq)

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

	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Match(ProjectParserLEFT_KEY)
		}
		{
			p.SetState(161)

			var _x = p.Instructions()

			localctx.(*BloqContext)._instructions = _x
		}
		{
			p.SetState(162)
			p.Match(ProjectParserRIGHT_KEY)
		}
		localctx.(*BloqContext).content = localctx.(*BloqContext).Get_instructions().GetL()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(ProjectParserLEFT_KEY)
		}
		{
			p.SetState(166)
			p.Match(ProjectParserRIGHT_KEY)
		}
		localctx.(*BloqContext).content = arrayList.New()

	}

	return localctx
}

// IBloq_matchContext is an interface to support dynamic dispatch.
type IBloq_matchContext interface {
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

	// IsBloq_matchContext differentiates from other interfaces.
	IsBloq_matchContext()
}

type Bloq_matchContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	content       *arrayList.List
	_instructions IInstructionsContext
}

func NewEmptyBloq_matchContext() *Bloq_matchContext {
	var p = new(Bloq_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_bloq_match
	return p
}

func (*Bloq_matchContext) IsBloq_matchContext() {}

func NewBloq_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bloq_matchContext {
	var p = new(Bloq_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_bloq_match

	return p
}

func (s *Bloq_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Bloq_matchContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *Bloq_matchContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *Bloq_matchContext) GetContent() *arrayList.List { return s.content }

func (s *Bloq_matchContext) SetContent(v *arrayList.List) { s.content = v }

func (s *Bloq_matchContext) LEFT_KEY() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_KEY, 0)
}

func (s *Bloq_matchContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *Bloq_matchContext) RIGHT_KEY() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_KEY, 0)
}

func (s *Bloq_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bloq_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bloq_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterBloq_match(s)
	}
}

func (s *Bloq_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitBloq_match(s)
	}
}

func (p *ProjectParser) Bloq_match() (localctx IBloq_matchContext) {
	localctx = NewBloq_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ProjectParserRULE_bloq_match)

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

	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.Match(ProjectParserLEFT_KEY)
		}
		{
			p.SetState(171)

			var _x = p.Instructions()

			localctx.(*Bloq_matchContext)._instructions = _x
		}
		{
			p.SetState(172)
			p.Match(ProjectParserRIGHT_KEY)
		}
		localctx.(*Bloq_matchContext).content = localctx.(*Bloq_matchContext).Get_instructions().GetL()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)

			var _x = p.Instructions()

			localctx.(*Bloq_matchContext)._instructions = _x
		}
		localctx.(*Bloq_matchContext).content = localctx.(*Bloq_matchContext).Get_instructions().GetL()

	}

	return localctx
}

// IListParamsContext is an interface to support dynamic dispatch.
type IListParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetList returns the List rule contexts.
	GetList() IListParamsContext

	// Get_data_type returns the _data_type rule contexts.
	Get_data_type() IData_typeContext

	// SetList sets the List rule contexts.
	SetList(IListParamsContext)

	// Set_data_type sets the _data_type rule contexts.
	Set_data_type(IData_typeContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListParamsContext differentiates from other interfaces.
	IsListParamsContext()
}

type ListParamsContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	l          *arrayList.List
	List       IListParamsContext
	_ID        antlr.Token
	_data_type IData_typeContext
}

func NewEmptyListParamsContext() *ListParamsContext {
	var p = new(ListParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_listParams
	return p
}

func (*ListParamsContext) IsListParamsContext() {}

func NewListParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListParamsContext {
	var p = new(ListParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_listParams

	return p
}

func (s *ListParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListParamsContext) Get_ID() antlr.Token { return s._ID }

func (s *ListParamsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListParamsContext) GetList() IListParamsContext { return s.List }

func (s *ListParamsContext) Get_data_type() IData_typeContext { return s._data_type }

func (s *ListParamsContext) SetList(v IListParamsContext) { s.List = v }

func (s *ListParamsContext) Set_data_type(v IData_typeContext) { s._data_type = v }

func (s *ListParamsContext) GetL() *arrayList.List { return s.l }

func (s *ListParamsContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListParamsContext) ID() antlr.TerminalNode {
	return s.GetToken(ProjectParserID, 0)
}

func (s *ListParamsContext) COLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOLON, 0)
}

func (s *ListParamsContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *ListParamsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *ListParamsContext) ListParams() IListParamsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListParamsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListParamsContext)
}

func (s *ListParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterListParams(s)
	}
}

func (s *ListParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitListParams(s)
	}
}

func (p *ProjectParser) ListParams() (localctx IListParamsContext) {
	return p.listParams(0)
}

func (p *ProjectParser) listParams(_p int) (localctx IListParamsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListParamsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListParamsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, ProjectParserRULE_listParams, _p)

	localctx.(*ListParamsContext).l = arrayList.New()

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
		p.SetState(181)

		var _m = p.Match(ProjectParserID)

		localctx.(*ListParamsContext)._ID = _m
	}
	{
		p.SetState(182)
		p.Match(ProjectParserCOLON)
	}
	{
		p.SetState(183)

		var _x = p.Data_type()

		localctx.(*ListParamsContext)._data_type = _x
	}

	listIds := arrayList.New()
	listIds.Add(Expression.NewIdentifier((func() string {
		if localctx.(*ListParamsContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListParamsContext).Get_ID().GetText()
		}
	}()), (func() int {
		if localctx.(*ListParamsContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*ListParamsContext).Get_ID().GetLine()
		}
	}()), localctx.(*ListParamsContext).Get_ID().GetColumn()))
	newDec := Natives.NewDeclaration(listIds, false, localctx.(*ListParamsContext).Get_data_type().GetData())
	localctx.(*ListParamsContext).l.Add(newDec)

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListParamsContext(p, _parentctx, _parentState)
			localctx.(*ListParamsContext).List = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listParams)
			p.SetState(186)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(187)
				p.Match(ProjectParserCOMMA)
			}
			{
				p.SetState(188)

				var _m = p.Match(ProjectParserID)

				localctx.(*ListParamsContext)._ID = _m
			}
			{
				p.SetState(189)
				p.Match(ProjectParserCOLON)
			}
			{
				p.SetState(190)

				var _x = p.Data_type()

				localctx.(*ListParamsContext)._data_type = _x
			}

			listIds := arrayList.New()
			listIds.Add(Expression.NewIdentifier((func() string {
				if localctx.(*ListParamsContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*ListParamsContext).Get_ID().GetText()
				}
			}()), (func() int {
				if localctx.(*ListParamsContext).Get_ID() == nil {
					return 0
				} else {
					return localctx.(*ListParamsContext).Get_ID().GetLine()
				}
			}()), localctx.(*ListParamsContext).Get_ID().GetColumn()))
			newDec := Natives.NewDeclaration(listIds, false, localctx.(*ListParamsContext).Get_data_type().GetData())
			localctx.(*ListParamsContext).GetList().GetL().Add(newDec)
			localctx.(*ListParamsContext).l = localctx.(*ListParamsContext).GetList().GetL()

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

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
	p.EnterRule(localctx, 14, ProjectParserRULE_instructions)

	localctx.(*InstructionsContext).l = arrayList.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(198)

				var _x = p.Instruction()

				localctx.(*InstructionsContext)._instruction = _x
			}
			localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
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

	// Get_bucle_prod returns the _bucle_prod rule contexts.
	Get_bucle_prod() IBucle_prodContext

	// Get_expr_rel returns the _expr_rel rule contexts.
	Get_expr_rel() IExpr_relContext

	// Get_dec_arr returns the _dec_arr rule contexts.
	Get_dec_arr() IDec_arrContext

	// Get_transfer_prod returns the _transfer_prod rule contexts.
	Get_transfer_prod() ITransfer_prodContext

	// Set_print_prod sets the _print_prod rule contexts.
	Set_print_prod(IPrint_prodContext)

	// Set_declaration_prod sets the _declaration_prod rule contexts.
	Set_declaration_prod(IDeclaration_prodContext)

	// Set_assign_prod sets the _assign_prod rule contexts.
	Set_assign_prod(IAssign_prodContext)

	// Set_conditional_prod sets the _conditional_prod rule contexts.
	Set_conditional_prod(IConditional_prodContext)

	// Set_bucle_prod sets the _bucle_prod rule contexts.
	Set_bucle_prod(IBucle_prodContext)

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

	// Set_dec_arr sets the _dec_arr rule contexts.
	Set_dec_arr(IDec_arrContext)

	// Set_transfer_prod sets the _transfer_prod rule contexts.
	Set_transfer_prod(ITransfer_prodContext)

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
	_bucle_prod       IBucle_prodContext
	_expr_rel         IExpr_relContext
	_dec_arr          IDec_arrContext
	_transfer_prod    ITransfer_prodContext
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

func (s *InstructionContext) Get_bucle_prod() IBucle_prodContext { return s._bucle_prod }

func (s *InstructionContext) Get_expr_rel() IExpr_relContext { return s._expr_rel }

func (s *InstructionContext) Get_dec_arr() IDec_arrContext { return s._dec_arr }

func (s *InstructionContext) Get_transfer_prod() ITransfer_prodContext { return s._transfer_prod }

func (s *InstructionContext) Set_print_prod(v IPrint_prodContext) { s._print_prod = v }

func (s *InstructionContext) Set_declaration_prod(v IDeclaration_prodContext) {
	s._declaration_prod = v
}

func (s *InstructionContext) Set_assign_prod(v IAssign_prodContext) { s._assign_prod = v }

func (s *InstructionContext) Set_conditional_prod(v IConditional_prodContext) {
	s._conditional_prod = v
}

func (s *InstructionContext) Set_bucle_prod(v IBucle_prodContext) { s._bucle_prod = v }

func (s *InstructionContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *InstructionContext) Set_dec_arr(v IDec_arrContext) { s._dec_arr = v }

func (s *InstructionContext) Set_transfer_prod(v ITransfer_prodContext) { s._transfer_prod = v }

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

func (s *InstructionContext) Bucle_prod() IBucle_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBucle_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBucle_prodContext)
}

func (s *InstructionContext) Expr_rel() IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *InstructionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *InstructionContext) Dec_arr() IDec_arrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDec_arrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDec_arrContext)
}

func (s *InstructionContext) Transfer_prod() ITransfer_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransfer_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransfer_prodContext)
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
	p.EnterRule(localctx, 16, ProjectParserRULE_instruction)

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

	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)

			var _x = p.Print_prod()

			localctx.(*InstructionContext)._print_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_print_prod().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)

			var _x = p.Declaration_prod()

			localctx.(*InstructionContext)._declaration_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_declaration_prod().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(212)

			var _x = p.Assign_prod()

			localctx.(*InstructionContext)._assign_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_assign_prod().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(215)

			var _x = p.Conditional_prod()

			localctx.(*InstructionContext)._conditional_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_conditional_prod().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)

			var _x = p.Bucle_prod()

			localctx.(*InstructionContext)._bucle_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_bucle_prod().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(221)

			var _x = p.expr_rel(0)

			localctx.(*InstructionContext)._expr_rel = _x
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(222)
				p.Match(ProjectParserCOMMA)
			}

		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_expr_rel().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(227)

			var _x = p.Dec_arr()

			localctx.(*InstructionContext)._dec_arr = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_dec_arr().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(230)

			var _x = p.Transfer_prod()

			localctx.(*InstructionContext)._transfer_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_transfer_prod().GetInstr()

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

func (s *Print_prodContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ProjectParserCOMMA)
}

func (s *Print_prodContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, i)
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
	p.EnterRule(localctx, 18, ProjectParserRULE_print_prod)
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

	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.Match(ProjectParserPRINT)
		}
		{
			p.SetState(236)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(237)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(238)

			var _x = p.Expression()

			localctx.(*Print_prodContext)._expression = _x
		}
		{
			p.SetState(239)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(240)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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
			p.SetState(243)
			p.Match(ProjectParserPRINT)
		}
		{
			p.SetState(244)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(245)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(246)

			var _x = p.Expression()

			localctx.(*Print_prodContext).opBefore = _x
		}
		{
			p.SetState(247)
			p.Match(ProjectParserCOMMA)
		}
		{
			p.SetState(248)

			var _x = p.listVars(0)

			localctx.(*Print_prodContext)._listVars = _x
		}
		{
			p.SetState(249)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(250)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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
			p.SetState(253)
			p.Match(ProjectParserPRINTLN)
		}
		{
			p.SetState(254)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(255)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(256)

			var _x = p.Expression()

			localctx.(*Print_prodContext)._expression = _x
		}
		{
			p.SetState(257)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(258)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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
			p.SetState(261)
			p.Match(ProjectParserPRINTLN)
		}
		{
			p.SetState(262)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(263)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(264)

			var _x = p.Expression()

			localctx.(*Print_prodContext).opBefore = _x
		}
		{
			p.SetState(265)
			p.Match(ProjectParserCOMMA)
		}
		{
			p.SetState(266)

			var _x = p.listVars(0)

			localctx.(*Print_prodContext)._listVars = _x
		}
		{
			p.SetState(267)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(268)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, ProjectParserRULE_listVars, _p)

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
		p.SetState(274)

		var _x = p.Expression()

		localctx.(*ListVarsContext)._expression = _x
	}

	localctx.(*ListVarsContext).list.Add(localctx.(*ListVarsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListVarsContext(p, _parentctx, _parentState)
			localctx.(*ListVarsContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listVars)
			p.SetState(277)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(278)
				p.Match(ProjectParserCOMMA)
			}
			{
				p.SetState(279)

				var _x = p.Expression()

				localctx.(*ListVarsContext)._expression = _x
			}

			localctx.(*ListVarsContext).GetSub().GetList().Add(localctx.(*ListVarsContext).Get_expression().GetP())
			localctx.(*ListVarsContext).list = localctx.(*ListVarsContext).GetSub().GetList()

		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IDeclaration_prodContext is an interface to support dynamic dispatch.
type IDeclaration_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MUT returns the _MUT token.
	Get_MUT() antlr.Token

	// Set_MUT sets the _MUT token.
	Set_MUT(antlr.Token)

	// Get_ids_type returns the _ids_type rule contexts.
	Get_ids_type() IIds_typeContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_data_type returns the _data_type rule contexts.
	Get_data_type() IData_typeContext

	// Set_ids_type sets the _ids_type rule contexts.
	Set_ids_type(IIds_typeContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_data_type sets the _data_type rule contexts.
	Set_data_type(IData_typeContext)

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
	_MUT        antlr.Token
	_ids_type   IIds_typeContext
	_expression IExpressionContext
	_data_type  IData_typeContext
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

func (s *Declaration_prodContext) Get_MUT() antlr.Token { return s._MUT }

func (s *Declaration_prodContext) Set_MUT(v antlr.Token) { s._MUT = v }

func (s *Declaration_prodContext) Get_ids_type() IIds_typeContext { return s._ids_type }

func (s *Declaration_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Declaration_prodContext) Get_data_type() IData_typeContext { return s._data_type }

func (s *Declaration_prodContext) Set_ids_type(v IIds_typeContext) { s._ids_type = v }

func (s *Declaration_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Declaration_prodContext) Set_data_type(v IData_typeContext) { s._data_type = v }

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

func (s *Declaration_prodContext) MUT() antlr.TerminalNode {
	return s.GetToken(ProjectParserMUT, 0)
}

func (s *Declaration_prodContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *Declaration_prodContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *Declaration_prodContext) COLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOLON, 0)
}

func (s *Declaration_prodContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
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
	p.EnterRule(localctx, 22, ProjectParserRULE_declaration_prod)
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

	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)
			p.Match(ProjectParserDECLARAR)
		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProjectParserMUT {
			{
				p.SetState(288)

				var _m = p.Match(ProjectParserMUT)

				localctx.(*Declaration_prodContext)._MUT = _m
			}

		}
		{
			p.SetState(291)

			var _x = p.Ids_type()

			localctx.(*Declaration_prodContext)._ids_type = _x
		}
		{
			p.SetState(292)
			p.Match(ProjectParserEQUAL)
		}
		{
			p.SetState(293)

			var _x = p.Expression()

			localctx.(*Declaration_prodContext)._expression = _x
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(294)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

		if (func() string {
			if localctx.(*Declaration_prodContext).Get_MUT() == nil {
				return ""
			} else {
				return localctx.(*Declaration_prodContext).Get_MUT().GetText()
			}
		}()) != "" {
			localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP(), true, "")
		} else {
			localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP(), false, "")
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(299)
			p.Match(ProjectParserDECLARAR)
		}
		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProjectParserMUT {
			{
				p.SetState(300)

				var _m = p.Match(ProjectParserMUT)

				localctx.(*Declaration_prodContext)._MUT = _m
			}

		}
		{
			p.SetState(303)

			var _x = p.Ids_type()

			localctx.(*Declaration_prodContext)._ids_type = _x
		}
		{
			p.SetState(304)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(305)

			var _x = p.Data_type()

			localctx.(*Declaration_prodContext)._data_type = _x
		}
		{
			p.SetState(306)
			p.Match(ProjectParserEQUAL)
		}
		{
			p.SetState(307)

			var _x = p.Expression()

			localctx.(*Declaration_prodContext)._expression = _x
		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(308)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

		if (func() string {
			if localctx.(*Declaration_prodContext).Get_MUT() == nil {
				return ""
			} else {
				return localctx.(*Declaration_prodContext).Get_MUT().GetText()
			}
		}()) != "" {
			localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP(), true, localctx.(*Declaration_prodContext).Get_data_type().GetData())
		} else {
			localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP(), false, localctx.(*Declaration_prodContext).Get_data_type().GetData())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(313)
			p.Match(ProjectParserDECLARAR)
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProjectParserMUT {
			{
				p.SetState(314)

				var _m = p.Match(ProjectParserMUT)

				localctx.(*Declaration_prodContext)._MUT = _m
			}

		}
		{
			p.SetState(317)

			var _x = p.Ids_type()

			localctx.(*Declaration_prodContext)._ids_type = _x
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(318)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

		if (func() string {
			if localctx.(*Declaration_prodContext).Get_MUT() == nil {
				return ""
			} else {
				return localctx.(*Declaration_prodContext).Get_MUT().GetText()
			}
		}()) != "" {
			localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), true, "")
		} else {
			localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), true, "")
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(323)
			p.Match(ProjectParserDECLARAR)
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProjectParserMUT {
			{
				p.SetState(324)

				var _m = p.Match(ProjectParserMUT)

				localctx.(*Declaration_prodContext)._MUT = _m
			}

		}
		{
			p.SetState(327)

			var _x = p.Ids_type()

			localctx.(*Declaration_prodContext)._ids_type = _x
		}
		{
			p.SetState(328)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(329)

			var _x = p.Data_type()

			localctx.(*Declaration_prodContext)._data_type = _x
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(330)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

		if (func() string {
			if localctx.(*Declaration_prodContext).Get_MUT() == nil {
				return ""
			} else {
				return localctx.(*Declaration_prodContext).Get_MUT().GetText()
			}
		}()) != "" {
			localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), true, localctx.(*Declaration_prodContext).Get_data_type().GetData())
		} else {
			localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), true, localctx.(*Declaration_prodContext).Get_data_type().GetData())
		}

	}

	return localctx
}

// IAssign_prodContext is an interface to support dynamic dispatch.
type IAssign_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listIds returns the _listIds rule contexts.
	Get_listIds() IListIdsContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_listInArray returns the _listInArray rule contexts.
	Get_listInArray() IListInArrayContext

	// Set_listIds sets the _listIds rule contexts.
	Set_listIds(IListIdsContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_listInArray sets the _listInArray rule contexts.
	Set_listInArray(IListInArrayContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsAssign_prodContext differentiates from other interfaces.
	IsAssign_prodContext()
}

type Assign_prodContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        Abstract.Instruction
	_listIds     IListIdsContext
	_expression  IExpressionContext
	_ID          antlr.Token
	_listInArray IListInArrayContext
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

func (s *Assign_prodContext) Get_ID() antlr.Token { return s._ID }

func (s *Assign_prodContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Assign_prodContext) Get_listIds() IListIdsContext { return s._listIds }

func (s *Assign_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Assign_prodContext) Get_listInArray() IListInArrayContext { return s._listInArray }

func (s *Assign_prodContext) Set_listIds(v IListIdsContext) { s._listIds = v }

func (s *Assign_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Assign_prodContext) Set_listInArray(v IListInArrayContext) { s._listInArray = v }

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

func (s *Assign_prodContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *Assign_prodContext) ID() antlr.TerminalNode {
	return s.GetToken(ProjectParserID, 0)
}

func (s *Assign_prodContext) ListInArray() IListInArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListInArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListInArrayContext)
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
	p.EnterRule(localctx, 24, ProjectParserRULE_assign_prod)
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

	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(337)

			var _x = p.listIds(0)

			localctx.(*Assign_prodContext)._listIds = _x
		}
		{
			p.SetState(338)
			p.Match(ProjectParserEQUAL)
		}
		{
			p.SetState(339)

			var _x = p.Expression()

			localctx.(*Assign_prodContext)._expression = _x
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(340)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

		localctx.(*Assign_prodContext).instr = Natives.NewAssign(localctx.(*Assign_prodContext).Get_listIds().GetList(), localctx.(*Assign_prodContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)

			var _m = p.Match(ProjectParserID)

			localctx.(*Assign_prodContext)._ID = _m
		}
		{
			p.SetState(346)

			var _x = p.listInArray(0)

			localctx.(*Assign_prodContext)._listInArray = _x
		}
		{
			p.SetState(347)
			p.Match(ProjectParserEQUAL)
		}
		{
			p.SetState(348)

			var _x = p.Expression()

			localctx.(*Assign_prodContext)._expression = _x
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(349)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

		localctx.(*Assign_prodContext).instr = DecArrays.NewAssignArray((func() string {
			if localctx.(*Assign_prodContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Assign_prodContext).Get_ID().GetText()
			}
		}()), localctx.(*Assign_prodContext).Get_listInArray().GetL(), localctx.(*Assign_prodContext).Get_expression().GetP())

	}

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
	p.EnterRule(localctx, 26, ProjectParserRULE_ids_type)

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

	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		localctx.(*Ids_typeContext).list = localctx.(*Ids_typeContext).Get_listIds().GetList()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		{
			p.SetState(360)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(361)
			p.Match(ProjectParserRSTRING)
		}
		localctx.(*Ids_typeContext).list = localctx.(*Ids_typeContext).Get_listIds().GetList()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(364)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		{
			p.SetState(365)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(366)
			p.Match(ProjectParserREFERENCE)
		}
		{
			p.SetState(367)
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
	_startState := 28
	p.EnterRecursionRule(localctx, 28, ProjectParserRULE_listIds, _p)

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
		p.SetState(373)

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
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListIdsContext(p, _parentctx, _parentState)
			localctx.(*ListIdsContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listIds)
			p.SetState(376)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(377)
				p.Match(ProjectParserCOMMA)
			}
			{
				p.SetState(378)

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
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IConditional_prodContext is an interface to support dynamic dispatch.
type IConditional_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_if_prod returns the _if_prod rule contexts.
	Get_if_prod() IIf_prodContext

	// Get_match_prod returns the _match_prod rule contexts.
	Get_match_prod() IMatch_prodContext

	// Set_if_prod sets the _if_prod rule contexts.
	Set_if_prod(IIf_prodContext)

	// Set_match_prod sets the _match_prod rule contexts.
	Set_match_prod(IMatch_prodContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsConditional_prodContext differentiates from other interfaces.
	IsConditional_prodContext()
}

type Conditional_prodContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	p           Abstract.Expression
	_if_prod    IIf_prodContext
	_match_prod IMatch_prodContext
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

func (s *Conditional_prodContext) Get_if_prod() IIf_prodContext { return s._if_prod }

func (s *Conditional_prodContext) Get_match_prod() IMatch_prodContext { return s._match_prod }

func (s *Conditional_prodContext) Set_if_prod(v IIf_prodContext) { s._if_prod = v }

func (s *Conditional_prodContext) Set_match_prod(v IMatch_prodContext) { s._match_prod = v }

func (s *Conditional_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Conditional_prodContext) GetP() Abstract.Expression { return s.p }

func (s *Conditional_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Conditional_prodContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Conditional_prodContext) If_prod() IIf_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_prodContext)
}

func (s *Conditional_prodContext) Match_prod() IMatch_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatch_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatch_prodContext)
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
	p.EnterRule(localctx, 30, ProjectParserRULE_conditional_prod)

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

	p.SetState(391)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserRIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)

			var _x = p.If_prod()

			localctx.(*Conditional_prodContext)._if_prod = _x
		}

		localctx.(*Conditional_prodContext).instr = localctx.(*Conditional_prodContext).Get_if_prod().GetInstr()
		localctx.(*Conditional_prodContext).p = localctx.(*Conditional_prodContext).Get_if_prod().GetP()

	case ProjectParserRMATCH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)

			var _x = p.Match_prod()

			localctx.(*Conditional_prodContext)._match_prod = _x
		}

		localctx.(*Conditional_prodContext).instr = localctx.(*Conditional_prodContext).Get_match_prod().GetInstr()
		localctx.(*Conditional_prodContext).p = localctx.(*Conditional_prodContext).Get_match_prod().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIf_prodContext is an interface to support dynamic dispatch.
type IIf_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RIF returns the _RIF token.
	Get_RIF() antlr.Token

	// Set_RIF sets the _RIF token.
	Set_RIF(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_bloq returns the _bloq rule contexts.
	Get_bloq() IBloqContext

	// GetBif returns the bif rule contexts.
	GetBif() IBloqContext

	// GetBelse returns the belse rule contexts.
	GetBelse() IBloqContext

	// Get_list_else_if returns the _list_else_if rule contexts.
	Get_list_else_if() IList_else_ifContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_bloq sets the _bloq rule contexts.
	Set_bloq(IBloqContext)

	// SetBif sets the bif rule contexts.
	SetBif(IBloqContext)

	// SetBelse sets the belse rule contexts.
	SetBelse(IBloqContext)

	// Set_list_else_if sets the _list_else_if rule contexts.
	Set_list_else_if(IList_else_ifContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsIf_prodContext differentiates from other interfaces.
	IsIf_prodContext()
}

type If_prodContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	instr         Abstract.Instruction
	p             Abstract.Expression
	_RIF          antlr.Token
	_expression   IExpressionContext
	_bloq         IBloqContext
	bif           IBloqContext
	belse         IBloqContext
	_list_else_if IList_else_ifContext
}

func NewEmptyIf_prodContext() *If_prodContext {
	var p = new(If_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_if_prod
	return p
}

func (*If_prodContext) IsIf_prodContext() {}

func NewIf_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_prodContext {
	var p = new(If_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_if_prod

	return p
}

func (s *If_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *If_prodContext) Get_RIF() antlr.Token { return s._RIF }

func (s *If_prodContext) Set_RIF(v antlr.Token) { s._RIF = v }

func (s *If_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *If_prodContext) Get_bloq() IBloqContext { return s._bloq }

func (s *If_prodContext) GetBif() IBloqContext { return s.bif }

func (s *If_prodContext) GetBelse() IBloqContext { return s.belse }

func (s *If_prodContext) Get_list_else_if() IList_else_ifContext { return s._list_else_if }

func (s *If_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *If_prodContext) Set_bloq(v IBloqContext) { s._bloq = v }

func (s *If_prodContext) SetBif(v IBloqContext) { s.bif = v }

func (s *If_prodContext) SetBelse(v IBloqContext) { s.belse = v }

func (s *If_prodContext) Set_list_else_if(v IList_else_ifContext) { s._list_else_if = v }

func (s *If_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *If_prodContext) GetP() Abstract.Expression { return s.p }

func (s *If_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *If_prodContext) SetP(v Abstract.Expression) { s.p = v }

func (s *If_prodContext) RIF() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIF, 0)
}

func (s *If_prodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_prodContext) AllBloq() []IBloqContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqContext)(nil)).Elem())
	var tst = make([]IBloqContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqContext)
		}
	}

	return tst
}

func (s *If_prodContext) Bloq(i int) IBloqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqContext)
}

func (s *If_prodContext) RELSE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRELSE, 0)
}

func (s *If_prodContext) List_else_if() IList_else_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_else_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_else_ifContext)
}

func (s *If_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterIf_prod(s)
	}
}

func (s *If_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitIf_prod(s)
	}
}

func (p *ProjectParser) If_prod() (localctx IIf_prodContext) {
	localctx = NewIf_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ProjectParserRULE_if_prod)

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

	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(393)

			var _m = p.Match(ProjectParserRIF)

			localctx.(*If_prodContext)._RIF = _m
		}
		{
			p.SetState(394)

			var _x = p.Expression()

			localctx.(*If_prodContext)._expression = _x
		}
		{
			p.SetState(395)

			var _x = p.Bloq()

			localctx.(*If_prodContext)._bloq = _x
		}

		localctx.(*If_prodContext).instr = Natives.NewIf(localctx.(*If_prodContext).Get_expression().GetP(), localctx.(*If_prodContext).Get_bloq().GetContent(), nil, nil, (func() int {
			if localctx.(*If_prodContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*If_prodContext).Get_RIF().GetLine()
			}
		}()), localctx.(*If_prodContext).Get_RIF().GetColumn())
		localctx.(*If_prodContext).p = Natives.NewIf(localctx.(*If_prodContext).Get_expression().GetP(), localctx.(*If_prodContext).Get_bloq().GetContent(), nil, nil, (func() int {
			if localctx.(*If_prodContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*If_prodContext).Get_RIF().GetLine()
			}
		}()), localctx.(*If_prodContext).Get_RIF().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(398)

			var _m = p.Match(ProjectParserRIF)

			localctx.(*If_prodContext)._RIF = _m
		}
		{
			p.SetState(399)

			var _x = p.Expression()

			localctx.(*If_prodContext)._expression = _x
		}
		{
			p.SetState(400)

			var _x = p.Bloq()

			localctx.(*If_prodContext).bif = _x
		}
		{
			p.SetState(401)
			p.Match(ProjectParserRELSE)
		}
		{
			p.SetState(402)

			var _x = p.Bloq()

			localctx.(*If_prodContext).belse = _x
		}

		localctx.(*If_prodContext).instr = Natives.NewIf(localctx.(*If_prodContext).Get_expression().GetP(), localctx.(*If_prodContext).GetBif().GetContent(), nil, localctx.(*If_prodContext).GetBelse().GetContent(), (func() int {
			if localctx.(*If_prodContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*If_prodContext).Get_RIF().GetLine()
			}
		}()), localctx.(*If_prodContext).Get_RIF().GetColumn())
		localctx.(*If_prodContext).p = Natives.NewIf(localctx.(*If_prodContext).Get_expression().GetP(), localctx.(*If_prodContext).GetBif().GetContent(), nil, localctx.(*If_prodContext).GetBelse().GetContent(), (func() int {
			if localctx.(*If_prodContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*If_prodContext).Get_RIF().GetLine()
			}
		}()), localctx.(*If_prodContext).Get_RIF().GetColumn())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(405)

			var _m = p.Match(ProjectParserRIF)

			localctx.(*If_prodContext)._RIF = _m
		}
		{
			p.SetState(406)

			var _x = p.Expression()

			localctx.(*If_prodContext)._expression = _x
		}
		{
			p.SetState(407)

			var _x = p.Bloq()

			localctx.(*If_prodContext).bif = _x
		}
		{
			p.SetState(408)

			var _x = p.List_else_if()

			localctx.(*If_prodContext)._list_else_if = _x
		}
		{
			p.SetState(409)
			p.Match(ProjectParserRELSE)
		}
		{
			p.SetState(410)

			var _x = p.Bloq()

			localctx.(*If_prodContext).belse = _x
		}

		localctx.(*If_prodContext).instr = Natives.NewIf(localctx.(*If_prodContext).Get_expression().GetP(), localctx.(*If_prodContext).GetBif().GetContent(), localctx.(*If_prodContext).Get_list_else_if().GetList(), localctx.(*If_prodContext).GetBelse().GetContent(), (func() int {
			if localctx.(*If_prodContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*If_prodContext).Get_RIF().GetLine()
			}
		}()), localctx.(*If_prodContext).Get_RIF().GetColumn())
		localctx.(*If_prodContext).p = Natives.NewIf(localctx.(*If_prodContext).Get_expression().GetP(), localctx.(*If_prodContext).GetBif().GetContent(), localctx.(*If_prodContext).Get_list_else_if().GetList(), localctx.(*If_prodContext).GetBelse().GetContent(), (func() int {
			if localctx.(*If_prodContext).Get_RIF() == nil {
				return 0
			} else {
				return localctx.(*If_prodContext).Get_RIF().GetLine()
			}
		}()), localctx.(*If_prodContext).Get_RIF().GetColumn())

	}

	return localctx
}

// IList_else_ifContext is an interface to support dynamic dispatch.
type IList_else_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_else_if returns the _else_if rule contexts.
	Get_else_if() IElse_ifContext

	// Set_else_if sets the _else_if rule contexts.
	Set_else_if(IElse_ifContext)

	// GetNewList returns the newList rule context list.
	GetNewList() []IElse_ifContext

	// SetNewList sets the newList rule context list.
	SetNewList([]IElse_ifContext)

	// GetList returns the list attribute.
	GetList() *arrayList.List

	// SetList sets the list attribute.
	SetList(*arrayList.List)

	// IsList_else_ifContext differentiates from other interfaces.
	IsList_else_ifContext()
}

type List_else_ifContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	list     *arrayList.List
	_else_if IElse_ifContext
	newList  []IElse_ifContext
}

func NewEmptyList_else_ifContext() *List_else_ifContext {
	var p = new(List_else_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_list_else_if
	return p
}

func (*List_else_ifContext) IsList_else_ifContext() {}

func NewList_else_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_else_ifContext {
	var p = new(List_else_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_list_else_if

	return p
}

func (s *List_else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *List_else_ifContext) Get_else_if() IElse_ifContext { return s._else_if }

func (s *List_else_ifContext) Set_else_if(v IElse_ifContext) { s._else_if = v }

func (s *List_else_ifContext) GetNewList() []IElse_ifContext { return s.newList }

func (s *List_else_ifContext) SetNewList(v []IElse_ifContext) { s.newList = v }

func (s *List_else_ifContext) GetList() *arrayList.List { return s.list }

func (s *List_else_ifContext) SetList(v *arrayList.List) { s.list = v }

func (s *List_else_ifContext) AllElse_if() []IElse_ifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElse_ifContext)(nil)).Elem())
	var tst = make([]IElse_ifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElse_ifContext)
		}
	}

	return tst
}

func (s *List_else_ifContext) Else_if(i int) IElse_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_ifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElse_ifContext)
}

func (s *List_else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterList_else_if(s)
	}
}

func (s *List_else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitList_else_if(s)
	}
}

func (p *ProjectParser) List_else_if() (localctx IList_else_ifContext) {
	localctx = NewList_else_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ProjectParserRULE_list_else_if)

	localctx.(*List_else_ifContext).list = arrayList.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(415)

				var _x = p.Else_if()

				localctx.(*List_else_ifContext)._else_if = _x
			}
			localctx.(*List_else_ifContext).newList = append(localctx.(*List_else_ifContext).newList, localctx.(*List_else_ifContext)._else_if)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	listInt := localctx.(*List_else_ifContext).GetNewList()
	for _, e := range listInt {
		localctx.(*List_else_ifContext).list.Add(e.GetInstr())
	}

	return localctx
}

// IElse_ifContext is an interface to support dynamic dispatch.
type IElse_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RIF returns the _RIF token.
	Get_RIF() antlr.Token

	// Set_RIF sets the _RIF token.
	Set_RIF(antlr.Token)

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

	// IsElse_ifContext differentiates from other interfaces.
	IsElse_ifContext()
}

type Else_ifContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_RIF        antlr.Token
	_expression IExpressionContext
	_bloq       IBloqContext
}

func NewEmptyElse_ifContext() *Else_ifContext {
	var p = new(Else_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_else_if
	return p
}

func (*Else_ifContext) IsElse_ifContext() {}

func NewElse_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_ifContext {
	var p = new(Else_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_else_if

	return p
}

func (s *Else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_ifContext) Get_RIF() antlr.Token { return s._RIF }

func (s *Else_ifContext) Set_RIF(v antlr.Token) { s._RIF = v }

func (s *Else_ifContext) Get_expression() IExpressionContext { return s._expression }

func (s *Else_ifContext) Get_bloq() IBloqContext { return s._bloq }

func (s *Else_ifContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Else_ifContext) Set_bloq(v IBloqContext) { s._bloq = v }

func (s *Else_ifContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Else_ifContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Else_ifContext) RELSE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRELSE, 0)
}

func (s *Else_ifContext) RIF() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIF, 0)
}

func (s *Else_ifContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Else_ifContext) Bloq() IBloqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqContext)
}

func (s *Else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterElse_if(s)
	}
}

func (s *Else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitElse_if(s)
	}
}

func (p *ProjectParser) Else_if() (localctx IElse_ifContext) {
	localctx = NewElse_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ProjectParserRULE_else_if)

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
		p.SetState(422)
		p.Match(ProjectParserRELSE)
	}
	{
		p.SetState(423)

		var _m = p.Match(ProjectParserRIF)

		localctx.(*Else_ifContext)._RIF = _m
	}
	{
		p.SetState(424)

		var _x = p.Expression()

		localctx.(*Else_ifContext)._expression = _x
	}
	{
		p.SetState(425)

		var _x = p.Bloq()

		localctx.(*Else_ifContext)._bloq = _x
	}
	localctx.(*Else_ifContext).instr = Natives.NewIf(localctx.(*Else_ifContext).Get_expression().GetP(), localctx.(*Else_ifContext).Get_bloq().GetContent(), nil, nil, (func() int {
		if localctx.(*Else_ifContext).Get_RIF() == nil {
			return 0
		} else {
			return localctx.(*Else_ifContext).Get_RIF().GetLine()
		}
	}()), localctx.(*Else_ifContext).Get_RIF().GetColumn())

	return localctx
}

// IMatch_prodContext is an interface to support dynamic dispatch.
type IMatch_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_list_instr_match returns the _list_instr_match rule contexts.
	Get_list_instr_match() IList_instr_matchContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_list_instr_match sets the _list_instr_match rule contexts.
	Set_list_instr_match(IList_instr_matchContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsMatch_prodContext differentiates from other interfaces.
	IsMatch_prodContext()
}

type Match_prodContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	instr             Abstract.Instruction
	p                 Abstract.Expression
	_expression       IExpressionContext
	_list_instr_match IList_instr_matchContext
}

func NewEmptyMatch_prodContext() *Match_prodContext {
	var p = new(Match_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_match_prod
	return p
}

func (*Match_prodContext) IsMatch_prodContext() {}

func NewMatch_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Match_prodContext {
	var p = new(Match_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_match_prod

	return p
}

func (s *Match_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Match_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Match_prodContext) Get_list_instr_match() IList_instr_matchContext {
	return s._list_instr_match
}

func (s *Match_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Match_prodContext) Set_list_instr_match(v IList_instr_matchContext) { s._list_instr_match = v }

func (s *Match_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Match_prodContext) GetP() Abstract.Expression { return s.p }

func (s *Match_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Match_prodContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Match_prodContext) RMATCH() antlr.TerminalNode {
	return s.GetToken(ProjectParserRMATCH, 0)
}

func (s *Match_prodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Match_prodContext) LEFT_KEY() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_KEY, 0)
}

func (s *Match_prodContext) List_instr_match() IList_instr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_instr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_instr_matchContext)
}

func (s *Match_prodContext) RIGHT_KEY() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_KEY, 0)
}

func (s *Match_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Match_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Match_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterMatch_prod(s)
	}
}

func (s *Match_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitMatch_prod(s)
	}
}

func (p *ProjectParser) Match_prod() (localctx IMatch_prodContext) {
	localctx = NewMatch_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ProjectParserRULE_match_prod)

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
		p.SetState(428)
		p.Match(ProjectParserRMATCH)
	}
	{
		p.SetState(429)

		var _x = p.Expression()

		localctx.(*Match_prodContext)._expression = _x
	}
	{
		p.SetState(430)
		p.Match(ProjectParserLEFT_KEY)
	}
	{
		p.SetState(431)

		var _x = p.list_instr_match(0)

		localctx.(*Match_prodContext)._list_instr_match = _x
	}
	{
		p.SetState(432)
		p.Match(ProjectParserRIGHT_KEY)
	}

	localctx.(*Match_prodContext).instr = Natives.NewMatch(localctx.(*Match_prodContext).Get_expression().GetP(), localctx.(*Match_prodContext).Get_list_instr_match().GetList())
	localctx.(*Match_prodContext).p = Natives.NewMatch(localctx.(*Match_prodContext).Get_expression().GetP(), localctx.(*Match_prodContext).Get_list_instr_match().GetList())

	return localctx
}

// IList_instr_matchContext is an interface to support dynamic dispatch.
type IList_instr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSub returns the sub rule contexts.
	GetSub() IList_instr_matchContext

	// Get_instr_match returns the _instr_match rule contexts.
	Get_instr_match() IInstr_matchContext

	// SetSub sets the sub rule contexts.
	SetSub(IList_instr_matchContext)

	// Set_instr_match sets the _instr_match rule contexts.
	Set_instr_match(IInstr_matchContext)

	// GetList returns the list attribute.
	GetList() *arrayList.List

	// SetList sets the list attribute.
	SetList(*arrayList.List)

	// IsList_instr_matchContext differentiates from other interfaces.
	IsList_instr_matchContext()
}

type List_instr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	list         *arrayList.List
	sub          IList_instr_matchContext
	_instr_match IInstr_matchContext
}

func NewEmptyList_instr_matchContext() *List_instr_matchContext {
	var p = new(List_instr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_list_instr_match
	return p
}

func (*List_instr_matchContext) IsList_instr_matchContext() {}

func NewList_instr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_instr_matchContext {
	var p = new(List_instr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_list_instr_match

	return p
}

func (s *List_instr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *List_instr_matchContext) GetSub() IList_instr_matchContext { return s.sub }

func (s *List_instr_matchContext) Get_instr_match() IInstr_matchContext { return s._instr_match }

func (s *List_instr_matchContext) SetSub(v IList_instr_matchContext) { s.sub = v }

func (s *List_instr_matchContext) Set_instr_match(v IInstr_matchContext) { s._instr_match = v }

func (s *List_instr_matchContext) GetList() *arrayList.List { return s.list }

func (s *List_instr_matchContext) SetList(v *arrayList.List) { s.list = v }

func (s *List_instr_matchContext) Instr_match() IInstr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstr_matchContext)
}

func (s *List_instr_matchContext) List_instr_match() IList_instr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_instr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_instr_matchContext)
}

func (s *List_instr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_instr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_instr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterList_instr_match(s)
	}
}

func (s *List_instr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitList_instr_match(s)
	}
}

func (p *ProjectParser) List_instr_match() (localctx IList_instr_matchContext) {
	return p.list_instr_match(0)
}

func (p *ProjectParser) list_instr_match(_p int) (localctx IList_instr_matchContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewList_instr_matchContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IList_instr_matchContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, ProjectParserRULE_list_instr_match, _p)

	localctx.(*List_instr_matchContext).list = arrayList.New()

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
		p.SetState(436)

		var _x = p.Instr_match()

		localctx.(*List_instr_matchContext)._instr_match = _x
	}
	localctx.(*List_instr_matchContext).list.Add(localctx.(*List_instr_matchContext).Get_instr_match().GetInstr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewList_instr_matchContext(p, _parentctx, _parentState)
			localctx.(*List_instr_matchContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_list_instr_match)
			p.SetState(439)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(440)

				var _x = p.Instr_match()

				localctx.(*List_instr_matchContext)._instr_match = _x
			}

			localctx.(*List_instr_matchContext).GetSub().GetList().Add(localctx.(*List_instr_matchContext).Get_instr_match().GetInstr())
			localctx.(*List_instr_matchContext).list = localctx.(*List_instr_matchContext).GetSub().GetList()

		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}

	return localctx
}

// IInstr_matchContext is an interface to support dynamic dispatch.
type IInstr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_EQUAL_ARROW returns the _EQUAL_ARROW token.
	Get_EQUAL_ARROW() antlr.Token

	// Set_EQUAL_ARROW sets the _EQUAL_ARROW token.
	Set_EQUAL_ARROW(antlr.Token)

	// Get_expr_match returns the _expr_match rule contexts.
	Get_expr_match() IExpr_matchContext

	// Get_bloq_match returns the _bloq_match rule contexts.
	Get_bloq_match() IBloq_matchContext

	// Set_expr_match sets the _expr_match rule contexts.
	Set_expr_match(IExpr_matchContext)

	// Set_bloq_match sets the _bloq_match rule contexts.
	Set_bloq_match(IBloq_matchContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsInstr_matchContext differentiates from other interfaces.
	IsInstr_matchContext()
}

type Instr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	instr        Abstract.Instruction
	_expr_match  IExpr_matchContext
	_EQUAL_ARROW antlr.Token
	_bloq_match  IBloq_matchContext
}

func NewEmptyInstr_matchContext() *Instr_matchContext {
	var p = new(Instr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_instr_match
	return p
}

func (*Instr_matchContext) IsInstr_matchContext() {}

func NewInstr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instr_matchContext {
	var p = new(Instr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_instr_match

	return p
}

func (s *Instr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Instr_matchContext) Get_EQUAL_ARROW() antlr.Token { return s._EQUAL_ARROW }

func (s *Instr_matchContext) Set_EQUAL_ARROW(v antlr.Token) { s._EQUAL_ARROW = v }

func (s *Instr_matchContext) Get_expr_match() IExpr_matchContext { return s._expr_match }

func (s *Instr_matchContext) Get_bloq_match() IBloq_matchContext { return s._bloq_match }

func (s *Instr_matchContext) Set_expr_match(v IExpr_matchContext) { s._expr_match = v }

func (s *Instr_matchContext) Set_bloq_match(v IBloq_matchContext) { s._bloq_match = v }

func (s *Instr_matchContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Instr_matchContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Instr_matchContext) Expr_match() IExpr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_matchContext)
}

func (s *Instr_matchContext) EQUAL_ARROW() antlr.TerminalNode {
	return s.GetToken(ProjectParserEQUAL_ARROW, 0)
}

func (s *Instr_matchContext) Bloq_match() IBloq_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloq_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloq_matchContext)
}

func (s *Instr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterInstr_match(s)
	}
}

func (s *Instr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitInstr_match(s)
	}
}

func (p *ProjectParser) Instr_match() (localctx IInstr_matchContext) {
	localctx = NewInstr_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ProjectParserRULE_instr_match)

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
		p.SetState(448)

		var _x = p.expr_match(0)

		localctx.(*Instr_matchContext)._expr_match = _x
	}
	{
		p.SetState(449)

		var _m = p.Match(ProjectParserEQUAL_ARROW)

		localctx.(*Instr_matchContext)._EQUAL_ARROW = _m
	}
	{
		p.SetState(450)

		var _x = p.Bloq_match()

		localctx.(*Instr_matchContext)._bloq_match = _x
	}

	localctx.(*Instr_matchContext).instr = Natives.NewCase(localctx.(*Instr_matchContext).Get_expr_match().GetList(), localctx.(*Instr_matchContext).Get_bloq_match().GetContent(), (func() int {
		if localctx.(*Instr_matchContext).Get_EQUAL_ARROW() == nil {
			return 0
		} else {
			return localctx.(*Instr_matchContext).Get_EQUAL_ARROW().GetLine()
		}
	}()), localctx.(*Instr_matchContext).Get_EQUAL_ARROW().GetColumn())

	return localctx
}

// IExpr_matchContext is an interface to support dynamic dispatch.
type IExpr_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSub returns the sub rule contexts.
	GetSub() IExpr_matchContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetSub sets the sub rule contexts.
	SetSub(IExpr_matchContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetList returns the list attribute.
	GetList() *arrayList.List

	// SetList sets the list attribute.
	SetList(*arrayList.List)

	// IsExpr_matchContext differentiates from other interfaces.
	IsExpr_matchContext()
}

type Expr_matchContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	list        *arrayList.List
	sub         IExpr_matchContext
	_expression IExpressionContext
}

func NewEmptyExpr_matchContext() *Expr_matchContext {
	var p = new(Expr_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_expr_match
	return p
}

func (*Expr_matchContext) IsExpr_matchContext() {}

func NewExpr_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_matchContext {
	var p = new(Expr_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_expr_match

	return p
}

func (s *Expr_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_matchContext) GetSub() IExpr_matchContext { return s.sub }

func (s *Expr_matchContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_matchContext) SetSub(v IExpr_matchContext) { s.sub = v }

func (s *Expr_matchContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_matchContext) GetList() *arrayList.List { return s.list }

func (s *Expr_matchContext) SetList(v *arrayList.List) { s.list = v }

func (s *Expr_matchContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expr_matchContext) PIPE() antlr.TerminalNode {
	return s.GetToken(ProjectParserPIPE, 0)
}

func (s *Expr_matchContext) Expr_match() IExpr_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_matchContext)
}

func (s *Expr_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterExpr_match(s)
	}
}

func (s *Expr_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitExpr_match(s)
	}
}

func (p *ProjectParser) Expr_match() (localctx IExpr_matchContext) {
	return p.expr_match(0)
}

func (p *ProjectParser) expr_match(_p int) (localctx IExpr_matchContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_matchContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_matchContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, ProjectParserRULE_expr_match, _p)

	localctx.(*Expr_matchContext).list = arrayList.New()

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
		p.SetState(454)

		var _x = p.Expression()

		localctx.(*Expr_matchContext)._expression = _x
	}

	localctx.(*Expr_matchContext).list.Add(localctx.(*Expr_matchContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_matchContext(p, _parentctx, _parentState)
			localctx.(*Expr_matchContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_expr_match)
			p.SetState(457)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(458)
				p.Match(ProjectParserPIPE)
			}
			{
				p.SetState(459)

				var _x = p.Expression()

				localctx.(*Expr_matchContext)._expression = _x
			}

			localctx.(*Expr_matchContext).GetSub().GetList().Add(localctx.(*Expr_matchContext).Get_expression().GetP())
			localctx.(*Expr_matchContext).list = localctx.(*Expr_matchContext).GetSub().GetList()

		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IBucle_prodContext is an interface to support dynamic dispatch.
type IBucle_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_while_prod returns the _while_prod rule contexts.
	Get_while_prod() IWhile_prodContext

	// Get_loop_prod returns the _loop_prod rule contexts.
	Get_loop_prod() ILoop_prodContext

	// Get_forin_prod returns the _forin_prod rule contexts.
	Get_forin_prod() IForin_prodContext

	// Set_while_prod sets the _while_prod rule contexts.
	Set_while_prod(IWhile_prodContext)

	// Set_loop_prod sets the _loop_prod rule contexts.
	Set_loop_prod(ILoop_prodContext)

	// Set_forin_prod sets the _forin_prod rule contexts.
	Set_forin_prod(IForin_prodContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsBucle_prodContext differentiates from other interfaces.
	IsBucle_prodContext()
}

type Bucle_prodContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_while_prod IWhile_prodContext
	_loop_prod  ILoop_prodContext
	_forin_prod IForin_prodContext
}

func NewEmptyBucle_prodContext() *Bucle_prodContext {
	var p = new(Bucle_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_bucle_prod
	return p
}

func (*Bucle_prodContext) IsBucle_prodContext() {}

func NewBucle_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bucle_prodContext {
	var p = new(Bucle_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_bucle_prod

	return p
}

func (s *Bucle_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Bucle_prodContext) Get_while_prod() IWhile_prodContext { return s._while_prod }

func (s *Bucle_prodContext) Get_loop_prod() ILoop_prodContext { return s._loop_prod }

func (s *Bucle_prodContext) Get_forin_prod() IForin_prodContext { return s._forin_prod }

func (s *Bucle_prodContext) Set_while_prod(v IWhile_prodContext) { s._while_prod = v }

func (s *Bucle_prodContext) Set_loop_prod(v ILoop_prodContext) { s._loop_prod = v }

func (s *Bucle_prodContext) Set_forin_prod(v IForin_prodContext) { s._forin_prod = v }

func (s *Bucle_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Bucle_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Bucle_prodContext) While_prod() IWhile_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_prodContext)
}

func (s *Bucle_prodContext) Loop_prod() ILoop_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoop_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoop_prodContext)
}

func (s *Bucle_prodContext) Forin_prod() IForin_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForin_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForin_prodContext)
}

func (s *Bucle_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bucle_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bucle_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterBucle_prod(s)
	}
}

func (s *Bucle_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitBucle_prod(s)
	}
}

func (p *ProjectParser) Bucle_prod() (localctx IBucle_prodContext) {
	localctx = NewBucle_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ProjectParserRULE_bucle_prod)

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

	p.SetState(476)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserRWHILE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(467)

			var _x = p.While_prod()

			localctx.(*Bucle_prodContext)._while_prod = _x
		}
		localctx.(*Bucle_prodContext).instr = localctx.(*Bucle_prodContext).Get_while_prod().GetInstr()

	case ProjectParserRLOOP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(470)

			var _x = p.Loop_prod()

			localctx.(*Bucle_prodContext)._loop_prod = _x
		}
		localctx.(*Bucle_prodContext).instr = localctx.(*Bucle_prodContext).Get_loop_prod().GetInstr()

	case ProjectParserRFOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(473)

			var _x = p.Forin_prod()

			localctx.(*Bucle_prodContext)._forin_prod = _x
		}
		localctx.(*Bucle_prodContext).instr = localctx.(*Bucle_prodContext).Get_forin_prod().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWhile_prodContext is an interface to support dynamic dispatch.
type IWhile_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RWHILE returns the _RWHILE token.
	Get_RWHILE() antlr.Token

	// Set_RWHILE sets the _RWHILE token.
	Set_RWHILE(antlr.Token)

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

	// IsWhile_prodContext differentiates from other interfaces.
	IsWhile_prodContext()
}

type While_prodContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_RWHILE     antlr.Token
	_expression IExpressionContext
	_bloq       IBloqContext
}

func NewEmptyWhile_prodContext() *While_prodContext {
	var p = new(While_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_while_prod
	return p
}

func (*While_prodContext) IsWhile_prodContext() {}

func NewWhile_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_prodContext {
	var p = new(While_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_while_prod

	return p
}

func (s *While_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *While_prodContext) Get_RWHILE() antlr.Token { return s._RWHILE }

func (s *While_prodContext) Set_RWHILE(v antlr.Token) { s._RWHILE = v }

func (s *While_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *While_prodContext) Get_bloq() IBloqContext { return s._bloq }

func (s *While_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *While_prodContext) Set_bloq(v IBloqContext) { s._bloq = v }

func (s *While_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *While_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *While_prodContext) RWHILE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRWHILE, 0)
}

func (s *While_prodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *While_prodContext) Bloq() IBloqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqContext)
}

func (s *While_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterWhile_prod(s)
	}
}

func (s *While_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitWhile_prod(s)
	}
}

func (p *ProjectParser) While_prod() (localctx IWhile_prodContext) {
	localctx = NewWhile_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ProjectParserRULE_while_prod)

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
		p.SetState(478)

		var _m = p.Match(ProjectParserRWHILE)

		localctx.(*While_prodContext)._RWHILE = _m
	}
	{
		p.SetState(479)

		var _x = p.Expression()

		localctx.(*While_prodContext)._expression = _x
	}
	{
		p.SetState(480)

		var _x = p.Bloq()

		localctx.(*While_prodContext)._bloq = _x
	}
	localctx.(*While_prodContext).instr = Natives.NewWhile(localctx.(*While_prodContext).Get_expression().GetP(), localctx.(*While_prodContext).Get_bloq().GetContent(), (func() int {
		if localctx.(*While_prodContext).Get_RWHILE() == nil {
			return 0
		} else {
			return localctx.(*While_prodContext).Get_RWHILE().GetLine()
		}
	}()), localctx.(*While_prodContext).Get_RWHILE().GetColumn())

	return localctx
}

// ILoop_prodContext is an interface to support dynamic dispatch.
type ILoop_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_bloq returns the _bloq rule contexts.
	Get_bloq() IBloqContext

	// Set_bloq sets the _bloq rule contexts.
	Set_bloq(IBloqContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsLoop_prodContext differentiates from other interfaces.
	IsLoop_prodContext()
}

type Loop_prodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	instr  Abstract.Instruction
	p      Abstract.Expression
	_bloq  IBloqContext
}

func NewEmptyLoop_prodContext() *Loop_prodContext {
	var p = new(Loop_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_loop_prod
	return p
}

func (*Loop_prodContext) IsLoop_prodContext() {}

func NewLoop_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loop_prodContext {
	var p = new(Loop_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_loop_prod

	return p
}

func (s *Loop_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Loop_prodContext) Get_bloq() IBloqContext { return s._bloq }

func (s *Loop_prodContext) Set_bloq(v IBloqContext) { s._bloq = v }

func (s *Loop_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Loop_prodContext) GetP() Abstract.Expression { return s.p }

func (s *Loop_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Loop_prodContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Loop_prodContext) RLOOP() antlr.TerminalNode {
	return s.GetToken(ProjectParserRLOOP, 0)
}

func (s *Loop_prodContext) Bloq() IBloqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqContext)
}

func (s *Loop_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loop_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loop_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterLoop_prod(s)
	}
}

func (s *Loop_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitLoop_prod(s)
	}
}

func (p *ProjectParser) Loop_prod() (localctx ILoop_prodContext) {
	localctx = NewLoop_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ProjectParserRULE_loop_prod)

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
		p.SetState(483)
		p.Match(ProjectParserRLOOP)
	}
	{
		p.SetState(484)

		var _x = p.Bloq()

		localctx.(*Loop_prodContext)._bloq = _x
	}

	localctx.(*Loop_prodContext).instr = Natives.NewLoop(localctx.(*Loop_prodContext).Get_bloq().GetContent())
	localctx.(*Loop_prodContext).p = Natives.NewLoop(localctx.(*Loop_prodContext).Get_bloq().GetContent())

	return localctx
}

// IForin_prodContext is an interface to support dynamic dispatch.
type IForin_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_range_prod returns the _range_prod rule contexts.
	Get_range_prod() IRange_prodContext

	// Get_bloq returns the _bloq rule contexts.
	Get_bloq() IBloqContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_range_prod sets the _range_prod rule contexts.
	Set_range_prod(IRange_prodContext)

	// Set_bloq sets the _bloq rule contexts.
	Set_bloq(IBloqContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsForin_prodContext differentiates from other interfaces.
	IsForin_prodContext()
}

type Forin_prodContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_expression IExpressionContext
	_range_prod IRange_prodContext
	_bloq       IBloqContext
}

func NewEmptyForin_prodContext() *Forin_prodContext {
	var p = new(Forin_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_forin_prod
	return p
}

func (*Forin_prodContext) IsForin_prodContext() {}

func NewForin_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Forin_prodContext {
	var p = new(Forin_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_forin_prod

	return p
}

func (s *Forin_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Forin_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Forin_prodContext) Get_range_prod() IRange_prodContext { return s._range_prod }

func (s *Forin_prodContext) Get_bloq() IBloqContext { return s._bloq }

func (s *Forin_prodContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Forin_prodContext) Set_range_prod(v IRange_prodContext) { s._range_prod = v }

func (s *Forin_prodContext) Set_bloq(v IBloqContext) { s._bloq = v }

func (s *Forin_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Forin_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Forin_prodContext) RFOR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRFOR, 0)
}

func (s *Forin_prodContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Forin_prodContext) RIN() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIN, 0)
}

func (s *Forin_prodContext) Range_prod() IRange_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRange_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRange_prodContext)
}

func (s *Forin_prodContext) Bloq() IBloqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqContext)
}

func (s *Forin_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Forin_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Forin_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterForin_prod(s)
	}
}

func (s *Forin_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitForin_prod(s)
	}
}

func (p *ProjectParser) Forin_prod() (localctx IForin_prodContext) {
	localctx = NewForin_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ProjectParserRULE_forin_prod)

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
		p.SetState(487)
		p.Match(ProjectParserRFOR)
	}
	{
		p.SetState(488)

		var _x = p.Expression()

		localctx.(*Forin_prodContext)._expression = _x
	}
	{
		p.SetState(489)
		p.Match(ProjectParserRIN)
	}
	{
		p.SetState(490)

		var _x = p.Range_prod()

		localctx.(*Forin_prodContext)._range_prod = _x
	}
	{
		p.SetState(491)

		var _x = p.Bloq()

		localctx.(*Forin_prodContext)._bloq = _x
	}
	localctx.(*Forin_prodContext).instr = Natives.NewForIn(localctx.(*Forin_prodContext).Get_expression().GetP(), localctx.(*Forin_prodContext).Get_range_prod().GetP(), localctx.(*Forin_prodContext).Get_bloq().GetContent())

	return localctx
}

// IRange_prodContext is an interface to support dynamic dispatch.
type IRange_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RANGE returns the _RANGE token.
	Get_RANGE() antlr.Token

	// Set_RANGE sets the _RANGE token.
	Set_RANGE(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpressionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpressionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpressionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpressionContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsRange_prodContext differentiates from other interfaces.
	IsRange_prodContext()
}

type Range_prodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	p      Abstract.Expression
	e1     IExpressionContext
	_RANGE antlr.Token
	e2     IExpressionContext
}

func NewEmptyRange_prodContext() *Range_prodContext {
	var p = new(Range_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_range_prod
	return p
}

func (*Range_prodContext) IsRange_prodContext() {}

func NewRange_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_prodContext {
	var p = new(Range_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_range_prod

	return p
}

func (s *Range_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Range_prodContext) Get_RANGE() antlr.Token { return s._RANGE }

func (s *Range_prodContext) Set_RANGE(v antlr.Token) { s._RANGE = v }

func (s *Range_prodContext) GetE1() IExpressionContext { return s.e1 }

func (s *Range_prodContext) GetE2() IExpressionContext { return s.e2 }

func (s *Range_prodContext) SetE1(v IExpressionContext) { s.e1 = v }

func (s *Range_prodContext) SetE2(v IExpressionContext) { s.e2 = v }

func (s *Range_prodContext) GetP() Abstract.Expression { return s.p }

func (s *Range_prodContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Range_prodContext) RANGE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRANGE, 0)
}

func (s *Range_prodContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Range_prodContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Range_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterRange_prod(s)
	}
}

func (s *Range_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitRange_prod(s)
	}
}

func (p *ProjectParser) Range_prod() (localctx IRange_prodContext) {
	localctx = NewRange_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ProjectParserRULE_range_prod)

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
		p.SetState(494)

		var _x = p.Expression()

		localctx.(*Range_prodContext).e1 = _x
	}
	{
		p.SetState(495)

		var _m = p.Match(ProjectParserRANGE)

		localctx.(*Range_prodContext)._RANGE = _m
	}
	{
		p.SetState(496)

		var _x = p.Expression()

		localctx.(*Range_prodContext).e2 = _x
	}

	localctx.(*Range_prodContext).p = Expression.NewRange((func() int {
		if localctx.(*Range_prodContext).Get_RANGE() == nil {
			return 0
		} else {
			return localctx.(*Range_prodContext).Get_RANGE().GetLine()
		}
	}()), localctx.(*Range_prodContext).Get_RANGE().GetColumn(), localctx.(*Range_prodContext).GetE1().GetP(), localctx.(*Range_prodContext).GetE2().GetP())

	return localctx
}

// ICalled_funcContext is an interface to support dynamic dispatch.
type ICalled_funcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listExpressions returns the _listExpressions rule contexts.
	Get_listExpressions() IListExpressionsContext

	// Set_listExpressions sets the _listExpressions rule contexts.
	Set_listExpressions(IListExpressionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsCalled_funcContext differentiates from other interfaces.
	IsCalled_funcContext()
}

type Called_funcContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            Abstract.Instruction
	p                Abstract.Expression
	_ID              antlr.Token
	_listExpressions IListExpressionsContext
}

func NewEmptyCalled_funcContext() *Called_funcContext {
	var p = new(Called_funcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_called_func
	return p
}

func (*Called_funcContext) IsCalled_funcContext() {}

func NewCalled_funcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Called_funcContext {
	var p = new(Called_funcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_called_func

	return p
}

func (s *Called_funcContext) GetParser() antlr.Parser { return s.parser }

func (s *Called_funcContext) Get_ID() antlr.Token { return s._ID }

func (s *Called_funcContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Called_funcContext) Get_listExpressions() IListExpressionsContext { return s._listExpressions }

func (s *Called_funcContext) Set_listExpressions(v IListExpressionsContext) { s._listExpressions = v }

func (s *Called_funcContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Called_funcContext) GetP() Abstract.Expression { return s.p }

func (s *Called_funcContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Called_funcContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Called_funcContext) ID() antlr.TerminalNode {
	return s.GetToken(ProjectParserID, 0)
}

func (s *Called_funcContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *Called_funcContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_PAR, 0)
}

func (s *Called_funcContext) ListExpressions() IListExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListExpressionsContext)
}

func (s *Called_funcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Called_funcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Called_funcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterCalled_func(s)
	}
}

func (s *Called_funcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitCalled_func(s)
	}
}

func (p *ProjectParser) Called_func() (localctx ICalled_funcContext) {
	localctx = NewCalled_funcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ProjectParserRULE_called_func)

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

	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)

			var _m = p.Match(ProjectParserID)

			localctx.(*Called_funcContext)._ID = _m
		}
		{
			p.SetState(500)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(501)
			p.Match(ProjectParserRIGHT_PAR)
		}

		localctx.(*Called_funcContext).instr = ExpressionSpecial.NewCallFunction((func() string {
			if localctx.(*Called_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Called_funcContext).Get_ID().GetText()
			}
		}()), arrayList.New())
		localctx.(*Called_funcContext).p = ExpressionSpecial.NewCallFunction((func() string {
			if localctx.(*Called_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Called_funcContext).Get_ID().GetText()
			}
		}()), arrayList.New())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(503)

			var _m = p.Match(ProjectParserID)

			localctx.(*Called_funcContext)._ID = _m
		}
		{
			p.SetState(504)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(505)

			var _x = p.listExpressions(0)

			localctx.(*Called_funcContext)._listExpressions = _x
		}
		{
			p.SetState(506)
			p.Match(ProjectParserRIGHT_PAR)
		}

		localctx.(*Called_funcContext).instr = ExpressionSpecial.NewCallFunction((func() string {
			if localctx.(*Called_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Called_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Called_funcContext).Get_listExpressions().GetL())
		localctx.(*Called_funcContext).p = ExpressionSpecial.NewCallFunction((func() string {
			if localctx.(*Called_funcContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Called_funcContext).Get_ID().GetText()
			}
		}()), localctx.(*Called_funcContext).Get_listExpressions().GetL())

	}

	return localctx
}

// IListExpressionsContext is an interface to support dynamic dispatch.
type IListExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the List rule contexts.
	GetList() IListExpressionsContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the List rule contexts.
	SetList(IListExpressionsContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListExpressionsContext differentiates from other interfaces.
	IsListExpressionsContext()
}

type ListExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	List        IListExpressionsContext
	_expression IExpressionContext
}

func NewEmptyListExpressionsContext() *ListExpressionsContext {
	var p = new(ListExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_listExpressions
	return p
}

func (*ListExpressionsContext) IsListExpressionsContext() {}

func NewListExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListExpressionsContext {
	var p = new(ListExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_listExpressions

	return p
}

func (s *ListExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListExpressionsContext) GetList() IListExpressionsContext { return s.List }

func (s *ListExpressionsContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListExpressionsContext) SetList(v IListExpressionsContext) { s.List = v }

func (s *ListExpressionsContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListExpressionsContext) GetL() *arrayList.List { return s.l }

func (s *ListExpressionsContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListExpressionsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListExpressionsContext) ListExpressions() IListExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListExpressionsContext)
}

func (s *ListExpressionsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *ListExpressionsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *ListExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterListExpressions(s)
	}
}

func (s *ListExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitListExpressions(s)
	}
}

func (p *ProjectParser) ListExpressions() (localctx IListExpressionsContext) {
	return p.listExpressions(0)
}

func (p *ProjectParser) listExpressions(_p int) (localctx IListExpressionsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListExpressionsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListExpressionsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, ProjectParserRULE_listExpressions, _p)

	localctx.(*ListExpressionsContext).l = arrayList.New()

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
		p.SetState(512)

		var _x = p.Expression()

		localctx.(*ListExpressionsContext)._expression = _x
	}
	localctx.(*ListExpressionsContext).l.Add(localctx.(*ListExpressionsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListExpressionsContext(p, _parentctx, _parentState)
			localctx.(*ListExpressionsContext).List = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listExpressions)
			p.SetState(515)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(516)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(517)

				var _x = p.Expression()

				localctx.(*ListExpressionsContext)._expression = _x
			}

			localctx.(*ListExpressionsContext).GetList().GetL().Add(localctx.(*ListExpressionsContext).Get_expression().GetP())
			localctx.(*ListExpressionsContext).l = localctx.(*ListExpressionsContext).GetList().GetL()

		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// IDec_arrContext is an interface to support dynamic dispatch.
type IDec_arrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MUT returns the _MUT token.
	Get_MUT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_MUT sets the _MUT token.
	Set_MUT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listDim returns the _listDim rule contexts.
	Get_listDim() IListDimContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_listDim sets the _listDim rule contexts.
	Set_listDim(IListDimContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsDec_arrContext differentiates from other interfaces.
	IsDec_arrContext()
}

type Dec_arrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_MUT        antlr.Token
	_ID         antlr.Token
	_listDim    IListDimContext
	_expression IExpressionContext
}

func NewEmptyDec_arrContext() *Dec_arrContext {
	var p = new(Dec_arrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_dec_arr
	return p
}

func (*Dec_arrContext) IsDec_arrContext() {}

func NewDec_arrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_arrContext {
	var p = new(Dec_arrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_dec_arr

	return p
}

func (s *Dec_arrContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_arrContext) Get_MUT() antlr.Token { return s._MUT }

func (s *Dec_arrContext) Get_ID() antlr.Token { return s._ID }

func (s *Dec_arrContext) Set_MUT(v antlr.Token) { s._MUT = v }

func (s *Dec_arrContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Dec_arrContext) Get_listDim() IListDimContext { return s._listDim }

func (s *Dec_arrContext) Get_expression() IExpressionContext { return s._expression }

func (s *Dec_arrContext) Set_listDim(v IListDimContext) { s._listDim = v }

func (s *Dec_arrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Dec_arrContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Dec_arrContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Dec_arrContext) DECLARAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserDECLARAR, 0)
}

func (s *Dec_arrContext) ID() antlr.TerminalNode {
	return s.GetToken(ProjectParserID, 0)
}

func (s *Dec_arrContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ProjectParserEQUAL, 0)
}

func (s *Dec_arrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Dec_arrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *Dec_arrContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
}

func (s *Dec_arrContext) MUT() antlr.TerminalNode {
	return s.GetToken(ProjectParserMUT, 0)
}

func (s *Dec_arrContext) COLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOLON, 0)
}

func (s *Dec_arrContext) ListDim() IListDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListDimContext)
}

func (s *Dec_arrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_arrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dec_arrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterDec_arr(s)
	}
}

func (s *Dec_arrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitDec_arr(s)
	}
}

func (p *ProjectParser) Dec_arr() (localctx IDec_arrContext) {
	localctx = NewDec_arrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ProjectParserRULE_dec_arr)
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
		p.SetState(525)
		p.Match(ProjectParserDECLARAR)
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProjectParserMUT {
		{
			p.SetState(526)

			var _m = p.Match(ProjectParserMUT)

			localctx.(*Dec_arrContext)._MUT = _m
		}

	}
	{
		p.SetState(529)

		var _m = p.Match(ProjectParserID)

		localctx.(*Dec_arrContext)._ID = _m
	}
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ProjectParserCOLON {
		{
			p.SetState(530)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(531)

			var _x = p.ListDim()

			localctx.(*Dec_arrContext)._listDim = _x
		}

	}
	{
		p.SetState(534)
		p.Match(ProjectParserEQUAL)
	}
	{
		p.SetState(535)

		var _x = p.Expression()

		localctx.(*Dec_arrContext)._expression = _x
	}
	{
		p.SetState(536)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ProjectParserSEMICOLON || _la == ProjectParserCOMMA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	if (func() string {
		if localctx.(*Dec_arrContext).Get_MUT() == nil {
			return ""
		} else {
			return localctx.(*Dec_arrContext).Get_MUT().GetText()
		}
	}()) != "" {
		localctx.(*Dec_arrContext).instr = DecArrays.NewDecArray(localctx.(*Dec_arrContext).Get_listDim().GetLength(), (func() string {
			if localctx.(*Dec_arrContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Dec_arrContext).Get_ID().GetText()
			}
		}()), localctx.(*Dec_arrContext).Get_expression().GetP(), localctx.(*Dec_arrContext).Get_listDim().GetData(), true, localctx.(*Dec_arrContext).Get_listDim().GetPos())
	} else {
		localctx.(*Dec_arrContext).instr = DecArrays.NewDecArray(localctx.(*Dec_arrContext).Get_listDim().GetLength(), (func() string {
			if localctx.(*Dec_arrContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Dec_arrContext).Get_ID().GetText()
			}
		}()), localctx.(*Dec_arrContext).Get_expression().GetP(), localctx.(*Dec_arrContext).Get_listDim().GetData(), false, localctx.(*Dec_arrContext).Get_listDim().GetPos())
	}

	return localctx
}

// IListDimContext is an interface to support dynamic dispatch.
type IListDimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l rule contexts.
	GetL() IListDimContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_data_type returns the _data_type rule contexts.
	Get_data_type() IData_typeContext

	// SetL sets the l rule contexts.
	SetL(IListDimContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_data_type sets the _data_type rule contexts.
	Set_data_type(IData_typeContext)

	// GetLength returns the length attribute.
	GetLength() int

	// GetData returns the data attribute.
	GetData() SymbolTable.DataType

	// GetPos returns the pos attribute.
	GetPos() Abstract.Expression

	// SetLength sets the length attribute.
	SetLength(int)

	// SetData sets the data attribute.
	SetData(SymbolTable.DataType)

	// SetPos sets the pos attribute.
	SetPos(Abstract.Expression)

	// IsListDimContext differentiates from other interfaces.
	IsListDimContext()
}

type ListDimContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	length      int
	data        SymbolTable.DataType
	pos         Abstract.Expression
	l           IListDimContext
	_expression IExpressionContext
	_data_type  IData_typeContext
}

func NewEmptyListDimContext() *ListDimContext {
	var p = new(ListDimContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_listDim
	return p
}

func (*ListDimContext) IsListDimContext() {}

func NewListDimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListDimContext {
	var p = new(ListDimContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_listDim

	return p
}

func (s *ListDimContext) GetParser() antlr.Parser { return s.parser }

func (s *ListDimContext) GetL() IListDimContext { return s.l }

func (s *ListDimContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListDimContext) Get_data_type() IData_typeContext { return s._data_type }

func (s *ListDimContext) SetL(v IListDimContext) { s.l = v }

func (s *ListDimContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListDimContext) Set_data_type(v IData_typeContext) { s._data_type = v }

func (s *ListDimContext) GetLength() int { return s.length }

func (s *ListDimContext) GetData() SymbolTable.DataType { return s.data }

func (s *ListDimContext) GetPos() Abstract.Expression { return s.pos }

func (s *ListDimContext) SetLength(v int) { s.length = v }

func (s *ListDimContext) SetData(v SymbolTable.DataType) { s.data = v }

func (s *ListDimContext) SetPos(v Abstract.Expression) { s.pos = v }

func (s *ListDimContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_BRACKET, 0)
}

func (s *ListDimContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *ListDimContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListDimContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_BRACKET, 0)
}

func (s *ListDimContext) ListDim() IListDimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListDimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListDimContext)
}

func (s *ListDimContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *ListDimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListDimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListDimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterListDim(s)
	}
}

func (s *ListDimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitListDim(s)
	}
}

func (p *ProjectParser) ListDim() (localctx IListDimContext) {
	localctx = NewListDimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ProjectParserRULE_listDim)
	localctx.(*ListDimContext).length = 0

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

	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(539)
			p.Match(ProjectParserLEFT_BRACKET)
		}
		{
			p.SetState(540)

			var _x = p.ListDim()

			localctx.(*ListDimContext).l = _x
		}
		{
			p.SetState(541)
			p.Match(ProjectParserSEMICOLON)
		}
		{
			p.SetState(542)

			var _x = p.Expression()

			localctx.(*ListDimContext)._expression = _x
		}
		{
			p.SetState(543)
			p.Match(ProjectParserRIGHT_BRACKET)
		}
		localctx.(*ListDimContext).length = localctx.(*ListDimContext).GetL().GetLength() + 1
		localctx.(*ListDimContext).data = localctx.(*ListDimContext).GetL().GetData()
		localctx.(*ListDimContext).pos = localctx.(*ListDimContext).Get_expression().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(546)
			p.Match(ProjectParserLEFT_BRACKET)
		}
		{
			p.SetState(547)

			var _x = p.Data_type()

			localctx.(*ListDimContext)._data_type = _x
		}
		{
			p.SetState(548)
			p.Match(ProjectParserSEMICOLON)
		}
		{
			p.SetState(549)

			var _x = p.Expression()

			localctx.(*ListDimContext)._expression = _x
		}
		{
			p.SetState(550)
			p.Match(ProjectParserRIGHT_BRACKET)
		}

		localctx.(*ListDimContext).length = 1
		localctx.(*ListDimContext).pos = localctx.(*ListDimContext).Get_expression().GetP()
		switch localctx.(*ListDimContext).Get_data_type().GetData() {
		case "i64":
			localctx.(*ListDimContext).data = SymbolTable.INTEGER
		case "f64":
			localctx.(*ListDimContext).data = SymbolTable.FLOAT
		case "&str":
			localctx.(*ListDimContext).data = SymbolTable.STR
		case "String":
			localctx.(*ListDimContext).data = SymbolTable.STRING
		case "bool":
			localctx.(*ListDimContext).data = SymbolTable.BOOLEAN
		}

	}

	return localctx
}

// ITransfer_prodContext is an interface to support dynamic dispatch.
type ITransfer_prodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_break_instr returns the _break_instr rule contexts.
	Get_break_instr() IBreak_instrContext

	// Get_continue_instr returns the _continue_instr rule contexts.
	Get_continue_instr() IContinue_instrContext

	// Get_return_instr returns the _return_instr rule contexts.
	Get_return_instr() IReturn_instrContext

	// Set_break_instr sets the _break_instr rule contexts.
	Set_break_instr(IBreak_instrContext)

	// Set_continue_instr sets the _continue_instr rule contexts.
	Set_continue_instr(IContinue_instrContext)

	// Set_return_instr sets the _return_instr rule contexts.
	Set_return_instr(IReturn_instrContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsTransfer_prodContext differentiates from other interfaces.
	IsTransfer_prodContext()
}

type Transfer_prodContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	instr           Abstract.Instruction
	_break_instr    IBreak_instrContext
	_continue_instr IContinue_instrContext
	_return_instr   IReturn_instrContext
}

func NewEmptyTransfer_prodContext() *Transfer_prodContext {
	var p = new(Transfer_prodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_transfer_prod
	return p
}

func (*Transfer_prodContext) IsTransfer_prodContext() {}

func NewTransfer_prodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Transfer_prodContext {
	var p = new(Transfer_prodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_transfer_prod

	return p
}

func (s *Transfer_prodContext) GetParser() antlr.Parser { return s.parser }

func (s *Transfer_prodContext) Get_break_instr() IBreak_instrContext { return s._break_instr }

func (s *Transfer_prodContext) Get_continue_instr() IContinue_instrContext { return s._continue_instr }

func (s *Transfer_prodContext) Get_return_instr() IReturn_instrContext { return s._return_instr }

func (s *Transfer_prodContext) Set_break_instr(v IBreak_instrContext) { s._break_instr = v }

func (s *Transfer_prodContext) Set_continue_instr(v IContinue_instrContext) { s._continue_instr = v }

func (s *Transfer_prodContext) Set_return_instr(v IReturn_instrContext) { s._return_instr = v }

func (s *Transfer_prodContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Transfer_prodContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Transfer_prodContext) Break_instr() IBreak_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreak_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreak_instrContext)
}

func (s *Transfer_prodContext) Continue_instr() IContinue_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinue_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinue_instrContext)
}

func (s *Transfer_prodContext) Return_instr() IReturn_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_instrContext)
}

func (s *Transfer_prodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Transfer_prodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Transfer_prodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterTransfer_prod(s)
	}
}

func (s *Transfer_prodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitTransfer_prod(s)
	}
}

func (p *ProjectParser) Transfer_prod() (localctx ITransfer_prodContext) {
	localctx = NewTransfer_prodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ProjectParserRULE_transfer_prod)

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

	p.SetState(564)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserRBREAK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(555)

			var _x = p.Break_instr()

			localctx.(*Transfer_prodContext)._break_instr = _x
		}
		localctx.(*Transfer_prodContext).instr = localctx.(*Transfer_prodContext).Get_break_instr().GetInstr()

	case ProjectParserRCONTINUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(558)

			var _x = p.Continue_instr()

			localctx.(*Transfer_prodContext)._continue_instr = _x
		}
		localctx.(*Transfer_prodContext).instr = localctx.(*Transfer_prodContext).Get_continue_instr().GetInstr()

	case ProjectParserRRETURN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(561)

			var _x = p.Return_instr()

			localctx.(*Transfer_prodContext)._return_instr = _x
		}
		localctx.(*Transfer_prodContext).instr = localctx.(*Transfer_prodContext).Get_return_instr().GetInstr()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBreak_instrContext is an interface to support dynamic dispatch.
type IBreak_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RBREAK returns the _RBREAK token.
	Get_RBREAK() antlr.Token

	// Set_RBREAK sets the _RBREAK token.
	Set_RBREAK(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsBreak_instrContext differentiates from other interfaces.
	IsBreak_instrContext()
}

type Break_instrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_RBREAK     antlr.Token
	_expression IExpressionContext
}

func NewEmptyBreak_instrContext() *Break_instrContext {
	var p = new(Break_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_break_instr
	return p
}

func (*Break_instrContext) IsBreak_instrContext() {}

func NewBreak_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_instrContext {
	var p = new(Break_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_break_instr

	return p
}

func (s *Break_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *Break_instrContext) Get_RBREAK() antlr.Token { return s._RBREAK }

func (s *Break_instrContext) Set_RBREAK(v antlr.Token) { s._RBREAK = v }

func (s *Break_instrContext) Get_expression() IExpressionContext { return s._expression }

func (s *Break_instrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Break_instrContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Break_instrContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Break_instrContext) RBREAK() antlr.TerminalNode {
	return s.GetToken(ProjectParserRBREAK, 0)
}

func (s *Break_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *Break_instrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Break_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterBreak_instr(s)
	}
}

func (s *Break_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitBreak_instr(s)
	}
}

func (p *ProjectParser) Break_instr() (localctx IBreak_instrContext) {
	localctx = NewBreak_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ProjectParserRULE_break_instr)

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

	p.SetState(574)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(566)

			var _m = p.Match(ProjectParserRBREAK)

			localctx.(*Break_instrContext)._RBREAK = _m
		}
		{
			p.SetState(567)
			p.Match(ProjectParserSEMICOLON)
		}
		localctx.(*Break_instrContext).instr = Natives.NewBreak((func() int {
			if localctx.(*Break_instrContext).Get_RBREAK() == nil {
				return 0
			} else {
				return localctx.(*Break_instrContext).Get_RBREAK().GetLine()
			}
		}()), localctx.(*Break_instrContext).Get_RBREAK().GetColumn(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(569)

			var _m = p.Match(ProjectParserRBREAK)

			localctx.(*Break_instrContext)._RBREAK = _m
		}
		{
			p.SetState(570)

			var _x = p.Expression()

			localctx.(*Break_instrContext)._expression = _x
		}
		{
			p.SetState(571)
			p.Match(ProjectParserSEMICOLON)
		}
		localctx.(*Break_instrContext).instr = Natives.NewBreak((func() int {
			if localctx.(*Break_instrContext).Get_RBREAK() == nil {
				return 0
			} else {
				return localctx.(*Break_instrContext).Get_RBREAK().GetLine()
			}
		}()), localctx.(*Break_instrContext).Get_RBREAK().GetColumn(), localctx.(*Break_instrContext).Get_expression().GetP())

	}

	return localctx
}

// IContinue_instrContext is an interface to support dynamic dispatch.
type IContinue_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RCONTINUE returns the _RCONTINUE token.
	Get_RCONTINUE() antlr.Token

	// Set_RCONTINUE sets the _RCONTINUE token.
	Set_RCONTINUE(antlr.Token)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsContinue_instrContext differentiates from other interfaces.
	IsContinue_instrContext()
}

type Continue_instrContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	instr      Abstract.Instruction
	_RCONTINUE antlr.Token
}

func NewEmptyContinue_instrContext() *Continue_instrContext {
	var p = new(Continue_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_continue_instr
	return p
}

func (*Continue_instrContext) IsContinue_instrContext() {}

func NewContinue_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_instrContext {
	var p = new(Continue_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_continue_instr

	return p
}

func (s *Continue_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *Continue_instrContext) Get_RCONTINUE() antlr.Token { return s._RCONTINUE }

func (s *Continue_instrContext) Set_RCONTINUE(v antlr.Token) { s._RCONTINUE = v }

func (s *Continue_instrContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Continue_instrContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Continue_instrContext) RCONTINUE() antlr.TerminalNode {
	return s.GetToken(ProjectParserRCONTINUE, 0)
}

func (s *Continue_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *Continue_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterContinue_instr(s)
	}
}

func (s *Continue_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitContinue_instr(s)
	}
}

func (p *ProjectParser) Continue_instr() (localctx IContinue_instrContext) {
	localctx = NewContinue_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ProjectParserRULE_continue_instr)

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
		p.SetState(576)

		var _m = p.Match(ProjectParserRCONTINUE)

		localctx.(*Continue_instrContext)._RCONTINUE = _m
	}
	{
		p.SetState(577)
		p.Match(ProjectParserSEMICOLON)
	}
	localctx.(*Continue_instrContext).instr = Natives.NewContinue((func() int {
		if localctx.(*Continue_instrContext).Get_RCONTINUE() == nil {
			return 0
		} else {
			return localctx.(*Continue_instrContext).Get_RCONTINUE().GetLine()
		}
	}()), localctx.(*Continue_instrContext).Get_RCONTINUE().GetColumn())

	return localctx
}

// IReturn_instrContext is an interface to support dynamic dispatch.
type IReturn_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RRETURN returns the _RRETURN token.
	Get_RRETURN() antlr.Token

	// Set_RRETURN sets the _RRETURN token.
	Set_RRETURN(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsReturn_instrContext differentiates from other interfaces.
	IsReturn_instrContext()
}

type Return_instrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       Abstract.Instruction
	_RRETURN    antlr.Token
	_expression IExpressionContext
}

func NewEmptyReturn_instrContext() *Return_instrContext {
	var p = new(Return_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_return_instr
	return p
}

func (*Return_instrContext) IsReturn_instrContext() {}

func NewReturn_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_instrContext {
	var p = new(Return_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_return_instr

	return p
}

func (s *Return_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_instrContext) Get_RRETURN() antlr.Token { return s._RRETURN }

func (s *Return_instrContext) Set_RRETURN(v antlr.Token) { s._RRETURN = v }

func (s *Return_instrContext) Get_expression() IExpressionContext { return s._expression }

func (s *Return_instrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Return_instrContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Return_instrContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Return_instrContext) RRETURN() antlr.TerminalNode {
	return s.GetToken(ProjectParserRRETURN, 0)
}

func (s *Return_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(ProjectParserSEMICOLON, 0)
}

func (s *Return_instrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterReturn_instr(s)
	}
}

func (s *Return_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitReturn_instr(s)
	}
}

func (p *ProjectParser) Return_instr() (localctx IReturn_instrContext) {
	localctx = NewReturn_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ProjectParserRULE_return_instr)

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

	p.SetState(590)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(580)

			var _m = p.Match(ProjectParserRRETURN)

			localctx.(*Return_instrContext)._RRETURN = _m
		}
		{
			p.SetState(581)
			p.Match(ProjectParserSEMICOLON)
		}
		localctx.(*Return_instrContext).instr = Natives.NewReturn((func() int {
			if localctx.(*Return_instrContext).Get_RRETURN() == nil {
				return 0
			} else {
				return localctx.(*Return_instrContext).Get_RRETURN().GetLine()
			}
		}()), localctx.(*Return_instrContext).Get_RRETURN().GetColumn(), nil)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(583)

			var _m = p.Match(ProjectParserRRETURN)

			localctx.(*Return_instrContext)._RRETURN = _m
		}
		{
			p.SetState(584)

			var _x = p.Expression()

			localctx.(*Return_instrContext)._expression = _x
		}
		p.SetState(586)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(585)
				p.Match(ProjectParserSEMICOLON)
			}

		}
		localctx.(*Return_instrContext).instr = Natives.NewReturn((func() int {
			if localctx.(*Return_instrContext).Get_RRETURN() == nil {
				return 0
			} else {
				return localctx.(*Return_instrContext).Get_RRETURN().GetLine()
			}
		}()), localctx.(*Return_instrContext).Get_RRETURN().GetColumn(), localctx.(*Return_instrContext).Get_expression().GetP())

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// Get_conditional_prod returns the _conditional_prod rule contexts.
	Get_conditional_prod() IConditional_prodContext

	// Get_loop_prod returns the _loop_prod rule contexts.
	Get_loop_prod() ILoop_prodContext

	// Get_expr_rel returns the _expr_rel rule contexts.
	Get_expr_rel() IExpr_relContext

	// Get_expr_logic returns the _expr_logic rule contexts.
	Get_expr_logic() IExpr_logicContext

	// Get_arraydata returns the _arraydata rule contexts.
	Get_arraydata() IArraydataContext

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// Set_conditional_prod sets the _conditional_prod rule contexts.
	Set_conditional_prod(IConditional_prodContext)

	// Set_loop_prod sets the _loop_prod rule contexts.
	Set_loop_prod(ILoop_prodContext)

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

	// Set_expr_logic sets the _expr_logic rule contexts.
	Set_expr_logic(IExpr_logicContext)

	// Set_arraydata sets the _arraydata rule contexts.
	Set_arraydata(IArraydataContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	p                 Abstract.Expression
	_expr_valor       IExpr_valorContext
	_conditional_prod IConditional_prodContext
	_loop_prod        ILoop_prodContext
	_expr_rel         IExpr_relContext
	_expr_logic       IExpr_logicContext
	_arraydata        IArraydataContext
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

func (s *ExpressionContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *ExpressionContext) Get_conditional_prod() IConditional_prodContext {
	return s._conditional_prod
}

func (s *ExpressionContext) Get_loop_prod() ILoop_prodContext { return s._loop_prod }

func (s *ExpressionContext) Get_expr_rel() IExpr_relContext { return s._expr_rel }

func (s *ExpressionContext) Get_expr_logic() IExpr_logicContext { return s._expr_logic }

func (s *ExpressionContext) Get_arraydata() IArraydataContext { return s._arraydata }

func (s *ExpressionContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *ExpressionContext) Set_conditional_prod(v IConditional_prodContext) { s._conditional_prod = v }

func (s *ExpressionContext) Set_loop_prod(v ILoop_prodContext) { s._loop_prod = v }

func (s *ExpressionContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *ExpressionContext) Set_expr_logic(v IExpr_logicContext) { s._expr_logic = v }

func (s *ExpressionContext) Set_arraydata(v IArraydataContext) { s._arraydata = v }

func (s *ExpressionContext) GetP() Abstract.Expression { return s.p }

func (s *ExpressionContext) SetP(v Abstract.Expression) { s.p = v }

func (s *ExpressionContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *ExpressionContext) Conditional_prod() IConditional_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditional_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditional_prodContext)
}

func (s *ExpressionContext) Loop_prod() ILoop_prodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoop_prodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoop_prodContext)
}

func (s *ExpressionContext) Expr_rel() IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *ExpressionContext) Expr_logic() IExpr_logicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_logicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_logicContext)
}

func (s *ExpressionContext) Arraydata() IArraydataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArraydataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArraydataContext)
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
	p.EnterRule(localctx, 72, ProjectParserRULE_expression)

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

	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(592)

			var _x = p.Expr_valor()

			localctx.(*ExpressionContext)._expr_valor = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_valor().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(595)

			var _x = p.Conditional_prod()

			localctx.(*ExpressionContext)._conditional_prod = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_conditional_prod().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(598)

			var _x = p.Loop_prod()

			localctx.(*ExpressionContext)._loop_prod = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_loop_prod().GetP()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(601)

			var _x = p.expr_rel(0)

			localctx.(*ExpressionContext)._expr_rel = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_rel().GetP()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(604)

			var _x = p.Expr_logic()

			localctx.(*ExpressionContext)._expr_logic = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_logic().GetP()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(607)

			var _x = p.Arraydata()

			localctx.(*ExpressionContext)._arraydata = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_arraydata().GetP()

	}

	return localctx
}

// IArraydataContext is an interface to support dynamic dispatch.
type IArraydataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listExpressions returns the _listExpressions rule contexts.
	Get_listExpressions() IListExpressionsContext

	// Set_listExpressions sets the _listExpressions rule contexts.
	Set_listExpressions(IListExpressionsContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsArraydataContext differentiates from other interfaces.
	IsArraydataContext()
}

type ArraydataContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	p                Abstract.Expression
	_listExpressions IListExpressionsContext
}

func NewEmptyArraydataContext() *ArraydataContext {
	var p = new(ArraydataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_arraydata
	return p
}

func (*ArraydataContext) IsArraydataContext() {}

func NewArraydataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraydataContext {
	var p = new(ArraydataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_arraydata

	return p
}

func (s *ArraydataContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraydataContext) Get_listExpressions() IListExpressionsContext { return s._listExpressions }

func (s *ArraydataContext) Set_listExpressions(v IListExpressionsContext) { s._listExpressions = v }

func (s *ArraydataContext) GetP() Abstract.Expression { return s.p }

func (s *ArraydataContext) SetP(v Abstract.Expression) { s.p = v }

func (s *ArraydataContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_BRACKET, 0)
}

func (s *ArraydataContext) ListExpressions() IListExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListExpressionsContext)
}

func (s *ArraydataContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_BRACKET, 0)
}

func (s *ArraydataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraydataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArraydataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterArraydata(s)
	}
}

func (s *ArraydataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitArraydata(s)
	}
}

func (p *ProjectParser) Arraydata() (localctx IArraydataContext) {
	localctx = NewArraydataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ProjectParserRULE_arraydata)

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
		p.SetState(612)
		p.Match(ProjectParserLEFT_BRACKET)
	}
	{
		p.SetState(613)

		var _x = p.listExpressions(0)

		localctx.(*ArraydataContext)._listExpressions = _x
	}
	{
		p.SetState(614)
		p.Match(ProjectParserRIGHT_BRACKET)
	}
	localctx.(*ArraydataContext).p = ExpressionSpecial.NewValueArray(localctx.(*ArraydataContext).Get_listExpressions().GetL())

	return localctx
}

// IAccess_arrayContext is an interface to support dynamic dispatch.
type IAccess_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_listInArray returns the _listInArray rule contexts.
	Get_listInArray() IListInArrayContext

	// Set_listInArray sets the _listInArray rule contexts.
	Set_listInArray(IListInArrayContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsAccess_arrayContext differentiates from other interfaces.
	IsAccess_arrayContext()
}

type Access_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	p            Abstract.Expression
	instr        Abstract.Instruction
	_ID          antlr.Token
	_listInArray IListInArrayContext
}

func NewEmptyAccess_arrayContext() *Access_arrayContext {
	var p = new(Access_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_access_array
	return p
}

func (*Access_arrayContext) IsAccess_arrayContext() {}

func NewAccess_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Access_arrayContext {
	var p = new(Access_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_access_array

	return p
}

func (s *Access_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Access_arrayContext) Get_ID() antlr.Token { return s._ID }

func (s *Access_arrayContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Access_arrayContext) Get_listInArray() IListInArrayContext { return s._listInArray }

func (s *Access_arrayContext) Set_listInArray(v IListInArrayContext) { s._listInArray = v }

func (s *Access_arrayContext) GetP() Abstract.Expression { return s.p }

func (s *Access_arrayContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Access_arrayContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Access_arrayContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Access_arrayContext) ID() antlr.TerminalNode {
	return s.GetToken(ProjectParserID, 0)
}

func (s *Access_arrayContext) ListInArray() IListInArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListInArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListInArrayContext)
}

func (s *Access_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Access_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Access_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterAccess_array(s)
	}
}

func (s *Access_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitAccess_array(s)
	}
}

func (p *ProjectParser) Access_array() (localctx IAccess_arrayContext) {
	localctx = NewAccess_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ProjectParserRULE_access_array)

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
		p.SetState(617)

		var _m = p.Match(ProjectParserID)

		localctx.(*Access_arrayContext)._ID = _m
	}
	{
		p.SetState(618)

		var _x = p.listInArray(0)

		localctx.(*Access_arrayContext)._listInArray = _x
	}

	localctx.(*Access_arrayContext).p = Access.NewAccessArray((func() string {
		if localctx.(*Access_arrayContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Access_arrayContext).Get_ID().GetText()
		}
	}()), localctx.(*Access_arrayContext).Get_listInArray().GetL())
	localctx.(*Access_arrayContext).instr = Access.NewAccessArray((func() string {
		if localctx.(*Access_arrayContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Access_arrayContext).Get_ID().GetText()
		}
	}()), localctx.(*Access_arrayContext).Get_listInArray().GetL())

	return localctx
}

// IListInArrayContext is an interface to support dynamic dispatch.
type IListInArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSublist returns the sublist rule contexts.
	GetSublist() IListInArrayContext

	// Get_inArray returns the _inArray rule contexts.
	Get_inArray() IInArrayContext

	// SetSublist sets the sublist rule contexts.
	SetSublist(IListInArrayContext)

	// Set_inArray sets the _inArray rule contexts.
	Set_inArray(IInArrayContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListInArrayContext differentiates from other interfaces.
	IsListInArrayContext()
}

type ListInArrayContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	l        *arrayList.List
	sublist  IListInArrayContext
	_inArray IInArrayContext
}

func NewEmptyListInArrayContext() *ListInArrayContext {
	var p = new(ListInArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_listInArray
	return p
}

func (*ListInArrayContext) IsListInArrayContext() {}

func NewListInArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListInArrayContext {
	var p = new(ListInArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_listInArray

	return p
}

func (s *ListInArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListInArrayContext) GetSublist() IListInArrayContext { return s.sublist }

func (s *ListInArrayContext) Get_inArray() IInArrayContext { return s._inArray }

func (s *ListInArrayContext) SetSublist(v IListInArrayContext) { s.sublist = v }

func (s *ListInArrayContext) Set_inArray(v IInArrayContext) { s._inArray = v }

func (s *ListInArrayContext) GetL() *arrayList.List { return s.l }

func (s *ListInArrayContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListInArrayContext) InArray() IInArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInArrayContext)
}

func (s *ListInArrayContext) ListInArray() IListInArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListInArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListInArrayContext)
}

func (s *ListInArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListInArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListInArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterListInArray(s)
	}
}

func (s *ListInArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitListInArray(s)
	}
}

func (p *ProjectParser) ListInArray() (localctx IListInArrayContext) {
	return p.listInArray(0)
}

func (p *ProjectParser) listInArray(_p int) (localctx IListInArrayContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListInArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListInArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, ProjectParserRULE_listInArray, _p)

	localctx.(*ListInArrayContext).l = arrayList.New()

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
		p.SetState(622)

		var _x = p.InArray()

		localctx.(*ListInArrayContext)._inArray = _x
	}
	localctx.(*ListInArrayContext).l.Add(localctx.(*ListInArrayContext).Get_inArray().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(631)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListInArrayContext(p, _parentctx, _parentState)
			localctx.(*ListInArrayContext).sublist = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listInArray)
			p.SetState(625)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(626)

				var _x = p.InArray()

				localctx.(*ListInArrayContext)._inArray = _x
			}

			localctx.(*ListInArrayContext).GetSublist().GetL().Add(localctx.(*ListInArrayContext).Get_inArray().GetP())
			localctx.(*ListInArrayContext).l = localctx.(*ListInArrayContext).GetSublist().GetL()

		}
		p.SetState(633)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
	}

	return localctx
}

// IInArrayContext is an interface to support dynamic dispatch.
type IInArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsInArrayContext differentiates from other interfaces.
	IsInArrayContext()
}

type InArrayContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           Abstract.Expression
	_expression IExpressionContext
}

func NewEmptyInArrayContext() *InArrayContext {
	var p = new(InArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_inArray
	return p
}

func (*InArrayContext) IsInArrayContext() {}

func NewInArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InArrayContext {
	var p = new(InArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_inArray

	return p
}

func (s *InArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *InArrayContext) Get_expression() IExpressionContext { return s._expression }

func (s *InArrayContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *InArrayContext) GetP() Abstract.Expression { return s.p }

func (s *InArrayContext) SetP(v Abstract.Expression) { s.p = v }

func (s *InArrayContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_BRACKET, 0)
}

func (s *InArrayContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InArrayContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_BRACKET, 0)
}

func (s *InArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterInArray(s)
	}
}

func (s *InArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitInArray(s)
	}
}

func (p *ProjectParser) InArray() (localctx IInArrayContext) {
	localctx = NewInArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ProjectParserRULE_inArray)

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
		p.SetState(634)
		p.Match(ProjectParserLEFT_BRACKET)
	}
	{
		p.SetState(635)

		var _x = p.Expression()

		localctx.(*InArrayContext)._expression = _x
	}
	{
		p.SetState(636)
		p.Match(ProjectParserRIGHT_BRACKET)
	}
	localctx.(*InArrayContext).p = localctx.(*InArrayContext).Get_expression().GetP()

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

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsExpr_relContext differentiates from other interfaces.
	IsExpr_relContext()
}

type Expr_relContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          Abstract.Expression
	instr      Abstract.Instruction
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

func (s *Expr_relContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Expr_relContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Expr_relContext) SetInstr(v Abstract.Instruction) { s.instr = v }

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

func (s *Expr_relContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(ProjectParserNOTEQUAL, 0)
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
	_startState := 82
	p.EnterRecursionRule(localctx, 82, ProjectParserRULE_expr_rel, _p)
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
		p.SetState(640)

		var _x = p.expr_arit(0)

		localctx.(*Expr_relContext)._expr_arit = _x
	}

	localctx.(*Expr_relContext).p = localctx.(*Expr_relContext).Get_expr_arit().GetP()
	localctx.(*Expr_relContext).instr = localctx.(*Expr_relContext).Get_expr_arit().GetInstr()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(650)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_relContext(p, _parentctx, _parentState)
			localctx.(*Expr_relContext).opLeft = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_expr_rel)
			p.SetState(643)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(644)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expr_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(ProjectParserEQUEAL_EQUAL-52))|(1<<(ProjectParserNOTEQUAL-52))|(1<<(ProjectParserGREATER_THAN-52))|(1<<(ProjectParserLESS_THAN-52))|(1<<(ProjectParserGREATER_EQUALTHAN-52))|(1<<(ProjectParserLESS_EQUEALTHAN-52)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expr_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(645)

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
			localctx.(*Expr_relContext).instr = Expression.NewOperation(localctx.(*Expr_relContext).GetOpLeft().GetP(), (func() string {
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
		p.SetState(652)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

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

	// Get_pow_op returns the _pow_op rule contexts.
	Get_pow_op() IPow_opContext

	// GetOpRight returns the opRight rule contexts.
	GetOpRight() IExpr_aritContext

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// SetOpLeft sets the opLeft rule contexts.
	SetOpLeft(IExpr_aritContext)

	// SetOpU sets the opU rule contexts.
	SetOpU(IExpressionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_pow_op sets the _pow_op rule contexts.
	Set_pow_op(IPow_opContext)

	// SetOpRight sets the opRight rule contexts.
	SetOpRight(IExpr_aritContext)

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsExpr_aritContext differentiates from other interfaces.
	IsExpr_aritContext()
}

type Expr_aritContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           Abstract.Expression
	instr       Abstract.Instruction
	opLeft      IExpr_aritContext
	opU         IExpressionContext
	_expression IExpressionContext
	_pow_op     IPow_opContext
	opRight     IExpr_aritContext
	_expr_valor IExpr_valorContext
	op          antlr.Token
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

func (s *Expr_aritContext) Get_pow_op() IPow_opContext { return s._pow_op }

func (s *Expr_aritContext) GetOpRight() IExpr_aritContext { return s.opRight }

func (s *Expr_aritContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *Expr_aritContext) SetOpLeft(v IExpr_aritContext) { s.opLeft = v }

func (s *Expr_aritContext) SetOpU(v IExpressionContext) { s.opU = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_pow_op(v IPow_opContext) { s._pow_op = v }

func (s *Expr_aritContext) SetOpRight(v IExpr_aritContext) { s.opRight = v }

func (s *Expr_aritContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *Expr_aritContext) GetP() Abstract.Expression { return s.p }

func (s *Expr_aritContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Expr_aritContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Expr_aritContext) SetInstr(v Abstract.Instruction) { s.instr = v }

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

func (s *Expr_aritContext) Pow_op() IPow_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPow_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPow_opContext)
}

func (s *Expr_aritContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *Expr_aritContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ProjectParserCOMMA, 0)
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

func (s *Expr_aritContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
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
	_startState := 84
	p.EnterRecursionRule(localctx, 84, ProjectParserRULE_expr_arit, _p)
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
	p.SetState(674)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(654)
			p.Match(ProjectParserSUB)
		}
		{
			p.SetState(655)

			var _x = p.Expression()

			localctx.(*Expr_aritContext).opU = _x
			localctx.(*Expr_aritContext)._expression = _x
		}

		localctx.(*Expr_aritContext).p = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpU().GetP(), "-", nil, true, localctx.(*Expr_aritContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_aritContext).GetOpU().GetStart().GetColumn())
		localctx.(*Expr_aritContext).instr = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpU().GetP(), "-", nil, true, localctx.(*Expr_aritContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_aritContext).GetOpU().GetStart().GetColumn())

	case 2:
		{
			p.SetState(658)

			var _x = p.Pow_op()

			localctx.(*Expr_aritContext)._pow_op = _x
		}
		{
			p.SetState(659)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(660)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opLeft = _x
		}
		{
			p.SetState(661)
			p.Match(ProjectParserCOMMA)
		}
		{
			p.SetState(662)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opRight = _x
		}
		{
			p.SetState(663)
			p.Match(ProjectParserRIGHT_PAR)
		}

		localctx.(*Expr_aritContext).p = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), localctx.(*Expr_aritContext).Get_pow_op().GetOp(), localctx.(*Expr_aritContext).GetOpRight().GetP(), false, localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetLine(), localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetColumn())
		localctx.(*Expr_aritContext).instr = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), localctx.(*Expr_aritContext).Get_pow_op().GetOp(), localctx.(*Expr_aritContext).GetOpRight().GetP(), false, localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetLine(), localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetColumn())

	case 3:
		{
			p.SetState(666)

			var _x = p.Expr_valor()

			localctx.(*Expr_aritContext)._expr_valor = _x
		}

		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expr_valor().GetP()
		localctx.(*Expr_aritContext).instr = localctx.(*Expr_aritContext).Get_expr_valor().GetInstr()

	case 4:
		{
			p.SetState(669)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(670)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(671)
			p.Match(ProjectParserRIGHT_PAR)
		}

		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()
		localctx.(*Expr_aritContext).instr = nil

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(688)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(686)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opLeft = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_expr_arit)
				p.SetState(676)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(677)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(ProjectParserMUL-58))|(1<<(ProjectParserDIV-58))|(1<<(ProjectParserMOD-58)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(678)

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
				localctx.(*Expr_aritContext).instr = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), (func() string {
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
				p.SetState(681)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(682)

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
					p.SetState(683)

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
				localctx.(*Expr_aritContext).instr = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), (func() string {
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
		p.SetState(690)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr_valorContext is an interface to support dynamic dispatch.
type IExpr_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_called_func returns the _called_func rule contexts.
	Get_called_func() ICalled_funcContext

	// Get_primitive returns the _primitive rule contexts.
	Get_primitive() IPrimitiveContext

	// Get_expr_cast returns the _expr_cast rule contexts.
	Get_expr_cast() IExpr_castContext

	// Get_access_array returns the _access_array rule contexts.
	Get_access_array() IAccess_arrayContext

	// Set_called_func sets the _called_func rule contexts.
	Set_called_func(ICalled_funcContext)

	// Set_primitive sets the _primitive rule contexts.
	Set_primitive(IPrimitiveContext)

	// Set_expr_cast sets the _expr_cast rule contexts.
	Set_expr_cast(IExpr_castContext)

	// Set_access_array sets the _access_array rule contexts.
	Set_access_array(IAccess_arrayContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsExpr_valorContext differentiates from other interfaces.
	IsExpr_valorContext()
}

type Expr_valorContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	p             Abstract.Expression
	instr         Abstract.Instruction
	_called_func  ICalled_funcContext
	_primitive    IPrimitiveContext
	_expr_cast    IExpr_castContext
	_access_array IAccess_arrayContext
}

func NewEmptyExpr_valorContext() *Expr_valorContext {
	var p = new(Expr_valorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_expr_valor
	return p
}

func (*Expr_valorContext) IsExpr_valorContext() {}

func NewExpr_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_valorContext {
	var p = new(Expr_valorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_expr_valor

	return p
}

func (s *Expr_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_valorContext) Get_called_func() ICalled_funcContext { return s._called_func }

func (s *Expr_valorContext) Get_primitive() IPrimitiveContext { return s._primitive }

func (s *Expr_valorContext) Get_expr_cast() IExpr_castContext { return s._expr_cast }

func (s *Expr_valorContext) Get_access_array() IAccess_arrayContext { return s._access_array }

func (s *Expr_valorContext) Set_called_func(v ICalled_funcContext) { s._called_func = v }

func (s *Expr_valorContext) Set_primitive(v IPrimitiveContext) { s._primitive = v }

func (s *Expr_valorContext) Set_expr_cast(v IExpr_castContext) { s._expr_cast = v }

func (s *Expr_valorContext) Set_access_array(v IAccess_arrayContext) { s._access_array = v }

func (s *Expr_valorContext) GetP() Abstract.Expression { return s.p }

func (s *Expr_valorContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Expr_valorContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Expr_valorContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Expr_valorContext) Called_func() ICalled_funcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalled_funcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalled_funcContext)
}

func (s *Expr_valorContext) Primitive() IPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *Expr_valorContext) Expr_cast() IExpr_castContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_castContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_castContext)
}

func (s *Expr_valorContext) Access_array() IAccess_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccess_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccess_arrayContext)
}

func (s *Expr_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_valorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterExpr_valor(s)
	}
}

func (s *Expr_valorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitExpr_valor(s)
	}
}

func (p *ProjectParser) Expr_valor() (localctx IExpr_valorContext) {
	localctx = NewExpr_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ProjectParserRULE_expr_valor)

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

	p.SetState(703)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(691)

			var _x = p.Called_func()

			localctx.(*Expr_valorContext)._called_func = _x
		}

		localctx.(*Expr_valorContext).p = localctx.(*Expr_valorContext).Get_called_func().GetP()
		localctx.(*Expr_valorContext).instr = localctx.(*Expr_valorContext).Get_called_func().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(694)

			var _x = p.Primitive()

			localctx.(*Expr_valorContext)._primitive = _x
		}

		localctx.(*Expr_valorContext).p = localctx.(*Expr_valorContext).Get_primitive().GetP()
		localctx.(*Expr_valorContext).instr = localctx.(*Expr_valorContext).Get_primitive().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(697)

			var _x = p.Expr_cast()

			localctx.(*Expr_valorContext)._expr_cast = _x
		}

		localctx.(*Expr_valorContext).p = localctx.(*Expr_valorContext).Get_expr_cast().GetP()
		localctx.(*Expr_valorContext).instr = localctx.(*Expr_valorContext).Get_expr_cast().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(700)

			var _x = p.Access_array()

			localctx.(*Expr_valorContext)._access_array = _x
		}

		localctx.(*Expr_valorContext).p = localctx.(*Expr_valorContext).Get_access_array().GetP()
		localctx.(*Expr_valorContext).instr = localctx.(*Expr_valorContext).Get_access_array().GetInstr()

	}

	return localctx
}

// IPow_opContext is an interface to support dynamic dispatch.
type IPow_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_POWI returns the _POWI token.
	Get_POWI() antlr.Token

	// Get_POWF returns the _POWF token.
	Get_POWF() antlr.Token

	// Set_POWI sets the _POWI token.
	Set_POWI(antlr.Token)

	// Set_POWF sets the _POWF token.
	Set_POWF(antlr.Token)

	// GetOp returns the op attribute.
	GetOp() string

	// SetOp sets the op attribute.
	SetOp(string)

	// IsPow_opContext differentiates from other interfaces.
	IsPow_opContext()
}

type Pow_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     string
	_POWI  antlr.Token
	_POWF  antlr.Token
}

func NewEmptyPow_opContext() *Pow_opContext {
	var p = new(Pow_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_pow_op
	return p
}

func (*Pow_opContext) IsPow_opContext() {}

func NewPow_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pow_opContext {
	var p = new(Pow_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_pow_op

	return p
}

func (s *Pow_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Pow_opContext) Get_POWI() antlr.Token { return s._POWI }

func (s *Pow_opContext) Get_POWF() antlr.Token { return s._POWF }

func (s *Pow_opContext) Set_POWI(v antlr.Token) { s._POWI = v }

func (s *Pow_opContext) Set_POWF(v antlr.Token) { s._POWF = v }

func (s *Pow_opContext) GetOp() string { return s.op }

func (s *Pow_opContext) SetOp(v string) { s.op = v }

func (s *Pow_opContext) RINTEGER() antlr.TerminalNode {
	return s.GetToken(ProjectParserRINTEGER, 0)
}

func (s *Pow_opContext) HERITAGE() antlr.TerminalNode {
	return s.GetToken(ProjectParserHERITAGE, 0)
}

func (s *Pow_opContext) POWI() antlr.TerminalNode {
	return s.GetToken(ProjectParserPOWI, 0)
}

func (s *Pow_opContext) RREAL() antlr.TerminalNode {
	return s.GetToken(ProjectParserRREAL, 0)
}

func (s *Pow_opContext) POWF() antlr.TerminalNode {
	return s.GetToken(ProjectParserPOWF, 0)
}

func (s *Pow_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pow_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pow_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterPow_op(s)
	}
}

func (s *Pow_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitPow_op(s)
	}
}

func (p *ProjectParser) Pow_op() (localctx IPow_opContext) {
	localctx = NewPow_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ProjectParserRULE_pow_op)

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

	p.SetState(713)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserRINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(705)
			p.Match(ProjectParserRINTEGER)
		}
		{
			p.SetState(706)
			p.Match(ProjectParserHERITAGE)
		}
		{
			p.SetState(707)

			var _m = p.Match(ProjectParserPOWI)

			localctx.(*Pow_opContext)._POWI = _m
		}
		localctx.(*Pow_opContext).op = (func() string {
			if localctx.(*Pow_opContext).Get_POWI() == nil {
				return ""
			} else {
				return localctx.(*Pow_opContext).Get_POWI().GetText()
			}
		}())

	case ProjectParserRREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(709)
			p.Match(ProjectParserRREAL)
		}
		{
			p.SetState(710)
			p.Match(ProjectParserHERITAGE)
		}
		{
			p.SetState(711)

			var _m = p.Match(ProjectParserPOWF)

			localctx.(*Pow_opContext)._POWF = _m
		}
		localctx.(*Pow_opContext).op = (func() string {
			if localctx.(*Pow_opContext).Get_POWF() == nil {
				return ""
			} else {
				return localctx.(*Pow_opContext).Get_POWF().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	// GetOpU returns the opU rule contexts.
	GetOpU() IExpressionContext

	// GetOpLeft returns the opLeft rule contexts.
	GetOpLeft() IExpr_relContext

	// GetOpRight returns the opRight rule contexts.
	GetOpRight() IExpr_relContext

	// SetOpU sets the opU rule contexts.
	SetOpU(IExpressionContext)

	// SetOpLeft sets the opLeft rule contexts.
	SetOpLeft(IExpr_relContext)

	// SetOpRight sets the opRight rule contexts.
	SetOpRight(IExpr_relContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsExpr_logicContext differentiates from other interfaces.
	IsExpr_logicContext()
}

type Expr_logicContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	p       Abstract.Expression
	instr   Abstract.Instruction
	opU     IExpressionContext
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

func (s *Expr_logicContext) GetOpU() IExpressionContext { return s.opU }

func (s *Expr_logicContext) GetOpLeft() IExpr_relContext { return s.opLeft }

func (s *Expr_logicContext) GetOpRight() IExpr_relContext { return s.opRight }

func (s *Expr_logicContext) SetOpU(v IExpressionContext) { s.opU = v }

func (s *Expr_logicContext) SetOpLeft(v IExpr_relContext) { s.opLeft = v }

func (s *Expr_logicContext) SetOpRight(v IExpr_relContext) { s.opRight = v }

func (s *Expr_logicContext) GetP() Abstract.Expression { return s.p }

func (s *Expr_logicContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Expr_logicContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Expr_logicContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Expr_logicContext) ADMIRATION() antlr.TerminalNode {
	return s.GetToken(ProjectParserADMIRATION, 0)
}

func (s *Expr_logicContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

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

func (s *Expr_logicContext) AND() antlr.TerminalNode {
	return s.GetToken(ProjectParserAND, 0)
}

func (s *Expr_logicContext) OR() antlr.TerminalNode {
	return s.GetToken(ProjectParserOR, 0)
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
	p.EnterRule(localctx, 90, ProjectParserRULE_expr_logic)
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

	p.SetState(724)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserADMIRATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(715)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(716)

			var _x = p.Expression()

			localctx.(*Expr_logicContext).opU = _x
		}

		localctx.(*Expr_logicContext).p = Expression.NewOperation(localctx.(*Expr_logicContext).GetOpU().GetP(), "!", nil, true, localctx.(*Expr_logicContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_logicContext).GetOpU().GetStart().GetColumn())
		localctx.(*Expr_logicContext).instr = Expression.NewOperation(localctx.(*Expr_logicContext).GetOpU().GetP(), "!", nil, true, localctx.(*Expr_logicContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_logicContext).GetOpU().GetStart().GetColumn())

	case ProjectParserRINTEGER, ProjectParserRREAL, ProjectParserINTEGER, ProjectParserFLOAT, ProjectParserCHAR, ProjectParserSTRING, ProjectParserBOOLEAN, ProjectParserID, ProjectParserLEFT_PAR, ProjectParserSUB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(719)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logicContext).opLeft = _x
		}
		{
			p.SetState(720)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_logicContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ProjectParserAND || _la == ProjectParserOR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_logicContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(721)

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
		localctx.(*Expr_logicContext).instr = Expression.NewOperation(localctx.(*Expr_logicContext).GetOpLeft().GetP(), (func() string {
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

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpr_castContext is an interface to support dynamic dispatch.
type IExpr_castContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RAS returns the _RAS token.
	Get_RAS() antlr.Token

	// Set_RAS sets the _RAS token.
	Set_RAS(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_data_type returns the _data_type rule contexts.
	Get_data_type() IData_typeContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_data_type sets the _data_type rule contexts.
	Set_data_type(IData_typeContext)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsExpr_castContext differentiates from other interfaces.
	IsExpr_castContext()
}

type Expr_castContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           Abstract.Expression
	instr       Abstract.Instruction
	_expression IExpressionContext
	_RAS        antlr.Token
	_data_type  IData_typeContext
}

func NewEmptyExpr_castContext() *Expr_castContext {
	var p = new(Expr_castContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_expr_cast
	return p
}

func (*Expr_castContext) IsExpr_castContext() {}

func NewExpr_castContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_castContext {
	var p = new(Expr_castContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_expr_cast

	return p
}

func (s *Expr_castContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_castContext) Get_RAS() antlr.Token { return s._RAS }

func (s *Expr_castContext) Set_RAS(v antlr.Token) { s._RAS = v }

func (s *Expr_castContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_castContext) Get_data_type() IData_typeContext { return s._data_type }

func (s *Expr_castContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_castContext) Set_data_type(v IData_typeContext) { s._data_type = v }

func (s *Expr_castContext) GetP() Abstract.Expression { return s.p }

func (s *Expr_castContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *Expr_castContext) SetP(v Abstract.Expression) { s.p = v }

func (s *Expr_castContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *Expr_castContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *Expr_castContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expr_castContext) RAS() antlr.TerminalNode {
	return s.GetToken(ProjectParserRAS, 0)
}

func (s *Expr_castContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Expr_castContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_PAR, 0)
}

func (s *Expr_castContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_castContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_castContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterExpr_cast(s)
	}
}

func (s *Expr_castContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitExpr_cast(s)
	}
}

func (p *ProjectParser) Expr_cast() (localctx IExpr_castContext) {
	localctx = NewExpr_castContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ProjectParserRULE_expr_cast)

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
		p.SetState(726)
		p.Match(ProjectParserLEFT_PAR)
	}
	{
		p.SetState(727)

		var _x = p.Expression()

		localctx.(*Expr_castContext)._expression = _x
	}
	{
		p.SetState(728)

		var _m = p.Match(ProjectParserRAS)

		localctx.(*Expr_castContext)._RAS = _m
	}
	{
		p.SetState(729)

		var _x = p.Data_type()

		localctx.(*Expr_castContext)._data_type = _x
	}
	{
		p.SetState(730)
		p.Match(ProjectParserRIGHT_PAR)
	}

	if localctx.(*Expr_castContext).Get_data_type().GetData() == "i64" {
		localctx.(*Expr_castContext).p = Expression.NewCast(localctx.(*Expr_castContext).Get_expression().GetP(), SymbolTable.INTEGER, (func() int {
			if localctx.(*Expr_castContext).Get_RAS() == nil {
				return 0
			} else {
				return localctx.(*Expr_castContext).Get_RAS().GetLine()
			}
		}()), localctx.(*Expr_castContext).Get_RAS().GetColumn())
		localctx.(*Expr_castContext).instr = Expression.NewCast(localctx.(*Expr_castContext).Get_expression().GetP(), SymbolTable.INTEGER, (func() int {
			if localctx.(*Expr_castContext).Get_RAS() == nil {
				return 0
			} else {
				return localctx.(*Expr_castContext).Get_RAS().GetLine()
			}
		}()), localctx.(*Expr_castContext).Get_RAS().GetColumn())
	} else if localctx.(*Expr_castContext).Get_data_type().GetData() == "f64" {
		localctx.(*Expr_castContext).p = Expression.NewCast(localctx.(*Expr_castContext).Get_expression().GetP(), SymbolTable.FLOAT, (func() int {
			if localctx.(*Expr_castContext).Get_RAS() == nil {
				return 0
			} else {
				return localctx.(*Expr_castContext).Get_RAS().GetLine()
			}
		}()), localctx.(*Expr_castContext).Get_RAS().GetColumn())
		localctx.(*Expr_castContext).instr = Expression.NewCast(localctx.(*Expr_castContext).Get_expression().GetP(), SymbolTable.FLOAT, (func() int {
			if localctx.(*Expr_castContext).Get_RAS() == nil {
				return 0
			} else {
				return localctx.(*Expr_castContext).Get_RAS().GetLine()
			}
		}()), localctx.(*Expr_castContext).Get_RAS().GetColumn())
	}

	return localctx
}

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RINTEGER returns the _RINTEGER token.
	Get_RINTEGER() antlr.Token

	// Get_RREAL returns the _RREAL token.
	Get_RREAL() antlr.Token

	// Get_RSTR returns the _RSTR token.
	Get_RSTR() antlr.Token

	// Get_RSTRING returns the _RSTRING token.
	Get_RSTRING() antlr.Token

	// Get_RBOOLEAN returns the _RBOOLEAN token.
	Get_RBOOLEAN() antlr.Token

	// Get_RCHAR returns the _RCHAR token.
	Get_RCHAR() antlr.Token

	// Set_RINTEGER sets the _RINTEGER token.
	Set_RINTEGER(antlr.Token)

	// Set_RREAL sets the _RREAL token.
	Set_RREAL(antlr.Token)

	// Set_RSTR sets the _RSTR token.
	Set_RSTR(antlr.Token)

	// Set_RSTRING sets the _RSTRING token.
	Set_RSTRING(antlr.Token)

	// Set_RBOOLEAN sets the _RBOOLEAN token.
	Set_RBOOLEAN(antlr.Token)

	// Set_RCHAR sets the _RCHAR token.
	Set_RCHAR(antlr.Token)

	// GetData returns the data attribute.
	GetData() string

	// SetData sets the data attribute.
	SetData(string)

	// IsData_typeContext differentiates from other interfaces.
	IsData_typeContext()
}

type Data_typeContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	data      string
	_RINTEGER antlr.Token
	_RREAL    antlr.Token
	_RSTR     antlr.Token
	_RSTRING  antlr.Token
	_RBOOLEAN antlr.Token
	_RCHAR    antlr.Token
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_data_type
	return p
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) Get_RINTEGER() antlr.Token { return s._RINTEGER }

func (s *Data_typeContext) Get_RREAL() antlr.Token { return s._RREAL }

func (s *Data_typeContext) Get_RSTR() antlr.Token { return s._RSTR }

func (s *Data_typeContext) Get_RSTRING() antlr.Token { return s._RSTRING }

func (s *Data_typeContext) Get_RBOOLEAN() antlr.Token { return s._RBOOLEAN }

func (s *Data_typeContext) Get_RCHAR() antlr.Token { return s._RCHAR }

func (s *Data_typeContext) Set_RINTEGER(v antlr.Token) { s._RINTEGER = v }

func (s *Data_typeContext) Set_RREAL(v antlr.Token) { s._RREAL = v }

func (s *Data_typeContext) Set_RSTR(v antlr.Token) { s._RSTR = v }

func (s *Data_typeContext) Set_RSTRING(v antlr.Token) { s._RSTRING = v }

func (s *Data_typeContext) Set_RBOOLEAN(v antlr.Token) { s._RBOOLEAN = v }

func (s *Data_typeContext) Set_RCHAR(v antlr.Token) { s._RCHAR = v }

func (s *Data_typeContext) GetData() string { return s.data }

func (s *Data_typeContext) SetData(v string) { s.data = v }

func (s *Data_typeContext) RINTEGER() antlr.TerminalNode {
	return s.GetToken(ProjectParserRINTEGER, 0)
}

func (s *Data_typeContext) RREAL() antlr.TerminalNode {
	return s.GetToken(ProjectParserRREAL, 0)
}

func (s *Data_typeContext) RSTR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRSTR, 0)
}

func (s *Data_typeContext) RSTRING() antlr.TerminalNode {
	return s.GetToken(ProjectParserRSTRING, 0)
}

func (s *Data_typeContext) RBOOLEAN() antlr.TerminalNode {
	return s.GetToken(ProjectParserRBOOLEAN, 0)
}

func (s *Data_typeContext) RCHAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRCHAR, 0)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.EnterData_type(s)
	}
}

func (s *Data_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectParserListener); ok {
		listenerT.ExitData_type(s)
	}
}

func (p *ProjectParser) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ProjectParserRULE_data_type)

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

	p.SetState(745)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserRINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(733)

			var _m = p.Match(ProjectParserRINTEGER)

			localctx.(*Data_typeContext)._RINTEGER = _m
		}
		localctx.(*Data_typeContext).data = (func() string {
			if localctx.(*Data_typeContext).Get_RINTEGER() == nil {
				return ""
			} else {
				return localctx.(*Data_typeContext).Get_RINTEGER().GetText()
			}
		}())

	case ProjectParserRREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(735)

			var _m = p.Match(ProjectParserRREAL)

			localctx.(*Data_typeContext)._RREAL = _m
		}
		localctx.(*Data_typeContext).data = (func() string {
			if localctx.(*Data_typeContext).Get_RREAL() == nil {
				return ""
			} else {
				return localctx.(*Data_typeContext).Get_RREAL().GetText()
			}
		}())

	case ProjectParserRSTR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(737)

			var _m = p.Match(ProjectParserRSTR)

			localctx.(*Data_typeContext)._RSTR = _m
		}
		localctx.(*Data_typeContext).data = (func() string {
			if localctx.(*Data_typeContext).Get_RSTR() == nil {
				return ""
			} else {
				return localctx.(*Data_typeContext).Get_RSTR().GetText()
			}
		}())

	case ProjectParserRSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(739)

			var _m = p.Match(ProjectParserRSTRING)

			localctx.(*Data_typeContext)._RSTRING = _m
		}
		localctx.(*Data_typeContext).data = (func() string {
			if localctx.(*Data_typeContext).Get_RSTRING() == nil {
				return ""
			} else {
				return localctx.(*Data_typeContext).Get_RSTRING().GetText()
			}
		}())

	case ProjectParserRBOOLEAN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(741)

			var _m = p.Match(ProjectParserRBOOLEAN)

			localctx.(*Data_typeContext)._RBOOLEAN = _m
		}
		localctx.(*Data_typeContext).data = (func() string {
			if localctx.(*Data_typeContext).Get_RBOOLEAN() == nil {
				return ""
			} else {
				return localctx.(*Data_typeContext).Get_RBOOLEAN().GetText()
			}
		}())

	case ProjectParserRCHAR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(743)

			var _m = p.Match(ProjectParserRCHAR)

			localctx.(*Data_typeContext)._RCHAR = _m
		}
		localctx.(*Data_typeContext).data = (func() string {
			if localctx.(*Data_typeContext).Get_RCHAR() == nil {
				return ""
			} else {
				return localctx.(*Data_typeContext).Get_RCHAR().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

	// Get_TOSTRING returns the _TOSTRING token.
	Get_TOSTRING() antlr.Token

	// Get_TOOWNED returns the _TOOWNED token.
	Get_TOOWNED() antlr.Token

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

	// Set_TOSTRING sets the _TOSTRING token.
	Set_TOSTRING(antlr.Token)

	// Set_TOOWNED sets the _TOOWNED token.
	Set_TOOWNED(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// Set_BOOLEAN sets the _BOOLEAN token.
	Set_BOOLEAN(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetP returns the p attribute.
	GetP() Abstract.Expression

	// GetInstr returns the instr attribute.
	GetInstr() Abstract.Instruction

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// SetInstr sets the instr attribute.
	SetInstr(Abstract.Instruction)

	// IsPrimitiveContext differentiates from other interfaces.
	IsPrimitiveContext()
}

type PrimitiveContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         Abstract.Expression
	instr     Abstract.Instruction
	_INTEGER  antlr.Token
	_FLOAT    antlr.Token
	_STRING   antlr.Token
	_TOSTRING antlr.Token
	_TOOWNED  antlr.Token
	_CHAR     antlr.Token
	_BOOLEAN  antlr.Token
	_ID       antlr.Token
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

func (s *PrimitiveContext) Get_TOSTRING() antlr.Token { return s._TOSTRING }

func (s *PrimitiveContext) Get_TOOWNED() antlr.Token { return s._TOOWNED }

func (s *PrimitiveContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *PrimitiveContext) Get_BOOLEAN() antlr.Token { return s._BOOLEAN }

func (s *PrimitiveContext) Get_ID() antlr.Token { return s._ID }

func (s *PrimitiveContext) Set_INTEGER(v antlr.Token) { s._INTEGER = v }

func (s *PrimitiveContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *PrimitiveContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitiveContext) Set_TOSTRING(v antlr.Token) { s._TOSTRING = v }

func (s *PrimitiveContext) Set_TOOWNED(v antlr.Token) { s._TOOWNED = v }

func (s *PrimitiveContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *PrimitiveContext) Set_BOOLEAN(v antlr.Token) { s._BOOLEAN = v }

func (s *PrimitiveContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *PrimitiveContext) GetP() Abstract.Expression { return s.p }

func (s *PrimitiveContext) GetInstr() Abstract.Instruction { return s.instr }

func (s *PrimitiveContext) SetP(v Abstract.Expression) { s.p = v }

func (s *PrimitiveContext) SetInstr(v Abstract.Instruction) { s.instr = v }

func (s *PrimitiveContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ProjectParserINTEGER, 0)
}

func (s *PrimitiveContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ProjectParserFLOAT, 0)
}

func (s *PrimitiveContext) STRING() antlr.TerminalNode {
	return s.GetToken(ProjectParserSTRING, 0)
}

func (s *PrimitiveContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserLEFT_PAR, 0)
}

func (s *PrimitiveContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(ProjectParserRIGHT_PAR, 0)
}

func (s *PrimitiveContext) DOT() antlr.TerminalNode {
	return s.GetToken(ProjectParserDOT, 0)
}

func (s *PrimitiveContext) TOSTRING() antlr.TerminalNode {
	return s.GetToken(ProjectParserTOSTRING, 0)
}

func (s *PrimitiveContext) TOOWNED() antlr.TerminalNode {
	return s.GetToken(ProjectParserTOOWNED, 0)
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
	p.EnterRule(localctx, 96, ProjectParserRULE_primitive)

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

	p.SetState(769)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(747)

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
		localctx.(*PrimitiveContext).instr = Expression.NewPrimitive(num, SymbolTable.INTEGER, (func() int {
			if localctx.(*PrimitiveContext).Get_INTEGER() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_INTEGER().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_INTEGER().GetColumn())

	case ProjectParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(749)

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
		localctx.(*PrimitiveContext).instr = Expression.NewPrimitive(num, SymbolTable.FLOAT, (func() int {
			if localctx.(*PrimitiveContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_FLOAT().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_FLOAT().GetColumn())

	case ProjectParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(751)

			var _m = p.Match(ProjectParserSTRING)

			localctx.(*PrimitiveContext)._STRING = _m
		}
		p.SetState(760)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
			p.SetState(756)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(752)
					p.Match(ProjectParserDOT)
				}
				{
					p.SetState(753)

					var _m = p.Match(ProjectParserTOSTRING)

					localctx.(*PrimitiveContext)._TOSTRING = _m
				}

			case 2:
				{
					p.SetState(754)
					p.Match(ProjectParserDOT)
				}
				{
					p.SetState(755)

					var _m = p.Match(ProjectParserTOOWNED)

					localctx.(*PrimitiveContext)._TOOWNED = _m
				}

			}
			{
				p.SetState(758)
				p.Match(ProjectParserLEFT_PAR)
			}
			{
				p.SetState(759)
				p.Match(ProjectParserRIGHT_PAR)
			}

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
		if (func() string {
			if localctx.(*PrimitiveContext).Get_TOSTRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_TOSTRING().GetText()
			}
		}()) != "" || (func() string {
			if localctx.(*PrimitiveContext).Get_TOOWNED() == nil {
				return ""
			} else {
				return localctx.(*PrimitiveContext).Get_TOOWNED().GetText()
			}
		}()) != "" {
			localctx.(*PrimitiveContext).p = Expression.NewPrimitive(str, SymbolTable.STRING, (func() int {
				if localctx.(*PrimitiveContext).Get_STRING() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_STRING().GetLine()
				}
			}()), localctx.(*PrimitiveContext).Get_STRING().GetColumn())
			localctx.(*PrimitiveContext).instr = Expression.NewPrimitive(str, SymbolTable.STRING, (func() int {
				if localctx.(*PrimitiveContext).Get_STRING() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_STRING().GetLine()
				}
			}()), localctx.(*PrimitiveContext).Get_STRING().GetColumn())
		} else {
			localctx.(*PrimitiveContext).p = Expression.NewPrimitive(str, SymbolTable.STR, (func() int {
				if localctx.(*PrimitiveContext).Get_STRING() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_STRING().GetLine()
				}
			}()), localctx.(*PrimitiveContext).Get_STRING().GetColumn())
			localctx.(*PrimitiveContext).instr = Expression.NewPrimitive(str, SymbolTable.STR, (func() int {
				if localctx.(*PrimitiveContext).Get_STRING() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_STRING().GetLine()
				}
			}()), localctx.(*PrimitiveContext).Get_STRING().GetColumn())
		}

	case ProjectParserCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(763)

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
		localctx.(*PrimitiveContext).instr = Expression.NewPrimitive(chr, SymbolTable.CHAR, (func() int {
			if localctx.(*PrimitiveContext).Get_CHAR() == nil {
				return 0
			} else {
				return localctx.(*PrimitiveContext).Get_CHAR().GetLine()
			}
		}()), localctx.(*PrimitiveContext).Get_CHAR().GetColumn())

	case ProjectParserBOOLEAN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(765)

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
		if str == "true" {
			localctx.(*PrimitiveContext).p = Expression.NewPrimitive(true, SymbolTable.BOOLEAN, (func() int {
				if localctx.(*PrimitiveContext).Get_BOOLEAN() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_BOOLEAN().GetLine()
				}
			}()), localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
			localctx.(*PrimitiveContext).instr = Expression.NewPrimitive(true, SymbolTable.BOOLEAN, (func() int {
				if localctx.(*PrimitiveContext).Get_BOOLEAN() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_BOOLEAN().GetLine()
				}
			}()), localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
		} else {
			localctx.(*PrimitiveContext).p = Expression.NewPrimitive(false, SymbolTable.BOOLEAN, (func() int {
				if localctx.(*PrimitiveContext).Get_BOOLEAN() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_BOOLEAN().GetLine()
				}
			}()), localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
			localctx.(*PrimitiveContext).instr = Expression.NewPrimitive(false, SymbolTable.BOOLEAN, (func() int {
				if localctx.(*PrimitiveContext).Get_BOOLEAN() == nil {
					return 0
				} else {
					return localctx.(*PrimitiveContext).Get_BOOLEAN().GetLine()
				}
			}()), localctx.(*PrimitiveContext).Get_BOOLEAN().GetColumn())
		}

	case ProjectParserID:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(767)

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
		localctx.(*PrimitiveContext).instr = Expression.NewIdentifier((func() string {
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
	case 1:
		var t *ListFuncsContext = nil
		if localctx != nil {
			t = localctx.(*ListFuncsContext)
		}
		return p.ListFuncs_Sempred(t, predIndex)

	case 6:
		var t *ListParamsContext = nil
		if localctx != nil {
			t = localctx.(*ListParamsContext)
		}
		return p.ListParams_Sempred(t, predIndex)

	case 10:
		var t *ListVarsContext = nil
		if localctx != nil {
			t = localctx.(*ListVarsContext)
		}
		return p.ListVars_Sempred(t, predIndex)

	case 14:
		var t *ListIdsContext = nil
		if localctx != nil {
			t = localctx.(*ListIdsContext)
		}
		return p.ListIds_Sempred(t, predIndex)

	case 20:
		var t *List_instr_matchContext = nil
		if localctx != nil {
			t = localctx.(*List_instr_matchContext)
		}
		return p.List_instr_match_Sempred(t, predIndex)

	case 22:
		var t *Expr_matchContext = nil
		if localctx != nil {
			t = localctx.(*Expr_matchContext)
		}
		return p.Expr_match_Sempred(t, predIndex)

	case 29:
		var t *ListExpressionsContext = nil
		if localctx != nil {
			t = localctx.(*ListExpressionsContext)
		}
		return p.ListExpressions_Sempred(t, predIndex)

	case 39:
		var t *ListInArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListInArrayContext)
		}
		return p.ListInArray_Sempred(t, predIndex)

	case 41:
		var t *Expr_relContext = nil
		if localctx != nil {
			t = localctx.(*Expr_relContext)
		}
		return p.Expr_rel_Sempred(t, predIndex)

	case 42:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ProjectParser) ListFuncs_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) ListParams_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) ListVars_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) ListIds_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) List_instr_match_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) Expr_match_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) ListExpressions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) ListInArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) Expr_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ProjectParser) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
