// Code generated from OptimizerParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // OptimizerParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2_Project1/server/interpreter/AST"
import "OLC2_Project1/server/interpreter/Optimizer/AbstractOptimizer"
import "OLC2_Project1/server/interpreter/Optimizer/Primitive"
import "OLC2_Project1/server/interpreter/Optimizer/Assignment"
import "OLC2_Project1/server/interpreter/Optimizer/Control"
import "OLC2_Project1/server/interpreter/Optimizer/Function"
import "OLC2_Project1/server/interpreter/Optimizer/Print"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 336,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	93, 10, 5, 12, 5, 14, 5, 96, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 7, 6, 106, 10, 6, 12, 6, 14, 6, 109, 11, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 7, 8, 121, 10, 8, 12, 8, 14,
	8, 124, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 158,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 226, 10, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 267, 10, 21, 3, 22, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 281,
	10, 23, 12, 23, 14, 23, 284, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 299, 10,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	7, 24, 311, 10, 24, 12, 24, 14, 24, 314, 11, 24, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 5, 25, 322, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 334, 10, 26, 3, 26, 2, 6, 8,
	10, 44, 46, 27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 2, 5, 3, 2, 38, 43, 3, 2, 44, 45,
	3, 2, 47, 48, 2, 336, 2, 52, 3, 2, 2, 2, 4, 57, 3, 2, 2, 2, 6, 79, 3, 2,
	2, 2, 8, 84, 3, 2, 2, 2, 10, 97, 3, 2, 2, 2, 12, 110, 3, 2, 2, 2, 14, 122,
	3, 2, 2, 2, 16, 157, 3, 2, 2, 2, 18, 159, 3, 2, 2, 2, 20, 171, 3, 2, 2,
	2, 22, 183, 3, 2, 2, 2, 24, 189, 3, 2, 2, 2, 26, 198, 3, 2, 2, 2, 28, 203,
	3, 2, 2, 2, 30, 207, 3, 2, 2, 2, 32, 225, 3, 2, 2, 2, 34, 227, 3, 2, 2,
	2, 36, 239, 3, 2, 2, 2, 38, 251, 3, 2, 2, 2, 40, 266, 3, 2, 2, 2, 42, 268,
	3, 2, 2, 2, 44, 271, 3, 2, 2, 2, 46, 298, 3, 2, 2, 2, 48, 321, 3, 2, 2,
	2, 50, 333, 3, 2, 2, 2, 52, 53, 5, 4, 3, 2, 53, 54, 5, 6, 4, 2, 54, 55,
	5, 10, 6, 2, 55, 56, 8, 2, 1, 2, 56, 3, 3, 2, 2, 2, 57, 58, 7, 7, 2, 2,
	58, 59, 7, 8, 2, 2, 59, 60, 7, 4, 2, 2, 60, 61, 7, 10, 2, 2, 61, 62, 7,
	36, 2, 2, 62, 63, 7, 22, 2, 2, 63, 64, 7, 37, 2, 2, 64, 65, 7, 28, 2, 2,
	65, 66, 7, 4, 2, 2, 66, 67, 7, 11, 2, 2, 67, 68, 7, 36, 2, 2, 68, 69, 7,
	22, 2, 2, 69, 70, 7, 37, 2, 2, 70, 71, 7, 28, 2, 2, 71, 72, 7, 4, 2, 2,
	72, 73, 7, 18, 2, 2, 73, 74, 7, 28, 2, 2, 74, 75, 7, 4, 2, 2, 75, 76, 7,
	19, 2, 2, 76, 77, 7, 28, 2, 2, 77, 78, 8, 3, 1, 2, 78, 5, 3, 2, 2, 2, 79,
	80, 7, 4, 2, 2, 80, 81, 5, 8, 5, 2, 81, 82, 7, 28, 2, 2, 82, 83, 8, 4,
	1, 2, 83, 7, 3, 2, 2, 2, 84, 85, 8, 5, 1, 2, 85, 86, 7, 26, 2, 2, 86, 87,
	8, 5, 1, 2, 87, 94, 3, 2, 2, 2, 88, 89, 12, 4, 2, 2, 89, 90, 7, 29, 2,
	2, 90, 91, 7, 26, 2, 2, 91, 93, 8, 5, 1, 2, 92, 88, 3, 2, 2, 2, 93, 96,
	3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 9, 3, 2, 2, 2,
	96, 94, 3, 2, 2, 2, 97, 98, 8, 6, 1, 2, 98, 99, 5, 12, 7, 2, 99, 100, 8,
	6, 1, 2, 100, 107, 3, 2, 2, 2, 101, 102, 12, 4, 2, 2, 102, 103, 5, 12,
	7, 2, 103, 104, 8, 6, 1, 2, 104, 106, 3, 2, 2, 2, 105, 101, 3, 2, 2, 2,
	106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	11, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 111, 7, 6, 2, 2, 111, 112, 7,
	26, 2, 2, 112, 113, 7, 32, 2, 2, 113, 114, 7, 33, 2, 2, 114, 115, 7, 34,
	2, 2, 115, 116, 5, 14, 8, 2, 116, 117, 7, 35, 2, 2, 117, 118, 8, 7, 1,
	2, 118, 13, 3, 2, 2, 2, 119, 121, 5, 16, 9, 2, 120, 119, 3, 2, 2, 2, 121,
	124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 125,
	3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 126, 8, 8, 1, 2, 126, 15, 3, 2,
	2, 2, 127, 128, 5, 18, 10, 2, 128, 129, 8, 9, 1, 2, 129, 158, 3, 2, 2,
	2, 130, 131, 5, 20, 11, 2, 131, 132, 8, 9, 1, 2, 132, 158, 3, 2, 2, 2,
	133, 134, 5, 22, 12, 2, 134, 135, 8, 9, 1, 2, 135, 158, 3, 2, 2, 2, 136,
	137, 5, 24, 13, 2, 137, 138, 8, 9, 1, 2, 138, 158, 3, 2, 2, 2, 139, 140,
	5, 26, 14, 2, 140, 141, 8, 9, 1, 2, 141, 158, 3, 2, 2, 2, 142, 143, 5,
	28, 15, 2, 143, 144, 8, 9, 1, 2, 144, 158, 3, 2, 2, 2, 145, 146, 5, 30,
	16, 2, 146, 147, 8, 9, 1, 2, 147, 158, 3, 2, 2, 2, 148, 149, 5, 32, 17,
	2, 149, 150, 8, 9, 1, 2, 150, 158, 3, 2, 2, 2, 151, 152, 5, 34, 18, 2,
	152, 153, 8, 9, 1, 2, 153, 158, 3, 2, 2, 2, 154, 155, 5, 36, 19, 2, 155,
	156, 8, 9, 1, 2, 156, 158, 3, 2, 2, 2, 157, 127, 3, 2, 2, 2, 157, 130,
	3, 2, 2, 2, 157, 133, 3, 2, 2, 2, 157, 136, 3, 2, 2, 2, 157, 139, 3, 2,
	2, 2, 157, 142, 3, 2, 2, 2, 157, 145, 3, 2, 2, 2, 157, 148, 3, 2, 2, 2,
	157, 151, 3, 2, 2, 2, 157, 154, 3, 2, 2, 2, 158, 17, 3, 2, 2, 2, 159, 160,
	7, 11, 2, 2, 160, 161, 7, 36, 2, 2, 161, 162, 7, 32, 2, 2, 162, 163, 5,
	48, 25, 2, 163, 164, 7, 33, 2, 2, 164, 165, 5, 50, 26, 2, 165, 166, 7,
	37, 2, 2, 166, 167, 7, 27, 2, 2, 167, 168, 5, 42, 22, 2, 168, 169, 7, 28,
	2, 2, 169, 170, 8, 10, 1, 2, 170, 19, 3, 2, 2, 2, 171, 172, 7, 10, 2, 2,
	172, 173, 7, 36, 2, 2, 173, 174, 7, 32, 2, 2, 174, 175, 5, 48, 25, 2, 175,
	176, 7, 33, 2, 2, 176, 177, 5, 50, 26, 2, 177, 178, 7, 37, 2, 2, 178, 179,
	7, 27, 2, 2, 179, 180, 5, 42, 22, 2, 180, 181, 7, 28, 2, 2, 181, 182, 8,
	11, 1, 2, 182, 21, 3, 2, 2, 2, 183, 184, 5, 50, 26, 2, 184, 185, 7, 27,
	2, 2, 185, 186, 5, 42, 22, 2, 186, 187, 7, 28, 2, 2, 187, 188, 8, 12, 1,
	2, 188, 23, 3, 2, 2, 2, 189, 190, 7, 12, 2, 2, 190, 191, 7, 32, 2, 2, 191,
	192, 5, 42, 22, 2, 192, 193, 7, 33, 2, 2, 193, 194, 7, 13, 2, 2, 194, 195,
	7, 26, 2, 2, 195, 196, 7, 28, 2, 2, 196, 197, 8, 13, 1, 2, 197, 25, 3,
	2, 2, 2, 198, 199, 7, 13, 2, 2, 199, 200, 7, 26, 2, 2, 200, 201, 7, 28,
	2, 2, 201, 202, 8, 14, 1, 2, 202, 27, 3, 2, 2, 2, 203, 204, 7, 26, 2, 2,
	204, 205, 7, 30, 2, 2, 205, 206, 8, 15, 1, 2, 206, 29, 3, 2, 2, 2, 207,
	208, 7, 15, 2, 2, 208, 209, 7, 32, 2, 2, 209, 210, 7, 25, 2, 2, 210, 211,
	7, 29, 2, 2, 211, 212, 5, 40, 21, 2, 212, 213, 5, 38, 20, 2, 213, 214,
	7, 33, 2, 2, 214, 215, 7, 28, 2, 2, 215, 216, 8, 16, 1, 2, 216, 31, 3,
	2, 2, 2, 217, 218, 7, 14, 2, 2, 218, 219, 5, 42, 22, 2, 219, 220, 7, 28,
	2, 2, 220, 221, 8, 17, 1, 2, 221, 226, 3, 2, 2, 2, 222, 223, 7, 14, 2,
	2, 223, 224, 7, 28, 2, 2, 224, 226, 8, 17, 1, 2, 225, 217, 3, 2, 2, 2,
	225, 222, 3, 2, 2, 2, 226, 33, 3, 2, 2, 2, 227, 228, 5, 50, 26, 2, 228,
	229, 7, 27, 2, 2, 229, 230, 7, 10, 2, 2, 230, 231, 7, 36, 2, 2, 231, 232,
	7, 32, 2, 2, 232, 233, 5, 48, 25, 2, 233, 234, 7, 33, 2, 2, 234, 235, 5,
	42, 22, 2, 235, 236, 7, 37, 2, 2, 236, 237, 7, 28, 2, 2, 237, 238, 8, 18,
	1, 2, 238, 35, 3, 2, 2, 2, 239, 240, 5, 50, 26, 2, 240, 241, 7, 27, 2,
	2, 241, 242, 7, 11, 2, 2, 242, 243, 7, 36, 2, 2, 243, 244, 7, 32, 2, 2,
	244, 245, 5, 48, 25, 2, 245, 246, 7, 33, 2, 2, 246, 247, 5, 42, 22, 2,
	247, 248, 7, 37, 2, 2, 248, 249, 7, 28, 2, 2, 249, 250, 8, 19, 1, 2, 250,
	37, 3, 2, 2, 2, 251, 252, 5, 50, 26, 2, 252, 253, 8, 20, 1, 2, 253, 39,
	3, 2, 2, 2, 254, 255, 7, 32, 2, 2, 255, 256, 7, 3, 2, 2, 256, 257, 7, 33,
	2, 2, 257, 267, 8, 21, 1, 2, 258, 259, 7, 32, 2, 2, 259, 260, 7, 4, 2,
	2, 260, 261, 7, 33, 2, 2, 261, 267, 8, 21, 1, 2, 262, 263, 7, 32, 2, 2,
	263, 264, 7, 5, 2, 2, 264, 265, 7, 33, 2, 2, 265, 267, 8, 21, 1, 2, 266,
	254, 3, 2, 2, 2, 266, 258, 3, 2, 2, 2, 266, 262, 3, 2, 2, 2, 267, 41, 3,
	2, 2, 2, 268, 269, 5, 44, 23, 2, 269, 270, 8, 22, 1, 2, 270, 43, 3, 2,
	2, 2, 271, 272, 8, 23, 1, 2, 272, 273, 5, 46, 24, 2, 273, 274, 8, 23, 1,
	2, 274, 282, 3, 2, 2, 2, 275, 276, 12, 4, 2, 2, 276, 277, 9, 2, 2, 2, 277,
	278, 5, 44, 23, 5, 278, 279, 8, 23, 1, 2, 279, 281, 3, 2, 2, 2, 280, 275,
	3, 2, 2, 2, 281, 284, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 282, 283, 3, 2,
	2, 2, 283, 45, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 285, 286, 8, 24, 1, 2,
	286, 287, 7, 48, 2, 2, 287, 288, 5, 42, 22, 2, 288, 289, 8, 24, 1, 2, 289,
	299, 3, 2, 2, 2, 290, 291, 5, 50, 26, 2, 291, 292, 8, 24, 1, 2, 292, 299,
	3, 2, 2, 2, 293, 294, 7, 32, 2, 2, 294, 295, 5, 42, 22, 2, 295, 296, 7,
	33, 2, 2, 296, 297, 8, 24, 1, 2, 297, 299, 3, 2, 2, 2, 298, 285, 3, 2,
	2, 2, 298, 290, 3, 2, 2, 2, 298, 293, 3, 2, 2, 2, 299, 312, 3, 2, 2, 2,
	300, 301, 12, 6, 2, 2, 301, 302, 9, 3, 2, 2, 302, 303, 5, 46, 24, 7, 303,
	304, 8, 24, 1, 2, 304, 311, 3, 2, 2, 2, 305, 306, 12, 5, 2, 2, 306, 307,
	9, 4, 2, 2, 307, 308, 5, 46, 24, 6, 308, 309, 8, 24, 1, 2, 309, 311, 3,
	2, 2, 2, 310, 300, 3, 2, 2, 2, 310, 305, 3, 2, 2, 2, 311, 314, 3, 2, 2,
	2, 312, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 47, 3, 2, 2, 2, 314,
	312, 3, 2, 2, 2, 315, 316, 7, 3, 2, 2, 316, 322, 8, 25, 1, 2, 317, 318,
	7, 4, 2, 2, 318, 322, 8, 25, 1, 2, 319, 320, 7, 5, 2, 2, 320, 322, 8, 25,
	1, 2, 321, 315, 3, 2, 2, 2, 321, 317, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2,
	322, 49, 3, 2, 2, 2, 323, 324, 7, 23, 2, 2, 324, 334, 8, 26, 1, 2, 325,
	326, 7, 22, 2, 2, 326, 334, 8, 26, 1, 2, 327, 328, 7, 26, 2, 2, 328, 334,
	8, 26, 1, 2, 329, 330, 7, 18, 2, 2, 330, 334, 8, 26, 1, 2, 331, 332, 7,
	19, 2, 2, 332, 334, 8, 26, 1, 2, 333, 323, 3, 2, 2, 2, 333, 325, 3, 2,
	2, 2, 333, 327, 3, 2, 2, 2, 333, 329, 3, 2, 2, 2, 333, 331, 3, 2, 2, 2,
	334, 51, 3, 2, 2, 2, 14, 94, 107, 122, 157, 225, 266, 282, 298, 310, 312,
	321, 333,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'int'", "'float'", "'char'", "'void'", "'#include'", "'<stdio.h>'",
	"'<math.h>'", "'heap'", "'stack'", "'if'", "'goto'", "'return'", "'printf'",
	"'pow'", "'mod'", "'P'", "'H'", "", "", "", "", "", "", "", "'='", "';'",
	"','", "':'", "'!'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'=='",
	"'!='", "'>'", "'<'", "'>='", "'<='", "'*'", "'/'", "'%'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "RINT", "RFLOAT", "RCHAR", "RVOID", "RINCLUDE", "RSTDIO", "RMATH",
	"RHEAP", "RSTACK", "RIF", "RGOTO", "RRETURN", "RPRINTF", "RPOW", "RMOD",
	"RPSTACK", "RPHEAP", "MULTILINE", "INLINE", "INTEGER", "FLOAT", "CHAR",
	"STRING", "ID", "EQUAL", "SEMICOLON", "COMMA", "COLON", "ADMIRATION", "LEFT_PAR",
	"RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET", "RIGHT_BRACKET",
	"EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN", "GREATER_EQUALTHAN",
	"LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB", "WHITESPACE",
}

