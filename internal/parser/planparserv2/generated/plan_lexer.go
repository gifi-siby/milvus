// Code generated from Plan.g4 by ANTLR 4.13.2. DO NOT EDIT.

package planparserv2

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PlanLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PlanLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func planlexerLexerInit() {
	staticData := &PlanLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'['", "','", "']'", "'{'", "'}'", "'<'", "'<='",
		"'>'", "'>='", "'=='", "'!='", "", "", "", "", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'**'", "'<<'", "'>>'", "'&'", "'|'", "'^'", "", "", "",
		"", "'~'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'$meta'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "LBRACE", "RBRACE", "LT", "LE", "GT", "GE",
		"EQ", "NE", "LIKE", "EXISTS", "TEXTMATCH", "PHRASEMATCH", "ADD", "SUB",
		"MUL", "DIV", "MOD", "POW", "SHL", "SHR", "BAND", "BOR", "BXOR", "AND",
		"OR", "ISNULL", "ISNOTNULL", "BNOT", "NOT", "IN", "EmptyArray", "JSONContains",
		"JSONContainsAll", "JSONContainsAny", "ArrayContains", "ArrayContainsAll",
		"ArrayContainsAny", "ArrayLength", "BooleanConstant", "IntegerConstant",
		"FloatingConstant", "Identifier", "Meta", "StringLiteral", "JSONIdentifier",
		"Whitespace", "Newline",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "LBRACE", "RBRACE", "LT", "LE",
		"GT", "GE", "EQ", "NE", "LIKE", "EXISTS", "TEXTMATCH", "PHRASEMATCH",
		"ADD", "SUB", "MUL", "DIV", "MOD", "POW", "SHL", "SHR", "BAND", "BOR",
		"BXOR", "AND", "OR", "ISNULL", "ISNOTNULL", "BNOT", "NOT", "IN", "EmptyArray",
		"JSONContains", "JSONContainsAll", "JSONContainsAny", "ArrayContains",
		"ArrayContainsAll", "ArrayContainsAny", "ArrayLength", "BooleanConstant",
		"IntegerConstant", "FloatingConstant", "Identifier", "Meta", "StringLiteral",
		"JSONIdentifier", "EncodingPrefix", "DoubleSCharSequence", "SingleSCharSequence",
		"DoubleSChar", "SingleSChar", "Nondigit", "Digit", "BinaryConstant",
		"DecimalConstant", "OctalConstant", "HexadecimalConstant", "NonzeroDigit",
		"OctalDigit", "HexadecimalDigit", "HexQuad", "UniversalCharacterName",
		"DecimalFloatingConstant", "HexadecimalFloatingConstant", "FractionalConstant",
		"ExponentPart", "DigitSequence", "HexadecimalFractionalConstant", "HexadecimalDigitSequence",
		"BinaryExponentPart", "EscapeSequence", "Whitespace", "Newline",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 52, 862, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 194, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 208, 8, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 230, 8,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 3, 16, 256, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 291, 8,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 299, 8, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 3, 30, 315, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 339, 8, 31, 1, 32,
		1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 350, 8,
		33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 356, 8, 34, 1, 35, 1, 35, 1, 35,
		5, 35, 361, 8, 35, 10, 35, 12, 35, 364, 9, 35, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 3, 36, 394, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 430,
		8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 3, 38, 466, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 3, 39, 496, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		3, 40, 534, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 572, 8, 41,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 3, 42, 598, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1,
		43, 1, 43, 3, 43, 627, 8, 43, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 633, 8,
		44, 1, 45, 1, 45, 3, 45, 637, 8, 45, 1, 46, 1, 46, 1, 46, 5, 46, 642, 8,
		46, 10, 46, 12, 46, 645, 9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 48, 3, 48, 654, 8, 48, 1, 48, 1, 48, 3, 48, 658, 8, 48, 1, 48, 1, 48,
		1, 48, 3, 48, 663, 8, 48, 1, 48, 3, 48, 666, 8, 48, 1, 49, 1, 49, 3, 49,
		670, 8, 49, 1, 49, 1, 49, 1, 49, 3, 49, 675, 8, 49, 1, 49, 1, 49, 4, 49,
		679, 8, 49, 11, 49, 12, 49, 680, 1, 50, 1, 50, 1, 50, 3, 50, 686, 8, 50,
		1, 51, 4, 51, 689, 8, 51, 11, 51, 12, 51, 690, 1, 52, 4, 52, 694, 8, 52,
		11, 52, 12, 52, 695, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 3,
		53, 705, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54,
		714, 8, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 4, 57, 723,
		8, 57, 11, 57, 12, 57, 724, 1, 58, 1, 58, 5, 58, 729, 8, 58, 10, 58, 12,
		58, 732, 9, 58, 1, 58, 3, 58, 735, 8, 58, 1, 59, 1, 59, 5, 59, 739, 8,
		59, 10, 59, 12, 59, 742, 9, 59, 1, 60, 1, 60, 1, 60, 1, 60, 1, 61, 1, 61,
		1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1,
		65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 3, 65, 769,
		8, 65, 1, 66, 1, 66, 3, 66, 773, 8, 66, 1, 66, 1, 66, 1, 66, 3, 66, 778,
		8, 66, 1, 67, 1, 67, 1, 67, 1, 67, 3, 67, 784, 8, 67, 1, 67, 1, 67, 1,
		68, 3, 68, 789, 8, 68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 3, 68, 796, 8,
		68, 1, 69, 1, 69, 3, 69, 800, 8, 69, 1, 69, 1, 69, 1, 70, 4, 70, 805, 8,
		70, 11, 70, 12, 70, 806, 1, 71, 3, 71, 810, 8, 71, 1, 71, 1, 71, 1, 71,
		1, 71, 1, 71, 3, 71, 817, 8, 71, 1, 72, 4, 72, 820, 8, 72, 11, 72, 12,
		72, 821, 1, 73, 1, 73, 3, 73, 826, 8, 73, 1, 73, 1, 73, 1, 74, 1, 74, 1,
		74, 1, 74, 1, 74, 3, 74, 835, 8, 74, 1, 74, 3, 74, 838, 8, 74, 1, 74, 1,
		74, 1, 74, 1, 74, 1, 74, 3, 74, 845, 8, 74, 1, 75, 4, 75, 848, 8, 75, 11,
		75, 12, 75, 849, 1, 75, 1, 75, 1, 76, 1, 76, 3, 76, 856, 8, 76, 1, 76,
		3, 76, 859, 8, 76, 1, 76, 1, 76, 0, 0, 77, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29,
		15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47,
		24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65,
		33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83,
		42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101,
		0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113, 0, 115, 0, 117, 0, 119,
		0, 121, 0, 123, 0, 125, 0, 127, 0, 129, 0, 131, 0, 133, 0, 135, 0, 137,
		0, 139, 0, 141, 0, 143, 0, 145, 0, 147, 0, 149, 0, 151, 51, 153, 52, 1,
		0, 16, 3, 0, 76, 76, 85, 85, 117, 117, 4, 0, 10, 10, 13, 13, 34, 34, 92,
		92, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 3, 0, 65, 90, 95, 95, 97, 122,
		1, 0, 48, 57, 2, 0, 66, 66, 98, 98, 1, 0, 48, 49, 2, 0, 88, 88, 120, 120,
		1, 0, 49, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 69, 69,
		101, 101, 2, 0, 43, 43, 45, 45, 2, 0, 80, 80, 112, 112, 10, 0, 34, 34,
		39, 39, 63, 63, 92, 92, 97, 98, 102, 102, 110, 110, 114, 114, 116, 116,
		118, 118, 2, 0, 9, 9, 32, 32, 909, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0,
		0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0,
		0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0,
		0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0,
		0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1,
		0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43,
		1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0,
		51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0,
		0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0,
		0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0,
		0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1,
		0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89,
		1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0,
		97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 151, 1, 0, 0, 0, 0, 153, 1, 0, 0,
		0, 1, 155, 1, 0, 0, 0, 3, 157, 1, 0, 0, 0, 5, 159, 1, 0, 0, 0, 7, 161,
		1, 0, 0, 0, 9, 163, 1, 0, 0, 0, 11, 165, 1, 0, 0, 0, 13, 167, 1, 0, 0,
		0, 15, 169, 1, 0, 0, 0, 17, 171, 1, 0, 0, 0, 19, 174, 1, 0, 0, 0, 21, 176,
		1, 0, 0, 0, 23, 179, 1, 0, 0, 0, 25, 182, 1, 0, 0, 0, 27, 193, 1, 0, 0,
		0, 29, 207, 1, 0, 0, 0, 31, 229, 1, 0, 0, 0, 33, 255, 1, 0, 0, 0, 35, 257,
		1, 0, 0, 0, 37, 259, 1, 0, 0, 0, 39, 261, 1, 0, 0, 0, 41, 263, 1, 0, 0,
		0, 43, 265, 1, 0, 0, 0, 45, 267, 1, 0, 0, 0, 47, 270, 1, 0, 0, 0, 49, 273,
		1, 0, 0, 0, 51, 276, 1, 0, 0, 0, 53, 278, 1, 0, 0, 0, 55, 280, 1, 0, 0,
		0, 57, 290, 1, 0, 0, 0, 59, 298, 1, 0, 0, 0, 61, 314, 1, 0, 0, 0, 63, 338,
		1, 0, 0, 0, 65, 340, 1, 0, 0, 0, 67, 349, 1, 0, 0, 0, 69, 355, 1, 0, 0,
		0, 71, 357, 1, 0, 0, 0, 73, 393, 1, 0, 0, 0, 75, 429, 1, 0, 0, 0, 77, 465,
		1, 0, 0, 0, 79, 495, 1, 0, 0, 0, 81, 533, 1, 0, 0, 0, 83, 571, 1, 0, 0,
		0, 85, 597, 1, 0, 0, 0, 87, 626, 1, 0, 0, 0, 89, 632, 1, 0, 0, 0, 91, 636,
		1, 0, 0, 0, 93, 638, 1, 0, 0, 0, 95, 646, 1, 0, 0, 0, 97, 653, 1, 0, 0,
		0, 99, 669, 1, 0, 0, 0, 101, 685, 1, 0, 0, 0, 103, 688, 1, 0, 0, 0, 105,
		693, 1, 0, 0, 0, 107, 704, 1, 0, 0, 0, 109, 713, 1, 0, 0, 0, 111, 715,
		1, 0, 0, 0, 113, 717, 1, 0, 0, 0, 115, 719, 1, 0, 0, 0, 117, 734, 1, 0,
		0, 0, 119, 736, 1, 0, 0, 0, 121, 743, 1, 0, 0, 0, 123, 747, 1, 0, 0, 0,
		125, 749, 1, 0, 0, 0, 127, 751, 1, 0, 0, 0, 129, 753, 1, 0, 0, 0, 131,
		768, 1, 0, 0, 0, 133, 777, 1, 0, 0, 0, 135, 779, 1, 0, 0, 0, 137, 795,
		1, 0, 0, 0, 139, 797, 1, 0, 0, 0, 141, 804, 1, 0, 0, 0, 143, 816, 1, 0,
		0, 0, 145, 819, 1, 0, 0, 0, 147, 823, 1, 0, 0, 0, 149, 844, 1, 0, 0, 0,
		151, 847, 1, 0, 0, 0, 153, 858, 1, 0, 0, 0, 155, 156, 5, 40, 0, 0, 156,
		2, 1, 0, 0, 0, 157, 158, 5, 41, 0, 0, 158, 4, 1, 0, 0, 0, 159, 160, 5,
		91, 0, 0, 160, 6, 1, 0, 0, 0, 161, 162, 5, 44, 0, 0, 162, 8, 1, 0, 0, 0,
		163, 164, 5, 93, 0, 0, 164, 10, 1, 0, 0, 0, 165, 166, 5, 123, 0, 0, 166,
		12, 1, 0, 0, 0, 167, 168, 5, 125, 0, 0, 168, 14, 1, 0, 0, 0, 169, 170,
		5, 60, 0, 0, 170, 16, 1, 0, 0, 0, 171, 172, 5, 60, 0, 0, 172, 173, 5, 61,
		0, 0, 173, 18, 1, 0, 0, 0, 174, 175, 5, 62, 0, 0, 175, 20, 1, 0, 0, 0,
		176, 177, 5, 62, 0, 0, 177, 178, 5, 61, 0, 0, 178, 22, 1, 0, 0, 0, 179,
		180, 5, 61, 0, 0, 180, 181, 5, 61, 0, 0, 181, 24, 1, 0, 0, 0, 182, 183,
		5, 33, 0, 0, 183, 184, 5, 61, 0, 0, 184, 26, 1, 0, 0, 0, 185, 186, 5, 108,
		0, 0, 186, 187, 5, 105, 0, 0, 187, 188, 5, 107, 0, 0, 188, 194, 5, 101,
		0, 0, 189, 190, 5, 76, 0, 0, 190, 191, 5, 73, 0, 0, 191, 192, 5, 75, 0,
		0, 192, 194, 5, 69, 0, 0, 193, 185, 1, 0, 0, 0, 193, 189, 1, 0, 0, 0, 194,
		28, 1, 0, 0, 0, 195, 196, 5, 101, 0, 0, 196, 197, 5, 120, 0, 0, 197, 198,
		5, 105, 0, 0, 198, 199, 5, 115, 0, 0, 199, 200, 5, 116, 0, 0, 200, 208,
		5, 115, 0, 0, 201, 202, 5, 69, 0, 0, 202, 203, 5, 88, 0, 0, 203, 204, 5,
		73, 0, 0, 204, 205, 5, 83, 0, 0, 205, 206, 5, 84, 0, 0, 206, 208, 5, 83,
		0, 0, 207, 195, 1, 0, 0, 0, 207, 201, 1, 0, 0, 0, 208, 30, 1, 0, 0, 0,
		209, 210, 5, 116, 0, 0, 210, 211, 5, 101, 0, 0, 211, 212, 5, 120, 0, 0,
		212, 213, 5, 116, 0, 0, 213, 214, 5, 95, 0, 0, 214, 215, 5, 109, 0, 0,
		215, 216, 5, 97, 0, 0, 216, 217, 5, 116, 0, 0, 217, 218, 5, 99, 0, 0, 218,
		230, 5, 104, 0, 0, 219, 220, 5, 84, 0, 0, 220, 221, 5, 69, 0, 0, 221, 222,
		5, 88, 0, 0, 222, 223, 5, 84, 0, 0, 223, 224, 5, 95, 0, 0, 224, 225, 5,
		77, 0, 0, 225, 226, 5, 65, 0, 0, 226, 227, 5, 84, 0, 0, 227, 228, 5, 67,
		0, 0, 228, 230, 5, 72, 0, 0, 229, 209, 1, 0, 0, 0, 229, 219, 1, 0, 0, 0,
		230, 32, 1, 0, 0, 0, 231, 232, 5, 112, 0, 0, 232, 233, 5, 104, 0, 0, 233,
		234, 5, 114, 0, 0, 234, 235, 5, 97, 0, 0, 235, 236, 5, 115, 0, 0, 236,
		237, 5, 101, 0, 0, 237, 238, 5, 95, 0, 0, 238, 239, 5, 109, 0, 0, 239,
		240, 5, 97, 0, 0, 240, 241, 5, 116, 0, 0, 241, 242, 5, 99, 0, 0, 242, 256,
		5, 104, 0, 0, 243, 244, 5, 80, 0, 0, 244, 245, 5, 72, 0, 0, 245, 246, 5,
		82, 0, 0, 246, 247, 5, 65, 0, 0, 247, 248, 5, 83, 0, 0, 248, 249, 5, 69,
		0, 0, 249, 250, 5, 95, 0, 0, 250, 251, 5, 77, 0, 0, 251, 252, 5, 65, 0,
		0, 252, 253, 5, 84, 0, 0, 253, 254, 5, 67, 0, 0, 254, 256, 5, 72, 0, 0,
		255, 231, 1, 0, 0, 0, 255, 243, 1, 0, 0, 0, 256, 34, 1, 0, 0, 0, 257, 258,
		5, 43, 0, 0, 258, 36, 1, 0, 0, 0, 259, 260, 5, 45, 0, 0, 260, 38, 1, 0,
		0, 0, 261, 262, 5, 42, 0, 0, 262, 40, 1, 0, 0, 0, 263, 264, 5, 47, 0, 0,
		264, 42, 1, 0, 0, 0, 265, 266, 5, 37, 0, 0, 266, 44, 1, 0, 0, 0, 267, 268,
		5, 42, 0, 0, 268, 269, 5, 42, 0, 0, 269, 46, 1, 0, 0, 0, 270, 271, 5, 60,
		0, 0, 271, 272, 5, 60, 0, 0, 272, 48, 1, 0, 0, 0, 273, 274, 5, 62, 0, 0,
		274, 275, 5, 62, 0, 0, 275, 50, 1, 0, 0, 0, 276, 277, 5, 38, 0, 0, 277,
		52, 1, 0, 0, 0, 278, 279, 5, 124, 0, 0, 279, 54, 1, 0, 0, 0, 280, 281,
		5, 94, 0, 0, 281, 56, 1, 0, 0, 0, 282, 283, 5, 38, 0, 0, 283, 291, 5, 38,
		0, 0, 284, 285, 5, 97, 0, 0, 285, 286, 5, 110, 0, 0, 286, 291, 5, 100,
		0, 0, 287, 288, 5, 65, 0, 0, 288, 289, 5, 78, 0, 0, 289, 291, 5, 68, 0,
		0, 290, 282, 1, 0, 0, 0, 290, 284, 1, 0, 0, 0, 290, 287, 1, 0, 0, 0, 291,
		58, 1, 0, 0, 0, 292, 293, 5, 124, 0, 0, 293, 299, 5, 124, 0, 0, 294, 295,
		5, 111, 0, 0, 295, 299, 5, 114, 0, 0, 296, 297, 5, 79, 0, 0, 297, 299,
		5, 82, 0, 0, 298, 292, 1, 0, 0, 0, 298, 294, 1, 0, 0, 0, 298, 296, 1, 0,
		0, 0, 299, 60, 1, 0, 0, 0, 300, 301, 5, 105, 0, 0, 301, 302, 5, 115, 0,
		0, 302, 303, 5, 32, 0, 0, 303, 304, 5, 110, 0, 0, 304, 305, 5, 117, 0,
		0, 305, 306, 5, 108, 0, 0, 306, 315, 5, 108, 0, 0, 307, 308, 5, 73, 0,
		0, 308, 309, 5, 83, 0, 0, 309, 310, 5, 32, 0, 0, 310, 311, 5, 78, 0, 0,
		311, 312, 5, 85, 0, 0, 312, 313, 5, 76, 0, 0, 313, 315, 5, 76, 0, 0, 314,
		300, 1, 0, 0, 0, 314, 307, 1, 0, 0, 0, 315, 62, 1, 0, 0, 0, 316, 317, 5,
		105, 0, 0, 317, 318, 5, 115, 0, 0, 318, 319, 5, 32, 0, 0, 319, 320, 5,
		110, 0, 0, 320, 321, 5, 111, 0, 0, 321, 322, 5, 116, 0, 0, 322, 323, 5,
		32, 0, 0, 323, 324, 5, 110, 0, 0, 324, 325, 5, 117, 0, 0, 325, 326, 5,
		108, 0, 0, 326, 339, 5, 108, 0, 0, 327, 328, 5, 73, 0, 0, 328, 329, 5,
		83, 0, 0, 329, 330, 5, 32, 0, 0, 330, 331, 5, 78, 0, 0, 331, 332, 5, 79,
		0, 0, 332, 333, 5, 84, 0, 0, 333, 334, 5, 32, 0, 0, 334, 335, 5, 78, 0,
		0, 335, 336, 5, 85, 0, 0, 336, 337, 5, 76, 0, 0, 337, 339, 5, 76, 0, 0,
		338, 316, 1, 0, 0, 0, 338, 327, 1, 0, 0, 0, 339, 64, 1, 0, 0, 0, 340, 341,
		5, 126, 0, 0, 341, 66, 1, 0, 0, 0, 342, 350, 5, 33, 0, 0, 343, 344, 5,
		110, 0, 0, 344, 345, 5, 111, 0, 0, 345, 350, 5, 116, 0, 0, 346, 347, 5,
		78, 0, 0, 347, 348, 5, 79, 0, 0, 348, 350, 5, 84, 0, 0, 349, 342, 1, 0,
		0, 0, 349, 343, 1, 0, 0, 0, 349, 346, 1, 0, 0, 0, 350, 68, 1, 0, 0, 0,
		351, 352, 5, 105, 0, 0, 352, 356, 5, 110, 0, 0, 353, 354, 5, 73, 0, 0,
		354, 356, 5, 78, 0, 0, 355, 351, 1, 0, 0, 0, 355, 353, 1, 0, 0, 0, 356,
		70, 1, 0, 0, 0, 357, 362, 5, 91, 0, 0, 358, 361, 3, 151, 75, 0, 359, 361,
		3, 153, 76, 0, 360, 358, 1, 0, 0, 0, 360, 359, 1, 0, 0, 0, 361, 364, 1,
		0, 0, 0, 362, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 365, 1, 0, 0,
		0, 364, 362, 1, 0, 0, 0, 365, 366, 5, 93, 0, 0, 366, 72, 1, 0, 0, 0, 367,
		368, 5, 106, 0, 0, 368, 369, 5, 115, 0, 0, 369, 370, 5, 111, 0, 0, 370,
		371, 5, 110, 0, 0, 371, 372, 5, 95, 0, 0, 372, 373, 5, 99, 0, 0, 373, 374,
		5, 111, 0, 0, 374, 375, 5, 110, 0, 0, 375, 376, 5, 116, 0, 0, 376, 377,
		5, 97, 0, 0, 377, 378, 5, 105, 0, 0, 378, 379, 5, 110, 0, 0, 379, 394,
		5, 115, 0, 0, 380, 381, 5, 74, 0, 0, 381, 382, 5, 83, 0, 0, 382, 383, 5,
		79, 0, 0, 383, 384, 5, 78, 0, 0, 384, 385, 5, 95, 0, 0, 385, 386, 5, 67,
		0, 0, 386, 387, 5, 79, 0, 0, 387, 388, 5, 78, 0, 0, 388, 389, 5, 84, 0,
		0, 389, 390, 5, 65, 0, 0, 390, 391, 5, 73, 0, 0, 391, 392, 5, 78, 0, 0,
		392, 394, 5, 83, 0, 0, 393, 367, 1, 0, 0, 0, 393, 380, 1, 0, 0, 0, 394,
		74, 1, 0, 0, 0, 395, 396, 5, 106, 0, 0, 396, 397, 5, 115, 0, 0, 397, 398,
		5, 111, 0, 0, 398, 399, 5, 110, 0, 0, 399, 400, 5, 95, 0, 0, 400, 401,
		5, 99, 0, 0, 401, 402, 5, 111, 0, 0, 402, 403, 5, 110, 0, 0, 403, 404,
		5, 116, 0, 0, 404, 405, 5, 97, 0, 0, 405, 406, 5, 105, 0, 0, 406, 407,
		5, 110, 0, 0, 407, 408, 5, 115, 0, 0, 408, 409, 5, 95, 0, 0, 409, 410,
		5, 97, 0, 0, 410, 411, 5, 108, 0, 0, 411, 430, 5, 108, 0, 0, 412, 413,
		5, 74, 0, 0, 413, 414, 5, 83, 0, 0, 414, 415, 5, 79, 0, 0, 415, 416, 5,
		78, 0, 0, 416, 417, 5, 95, 0, 0, 417, 418, 5, 67, 0, 0, 418, 419, 5, 79,
		0, 0, 419, 420, 5, 78, 0, 0, 420, 421, 5, 84, 0, 0, 421, 422, 5, 65, 0,
		0, 422, 423, 5, 73, 0, 0, 423, 424, 5, 78, 0, 0, 424, 425, 5, 83, 0, 0,
		425, 426, 5, 95, 0, 0, 426, 427, 5, 65, 0, 0, 427, 428, 5, 76, 0, 0, 428,
		430, 5, 76, 0, 0, 429, 395, 1, 0, 0, 0, 429, 412, 1, 0, 0, 0, 430, 76,
		1, 0, 0, 0, 431, 432, 5, 106, 0, 0, 432, 433, 5, 115, 0, 0, 433, 434, 5,
		111, 0, 0, 434, 435, 5, 110, 0, 0, 435, 436, 5, 95, 0, 0, 436, 437, 5,
		99, 0, 0, 437, 438, 5, 111, 0, 0, 438, 439, 5, 110, 0, 0, 439, 440, 5,
		116, 0, 0, 440, 441, 5, 97, 0, 0, 441, 442, 5, 105, 0, 0, 442, 443, 5,
		110, 0, 0, 443, 444, 5, 115, 0, 0, 444, 445, 5, 95, 0, 0, 445, 446, 5,
		97, 0, 0, 446, 447, 5, 110, 0, 0, 447, 466, 5, 121, 0, 0, 448, 449, 5,
		74, 0, 0, 449, 450, 5, 83, 0, 0, 450, 451, 5, 79, 0, 0, 451, 452, 5, 78,
		0, 0, 452, 453, 5, 95, 0, 0, 453, 454, 5, 67, 0, 0, 454, 455, 5, 79, 0,
		0, 455, 456, 5, 78, 0, 0, 456, 457, 5, 84, 0, 0, 457, 458, 5, 65, 0, 0,
		458, 459, 5, 73, 0, 0, 459, 460, 5, 78, 0, 0, 460, 461, 5, 83, 0, 0, 461,
		462, 5, 95, 0, 0, 462, 463, 5, 65, 0, 0, 463, 464, 5, 78, 0, 0, 464, 466,
		5, 89, 0, 0, 465, 431, 1, 0, 0, 0, 465, 448, 1, 0, 0, 0, 466, 78, 1, 0,
		0, 0, 467, 468, 5, 97, 0, 0, 468, 469, 5, 114, 0, 0, 469, 470, 5, 114,
		0, 0, 470, 471, 5, 97, 0, 0, 471, 472, 5, 121, 0, 0, 472, 473, 5, 95, 0,
		0, 473, 474, 5, 99, 0, 0, 474, 475, 5, 111, 0, 0, 475, 476, 5, 110, 0,
		0, 476, 477, 5, 116, 0, 0, 477, 478, 5, 97, 0, 0, 478, 479, 5, 105, 0,
		0, 479, 480, 5, 110, 0, 0, 480, 496, 5, 115, 0, 0, 481, 482, 5, 65, 0,
		0, 482, 483, 5, 82, 0, 0, 483, 484, 5, 82, 0, 0, 484, 485, 5, 65, 0, 0,
		485, 486, 5, 89, 0, 0, 486, 487, 5, 95, 0, 0, 487, 488, 5, 67, 0, 0, 488,
		489, 5, 79, 0, 0, 489, 490, 5, 78, 0, 0, 490, 491, 5, 84, 0, 0, 491, 492,
		5, 65, 0, 0, 492, 493, 5, 73, 0, 0, 493, 494, 5, 78, 0, 0, 494, 496, 5,
		83, 0, 0, 495, 467, 1, 0, 0, 0, 495, 481, 1, 0, 0, 0, 496, 80, 1, 0, 0,
		0, 497, 498, 5, 97, 0, 0, 498, 499, 5, 114, 0, 0, 499, 500, 5, 114, 0,
		0, 500, 501, 5, 97, 0, 0, 501, 502, 5, 121, 0, 0, 502, 503, 5, 95, 0, 0,
		503, 504, 5, 99, 0, 0, 504, 505, 5, 111, 0, 0, 505, 506, 5, 110, 0, 0,
		506, 507, 5, 116, 0, 0, 507, 508, 5, 97, 0, 0, 508, 509, 5, 105, 0, 0,
		509, 510, 5, 110, 0, 0, 510, 511, 5, 115, 0, 0, 511, 512, 5, 95, 0, 0,
		512, 513, 5, 97, 0, 0, 513, 514, 5, 108, 0, 0, 514, 534, 5, 108, 0, 0,
		515, 516, 5, 65, 0, 0, 516, 517, 5, 82, 0, 0, 517, 518, 5, 82, 0, 0, 518,
		519, 5, 65, 0, 0, 519, 520, 5, 89, 0, 0, 520, 521, 5, 95, 0, 0, 521, 522,
		5, 67, 0, 0, 522, 523, 5, 79, 0, 0, 523, 524, 5, 78, 0, 0, 524, 525, 5,
		84, 0, 0, 525, 526, 5, 65, 0, 0, 526, 527, 5, 73, 0, 0, 527, 528, 5, 78,
		0, 0, 528, 529, 5, 83, 0, 0, 529, 530, 5, 95, 0, 0, 530, 531, 5, 65, 0,
		0, 531, 532, 5, 76, 0, 0, 532, 534, 5, 76, 0, 0, 533, 497, 1, 0, 0, 0,
		533, 515, 1, 0, 0, 0, 534, 82, 1, 0, 0, 0, 535, 536, 5, 97, 0, 0, 536,
		537, 5, 114, 0, 0, 537, 538, 5, 114, 0, 0, 538, 539, 5, 97, 0, 0, 539,
		540, 5, 121, 0, 0, 540, 541, 5, 95, 0, 0, 541, 542, 5, 99, 0, 0, 542, 543,
		5, 111, 0, 0, 543, 544, 5, 110, 0, 0, 544, 545, 5, 116, 0, 0, 545, 546,
		5, 97, 0, 0, 546, 547, 5, 105, 0, 0, 547, 548, 5, 110, 0, 0, 548, 549,
		5, 115, 0, 0, 549, 550, 5, 95, 0, 0, 550, 551, 5, 97, 0, 0, 551, 552, 5,
		110, 0, 0, 552, 572, 5, 121, 0, 0, 553, 554, 5, 65, 0, 0, 554, 555, 5,
		82, 0, 0, 555, 556, 5, 82, 0, 0, 556, 557, 5, 65, 0, 0, 557, 558, 5, 89,
		0, 0, 558, 559, 5, 95, 0, 0, 559, 560, 5, 67, 0, 0, 560, 561, 5, 79, 0,
		0, 561, 562, 5, 78, 0, 0, 562, 563, 5, 84, 0, 0, 563, 564, 5, 65, 0, 0,
		564, 565, 5, 73, 0, 0, 565, 566, 5, 78, 0, 0, 566, 567, 5, 83, 0, 0, 567,
		568, 5, 95, 0, 0, 568, 569, 5, 65, 0, 0, 569, 570, 5, 78, 0, 0, 570, 572,
		5, 89, 0, 0, 571, 535, 1, 0, 0, 0, 571, 553, 1, 0, 0, 0, 572, 84, 1, 0,
		0, 0, 573, 574, 5, 97, 0, 0, 574, 575, 5, 114, 0, 0, 575, 576, 5, 114,
		0, 0, 576, 577, 5, 97, 0, 0, 577, 578, 5, 121, 0, 0, 578, 579, 5, 95, 0,
		0, 579, 580, 5, 108, 0, 0, 580, 581, 5, 101, 0, 0, 581, 582, 5, 110, 0,
		0, 582, 583, 5, 103, 0, 0, 583, 584, 5, 116, 0, 0, 584, 598, 5, 104, 0,
		0, 585, 586, 5, 65, 0, 0, 586, 587, 5, 82, 0, 0, 587, 588, 5, 82, 0, 0,
		588, 589, 5, 65, 0, 0, 589, 590, 5, 89, 0, 0, 590, 591, 5, 95, 0, 0, 591,
		592, 5, 76, 0, 0, 592, 593, 5, 69, 0, 0, 593, 594, 5, 78, 0, 0, 594, 595,
		5, 71, 0, 0, 595, 596, 5, 84, 0, 0, 596, 598, 5, 72, 0, 0, 597, 573, 1,
		0, 0, 0, 597, 585, 1, 0, 0, 0, 598, 86, 1, 0, 0, 0, 599, 600, 5, 116, 0,
		0, 600, 601, 5, 114, 0, 0, 601, 602, 5, 117, 0, 0, 602, 627, 5, 101, 0,
		0, 603, 604, 5, 84, 0, 0, 604, 605, 5, 114, 0, 0, 605, 606, 5, 117, 0,
		0, 606, 627, 5, 101, 0, 0, 607, 608, 5, 84, 0, 0, 608, 609, 5, 82, 0, 0,
		609, 610, 5, 85, 0, 0, 610, 627, 5, 69, 0, 0, 611, 612, 5, 102, 0, 0, 612,
		613, 5, 97, 0, 0, 613, 614, 5, 108, 0, 0, 614, 615, 5, 115, 0, 0, 615,
		627, 5, 101, 0, 0, 616, 617, 5, 70, 0, 0, 617, 618, 5, 97, 0, 0, 618, 619,
		5, 108, 0, 0, 619, 620, 5, 115, 0, 0, 620, 627, 5, 101, 0, 0, 621, 622,
		5, 70, 0, 0, 622, 623, 5, 65, 0, 0, 623, 624, 5, 76, 0, 0, 624, 625, 5,
		83, 0, 0, 625, 627, 5, 69, 0, 0, 626, 599, 1, 0, 0, 0, 626, 603, 1, 0,
		0, 0, 626, 607, 1, 0, 0, 0, 626, 611, 1, 0, 0, 0, 626, 616, 1, 0, 0, 0,
		626, 621, 1, 0, 0, 0, 627, 88, 1, 0, 0, 0, 628, 633, 3, 117, 58, 0, 629,
		633, 3, 119, 59, 0, 630, 633, 3, 121, 60, 0, 631, 633, 3, 115, 57, 0, 632,
		628, 1, 0, 0, 0, 632, 629, 1, 0, 0, 0, 632, 630, 1, 0, 0, 0, 632, 631,
		1, 0, 0, 0, 633, 90, 1, 0, 0, 0, 634, 637, 3, 133, 66, 0, 635, 637, 3,
		135, 67, 0, 636, 634, 1, 0, 0, 0, 636, 635, 1, 0, 0, 0, 637, 92, 1, 0,
		0, 0, 638, 643, 3, 111, 55, 0, 639, 642, 3, 111, 55, 0, 640, 642, 3, 113,
		56, 0, 641, 639, 1, 0, 0, 0, 641, 640, 1, 0, 0, 0, 642, 645, 1, 0, 0, 0,
		643, 641, 1, 0, 0, 0, 643, 644, 1, 0, 0, 0, 644, 94, 1, 0, 0, 0, 645, 643,
		1, 0, 0, 0, 646, 647, 5, 36, 0, 0, 647, 648, 5, 109, 0, 0, 648, 649, 5,
		101, 0, 0, 649, 650, 5, 116, 0, 0, 650, 651, 5, 97, 0, 0, 651, 96, 1, 0,
		0, 0, 652, 654, 3, 101, 50, 0, 653, 652, 1, 0, 0, 0, 653, 654, 1, 0, 0,
		0, 654, 665, 1, 0, 0, 0, 655, 657, 5, 34, 0, 0, 656, 658, 3, 103, 51, 0,
		657, 656, 1, 0, 0, 0, 657, 658, 1, 0, 0, 0, 658, 659, 1, 0, 0, 0, 659,
		666, 5, 34, 0, 0, 660, 662, 5, 39, 0, 0, 661, 663, 3, 105, 52, 0, 662,
		661, 1, 0, 0, 0, 662, 663, 1, 0, 0, 0, 663, 664, 1, 0, 0, 0, 664, 666,
		5, 39, 0, 0, 665, 655, 1, 0, 0, 0, 665, 660, 1, 0, 0, 0, 666, 98, 1, 0,
		0, 0, 667, 670, 3, 93, 46, 0, 668, 670, 3, 95, 47, 0, 669, 667, 1, 0, 0,
		0, 669, 668, 1, 0, 0, 0, 670, 678, 1, 0, 0, 0, 671, 674, 5, 91, 0, 0, 672,
		675, 3, 97, 48, 0, 673, 675, 3, 117, 58, 0, 674, 672, 1, 0, 0, 0, 674,
		673, 1, 0, 0, 0, 675, 676, 1, 0, 0, 0, 676, 677, 5, 93, 0, 0, 677, 679,
		1, 0, 0, 0, 678, 671, 1, 0, 0, 0, 679, 680, 1, 0, 0, 0, 680, 678, 1, 0,
		0, 0, 680, 681, 1, 0, 0, 0, 681, 100, 1, 0, 0, 0, 682, 683, 5, 117, 0,
		0, 683, 686, 5, 56, 0, 0, 684, 686, 7, 0, 0, 0, 685, 682, 1, 0, 0, 0, 685,
		684, 1, 0, 0, 0, 686, 102, 1, 0, 0, 0, 687, 689, 3, 107, 53, 0, 688, 687,
		1, 0, 0, 0, 689, 690, 1, 0, 0, 0, 690, 688, 1, 0, 0, 0, 690, 691, 1, 0,
		0, 0, 691, 104, 1, 0, 0, 0, 692, 694, 3, 109, 54, 0, 693, 692, 1, 0, 0,
		0, 694, 695, 1, 0, 0, 0, 695, 693, 1, 0, 0, 0, 695, 696, 1, 0, 0, 0, 696,
		106, 1, 0, 0, 0, 697, 705, 8, 1, 0, 0, 698, 705, 3, 149, 74, 0, 699, 700,
		5, 92, 0, 0, 700, 705, 5, 10, 0, 0, 701, 702, 5, 92, 0, 0, 702, 703, 5,
		13, 0, 0, 703, 705, 5, 10, 0, 0, 704, 697, 1, 0, 0, 0, 704, 698, 1, 0,
		0, 0, 704, 699, 1, 0, 0, 0, 704, 701, 1, 0, 0, 0, 705, 108, 1, 0, 0, 0,
		706, 714, 8, 2, 0, 0, 707, 714, 3, 149, 74, 0, 708, 709, 5, 92, 0, 0, 709,
		714, 5, 10, 0, 0, 710, 711, 5, 92, 0, 0, 711, 712, 5, 13, 0, 0, 712, 714,
		5, 10, 0, 0, 713, 706, 1, 0, 0, 0, 713, 707, 1, 0, 0, 0, 713, 708, 1, 0,
		0, 0, 713, 710, 1, 0, 0, 0, 714, 110, 1, 0, 0, 0, 715, 716, 7, 3, 0, 0,
		716, 112, 1, 0, 0, 0, 717, 718, 7, 4, 0, 0, 718, 114, 1, 0, 0, 0, 719,
		720, 5, 48, 0, 0, 720, 722, 7, 5, 0, 0, 721, 723, 7, 6, 0, 0, 722, 721,
		1, 0, 0, 0, 723, 724, 1, 0, 0, 0, 724, 722, 1, 0, 0, 0, 724, 725, 1, 0,
		0, 0, 725, 116, 1, 0, 0, 0, 726, 730, 3, 123, 61, 0, 727, 729, 3, 113,
		56, 0, 728, 727, 1, 0, 0, 0, 729, 732, 1, 0, 0, 0, 730, 728, 1, 0, 0, 0,
		730, 731, 1, 0, 0, 0, 731, 735, 1, 0, 0, 0, 732, 730, 1, 0, 0, 0, 733,
		735, 5, 48, 0, 0, 734, 726, 1, 0, 0, 0, 734, 733, 1, 0, 0, 0, 735, 118,
		1, 0, 0, 0, 736, 740, 5, 48, 0, 0, 737, 739, 3, 125, 62, 0, 738, 737, 1,
		0, 0, 0, 739, 742, 1, 0, 0, 0, 740, 738, 1, 0, 0, 0, 740, 741, 1, 0, 0,
		0, 741, 120, 1, 0, 0, 0, 742, 740, 1, 0, 0, 0, 743, 744, 5, 48, 0, 0, 744,
		745, 7, 7, 0, 0, 745, 746, 3, 145, 72, 0, 746, 122, 1, 0, 0, 0, 747, 748,
		7, 8, 0, 0, 748, 124, 1, 0, 0, 0, 749, 750, 7, 9, 0, 0, 750, 126, 1, 0,
		0, 0, 751, 752, 7, 10, 0, 0, 752, 128, 1, 0, 0, 0, 753, 754, 3, 127, 63,
		0, 754, 755, 3, 127, 63, 0, 755, 756, 3, 127, 63, 0, 756, 757, 3, 127,
		63, 0, 757, 130, 1, 0, 0, 0, 758, 759, 5, 92, 0, 0, 759, 760, 5, 117, 0,
		0, 760, 761, 1, 0, 0, 0, 761, 769, 3, 129, 64, 0, 762, 763, 5, 92, 0, 0,
		763, 764, 5, 85, 0, 0, 764, 765, 1, 0, 0, 0, 765, 766, 3, 129, 64, 0, 766,
		767, 3, 129, 64, 0, 767, 769, 1, 0, 0, 0, 768, 758, 1, 0, 0, 0, 768, 762,
		1, 0, 0, 0, 769, 132, 1, 0, 0, 0, 770, 772, 3, 137, 68, 0, 771, 773, 3,
		139, 69, 0, 772, 771, 1, 0, 0, 0, 772, 773, 1, 0, 0, 0, 773, 778, 1, 0,
		0, 0, 774, 775, 3, 141, 70, 0, 775, 776, 3, 139, 69, 0, 776, 778, 1, 0,
		0, 0, 777, 770, 1, 0, 0, 0, 777, 774, 1, 0, 0, 0, 778, 134, 1, 0, 0, 0,
		779, 780, 5, 48, 0, 0, 780, 783, 7, 7, 0, 0, 781, 784, 3, 143, 71, 0, 782,
		784, 3, 145, 72, 0, 783, 781, 1, 0, 0, 0, 783, 782, 1, 0, 0, 0, 784, 785,
		1, 0, 0, 0, 785, 786, 3, 147, 73, 0, 786, 136, 1, 0, 0, 0, 787, 789, 3,
		141, 70, 0, 788, 787, 1, 0, 0, 0, 788, 789, 1, 0, 0, 0, 789, 790, 1, 0,
		0, 0, 790, 791, 5, 46, 0, 0, 791, 796, 3, 141, 70, 0, 792, 793, 3, 141,
		70, 0, 793, 794, 5, 46, 0, 0, 794, 796, 1, 0, 0, 0, 795, 788, 1, 0, 0,
		0, 795, 792, 1, 0, 0, 0, 796, 138, 1, 0, 0, 0, 797, 799, 7, 11, 0, 0, 798,
		800, 7, 12, 0, 0, 799, 798, 1, 0, 0, 0, 799, 800, 1, 0, 0, 0, 800, 801,
		1, 0, 0, 0, 801, 802, 3, 141, 70, 0, 802, 140, 1, 0, 0, 0, 803, 805, 3,
		113, 56, 0, 804, 803, 1, 0, 0, 0, 805, 806, 1, 0, 0, 0, 806, 804, 1, 0,
		0, 0, 806, 807, 1, 0, 0, 0, 807, 142, 1, 0, 0, 0, 808, 810, 3, 145, 72,
		0, 809, 808, 1, 0, 0, 0, 809, 810, 1, 0, 0, 0, 810, 811, 1, 0, 0, 0, 811,
		812, 5, 46, 0, 0, 812, 817, 3, 145, 72, 0, 813, 814, 3, 145, 72, 0, 814,
		815, 5, 46, 0, 0, 815, 817, 1, 0, 0, 0, 816, 809, 1, 0, 0, 0, 816, 813,
		1, 0, 0, 0, 817, 144, 1, 0, 0, 0, 818, 820, 3, 127, 63, 0, 819, 818, 1,
		0, 0, 0, 820, 821, 1, 0, 0, 0, 821, 819, 1, 0, 0, 0, 821, 822, 1, 0, 0,
		0, 822, 146, 1, 0, 0, 0, 823, 825, 7, 13, 0, 0, 824, 826, 7, 12, 0, 0,
		825, 824, 1, 0, 0, 0, 825, 826, 1, 0, 0, 0, 826, 827, 1, 0, 0, 0, 827,
		828, 3, 141, 70, 0, 828, 148, 1, 0, 0, 0, 829, 830, 5, 92, 0, 0, 830, 845,
		7, 14, 0, 0, 831, 832, 5, 92, 0, 0, 832, 834, 3, 125, 62, 0, 833, 835,
		3, 125, 62, 0, 834, 833, 1, 0, 0, 0, 834, 835, 1, 0, 0, 0, 835, 837, 1,
		0, 0, 0, 836, 838, 3, 125, 62, 0, 837, 836, 1, 0, 0, 0, 837, 838, 1, 0,
		0, 0, 838, 845, 1, 0, 0, 0, 839, 840, 5, 92, 0, 0, 840, 841, 5, 120, 0,
		0, 841, 842, 1, 0, 0, 0, 842, 845, 3, 145, 72, 0, 843, 845, 3, 131, 65,
		0, 844, 829, 1, 0, 0, 0, 844, 831, 1, 0, 0, 0, 844, 839, 1, 0, 0, 0, 844,
		843, 1, 0, 0, 0, 845, 150, 1, 0, 0, 0, 846, 848, 7, 15, 0, 0, 847, 846,
		1, 0, 0, 0, 848, 849, 1, 0, 0, 0, 849, 847, 1, 0, 0, 0, 849, 850, 1, 0,
		0, 0, 850, 851, 1, 0, 0, 0, 851, 852, 6, 75, 0, 0, 852, 152, 1, 0, 0, 0,
		853, 855, 5, 13, 0, 0, 854, 856, 5, 10, 0, 0, 855, 854, 1, 0, 0, 0, 855,
		856, 1, 0, 0, 0, 856, 859, 1, 0, 0, 0, 857, 859, 5, 10, 0, 0, 858, 853,
		1, 0, 0, 0, 858, 857, 1, 0, 0, 0, 859, 860, 1, 0, 0, 0, 860, 861, 6, 76,
		0, 0, 861, 154, 1, 0, 0, 0, 59, 0, 193, 207, 229, 255, 290, 298, 314, 338,
		349, 355, 360, 362, 393, 429, 465, 495, 533, 571, 597, 626, 632, 636, 641,
		643, 653, 657, 662, 665, 669, 674, 680, 685, 690, 695, 704, 713, 724, 730,
		734, 740, 768, 772, 777, 783, 788, 795, 799, 806, 809, 816, 821, 825, 834,
		837, 844, 849, 855, 858, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PlanLexerInit initializes any static state used to implement PlanLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPlanLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PlanLexerInit() {
	staticData := &PlanLexerLexerStaticData
	staticData.once.Do(planlexerLexerInit)
}

// NewPlanLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPlanLexer(input antlr.CharStream) *PlanLexer {
	PlanLexerInit()
	l := new(PlanLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PlanLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Plan.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PlanLexer tokens.
const (
	PlanLexerT__0             = 1
	PlanLexerT__1             = 2
	PlanLexerT__2             = 3
	PlanLexerT__3             = 4
	PlanLexerT__4             = 5
	PlanLexerLBRACE           = 6
	PlanLexerRBRACE           = 7
	PlanLexerLT               = 8
	PlanLexerLE               = 9
	PlanLexerGT               = 10
	PlanLexerGE               = 11
	PlanLexerEQ               = 12
	PlanLexerNE               = 13
	PlanLexerLIKE             = 14
	PlanLexerEXISTS           = 15
	PlanLexerTEXTMATCH        = 16
	PlanLexerPHRASEMATCH      = 17
	PlanLexerADD              = 18
	PlanLexerSUB              = 19
	PlanLexerMUL              = 20
	PlanLexerDIV              = 21
	PlanLexerMOD              = 22
	PlanLexerPOW              = 23
	PlanLexerSHL              = 24
	PlanLexerSHR              = 25
	PlanLexerBAND             = 26
	PlanLexerBOR              = 27
	PlanLexerBXOR             = 28
	PlanLexerAND              = 29
	PlanLexerOR               = 30
	PlanLexerISNULL           = 31
	PlanLexerISNOTNULL        = 32
	PlanLexerBNOT             = 33
	PlanLexerNOT              = 34
	PlanLexerIN               = 35
	PlanLexerEmptyArray       = 36
	PlanLexerJSONContains     = 37
	PlanLexerJSONContainsAll  = 38
	PlanLexerJSONContainsAny  = 39
	PlanLexerArrayContains    = 40
	PlanLexerArrayContainsAll = 41
	PlanLexerArrayContainsAny = 42
	PlanLexerArrayLength      = 43
	PlanLexerBooleanConstant  = 44
	PlanLexerIntegerConstant  = 45
	PlanLexerFloatingConstant = 46
	PlanLexerIdentifier       = 47
	PlanLexerMeta             = 48
	PlanLexerStringLiteral    = 49
	PlanLexerJSONIdentifier   = 50
	PlanLexerWhitespace       = 51
	PlanLexerNewline          = 52
)
