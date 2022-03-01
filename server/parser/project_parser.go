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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 334,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 45, 10, 3,
	12, 3, 14, 3, 48, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 64, 10, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 102, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 113, 10, 6,
	12, 6, 14, 6, 116, 11, 6, 3, 7, 3, 7, 5, 7, 120, 10, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 125, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 134,
	10, 7, 3, 7, 3, 7, 3, 7, 5, 7, 139, 10, 7, 3, 7, 3, 7, 3, 7, 5, 7, 144,
	10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 166, 10, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 176, 10,
	10, 12, 10, 14, 10, 179, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	196, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 5, 13, 210, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 221, 10, 14, 12, 14, 14, 14,
	224, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 250, 10, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 262, 10, 15,
	12, 15, 14, 15, 265, 11, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 275, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 5, 17, 286, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 308, 10, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 319, 10, 20, 3,
	20, 3, 20, 5, 20, 323, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 5, 20, 332, 10, 20, 3, 20, 2, 6, 10, 18, 26, 28, 21, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 6, 4, 2,
	42, 42, 44, 47, 3, 2, 48, 50, 3, 2, 51, 52, 3, 2, 53, 54, 2, 356, 2, 40,
	3, 2, 2, 2, 4, 46, 3, 2, 2, 2, 6, 63, 3, 2, 2, 2, 8, 101, 3, 2, 2, 2, 10,
	103, 3, 2, 2, 2, 12, 143, 3, 2, 2, 2, 14, 145, 3, 2, 2, 2, 16, 165, 3,
	2, 2, 2, 18, 167, 3, 2, 2, 2, 20, 180, 3, 2, 2, 2, 22, 195, 3, 2, 2, 2,
	24, 209, 3, 2, 2, 2, 26, 211, 3, 2, 2, 2, 28, 249, 3, 2, 2, 2, 30, 274,
	3, 2, 2, 2, 32, 285, 3, 2, 2, 2, 34, 287, 3, 2, 2, 2, 36, 307, 3, 2, 2,
	2, 38, 331, 3, 2, 2, 2, 40, 41, 5, 4, 3, 2, 41, 42, 8, 2, 1, 2, 42, 3,
	3, 2, 2, 2, 43, 45, 5, 6, 4, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2,
	46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 46, 3,
	2, 2, 2, 49, 50, 8, 3, 1, 2, 50, 5, 3, 2, 2, 2, 51, 52, 5, 8, 5, 2, 52,
	53, 8, 4, 1, 2, 53, 64, 3, 2, 2, 2, 54, 55, 5, 12, 7, 2, 55, 56, 8, 4,
	1, 2, 56, 64, 3, 2, 2, 2, 57, 58, 5, 14, 8, 2, 58, 59, 8, 4, 1, 2, 59,
	64, 3, 2, 2, 2, 60, 61, 5, 20, 11, 2, 61, 62, 8, 4, 1, 2, 62, 64, 3, 2,
	2, 2, 63, 51, 3, 2, 2, 2, 63, 54, 3, 2, 2, 2, 63, 57, 3, 2, 2, 2, 63, 60,
	3, 2, 2, 2, 64, 7, 3, 2, 2, 2, 65, 66, 7, 4, 2, 2, 66, 67, 7, 33, 2, 2,
	67, 68, 7, 36, 2, 2, 68, 69, 5, 24, 13, 2, 69, 70, 7, 37, 2, 2, 70, 71,
	7, 30, 2, 2, 71, 72, 8, 5, 1, 2, 72, 102, 3, 2, 2, 2, 73, 74, 7, 4, 2,
	2, 74, 75, 7, 33, 2, 2, 75, 76, 7, 36, 2, 2, 76, 77, 5, 24, 13, 2, 77,
	78, 7, 31, 2, 2, 78, 79, 5, 10, 6, 2, 79, 80, 7, 37, 2, 2, 80, 81, 7, 30,
	2, 2, 81, 82, 8, 5, 1, 2, 82, 102, 3, 2, 2, 2, 83, 84, 7, 3, 2, 2, 84,
	85, 7, 33, 2, 2, 85, 86, 7, 36, 2, 2, 86, 87, 5, 24, 13, 2, 87, 88, 7,
	37, 2, 2, 88, 89, 7, 30, 2, 2, 89, 90, 8, 5, 1, 2, 90, 102, 3, 2, 2, 2,
	91, 92, 7, 3, 2, 2, 92, 93, 7, 33, 2, 2, 93, 94, 7, 36, 2, 2, 94, 95, 5,
	24, 13, 2, 95, 96, 7, 31, 2, 2, 96, 97, 5, 10, 6, 2, 97, 98, 7, 37, 2,
	2, 98, 99, 7, 30, 2, 2, 99, 100, 8, 5, 1, 2, 100, 102, 3, 2, 2, 2, 101,
	65, 3, 2, 2, 2, 101, 73, 3, 2, 2, 2, 101, 83, 3, 2, 2, 2, 101, 91, 3, 2,
	2, 2, 102, 9, 3, 2, 2, 2, 103, 104, 8, 6, 1, 2, 104, 105, 5, 24, 13, 2,
	105, 106, 8, 6, 1, 2, 106, 114, 3, 2, 2, 2, 107, 108, 12, 4, 2, 2, 108,
	109, 7, 31, 2, 2, 109, 110, 5, 24, 13, 2, 110, 111, 8, 6, 1, 2, 111, 113,
	3, 2, 2, 2, 112, 107, 3, 2, 2, 2, 113, 116, 3, 2, 2, 2, 114, 112, 3, 2,
	2, 2, 114, 115, 3, 2, 2, 2, 115, 11, 3, 2, 2, 2, 116, 114, 3, 2, 2, 2,
	117, 119, 7, 5, 2, 2, 118, 120, 7, 6, 2, 2, 119, 118, 3, 2, 2, 2, 119,
	120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 124, 5, 16, 9, 2, 122, 123,
	7, 32, 2, 2, 123, 125, 5, 36, 19, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3,
	2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 7, 28, 2, 2, 127, 128, 5, 24,
	13, 2, 128, 129, 7, 30, 2, 2, 129, 130, 8, 7, 1, 2, 130, 144, 3, 2, 2,
	2, 131, 133, 7, 5, 2, 2, 132, 134, 7, 6, 2, 2, 133, 132, 3, 2, 2, 2, 133,
	134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 138, 5, 16, 9, 2, 136, 137,
	7, 32, 2, 2, 137, 139, 5, 36, 19, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3,
	2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141, 7, 30, 2, 2, 141, 142, 8, 7, 1,
	2, 142, 144, 3, 2, 2, 2, 143, 117, 3, 2, 2, 2, 143, 131, 3, 2, 2, 2, 144,
	13, 3, 2, 2, 2, 145, 146, 5, 18, 10, 2, 146, 147, 7, 28, 2, 2, 147, 148,
	5, 24, 13, 2, 148, 149, 7, 30, 2, 2, 149, 150, 8, 8, 1, 2, 150, 15, 3,
	2, 2, 2, 151, 152, 5, 18, 10, 2, 152, 153, 8, 9, 1, 2, 153, 166, 3, 2,
	2, 2, 154, 155, 5, 18, 10, 2, 155, 156, 7, 32, 2, 2, 156, 157, 7, 7, 2,
	2, 157, 158, 8, 9, 1, 2, 158, 166, 3, 2, 2, 2, 159, 160, 5, 18, 10, 2,
	160, 161, 7, 32, 2, 2, 161, 162, 7, 34, 2, 2, 162, 163, 7, 12, 2, 2, 163,
	164, 8, 9, 1, 2, 164, 166, 3, 2, 2, 2, 165, 151, 3, 2, 2, 2, 165, 154,
	3, 2, 2, 2, 165, 159, 3, 2, 2, 2, 166, 17, 3, 2, 2, 2, 167, 168, 8, 10,
	1, 2, 168, 169, 7, 27, 2, 2, 169, 170, 8, 10, 1, 2, 170, 177, 3, 2, 2,
	2, 171, 172, 12, 4, 2, 2, 172, 173, 7, 31, 2, 2, 173, 174, 7, 27, 2, 2,
	174, 176, 8, 10, 1, 2, 175, 171, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177,
	175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 19, 3, 2, 2, 2, 179, 177, 3,
	2, 2, 2, 180, 181, 7, 14, 2, 2, 181, 182, 7, 36, 2, 2, 182, 183, 5, 24,
	13, 2, 183, 184, 7, 37, 2, 2, 184, 185, 5, 22, 12, 2, 185, 186, 8, 11,
	1, 2, 186, 21, 3, 2, 2, 2, 187, 188, 7, 40, 2, 2, 188, 189, 5, 4, 3, 2,
	189, 190, 7, 41, 2, 2, 190, 191, 8, 12, 1, 2, 191, 196, 3, 2, 2, 2, 192,
	193, 7, 40, 2, 2, 193, 194, 7, 41, 2, 2, 194, 196, 8, 12, 1, 2, 195, 187,
	3, 2, 2, 2, 195, 192, 3, 2, 2, 2, 196, 23, 3, 2, 2, 2, 197, 198, 5, 26,
	14, 2, 198, 199, 8, 13, 1, 2, 199, 210, 3, 2, 2, 2, 200, 201, 5, 28, 15,
	2, 201, 202, 8, 13, 1, 2, 202, 210, 3, 2, 2, 2, 203, 204, 5, 32, 17, 2,
	204, 205, 8, 13, 1, 2, 205, 210, 3, 2, 2, 2, 206, 207, 5, 34, 18, 2, 207,
	208, 8, 13, 1, 2, 208, 210, 3, 2, 2, 2, 209, 197, 3, 2, 2, 2, 209, 200,
	3, 2, 2, 2, 209, 203, 3, 2, 2, 2, 209, 206, 3, 2, 2, 2, 210, 25, 3, 2,
	2, 2, 211, 212, 8, 14, 1, 2, 212, 213, 5, 28, 15, 2, 213, 214, 8, 14, 1,
	2, 214, 222, 3, 2, 2, 2, 215, 216, 12, 4, 2, 2, 216, 217, 9, 2, 2, 2, 217,
	218, 5, 26, 14, 5, 218, 219, 8, 14, 1, 2, 219, 221, 3, 2, 2, 2, 220, 215,
	3, 2, 2, 2, 221, 224, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 222, 223, 3, 2,
	2, 2, 223, 27, 3, 2, 2, 2, 224, 222, 3, 2, 2, 2, 225, 226, 8, 15, 1, 2,
	226, 227, 7, 52, 2, 2, 227, 228, 5, 24, 13, 2, 228, 229, 8, 15, 1, 2, 229,
	250, 3, 2, 2, 2, 230, 231, 5, 30, 16, 2, 231, 232, 7, 36, 2, 2, 232, 233,
	5, 28, 15, 2, 233, 234, 7, 31, 2, 2, 234, 235, 5, 28, 15, 2, 235, 236,
	7, 37, 2, 2, 236, 237, 8, 15, 1, 2, 237, 250, 3, 2, 2, 2, 238, 239, 5,
	38, 20, 2, 239, 240, 8, 15, 1, 2, 240, 250, 3, 2, 2, 2, 241, 242, 5, 34,
	18, 2, 242, 243, 8, 15, 1, 2, 243, 250, 3, 2, 2, 2, 244, 245, 7, 36, 2,
	2, 245, 246, 5, 24, 13, 2, 246, 247, 7, 37, 2, 2, 247, 248, 8, 15, 1, 2,
	248, 250, 3, 2, 2, 2, 249, 225, 3, 2, 2, 2, 249, 230, 3, 2, 2, 2, 249,
	238, 3, 2, 2, 2, 249, 241, 3, 2, 2, 2, 249, 244, 3, 2, 2, 2, 250, 263,
	3, 2, 2, 2, 251, 252, 12, 7, 2, 2, 252, 253, 9, 3, 2, 2, 253, 254, 5, 28,
	15, 8, 254, 255, 8, 15, 1, 2, 255, 262, 3, 2, 2, 2, 256, 257, 12, 6, 2,
	2, 257, 258, 9, 4, 2, 2, 258, 259, 5, 28, 15, 7, 259, 260, 8, 15, 1, 2,
	260, 262, 3, 2, 2, 2, 261, 251, 3, 2, 2, 2, 261, 256, 3, 2, 2, 2, 262,
	265, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 29, 3,
	2, 2, 2, 265, 263, 3, 2, 2, 2, 266, 267, 7, 8, 2, 2, 267, 268, 7, 35, 2,
	2, 268, 269, 7, 16, 2, 2, 269, 275, 8, 16, 1, 2, 270, 271, 7, 9, 2, 2,
	271, 272, 7, 35, 2, 2, 272, 273, 7, 17, 2, 2, 273, 275, 8, 16, 1, 2, 274,
	266, 3, 2, 2, 2, 274, 270, 3, 2, 2, 2, 275, 31, 3, 2, 2, 2, 276, 277, 7,
	33, 2, 2, 277, 278, 5, 24, 13, 2, 278, 279, 8, 17, 1, 2, 279, 286, 3, 2,
	2, 2, 280, 281, 5, 26, 14, 2, 281, 282, 9, 5, 2, 2, 282, 283, 5, 26, 14,
	2, 283, 284, 8, 17, 1, 2, 284, 286, 3, 2, 2, 2, 285, 276, 3, 2, 2, 2, 285,
	280, 3, 2, 2, 2, 286, 33, 3, 2, 2, 2, 287, 288, 7, 36, 2, 2, 288, 289,
	5, 24, 13, 2, 289, 290, 7, 13, 2, 2, 290, 291, 5, 36, 19, 2, 291, 292,
	7, 37, 2, 2, 292, 293, 8, 18, 1, 2, 293, 35, 3, 2, 2, 2, 294, 295, 7, 8,
	2, 2, 295, 308, 8, 19, 1, 2, 296, 297, 7, 9, 2, 2, 297, 308, 8, 19, 1,
	2, 298, 299, 7, 12, 2, 2, 299, 308, 8, 19, 1, 2, 300, 301, 7, 7, 2, 2,
	301, 308, 8, 19, 1, 2, 302, 303, 7, 10, 2, 2, 303, 308, 8, 19, 1, 2, 304,
	305, 7, 11, 2, 2, 305, 308, 8, 19, 1, 2, 306, 308, 3, 2, 2, 2, 307, 294,
	3, 2, 2, 2, 307, 296, 3, 2, 2, 2, 307, 298, 3, 2, 2, 2, 307, 300, 3, 2,
	2, 2, 307, 302, 3, 2, 2, 2, 307, 304, 3, 2, 2, 2, 307, 306, 3, 2, 2, 2,
	308, 37, 3, 2, 2, 2, 309, 310, 7, 22, 2, 2, 310, 332, 8, 20, 1, 2, 311,
	312, 7, 23, 2, 2, 312, 332, 8, 20, 1, 2, 313, 322, 7, 25, 2, 2, 314, 315,
	7, 29, 2, 2, 315, 319, 7, 18, 2, 2, 316, 317, 7, 29, 2, 2, 317, 319, 7,
	19, 2, 2, 318, 314, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 319, 320, 3, 2, 2,
	2, 320, 321, 7, 36, 2, 2, 321, 323, 7, 37, 2, 2, 322, 318, 3, 2, 2, 2,
	322, 323, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 332, 8, 20, 1, 2, 325,
	326, 7, 24, 2, 2, 326, 332, 8, 20, 1, 2, 327, 328, 7, 26, 2, 2, 328, 332,
	8, 20, 1, 2, 329, 330, 7, 27, 2, 2, 330, 332, 8, 20, 1, 2, 331, 309, 3,
	2, 2, 2, 331, 311, 3, 2, 2, 2, 331, 313, 3, 2, 2, 2, 331, 325, 3, 2, 2,
	2, 331, 327, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 332, 39, 3, 2, 2, 2, 25,
	46, 63, 101, 114, 119, 124, 133, 138, 143, 165, 177, 195, 209, 222, 249,
	261, 263, 274, 285, 307, 318, 322, 331,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", "'f64'",
	"'bool'", "'char'", "'&str'", "'as'", "'If'", "'entonces'", "'pow'", "'powf'",
	"'to_string'", "'to_owned'", "", "", "", "", "", "", "", "", "'='", "'.'",
	"';'", "','", "':'", "'!'", "'&'", "'::'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'*'", "'/'",
	"'%'", "'+'", "'-'", "'&&'", "'||'",
}
var symbolicNames = []string{
	"", "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL",
	"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "POWI", "POWF", "TOSTRING",
	"TOOWNED", "MULTILINE", "INLINE", "INTEGER", "FLOAT", "CHAR", "STRING",
	"BOOLEAN", "ID", "EQUAL", "DOT", "SEMICOLON", "COMMA", "COLON", "ADMIRATION",
	"REFERENCE", "HERITAGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY",
	"LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN",
	"LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD",
	"ADD", "SUB", "AND", "OR", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instructions", "instruction", "print_prod", "listVars", "declaration_prod",
	"assign_prod", "ids_type", "listIds", "conditional_prod", "bloq", "expression",
	"expr_rel", "expr_arit", "pow_op", "expr_logic", "expr_cast", "data_type",
	"primitive",
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
	ProjectParserPOWI              = 14
	ProjectParserPOWF              = 15
	ProjectParserTOSTRING          = 16
	ProjectParserTOOWNED           = 17
	ProjectParserMULTILINE         = 18
	ProjectParserINLINE            = 19
	ProjectParserINTEGER           = 20
	ProjectParserFLOAT             = 21
	ProjectParserCHAR              = 22
	ProjectParserSTRING            = 23
	ProjectParserBOOLEAN           = 24
	ProjectParserID                = 25
	ProjectParserEQUAL             = 26
	ProjectParserDOT               = 27
	ProjectParserSEMICOLON         = 28
	ProjectParserCOMMA             = 29
	ProjectParserCOLON             = 30
	ProjectParserADMIRATION        = 31
	ProjectParserREFERENCE         = 32
	ProjectParserHERITAGE          = 33
	ProjectParserLEFT_PAR          = 34
	ProjectParserRIGHT_PAR         = 35
	ProjectParserLEFT_KEY          = 36
	ProjectParserRIGHT_KEY         = 37
	ProjectParserLEFT_BRACKET      = 38
	ProjectParserRIGHT_BRACKET     = 39
	ProjectParserEQUEAL_EQUAL      = 40
	ProjectParserNOTEQUAL          = 41
	ProjectParserGREATER_THAN      = 42
	ProjectParserLESS_THAN         = 43
	ProjectParserGREATER_EQUALTHAN = 44
	ProjectParserLESS_EQUEALTHAN   = 45
	ProjectParserMUL               = 46
	ProjectParserDIV               = 47
	ProjectParserMOD               = 48
	ProjectParserADD               = 49
	ProjectParserSUB               = 50
	ProjectParserAND               = 51
	ProjectParserOR                = 52
	ProjectParserWHITESPACE        = 53
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
	ProjectParserRULE_expr_arit        = 13
	ProjectParserRULE_pow_op           = 14
	ProjectParserRULE_expr_logic       = 15
	ProjectParserRULE_expr_cast        = 16
	ProjectParserRULE_data_type        = 17
	ProjectParserRULE_primitive        = 18
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
		p.SetState(38)

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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ProjectParserPRINTLN)|(1<<ProjectParserPRINT)|(1<<ProjectParserDECLARAR)|(1<<ProjectParserRIF)|(1<<ProjectParserID))) != 0 {
		{
			p.SetState(41)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(46)
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

	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserPRINTLN, ProjectParserPRINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)

			var _x = p.Print_prod()

			localctx.(*InstructionContext)._print_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_print_prod().GetInstr()

	case ProjectParserDECLARAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)

			var _x = p.Declaration_prod()

			localctx.(*InstructionContext)._declaration_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_declaration_prod().GetInstr()

	case ProjectParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)

			var _x = p.Assign_prod()

			localctx.(*InstructionContext)._assign_prod = _x
		}
		localctx.(*InstructionContext).instr = localctx.(*InstructionContext).Get_assign_prod().GetInstr()

	case ProjectParserRIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)

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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(ProjectParserPRINT)
		}
		{
			p.SetState(64)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(65)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(66)

			var _x = p.Expression()

			localctx.(*Print_prodContext)._expression = _x
		}
		{
			p.SetState(67)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(68)
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
			p.SetState(71)
			p.Match(ProjectParserPRINT)
		}
		{
			p.SetState(72)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(73)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(74)

			var _x = p.Expression()

			localctx.(*Print_prodContext).opBefore = _x
		}
		{
			p.SetState(75)
			p.Match(ProjectParserCOMMA)
		}
		{
			p.SetState(76)

			var _x = p.listVars(0)

			localctx.(*Print_prodContext)._listVars = _x
		}
		{
			p.SetState(77)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(78)
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
			p.SetState(81)
			p.Match(ProjectParserPRINTLN)
		}
		{
			p.SetState(82)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(83)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(84)

			var _x = p.Expression()

			localctx.(*Print_prodContext)._expression = _x
		}
		{
			p.SetState(85)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(86)
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
			p.SetState(89)
			p.Match(ProjectParserPRINTLN)
		}
		{
			p.SetState(90)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(91)

			var _m = p.Match(ProjectParserLEFT_PAR)

			localctx.(*Print_prodContext)._LEFT_PAR = _m
		}
		{
			p.SetState(92)

			var _x = p.Expression()

			localctx.(*Print_prodContext).opBefore = _x
		}
		{
			p.SetState(93)
			p.Match(ProjectParserCOMMA)
		}
		{
			p.SetState(94)

			var _x = p.listVars(0)

			localctx.(*Print_prodContext)._listVars = _x
		}
		{
			p.SetState(95)
			p.Match(ProjectParserRIGHT_PAR)
		}
		{
			p.SetState(96)
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
		p.SetState(102)

		var _x = p.Expression()

		localctx.(*ListVarsContext)._expression = _x
	}

	localctx.(*ListVarsContext).list.Add(localctx.(*ListVarsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(112)
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
			p.SetState(105)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(106)
				p.Match(ProjectParserCOMMA)
			}
			{
				p.SetState(107)

				var _x = p.Expression()

				localctx.(*ListVarsContext)._expression = _x
			}

			localctx.(*ListVarsContext).GetSub().GetList().Add(localctx.(*ListVarsContext).Get_expression().GetP())
			localctx.(*ListVarsContext).list = localctx.(*ListVarsContext).GetSub().GetList()

		}
		p.SetState(114)
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

	// Get_MUT returns the _MUT token.
	Get_MUT() antlr.Token

	// Set_MUT sets the _MUT token.
	Set_MUT(antlr.Token)

	// Get_ids_type returns the _ids_type rule contexts.
	Get_ids_type() IIds_typeContext

	// Get_data_type returns the _data_type rule contexts.
	Get_data_type() IData_typeContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_ids_type sets the _ids_type rule contexts.
	Set_ids_type(IIds_typeContext)

	// Set_data_type sets the _data_type rule contexts.
	Set_data_type(IData_typeContext)

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
	_MUT        antlr.Token
	_ids_type   IIds_typeContext
	_data_type  IData_typeContext
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