var ruleNames = []string{
	"start", "headers", "list_temps", "temps", "listFuncs", "function", "instructions",
	"instruction", "assign_stack", "assign_heap", "assign", "if_instr", "goto_instr",
	"label_instr", "printf_instr", "return_instr", "getHeap_instr", "getStack_instr",
	"expr_print", "convert", "expression", "expr_rel", "expr_arit", "data_type",
	"expr_valor",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type OptimizerParser struct {
	*antlr.BaseParser
}

func NewOptimizerParser(input antlr.TokenStream) *OptimizerParser {
	this := new(OptimizerParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "OptimizerParser.g4"

	return this
}

// OptimizerParser tokens.
const (
	OptimizerParserEOF               = antlr.TokenEOF
	OptimizerParserRINT              = 1
	OptimizerParserRFLOAT            = 2
	OptimizerParserRCHAR             = 3
	OptimizerParserRVOID             = 4
	OptimizerParserRINCLUDE          = 5
	OptimizerParserRSTDIO            = 6
	OptimizerParserRMATH             = 7
	OptimizerParserRHEAP             = 8
	OptimizerParserRSTACK            = 9
	OptimizerParserRIF               = 10
	OptimizerParserRGOTO             = 11
	OptimizerParserRRETURN           = 12
	OptimizerParserRPRINTF           = 13
	OptimizerParserRPOW              = 14
	OptimizerParserRMOD              = 15
	OptimizerParserRPSTACK           = 16
	OptimizerParserRPHEAP            = 17
	OptimizerParserMULTILINE         = 18
	OptimizerParserINLINE            = 19
	OptimizerParserINTEGER           = 20
	OptimizerParserFLOAT             = 21
	OptimizerParserCHAR              = 22
	OptimizerParserSTRING            = 23
	OptimizerParserID                = 24
	OptimizerParserEQUAL             = 25
	OptimizerParserSEMICOLON         = 26
	OptimizerParserCOMMA             = 27
	OptimizerParserCOLON             = 28
	OptimizerParserADMIRATION        = 29
	OptimizerParserLEFT_PAR          = 30
	OptimizerParserRIGHT_PAR         = 31
	OptimizerParserLEFT_KEY          = 32
	OptimizerParserRIGHT_KEY         = 33
	OptimizerParserLEFT_BRACKET      = 34
	OptimizerParserRIGHT_BRACKET     = 35
	OptimizerParserEQUEAL_EQUAL      = 36
	OptimizerParserNOTEQUAL          = 37
	OptimizerParserGREATER_THAN      = 38
	OptimizerParserLESS_THAN         = 39
	OptimizerParserGREATER_EQUALTHAN = 40
	OptimizerParserLESS_EQUEALTHAN   = 41
	OptimizerParserMUL               = 42
	OptimizerParserDIV               = 43
	OptimizerParserMOD               = 44
	OptimizerParserADD               = 45
	OptimizerParserSUB               = 46
	OptimizerParserWHITESPACE        = 47
)

// OptimizerParser rules.
const (
	OptimizerParserRULE_start          = 0
	OptimizerParserRULE_headers        = 1
	OptimizerParserRULE_list_temps     = 2
	OptimizerParserRULE_temps          = 3
	OptimizerParserRULE_listFuncs      = 4
	OptimizerParserRULE_function       = 5
	OptimizerParserRULE_instructions   = 6
	OptimizerParserRULE_instruction    = 7
	OptimizerParserRULE_assign_stack   = 8
	OptimizerParserRULE_assign_heap    = 9
	OptimizerParserRULE_assign         = 10
	OptimizerParserRULE_if_instr       = 11
	OptimizerParserRULE_goto_instr     = 12
	OptimizerParserRULE_label_instr    = 13
	OptimizerParserRULE_printf_instr   = 14
	OptimizerParserRULE_return_instr   = 15
	OptimizerParserRULE_getHeap_instr  = 16
	OptimizerParserRULE_getStack_instr = 17
	OptimizerParserRULE_expr_print     = 18
	OptimizerParserRULE_convert        = 19
	OptimizerParserRULE_expression     = 20
	OptimizerParserRULE_expr_rel       = 21
	OptimizerParserRULE_expr_arit      = 22
	OptimizerParserRULE_data_type      = 23
	OptimizerParserRULE_expr_valor     = 24
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_list_temps returns the _list_temps rule contexts.
	Get_list_temps() IList_tempsContext

	// Get_listFuncs returns the _listFuncs rule contexts.
	Get_listFuncs() IListFuncsContext

	// Set_list_temps sets the _list_temps rule contexts.
	Set_list_temps(IList_tempsContext)

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
	parser      antlr.Parser
	tree        AST.Tree
	_list_temps IList_tempsContext
	_listFuncs  IListFuncsContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_list_temps() IList_tempsContext { return s._list_temps }

func (s *StartContext) Get_listFuncs() IListFuncsContext { return s._listFuncs }

func (s *StartContext) Set_list_temps(v IList_tempsContext) { s._list_temps = v }

func (s *StartContext) Set_listFuncs(v IListFuncsContext) { s._listFuncs = v }

func (s *StartContext) GetTree() AST.Tree { return s.tree }

func (s *StartContext) SetTree(v AST.Tree) { s.tree = v }

func (s *StartContext) Headers() IHeadersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeadersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeadersContext)
}

func (s *StartContext) List_temps() IList_tempsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_tempsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_tempsContext)
}

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
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *OptimizerParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OptimizerParserRULE_start)

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
		p.SetState(50)
		p.Headers()
	}
	{
		p.SetState(51)

		var _x = p.List_temps()

		localctx.(*StartContext)._list_temps = _x
	}
	{
		p.SetState(52)

		var _x = p.listFuncs(0)

		localctx.(*StartContext)._listFuncs = _x
	}
	localctx.(*StartContext).tree = AST.NewTreeOptimizer(localctx.(*StartContext).Get_listFuncs().GetL(), localctx.(*StartContext).Get_list_temps().GetL())

	return localctx
}

// IHeadersContext is an interface to support dynamic dispatch.
type IHeadersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RINCLUDE returns the _RINCLUDE token.
	Get_RINCLUDE() antlr.Token

	// Set_RINCLUDE sets the _RINCLUDE token.
	Set_RINCLUDE(antlr.Token)

	// GetCode returns the code attribute.
	GetCode() string

	// SetCode sets the code attribute.
	SetCode(string)

	// IsHeadersContext differentiates from other interfaces.
	IsHeadersContext()
}

type HeadersContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	code      string
	_RINCLUDE antlr.Token
}

func NewEmptyHeadersContext() *HeadersContext {
	var p = new(HeadersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_headers
	return p
}

func (*HeadersContext) IsHeadersContext() {}

func NewHeadersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeadersContext {
	var p = new(HeadersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_headers

	return p
}

func (s *HeadersContext) GetParser() antlr.Parser { return s.parser }

func (s *HeadersContext) Get_RINCLUDE() antlr.Token { return s._RINCLUDE }

func (s *HeadersContext) Set_RINCLUDE(v antlr.Token) { s._RINCLUDE = v }

func (s *HeadersContext) GetCode() string { return s.code }

func (s *HeadersContext) SetCode(v string) { s.code = v }

func (s *HeadersContext) RINCLUDE() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRINCLUDE, 0)
}

func (s *HeadersContext) RSTDIO() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRSTDIO, 0)
}

func (s *HeadersContext) AllRFLOAT() []antlr.TerminalNode {
	return s.GetTokens(OptimizerParserRFLOAT)
}

func (s *HeadersContext) RFLOAT(i int) antlr.TerminalNode {
	return s.GetToken(OptimizerParserRFLOAT, i)
}

func (s *HeadersContext) RHEAP() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRHEAP, 0)
}

func (s *HeadersContext) AllLEFT_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(OptimizerParserLEFT_BRACKET)
}

func (s *HeadersContext) LEFT_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_BRACKET, i)
}

func (s *HeadersContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(OptimizerParserINTEGER)
}

func (s *HeadersContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(OptimizerParserINTEGER, i)
}

func (s *HeadersContext) AllRIGHT_BRACKET() []antlr.TerminalNode {
	return s.GetTokens(OptimizerParserRIGHT_BRACKET)
}

func (s *HeadersContext) RIGHT_BRACKET(i int) antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_BRACKET, i)
}

func (s *HeadersContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(OptimizerParserSEMICOLON)
}

func (s *HeadersContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, i)
}

func (s *HeadersContext) RSTACK() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRSTACK, 0)
}

func (s *HeadersContext) RPSTACK() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRPSTACK, 0)
}

func (s *HeadersContext) RPHEAP() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRPHEAP, 0)
}

func (s *HeadersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeadersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterHeaders(s)
	}
}

func (s *HeadersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitHeaders(s)
	}
}

