// Code generated from ProjectLexer.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 77, 508,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 36, 3, 36, 6, 36, 359, 10, 36, 13, 36, 14, 36, 360, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 6, 37, 372, 10, 37, 13, 37,
	14, 37, 373, 3, 37, 3, 37, 3, 38, 6, 38, 379, 10, 38, 13, 38, 14, 38, 380,
	3, 39, 6, 39, 384, 10, 39, 13, 39, 14, 39, 385, 3, 39, 3, 39, 6, 39, 390,
	10, 39, 13, 39, 14, 39, 391, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41,
	7, 41, 400, 10, 41, 12, 41, 14, 41, 403, 11, 41, 3, 41, 3, 41, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 416, 10, 42,
	3, 43, 3, 43, 7, 43, 420, 10, 43, 12, 43, 14, 43, 423, 11, 43, 3, 44, 3,
	44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49,
	3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3,
	53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58,
	3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3,
	63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67,
	3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3,
	72, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76,
	6, 76, 500, 10, 76, 13, 76, 14, 76, 501, 3, 76, 3, 76, 3, 77, 3, 77, 3,
	77, 2, 2, 78, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73,
	38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91,
	47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55,
	109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63,
	125, 64, 127, 65, 129, 66, 131, 67, 133, 68, 135, 69, 137, 70, 139, 71,
	141, 72, 143, 73, 145, 74, 147, 75, 149, 76, 151, 77, 153, 2, 3, 2, 10,
	3, 2, 49, 49, 3, 2, 12, 12, 3, 2, 50, 59, 3, 2, 36, 36, 7, 2, 67, 92, 97,
	97, 99, 124, 211, 211, 243, 243, 8, 2, 50, 59, 67, 92, 97, 97, 99, 124,
	211, 211, 243, 243, 5, 2, 11, 12, 15, 15, 34, 34, 9, 2, 34, 35, 37, 37,
	45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 515, 2, 3, 3, 2, 2, 2, 2, 5,
	3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2,
	105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2,
	2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119,
	3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2,
	2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3,
	2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2,
	141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2,
	2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 3, 155, 3, 2, 2, 2, 5, 163,
	3, 2, 2, 2, 7, 169, 3, 2, 2, 2, 9, 173, 3, 2, 2, 2, 11, 177, 3, 2, 2, 2,
	13, 184, 3, 2, 2, 2, 15, 188, 3, 2, 2, 2, 17, 192, 3, 2, 2, 2, 19, 197,
	3, 2, 2, 2, 21, 202, 3, 2, 2, 2, 23, 207, 3, 2, 2, 2, 25, 210, 3, 2, 2,
	2, 27, 213, 3, 2, 2, 2, 29, 218, 3, 2, 2, 2, 31, 224, 3, 2, 2, 2, 33, 230,
	3, 2, 2, 2, 35, 235, 3, 2, 2, 2, 37, 239, 3, 2, 2, 2, 39, 242, 3, 2, 2,
	2, 41, 248, 3, 2, 2, 2, 43, 257, 3, 2, 2, 2, 45, 264, 3, 2, 2, 2, 47, 268,
	3, 2, 2, 2, 49, 273, 3, 2, 2, 2, 51, 283, 3, 2, 2, 2, 53, 292, 3, 2, 2,
	2, 55, 297, 3, 2, 2, 2, 57, 300, 3, 2, 2, 2, 59, 307, 3, 2, 2, 2, 61, 311,
	3, 2, 2, 2, 63, 315, 3, 2, 2, 2, 65, 319, 3, 2, 2, 2, 67, 330, 3, 2, 2,
	2, 69, 349, 3, 2, 2, 2, 71, 354, 3, 2, 2, 2, 73, 367, 3, 2, 2, 2, 75, 378,
	3, 2, 2, 2, 77, 383, 3, 2, 2, 2, 79, 393, 3, 2, 2, 2, 81, 397, 3, 2, 2,
	2, 83, 415, 3, 2, 2, 2, 85, 417, 3, 2, 2, 2, 87, 424, 3, 2, 2, 2, 89, 426,
	3, 2, 2, 2, 91, 428, 3, 2, 2, 2, 93, 430, 3, 2, 2, 2, 95, 432, 3, 2, 2,
	2, 97, 434, 3, 2, 2, 2, 99, 436, 3, 2, 2, 2, 101, 438, 3, 2, 2, 2, 103,
	441, 3, 2, 2, 2, 105, 444, 3, 2, 2, 2, 107, 447, 3, 2, 2, 2, 109, 449,
	3, 2, 2, 2, 111, 451, 3, 2, 2, 2, 113, 453, 3, 2, 2, 2, 115, 455, 3, 2,
	2, 2, 117, 457, 3, 2, 2, 2, 119, 459, 3, 2, 2, 2, 121, 462, 3, 2, 2, 2,
	123, 465, 3, 2, 2, 2, 125, 467, 3, 2, 2, 2, 127, 469, 3, 2, 2, 2, 129,
	472, 3, 2, 2, 2, 131, 475, 3, 2, 2, 2, 133, 477, 3, 2, 2, 2, 135, 479,
	3, 2, 2, 2, 137, 481, 3, 2, 2, 2, 139, 483, 3, 2, 2, 2, 141, 485, 3, 2,
	2, 2, 143, 488, 3, 2, 2, 2, 145, 491, 3, 2, 2, 2, 147, 493, 3, 2, 2, 2,
	149, 496, 3, 2, 2, 2, 151, 499, 3, 2, 2, 2, 153, 505, 3, 2, 2, 2, 155,
	156, 7, 114, 2, 2, 156, 157, 7, 116, 2, 2, 157, 158, 7, 107, 2, 2, 158,
	159, 7, 112, 2, 2, 159, 160, 7, 118, 2, 2, 160, 161, 7, 110, 2, 2, 161,
	162, 7, 112, 2, 2, 162, 4, 3, 2, 2, 2, 163, 164, 7, 114, 2, 2, 164, 165,
	7, 116, 2, 2, 165, 166, 7, 107, 2, 2, 166, 167, 7, 112, 2, 2, 167, 168,
	7, 118, 2, 2, 168, 6, 3, 2, 2, 2, 169, 170, 7, 110, 2, 2, 170, 171, 7,
	103, 2, 2, 171, 172, 7, 118, 2, 2, 172, 8, 3, 2, 2, 2, 173, 174, 7, 111,
	2, 2, 174, 175, 7, 119, 2, 2, 175, 176, 7, 118, 2, 2, 176, 10, 3, 2, 2,
	2, 177, 178, 7, 85, 2, 2, 178, 179, 7, 118, 2, 2, 179, 180, 7, 116, 2,
	2, 180, 181, 7, 107, 2, 2, 181, 182, 7, 112, 2, 2, 182, 183, 7, 105, 2,
	2, 183, 12, 3, 2, 2, 2, 184, 185, 7, 107, 2, 2, 185, 186, 7, 56, 2, 2,
	186, 187, 7, 54, 2, 2, 187, 14, 3, 2, 2, 2, 188, 189, 7, 104, 2, 2, 189,
	190, 7, 56, 2, 2, 190, 191, 7, 54, 2, 2, 191, 16, 3, 2, 2, 2, 192, 193,
	7, 100, 2, 2, 193, 194, 7, 113, 2, 2, 194, 195, 7, 113, 2, 2, 195, 196,
	7, 110, 2, 2, 196, 18, 3, 2, 2, 2, 197, 198, 7, 101, 2, 2, 198, 199, 7,
	106, 2, 2, 199, 200, 7, 99, 2, 2, 200, 201, 7, 116, 2, 2, 201, 20, 3, 2,
	2, 2, 202, 203, 7, 40, 2, 2, 203, 204, 7, 117, 2, 2, 204, 205, 7, 118,
	2, 2, 205, 206, 7, 116, 2, 2, 206, 22, 3, 2, 2, 2, 207, 208, 7, 99, 2,
	2, 208, 209, 7, 117, 2, 2, 209, 24, 3, 2, 2, 2, 210, 211, 7, 107, 2, 2,
	211, 212, 7, 104, 2, 2, 212, 26, 3, 2, 2, 2, 213, 214, 7, 103, 2, 2, 214,
	215, 7, 110, 2, 2, 215, 216, 7, 117, 2, 2, 216, 217, 7, 103, 2, 2, 217,
	28, 3, 2, 2, 2, 218, 219, 7, 111, 2, 2, 219, 220, 7, 99, 2, 2, 220, 221,
	7, 118, 2, 2, 221, 222, 7, 101, 2, 2, 222, 223, 7, 106, 2, 2, 223, 30,
	3, 2, 2, 2, 224, 225, 7, 121, 2, 2, 225, 226, 7, 106, 2, 2, 226, 227, 7,
	107, 2, 2, 227, 228, 7, 110, 2, 2, 228, 229, 7, 103, 2, 2, 229, 32, 3,
	2, 2, 2, 230, 231, 7, 110, 2, 2, 231, 232, 7, 113, 2, 2, 232, 233, 7, 113,
	2, 2, 233, 234, 7, 114, 2, 2, 234, 34, 3, 2, 2, 2, 235, 236, 7, 104, 2,
	2, 236, 237, 7, 113, 2, 2, 237, 238, 7, 116, 2, 2, 238, 36, 3, 2, 2, 2,
	239, 240, 7, 107, 2, 2, 240, 241, 7, 112, 2, 2, 241, 38, 3, 2, 2, 2, 242,
	243, 7, 100, 2, 2, 243, 244, 7, 116, 2, 2, 244, 245, 7, 103, 2, 2, 245,
	246, 7, 99, 2, 2, 246, 247, 7, 109, 2, 2, 247, 40, 3, 2, 2, 2, 248, 249,
	7, 101, 2, 2, 249, 250, 7, 113, 2, 2, 250, 251, 7, 112, 2, 2, 251, 252,
	7, 118, 2, 2, 252, 253, 7, 107, 2, 2, 253, 254, 7, 112, 2, 2, 254, 255,
	7, 119, 2, 2, 255, 256, 7, 103, 2, 2, 256, 42, 3, 2, 2, 2, 257, 258, 7,
	116, 2, 2, 258, 259, 7, 103, 2, 2, 259, 260, 7, 118, 2, 2, 260, 261, 7,
	119, 2, 2, 261, 262, 7, 116, 2, 2, 262, 263, 7, 112, 2, 2, 263, 44, 3,
	2, 2, 2, 264, 265, 7, 114, 2, 2, 265, 266, 7, 113, 2, 2, 266, 267, 7, 121,
	2, 2, 267, 46, 3, 2, 2, 2, 268, 269, 7, 114, 2, 2, 269, 270, 7, 113, 2,
	2, 270, 271, 7, 121, 2, 2, 271, 272, 7, 104, 2, 2, 272, 48, 3, 2, 2, 2,
	273, 274, 7, 118, 2, 2, 274, 275, 7, 113, 2, 2, 275, 276, 7, 97, 2, 2,
	276, 277, 7, 117, 2, 2, 277, 278, 7, 118, 2, 2, 278, 279, 7, 116, 2, 2,
	279, 280, 7, 107, 2, 2, 280, 281, 7, 112, 2, 2, 281, 282, 7, 105, 2, 2,
	282, 50, 3, 2, 2, 2, 283, 284, 7, 118, 2, 2, 284, 285, 7, 113, 2, 2, 285,
	286, 7, 97, 2, 2, 286, 287, 7, 113, 2, 2, 287, 288, 7, 121, 2, 2, 288,
	289, 7, 112, 2, 2, 289, 290, 7, 103, 2, 2, 290, 291, 7, 102, 2, 2, 291,
	52, 3, 2, 2, 2, 292, 293, 7, 111, 2, 2, 293, 294, 7, 99, 2, 2, 294, 295,
	7, 107, 2, 2, 295, 296, 7, 112, 2, 2, 296, 54, 3, 2, 2, 2, 297, 298, 7,
	104, 2, 2, 298, 299, 7, 112, 2, 2, 299, 56, 3, 2, 2, 2, 300, 301, 7, 120,
	2, 2, 301, 302, 7, 103, 2, 2, 302, 303, 7, 101, 2, 2, 303, 304, 7, 118,
	2, 2, 304, 305, 7, 113, 2, 2, 305, 306, 7, 116, 2, 2, 306, 58, 3, 2, 2,
	2, 307, 308, 7, 120, 2, 2, 308, 309, 7, 103, 2, 2, 309, 310, 7, 101, 2,
	2, 310, 60, 3, 2, 2, 2, 311, 312, 7, 88, 2, 2, 312, 313, 7, 103, 2, 2,
	313, 314, 7, 101, 2, 2, 314, 62, 3, 2, 2, 2, 315, 316, 7, 112, 2, 2, 316,
	317, 7, 103, 2, 2, 317, 318, 7, 121, 2, 2, 318, 64, 3, 2, 2, 2, 319, 320,
	7, 88, 2, 2, 320, 321, 7, 103, 2, 2, 321, 322, 7, 101, 2, 2, 322, 323,
	7, 60, 2, 2, 323, 324, 7, 60, 2, 2, 324, 325, 7, 112, 2, 2, 325, 326, 7,
	103, 2, 2, 326, 327, 7, 121, 2, 2, 327, 328, 7, 42, 2, 2, 328, 329, 7,
	43, 2, 2, 329, 66, 3, 2, 2, 2, 330, 331, 7, 88, 2, 2, 331, 332, 7, 103,
	2, 2, 332, 333, 7, 101, 2, 2, 333, 334, 7, 60, 2, 2, 334, 335, 7, 60, 2,
	2, 335, 336, 7, 121, 2, 2, 336, 337, 7, 107, 2, 2, 337, 338, 7, 118, 2,
	2, 338, 339, 7, 106, 2, 2, 339, 340, 7, 97, 2, 2, 340, 341, 7, 101, 2,
	2, 341, 342, 7, 99, 2, 2, 342, 343, 7, 114, 2, 2, 343, 344, 7, 99, 2, 2,
	344, 345, 7, 101, 2, 2, 345, 346, 7, 107, 2, 2, 346, 347, 7, 118, 2, 2,
	347, 348, 7, 123, 2, 2, 348, 68, 3, 2, 2, 2, 349, 350, 7, 114, 2, 2, 350,
	351, 7, 119, 2, 2, 351, 352, 7, 117, 2, 2, 352, 353, 7, 106, 2, 2, 353,
	70, 3, 2, 2, 2, 354, 355, 7, 49, 2, 2, 355, 356, 7, 44, 2, 2, 356, 358,
	3, 2, 2, 2, 357, 359, 10, 2, 2, 2, 358, 357, 3, 2, 2, 2, 359, 360, 3, 2,
	2, 2, 360, 358, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2,
	362, 363, 7, 44, 2, 2, 363, 364, 7, 49, 2, 2, 364, 365, 3, 2, 2, 2, 365,
	366, 8, 36, 2, 2, 366, 72, 3, 2, 2, 2, 367, 368, 7, 49, 2, 2, 368, 369,
	7, 49, 2, 2, 369, 371, 3, 2, 2, 2, 370, 372, 10, 3, 2, 2, 371, 370, 3,
	2, 2, 2, 372, 373, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 373, 374, 3, 2, 2,
	2, 374, 375, 3, 2, 2, 2, 375, 376, 8, 37, 2, 2, 376, 74, 3, 2, 2, 2, 377,
	379, 9, 4, 2, 2, 378, 377, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 378,
	3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 76, 3, 2, 2, 2, 382, 384, 9, 4,
	2, 2, 383, 382, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2,
	385, 386, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 389, 7, 48, 2, 2, 388,
	390, 9, 4, 2, 2, 389, 388, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 389,
	3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 78, 3, 2, 2, 2, 393, 394, 7, 41,
	2, 2, 394, 395, 10, 5, 2, 2, 395, 396, 7, 41, 2, 2, 396, 80, 3, 2, 2, 2,
	397, 401, 7, 36, 2, 2, 398, 400, 10, 5, 2, 2, 399, 398, 3, 2, 2, 2, 400,
	403, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2, 401, 402, 3, 2, 2, 2, 402, 404,
	3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 404, 405, 7, 36, 2, 2, 405, 82, 3, 2,
	2, 2, 406, 407, 7, 118, 2, 2, 407, 408, 7, 116, 2, 2, 408, 409, 7, 119,
	2, 2, 409, 416, 7, 103, 2, 2, 410, 411, 7, 104, 2, 2, 411, 412, 7, 99,
	2, 2, 412, 413, 7, 110, 2, 2, 413, 414, 7, 117, 2, 2, 414, 416, 7, 103,
	2, 2, 415, 406, 3, 2, 2, 2, 415, 410, 3, 2, 2, 2, 416, 84, 3, 2, 2, 2,
	417, 421, 9, 6, 2, 2, 418, 420, 9, 7, 2, 2, 419, 418, 3, 2, 2, 2, 420,
	423, 3, 2, 2, 2, 421, 419, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 86, 3,
	2, 2, 2, 423, 421, 3, 2, 2, 2, 424, 425, 7, 63, 2, 2, 425, 88, 3, 2, 2,
	2, 426, 427, 7, 48, 2, 2, 427, 90, 3, 2, 2, 2, 428, 429, 7, 61, 2, 2, 429,
	92, 3, 2, 2, 2, 430, 431, 7, 46, 2, 2, 431, 94, 3, 2, 2, 2, 432, 433, 7,
	60, 2, 2, 433, 96, 3, 2, 2, 2, 434, 435, 7, 35, 2, 2, 435, 98, 3, 2, 2,
	2, 436, 437, 7, 40, 2, 2, 437, 100, 3, 2, 2, 2, 438, 439, 7, 60, 2, 2,
	439, 440, 7, 60, 2, 2, 440, 102, 3, 2, 2, 2, 441, 442, 7, 47, 2, 2, 442,
	443, 7, 64, 2, 2, 443, 104, 3, 2, 2, 2, 444, 445, 7, 48, 2, 2, 445, 446,
	7, 48, 2, 2, 446, 106, 3, 2, 2, 2, 447, 448, 7, 42, 2, 2, 448, 108, 3,
	2, 2, 2, 449, 450, 7, 43, 2, 2, 450, 110, 3, 2, 2, 2, 451, 452, 7, 125,
	2, 2, 452, 112, 3, 2, 2, 2, 453, 454, 7, 127, 2, 2, 454, 114, 3, 2, 2,
	2, 455, 456, 7, 93, 2, 2, 456, 116, 3, 2, 2, 2, 457, 458, 7, 95, 2, 2,
	458, 118, 3, 2, 2, 2, 459, 460, 7, 63, 2, 2, 460, 461, 7, 63, 2, 2, 461,
	120, 3, 2, 2, 2, 462, 463, 7, 35, 2, 2, 463, 464, 7, 63, 2, 2, 464, 122,
	3, 2, 2, 2, 465, 466, 7, 64, 2, 2, 466, 124, 3, 2, 2, 2, 467, 468, 7, 62,
	2, 2, 468, 126, 3, 2, 2, 2, 469, 470, 7, 64, 2, 2, 470, 471, 7, 63, 2,
	2, 471, 128, 3, 2, 2, 2, 472, 473, 7, 62, 2, 2, 473, 474, 7, 63, 2, 2,
	474, 130, 3, 2, 2, 2, 475, 476, 7, 44, 2, 2, 476, 132, 3, 2, 2, 2, 477,
	478, 7, 49, 2, 2, 478, 134, 3, 2, 2, 2, 479, 480, 7, 39, 2, 2, 480, 136,
	3, 2, 2, 2, 481, 482, 7, 45, 2, 2, 482, 138, 3, 2, 2, 2, 483, 484, 7, 47,
	2, 2, 484, 140, 3, 2, 2, 2, 485, 486, 7, 40, 2, 2, 486, 487, 7, 40, 2,
	2, 487, 142, 3, 2, 2, 2, 488, 489, 7, 126, 2, 2, 489, 490, 7, 126, 2, 2,
	490, 144, 3, 2, 2, 2, 491, 492, 7, 126, 2, 2, 492, 146, 3, 2, 2, 2, 493,
	494, 7, 63, 2, 2, 494, 495, 7, 64, 2, 2, 495, 148, 3, 2, 2, 2, 496, 497,
	7, 97, 2, 2, 497, 150, 3, 2, 2, 2, 498, 500, 9, 8, 2, 2, 499, 498, 3, 2,
	2, 2, 500, 501, 3, 2, 2, 2, 501, 499, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2,
	502, 503, 3, 2, 2, 2, 503, 504, 8, 76, 2, 2, 504, 152, 3, 2, 2, 2, 505,
	506, 7, 94, 2, 2, 506, 507, 9, 9, 2, 2, 507, 154, 3, 2, 2, 2, 12, 2, 360,
	373, 380, 385, 391, 401, 415, 421, 501, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'println'", "'print'", "'let'", "'mut'", "'String'", "'i64'", "'f64'",
	"'bool'", "'char'", "'&str'", "'as'", "'if'", "'else'", "'match'", "'while'",
	"'loop'", "'for'", "'in'", "'break'", "'continue'", "'return'", "'pow'",
	"'powf'", "'to_string'", "'to_owned'", "'main'", "'fn'", "'vector'", "'vec'",
	"'Vec'", "'new'", "'Vec::new()'", "'Vec::with_capacity'", "'push'", "",
	"", "", "", "", "", "", "", "'='", "'.'", "';'", "','", "':'", "'!'", "'&'",
	"'::'", "'->'", "'..'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'=='",
	"'!='", "'>'", "'<'", "'>='", "'<='", "'*'", "'/'", "'%'", "'+'", "'-'",
	"'&&'", "'||'", "'|'", "'=>'", "'_'",
}

