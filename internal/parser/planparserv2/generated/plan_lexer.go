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
		"'>'", "'>='", "'=='", "'!='", "", "", "", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'**'", "'<<'", "'>>'", "'&'", "'|'", "'^'", "", "", "", "",
		"'~'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'$meta'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "LBRACE", "RBRACE", "LT", "LE", "GT", "GE",
		"EQ", "NE", "LIKE", "EXISTS", "TEXTMATCH", "ADD", "SUB", "MUL", "DIV",
		"MOD", "POW", "SHL", "SHR", "BAND", "BOR", "BXOR", "AND", "OR", "ISNULL",
		"ISNOTNULL", "BNOT", "NOT", "IN", "EmptyArray", "JSONContains", "JSONContainsAll",
		"JSONContainsAny", "ArrayContains", "ArrayContainsAll", "ArrayContainsAny",
		"ArrayLength", "BooleanConstant", "IntegerConstant", "FloatingConstant",
		"Identifier", "Meta", "StringLiteral", "JSONIdentifier", "Whitespace",
		"Newline",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "LBRACE", "RBRACE", "LT", "LE",
		"GT", "GE", "EQ", "NE", "LIKE", "EXISTS", "TEXTMATCH", "ADD", "SUB",
		"MUL", "DIV", "MOD", "POW", "SHL", "SHR", "BAND", "BOR", "BXOR", "AND",
		"OR", "ISNULL", "ISNOTNULL", "BNOT", "NOT", "IN", "EmptyArray", "JSONContains",
		"JSONContainsAll", "JSONContainsAny", "ArrayContains", "ArrayContainsAll",
		"ArrayContainsAny", "ArrayLength", "BooleanConstant", "IntegerConstant",
		"FloatingConstant", "Identifier", "Meta", "StringLiteral", "JSONIdentifier",
		"EncodingPrefix", "DoubleSCharSequence", "SingleSCharSequence", "DoubleSChar",
		"SingleSChar", "Nondigit", "Digit", "BinaryConstant", "DecimalConstant",
		"OctalConstant", "HexadecimalConstant", "NonzeroDigit", "OctalDigit",
		"HexadecimalDigit", "HexQuad", "UniversalCharacterName", "DecimalFloatingConstant",
		"HexadecimalFloatingConstant", "FractionalConstant", "ExponentPart",
		"DigitSequence", "HexadecimalFractionalConstant", "HexadecimalDigitSequence",
		"BinaryExponentPart", "EscapeSequence", "Whitespace", "Newline",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 51, 834, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13,
		192, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 3, 14, 206, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 228, 8, 15, 1, 16, 1, 16,
		1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3,
		27, 263, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 271, 8,
		28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 287, 8, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 311, 8,
		30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32,
		322, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 328, 8, 33, 1, 34, 1, 34,
		1, 34, 5, 34, 333, 8, 34, 10, 34, 12, 34, 336, 9, 34, 1, 34, 1, 34, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 366, 8, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3,
		36, 402, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 438, 8, 37, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 468, 8, 38, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 3, 39, 506, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1,
		40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 544,
		8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 570, 8, 41, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 3, 42, 599, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 3, 43,
		605, 8, 43, 1, 44, 1, 44, 3, 44, 609, 8, 44, 1, 45, 1, 45, 1, 45, 5, 45,
		614, 8, 45, 10, 45, 12, 45, 617, 9, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 47, 3, 47, 626, 8, 47, 1, 47, 1, 47, 3, 47, 630, 8, 47, 1,
		47, 1, 47, 1, 47, 3, 47, 635, 8, 47, 1, 47, 3, 47, 638, 8, 47, 1, 48, 1,
		48, 3, 48, 642, 8, 48, 1, 48, 1, 48, 1, 48, 3, 48, 647, 8, 48, 1, 48, 1,
		48, 4, 48, 651, 8, 48, 11, 48, 12, 48, 652, 1, 49, 1, 49, 1, 49, 3, 49,
		658, 8, 49, 1, 50, 4, 50, 661, 8, 50, 11, 50, 12, 50, 662, 1, 51, 4, 51,
		666, 8, 51, 11, 51, 12, 51, 667, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1,
		52, 1, 52, 3, 52, 677, 8, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		1, 53, 3, 53, 686, 8, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1,
		56, 4, 56, 695, 8, 56, 11, 56, 12, 56, 696, 1, 57, 1, 57, 5, 57, 701, 8,
		57, 10, 57, 12, 57, 704, 9, 57, 1, 57, 3, 57, 707, 8, 57, 1, 58, 1, 58,
		5, 58, 711, 8, 58, 10, 58, 12, 58, 714, 9, 58, 1, 59, 1, 59, 1, 59, 1,
		59, 1, 60, 1, 60, 1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63,
		1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1,
		64, 3, 64, 741, 8, 64, 1, 65, 1, 65, 3, 65, 745, 8, 65, 1, 65, 1, 65, 1,
		65, 3, 65, 750, 8, 65, 1, 66, 1, 66, 1, 66, 1, 66, 3, 66, 756, 8, 66, 1,
		66, 1, 66, 1, 67, 3, 67, 761, 8, 67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 67,
		3, 67, 768, 8, 67, 1, 68, 1, 68, 3, 68, 772, 8, 68, 1, 68, 1, 68, 1, 69,
		4, 69, 777, 8, 69, 11, 69, 12, 69, 778, 1, 70, 3, 70, 782, 8, 70, 1, 70,
		1, 70, 1, 70, 1, 70, 1, 70, 3, 70, 789, 8, 70, 1, 71, 4, 71, 792, 8, 71,
		11, 71, 12, 71, 793, 1, 72, 1, 72, 3, 72, 798, 8, 72, 1, 72, 1, 72, 1,
		73, 1, 73, 1, 73, 1, 73, 1, 73, 3, 73, 807, 8, 73, 1, 73, 3, 73, 810, 8,
		73, 1, 73, 1, 73, 1, 73, 1, 73, 1, 73, 3, 73, 817, 8, 73, 1, 74, 4, 74,
		820, 8, 74, 11, 74, 12, 74, 821, 1, 74, 1, 74, 1, 75, 1, 75, 3, 75, 828,
		8, 75, 1, 75, 3, 75, 831, 8, 75, 1, 75, 1, 75, 0, 0, 76, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61,
		31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79,
		40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97,
		49, 99, 0, 101, 0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113, 0, 115,
		0, 117, 0, 119, 0, 121, 0, 123, 0, 125, 0, 127, 0, 129, 0, 131, 0, 133,
		0, 135, 0, 137, 0, 139, 0, 141, 0, 143, 0, 145, 0, 147, 0, 149, 50, 151,
		51, 1, 0, 16, 3, 0, 76, 76, 85, 85, 117, 117, 4, 0, 10, 10, 13, 13, 34,
		34, 92, 92, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 3, 0, 65, 90, 95, 95,
		97, 122, 1, 0, 48, 57, 2, 0, 66, 66, 98, 98, 1, 0, 48, 49, 2, 0, 88, 88,
		120, 120, 1, 0, 49, 57, 1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 2,
		0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 2, 0, 80, 80, 112, 112, 10,
		0, 34, 34, 39, 39, 63, 63, 92, 92, 97, 98, 102, 102, 110, 110, 114, 114,
		116, 116, 118, 118, 2, 0, 9, 9, 32, 32, 880, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65,
		1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0,
		73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0,
		0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0,
		0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0,
		0, 0, 0, 97, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 0, 151, 1, 0, 0, 0, 1, 153,
		1, 0, 0, 0, 3, 155, 1, 0, 0, 0, 5, 157, 1, 0, 0, 0, 7, 159, 1, 0, 0, 0,
		9, 161, 1, 0, 0, 0, 11, 163, 1, 0, 0, 0, 13, 165, 1, 0, 0, 0, 15, 167,
		1, 0, 0, 0, 17, 169, 1, 0, 0, 0, 19, 172, 1, 0, 0, 0, 21, 174, 1, 0, 0,
		0, 23, 177, 1, 0, 0, 0, 25, 180, 1, 0, 0, 0, 27, 191, 1, 0, 0, 0, 29, 205,
		1, 0, 0, 0, 31, 227, 1, 0, 0, 0, 33, 229, 1, 0, 0, 0, 35, 231, 1, 0, 0,
		0, 37, 233, 1, 0, 0, 0, 39, 235, 1, 0, 0, 0, 41, 237, 1, 0, 0, 0, 43, 239,
		1, 0, 0, 0, 45, 242, 1, 0, 0, 0, 47, 245, 1, 0, 0, 0, 49, 248, 1, 0, 0,
		0, 51, 250, 1, 0, 0, 0, 53, 252, 1, 0, 0, 0, 55, 262, 1, 0, 0, 0, 57, 270,
		1, 0, 0, 0, 59, 286, 1, 0, 0, 0, 61, 310, 1, 0, 0, 0, 63, 312, 1, 0, 0,
		0, 65, 321, 1, 0, 0, 0, 67, 327, 1, 0, 0, 0, 69, 329, 1, 0, 0, 0, 71, 365,
		1, 0, 0, 0, 73, 401, 1, 0, 0, 0, 75, 437, 1, 0, 0, 0, 77, 467, 1, 0, 0,
		0, 79, 505, 1, 0, 0, 0, 81, 543, 1, 0, 0, 0, 83, 569, 1, 0, 0, 0, 85, 598,
		1, 0, 0, 0, 87, 604, 1, 0, 0, 0, 89, 608, 1, 0, 0, 0, 91, 610, 1, 0, 0,
		0, 93, 618, 1, 0, 0, 0, 95, 625, 1, 0, 0, 0, 97, 641, 1, 0, 0, 0, 99, 657,
		1, 0, 0, 0, 101, 660, 1, 0, 0, 0, 103, 665, 1, 0, 0, 0, 105, 676, 1, 0,
		0, 0, 107, 685, 1, 0, 0, 0, 109, 687, 1, 0, 0, 0, 111, 689, 1, 0, 0, 0,
		113, 691, 1, 0, 0, 0, 115, 706, 1, 0, 0, 0, 117, 708, 1, 0, 0, 0, 119,
		715, 1, 0, 0, 0, 121, 719, 1, 0, 0, 0, 123, 721, 1, 0, 0, 0, 125, 723,
		1, 0, 0, 0, 127, 725, 1, 0, 0, 0, 129, 740, 1, 0, 0, 0, 131, 749, 1, 0,
		0, 0, 133, 751, 1, 0, 0, 0, 135, 767, 1, 0, 0, 0, 137, 769, 1, 0, 0, 0,
		139, 776, 1, 0, 0, 0, 141, 788, 1, 0, 0, 0, 143, 791, 1, 0, 0, 0, 145,
		795, 1, 0, 0, 0, 147, 816, 1, 0, 0, 0, 149, 819, 1, 0, 0, 0, 151, 830,
		1, 0, 0, 0, 153, 154, 5, 40, 0, 0, 154, 2, 1, 0, 0, 0, 155, 156, 5, 41,
		0, 0, 156, 4, 1, 0, 0, 0, 157, 158, 5, 91, 0, 0, 158, 6, 1, 0, 0, 0, 159,
		160, 5, 44, 0, 0, 160, 8, 1, 0, 0, 0, 161, 162, 5, 93, 0, 0, 162, 10, 1,
		0, 0, 0, 163, 164, 5, 123, 0, 0, 164, 12, 1, 0, 0, 0, 165, 166, 5, 125,
		0, 0, 166, 14, 1, 0, 0, 0, 167, 168, 5, 60, 0, 0, 168, 16, 1, 0, 0, 0,
		169, 170, 5, 60, 0, 0, 170, 171, 5, 61, 0, 0, 171, 18, 1, 0, 0, 0, 172,
		173, 5, 62, 0, 0, 173, 20, 1, 0, 0, 0, 174, 175, 5, 62, 0, 0, 175, 176,
		5, 61, 0, 0, 176, 22, 1, 0, 0, 0, 177, 178, 5, 61, 0, 0, 178, 179, 5, 61,
		0, 0, 179, 24, 1, 0, 0, 0, 180, 181, 5, 33, 0, 0, 181, 182, 5, 61, 0, 0,
		182, 26, 1, 0, 0, 0, 183, 184, 5, 108, 0, 0, 184, 185, 5, 105, 0, 0, 185,
		186, 5, 107, 0, 0, 186, 192, 5, 101, 0, 0, 187, 188, 5, 76, 0, 0, 188,
		189, 5, 73, 0, 0, 189, 190, 5, 75, 0, 0, 190, 192, 5, 69, 0, 0, 191, 183,
		1, 0, 0, 0, 191, 187, 1, 0, 0, 0, 192, 28, 1, 0, 0, 0, 193, 194, 5, 101,
		0, 0, 194, 195, 5, 120, 0, 0, 195, 196, 5, 105, 0, 0, 196, 197, 5, 115,
		0, 0, 197, 198, 5, 116, 0, 0, 198, 206, 5, 115, 0, 0, 199, 200, 5, 69,
		0, 0, 200, 201, 5, 88, 0, 0, 201, 202, 5, 73, 0, 0, 202, 203, 5, 83, 0,
		0, 203, 204, 5, 84, 0, 0, 204, 206, 5, 83, 0, 0, 205, 193, 1, 0, 0, 0,
		205, 199, 1, 0, 0, 0, 206, 30, 1, 0, 0, 0, 207, 208, 5, 116, 0, 0, 208,
		209, 5, 101, 0, 0, 209, 210, 5, 120, 0, 0, 210, 211, 5, 116, 0, 0, 211,
		212, 5, 95, 0, 0, 212, 213, 5, 109, 0, 0, 213, 214, 5, 97, 0, 0, 214, 215,
		5, 116, 0, 0, 215, 216, 5, 99, 0, 0, 216, 228, 5, 104, 0, 0, 217, 218,
		5, 84, 0, 0, 218, 219, 5, 69, 0, 0, 219, 220, 5, 88, 0, 0, 220, 221, 5,
		84, 0, 0, 221, 222, 5, 95, 0, 0, 222, 223, 5, 77, 0, 0, 223, 224, 5, 65,
		0, 0, 224, 225, 5, 84, 0, 0, 225, 226, 5, 67, 0, 0, 226, 228, 5, 72, 0,
		0, 227, 207, 1, 0, 0, 0, 227, 217, 1, 0, 0, 0, 228, 32, 1, 0, 0, 0, 229,
		230, 5, 43, 0, 0, 230, 34, 1, 0, 0, 0, 231, 232, 5, 45, 0, 0, 232, 36,
		1, 0, 0, 0, 233, 234, 5, 42, 0, 0, 234, 38, 1, 0, 0, 0, 235, 236, 5, 47,
		0, 0, 236, 40, 1, 0, 0, 0, 237, 238, 5, 37, 0, 0, 238, 42, 1, 0, 0, 0,
		239, 240, 5, 42, 0, 0, 240, 241, 5, 42, 0, 0, 241, 44, 1, 0, 0, 0, 242,
		243, 5, 60, 0, 0, 243, 244, 5, 60, 0, 0, 244, 46, 1, 0, 0, 0, 245, 246,
		5, 62, 0, 0, 246, 247, 5, 62, 0, 0, 247, 48, 1, 0, 0, 0, 248, 249, 5, 38,
		0, 0, 249, 50, 1, 0, 0, 0, 250, 251, 5, 124, 0, 0, 251, 52, 1, 0, 0, 0,
		252, 253, 5, 94, 0, 0, 253, 54, 1, 0, 0, 0, 254, 255, 5, 38, 0, 0, 255,
		263, 5, 38, 0, 0, 256, 257, 5, 97, 0, 0, 257, 258, 5, 110, 0, 0, 258, 263,
		5, 100, 0, 0, 259, 260, 5, 65, 0, 0, 260, 261, 5, 78, 0, 0, 261, 263, 5,
		68, 0, 0, 262, 254, 1, 0, 0, 0, 262, 256, 1, 0, 0, 0, 262, 259, 1, 0, 0,
		0, 263, 56, 1, 0, 0, 0, 264, 265, 5, 124, 0, 0, 265, 271, 5, 124, 0, 0,
		266, 267, 5, 111, 0, 0, 267, 271, 5, 114, 0, 0, 268, 269, 5, 79, 0, 0,
		269, 271, 5, 82, 0, 0, 270, 264, 1, 0, 0, 0, 270, 266, 1, 0, 0, 0, 270,
		268, 1, 0, 0, 0, 271, 58, 1, 0, 0, 0, 272, 273, 5, 105, 0, 0, 273, 274,
		5, 115, 0, 0, 274, 275, 5, 32, 0, 0, 275, 276, 5, 110, 0, 0, 276, 277,
		5, 117, 0, 0, 277, 278, 5, 108, 0, 0, 278, 287, 5, 108, 0, 0, 279, 280,
		5, 73, 0, 0, 280, 281, 5, 83, 0, 0, 281, 282, 5, 32, 0, 0, 282, 283, 5,
		78, 0, 0, 283, 284, 5, 85, 0, 0, 284, 285, 5, 76, 0, 0, 285, 287, 5, 76,
		0, 0, 286, 272, 1, 0, 0, 0, 286, 279, 1, 0, 0, 0, 287, 60, 1, 0, 0, 0,
		288, 289, 5, 105, 0, 0, 289, 290, 5, 115, 0, 0, 290, 291, 5, 32, 0, 0,
		291, 292, 5, 110, 0, 0, 292, 293, 5, 111, 0, 0, 293, 294, 5, 116, 0, 0,
		294, 295, 5, 32, 0, 0, 295, 296, 5, 110, 0, 0, 296, 297, 5, 117, 0, 0,
		297, 298, 5, 108, 0, 0, 298, 311, 5, 108, 0, 0, 299, 300, 5, 73, 0, 0,
		300, 301, 5, 83, 0, 0, 301, 302, 5, 32, 0, 0, 302, 303, 5, 78, 0, 0, 303,
		304, 5, 79, 0, 0, 304, 305, 5, 84, 0, 0, 305, 306, 5, 32, 0, 0, 306, 307,
		5, 78, 0, 0, 307, 308, 5, 85, 0, 0, 308, 309, 5, 76, 0, 0, 309, 311, 5,
		76, 0, 0, 310, 288, 1, 0, 0, 0, 310, 299, 1, 0, 0, 0, 311, 62, 1, 0, 0,
		0, 312, 313, 5, 126, 0, 0, 313, 64, 1, 0, 0, 0, 314, 322, 5, 33, 0, 0,
		315, 316, 5, 110, 0, 0, 316, 317, 5, 111, 0, 0, 317, 322, 5, 116, 0, 0,
		318, 319, 5, 78, 0, 0, 319, 320, 5, 79, 0, 0, 320, 322, 5, 84, 0, 0, 321,
		314, 1, 0, 0, 0, 321, 315, 1, 0, 0, 0, 321, 318, 1, 0, 0, 0, 322, 66, 1,
		0, 0, 0, 323, 324, 5, 105, 0, 0, 324, 328, 5, 110, 0, 0, 325, 326, 5, 73,
		0, 0, 326, 328, 5, 78, 0, 0, 327, 323, 1, 0, 0, 0, 327, 325, 1, 0, 0, 0,
		328, 68, 1, 0, 0, 0, 329, 334, 5, 91, 0, 0, 330, 333, 3, 149, 74, 0, 331,
		333, 3, 151, 75, 0, 332, 330, 1, 0, 0, 0, 332, 331, 1, 0, 0, 0, 333, 336,
		1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 337, 1, 0,
		0, 0, 336, 334, 1, 0, 0, 0, 337, 338, 5, 93, 0, 0, 338, 70, 1, 0, 0, 0,
		339, 340, 5, 106, 0, 0, 340, 341, 5, 115, 0, 0, 341, 342, 5, 111, 0, 0,
		342, 343, 5, 110, 0, 0, 343, 344, 5, 95, 0, 0, 344, 345, 5, 99, 0, 0, 345,
		346, 5, 111, 0, 0, 346, 347, 5, 110, 0, 0, 347, 348, 5, 116, 0, 0, 348,
		349, 5, 97, 0, 0, 349, 350, 5, 105, 0, 0, 350, 351, 5, 110, 0, 0, 351,
		366, 5, 115, 0, 0, 352, 353, 5, 74, 0, 0, 353, 354, 5, 83, 0, 0, 354, 355,
		5, 79, 0, 0, 355, 356, 5, 78, 0, 0, 356, 357, 5, 95, 0, 0, 357, 358, 5,
		67, 0, 0, 358, 359, 5, 79, 0, 0, 359, 360, 5, 78, 0, 0, 360, 361, 5, 84,
		0, 0, 361, 362, 5, 65, 0, 0, 362, 363, 5, 73, 0, 0, 363, 364, 5, 78, 0,
		0, 364, 366, 5, 83, 0, 0, 365, 339, 1, 0, 0, 0, 365, 352, 1, 0, 0, 0, 366,
		72, 1, 0, 0, 0, 367, 368, 5, 106, 0, 0, 368, 369, 5, 115, 0, 0, 369, 370,
		5, 111, 0, 0, 370, 371, 5, 110, 0, 0, 371, 372, 5, 95, 0, 0, 372, 373,
		5, 99, 0, 0, 373, 374, 5, 111, 0, 0, 374, 375, 5, 110, 0, 0, 375, 376,
		5, 116, 0, 0, 376, 377, 5, 97, 0, 0, 377, 378, 5, 105, 0, 0, 378, 379,
		5, 110, 0, 0, 379, 380, 5, 115, 0, 0, 380, 381, 5, 95, 0, 0, 381, 382,
		5, 97, 0, 0, 382, 383, 5, 108, 0, 0, 383, 402, 5, 108, 0, 0, 384, 385,
		5, 74, 0, 0, 385, 386, 5, 83, 0, 0, 386, 387, 5, 79, 0, 0, 387, 388, 5,
		78, 0, 0, 388, 389, 5, 95, 0, 0, 389, 390, 5, 67, 0, 0, 390, 391, 5, 79,
		0, 0, 391, 392, 5, 78, 0, 0, 392, 393, 5, 84, 0, 0, 393, 394, 5, 65, 0,
		0, 394, 395, 5, 73, 0, 0, 395, 396, 5, 78, 0, 0, 396, 397, 5, 83, 0, 0,
		397, 398, 5, 95, 0, 0, 398, 399, 5, 65, 0, 0, 399, 400, 5, 76, 0, 0, 400,
		402, 5, 76, 0, 0, 401, 367, 1, 0, 0, 0, 401, 384, 1, 0, 0, 0, 402, 74,
		1, 0, 0, 0, 403, 404, 5, 106, 0, 0, 404, 405, 5, 115, 0, 0, 405, 406, 5,
		111, 0, 0, 406, 407, 5, 110, 0, 0, 407, 408, 5, 95, 0, 0, 408, 409, 5,
		99, 0, 0, 409, 410, 5, 111, 0, 0, 410, 411, 5, 110, 0, 0, 411, 412, 5,
		116, 0, 0, 412, 413, 5, 97, 0, 0, 413, 414, 5, 105, 0, 0, 414, 415, 5,
		110, 0, 0, 415, 416, 5, 115, 0, 0, 416, 417, 5, 95, 0, 0, 417, 418, 5,
		97, 0, 0, 418, 419, 5, 110, 0, 0, 419, 438, 5, 121, 0, 0, 420, 421, 5,
		74, 0, 0, 421, 422, 5, 83, 0, 0, 422, 423, 5, 79, 0, 0, 423, 424, 5, 78,
		0, 0, 424, 425, 5, 95, 0, 0, 425, 426, 5, 67, 0, 0, 426, 427, 5, 79, 0,
		0, 427, 428, 5, 78, 0, 0, 428, 429, 5, 84, 0, 0, 429, 430, 5, 65, 0, 0,
		430, 431, 5, 73, 0, 0, 431, 432, 5, 78, 0, 0, 432, 433, 5, 83, 0, 0, 433,
		434, 5, 95, 0, 0, 434, 435, 5, 65, 0, 0, 435, 436, 5, 78, 0, 0, 436, 438,
		5, 89, 0, 0, 437, 403, 1, 0, 0, 0, 437, 420, 1, 0, 0, 0, 438, 76, 1, 0,
		0, 0, 439, 440, 5, 97, 0, 0, 440, 441, 5, 114, 0, 0, 441, 442, 5, 114,
		0, 0, 442, 443, 5, 97, 0, 0, 443, 444, 5, 121, 0, 0, 444, 445, 5, 95, 0,
		0, 445, 446, 5, 99, 0, 0, 446, 447, 5, 111, 0, 0, 447, 448, 5, 110, 0,
		0, 448, 449, 5, 116, 0, 0, 449, 450, 5, 97, 0, 0, 450, 451, 5, 105, 0,
		0, 451, 452, 5, 110, 0, 0, 452, 468, 5, 115, 0, 0, 453, 454, 5, 65, 0,
		0, 454, 455, 5, 82, 0, 0, 455, 456, 5, 82, 0, 0, 456, 457, 5, 65, 0, 0,
		457, 458, 5, 89, 0, 0, 458, 459, 5, 95, 0, 0, 459, 460, 5, 67, 0, 0, 460,
		461, 5, 79, 0, 0, 461, 462, 5, 78, 0, 0, 462, 463, 5, 84, 0, 0, 463, 464,
		5, 65, 0, 0, 464, 465, 5, 73, 0, 0, 465, 466, 5, 78, 0, 0, 466, 468, 5,
		83, 0, 0, 467, 439, 1, 0, 0, 0, 467, 453, 1, 0, 0, 0, 468, 78, 1, 0, 0,
		0, 469, 470, 5, 97, 0, 0, 470, 471, 5, 114, 0, 0, 471, 472, 5, 114, 0,
		0, 472, 473, 5, 97, 0, 0, 473, 474, 5, 121, 0, 0, 474, 475, 5, 95, 0, 0,
		475, 476, 5, 99, 0, 0, 476, 477, 5, 111, 0, 0, 477, 478, 5, 110, 0, 0,
		478, 479, 5, 116, 0, 0, 479, 480, 5, 97, 0, 0, 480, 481, 5, 105, 0, 0,
		481, 482, 5, 110, 0, 0, 482, 483, 5, 115, 0, 0, 483, 484, 5, 95, 0, 0,
		484, 485, 5, 97, 0, 0, 485, 486, 5, 108, 0, 0, 486, 506, 5, 108, 0, 0,
		487, 488, 5, 65, 0, 0, 488, 489, 5, 82, 0, 0, 489, 490, 5, 82, 0, 0, 490,
		491, 5, 65, 0, 0, 491, 492, 5, 89, 0, 0, 492, 493, 5, 95, 0, 0, 493, 494,
		5, 67, 0, 0, 494, 495, 5, 79, 0, 0, 495, 496, 5, 78, 0, 0, 496, 497, 5,
		84, 0, 0, 497, 498, 5, 65, 0, 0, 498, 499, 5, 73, 0, 0, 499, 500, 5, 78,
		0, 0, 500, 501, 5, 83, 0, 0, 501, 502, 5, 95, 0, 0, 502, 503, 5, 65, 0,
		0, 503, 504, 5, 76, 0, 0, 504, 506, 5, 76, 0, 0, 505, 469, 1, 0, 0, 0,
		505, 487, 1, 0, 0, 0, 506, 80, 1, 0, 0, 0, 507, 508, 5, 97, 0, 0, 508,
		509, 5, 114, 0, 0, 509, 510, 5, 114, 0, 0, 510, 511, 5, 97, 0, 0, 511,
		512, 5, 121, 0, 0, 512, 513, 5, 95, 0, 0, 513, 514, 5, 99, 0, 0, 514, 515,
		5, 111, 0, 0, 515, 516, 5, 110, 0, 0, 516, 517, 5, 116, 0, 0, 517, 518,
		5, 97, 0, 0, 518, 519, 5, 105, 0, 0, 519, 520, 5, 110, 0, 0, 520, 521,
		5, 115, 0, 0, 521, 522, 5, 95, 0, 0, 522, 523, 5, 97, 0, 0, 523, 524, 5,
		110, 0, 0, 524, 544, 5, 121, 0, 0, 525, 526, 5, 65, 0, 0, 526, 527, 5,
		82, 0, 0, 527, 528, 5, 82, 0, 0, 528, 529, 5, 65, 0, 0, 529, 530, 5, 89,
		0, 0, 530, 531, 5, 95, 0, 0, 531, 532, 5, 67, 0, 0, 532, 533, 5, 79, 0,
		0, 533, 534, 5, 78, 0, 0, 534, 535, 5, 84, 0, 0, 535, 536, 5, 65, 0, 0,
		536, 537, 5, 73, 0, 0, 537, 538, 5, 78, 0, 0, 538, 539, 5, 83, 0, 0, 539,
		540, 5, 95, 0, 0, 540, 541, 5, 65, 0, 0, 541, 542, 5, 78, 0, 0, 542, 544,
		5, 89, 0, 0, 543, 507, 1, 0, 0, 0, 543, 525, 1, 0, 0, 0, 544, 82, 1, 0,
		0, 0, 545, 546, 5, 97, 0, 0, 546, 547, 5, 114, 0, 0, 547, 548, 5, 114,
		0, 0, 548, 549, 5, 97, 0, 0, 549, 550, 5, 121, 0, 0, 550, 551, 5, 95, 0,
		0, 551, 552, 5, 108, 0, 0, 552, 553, 5, 101, 0, 0, 553, 554, 5, 110, 0,
		0, 554, 555, 5, 103, 0, 0, 555, 556, 5, 116, 0, 0, 556, 570, 5, 104, 0,
		0, 557, 558, 5, 65, 0, 0, 558, 559, 5, 82, 0, 0, 559, 560, 5, 82, 0, 0,
		560, 561, 5, 65, 0, 0, 561, 562, 5, 89, 0, 0, 562, 563, 5, 95, 0, 0, 563,
		564, 5, 76, 0, 0, 564, 565, 5, 69, 0, 0, 565, 566, 5, 78, 0, 0, 566, 567,
		5, 71, 0, 0, 567, 568, 5, 84, 0, 0, 568, 570, 5, 72, 0, 0, 569, 545, 1,
		0, 0, 0, 569, 557, 1, 0, 0, 0, 570, 84, 1, 0, 0, 0, 571, 572, 5, 116, 0,
		0, 572, 573, 5, 114, 0, 0, 573, 574, 5, 117, 0, 0, 574, 599, 5, 101, 0,
		0, 575, 576, 5, 84, 0, 0, 576, 577, 5, 114, 0, 0, 577, 578, 5, 117, 0,
		0, 578, 599, 5, 101, 0, 0, 579, 580, 5, 84, 0, 0, 580, 581, 5, 82, 0, 0,
		581, 582, 5, 85, 0, 0, 582, 599, 5, 69, 0, 0, 583, 584, 5, 102, 0, 0, 584,
		585, 5, 97, 0, 0, 585, 586, 5, 108, 0, 0, 586, 587, 5, 115, 0, 0, 587,
		599, 5, 101, 0, 0, 588, 589, 5, 70, 0, 0, 589, 590, 5, 97, 0, 0, 590, 591,
		5, 108, 0, 0, 591, 592, 5, 115, 0, 0, 592, 599, 5, 101, 0, 0, 593, 594,
		5, 70, 0, 0, 594, 595, 5, 65, 0, 0, 595, 596, 5, 76, 0, 0, 596, 597, 5,
		83, 0, 0, 597, 599, 5, 69, 0, 0, 598, 571, 1, 0, 0, 0, 598, 575, 1, 0,
		0, 0, 598, 579, 1, 0, 0, 0, 598, 583, 1, 0, 0, 0, 598, 588, 1, 0, 0, 0,
		598, 593, 1, 0, 0, 0, 599, 86, 1, 0, 0, 0, 600, 605, 3, 115, 57, 0, 601,
		605, 3, 117, 58, 0, 602, 605, 3, 119, 59, 0, 603, 605, 3, 113, 56, 0, 604,
		600, 1, 0, 0, 0, 604, 601, 1, 0, 0, 0, 604, 602, 1, 0, 0, 0, 604, 603,
		1, 0, 0, 0, 605, 88, 1, 0, 0, 0, 606, 609, 3, 131, 65, 0, 607, 609, 3,
		133, 66, 0, 608, 606, 1, 0, 0, 0, 608, 607, 1, 0, 0, 0, 609, 90, 1, 0,
		0, 0, 610, 615, 3, 109, 54, 0, 611, 614, 3, 109, 54, 0, 612, 614, 3, 111,
		55, 0, 613, 611, 1, 0, 0, 0, 613, 612, 1, 0, 0, 0, 614, 617, 1, 0, 0, 0,
		615, 613, 1, 0, 0, 0, 615, 616, 1, 0, 0, 0, 616, 92, 1, 0, 0, 0, 617, 615,
		1, 0, 0, 0, 618, 619, 5, 36, 0, 0, 619, 620, 5, 109, 0, 0, 620, 621, 5,
		101, 0, 0, 621, 622, 5, 116, 0, 0, 622, 623, 5, 97, 0, 0, 623, 94, 1, 0,
		0, 0, 624, 626, 3, 99, 49, 0, 625, 624, 1, 0, 0, 0, 625, 626, 1, 0, 0,
		0, 626, 637, 1, 0, 0, 0, 627, 629, 5, 34, 0, 0, 628, 630, 3, 101, 50, 0,
		629, 628, 1, 0, 0, 0, 629, 630, 1, 0, 0, 0, 630, 631, 1, 0, 0, 0, 631,
		638, 5, 34, 0, 0, 632, 634, 5, 39, 0, 0, 633, 635, 3, 103, 51, 0, 634,
		633, 1, 0, 0, 0, 634, 635, 1, 0, 0, 0, 635, 636, 1, 0, 0, 0, 636, 638,
		5, 39, 0, 0, 637, 627, 1, 0, 0, 0, 637, 632, 1, 0, 0, 0, 638, 96, 1, 0,
		0, 0, 639, 642, 3, 91, 45, 0, 640, 642, 3, 93, 46, 0, 641, 639, 1, 0, 0,
		0, 641, 640, 1, 0, 0, 0, 642, 650, 1, 0, 0, 0, 643, 646, 5, 91, 0, 0, 644,
		647, 3, 95, 47, 0, 645, 647, 3, 115, 57, 0, 646, 644, 1, 0, 0, 0, 646,
		645, 1, 0, 0, 0, 647, 648, 1, 0, 0, 0, 648, 649, 5, 93, 0, 0, 649, 651,
		1, 0, 0, 0, 650, 643, 1, 0, 0, 0, 651, 652, 1, 0, 0, 0, 652, 650, 1, 0,
		0, 0, 652, 653, 1, 0, 0, 0, 653, 98, 1, 0, 0, 0, 654, 655, 5, 117, 0, 0,
		655, 658, 5, 56, 0, 0, 656, 658, 7, 0, 0, 0, 657, 654, 1, 0, 0, 0, 657,
		656, 1, 0, 0, 0, 658, 100, 1, 0, 0, 0, 659, 661, 3, 105, 52, 0, 660, 659,
		1, 0, 0, 0, 661, 662, 1, 0, 0, 0, 662, 660, 1, 0, 0, 0, 662, 663, 1, 0,
		0, 0, 663, 102, 1, 0, 0, 0, 664, 666, 3, 107, 53, 0, 665, 664, 1, 0, 0,
		0, 666, 667, 1, 0, 0, 0, 667, 665, 1, 0, 0, 0, 667, 668, 1, 0, 0, 0, 668,
		104, 1, 0, 0, 0, 669, 677, 8, 1, 0, 0, 670, 677, 3, 147, 73, 0, 671, 672,
		5, 92, 0, 0, 672, 677, 5, 10, 0, 0, 673, 674, 5, 92, 0, 0, 674, 675, 5,
		13, 0, 0, 675, 677, 5, 10, 0, 0, 676, 669, 1, 0, 0, 0, 676, 670, 1, 0,
		0, 0, 676, 671, 1, 0, 0, 0, 676, 673, 1, 0, 0, 0, 677, 106, 1, 0, 0, 0,
		678, 686, 8, 2, 0, 0, 679, 686, 3, 147, 73, 0, 680, 681, 5, 92, 0, 0, 681,
		686, 5, 10, 0, 0, 682, 683, 5, 92, 0, 0, 683, 684, 5, 13, 0, 0, 684, 686,
		5, 10, 0, 0, 685, 678, 1, 0, 0, 0, 685, 679, 1, 0, 0, 0, 685, 680, 1, 0,
		0, 0, 685, 682, 1, 0, 0, 0, 686, 108, 1, 0, 0, 0, 687, 688, 7, 3, 0, 0,
		688, 110, 1, 0, 0, 0, 689, 690, 7, 4, 0, 0, 690, 112, 1, 0, 0, 0, 691,
		692, 5, 48, 0, 0, 692, 694, 7, 5, 0, 0, 693, 695, 7, 6, 0, 0, 694, 693,
		1, 0, 0, 0, 695, 696, 1, 0, 0, 0, 696, 694, 1, 0, 0, 0, 696, 697, 1, 0,
		0, 0, 697, 114, 1, 0, 0, 0, 698, 702, 3, 121, 60, 0, 699, 701, 3, 111,
		55, 0, 700, 699, 1, 0, 0, 0, 701, 704, 1, 0, 0, 0, 702, 700, 1, 0, 0, 0,
		702, 703, 1, 0, 0, 0, 703, 707, 1, 0, 0, 0, 704, 702, 1, 0, 0, 0, 705,
		707, 5, 48, 0, 0, 706, 698, 1, 0, 0, 0, 706, 705, 1, 0, 0, 0, 707, 116,
		1, 0, 0, 0, 708, 712, 5, 48, 0, 0, 709, 711, 3, 123, 61, 0, 710, 709, 1,
		0, 0, 0, 711, 714, 1, 0, 0, 0, 712, 710, 1, 0, 0, 0, 712, 713, 1, 0, 0,
		0, 713, 118, 1, 0, 0, 0, 714, 712, 1, 0, 0, 0, 715, 716, 5, 48, 0, 0, 716,
		717, 7, 7, 0, 0, 717, 718, 3, 143, 71, 0, 718, 120, 1, 0, 0, 0, 719, 720,
		7, 8, 0, 0, 720, 122, 1, 0, 0, 0, 721, 722, 7, 9, 0, 0, 722, 124, 1, 0,
		0, 0, 723, 724, 7, 10, 0, 0, 724, 126, 1, 0, 0, 0, 725, 726, 3, 125, 62,
		0, 726, 727, 3, 125, 62, 0, 727, 728, 3, 125, 62, 0, 728, 729, 3, 125,
		62, 0, 729, 128, 1, 0, 0, 0, 730, 731, 5, 92, 0, 0, 731, 732, 5, 117, 0,
		0, 732, 733, 1, 0, 0, 0, 733, 741, 3, 127, 63, 0, 734, 735, 5, 92, 0, 0,
		735, 736, 5, 85, 0, 0, 736, 737, 1, 0, 0, 0, 737, 738, 3, 127, 63, 0, 738,
		739, 3, 127, 63, 0, 739, 741, 1, 0, 0, 0, 740, 730, 1, 0, 0, 0, 740, 734,
		1, 0, 0, 0, 741, 130, 1, 0, 0, 0, 742, 744, 3, 135, 67, 0, 743, 745, 3,
		137, 68, 0, 744, 743, 1, 0, 0, 0, 744, 745, 1, 0, 0, 0, 745, 750, 1, 0,
		0, 0, 746, 747, 3, 139, 69, 0, 747, 748, 3, 137, 68, 0, 748, 750, 1, 0,
		0, 0, 749, 742, 1, 0, 0, 0, 749, 746, 1, 0, 0, 0, 750, 132, 1, 0, 0, 0,
		751, 752, 5, 48, 0, 0, 752, 755, 7, 7, 0, 0, 753, 756, 3, 141, 70, 0, 754,
		756, 3, 143, 71, 0, 755, 753, 1, 0, 0, 0, 755, 754, 1, 0, 0, 0, 756, 757,
		1, 0, 0, 0, 757, 758, 3, 145, 72, 0, 758, 134, 1, 0, 0, 0, 759, 761, 3,
		139, 69, 0, 760, 759, 1, 0, 0, 0, 760, 761, 1, 0, 0, 0, 761, 762, 1, 0,
		0, 0, 762, 763, 5, 46, 0, 0, 763, 768, 3, 139, 69, 0, 764, 765, 3, 139,
		69, 0, 765, 766, 5, 46, 0, 0, 766, 768, 1, 0, 0, 0, 767, 760, 1, 0, 0,
		0, 767, 764, 1, 0, 0, 0, 768, 136, 1, 0, 0, 0, 769, 771, 7, 11, 0, 0, 770,
		772, 7, 12, 0, 0, 771, 770, 1, 0, 0, 0, 771, 772, 1, 0, 0, 0, 772, 773,
		1, 0, 0, 0, 773, 774, 3, 139, 69, 0, 774, 138, 1, 0, 0, 0, 775, 777, 3,
		111, 55, 0, 776, 775, 1, 0, 0, 0, 777, 778, 1, 0, 0, 0, 778, 776, 1, 0,
		0, 0, 778, 779, 1, 0, 0, 0, 779, 140, 1, 0, 0, 0, 780, 782, 3, 143, 71,
		0, 781, 780, 1, 0, 0, 0, 781, 782, 1, 0, 0, 0, 782, 783, 1, 0, 0, 0, 783,
		784, 5, 46, 0, 0, 784, 789, 3, 143, 71, 0, 785, 786, 3, 143, 71, 0, 786,
		787, 5, 46, 0, 0, 787, 789, 1, 0, 0, 0, 788, 781, 1, 0, 0, 0, 788, 785,
		1, 0, 0, 0, 789, 142, 1, 0, 0, 0, 790, 792, 3, 125, 62, 0, 791, 790, 1,
		0, 0, 0, 792, 793, 1, 0, 0, 0, 793, 791, 1, 0, 0, 0, 793, 794, 1, 0, 0,
		0, 794, 144, 1, 0, 0, 0, 795, 797, 7, 13, 0, 0, 796, 798, 7, 12, 0, 0,
		797, 796, 1, 0, 0, 0, 797, 798, 1, 0, 0, 0, 798, 799, 1, 0, 0, 0, 799,
		800, 3, 139, 69, 0, 800, 146, 1, 0, 0, 0, 801, 802, 5, 92, 0, 0, 802, 817,
		7, 14, 0, 0, 803, 804, 5, 92, 0, 0, 804, 806, 3, 123, 61, 0, 805, 807,
		3, 123, 61, 0, 806, 805, 1, 0, 0, 0, 806, 807, 1, 0, 0, 0, 807, 809, 1,
		0, 0, 0, 808, 810, 3, 123, 61, 0, 809, 808, 1, 0, 0, 0, 809, 810, 1, 0,
		0, 0, 810, 817, 1, 0, 0, 0, 811, 812, 5, 92, 0, 0, 812, 813, 5, 120, 0,
		0, 813, 814, 1, 0, 0, 0, 814, 817, 3, 143, 71, 0, 815, 817, 3, 129, 64,
		0, 816, 801, 1, 0, 0, 0, 816, 803, 1, 0, 0, 0, 816, 811, 1, 0, 0, 0, 816,
		815, 1, 0, 0, 0, 817, 148, 1, 0, 0, 0, 818, 820, 7, 15, 0, 0, 819, 818,
		1, 0, 0, 0, 820, 821, 1, 0, 0, 0, 821, 819, 1, 0, 0, 0, 821, 822, 1, 0,
		0, 0, 822, 823, 1, 0, 0, 0, 823, 824, 6, 74, 0, 0, 824, 150, 1, 0, 0, 0,
		825, 827, 5, 13, 0, 0, 826, 828, 5, 10, 0, 0, 827, 826, 1, 0, 0, 0, 827,
		828, 1, 0, 0, 0, 828, 831, 1, 0, 0, 0, 829, 831, 5, 10, 0, 0, 830, 825,
		1, 0, 0, 0, 830, 829, 1, 0, 0, 0, 831, 832, 1, 0, 0, 0, 832, 833, 6, 75,
		0, 0, 833, 152, 1, 0, 0, 0, 58, 0, 191, 205, 227, 262, 270, 286, 310, 321,
		327, 332, 334, 365, 401, 437, 467, 505, 543, 569, 598, 604, 608, 613, 615,
		625, 629, 634, 637, 641, 646, 652, 657, 662, 667, 676, 685, 696, 702, 706,
		712, 740, 744, 749, 755, 760, 767, 771, 778, 781, 788, 793, 797, 806, 809,
		816, 821, 827, 830, 1, 6, 0, 0,
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
	PlanLexerADD              = 17
	PlanLexerSUB              = 18
	PlanLexerMUL              = 19
	PlanLexerDIV              = 20
	PlanLexerMOD              = 21
	PlanLexerPOW              = 22
	PlanLexerSHL              = 23
	PlanLexerSHR              = 24
	PlanLexerBAND             = 25
	PlanLexerBOR              = 26
	PlanLexerBXOR             = 27
	PlanLexerAND              = 28
	PlanLexerOR               = 29
	PlanLexerISNULL           = 30
	PlanLexerISNOTNULL        = 31
	PlanLexerBNOT             = 32
	PlanLexerNOT              = 33
	PlanLexerIN               = 34
	PlanLexerEmptyArray       = 35
	PlanLexerJSONContains     = 36
	PlanLexerJSONContainsAll  = 37
	PlanLexerJSONContainsAny  = 38
	PlanLexerArrayContains    = 39
	PlanLexerArrayContainsAll = 40
	PlanLexerArrayContainsAny = 41
	PlanLexerArrayLength      = 42
	PlanLexerBooleanConstant  = 43
	PlanLexerIntegerConstant  = 44
	PlanLexerFloatingConstant = 45
	PlanLexerIdentifier       = 46
	PlanLexerMeta             = 47
	PlanLexerStringLiteral    = 48
	PlanLexerJSONIdentifier   = 49
	PlanLexerWhitespace       = 50
	PlanLexerNewline          = 51
)