func (p *OptimizerParser) Headers() (localctx IHeadersContext) {
	localctx = NewHeadersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OptimizerParserRULE_headers)

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
		p.SetState(55)

		var _m = p.Match(OptimizerParserRINCLUDE)

		localctx.(*HeadersContext)._RINCLUDE = _m
	}
	{
		p.SetState(56)
		p.Match(OptimizerParserRSTDIO)
	}
	{
		p.SetState(57)
		p.Match(OptimizerParserRFLOAT)
	}
	{
		p.SetState(58)
		p.Match(OptimizerParserRHEAP)
	}
	{
		p.SetState(59)
		p.Match(OptimizerParserLEFT_BRACKET)
	}
	{
		p.SetState(60)
		p.Match(OptimizerParserINTEGER)
	}
	{
		p.SetState(61)
		p.Match(OptimizerParserRIGHT_BRACKET)
	}
	{
		p.SetState(62)
		p.Match(OptimizerParserSEMICOLON)
	}
	{
		p.SetState(63)
		p.Match(OptimizerParserRFLOAT)
	}
	{
		p.SetState(64)
		p.Match(OptimizerParserRSTACK)
	}
	{
		p.SetState(65)
		p.Match(OptimizerParserLEFT_BRACKET)
	}
	{
		p.SetState(66)
		p.Match(OptimizerParserINTEGER)
	}
	{
		p.SetState(67)
		p.Match(OptimizerParserRIGHT_BRACKET)
	}
	{
		p.SetState(68)
		p.Match(OptimizerParserSEMICOLON)
	}
	{
		p.SetState(69)
		p.Match(OptimizerParserRFLOAT)
	}
	{
		p.SetState(70)
		p.Match(OptimizerParserRPSTACK)
	}
	{
		p.SetState(71)
		p.Match(OptimizerParserSEMICOLON)
	}
	{
		p.SetState(72)
		p.Match(OptimizerParserRFLOAT)
	}
	{
		p.SetState(73)
		p.Match(OptimizerParserRPHEAP)
	}
	{
		p.SetState(74)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*HeadersContext).code = (func() string {
		if localctx.(*HeadersContext).Get_RINCLUDE() == nil {
			return ""
		} else {
			return localctx.(*HeadersContext).Get_RINCLUDE().GetText()
		}
	}())

	return localctx
}

// IList_tempsContext is an interface to support dynamic dispatch.
type IList_tempsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_temps returns the _temps rule contexts.
	Get_temps() ITempsContext

	// Set_temps sets the _temps rule contexts.
	Set_temps(ITempsContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsList_tempsContext differentiates from other interfaces.
	IsList_tempsContext()
}

type List_tempsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	l      *arrayList.List
	_temps ITempsContext
}

func NewEmptyList_tempsContext() *List_tempsContext {
	var p = new(List_tempsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_list_temps
	return p
}

func (*List_tempsContext) IsList_tempsContext() {}

func NewList_tempsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_tempsContext {
	var p = new(List_tempsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_list_temps

	return p
}

func (s *List_tempsContext) GetParser() antlr.Parser { return s.parser }

func (s *List_tempsContext) Get_temps() ITempsContext { return s._temps }

func (s *List_tempsContext) Set_temps(v ITempsContext) { s._temps = v }

func (s *List_tempsContext) GetL() *arrayList.List { return s.l }

func (s *List_tempsContext) SetL(v *arrayList.List) { s.l = v }

func (s *List_tempsContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRFLOAT, 0)
}

func (s *List_tempsContext) Temps() ITempsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITempsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITempsContext)
}

func (s *List_tempsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *List_tempsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_tempsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_tempsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterList_temps(s)
	}
}

func (s *List_tempsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitList_temps(s)
	}
}

func (p *OptimizerParser) List_temps() (localctx IList_tempsContext) {
	localctx = NewList_tempsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, OptimizerParserRULE_list_temps)

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
		p.SetState(77)
		p.Match(OptimizerParserRFLOAT)
	}
	{
		p.SetState(78)

		var _x = p.temps(0)

		localctx.(*List_tempsContext)._temps = _x
	}
	{
		p.SetState(79)
		p.Match(OptimizerParserSEMICOLON)
	}
	localctx.(*List_tempsContext).l = localctx.(*List_tempsContext).Get_temps().GetL()

	return localctx
}

// ITempsContext is an interface to support dynamic dispatch.
type ITempsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetSub returns the sub rule contexts.
	GetSub() ITempsContext

	// SetSub sets the sub rule contexts.
	SetSub(ITempsContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsTempsContext differentiates from other interfaces.
	IsTempsContext()
}

type TempsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	l      *arrayList.List
	sub    ITempsContext
	_ID    antlr.Token
}

func NewEmptyTempsContext() *TempsContext {
	var p = new(TempsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_temps
	return p
}

func (*TempsContext) IsTempsContext() {}

func NewTempsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TempsContext {
	var p = new(TempsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_temps

	return p
}

func (s *TempsContext) GetParser() antlr.Parser { return s.parser }

func (s *TempsContext) Get_ID() antlr.Token { return s._ID }

func (s *TempsContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *TempsContext) GetSub() ITempsContext { return s.sub }

func (s *TempsContext) SetSub(v ITempsContext) { s.sub = v }

func (s *TempsContext) GetL() *arrayList.List { return s.l }

func (s *TempsContext) SetL(v *arrayList.List) { s.l = v }

func (s *TempsContext) ID() antlr.TerminalNode {
	return s.GetToken(OptimizerParserID, 0)
}

func (s *TempsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(OptimizerParserCOMMA, 0)
}

func (s *TempsContext) Temps() ITempsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITempsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITempsContext)
}

func (s *TempsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TempsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TempsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterTemps(s)
	}
}

func (s *TempsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitTemps(s)
	}
}

func (p *OptimizerParser) Temps() (localctx ITempsContext) {
	return p.temps(0)
}

func (p *OptimizerParser) temps(_p int) (localctx ITempsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTempsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITempsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, OptimizerParserRULE_temps, _p)

	localctx.(*TempsContext).l = arrayList.New()

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
		p.SetState(83)

		var _m = p.Match(OptimizerParserID)

		localctx.(*TempsContext)._ID = _m
	}
	localctx.(*TempsContext).l.Add((func() string {
		if localctx.(*TempsContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*TempsContext).Get_ID().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTempsContext(p, _parentctx, _parentState)
			localctx.(*TempsContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, OptimizerParserRULE_temps)
			p.SetState(86)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(87)
				p.Match(OptimizerParserCOMMA)
			}
			{
				p.SetState(88)

				var _m = p.Match(OptimizerParserID)

				localctx.(*TempsContext)._ID = _m
			}

			localctx.(*TempsContext).GetSub().GetL().Add((func() string {
				if localctx.(*TempsContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*TempsContext).Get_ID().GetText()
				}
			}()))
			localctx.(*TempsContext).l = localctx.(*TempsContext).GetSub().GetL()

		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

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
	p.RuleIndex = OptimizerParserRULE_listFuncs
	return p
}

func (*ListFuncsContext) IsListFuncsContext() {}

func NewListFuncsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListFuncsContext {
	var p = new(ListFuncsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_listFuncs

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
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterListFuncs(s)
	}
}

func (s *ListFuncsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitListFuncs(s)
	}
}

func (p *OptimizerParser) ListFuncs() (localctx IListFuncsContext) {
	return p.listFuncs(0)
}

func (p *OptimizerParser) listFuncs(_p int) (localctx IListFuncsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListFuncsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListFuncsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, OptimizerParserRULE_listFuncs, _p)

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
		p.SetState(96)

		var _x = p.Function()

		localctx.(*ListFuncsContext)._function = _x
	}
	localctx.(*ListFuncsContext).l.Add(localctx.(*ListFuncsContext).Get_function().GetInstr())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListFuncsContext(p, _parentctx, _parentState)
			localctx.(*ListFuncsContext).subList = _prevctx
			p.PushNewRecursionContext(localctx, _startState, OptimizerParserRULE_listFuncs)
			p.SetState(99)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(100)

				var _x = p.Function()

				localctx.(*ListFuncsContext)._function = _x
			}

			localctx.(*ListFuncsContext).GetSubList().GetL().Add(localctx.(*ListFuncsContext).Get_function().GetInstr())
			localctx.(*ListFuncsContext).l = localctx.(*ListFuncsContext).GetSubList().GetL()

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_instructions returns the _instructions rule contexts.
	Get_instructions() IInstructionsContext

	// Set_instructions sets the _instructions rule contexts.
	Set_instructions(IInstructionsContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	instr         AbstractOptimizer.Instruction
	_ID           antlr.Token
	_instructions IInstructionsContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *FunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FunctionContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *FunctionContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *FunctionContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *FunctionContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *FunctionContext) RVOID() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRVOID, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(OptimizerParserID, 0)
}

func (s *FunctionContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *FunctionContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
}

func (s *FunctionContext) LEFT_KEY() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_KEY, 0)
}

func (s *FunctionContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *FunctionContext) RIGHT_KEY() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_KEY, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *OptimizerParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, OptimizerParserRULE_function)

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
		p.SetState(108)
		p.Match(OptimizerParserRVOID)
	}
	{
		p.SetState(109)

		var _m = p.Match(OptimizerParserID)

		localctx.(*FunctionContext)._ID = _m
	}
	{
		p.SetState(110)
		p.Match(OptimizerParserLEFT_PAR)
	}
	{
		p.SetState(111)
		p.Match(OptimizerParserRIGHT_PAR)
	}
	{
		p.SetState(112)
		p.Match(OptimizerParserLEFT_KEY)
	}
	{
		p.SetState(113)

		var _x = p.Instructions()

		localctx.(*FunctionContext)._instructions = _x
	}
	{
		p.SetState(114)
		p.Match(OptimizerParserRIGHT_KEY)
	}

	localctx.(*FunctionContext).instr = Function.NewFunction((func() int {
		if localctx.(*FunctionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*FunctionContext).Get_ID().GetLine()
		}
	}()), localctx.(*FunctionContext).Get_ID().GetColumn(), (func() string {
		if localctx.(*FunctionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*FunctionContext).Get_ID().GetText()
		}
	}()), localctx.(*FunctionContext).Get_instructions().GetL())

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
	p.RuleIndex = OptimizerParserRULE_instructions
	return p
}

func (*InstructionsContext) IsInstructionsContext() {}

func NewInstructionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsContext {
	var p = new(InstructionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_instructions

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
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterInstructions(s)
	}
}

func (s *InstructionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitInstructions(s)
	}
}