var lexerSymbolicNames = []string{
	"", "PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL",
	"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE",
	"RLOOP", "RFOR", "RIN", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF",
	"TOSTRING", "TOOWNED", "RMAIN", "RFN", "RVECTOR", "RVEC", "RVECMayus",
	"RNEW", "REVECTORNEW", "REVECCAPACITY", "RPUSH", "MULTILINE", "INLINE",
	"INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT",
	"SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "HERITAGE", "ARROW",
	"RANGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET",
	"RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN",
	"GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB",
	"AND", "OR", "PIPE", "EQUAL_ARROW", "UNDERSCORE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"PRINTLN", "PRINT", "DECLARAR", "MUT", "RSTRING", "RINTEGER", "RREAL",
	"RBOOLEAN", "RCHAR", "RSTR", "RAS", "RIF", "RELSE", "RMATCH", "RWHILE",
	"RLOOP", "RFOR", "RIN", "RBREAK", "RCONTINUE", "RRETURN", "POWI", "POWF",
	"TOSTRING", "TOOWNED", "RMAIN", "RFN", "RVECTOR", "RVEC", "RVECMayus",
	"RNEW", "REVECTORNEW", "REVECCAPACITY", "RPUSH", "MULTILINE", "INLINE",
	"INTEGER", "FLOAT", "CHAR", "STRING", "BOOLEAN", "ID", "EQUAL", "DOT",
	"SEMICOLON", "COMMA", "COLON", "ADMIRATION", "REFERENCE", "HERITAGE", "ARROW",
	"RANGE", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY", "LEFT_BRACKET",
	"RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN", "LESS_THAN",
	"GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "MOD", "ADD", "SUB",
	"AND", "OR", "PIPE", "EQUAL_ARROW", "UNDERSCORE", "WHITESPACE", "ESC_SEQ",
}

type ProjectLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewProjectLexer(input antlr.CharStream) *ProjectLexer {

	l := new(ProjectLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "ProjectLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ProjectLexer tokens.
const (
	ProjectLexerPRINTLN           = 1
	ProjectLexerPRINT             = 2
	ProjectLexerDECLARAR          = 3
	ProjectLexerMUT               = 4
	ProjectLexerRSTRING           = 5
	ProjectLexerRINTEGER          = 6
	ProjectLexerRREAL             = 7
	ProjectLexerRBOOLEAN          = 8
	ProjectLexerRCHAR             = 9
	ProjectLexerRSTR              = 10
	ProjectLexerRAS               = 11
	ProjectLexerRIF               = 12
	ProjectLexerRELSE             = 13
	ProjectLexerRMATCH            = 14
	ProjectLexerRWHILE            = 15
	ProjectLexerRLOOP             = 16
	ProjectLexerRFOR              = 17
	ProjectLexerRIN               = 18
	ProjectLexerRBREAK            = 19
	ProjectLexerRCONTINUE         = 20
	ProjectLexerRRETURN           = 21
	ProjectLexerPOWI              = 22
	ProjectLexerPOWF              = 23
	ProjectLexerTOSTRING          = 24
	ProjectLexerTOOWNED           = 25
	ProjectLexerRMAIN             = 26
	ProjectLexerRFN               = 27
	ProjectLexerRVECTOR           = 28
	ProjectLexerRVEC              = 29
	ProjectLexerRVECMayus         = 30
	ProjectLexerRNEW              = 31
	ProjectLexerREVECTORNEW       = 32
	ProjectLexerREVECCAPACITY     = 33
	ProjectLexerRPUSH             = 34
	ProjectLexerMULTILINE         = 35
	ProjectLexerINLINE            = 36
	ProjectLexerINTEGER           = 37
	ProjectLexerFLOAT             = 38
	ProjectLexerCHAR              = 39
	ProjectLexerSTRING            = 40
	ProjectLexerBOOLEAN           = 41
	ProjectLexerID                = 42
	ProjectLexerEQUAL             = 43
	ProjectLexerDOT               = 44
	ProjectLexerSEMICOLON         = 45
	ProjectLexerCOMMA             = 46
	ProjectLexerCOLON             = 47
	ProjectLexerADMIRATION        = 48
	ProjectLexerREFERENCE         = 49
	ProjectLexerHERITAGE          = 50
	ProjectLexerARROW             = 51
	ProjectLexerRANGE             = 52
	ProjectLexerLEFT_PAR          = 53
	ProjectLexerRIGHT_PAR         = 54
	ProjectLexerLEFT_KEY          = 55
	ProjectLexerRIGHT_KEY         = 56
	ProjectLexerLEFT_BRACKET      = 57
	ProjectLexerRIGHT_BRACKET     = 58
	ProjectLexerEQUEAL_EQUAL      = 59
	ProjectLexerNOTEQUAL          = 60
	ProjectLexerGREATER_THAN      = 61
	ProjectLexerLESS_THAN         = 62
	ProjectLexerGREATER_EQUALTHAN = 63
	ProjectLexerLESS_EQUEALTHAN   = 64
	ProjectLexerMUL               = 65
	ProjectLexerDIV               = 66
	ProjectLexerMOD               = 67
	ProjectLexerADD               = 68
	ProjectLexerSUB               = 69
	ProjectLexerAND               = 70
	ProjectLexerOR                = 71
	ProjectLexerPIPE              = 72
	ProjectLexerEQUAL_ARROW       = 73
	ProjectLexerUNDERSCORE        = 74
	ProjectLexerWHITESPACE        = 75
)