func (s *Declaration_prodContext) Get_MUT() antlr.Token { return s._MUT }

func (s *Declaration_prodContext) Set_MUT(v antlr.Token) { s._MUT = v }

func (s *Declaration_prodContext) Get_ids_type() IIds_typeContext { return s._ids_type }

func (s *Declaration_prodContext) Get_data_type() IData_typeContext { return s._data_type }

func (s *Declaration_prodContext) Get_expression() IExpressionContext { return s._expression }

func (s *Declaration_prodContext) Set_ids_type(v IIds_typeContext) { s._ids_type = v }

func (s *Declaration_prodContext) Set_data_type(v IData_typeContext) { s._data_type = v }

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

func (s *Declaration_prodContext) MUT() antlr.TerminalNode {
	return s.GetToken(ProjectParserMUT, 0)
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
	p.EnterRule(localctx, 10, ProjectParserRULE_declaration_prod)
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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Match(ProjectParserDECLARAR)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProjectParserMUT {
			{
				p.SetState(116)

				var _m = p.Match(ProjectParserMUT)

				localctx.(*Declaration_prodContext)._MUT = _m
			}

		}
		{
			p.SetState(119)

			var _x = p.Ids_type()

			localctx.(*Declaration_prodContext)._ids_type = _x
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProjectParserCOLON {
			{
				p.SetState(120)
				p.Match(ProjectParserCOLON)
			}
			{
				p.SetState(121)

				var _x = p.Data_type()

				localctx.(*Declaration_prodContext)._data_type = _x
			}

		}
		{
			p.SetState(124)
			p.Match(ProjectParserEQUAL)
		}
		{
			p.SetState(125)

			var _x = p.Expression()

			localctx.(*Declaration_prodContext)._expression = _x
		}
		{
			p.SetState(126)
			p.Match(ProjectParserSEMICOLON)
		}

		if (func() string {
			if localctx.(*Declaration_prodContext).Get_MUT() == nil {
				return ""
			} else {
				return localctx.(*Declaration_prodContext).Get_MUT().GetText()
			}
		}()) != "" {
			if localctx.(*Declaration_prodContext).Get_data_type() != nil {
				localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP(), true, localctx.(*Declaration_prodContext).Get_data_type().GetData())
			} else {
				localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP(), true, "")
			}
		} else {
			if localctx.(*Declaration_prodContext).Get_data_type() != nil {
				localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP(), false, localctx.(*Declaration_prodContext).Get_data_type().GetData())
			} else {
				localctx.(*Declaration_prodContext).instr = Natives.NewDeclarationInit(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), localctx.(*Declaration_prodContext).Get_expression().GetP(), false, "")
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(129)
			p.Match(ProjectParserDECLARAR)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProjectParserMUT {
			{
				p.SetState(130)

				var _m = p.Match(ProjectParserMUT)

				localctx.(*Declaration_prodContext)._MUT = _m
			}

		}
		{
			p.SetState(133)

			var _x = p.Ids_type()

			localctx.(*Declaration_prodContext)._ids_type = _x
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ProjectParserCOLON {
			{
				p.SetState(134)
				p.Match(ProjectParserCOLON)
			}
			{
				p.SetState(135)

				var _x = p.Data_type()

				localctx.(*Declaration_prodContext)._data_type = _x
			}

		}
		{
			p.SetState(138)
			p.Match(ProjectParserSEMICOLON)
		}

		if (func() string {
			if localctx.(*Declaration_prodContext).Get_MUT() == nil {
				return ""
			} else {
				return localctx.(*Declaration_prodContext).Get_MUT().GetText()
			}
		}()) != "" {
			if localctx.(*Declaration_prodContext).Get_data_type() != nil {
				localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), true, localctx.(*Declaration_prodContext).Get_data_type().GetData())
			} else {
				localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), true, "")
			}
		} else {
			if localctx.(*Declaration_prodContext).Get_data_type() != nil {
				localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), true, localctx.(*Declaration_prodContext).Get_data_type().GetData())
			} else {
				localctx.(*Declaration_prodContext).instr = Natives.NewDeclaration(localctx.(*Declaration_prodContext).Get_ids_type().GetList(), true, "")
			}
		}

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
		p.SetState(143)

		var _x = p.listIds(0)

		localctx.(*Assign_prodContext)._listIds = _x
	}
	{
		p.SetState(144)
		p.Match(ProjectParserEQUAL)
	}
	{
		p.SetState(145)

		var _x = p.Expression()

		localctx.(*Assign_prodContext)._expression = _x
	}
	{
		p.SetState(146)
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

	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		localctx.(*Ids_typeContext).list = localctx.(*Ids_typeContext).Get_listIds().GetList()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		{
			p.SetState(153)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(154)
			p.Match(ProjectParserRSTRING)
		}
		localctx.(*Ids_typeContext).list = localctx.(*Ids_typeContext).Get_listIds().GetList()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)

			var _x = p.listIds(0)

			localctx.(*Ids_typeContext)._listIds = _x
		}
		{
			p.SetState(158)
			p.Match(ProjectParserCOLON)
		}
		{
			p.SetState(159)
			p.Match(ProjectParserREFERENCE)
		}
		{
			p.SetState(160)
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
		p.SetState(166)

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
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListIdsContext(p, _parentctx, _parentState)
			localctx.(*ListIdsContext).sub = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_listIds)
			p.SetState(169)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(170)
				p.Match(ProjectParserCOMMA)
			}
			{
				p.SetState(171)

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
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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
		p.SetState(178)
		p.Match(ProjectParserRIF)
	}
	{
		p.SetState(179)
		p.Match(ProjectParserLEFT_PAR)
	}
	{
		p.SetState(180)

		var _x = p.Expression()

		localctx.(*Conditional_prodContext)._expression = _x
	}
	{
		p.SetState(181)
		p.Match(ProjectParserRIGHT_PAR)
	}
	{
		p.SetState(182)

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

	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(ProjectParserLEFT_BRACKET)
		}
		{
			p.SetState(186)

			var _x = p.Instructions()

			localctx.(*BloqContext)._instructions = _x
		}
		{
			p.SetState(187)
			p.Match(ProjectParserRIGHT_BRACKET)
		}
		localctx.(*BloqContext).content = localctx.(*BloqContext).Get_instructions().GetL()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(ProjectParserLEFT_BRACKET)
		}
		{
			p.SetState(191)
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

	// Get_expr_cast returns the _expr_cast rule contexts.
	Get_expr_cast() IExpr_castContext

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// Set_expr_logic sets the _expr_logic rule contexts.
	Set_expr_logic(IExpr_logicContext)

	// Set_expr_cast sets the _expr_cast rule contexts.
	Set_expr_cast(IExpr_castContext)

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
	_expr_cast  IExpr_castContext
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

func (s *ExpressionContext) Get_expr_cast() IExpr_castContext { return s._expr_cast }

func (s *ExpressionContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *ExpressionContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *ExpressionContext) Set_expr_logic(v IExpr_logicContext) { s._expr_logic = v }

func (s *ExpressionContext) Set_expr_cast(v IExpr_castContext) { s._expr_cast = v }

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

func (s *ExpressionContext) Expr_cast() IExpr_castContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_castContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_castContext)
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

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)

			var _x = p.expr_rel(0)

			localctx.(*ExpressionContext)._expr_rel = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_rel().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)

			var _x = p.expr_arit(0)

			localctx.(*ExpressionContext)._expr_arit = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(201)

			var _x = p.Expr_logic()

			localctx.(*ExpressionContext)._expr_logic = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_logic().GetP()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)

			var _x = p.Expr_cast()

			localctx.(*ExpressionContext)._expr_cast = _x
		}
		localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_cast().GetP()

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
		p.SetState(210)

		var _x = p.expr_arit(0)

		localctx.(*Expr_relContext)._expr_arit = _x
	}
	localctx.(*Expr_relContext).p = localctx.(*Expr_relContext).Get_expr_arit().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_relContext(p, _parentctx, _parentState)
			localctx.(*Expr_relContext).opLeft = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_expr_rel)
			p.SetState(213)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(214)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expr_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ProjectParserEQUEAL_EQUAL-40))|(1<<(ProjectParserGREATER_THAN-40))|(1<<(ProjectParserLESS_THAN-40))|(1<<(ProjectParserGREATER_EQUALTHAN-40))|(1<<(ProjectParserLESS_EQUEALTHAN-40)))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expr_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(215)

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
		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
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

	// Get_primitive returns the _primitive rule contexts.
	Get_primitive() IPrimitiveContext

	// Get_expr_cast returns the _expr_cast rule contexts.
	Get_expr_cast() IExpr_castContext

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

	// Set_primitive sets the _primitive rule contexts.
	Set_primitive(IPrimitiveContext)

	// Set_expr_cast sets the _expr_cast rule contexts.
	Set_expr_cast(IExpr_castContext)

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
	_pow_op     IPow_opContext
	opRight     IExpr_aritContext
	_primitive  IPrimitiveContext
	_expr_cast  IExpr_castContext
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