func (p *OptimizerParser) Instructions() (localctx IInstructionsContext) {
	localctx = NewInstructionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, OptimizerParserRULE_instructions)

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
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<OptimizerParserRHEAP)|(1<<OptimizerParserRSTACK)|(1<<OptimizerParserRIF)|(1<<OptimizerParserRGOTO)|(1<<OptimizerParserRRETURN)|(1<<OptimizerParserRPRINTF)|(1<<OptimizerParserRPSTACK)|(1<<OptimizerParserRPHEAP)|(1<<OptimizerParserINTEGER)|(1<<OptimizerParserFLOAT)|(1<<OptimizerParserID))) != 0 {
		{
			p.SetState(117)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstructionsContext).GetE()
	for _, e := range listInt {
		if e.GetInstr() != nil {
			localctx.(*InstructionsContext).l.Add(e.GetInstr())
		}
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_assign_stack returns the _assign_stack rule contexts.
	Get_assign_stack() IAssign_stackContext

	// Get_assign_heap returns the _assign_heap rule contexts.
	Get_assign_heap() IAssign_heapContext

	// Get_assign returns the _assign rule contexts.
	Get_assign() IAssignContext

	// Get_if_instr returns the _if_instr rule contexts.
	Get_if_instr() IIf_instrContext

	// Get_goto_instr returns the _goto_instr rule contexts.
	Get_goto_instr() IGoto_instrContext

	// Get_label_instr returns the _label_instr rule contexts.
	Get_label_instr() ILabel_instrContext

	// Get_printf_instr returns the _printf_instr rule contexts.
	Get_printf_instr() IPrintf_instrContext

	// Get_return_instr returns the _return_instr rule contexts.
	Get_return_instr() IReturn_instrContext

	// Get_getHeap_instr returns the _getHeap_instr rule contexts.
	Get_getHeap_instr() IGetHeap_instrContext

	// Get_getStack_instr returns the _getStack_instr rule contexts.
	Get_getStack_instr() IGetStack_instrContext

	// Set_assign_stack sets the _assign_stack rule contexts.
	Set_assign_stack(IAssign_stackContext)

	// Set_assign_heap sets the _assign_heap rule contexts.
	Set_assign_heap(IAssign_heapContext)

	// Set_assign sets the _assign rule contexts.
	Set_assign(IAssignContext)

	// Set_if_instr sets the _if_instr rule contexts.
	Set_if_instr(IIf_instrContext)

	// Set_goto_instr sets the _goto_instr rule contexts.
	Set_goto_instr(IGoto_instrContext)

	// Set_label_instr sets the _label_instr rule contexts.
	Set_label_instr(ILabel_instrContext)

	// Set_printf_instr sets the _printf_instr rule contexts.
	Set_printf_instr(IPrintf_instrContext)

	// Set_return_instr sets the _return_instr rule contexts.
	Set_return_instr(IReturn_instrContext)

	// Set_getHeap_instr sets the _getHeap_instr rule contexts.
	Set_getHeap_instr(IGetHeap_instrContext)

	// Set_getStack_instr sets the _getStack_instr rule contexts.
	Set_getStack_instr(IGetStack_instrContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	instr           AbstractOptimizer.Instruction
	_assign_stack   IAssign_stackContext
	_assign_heap    IAssign_heapContext
	_assign         IAssignContext
	_if_instr       IIf_instrContext
	_goto_instr     IGoto_instrContext
	_label_instr    ILabel_instrContext
	_printf_instr   IPrintf_instrContext
	_return_instr   IReturn_instrContext
	_getHeap_instr  IGetHeap_instrContext
	_getStack_instr IGetStack_instrContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Get_assign_stack() IAssign_stackContext { return s._assign_stack }

func (s *InstructionContext) Get_assign_heap() IAssign_heapContext { return s._assign_heap }

func (s *InstructionContext) Get_assign() IAssignContext { return s._assign }

func (s *InstructionContext) Get_if_instr() IIf_instrContext { return s._if_instr }

func (s *InstructionContext) Get_goto_instr() IGoto_instrContext { return s._goto_instr }

func (s *InstructionContext) Get_label_instr() ILabel_instrContext { return s._label_instr }

func (s *InstructionContext) Get_printf_instr() IPrintf_instrContext { return s._printf_instr }

func (s *InstructionContext) Get_return_instr() IReturn_instrContext { return s._return_instr }

func (s *InstructionContext) Get_getHeap_instr() IGetHeap_instrContext { return s._getHeap_instr }

func (s *InstructionContext) Get_getStack_instr() IGetStack_instrContext { return s._getStack_instr }

func (s *InstructionContext) Set_assign_stack(v IAssign_stackContext) { s._assign_stack = v }

func (s *InstructionContext) Set_assign_heap(v IAssign_heapContext) { s._assign_heap = v }

func (s *InstructionContext) Set_assign(v IAssignContext) { s._assign = v }

func (s *InstructionContext) Set_if_instr(v IIf_instrContext) { s._if_instr = v }

func (s *InstructionContext) Set_goto_instr(v IGoto_instrContext) { s._goto_instr = v }

func (s *InstructionContext) Set_label_instr(v ILabel_instrContext) { s._label_instr = v }

func (s *InstructionContext) Set_printf_instr(v IPrintf_instrContext) { s._printf_instr = v }

func (s *InstructionContext) Set_return_instr(v IReturn_instrContext) { s._return_instr = v }

func (s *InstructionContext) Set_getHeap_instr(v IGetHeap_instrContext) { s._getHeap_instr = v }

func (s *InstructionContext) Set_getStack_instr(v IGetStack_instrContext) { s._getStack_instr = v }

func (s *InstructionContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *InstructionContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *InstructionContext) Assign_stack() IAssign_stackContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_stackContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_stackContext)
}

func (s *InstructionContext) Assign_heap() IAssign_heapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_heapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_heapContext)
}

func (s *InstructionContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *InstructionContext) If_instr() IIf_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_instrContext)
}

func (s *InstructionContext) Goto_instr() IGoto_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGoto_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGoto_instrContext)
}

func (s *InstructionContext) Label_instr() ILabel_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabel_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabel_instrContext)
}

func (s *InstructionContext) Printf_instr() IPrintf_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintf_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintf_instrContext)
}

func (s *InstructionContext) Return_instr() IReturn_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_instrContext)
}

func (s *InstructionContext) GetHeap_instr() IGetHeap_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetHeap_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetHeap_instrContext)
}

func (s *InstructionContext) GetStack_instr() IGetStack_instrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetStack_instrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetStack_instrContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *OptimizerParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, OptimizerParserRULE_instruction)

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

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)

			var _x = p.Assign_stack()

			localctx.(*InstructionContext)._assign_stack = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_assign_stack().GetInstr()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)

			var _x = p.Assign_heap()

			localctx.(*InstructionContext)._assign_heap = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_assign_heap().GetInstr()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)

			var _x = p.Assign()

			localctx.(*InstructionContext)._assign = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_assign().GetInstr()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)

			var _x = p.If_instr()

			localctx.(*InstructionContext)._if_instr = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_if_instr().GetInstr()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(137)

			var _x = p.Goto_instr()

			localctx.(*InstructionContext)._goto_instr = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_goto_instr().GetInstr()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)

			var _x = p.Label_instr()

			localctx.(*InstructionContext)._label_instr = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_label_instr().GetInstr()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(143)

			var _x = p.Printf_instr()

			localctx.(*InstructionContext)._printf_instr = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_printf_instr().GetInstr()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(146)

			var _x = p.Return_instr()

			localctx.(*InstructionContext)._return_instr = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_return_instr().GetInstr()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(149)

			var _x = p.GetHeap_instr()

			localctx.(*InstructionContext)._getHeap_instr = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_getHeap_instr().GetInstr()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(152)

			var _x = p.GetStack_instr()

			localctx.(*InstructionContext)._getStack_instr = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_getStack_instr().GetInstr()

	}

	return localctx
}

// IAssign_stackContext is an interface to support dynamic dispatch.
type IAssign_stackContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RSTACK returns the _RSTACK token.
	Get_RSTACK() antlr.Token

	// Get_EQUAL returns the _EQUAL token.
	Get_EQUAL() antlr.Token

	// Set_RSTACK sets the _RSTACK token.
	Set_RSTACK(antlr.Token)

	// Set_EQUAL sets the _EQUAL token.
	Set_EQUAL(antlr.Token)

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsAssign_stackContext differentiates from other interfaces.
	IsAssign_stackContext()
}

type Assign_stackContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       AbstractOptimizer.Instruction
	_RSTACK     antlr.Token
	_expr_valor IExpr_valorContext
	_EQUAL      antlr.Token
	_expression IExpressionContext
}

func NewEmptyAssign_stackContext() *Assign_stackContext {
	var p = new(Assign_stackContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_assign_stack
	return p
}

func (*Assign_stackContext) IsAssign_stackContext() {}

func NewAssign_stackContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stackContext {
	var p = new(Assign_stackContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_assign_stack

	return p
}

func (s *Assign_stackContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stackContext) Get_RSTACK() antlr.Token { return s._RSTACK }

func (s *Assign_stackContext) Get_EQUAL() antlr.Token { return s._EQUAL }

func (s *Assign_stackContext) Set_RSTACK(v antlr.Token) { s._RSTACK = v }

func (s *Assign_stackContext) Set_EQUAL(v antlr.Token) { s._EQUAL = v }

func (s *Assign_stackContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *Assign_stackContext) Get_expression() IExpressionContext { return s._expression }

func (s *Assign_stackContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *Assign_stackContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Assign_stackContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *Assign_stackContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *Assign_stackContext) RSTACK() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRSTACK, 0)
}

func (s *Assign_stackContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_BRACKET, 0)
}

func (s *Assign_stackContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *Assign_stackContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Assign_stackContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
}

func (s *Assign_stackContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *Assign_stackContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_BRACKET, 0)
}

func (s *Assign_stackContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(OptimizerParserEQUAL, 0)
}

func (s *Assign_stackContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assign_stackContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *Assign_stackContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stackContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_stackContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterAssign_stack(s)
	}
}

func (s *Assign_stackContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitAssign_stack(s)
	}
}

func (p *OptimizerParser) Assign_stack() (localctx IAssign_stackContext) {
	localctx = NewAssign_stackContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, OptimizerParserRULE_assign_stack)

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
		p.SetState(157)

		var _m = p.Match(OptimizerParserRSTACK)

		localctx.(*Assign_stackContext)._RSTACK = _m
	}
	{
		p.SetState(158)
		p.Match(OptimizerParserLEFT_BRACKET)
	}
	{
		p.SetState(159)
		p.Match(OptimizerParserLEFT_PAR)
	}
	{
		p.SetState(160)
		p.Data_type()
	}
	{
		p.SetState(161)
		p.Match(OptimizerParserRIGHT_PAR)
	}
	{
		p.SetState(162)

		var _x = p.Expr_valor()

		localctx.(*Assign_stackContext)._expr_valor = _x
	}
	{
		p.SetState(163)
		p.Match(OptimizerParserRIGHT_BRACKET)
	}
	{
		p.SetState(164)

		var _m = p.Match(OptimizerParserEQUAL)

		localctx.(*Assign_stackContext)._EQUAL = _m
	}
	{
		p.SetState(165)

		var _x = p.Expression()

		localctx.(*Assign_stackContext)._expression = _x
	}
	{
		p.SetState(166)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*Assign_stackContext).instr = Assignment.NewAssignHeapStack(localctx.(*Assign_stackContext).Get_expr_valor().GetP(), localctx.(*Assign_stackContext).Get_expression().GetP(), (func() int {
		if localctx.(*Assign_stackContext).Get_EQUAL() == nil {
			return 0
		} else {
			return localctx.(*Assign_stackContext).Get_EQUAL().GetLine()
		}
	}()), localctx.(*Assign_stackContext).Get_EQUAL().GetColumn(), (func() string {
		if localctx.(*Assign_stackContext).Get_RSTACK() == nil {
			return ""
		} else {
			return localctx.(*Assign_stackContext).Get_RSTACK().GetText()
		}
	}()))

	return localctx
}

// IAssign_heapContext is an interface to support dynamic dispatch.
type IAssign_heapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RHEAP returns the _RHEAP token.
	Get_RHEAP() antlr.Token

	// Get_EQUAL returns the _EQUAL token.
	Get_EQUAL() antlr.Token

	// Set_RHEAP sets the _RHEAP token.
	Set_RHEAP(antlr.Token)

	// Set_EQUAL sets the _EQUAL token.
	Set_EQUAL(antlr.Token)

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsAssign_heapContext differentiates from other interfaces.
	IsAssign_heapContext()
}

type Assign_heapContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       AbstractOptimizer.Instruction
	_RHEAP      antlr.Token
	_expr_valor IExpr_valorContext
	_EQUAL      antlr.Token
	_expression IExpressionContext
}

func NewEmptyAssign_heapContext() *Assign_heapContext {
	var p = new(Assign_heapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_assign_heap
	return p
}

func (*Assign_heapContext) IsAssign_heapContext() {}

func NewAssign_heapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_heapContext {
	var p = new(Assign_heapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_assign_heap

	return p
}

func (s *Assign_heapContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_heapContext) Get_RHEAP() antlr.Token { return s._RHEAP }

func (s *Assign_heapContext) Get_EQUAL() antlr.Token { return s._EQUAL }

func (s *Assign_heapContext) Set_RHEAP(v antlr.Token) { s._RHEAP = v }

func (s *Assign_heapContext) Set_EQUAL(v antlr.Token) { s._EQUAL = v }

func (s *Assign_heapContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *Assign_heapContext) Get_expression() IExpressionContext { return s._expression }

func (s *Assign_heapContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *Assign_heapContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Assign_heapContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *Assign_heapContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *Assign_heapContext) RHEAP() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRHEAP, 0)
}

func (s *Assign_heapContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_BRACKET, 0)
}

func (s *Assign_heapContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *Assign_heapContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Assign_heapContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
}

func (s *Assign_heapContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *Assign_heapContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_BRACKET, 0)
}

func (s *Assign_heapContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(OptimizerParserEQUAL, 0)
}

func (s *Assign_heapContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assign_heapContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *Assign_heapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_heapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_heapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterAssign_heap(s)
	}
}

func (s *Assign_heapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitAssign_heap(s)
	}
}

