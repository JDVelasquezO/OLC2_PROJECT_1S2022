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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 45, 303,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 6, 14,
	180, 10, 14, 13, 14, 14, 14, 181, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 6, 15, 193, 10, 15, 13, 15, 14, 15, 194, 3, 15,
	3, 15, 3, 16, 6, 16, 200, 10, 16, 13, 16, 14, 16, 201, 3, 17, 6, 17, 205,
	10, 17, 13, 17, 14, 17, 206, 3, 17, 3, 17, 6, 17, 211, 10, 17, 13, 17,
	14, 17, 212, 3, 18, 3, 18, 7, 18, 217, 10, 18, 12, 18, 14, 18, 220, 11,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 5, 19, 233, 10, 19, 3, 20, 3, 20, 7, 20, 237, 10, 20, 12, 20, 14,
	20, 240, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33,
	3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 44, 6, 44, 295, 10, 44, 13, 44, 14, 44, 296, 3, 44, 3,
	44, 3, 45, 3, 45, 3, 45, 2, 2, 46, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44,
	87, 45, 89, 2, 3, 2, 10, 3, 2, 49, 49, 3, 2, 12, 12, 3, 2, 50, 59, 3, 2,
	36, 36, 7, 2, 67, 92, 97, 97, 99, 124, 211, 211, 243, 243, 7, 2, 50, 59,
	67, 92, 99, 124, 211, 211, 243, 243, 5, 2, 11, 12, 15, 15, 34, 34, 9, 2,
	34, 35, 37, 37, 45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 310, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3,
	2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65,
	3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2,
	73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2,
	2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2,
	2, 3, 91, 3, 2, 2, 2, 5, 101, 3, 2, 2, 2, 7, 109, 3, 2, 2, 2, 9, 119, 3,
	2, 2, 2, 11, 128, 3, 2, 2, 2, 13, 135, 3, 2, 2, 2, 15, 143, 3, 2, 2, 2,
	17, 148, 3, 2, 2, 2, 19, 156, 3, 2, 2, 2, 21, 159, 3, 2, 2, 2, 23, 168,
	3, 2, 2, 2, 25, 172, 3, 2, 2, 2, 27, 175, 3, 2, 2, 2, 29, 188, 3, 2, 2,
	2, 31, 199, 3, 2, 2, 2, 33, 204, 3, 2, 2, 2, 35, 214, 3, 2, 2, 2, 37, 232,
	3, 2, 2, 2, 39, 234, 3, 2, 2, 2, 41, 241, 3, 2, 2, 2, 43, 245, 3, 2, 2,
	2, 45, 247, 3, 2, 2, 2, 47, 249, 3, 2, 2, 2, 49, 251, 3, 2, 2, 2, 51, 253,
	3, 2, 2, 2, 53, 255, 3, 2, 2, 2, 55, 257, 3, 2, 2, 2, 57, 259, 3, 2, 2,
	2, 59, 261, 3, 2, 2, 2, 61, 263, 3, 2, 2, 2, 63, 265, 3, 2, 2, 2, 65, 267,
	3, 2, 2, 2, 67, 270, 3, 2, 2, 2, 69, 273, 3, 2, 2, 2, 71, 275, 3, 2, 2,
	2, 73, 277, 3, 2, 2, 2, 75, 280, 3, 2, 2, 2, 77, 283, 3, 2, 2, 2, 79, 285,
	3, 2, 2, 2, 81, 287, 3, 2, 2, 2, 83, 289, 3, 2, 2, 2, 85, 291, 3, 2, 2,
	2, 87, 294, 3, 2, 2, 2, 89, 300, 3, 2, 2, 2, 91, 92, 7, 85, 2, 2, 92, 93,
	7, 103, 2, 2, 93, 94, 7, 112, 2, 2, 94, 95, 7, 118, 2, 2, 95, 96, 7, 103,
	2, 2, 96, 97, 7, 112, 2, 2, 97, 98, 7, 101, 2, 2, 98, 99, 7, 107, 2, 2,
	99, 100, 7, 99, 2, 2, 100, 4, 3, 2, 2, 2, 101, 102, 7, 101, 2, 2, 102,
	103, 7, 113, 2, 2, 103, 104, 7, 112, 2, 2, 104, 105, 7, 117, 2, 2, 105,
	106, 7, 113, 2, 2, 106, 107, 7, 110, 2, 2, 107, 108, 7, 99, 2, 2, 108,
	6, 3, 2, 2, 2, 109, 110, 7, 101, 2, 2, 110, 111, 7, 113, 2, 2, 111, 112,
	7, 112, 2, 2, 112, 113, 7, 117, 2, 2, 113, 114, 7, 113, 2, 2, 114, 115,
	7, 110, 2, 2, 115, 116, 7, 99, 2, 2, 116, 117, 7, 110, 2, 2, 117, 118,
	7, 112, 2, 2, 118, 8, 3, 2, 2, 2, 119, 120, 7, 102, 2, 2, 120, 121, 7,
	103, 2, 2, 121, 122, 7, 101, 2, 2, 122, 123, 7, 110, 2, 2, 123, 124, 7,
	99, 2, 2, 124, 125, 7, 116, 2, 2, 125, 126, 7, 99, 2, 2, 126, 127, 7, 116,
	2, 2, 127, 10, 3, 2, 2, 2, 128, 129, 7, 85, 2, 2, 129, 130, 7, 118, 2,
	2, 130, 131, 7, 116, 2, 2, 131, 132, 7, 107, 2, 2, 132, 133, 7, 112, 2,
	2, 133, 134, 7, 105, 2, 2, 134, 12, 3, 2, 2, 2, 135, 136, 7, 75, 2, 2,
	136, 137, 7, 112, 2, 2, 137, 138, 7, 118, 2, 2, 138, 139, 7, 103, 2, 2,
	139, 140, 7, 105, 2, 2, 140, 141, 7, 103, 2, 2, 141, 142, 7, 116, 2, 2,
	142, 14, 3, 2, 2, 2, 143, 144, 7, 84, 2, 2, 144, 145, 7, 103, 2, 2, 145,
	146, 7, 99, 2, 2, 146, 147, 7, 110, 2, 2, 147, 16, 3, 2, 2, 2, 148, 149,
	7, 68, 2, 2, 149, 150, 7, 113, 2, 2, 150, 151, 7, 113, 2, 2, 151, 152,
	7, 110, 2, 2, 152, 153, 7, 103, 2, 2, 153, 154, 7, 99, 2, 2, 154, 155,
	7, 112, 2, 2, 155, 18, 3, 2, 2, 2, 156, 157, 7, 75, 2, 2, 157, 158, 7,
	104, 2, 2, 158, 20, 3, 2, 2, 2, 159, 160, 7, 103, 2, 2, 160, 161, 7, 112,
	2, 2, 161, 162, 7, 118, 2, 2, 162, 163, 7, 113, 2, 2, 163, 164, 7, 112,
	2, 2, 164, 165, 7, 101, 2, 2, 165, 166, 7, 103, 2, 2, 166, 167, 7, 117,
	2, 2, 167, 22, 3, 2, 2, 2, 168, 169, 7, 67, 2, 2, 169, 170, 7, 80, 2, 2,
	170, 171, 7, 70, 2, 2, 171, 24, 3, 2, 2, 2, 172, 173, 7, 81, 2, 2, 173,
	174, 7, 84, 2, 2, 174, 26, 3, 2, 2, 2, 175, 176, 7, 42, 2, 2, 176, 177,
	7, 44, 2, 2, 177, 179, 3, 2, 2, 2, 178, 180, 10, 2, 2, 2, 179, 178, 3,
	2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2,
	2, 182, 183, 3, 2, 2, 2, 183, 184, 7, 44, 2, 2, 184, 185, 7, 43, 2, 2,
	185, 186, 3, 2, 2, 2, 186, 187, 8, 14, 2, 2, 187, 28, 3, 2, 2, 2, 188,
	189, 7, 49, 2, 2, 189, 190, 7, 49, 2, 2, 190, 192, 3, 2, 2, 2, 191, 193,
	10, 3, 2, 2, 192, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 192, 3, 2,
	2, 2, 194, 195, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 197, 8, 15, 2, 2,
	197, 30, 3, 2, 2, 2, 198, 200, 9, 4, 2, 2, 199, 198, 3, 2, 2, 2, 200, 201,
	3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 32, 3, 2,
	2, 2, 203, 205, 9, 4, 2, 2, 204, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2,
	206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208,
	210, 7, 48, 2, 2, 209, 211, 9, 4, 2, 2, 210, 209, 3, 2, 2, 2, 211, 212,
	3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 34, 3, 2,
	2, 2, 214, 218, 7, 41, 2, 2, 215, 217, 10, 5, 2, 2, 216, 215, 3, 2, 2,
	2, 217, 220, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219,
	221, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 222, 7, 41, 2, 2, 222, 36,
	3, 2, 2, 2, 223, 224, 7, 118, 2, 2, 224, 225, 7, 116, 2, 2, 225, 226, 7,
	119, 2, 2, 226, 233, 7, 103, 2, 2, 227, 228, 7, 104, 2, 2, 228, 229, 7,
	99, 2, 2, 229, 230, 7, 110, 2, 2, 230, 231, 7, 117, 2, 2, 231, 233, 7,
	103, 2, 2, 232, 223, 3, 2, 2, 2, 232, 227, 3, 2, 2, 2, 233, 38, 3, 2, 2,
	2, 234, 238, 9, 6, 2, 2, 235, 237, 9, 7, 2, 2, 236, 235, 3, 2, 2, 2, 237,
	240, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 40, 3,
	2, 2, 2, 240, 238, 3, 2, 2, 2, 241, 242, 7, 107, 2, 2, 242, 243, 7, 112,
	2, 2, 243, 244, 7, 118, 2, 2, 244, 42, 3, 2, 2, 2, 245, 246, 7, 63, 2,
	2, 246, 44, 3, 2, 2, 2, 247, 248, 7, 48, 2, 2, 248, 46, 3, 2, 2, 2, 249,
	250, 7, 61, 2, 2, 250, 48, 3, 2, 2, 2, 251, 252, 7, 46, 2, 2, 252, 50,
	3, 2, 2, 2, 253, 254, 7, 60, 2, 2, 254, 52, 3, 2, 2, 2, 255, 256, 7, 42,
	2, 2, 256, 54, 3, 2, 2, 2, 257, 258, 7, 43, 2, 2, 258, 56, 3, 2, 2, 2,
	259, 260, 7, 125, 2, 2, 260, 58, 3, 2, 2, 2, 261, 262, 7, 127, 2, 2, 262,
	60, 3, 2, 2, 2, 263, 264, 7, 93, 2, 2, 264, 62, 3, 2, 2, 2, 265, 266, 7,
	95, 2, 2, 266, 64, 3, 2, 2, 2, 267, 268, 7, 63, 2, 2, 268, 269, 7, 63,
	2, 2, 269, 66, 3, 2, 2, 2, 270, 271, 7, 35, 2, 2, 271, 272, 7, 63, 2, 2,
	272, 68, 3, 2, 2, 2, 273, 274, 7, 64, 2, 2, 274, 70, 3, 2, 2, 2, 275, 276,
	7, 62, 2, 2, 276, 72, 3, 2, 2, 2, 277, 278, 7, 64, 2, 2, 278, 279, 7, 63,
	2, 2, 279, 74, 3, 2, 2, 2, 280, 281, 7, 62, 2, 2, 281, 282, 7, 63, 2, 2,
	282, 76, 3, 2, 2, 2, 283, 284, 7, 44, 2, 2, 284, 78, 3, 2, 2, 2, 285, 286,
	7, 49, 2, 2, 286, 80, 3, 2, 2, 2, 287, 288, 7, 45, 2, 2, 288, 82, 3, 2,
	2, 2, 289, 290, 7, 47, 2, 2, 290, 84, 3, 2, 2, 2, 291, 292, 7, 39, 2, 2,
	292, 86, 3, 2, 2, 2, 293, 295, 9, 8, 2, 2, 294, 293, 3, 2, 2, 2, 295, 296,
	3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 3, 2,
	2, 2, 298, 299, 8, 44, 2, 2, 299, 88, 3, 2, 2, 2, 300, 301, 7, 94, 2, 2,
	301, 302, 9, 9, 2, 2, 302, 90, 3, 2, 2, 2, 12, 2, 181, 194, 201, 206, 212,
	218, 232, 238, 296, 3, 8, 2, 2,
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
	"", "'Sentencia'", "'consola'", "'consolaln'", "'declarar'", "'String'",
	"'Integer'", "'Real'", "'Boolean'", "'If'", "'entonces'", "'AND'", "'OR'",
	"", "", "", "", "", "", "", "'int'", "'='", "'.'", "';'", "','", "':'",
	"'('", "')'", "'{'", "'}'", "'['", "']'", "'=='", "'!='", "'>'", "'<'",
	"'>='", "'<='", "'*'", "'/'", "'+'", "'-'", "'%'",
}