func (s *Expr_aritContext) Get_primitive() IPrimitiveContext { return s._primitive }

func (s *Expr_aritContext) Get_expr_cast() IExpr_castContext { return s._expr_cast }

func (s *Expr_aritContext) SetOpLeft(v IExpr_aritContext) { s.opLeft = v }

func (s *Expr_aritContext) SetOpU(v IExpressionContext) { s.opU = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_pow_op(v IPow_opContext) { s._pow_op = v }

func (s *Expr_aritContext) SetOpRight(v IExpr_aritContext) { s.opRight = v }

func (s *Expr_aritContext) Set_primitive(v IPrimitiveContext) { s._primitive = v }

func (s *Expr_aritContext) Set_expr_cast(v IExpr_castContext) { s._expr_cast = v }

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

func (s *Expr_aritContext) Primitive() IPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveContext)
}

func (s *Expr_aritContext) Expr_cast() IExpr_castContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_castContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_castContext)
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
	_startState := 26
	p.EnterRecursionRule(localctx, 26, ProjectParserRULE_expr_arit, _p)
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
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(224)
			p.Match(ProjectParserSUB)
		}
		{
			p.SetState(225)

			var _x = p.Expression()

			localctx.(*Expr_aritContext).opU = _x
			localctx.(*Expr_aritContext)._expression = _x
		}
		localctx.(*Expr_aritContext).p = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpU().GetP(), "-", nil, true, localctx.(*Expr_aritContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_aritContext).GetOpU().GetStart().GetColumn())

	case 2:
		{
			p.SetState(228)

			var _x = p.Pow_op()

			localctx.(*Expr_aritContext)._pow_op = _x
		}
		{
			p.SetState(229)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(230)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opLeft = _x
		}
		{
			p.SetState(231)
			p.Match(ProjectParserCOMMA)
		}
		{
			p.SetState(232)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opRight = _x
		}
		{
			p.SetState(233)
			p.Match(ProjectParserRIGHT_PAR)
		}
		localctx.(*Expr_aritContext).p = Expression.NewOperation(localctx.(*Expr_aritContext).GetOpLeft().GetP(), localctx.(*Expr_aritContext).Get_pow_op().GetOp(), localctx.(*Expr_aritContext).GetOpRight().GetP(), false, localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetLine(), localctx.(*Expr_aritContext).Get_pow_op().GetStart().GetColumn())

	case 3:
		{
			p.SetState(236)

			var _x = p.Primitive()

			localctx.(*Expr_aritContext)._primitive = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitive().GetP()

	case 4:
		{
			p.SetState(239)

			var _x = p.Expr_cast()

			localctx.(*Expr_aritContext)._expr_cast = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expr_cast().GetP()

	case 5:
		{
			p.SetState(242)
			p.Match(ProjectParserLEFT_PAR)
		}
		{
			p.SetState(243)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(244)
			p.Match(ProjectParserRIGHT_PAR)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(259)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opLeft = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ProjectParserRULE_expr_arit)
				p.SetState(249)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(250)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(ProjectParserMUL-46))|(1<<(ProjectParserDIV-46))|(1<<(ProjectParserMOD-46)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(251)

					var _x = p.expr_arit(6)

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
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(255)

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
					p.SetState(256)

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

			}

		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 28, ProjectParserRULE_pow_op)

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

	p.SetState(272)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserRINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Match(ProjectParserRINTEGER)
		}
		{
			p.SetState(265)
			p.Match(ProjectParserHERITAGE)
		}
		{
			p.SetState(266)

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
			p.SetState(268)
			p.Match(ProjectParserRREAL)
		}
		{
			p.SetState(269)
			p.Match(ProjectParserHERITAGE)
		}
		{
			p.SetState(270)

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

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsExpr_logicContext differentiates from other interfaces.
	IsExpr_logicContext()
}