func (p *OptimizerParser) Assign_heap() (localctx IAssign_heapContext) {
	localctx = NewAssign_heapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, OptimizerParserRULE_assign_heap)

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
		p.SetState(169)

		var _m = p.Match(OptimizerParserRHEAP)

		localctx.(*Assign_heapContext)._RHEAP = _m
	}
	{
		p.SetState(170)
		p.Match(OptimizerParserLEFT_BRACKET)
	}
	{
		p.SetState(171)
		p.Match(OptimizerParserLEFT_PAR)
	}
	{
		p.SetState(172)
		p.Data_type()
	}
	{
		p.SetState(173)
		p.Match(OptimizerParserRIGHT_PAR)
	}
	{
		p.SetState(174)

		var _x = p.Expr_valor()

		localctx.(*Assign_heapContext)._expr_valor = _x
	}
	{
		p.SetState(175)
		p.Match(OptimizerParserRIGHT_BRACKET)
	}
	{
		p.SetState(176)

		var _m = p.Match(OptimizerParserEQUAL)

		localctx.(*Assign_heapContext)._EQUAL = _m
	}
	{
		p.SetState(177)

		var _x = p.Expression()

		localctx.(*Assign_heapContext)._expression = _x
	}
	{
		p.SetState(178)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*Assign_heapContext).instr = Assignment.NewAssignHeapStack(localctx.(*Assign_heapContext).Get_expr_valor().GetP(), localctx.(*Assign_heapContext).Get_expression().GetP(), (func() int {
		if localctx.(*Assign_heapContext).Get_EQUAL() == nil {
			return 0
		} else {
			return localctx.(*Assign_heapContext).Get_EQUAL().GetLine()
		}
	}()), localctx.(*Assign_heapContext).Get_EQUAL().GetColumn(), (func() string {
		if localctx.(*Assign_heapContext).Get_RHEAP() == nil {
			return ""
		} else {
			return localctx.(*Assign_heapContext).Get_RHEAP().GetText()
		}
	}()))

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_EQUAL returns the _EQUAL token.
	Get_EQUAL() antlr.Token

	// Set_EQUAL sets the _EQUAL token.
	Set_EQUAL(antlr.Token)

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       AbstractOptimizer.Instruction
	_expr_valor IExpr_valorContext
	_EQUAL      antlr.Token
	_expression IExpressionContext
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) Get_EQUAL() antlr.Token { return s._EQUAL }

func (s *AssignContext) Set_EQUAL(v antlr.Token) { s._EQUAL = v }

func (s *AssignContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *AssignContext) Get_expression() IExpressionContext { return s._expression }

func (s *AssignContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *AssignContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AssignContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *AssignContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *AssignContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *AssignContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(OptimizerParserEQUAL, 0)
}

func (s *AssignContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *OptimizerParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, OptimizerParserRULE_assign)

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
		p.SetState(181)

		var _x = p.Expr_valor()

		localctx.(*AssignContext)._expr_valor = _x
	}
	{
		p.SetState(182)

		var _m = p.Match(OptimizerParserEQUAL)

		localctx.(*AssignContext)._EQUAL = _m
	}
	{
		p.SetState(183)

		var _x = p.Expression()

		localctx.(*AssignContext)._expression = _x
	}
	{
		p.SetState(184)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*AssignContext).instr = Assignment.NewAssign(localctx.(*AssignContext).Get_expr_valor().GetP(), localctx.(*AssignContext).Get_expression().GetP(), (func() int {
		if localctx.(*AssignContext).Get_EQUAL() == nil {
			return 0
		} else {
			return localctx.(*AssignContext).Get_EQUAL().GetLine()
		}
	}()), localctx.(*AssignContext).Get_EQUAL().GetColumn())

	return localctx
}

// IIf_instrContext is an interface to support dynamic dispatch.
type IIf_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RIF returns the _RIF token.
	Get_RIF() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_RIF sets the _RIF token.
	Set_RIF(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsIf_instrContext differentiates from other interfaces.
	IsIf_instrContext()
}

type If_instrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       AbstractOptimizer.Instruction
	_RIF        antlr.Token
	_expression IExpressionContext
	_ID         antlr.Token
}

func NewEmptyIf_instrContext() *If_instrContext {
	var p = new(If_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_if_instr
	return p
}

func (*If_instrContext) IsIf_instrContext() {}

func NewIf_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_instrContext {
	var p = new(If_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_if_instr

	return p
}

func (s *If_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *If_instrContext) Get_RIF() antlr.Token { return s._RIF }

func (s *If_instrContext) Get_ID() antlr.Token { return s._ID }

func (s *If_instrContext) Set_RIF(v antlr.Token) { s._RIF = v }

func (s *If_instrContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *If_instrContext) Get_expression() IExpressionContext { return s._expression }

func (s *If_instrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *If_instrContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *If_instrContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *If_instrContext) RIF() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIF, 0)
}

func (s *If_instrContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *If_instrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_instrContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
}

func (s *If_instrContext) RGOTO() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRGOTO, 0)
}

func (s *If_instrContext) ID() antlr.TerminalNode {
	return s.GetToken(OptimizerParserID, 0)
}

func (s *If_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *If_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterIf_instr(s)
	}
}

func (s *If_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitIf_instr(s)
	}
}

func (p *OptimizerParser) If_instr() (localctx IIf_instrContext) {
	localctx = NewIf_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, OptimizerParserRULE_if_instr)

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
		p.SetState(187)

		var _m = p.Match(OptimizerParserRIF)

		localctx.(*If_instrContext)._RIF = _m
	}
	{
		p.SetState(188)
		p.Match(OptimizerParserLEFT_PAR)
	}
	{
		p.SetState(189)

		var _x = p.Expression()

		localctx.(*If_instrContext)._expression = _x
	}
	{
		p.SetState(190)
		p.Match(OptimizerParserRIGHT_PAR)
	}
	{
		p.SetState(191)
		p.Match(OptimizerParserRGOTO)
	}
	{
		p.SetState(192)

		var _m = p.Match(OptimizerParserID)

		localctx.(*If_instrContext)._ID = _m
	}
	{
		p.SetState(193)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*If_instrContext).instr = Control.NewIf(localctx.(*If_instrContext).Get_expression().GetP(), (func() string {
		if localctx.(*If_instrContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*If_instrContext).Get_ID().GetText()
		}
	}()), (func() int {
		if localctx.(*If_instrContext).Get_RIF() == nil {
			return 0
		} else {
			return localctx.(*If_instrContext).Get_RIF().GetLine()
		}
	}()), localctx.(*If_instrContext).Get_ID().GetColumn())

	return localctx
}

// IGoto_instrContext is an interface to support dynamic dispatch.
type IGoto_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsGoto_instrContext differentiates from other interfaces.
	IsGoto_instrContext()
}

type Goto_instrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	instr  AbstractOptimizer.Instruction
	_ID    antlr.Token
}

func NewEmptyGoto_instrContext() *Goto_instrContext {
	var p = new(Goto_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_goto_instr
	return p
}

func (*Goto_instrContext) IsGoto_instrContext() {}

func NewGoto_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Goto_instrContext {
	var p = new(Goto_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_goto_instr

	return p
}

func (s *Goto_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *Goto_instrContext) Get_ID() antlr.Token { return s._ID }

func (s *Goto_instrContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Goto_instrContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *Goto_instrContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *Goto_instrContext) RGOTO() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRGOTO, 0)
}

func (s *Goto_instrContext) ID() antlr.TerminalNode {
	return s.GetToken(OptimizerParserID, 0)
}

func (s *Goto_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *Goto_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Goto_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Goto_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterGoto_instr(s)
	}
}

func (s *Goto_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitGoto_instr(s)
	}
}

func (p *OptimizerParser) Goto_instr() (localctx IGoto_instrContext) {
	localctx = NewGoto_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, OptimizerParserRULE_goto_instr)

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
		p.SetState(196)
		p.Match(OptimizerParserRGOTO)
	}
	{
		p.SetState(197)

		var _m = p.Match(OptimizerParserID)

		localctx.(*Goto_instrContext)._ID = _m
	}
	{
		p.SetState(198)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*Goto_instrContext).instr = Control.NewGoTo((func() string {
		if localctx.(*Goto_instrContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Goto_instrContext).Get_ID().GetText()
		}
	}()), (func() int {
		if localctx.(*Goto_instrContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Goto_instrContext).Get_ID().GetLine()
		}
	}()), localctx.(*Goto_instrContext).Get_ID().GetColumn())

	return localctx
}

// ILabel_instrContext is an interface to support dynamic dispatch.
type ILabel_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsLabel_instrContext differentiates from other interfaces.
	IsLabel_instrContext()
}

type Label_instrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	instr  AbstractOptimizer.Instruction
	_ID    antlr.Token
}

func NewEmptyLabel_instrContext() *Label_instrContext {
	var p = new(Label_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_label_instr
	return p
}

func (*Label_instrContext) IsLabel_instrContext() {}

func NewLabel_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Label_instrContext {
	var p = new(Label_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_label_instr

	return p
}

func (s *Label_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *Label_instrContext) Get_ID() antlr.Token { return s._ID }

func (s *Label_instrContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Label_instrContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *Label_instrContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *Label_instrContext) ID() antlr.TerminalNode {
	return s.GetToken(OptimizerParserID, 0)
}

func (s *Label_instrContext) COLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserCOLON, 0)
}

func (s *Label_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Label_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Label_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterLabel_instr(s)
	}
}

func (s *Label_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitLabel_instr(s)
	}
}

func (p *OptimizerParser) Label_instr() (localctx ILabel_instrContext) {
	localctx = NewLabel_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, OptimizerParserRULE_label_instr)

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
		p.SetState(201)

		var _m = p.Match(OptimizerParserID)

		localctx.(*Label_instrContext)._ID = _m
	}
	{
		p.SetState(202)
		p.Match(OptimizerParserCOLON)
	}

	localctx.(*Label_instrContext).instr = Control.NewLabel((func() string {
		if localctx.(*Label_instrContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*Label_instrContext).Get_ID().GetText()
		}
	}()), (func() int {
		if localctx.(*Label_instrContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*Label_instrContext).Get_ID().GetLine()
		}
	}()), localctx.(*Label_instrContext).Get_ID().GetColumn())

	return localctx
}

// IPrintf_instrContext is an interface to support dynamic dispatch.
type IPrintf_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Get_convert returns the _convert rule contexts.
	Get_convert() IConvertContext

	// Get_expr_print returns the _expr_print rule contexts.
	Get_expr_print() IExpr_printContext

	// Set_convert sets the _convert rule contexts.
	Set_convert(IConvertContext)

	// Set_expr_print sets the _expr_print rule contexts.
	Set_expr_print(IExpr_printContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsPrintf_instrContext differentiates from other interfaces.
	IsPrintf_instrContext()
}

type Printf_instrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       AbstractOptimizer.Instruction
	_STRING     antlr.Token
	_convert    IConvertContext
	_expr_print IExpr_printContext
}

func NewEmptyPrintf_instrContext() *Printf_instrContext {
	var p = new(Printf_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_printf_instr
	return p
}

func (*Printf_instrContext) IsPrintf_instrContext() {}

func NewPrintf_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Printf_instrContext {
	var p = new(Printf_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_printf_instr

	return p
}

func (s *Printf_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *Printf_instrContext) Get_STRING() antlr.Token { return s._STRING }

func (s *Printf_instrContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *Printf_instrContext) Get_convert() IConvertContext { return s._convert }

func (s *Printf_instrContext) Get_expr_print() IExpr_printContext { return s._expr_print }

func (s *Printf_instrContext) Set_convert(v IConvertContext) { s._convert = v }

func (s *Printf_instrContext) Set_expr_print(v IExpr_printContext) { s._expr_print = v }

func (s *Printf_instrContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *Printf_instrContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *Printf_instrContext) RPRINTF() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRPRINTF, 0)
}

func (s *Printf_instrContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *Printf_instrContext) STRING() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSTRING, 0)
}

func (s *Printf_instrContext) COMMA() antlr.TerminalNode {
	return s.GetToken(OptimizerParserCOMMA, 0)
}

func (s *Printf_instrContext) Convert() IConvertContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConvertContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConvertContext)
}

func (s *Printf_instrContext) Expr_print() IExpr_printContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_printContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_printContext)
}

func (s *Printf_instrContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
}

func (s *Printf_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *Printf_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Printf_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Printf_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterPrintf_instr(s)
	}
}

func (s *Printf_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitPrintf_instr(s)
	}
}

func (p *OptimizerParser) Printf_instr() (localctx IPrintf_instrContext) {
	localctx = NewPrintf_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, OptimizerParserRULE_printf_instr)

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
		p.SetState(205)
		p.Match(OptimizerParserRPRINTF)
	}
	{
		p.SetState(206)
		p.Match(OptimizerParserLEFT_PAR)
	}
	{
		p.SetState(207)

		var _m = p.Match(OptimizerParserSTRING)

		localctx.(*Printf_instrContext)._STRING = _m
	}
	{
		p.SetState(208)
		p.Match(OptimizerParserCOMMA)
	}
	{
		p.SetState(209)

		var _x = p.Convert()

		localctx.(*Printf_instrContext)._convert = _x
	}
	{
		p.SetState(210)

		var _x = p.Expr_print()

		localctx.(*Printf_instrContext)._expr_print = _x
	}
	{
		p.SetState(211)
		p.Match(OptimizerParserRIGHT_PAR)
	}
	{
		p.SetState(212)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*Printf_instrContext).instr = Print.NewPrintf((func() string {
		if localctx.(*Printf_instrContext).Get_STRING() == nil {
			return ""
		} else {
			return localctx.(*Printf_instrContext).Get_STRING().GetText()
		}
	}()), localctx.(*Printf_instrContext).Get_convert().GetDtype(), localctx.(*Printf_instrContext).Get_expr_print().GetId(), (func() int {
		if localctx.(*Printf_instrContext).Get_STRING() == nil {
			return 0
		} else {
			return localctx.(*Printf_instrContext).Get_STRING().GetLine()
		}
	}()), localctx.(*Printf_instrContext).Get_STRING().GetColumn())

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
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsReturn_instrContext differentiates from other interfaces.
	IsReturn_instrContext()
}