var lexerSymbolicNames = []string{
	"", "SENTENCIA", "CONSOLA", "CONSOLALN", "DECLARAR", "RSTRING", "RINTEGER",
	"RREAL", "RBOOLEAN", "RIF", "RELSE", "RAND", "ROR", "MULTILINE", "INLINE",
	"INTEGER", "FLOAT", "STRING", "BOOLEAN", "ID", "DATA_TYPE", "EQUAL", "DOT",
	"SEMICOLON", "COMMA", "COLON", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY",
	"LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN",
	"LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "ADD",
	"SUB", "MOD", "WHITESPACE",
}

var lexerRuleNames = []string{
	"SENTENCIA", "CONSOLA", "CONSOLALN", "DECLARAR", "RSTRING", "RINTEGER",
	"RREAL", "RBOOLEAN", "RIF", "RELSE", "RAND", "ROR", "MULTILINE", "INLINE",
	"INTEGER", "FLOAT", "STRING", "BOOLEAN", "ID", "DATA_TYPE", "EQUAL", "DOT",
	"SEMICOLON", "COMMA", "COLON", "LEFT_PAR", "RIGHT_PAR", "LEFT_KEY", "RIGHT_KEY",
	"LEFT_BRACKET", "RIGHT_BRACKET", "EQUEAL_EQUAL", "NOTEQUAL", "GREATER_THAN",
	"LESS_THAN", "GREATER_EQUALTHAN", "LESS_EQUEALTHAN", "MUL", "DIV", "ADD",
	"SUB", "MOD", "WHITESPACE", "ESC_SEQ",
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
	ProjectLexerSENTENCIA         = 1
	ProjectLexerCONSOLA           = 2
	ProjectLexerCONSOLALN         = 3
	ProjectLexerDECLARAR          = 4
	ProjectLexerRSTRING           = 5
	ProjectLexerRINTEGER          = 6
	ProjectLexerRREAL             = 7
	ProjectLexerRBOOLEAN          = 8
	ProjectLexerRIF               = 9
	ProjectLexerRELSE             = 10
	ProjectLexerRAND              = 11
	ProjectLexerROR               = 12
	ProjectLexerMULTILINE         = 13
	ProjectLexerINLINE            = 14
	ProjectLexerINTEGER           = 15
	ProjectLexerFLOAT             = 16
	ProjectLexerSTRING            = 17
	ProjectLexerBOOLEAN           = 18
	ProjectLexerID                = 19
	ProjectLexerDATA_TYPE         = 20
	ProjectLexerEQUAL             = 21
	ProjectLexerDOT               = 22
	ProjectLexerSEMICOLON         = 23
	ProjectLexerCOMMA             = 24
	ProjectLexerCOLON             = 25
	ProjectLexerLEFT_PAR          = 26
	ProjectLexerRIGHT_PAR         = 27
	ProjectLexerLEFT_KEY          = 28
	ProjectLexerRIGHT_KEY         = 29
	ProjectLexerLEFT_BRACKET      = 30
	ProjectLexerRIGHT_BRACKET     = 31
	ProjectLexerEQUEAL_EQUAL      = 32
	ProjectLexerNOTEQUAL          = 33
	ProjectLexerGREATER_THAN      = 34
	ProjectLexerLESS_THAN         = 35
	ProjectLexerGREATER_EQUALTHAN = 36
	ProjectLexerLESS_EQUEALTHAN   = 37
	ProjectLexerMUL               = 38
	ProjectLexerDIV               = 39
	ProjectLexerADD               = 40
	ProjectLexerSUB               = 41
	ProjectLexerMOD               = 42
	ProjectLexerWHITESPACE        = 43
)