type Expr_logicContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	p       Abstract.Expression
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

func (s *Expr_logicContext) SetP(v Abstract.Expression) { s.p = v }

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
	p.EnterRule(localctx, 30, ProjectParserRULE_expr_logic)
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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserADMIRATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Match(ProjectParserADMIRATION)
		}
		{
			p.SetState(275)

			var _x = p.Expression()

			localctx.(*Expr_logicContext).opU = _x
		}
		localctx.(*Expr_logicContext).p = Expression.NewOperation(localctx.(*Expr_logicContext).GetOpU().GetP(), "!", nil, true, localctx.(*Expr_logicContext).GetOpU().GetStart().GetLine(), localctx.(*Expr_logicContext).GetOpU().GetStart().GetColumn())

	case ProjectParserRINTEGER, ProjectParserRREAL, ProjectParserINTEGER, ProjectParserFLOAT, ProjectParserCHAR, ProjectParserSTRING, ProjectParserBOOLEAN, ProjectParserID, ProjectParserLEFT_PAR, ProjectParserSUB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logicContext).opLeft = _x
		}
		{
			p.SetState(279)

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
			p.SetState(280)

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

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsExpr_castContext differentiates from other interfaces.
	IsExpr_castContext()
}

type Expr_castContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           Abstract.Expression
	_expression IExpressionContext
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