type Return_instrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       AbstractOptimizer.Instruction
	_RRETURN    antlr.Token
	_expression IExpressionContext
}

func NewEmptyReturn_instrContext() *Return_instrContext {
	var p = new(Return_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_return_instr
	return p
}

func (*Return_instrContext) IsReturn_instrContext() {}

func NewReturn_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_instrContext {
	var p = new(Return_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_return_instr

	return p
}

func (s *Return_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_instrContext) Get_RRETURN() antlr.Token { return s._RRETURN }

func (s *Return_instrContext) Set_RRETURN(v antlr.Token) { s._RRETURN = v }

func (s *Return_instrContext) Get_expression() IExpressionContext { return s._expression }

func (s *Return_instrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Return_instrContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *Return_instrContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *Return_instrContext) RRETURN() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRRETURN, 0)
}

func (s *Return_instrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Return_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *Return_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterReturn_instr(s)
	}
}

func (s *Return_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitReturn_instr(s)
	}
}

func (p *OptimizerParser) Return_instr() (localctx IReturn_instrContext) {
	localctx = NewReturn_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, OptimizerParserRULE_return_instr)

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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(215)

			var _m = p.Match(OptimizerParserRRETURN)

			localctx.(*Return_instrContext)._RRETURN = _m
		}
		{
			p.SetState(216)

			var _x = p.Expression()

			localctx.(*Return_instrContext)._expression = _x
		}
		{
			p.SetState(217)
			p.Match(OptimizerParserSEMICOLON)
		}

		localctx.(*Return_instrContext).instr = Control.NewReturn(localctx.(*Return_instrContext).Get_expression().GetP(), (func() int {
			if localctx.(*Return_instrContext).Get_RRETURN() == nil {
				return 0
			} else {
				return localctx.(*Return_instrContext).Get_RRETURN().GetLine()
			}
		}()), localctx.(*Return_instrContext).Get_RRETURN().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)

			var _m = p.Match(OptimizerParserRRETURN)

			localctx.(*Return_instrContext)._RRETURN = _m
		}
		{
			p.SetState(221)
			p.Match(OptimizerParserSEMICOLON)
		}

		localctx.(*Return_instrContext).instr = Control.NewReturn(nil, (func() int {
			if localctx.(*Return_instrContext).Get_RRETURN() == nil {
				return 0
			} else {
				return localctx.(*Return_instrContext).Get_RRETURN().GetLine()
			}
		}()), localctx.(*Return_instrContext).Get_RRETURN().GetColumn())

	}

	return localctx
}

// IGetHeap_instrContext is an interface to support dynamic dispatch.
type IGetHeap_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_EQUAL returns the _EQUAL token.
	Get_EQUAL() antlr.Token

	// Get_RHEAP returns the _RHEAP token.
	Get_RHEAP() antlr.Token

	// Set_EQUAL sets the _EQUAL token.
	Set_EQUAL(antlr.Token)

	// Set_RHEAP sets the _RHEAP token.
	Set_RHEAP(antlr.Token)

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsGetHeap_instrContext differentiates from other interfaces.
	IsGetHeap_instrContext()
}

type GetHeap_instrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       AbstractOptimizer.Instruction
	_expr_valor IExpr_valorContext
	_EQUAL      antlr.Token
	_RHEAP      antlr.Token
	_expression IExpressionContext
}

func NewEmptyGetHeap_instrContext() *GetHeap_instrContext {
	var p = new(GetHeap_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_getHeap_instr
	return p
}

func (*GetHeap_instrContext) IsGetHeap_instrContext() {}

func NewGetHeap_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetHeap_instrContext {
	var p = new(GetHeap_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_getHeap_instr

	return p
}

func (s *GetHeap_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *GetHeap_instrContext) Get_EQUAL() antlr.Token { return s._EQUAL }

func (s *GetHeap_instrContext) Get_RHEAP() antlr.Token { return s._RHEAP }

func (s *GetHeap_instrContext) Set_EQUAL(v antlr.Token) { s._EQUAL = v }

func (s *GetHeap_instrContext) Set_RHEAP(v antlr.Token) { s._RHEAP = v }

func (s *GetHeap_instrContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *GetHeap_instrContext) Get_expression() IExpressionContext { return s._expression }

func (s *GetHeap_instrContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *GetHeap_instrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *GetHeap_instrContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *GetHeap_instrContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *GetHeap_instrContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *GetHeap_instrContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(OptimizerParserEQUAL, 0)
}

func (s *GetHeap_instrContext) RHEAP() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRHEAP, 0)
}

func (s *GetHeap_instrContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_BRACKET, 0)
}

func (s *GetHeap_instrContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *GetHeap_instrContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *GetHeap_instrContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
}

func (s *GetHeap_instrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GetHeap_instrContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_BRACKET, 0)
}

func (s *GetHeap_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *GetHeap_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetHeap_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetHeap_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterGetHeap_instr(s)
	}
}

func (s *GetHeap_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitGetHeap_instr(s)
	}
}

func (p *OptimizerParser) GetHeap_instr() (localctx IGetHeap_instrContext) {
	localctx = NewGetHeap_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, OptimizerParserRULE_getHeap_instr)

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
		p.SetState(225)

		var _x = p.Expr_valor()

		localctx.(*GetHeap_instrContext)._expr_valor = _x
	}
	{
		p.SetState(226)

		var _m = p.Match(OptimizerParserEQUAL)

		localctx.(*GetHeap_instrContext)._EQUAL = _m
	}
	{
		p.SetState(227)

		var _m = p.Match(OptimizerParserRHEAP)

		localctx.(*GetHeap_instrContext)._RHEAP = _m
	}
	{
		p.SetState(228)
		p.Match(OptimizerParserLEFT_BRACKET)
	}
	{
		p.SetState(229)
		p.Match(OptimizerParserLEFT_PAR)
	}
	{
		p.SetState(230)
		p.Data_type()
	}
	{
		p.SetState(231)
		p.Match(OptimizerParserRIGHT_PAR)
	}
	{
		p.SetState(232)

		var _x = p.Expression()

		localctx.(*GetHeap_instrContext)._expression = _x
	}
	{
		p.SetState(233)
		p.Match(OptimizerParserRIGHT_BRACKET)
	}
	{
		p.SetState(234)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*GetHeap_instrContext).instr = Assignment.NewGetHeapStack(localctx.(*GetHeap_instrContext).Get_expr_valor().GetP(), localctx.(*GetHeap_instrContext).Get_expression().GetP(), (func() int {
		if localctx.(*GetHeap_instrContext).Get_EQUAL() == nil {
			return 0
		} else {
			return localctx.(*GetHeap_instrContext).Get_EQUAL().GetLine()
		}
	}()), localctx.(*GetHeap_instrContext).Get_EQUAL().GetColumn(), (func() string {
		if localctx.(*GetHeap_instrContext).Get_RHEAP() == nil {
			return ""
		} else {
			return localctx.(*GetHeap_instrContext).Get_RHEAP().GetText()
		}
	}()))

	return localctx
}

// IGetStack_instrContext is an interface to support dynamic dispatch.
type IGetStack_instrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_EQUAL returns the _EQUAL token.
	Get_EQUAL() antlr.Token

	// Get_RSTACK returns the _RSTACK token.
	Get_RSTACK() antlr.Token

	// Set_EQUAL sets the _EQUAL token.
	Set_EQUAL(antlr.Token)

	// Set_RSTACK sets the _RSTACK token.
	Set_RSTACK(antlr.Token)

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetInstr returns the instr attribute.
	GetInstr() AbstractOptimizer.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(AbstractOptimizer.Instruction)

	// IsGetStack_instrContext differentiates from other interfaces.
	IsGetStack_instrContext()
}

type GetStack_instrContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	instr       AbstractOptimizer.Instruction
	_expr_valor IExpr_valorContext
	_EQUAL      antlr.Token
	_RSTACK     antlr.Token
	_expression IExpressionContext
}

func NewEmptyGetStack_instrContext() *GetStack_instrContext {
	var p = new(GetStack_instrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_getStack_instr
	return p
}

func (*GetStack_instrContext) IsGetStack_instrContext() {}

func NewGetStack_instrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetStack_instrContext {
	var p = new(GetStack_instrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_getStack_instr

	return p
}

func (s *GetStack_instrContext) GetParser() antlr.Parser { return s.parser }

func (s *GetStack_instrContext) Get_EQUAL() antlr.Token { return s._EQUAL }

func (s *GetStack_instrContext) Get_RSTACK() antlr.Token { return s._RSTACK }

func (s *GetStack_instrContext) Set_EQUAL(v antlr.Token) { s._EQUAL = v }

func (s *GetStack_instrContext) Set_RSTACK(v antlr.Token) { s._RSTACK = v }

func (s *GetStack_instrContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *GetStack_instrContext) Get_expression() IExpressionContext { return s._expression }

func (s *GetStack_instrContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *GetStack_instrContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *GetStack_instrContext) GetInstr() AbstractOptimizer.Instruction { return s.instr }

func (s *GetStack_instrContext) SetInstr(v AbstractOptimizer.Instruction) { s.instr = v }

func (s *GetStack_instrContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *GetStack_instrContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(OptimizerParserEQUAL, 0)
}

func (s *GetStack_instrContext) RSTACK() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRSTACK, 0)
}

func (s *GetStack_instrContext) LEFT_BRACKET() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_BRACKET, 0)
}

func (s *GetStack_instrContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *GetStack_instrContext) Data_type() IData_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *GetStack_instrContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
}

func (s *GetStack_instrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GetStack_instrContext) RIGHT_BRACKET() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_BRACKET, 0)
}

func (s *GetStack_instrContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSEMICOLON, 0)
}

func (s *GetStack_instrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetStack_instrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetStack_instrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterGetStack_instr(s)
	}
}

func (s *GetStack_instrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitGetStack_instr(s)
	}
}

func (p *OptimizerParser) GetStack_instr() (localctx IGetStack_instrContext) {
	localctx = NewGetStack_instrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, OptimizerParserRULE_getStack_instr)

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
		p.SetState(237)

		var _x = p.Expr_valor()

		localctx.(*GetStack_instrContext)._expr_valor = _x
	}
	{
		p.SetState(238)

		var _m = p.Match(OptimizerParserEQUAL)

		localctx.(*GetStack_instrContext)._EQUAL = _m
	}
	{
		p.SetState(239)

		var _m = p.Match(OptimizerParserRSTACK)

		localctx.(*GetStack_instrContext)._RSTACK = _m
	}
	{
		p.SetState(240)
		p.Match(OptimizerParserLEFT_BRACKET)
	}
	{
		p.SetState(241)
		p.Match(OptimizerParserLEFT_PAR)
	}
	{
		p.SetState(242)
		p.Data_type()
	}
	{
		p.SetState(243)
		p.Match(OptimizerParserRIGHT_PAR)
	}
	{
		p.SetState(244)

		var _x = p.Expression()

		localctx.(*GetStack_instrContext)._expression = _x
	}
	{
		p.SetState(245)
		p.Match(OptimizerParserRIGHT_BRACKET)
	}
	{
		p.SetState(246)
		p.Match(OptimizerParserSEMICOLON)
	}

	localctx.(*GetStack_instrContext).instr = Assignment.NewGetHeapStack(localctx.(*GetStack_instrContext).Get_expr_valor().GetP(), localctx.(*GetStack_instrContext).Get_expression().GetP(), (func() int {
		if localctx.(*GetStack_instrContext).Get_EQUAL() == nil {
			return 0
		} else {
			return localctx.(*GetStack_instrContext).Get_EQUAL().GetLine()
		}
	}()), localctx.(*GetStack_instrContext).Get_EQUAL().GetColumn(), (func() string {
		if localctx.(*GetStack_instrContext).Get_RSTACK() == nil {
			return ""
		} else {
			return localctx.(*GetStack_instrContext).Get_RSTACK().GetText()
		}
	}()))

	return localctx
}

// IExpr_printContext is an interface to support dynamic dispatch.
type IExpr_printContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// GetId returns the id attribute.
	GetId() AbstractOptimizer.Expression

	// SetId sets the id attribute.
	SetId(AbstractOptimizer.Expression)

	// IsExpr_printContext differentiates from other interfaces.
	IsExpr_printContext()
}

type Expr_printContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	id          AbstractOptimizer.Expression
	_expr_valor IExpr_valorContext
}

func NewEmptyExpr_printContext() *Expr_printContext {
	var p = new(Expr_printContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_expr_print
	return p
}

func (*Expr_printContext) IsExpr_printContext() {}

func NewExpr_printContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_printContext {
	var p = new(Expr_printContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_expr_print

	return p
}

func (s *Expr_printContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_printContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *Expr_printContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *Expr_printContext) GetId() AbstractOptimizer.Expression { return s.id }

func (s *Expr_printContext) SetId(v AbstractOptimizer.Expression) { s.id = v }

func (s *Expr_printContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *Expr_printContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_printContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_printContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterExpr_print(s)
	}
}

func (s *Expr_printContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitExpr_print(s)
	}
}

func (p *OptimizerParser) Expr_print() (localctx IExpr_printContext) {
	localctx = NewExpr_printContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, OptimizerParserRULE_expr_print)

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
		p.SetState(249)

		var _x = p.Expr_valor()

		localctx.(*Expr_printContext)._expr_valor = _x
	}

	localctx.(*Expr_printContext).id = localctx.(*Expr_printContext).Get_expr_valor().GetP()

	return localctx
}

// IConvertContext is an interface to support dynamic dispatch.
type IConvertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RINT returns the _RINT token.
	Get_RINT() antlr.Token

	// Get_RFLOAT returns the _RFLOAT token.
	Get_RFLOAT() antlr.Token

	// Get_RCHAR returns the _RCHAR token.
	Get_RCHAR() antlr.Token

	// Set_RINT sets the _RINT token.
	Set_RINT(antlr.Token)

	// Set_RFLOAT sets the _RFLOAT token.
	Set_RFLOAT(antlr.Token)

	// Set_RCHAR sets the _RCHAR token.
	Set_RCHAR(antlr.Token)

	// GetDtype returns the dtype attribute.
	GetDtype() string

	// SetDtype sets the dtype attribute.
	SetDtype(string)

	// IsConvertContext differentiates from other interfaces.
	IsConvertContext()
}

type ConvertContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	dtype   string
	_RINT   antlr.Token
	_RFLOAT antlr.Token
	_RCHAR  antlr.Token
}

func NewEmptyConvertContext() *ConvertContext {
	var p = new(ConvertContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_convert
	return p
}

func (*ConvertContext) IsConvertContext() {}

func NewConvertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConvertContext {
	var p = new(ConvertContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_convert

	return p
}

func (s *ConvertContext) GetParser() antlr.Parser { return s.parser }

func (s *ConvertContext) Get_RINT() antlr.Token { return s._RINT }

func (s *ConvertContext) Get_RFLOAT() antlr.Token { return s._RFLOAT }

func (s *ConvertContext) Get_RCHAR() antlr.Token { return s._RCHAR }

func (s *ConvertContext) Set_RINT(v antlr.Token) { s._RINT = v }

func (s *ConvertContext) Set_RFLOAT(v antlr.Token) { s._RFLOAT = v }

func (s *ConvertContext) Set_RCHAR(v antlr.Token) { s._RCHAR = v }

func (s *ConvertContext) GetDtype() string { return s.dtype }

func (s *ConvertContext) SetDtype(v string) { s.dtype = v }

func (s *ConvertContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *ConvertContext) RINT() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRINT, 0)
}

func (s *ConvertContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
}

func (s *ConvertContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRFLOAT, 0)
}

func (s *ConvertContext) RCHAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRCHAR, 0)
}

func (s *ConvertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConvertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterConvert(s)
	}
}

func (s *ConvertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitConvert(s)
	}
}

func (p *OptimizerParser) Convert() (localctx IConvertContext) {
	localctx = NewConvertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, OptimizerParserRULE_convert)

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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.Match(OptimizerParserLEFT_PAR)
		}
		{
			p.SetState(253)

			var _m = p.Match(OptimizerParserRINT)

			localctx.(*ConvertContext)._RINT = _m
		}
		{
			p.SetState(254)
			p.Match(OptimizerParserRIGHT_PAR)
		}

		localctx.(*ConvertContext).dtype = (func() string {
			if localctx.(*ConvertContext).Get_RINT() == nil {
				return ""
			} else {
				return localctx.(*ConvertContext).Get_RINT().GetText()
			}
		}())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.Match(OptimizerParserLEFT_PAR)
		}
		{
			p.SetState(257)

			var _m = p.Match(OptimizerParserRFLOAT)

			localctx.(*ConvertContext)._RFLOAT = _m
		}
		{
			p.SetState(258)
			p.Match(OptimizerParserRIGHT_PAR)
		}

		localctx.(*ConvertContext).dtype = (func() string {
			if localctx.(*ConvertContext).Get_RFLOAT() == nil {
				return ""
			} else {
				return localctx.(*ConvertContext).Get_RFLOAT().GetText()
			}
		}())

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.Match(OptimizerParserLEFT_PAR)
		}
		{
			p.SetState(261)

			var _m = p.Match(OptimizerParserRCHAR)

			localctx.(*ConvertContext)._RCHAR = _m
		}
		{
			p.SetState(262)
			p.Match(OptimizerParserRIGHT_PAR)
		}

		localctx.(*ConvertContext).dtype = (func() string {
			if localctx.(*ConvertContext).Get_RCHAR() == nil {
				return ""
			} else {
				return localctx.(*ConvertContext).Get_RCHAR().GetText()
			}
		}())

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

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

	// GetP returns the p attribute.
	GetP() AbstractOptimizer.Expression

	// SetP sets the p attribute.
	SetP(AbstractOptimizer.Expression)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         AbstractOptimizer.Expression
	_expr_rel IExpr_relContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Get_expr_rel() IExpr_relContext { return s._expr_rel }

func (s *ExpressionContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *ExpressionContext) GetP() AbstractOptimizer.Expression { return s.p }

func (s *ExpressionContext) SetP(v AbstractOptimizer.Expression) { s.p = v }

func (s *ExpressionContext) Expr_rel() IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OptimizerParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, OptimizerParserRULE_expression)

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
		p.SetState(266)

		var _x = p.expr_rel(0)

		localctx.(*ExpressionContext)._expr_rel = _x
	}
	localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_rel().GetP()

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
	GetP() AbstractOptimizer.Expression

	// SetP sets the p attribute.
	SetP(AbstractOptimizer.Expression)

	// IsExpr_relContext differentiates from other interfaces.
	IsExpr_relContext()
}

type Expr_relContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          AbstractOptimizer.Expression
	opLeft     IExpr_relContext
	_expr_arit IExpr_aritContext
	op         antlr.Token
	opRight    IExpr_relContext
}

func NewEmptyExpr_relContext() *Expr_relContext {
	var p = new(Expr_relContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_expr_rel
	return p
}

func (*Expr_relContext) IsExpr_relContext() {}

func NewExpr_relContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_relContext {
	var p = new(Expr_relContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_expr_rel

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

func (s *Expr_relContext) GetP() AbstractOptimizer.Expression { return s.p }

func (s *Expr_relContext) SetP(v AbstractOptimizer.Expression) { s.p = v }

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
	return s.GetToken(OptimizerParserGREATER_THAN, 0)
}

func (s *Expr_relContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLESS_THAN, 0)
}

func (s *Expr_relContext) GREATER_EQUALTHAN() antlr.TerminalNode {
	return s.GetToken(OptimizerParserGREATER_EQUALTHAN, 0)
}

func (s *Expr_relContext) LESS_EQUEALTHAN() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLESS_EQUEALTHAN, 0)
}

func (s *Expr_relContext) EQUEAL_EQUAL() antlr.TerminalNode {
	return s.GetToken(OptimizerParserEQUEAL_EQUAL, 0)
}

func (s *Expr_relContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(OptimizerParserNOTEQUAL, 0)
}

func (s *Expr_relContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_relContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_relContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterExpr_rel(s)
	}
}

func (s *Expr_relContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitExpr_rel(s)
	}
}

func (p *OptimizerParser) Expr_rel() (localctx IExpr_relContext) {
	return p.expr_rel(0)
}

func (p *OptimizerParser) expr_rel(_p int) (localctx IExpr_relContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_relContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_relContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, OptimizerParserRULE_expr_rel, _p)
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
		p.SetState(270)

		var _x = p.expr_arit(0)

		localctx.(*Expr_relContext)._expr_arit = _x
	}

	localctx.(*Expr_relContext).p = localctx.(*Expr_relContext).Get_expr_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_relContext(p, _parentctx, _parentState)
			localctx.(*Expr_relContext).opLeft = _prevctx
			p.PushNewRecursionContext(localctx, _startState, OptimizerParserRULE_expr_rel)
			p.SetState(273)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(274)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expr_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(OptimizerParserEQUEAL_EQUAL-36))|(1<<(OptimizerParserNOTEQUAL-36))|(1<<(OptimizerParserGREATER_THAN-36))|(1<<(OptimizerParserLESS_THAN-36))|(1<<(OptimizerParserGREATER_EQUALTHAN-36))|(1<<(OptimizerParserLESS_EQUEALTHAN-36)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expr_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(275)

				var _x = p.expr_rel(3)

				localctx.(*Expr_relContext).opRight = _x
			}

			localctx.(*Expr_relContext).p = Primitive.NewOperation(localctx.(*Expr_relContext).GetOpLeft().GetP(), (func() string {
				if localctx.(*Expr_relContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expr_relContext).GetOp().GetText()
				}
			}()), localctx.(*Expr_relContext).GetOpRight().GetP(), (func() int {
				if localctx.(*Expr_relContext).GetOp() == nil {
					return 0
				} else {
					return localctx.(*Expr_relContext).GetOp().GetLine()
				}
			}()), localctx.(*Expr_relContext).GetOp().GetColumn())

		}
		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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

	// Get_expr_valor returns the _expr_valor rule contexts.
	Get_expr_valor() IExpr_valorContext

	// GetOpRight returns the opRight rule contexts.
	GetOpRight() IExpr_aritContext

	// SetOpLeft sets the opLeft rule contexts.
	SetOpLeft(IExpr_aritContext)

	// SetOpU sets the opU rule contexts.
	SetOpU(IExpressionContext)

	// Set_expr_valor sets the _expr_valor rule contexts.
	Set_expr_valor(IExpr_valorContext)

	// SetOpRight sets the opRight rule contexts.
	SetOpRight(IExpr_aritContext)

	// GetP returns the p attribute.
	GetP() AbstractOptimizer.Expression

	// SetP sets the p attribute.
	SetP(AbstractOptimizer.Expression)

	// IsExpr_aritContext differentiates from other interfaces.
	IsExpr_aritContext()
}

type Expr_aritContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           AbstractOptimizer.Expression
	opLeft      IExpr_aritContext
	op          antlr.Token
	opU         IExpressionContext
	_expr_valor IExpr_valorContext
	opRight     IExpr_aritContext
}

func NewEmptyExpr_aritContext() *Expr_aritContext {
	var p = new(Expr_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_expr_arit
	return p
}

func (*Expr_aritContext) IsExpr_aritContext() {}

func NewExpr_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_aritContext {
	var p = new(Expr_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_expr_arit

	return p
}

func (s *Expr_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expr_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_aritContext) GetOpLeft() IExpr_aritContext { return s.opLeft }

func (s *Expr_aritContext) GetOpU() IExpressionContext { return s.opU }

func (s *Expr_aritContext) Get_expr_valor() IExpr_valorContext { return s._expr_valor }

func (s *Expr_aritContext) GetOpRight() IExpr_aritContext { return s.opRight }

func (s *Expr_aritContext) SetOpLeft(v IExpr_aritContext) { s.opLeft = v }

func (s *Expr_aritContext) SetOpU(v IExpressionContext) { s.opU = v }

func (s *Expr_aritContext) Set_expr_valor(v IExpr_valorContext) { s._expr_valor = v }

func (s *Expr_aritContext) SetOpRight(v IExpr_aritContext) { s.opRight = v }

func (s *Expr_aritContext) GetP() AbstractOptimizer.Expression { return s.p }

func (s *Expr_aritContext) SetP(v AbstractOptimizer.Expression) { s.p = v }

func (s *Expr_aritContext) SUB() antlr.TerminalNode {
	return s.GetToken(OptimizerParserSUB, 0)
}

func (s *Expr_aritContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expr_aritContext) Expr_valor() IExpr_valorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valorContext)
}

func (s *Expr_aritContext) LEFT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserLEFT_PAR, 0)
}

func (s *Expr_aritContext) RIGHT_PAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRIGHT_PAR, 0)
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
	return s.GetToken(OptimizerParserMUL, 0)
}

func (s *Expr_aritContext) DIV() antlr.TerminalNode {
	return s.GetToken(OptimizerParserDIV, 0)
}

func (s *Expr_aritContext) ADD() antlr.TerminalNode {
	return s.GetToken(OptimizerParserADD, 0)
}