func (s *Expr_castContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_castContext) Get_data_type() IData_typeContext { return s._data_type }

func (s *Expr_castContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_castContext) Set_data_type(v IData_typeContext) { s._data_type = v }

func (s *Expr_castContext) GetP() Abstract.Expression { return s.p }

func (s *Expr_castContext) SetP(v Abstract.Expression) { s.p = v }

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
	p.EnterRule(localctx, 32, ProjectParserRULE_expr_cast)

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
		p.SetState(285)
		p.Match(ProjectParserLEFT_PAR)
	}
	{
		p.SetState(286)

		var _x = p.Expression()

		localctx.(*Expr_castContext)._expression = _x
	}
	{
		p.SetState(287)
		p.Match(ProjectParserRAS)
	}
	{
		p.SetState(288)

		var _x = p.Data_type()

		localctx.(*Expr_castContext)._data_type = _x
	}
	{
		p.SetState(289)
		p.Match(ProjectParserRIGHT_PAR)
	}

	if localctx.(*Expr_castContext).Get_data_type().GetData() == "i64" {
		localctx.(*Expr_castContext).p = Expression.NewCast(localctx.(*Expr_castContext).Get_expression().GetP(), SymbolTable.INTEGER)
	} else if localctx.(*Expr_castContext).Get_data_type().GetData() == "f64" {
		localctx.(*Expr_castContext).p = Expression.NewCast(localctx.(*Expr_castContext).Get_expression().GetP(), SymbolTable.FLOAT)
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
	p.EnterRule(localctx, 34, ProjectParserRULE_data_type)

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

	p.SetState(305)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserRINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)

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
			p.SetState(294)

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
			p.SetState(296)

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
			p.SetState(298)

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
			p.SetState(300)

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
			p.SetState(302)

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

	case ProjectParserEQUAL, ProjectParserSEMICOLON, ProjectParserRIGHT_PAR:
		p.EnterOuterAlt(localctx, 7)

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

	// SetP sets the p attribute.
	SetP(Abstract.Expression)

	// IsPrimitiveContext differentiates from other interfaces.
	IsPrimitiveContext()
}

type PrimitiveContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         Abstract.Expression
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
	p.EnterRule(localctx, 36, ProjectParserRULE_primitive)

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

	p.SetState(329)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ProjectParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)

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
			p.SetState(309)

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
			p.SetState(311)

			var _m = p.Match(ProjectParserSTRING)

			localctx.(*PrimitiveContext)._STRING = _m
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(312)
					p.Match(ProjectParserDOT)
				}
				{
					p.SetState(313)

					var _m = p.Match(ProjectParserTOSTRING)

					localctx.(*PrimitiveContext)._TOSTRING = _m
				}

			case 2:
				{
					p.SetState(314)
					p.Match(ProjectParserDOT)
				}
				{
					p.SetState(315)

					var _m = p.Match(ProjectParserTOOWNED)

					localctx.(*PrimitiveContext)._TOOWNED = _m
				}

			}
			{
				p.SetState(318)
				p.Match(ProjectParserLEFT_PAR)
			}
			{
				p.SetState(319)
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
		} else {
			localctx.(*PrimitiveContext).p = Expression.NewPrimitive(str, SymbolTable.STR, (func() int {
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
			p.SetState(323)

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
			p.SetState(325)

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
			p.SetState(327)

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

	case 13:
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
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