func (s *Expr_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterExpr_arit(s)
	}
}

func (s *Expr_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitExpr_arit(s)
	}
}

func (p *OptimizerParser) Expr_arit() (localctx IExpr_aritContext) {
	return p.expr_arit(0)
}

func (p *OptimizerParser) expr_arit(_p int) (localctx IExpr_aritContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, OptimizerParserRULE_expr_arit, _p)
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
	p.SetState(296)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case OptimizerParserSUB:
		{
			p.SetState(284)

			var _m = p.Match(OptimizerParserSUB)

			localctx.(*Expr_aritContext).op = _m
		}
		{
			p.SetState(285)

			var _x = p.Expression()

			localctx.(*Expr_aritContext).opU = _x
		}

		localctx.(*Expr_aritContext).p = Primitive.NewOperation(localctx.(*Expr_aritContext).GetOpU().GetP(), "-", nil, (func() int {
			if localctx.(*Expr_aritContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*Expr_aritContext).GetOp().GetLine()
			}
		}()), localctx.(*Expr_aritContext).GetOp().GetColumn())

	case OptimizerParserRPSTACK, OptimizerParserRPHEAP, OptimizerParserINTEGER, OptimizerParserFLOAT, OptimizerParserID:
		{
			p.SetState(288)

			var _x = p.Expr_valor()

			localctx.(*Expr_aritContext)._expr_valor = _x
		}

		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expr_valor().GetP()

	case OptimizerParserLEFT_PAR:
		{
			p.SetState(291)
			p.Match(OptimizerParserLEFT_PAR)
		}
		{
			p.SetState(292)
			p.Expression()
		}
		{
			p.SetState(293)
			p.Match(OptimizerParserRIGHT_PAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(308)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opLeft = _prevctx
				p.PushNewRecursionContext(localctx, _startState, OptimizerParserRULE_expr_arit)
				p.SetState(298)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(299)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == OptimizerParserMUL || _la == OptimizerParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(300)

					var _x = p.expr_arit(5)

					localctx.(*Expr_aritContext).opRight = _x
				}

				localctx.(*Expr_aritContext).p = Primitive.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpRight().GetP(), (func() int {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expr_aritContext).GetOp().GetColumn())

			case 2:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opLeft = _prevctx
				p.PushNewRecursionContext(localctx, _startState, OptimizerParserRULE_expr_arit)
				p.SetState(303)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(304)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == OptimizerParserADD || _la == OptimizerParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(305)

					var _x = p.expr_arit(4)

					localctx.(*Expr_aritContext).opRight = _x
				}

				localctx.(*Expr_aritContext).p = Primitive.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpRight().GetP(), (func() int {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetLine()
					}
				}()), localctx.(*Expr_aritContext).GetOp().GetColumn())

			}

		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RINT returns the _RINT token.
	Get_RINT() antlr.Token

	// Get_RFLOAT returns the _RFLOAT token.
	Get_RFLOAT() antlr.Token

	// Get_RCHAR returns the _RCHAR token.
	Get_RCHAR() antlr.Token

	// Set_RINT sets the _RINT token.
	Set_RINT(antlr.Token)

	// Set_RFLOAT sets the _RFLOAT token.
	Set_RFLOAT(antlr.Token)

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
	parser  antlr.Parser
	data    string
	_RINT   antlr.Token
	_RFLOAT antlr.Token
	_RCHAR  antlr.Token
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_data_type
	return p
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) Get_RINT() antlr.Token { return s._RINT }

func (s *Data_typeContext) Get_RFLOAT() antlr.Token { return s._RFLOAT }

func (s *Data_typeContext) Get_RCHAR() antlr.Token { return s._RCHAR }

func (s *Data_typeContext) Set_RINT(v antlr.Token) { s._RINT = v }

func (s *Data_typeContext) Set_RFLOAT(v antlr.Token) { s._RFLOAT = v }

func (s *Data_typeContext) Set_RCHAR(v antlr.Token) { s._RCHAR = v }

func (s *Data_typeContext) GetData() string { return s.data }

func (s *Data_typeContext) SetData(v string) { s.data = v }

func (s *Data_typeContext) RINT() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRINT, 0)
}

func (s *Data_typeContext) RFLOAT() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRFLOAT, 0)
}

func (s *Data_typeContext) RCHAR() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRCHAR, 0)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterData_type(s)
	}
}

func (s *Data_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitData_type(s)
	}
}

func (p *OptimizerParser) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, OptimizerParserRULE_data_type)

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

	p.SetState(319)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case OptimizerParserRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)

			var _m = p.Match(OptimizerParserRINT)

			localctx.(*Data_typeContext)._RINT = _m
		}
		localctx.(*Data_typeContext).data = (func() string {
			if localctx.(*Data_typeContext).Get_RINT() == nil {
				return ""
			} else {
				return localctx.(*Data_typeContext).Get_RINT().GetText()
			}
		}())

	case OptimizerParserRFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)

			var _m = p.Match(OptimizerParserRFLOAT)

			localctx.(*Data_typeContext)._RFLOAT = _m
		}
		localctx.(*Data_typeContext).data = (func() string {
			if localctx.(*Data_typeContext).Get_RFLOAT() == nil {
				return ""
			} else {
				return localctx.(*Data_typeContext).Get_RFLOAT().GetText()
			}
		}())

	case OptimizerParserRCHAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(317)

			var _m = p.Match(OptimizerParserRCHAR)

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

// IExpr_valorContext is an interface to support dynamic dispatch.
type IExpr_valorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_INTEGER returns the _INTEGER token.
	Get_INTEGER() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_RPSTACK returns the _RPSTACK token.
	Get_RPSTACK() antlr.Token

	// Get_RPHEAP returns the _RPHEAP token.
	Get_RPHEAP() antlr.Token

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_INTEGER sets the _INTEGER token.
	Set_INTEGER(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_RPSTACK sets the _RPSTACK token.
	Set_RPSTACK(antlr.Token)

	// Set_RPHEAP sets the _RPHEAP token.
	Set_RPHEAP(antlr.Token)

	// GetP returns the p attribute.
	GetP() AbstractOptimizer.Expression

	// SetP sets the p attribute.
	SetP(AbstractOptimizer.Expression)

	// IsExpr_valorContext differentiates from other interfaces.
	IsExpr_valorContext()
}

type Expr_valorContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        AbstractOptimizer.Expression
	_FLOAT   antlr.Token
	_INTEGER antlr.Token
	_ID      antlr.Token
	_RPSTACK antlr.Token
	_RPHEAP  antlr.Token
}

func NewEmptyExpr_valorContext() *Expr_valorContext {
	var p = new(Expr_valorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OptimizerParserRULE_expr_valor
	return p
}

func (*Expr_valorContext) IsExpr_valorContext() {}

func NewExpr_valorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_valorContext {
	var p = new(Expr_valorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OptimizerParserRULE_expr_valor

	return p
}

func (s *Expr_valorContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_valorContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *Expr_valorContext) Get_INTEGER() antlr.Token { return s._INTEGER }

func (s *Expr_valorContext) Get_ID() antlr.Token { return s._ID }

func (s *Expr_valorContext) Get_RPSTACK() antlr.Token { return s._RPSTACK }

func (s *Expr_valorContext) Get_RPHEAP() antlr.Token { return s._RPHEAP }

func (s *Expr_valorContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *Expr_valorContext) Set_INTEGER(v antlr.Token) { s._INTEGER = v }

func (s *Expr_valorContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Expr_valorContext) Set_RPSTACK(v antlr.Token) { s._RPSTACK = v }

func (s *Expr_valorContext) Set_RPHEAP(v antlr.Token) { s._RPHEAP = v }

func (s *Expr_valorContext) GetP() AbstractOptimizer.Expression { return s.p }

func (s *Expr_valorContext) SetP(v AbstractOptimizer.Expression) { s.p = v }

func (s *Expr_valorContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(OptimizerParserFLOAT, 0)
}

func (s *Expr_valorContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(OptimizerParserINTEGER, 0)
}

func (s *Expr_valorContext) ID() antlr.TerminalNode {
	return s.GetToken(OptimizerParserID, 0)
}

func (s *Expr_valorContext) RPSTACK() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRPSTACK, 0)
}

func (s *Expr_valorContext) RPHEAP() antlr.TerminalNode {
	return s.GetToken(OptimizerParserRPHEAP, 0)
}

func (s *Expr_valorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_valorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_valorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.EnterExpr_valor(s)
	}
}

func (s *Expr_valorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OptimizerParserListener); ok {
		listenerT.ExitExpr_valor(s)
	}
}

func (p *OptimizerParser) Expr_valor() (localctx IExpr_valorContext) {
	localctx = NewExpr_valorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, OptimizerParserRULE_expr_valor)

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

	p.SetState(331)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case OptimizerParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)

			var _m = p.Match(OptimizerParserFLOAT)

			localctx.(*Expr_valorContext)._FLOAT = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*Expr_valorContext).Get_FLOAT() == nil {
				return ""
			} else {
				return localctx.(*Expr_valorContext).Get_FLOAT().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*Expr_valorContext).p = Primitive.NewNumber(num, (func() int {
			if localctx.(*Expr_valorContext).Get_FLOAT() == nil {
				return 0
			} else {
				return localctx.(*Expr_valorContext).Get_FLOAT().GetLine()
			}
		}()), localctx.(*Expr_valorContext).Get_FLOAT().GetColumn())

	case OptimizerParserINTEGER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(323)

			var _m = p.Match(OptimizerParserINTEGER)

			localctx.(*Expr_valorContext)._INTEGER = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*Expr_valorContext).Get_INTEGER() == nil {
				return ""
			} else {
				return localctx.(*Expr_valorContext).Get_INTEGER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*Expr_valorContext).p = Primitive.NewNumber(num, (func() int {
			if localctx.(*Expr_valorContext).Get_INTEGER() == nil {
				return 0
			} else {
				return localctx.(*Expr_valorContext).Get_INTEGER().GetLine()
			}
		}()), localctx.(*Expr_valorContext).Get_INTEGER().GetColumn())

	case OptimizerParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(325)

			var _m = p.Match(OptimizerParserID)

			localctx.(*Expr_valorContext)._ID = _m
		}

		localctx.(*Expr_valorContext).p = Primitive.NewTemp((func() string {
			if localctx.(*Expr_valorContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Expr_valorContext).Get_ID().GetText()
			}
		}()), (func() int {
			if localctx.(*Expr_valorContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Expr_valorContext).Get_ID().GetLine()
			}
		}()), localctx.(*Expr_valorContext).Get_ID().GetColumn())

	case OptimizerParserRPSTACK:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(327)

			var _m = p.Match(OptimizerParserRPSTACK)

			localctx.(*Expr_valorContext)._RPSTACK = _m
		}

		localctx.(*Expr_valorContext).p = Primitive.NewTemp((func() string {
			if localctx.(*Expr_valorContext).Get_RPSTACK() == nil {
				return ""
			} else {
				return localctx.(*Expr_valorContext).Get_RPSTACK().GetText()
			}
		}()), (func() int {
			if localctx.(*Expr_valorContext).Get_RPSTACK() == nil {
				return 0
			} else {
				return localctx.(*Expr_valorContext).Get_RPSTACK().GetLine()
			}
		}()), localctx.(*Expr_valorContext).Get_RPSTACK().GetColumn())

	case OptimizerParserRPHEAP:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(329)

			var _m = p.Match(OptimizerParserRPHEAP)

			localctx.(*Expr_valorContext)._RPHEAP = _m
		}

		localctx.(*Expr_valorContext).p = Primitive.NewTemp((func() string {
			if localctx.(*Expr_valorContext).Get_RPHEAP() == nil {
				return ""
			} else {
				return localctx.(*Expr_valorContext).Get_RPHEAP().GetText()
			}
		}()), (func() int {
			if localctx.(*Expr_valorContext).Get_RPHEAP() == nil {
				return 0
			} else {
				return localctx.(*Expr_valorContext).Get_RPHEAP().GetLine()
			}
		}()), localctx.(*Expr_valorContext).Get_RPHEAP().GetColumn())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *OptimizerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *TempsContext = nil
		if localctx != nil {
			t = localctx.(*TempsContext)
		}
		return p.Temps_Sempred(t, predIndex)

	case 4:
		var t *ListFuncsContext = nil
		if localctx != nil {
			t = localctx.(*ListFuncsContext)
		}
		return p.ListFuncs_Sempred(t, predIndex)

	case 21:
		var t *Expr_relContext = nil
		if localctx != nil {
			t = localctx.(*Expr_relContext)
		}
		return p.Expr_rel_Sempred(t, predIndex)

	case 22:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *OptimizerParser) Temps_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *OptimizerParser) ListFuncs_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *OptimizerParser) Expr_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *OptimizerParser) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
