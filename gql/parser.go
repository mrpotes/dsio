// Code generated by goyacc - DO NOT EDIT.

package gql

import __yyfmt__ "fmt"

import (
	"fmt"
	"strconv"
	"time"

	"cloud.google.com/go/datastore"
)

type Token struct {
	token   int
	literal string
}

type Expression interface{}

type SelectExpr struct {
	Field  FieldExpr
	From   *FromExpr
	Where  []ConditionExpr
	Order  []OrderExpr
	Limit  *LimitExpr
	Offset *OffsetExpr
}

type FieldExpr struct {
	Distinct        bool
	DistinctOnField []string
	Field           []string
	Asterisk        bool
}

type FromExpr struct {
	Kind *KindExpr
}

type ConditionExpr interface {
	GetPropertyName() string
	GetValue() ValueExpr
	GetComparator() ComparatorExpr
}

type IsNullConditionExpr struct {
	PropertyName string
}

func (c IsNullConditionExpr) GetPropertyName() string {
	return c.PropertyName
}

func (c IsNullConditionExpr) GetValue() ValueExpr {
	return ValueExpr{}
}

func (c IsNullConditionExpr) GetComparator() ComparatorExpr {
	return OP_IS_NULL
}

type ForwardConditionExpr struct {
	PropertyName string
	Comparator   ComparatorExpr
	Value        ValueExpr
}

func (c ForwardConditionExpr) GetPropertyName() string {
	return c.PropertyName
}

func (c ForwardConditionExpr) GetValue() ValueExpr {
	return c.Value
}

func (c ForwardConditionExpr) GetComparator() ComparatorExpr {
	return c.Comparator
}

type BackwardConditionExpr struct {
	Value        ValueExpr
	Comparator   ComparatorExpr
	PropertyName string
}

func (c BackwardConditionExpr) GetPropertyName() string {
	return c.PropertyName
}

func (c BackwardConditionExpr) GetValue() ValueExpr {
	return c.Value
}

func (c BackwardConditionExpr) GetComparator() ComparatorExpr {
	return c.Comparator
}

type OrderExpr struct {
	PropertyName string
	Sort         SortType
}

type LimitExpr struct {
	Cursor string
	Number int
}

type OffsetExpr struct {
	Cursor string
	Number int
}

type ResultPositionExpr struct {
	Number      int
	BindingSite string
}

type ValueExpr struct {
	Type ValueType
	V    interface{}
}

type KeyLiteralExpr struct {
	Project   string
	Namespace string
	KeyPath   []KeyPathElementExpr
}

func (k KeyLiteralExpr) ToDatastoreKey(defaultNamespace string) *datastore.Key {
	namespace := defaultNamespace
	if k.Namespace != "" {
		namespace = k.Namespace
	}
	return datastoreKeyOf(namespace, k.KeyPath) // what to do with project ID? Go API doesn't have a field for it
}

func datastoreKeyOf(namespace string, keyPath []KeyPathElementExpr) (k *datastore.Key) {
	if len(keyPath) == 0 {
		return nil
	}
	parent := datastoreKeyOf(namespace, keyPath[:len(keyPath)-1])
	i := keyPath[len(keyPath)-1]
	if i.Name != "" {
		k = datastore.NameKey(i.Kind, i.Name, parent)
	} else {
		k = datastore.IDKey(i.Kind, i.ID, parent)
	}
	if namespace != "" {
		k.Namespace = namespace
	}
	return
}

type BlobLiteralExpr struct {
	Blob string
}

type DatetimeLiteralExpr struct {
	Datetime time.Time
}

type KeyPathElementExpr struct {
	Kind string
	ID   int64
	Name string
}

type KindExpr struct {
	Name string
}

type PropertyNameExpr struct {
	Name string
}

type ComparatorExpr int

func (c ComparatorExpr) String() string {
	switch c {
	case OP_IS_NULL:
		return "IS NULL"
	case OP_CONTAINS:
		return "CONTAINS"
	case OP_HAS_ANCESTOR:
		return "HAS ANCESTOR"
	case OP_IN:
		return "IN"
	case OP_HAS_DESCENDANT:
		return "HAS DESCENDANT"
	case OP_EQUALS:
		return "="
	case OP_LESS:
		return "<"
	case OP_LESS_EQUALS:
		return "<="
	case OP_GREATER:
		return ">"
	case OP_GREATER_EQUALS:
		return ">="
	}
	return ""
}

const (
	OP_IS_NULL ComparatorExpr = 1 << iota
	OP_CONTAINS
	OP_HAS_ANCESTOR
	OP_IN
	OP_HAS_DESCENDANT
	OP_EQUALS         // =
	OP_LESS           // <
	OP_LESS_EQUALS    // <=
	OP_GREATER        // >
	OP_GREATER_EQUALS // >=
)

type ValueType int

const (
	TYPE_BINDING_SITE ValueType = 1 << iota
	TYPE_KEY
	TYPE_BLOB
	TYPE_DATETIME
	TYPE_STRING
	TYPE_INTEGER
	TYPE_DOUBLE
	TYPE_BOOL
	TYPE_NULL
)

type SortType int

const (
	SORT_NONE SortType = iota
	SORT_ASC
	SORT_DESC
)

type yySymType struct {
	yys   int
	token Token
	expr  Expression
}

type yyXError struct {
	state, xsym int
}

const (
	yyDefault      = 57389
	yyEofCode      = 57344
	ANCESTOR       = 57379
	AND            = 57374
	ASC            = 57367
	ASTERISK       = 57354
	BINDING_SITE   = 57350
	BLOB           = 57385
	BY             = 57370
	COMMA          = 57356
	CONTAINS       = 57377
	DATETIME       = 57386
	DESC           = 57368
	DESCENDANT     = 57380
	DISTINCT       = 57363
	DOUBLE         = 57353
	EOF            = 57347
	EQUAL          = 57357
	FALSE          = 57388
	FIRST          = 57372
	FROM           = 57365
	HAS            = 57378
	ILLEGAL        = 57346
	IN             = 57381
	INTEGER        = 57352
	IS             = 57375
	KEY            = 57382
	LEFT_BRACKETS  = 57358
	LEFT_ROUND     = 57360
	LIMIT          = 57371
	NAME           = 57349
	NAMESPACE      = 57384
	NULL           = 57376
	OFFSET         = 57373
	ON             = 57364
	ORDER          = 57369
	PLUS           = 57355
	PROJECT        = 57383
	RIGHT_BRACKETS = 57359
	RIGHT_ROUND    = 57361
	SELECT         = 57362
	STRING         = 57351
	TRUE           = 57387
	WHERE          = 57366
	WS             = 57348
	yyErrCode      = 57345

	yyMaxDepth = 200
	yyTabOfs   = -66
)

var (
	yyPrec = map[int]int{
		PLUS: 0,
	}

	yyXLAT = map[int]int{
		57344: 0,  // $end (48x)
		57373: 1,  // OFFSET (43x)
		57371: 2,  // LIMIT (38x)
		57369: 3,  // ORDER (30x)
		57356: 4,  // COMMA (26x)
		57349: 5,  // NAME (25x)
		57374: 6,  // AND (18x)
		57352: 7,  // INTEGER (17x)
		57350: 8,  // BINDING_SITE (16x)
		57357: 9,  // EQUAL (16x)
		57361: 10, // RIGHT_ROUND (16x)
		57351: 11, // STRING (16x)
		57378: 12, // HAS (14x)
		57358: 13, // LEFT_BRACKETS (14x)
		57359: 14, // RIGHT_BRACKETS (14x)
		57381: 15, // IN (12x)
		57376: 16, // NULL (12x)
		57366: 17, // WHERE (12x)
		57385: 18, // BLOB (11x)
		57386: 19, // DATETIME (11x)
		57353: 20, // DOUBLE (11x)
		57388: 21, // FALSE (11x)
		57382: 22, // KEY (11x)
		57387: 23, // TRUE (11x)
		57408: 24, // property_name (10x)
		57365: 25, // FROM (9x)
		57360: 26, // LEFT_ROUND (7x)
		57411: 27, // result_position (5x)
		57409: 28, // property_names (4x)
		57367: 29, // ASC (3x)
		57368: 30, // DESC (3x)
		57398: 31, // kind (3x)
		57384: 32, // NAMESPACE (3x)
		57355: 33, // PLUS (3x)
		57413: 34, // synthetic_literal (3x)
		57414: 35, // value (3x)
		57354: 36, // ASTERISK (2x)
		57393: 37, // condition (2x)
		57377: 38, // CONTAINS (2x)
		57394: 39, // either_comparator (2x)
		57375: 40, // IS (2x)
		57396: 41, // key_path_element (2x)
		57399: 42, // opt_asc_desc (2x)
		57379: 43, // ANCESTOR (1x)
		57391: 44, // backward_comparator (1x)
		57370: 45, // BY (1x)
		57392: 46, // compound_condition (1x)
		57380: 47, // DESCENDANT (1x)
		57363: 48, // DISTINCT (1x)
		57390: 49, // Field (1x)
		57372: 50, // FIRST (1x)
		57395: 51, // forward_comparator (1x)
		57397: 52, // key_path_elements (1x)
		57364: 53, // ON (1x)
		57400: 54, // opt_from (1x)
		57401: 55, // opt_limit (1x)
		57402: 56, // opt_namespace (1x)
		57403: 57, // opt_offset (1x)
		57404: 58, // opt_order (1x)
		57405: 59, // opt_project (1x)
		57406: 60, // opt_where (1x)
		57407: 61, // orders (1x)
		57383: 62, // PROJECT (1x)
		57410: 63, // query (1x)
		57412: 64, // select (1x)
		57362: 65, // SELECT (1x)
		57389: 66, // $default (0x)
		57347: 67, // EOF (0x)
		57345: 68, // error (0x)
		57346: 69, // ILLEGAL (0x)
		57348: 70, // WS (0x)
	}

	yySymNames = []string{
		"$end",
		"OFFSET",
		"LIMIT",
		"ORDER",
		"COMMA",
		"NAME",
		"AND",
		"INTEGER",
		"BINDING_SITE",
		"EQUAL",
		"RIGHT_ROUND",
		"STRING",
		"HAS",
		"LEFT_BRACKETS",
		"RIGHT_BRACKETS",
		"IN",
		"NULL",
		"WHERE",
		"BLOB",
		"DATETIME",
		"DOUBLE",
		"FALSE",
		"KEY",
		"TRUE",
		"property_name",
		"FROM",
		"LEFT_ROUND",
		"result_position",
		"property_names",
		"ASC",
		"DESC",
		"kind",
		"NAMESPACE",
		"PLUS",
		"synthetic_literal",
		"value",
		"ASTERISK",
		"condition",
		"CONTAINS",
		"either_comparator",
		"IS",
		"key_path_element",
		"opt_asc_desc",
		"ANCESTOR",
		"backward_comparator",
		"BY",
		"compound_condition",
		"DESCENDANT",
		"DISTINCT",
		"Field",
		"FIRST",
		"forward_comparator",
		"key_path_elements",
		"ON",
		"opt_from",
		"opt_limit",
		"opt_namespace",
		"opt_offset",
		"opt_order",
		"opt_project",
		"opt_where",
		"orders",
		"PROJECT",
		"query",
		"select",
		"SELECT",
		"$default",
		"EOF",
		"error",
		"ILLEGAL",
		"WS",
	}

	yyTokenLiteralStrings = map[int]string{}

	yyReductions = map[int]struct{ xsym, components int }{
		0:  {0, 1},
		1:  {63, 1},
		2:  {64, 7},
		3:  {49, 1},
		4:  {49, 1},
		5:  {49, 2},
		6:  {49, 6},
		7:  {49, 6},
		8:  {28, 1},
		9:  {28, 3},
		10: {54, 0},
		11: {54, 2},
		12: {60, 0},
		13: {60, 2},
		14: {46, 1},
		15: {46, 3},
		16: {37, 3},
		17: {37, 3},
		18: {37, 3},
		19: {51, 1},
		20: {51, 1},
		21: {51, 2},
		22: {44, 1},
		23: {44, 1},
		24: {44, 2},
		25: {39, 1},
		26: {39, 1},
		27: {39, 2},
		28: {39, 1},
		29: {39, 2},
		30: {58, 0},
		31: {58, 3},
		32: {61, 2},
		33: {61, 4},
		34: {42, 0},
		35: {42, 1},
		36: {42, 1},
		37: {55, 0},
		38: {55, 2},
		39: {55, 7},
		40: {57, 0},
		41: {57, 2},
		42: {57, 4},
		43: {27, 1},
		44: {27, 1},
		45: {35, 1},
		46: {35, 1},
		47: {35, 1},
		48: {35, 1},
		49: {35, 1},
		50: {35, 1},
		51: {35, 1},
		52: {35, 1},
		53: {34, 6},
		54: {34, 4},
		55: {34, 4},
		56: {59, 0},
		57: {59, 5},
		58: {56, 0},
		59: {56, 5},
		60: {52, 1},
		61: {52, 3},
		62: {41, 3},
		63: {41, 3},
		64: {31, 1},
		65: {24, 1},
	}

	yyXErrors = map[yyXError]string{}

	yyParseTab = [116][]uint16{
		// 0
		{63: 67, 68, 69},
		{66},
		{65},
		{5: 75, 24: 74, 28: 72, 36: 71, 48: 73, 70},
		{56, 56, 56, 56, 17: 56, 25: 86, 54: 85},
		// 5
		{63, 63, 63, 63, 17: 63, 25: 63},
		{62, 62, 62, 62, 81, 17: 62, 25: 62},
		{5: 75, 24: 74, 28: 76, 53: 77},
		{58, 58, 58, 58, 58, 10: 58, 17: 58, 25: 58},
		{1, 1, 1, 1, 1, 6: 1, 9: 1, 1, 12: 1, 1, 1, 17: 1, 25: 1, 29: 1, 1, 38: 1, 40: 1},
		// 10
		{61, 61, 61, 61, 81, 17: 61, 25: 61},
		{26: 78},
		{5: 75, 24: 74, 28: 79},
		{4: 81, 10: 80},
		{5: 75, 24: 74, 28: 84, 36: 83},
		// 15
		{5: 75, 24: 82},
		{57, 57, 57, 57, 57, 10: 57, 17: 57, 25: 57},
		{60, 60, 60, 60, 17: 60, 25: 60},
		{59, 59, 59, 59, 81, 17: 59, 25: 59},
		{54, 54, 54, 54, 17: 90, 60: 89},
		// 20
		{5: 88, 31: 87},
		{55, 55, 55, 55, 17: 55},
		{2, 2, 2, 2, 2, 17: 2},
		{36, 36, 36, 156, 58: 155},
		{5: 75, 7: 98, 95, 11: 97, 16: 102, 18: 104, 105, 99, 101, 103, 100, 93, 34: 96, 94, 37: 92, 46: 91},
		// 25
		{53, 53, 53, 53, 6: 153},
		{52, 52, 52, 52, 6: 52},
		{9: 138, 12: 149, 139, 140, 38: 148, 147, 145, 51: 146},
		{9: 138, 12: 137, 139, 140, 136, 39: 135, 44: 134},
		{21, 21, 21, 21, 6: 21, 9: 21, 12: 21, 21, 21, 21},
		// 30
		{20, 20, 20, 20, 6: 20, 9: 20, 12: 20, 20, 20, 20},
		{19, 19, 19, 19, 6: 19, 9: 19, 12: 19, 19, 19, 19},
		{18, 18, 18, 18, 6: 18, 9: 18, 12: 18, 18, 18, 18},
		{17, 17, 17, 17, 6: 17, 9: 17, 12: 17, 17, 17, 17},
		{16, 16, 16, 16, 6: 16, 9: 16, 12: 16, 16, 16, 16},
		// 35
		{15, 15, 15, 15, 6: 15, 9: 15, 12: 15, 15, 15, 15},
		{14, 14, 14, 14, 6: 14, 9: 14, 12: 14, 14, 14, 14},
		{26: 112},
		{26: 109},
		{26: 106},
		// 40
		{11: 107},
		{10: 108},
		{11, 11, 11, 11, 6: 11, 9: 11, 12: 11, 11, 11, 11},
		{11: 110},
		{10: 111},
		// 45
		{12, 12, 12, 12, 6: 12, 9: 12, 12: 12, 12, 12, 12},
		{5: 10, 32: 10, 59: 113, 62: 114},
		{5: 8, 32: 120, 56: 119},
		{26: 115},
		{11: 116},
		// 50
		{10: 117},
		{4: 118},
		{5: 9, 32: 9},
		{5: 88, 31: 127, 41: 126, 52: 125},
		{26: 121},
		// 55
		{11: 122},
		{10: 123},
		{4: 124},
		{5: 7},
		{4: 132, 10: 131},
		// 60
		{4: 6, 10: 6},
		{4: 128},
		{7: 129, 11: 130},
		{4: 4, 10: 4},
		{4: 3, 10: 3},
		// 65
		{13, 13, 13, 13, 6: 13, 9: 13, 12: 13, 13, 13, 13},
		{5: 88, 31: 127, 41: 133},
		{4: 5, 10: 5},
		{5: 75, 24: 144},
		{5: 44},
		// 70
		{5: 43},
		{47: 143},
		{5: 41, 7: 41, 41, 11: 41, 16: 41, 18: 41, 41, 41, 41, 41, 41},
		{5: 40, 7: 40, 40, 142, 11: 40, 16: 40, 18: 40, 40, 40, 40, 40, 40},
		{5: 38, 7: 38, 38, 141, 11: 38, 16: 38, 18: 38, 38, 38, 38, 38, 38},
		// 75
		{5: 37, 7: 37, 37, 11: 37, 16: 37, 18: 37, 37, 37, 37, 37, 37},
		{5: 39, 7: 39, 39, 11: 39, 16: 39, 18: 39, 39, 39, 39, 39, 39},
		{5: 42},
		{48, 48, 48, 48, 6: 48},
		{16: 152},
		// 80
		{7: 98, 95, 11: 97, 16: 102, 18: 104, 105, 99, 101, 103, 100, 34: 96, 151},
		{7: 47, 47, 11: 47, 16: 47, 18: 47, 47, 47, 47, 47, 47},
		{7: 46, 46, 11: 46, 16: 46, 18: 46, 46, 46, 46, 46, 46},
		{43: 150},
		{7: 45, 45, 11: 45, 16: 45, 18: 45, 45, 45, 45, 45, 45},
		// 85
		{49, 49, 49, 49, 6: 49},
		{50, 50, 50, 50, 6: 50},
		{5: 75, 7: 98, 95, 11: 97, 16: 102, 18: 104, 105, 99, 101, 103, 100, 93, 34: 96, 94, 37: 154},
		{51, 51, 51, 51, 6: 51},
		{29, 29, 167, 55: 166},
		// 90
		{45: 157},
		{5: 75, 24: 159, 61: 158},
		{35, 35, 35, 4: 163},
		{32, 32, 32, 4: 32, 29: 161, 162, 42: 160},
		{34, 34, 34, 4: 34},
		// 95
		{31, 31, 31, 4: 31},
		{30, 30, 30, 4: 30},
		{5: 75, 24: 164},
		{32, 32, 32, 4: 32, 29: 161, 162, 42: 165},
		{33, 33, 33, 4: 33},
		// 100
		{26, 178, 57: 177},
		{7: 171, 170, 27: 168, 50: 169},
		{28, 28},
		{26: 172},
		{23, 23, 4: 23, 10: 23, 33: 23},
		// 105
		{22, 22, 4: 22, 10: 22, 33: 22},
		{7: 171, 170, 27: 173},
		{4: 174},
		{7: 171, 170, 27: 175},
		{10: 176},
		// 110
		{27, 27},
		{64},
		{7: 171, 170, 27: 179},
		{25, 33: 180},
		{7: 171, 170, 27: 181},
		// 115
		{24},
	}
)

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyLexerEx interface {
	yyLexer
	Reduced(rule, state int, lval *yySymType) bool
}

func yySymName(c int) (s string) {
	x, ok := yyXLAT[c]
	if ok {
		return yySymNames[x]
	}

	if c < 0x7f {
		return __yyfmt__.Sprintf("%q", c)
	}

	return __yyfmt__.Sprintf("%d", c)
}

func yylex1(yylex yyLexer, lval *yySymType) (n int) {
	n = yylex.Lex(lval)
	if n <= 0 {
		n = yyEofCode
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("\nlex %s(%#x %d), lval: %+v\n", yySymName(n), n, n, lval)
	}
	return n
}

func yyParse(yylex yyLexer) int {
	const yyError = 68

	yyEx, _ := yylex.(yyLexerEx)
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, 200)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yyerrok := func() {
		if yyDebug >= 2 {
			__yyfmt__.Printf("yyerrok()\n")
		}
		Errflag = 0
	}
	_ = yyerrok
	yystate := 0
	yychar := -1
	var yyxchar int
	var yyshift int
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	if yychar < 0 {
		yylval.yys = yystate
		yychar = yylex1(yylex, &yylval)
		var ok bool
		if yyxchar, ok = yyXLAT[yychar]; !ok {
			yyxchar = len(yySymNames) // > tab width
		}
	}
	if yyDebug >= 4 {
		var a []int
		for _, v := range yyS[:yyp+1] {
			a = append(a, v.yys)
		}
		__yyfmt__.Printf("state stack %v\n", a)
	}
	row := yyParseTab[yystate]
	yyn = 0
	if yyxchar < len(row) {
		if yyn = int(row[yyxchar]); yyn != 0 {
			yyn += yyTabOfs
		}
	}
	switch {
	case yyn > 0: // shift
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		yyshift = yyn
		if yyDebug >= 2 {
			__yyfmt__.Printf("shift, and goto state %d\n", yystate)
		}
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	case yyn < 0: // reduce
	case yystate == 1: // accept
		if yyDebug >= 2 {
			__yyfmt__.Println("accept")
		}
		goto ret0
	}

	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			if yyDebug >= 1 {
				__yyfmt__.Printf("no action for %s in state %d\n", yySymName(yychar), yystate)
			}
			msg, ok := yyXErrors[yyXError{yystate, yyxchar}]
			if !ok {
				msg, ok = yyXErrors[yyXError{yystate, -1}]
			}
			if !ok && yyshift != 0 {
				msg, ok = yyXErrors[yyXError{yyshift, yyxchar}]
			}
			if !ok {
				msg, ok = yyXErrors[yyXError{yyshift, -1}]
			}
			if yychar > 0 {
				ls := yyTokenLiteralStrings[yychar]
				if ls == "" {
					ls = yySymName(yychar)
				}
				if ls != "" {
					switch {
					case msg == "":
						msg = __yyfmt__.Sprintf("unexpected %s", ls)
					default:
						msg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)
					}
				}
			}
			if msg == "" {
				msg = "syntax error"
			}
			yylex.Error(msg)
			Nerrs++
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				row := yyParseTab[yyS[yyp].yys]
				if yyError < len(row) {
					yyn = int(row[yyError]) + yyTabOfs
					if yyn > 0 { // hit
						if yyDebug >= 2 {
							__yyfmt__.Printf("error recovery found error shift in state %d\n", yyS[yyp].yys)
						}
						yystate = yyn /* simulate a shift of "error" */
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery failed\n")
			}
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yySymName(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}

			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	r := -yyn
	x0 := yyReductions[r]
	x, n := x0.xsym, x0.components
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= n
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	exState := yystate
	yystate = int(yyParseTab[yyS[yyp].yys][x]) + yyTabOfs
	/* reduction by production r */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce using rule %v (%s), and goto state %d\n", r, yySymNames[x], yystate)
	}

	switch r {
	case 1:
		{
			yyVAL.expr = yyS[yypt-0].expr
			yylex.(*Lexer).Result = yyVAL.expr
		}
	case 2:
		{
			fieldExpr := yyS[yypt-5].expr
			fromExpr := yyS[yypt-4].expr
			whereExpr := yyS[yypt-3].expr
			orderExpr := yyS[yypt-2].expr
			limitExpr := yyS[yypt-1].expr
			offsetExpr := yyS[yypt-0].expr

			var from *FromExpr
			if fromExpr != nil {
				from = fromExpr.(*FromExpr)
			}

			var limit *LimitExpr
			if limitExpr != nil {
				limit = limitExpr.(*LimitExpr)
			}

			var offset *OffsetExpr
			if offsetExpr != nil {
				offset = offsetExpr.(*OffsetExpr)
			}

			yyVAL.expr = SelectExpr{
				Field:  fieldExpr.(FieldExpr),
				From:   from,
				Where:  whereExpr.([]ConditionExpr),
				Order:  orderExpr.([]OrderExpr),
				Limit:  limit,
				Offset: offset,
			}
		}
	case 3:
		{
			yyVAL.expr = FieldExpr{Asterisk: true}
		}
	case 4:
		{
			yyVAL.expr = FieldExpr{Field: yyS[yypt-0].expr.([]string)}
		}
	case 5:
		{
			yyVAL.expr = FieldExpr{
				Distinct: true,
				Field:    yyS[yypt-0].expr.([]string),
			}
		}
	case 6:
		{
			yyVAL.expr = FieldExpr{
				DistinctOnField: yyS[yypt-2].expr.([]string),
				Asterisk:        true,
			}
		}
	case 7:
		{
			yyVAL.expr = FieldExpr{
				DistinctOnField: yyS[yypt-2].expr.([]string),
				Field:           yyS[yypt-0].expr.([]string),
			}
		}
	case 8:
		{
			yyVAL.expr = []string{yyS[yypt-0].expr.(string)}
		}
	case 9:
		{
			yyVAL.expr = append(yyS[yypt-2].expr.([]string), yyS[yypt-0].expr.(string))
		}
	case 10:
		{
			yyVAL.expr = nil
		}
	case 11:
		{
			kind := yyS[yypt-0].expr.(KindExpr)
			yyVAL.expr = &FromExpr{Kind: &kind}
		}
	case 12:
		{
			yyVAL.expr = make([]ConditionExpr, 0)
		}
	case 13:
		{
			yyVAL.expr = yyS[yypt-0].expr.([]ConditionExpr)
		}
	case 14:
		{
			yyVAL.expr = []ConditionExpr{yyS[yypt-0].expr.(ConditionExpr)}
		}
	case 15:
		{
			yyVAL.expr = append(yyS[yypt-2].expr.([]ConditionExpr), yyS[yypt-0].expr.(ConditionExpr))
		}
	case 16:
		{
			yyVAL.expr = IsNullConditionExpr{
				PropertyName: yyS[yypt-2].expr.(string),
			}
		}
	case 17:
		{
			yyVAL.expr = ForwardConditionExpr{
				PropertyName: yyS[yypt-2].expr.(string),
				Comparator:   yyS[yypt-1].expr.(ComparatorExpr),
				Value:        yyS[yypt-0].expr.(ValueExpr),
			}
		}
	case 18:
		{
			yyVAL.expr = BackwardConditionExpr{
				Value:        yyS[yypt-2].expr.(ValueExpr),
				Comparator:   yyS[yypt-1].expr.(ComparatorExpr),
				PropertyName: yyS[yypt-0].expr.(string),
			}
		}
	case 19:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 20:
		{
			yyVAL.expr = OP_CONTAINS
		}
	case 21:
		{
			yyVAL.expr = OP_HAS_ANCESTOR
		}
	case 22:
		{
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 23:
		{
			yyVAL.expr = OP_IN
		}
	case 24:
		{
			yyVAL.expr = OP_HAS_DESCENDANT
		}
	case 25:
		{
			yyVAL.expr = OP_EQUALS
		}
	case 26:
		{
			yyVAL.expr = OP_LESS
		}
	case 27:
		{
			yyVAL.expr = OP_LESS_EQUALS
		}
	case 28:
		{
			yyVAL.expr = OP_GREATER
		}
	case 29:
		{
			yyVAL.expr = OP_GREATER_EQUALS
		}
	case 30:
		{
			yyVAL.expr = []OrderExpr{}
		}
	case 31:
		{
			yyVAL.expr = yyS[yypt-0].expr.([]OrderExpr)
		}
	case 32:
		{
			yyVAL.expr = []OrderExpr{
				OrderExpr{PropertyName: yyS[yypt-1].expr.(string), Sort: yyS[yypt-0].expr.(SortType)},
			}
		}
	case 33:
		{
			o := OrderExpr{PropertyName: yyS[yypt-1].expr.(string), Sort: yyS[yypt-0].expr.(SortType)}
			yyVAL.expr = append(yyS[yypt-3].expr.([]OrderExpr), o)
		}
	case 34:
		{
			yyVAL.expr = SORT_NONE
		}
	case 35:
		{
			yyVAL.expr = SORT_ASC
		}
	case 36:
		{
			yyVAL.expr = SORT_DESC
		}
	case 37:
		{
			yyVAL.expr = nil
		}
	case 38:
		{
			yyVAL.expr = &LimitExpr{
				Cursor: yyS[yypt-0].expr.(ResultPositionExpr).BindingSite,
				Number: yyS[yypt-0].expr.(ResultPositionExpr).Number,
			}
		}
	case 39:
		{
			yyVAL.expr = &LimitExpr{
				Cursor: yyS[yypt-3].expr.(ResultPositionExpr).BindingSite,
				Number: yyS[yypt-1].expr.(ResultPositionExpr).Number,
			}
		}
	case 40:
		{
			yyVAL.expr = nil
		}
	case 41:
		{
			if yyS[yypt-0].expr.(ResultPositionExpr).BindingSite != "" {
				yyVAL.expr = &OffsetExpr{Cursor: yyS[yypt-0].expr.(ResultPositionExpr).BindingSite}
			} else {
				yyVAL.expr = &OffsetExpr{Number: yyS[yypt-0].expr.(ResultPositionExpr).Number}
			}
		}
	case 42:
		{
			yyVAL.expr = &OffsetExpr{
				Cursor: yyS[yypt-2].expr.(ResultPositionExpr).BindingSite,
				Number: yyS[yypt-0].expr.(ResultPositionExpr).Number,
			}
		}
	case 43:
		{
			yyVAL.expr = ResultPositionExpr{BindingSite: yyS[yypt-0].token.literal}
		}
	case 44:
		{
			number, err := strconv.Atoi(yyS[yypt-0].token.literal)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to integer", yyS[yypt-0].token.literal))
			}
			yyVAL.expr = ResultPositionExpr{Number: number}
		}
	case 45:
		{
			yyVAL.expr = ValueExpr{Type: TYPE_BINDING_SITE, V: yyS[yypt-0].token.literal}
		}
	case 46:
		{
			switch t := yyS[yypt-0].expr.(type) {
			case KeyLiteralExpr:
				yyVAL.expr = ValueExpr{Type: TYPE_KEY, V: yyS[yypt-0].expr}
			case BlobLiteralExpr:
				yyVAL.expr = ValueExpr{Type: TYPE_BLOB, V: yyS[yypt-0].expr}
			case DatetimeLiteralExpr:
				yyVAL.expr = ValueExpr{Type: TYPE_DATETIME, V: t.Datetime}
			default:
				panic(fmt.Sprintf("unkown synthetic_literal:%v", yyS[yypt-0].expr))
			}
		}
	case 47:
		{
			yyVAL.expr = ValueExpr{Type: TYPE_STRING, V: yyS[yypt-0].token.literal}
		}
	case 48:
		{
			number, err := strconv.ParseInt(yyS[yypt-0].token.literal, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to integer", yyS[yypt-0].token.literal))
			}
			yyVAL.expr = ValueExpr{Type: TYPE_INTEGER, V: number}
		}
	case 49:
		{
			double, err := strconv.ParseFloat(yyS[yypt-0].token.literal, 64)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to double", yyS[yypt-0].token.literal))
			}
			yyVAL.expr = ValueExpr{Type: TYPE_DOUBLE, V: double}
		}
	case 50:
		{
			yyVAL.expr = ValueExpr{Type: TYPE_BOOL, V: true}
		}
	case 51:
		{
			yyVAL.expr = ValueExpr{Type: TYPE_BOOL, V: false}
		}
	case 52:
		{
			yyVAL.expr = ValueExpr{Type: TYPE_NULL, V: nil}
		}
	case 53:
		{
			yyVAL.expr = KeyLiteralExpr{
				Project:   yyS[yypt-3].expr.(string),
				Namespace: yyS[yypt-2].expr.(string),
				KeyPath:   yyS[yypt-1].expr.([]KeyPathElementExpr),
			}
		}
	case 54:
		{
			yyVAL.expr = BlobLiteralExpr{Blob: yyS[yypt-1].token.literal}
		}
	case 55:
		{
			t, err := time.Parse(time.RFC3339, yyS[yypt-1].token.literal)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to datime", yyS[yypt-1].token.literal))
			}
			yyVAL.expr = DatetimeLiteralExpr{Datetime: t}
		}
	case 56:
		{
			yyVAL.expr = ""
		}
	case 57:
		{
			yyVAL.expr = yyS[yypt-2].token.literal
		}
	case 58:
		{
			yyVAL.expr = ""
		}
	case 59:
		{
			yyVAL.expr = yyS[yypt-2].token.literal
		}
	case 60:
		{
			yyVAL.expr = []KeyPathElementExpr{yyS[yypt-0].expr.(KeyPathElementExpr)}
		}
	case 61:
		{
			yyVAL.expr = append(yyS[yypt-2].expr.([]KeyPathElementExpr), yyS[yypt-0].expr.(KeyPathElementExpr))
		}
	case 62:
		{
			number, err := strconv.ParseInt(yyS[yypt-0].token.literal, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("can't convert %v to integer", yyS[yypt-0].token.literal))
			}
			yyVAL.expr = KeyPathElementExpr{Kind: yyS[yypt-2].expr.(KindExpr).Name, ID: number}
		}
	case 63:
		{
			yyVAL.expr = KeyPathElementExpr{Kind: yyS[yypt-2].expr.(KindExpr).Name, Name: yyS[yypt-0].token.literal}
		}
	case 64:
		{
			yyVAL.expr = KindExpr{Name: yyS[yypt-0].token.literal}
		}
	case 65:
		{
			yyVAL.expr = yyS[yypt-0].token.literal
		}

	}

	if yyEx != nil && yyEx.Reduced(r, exState, &yyVAL) {
		return -1
	}
	goto yystack /* stack new state and value */
}

type Lexer struct {
	Scanner    *Scanner
	Result     Expression
	parserErr  error
	scannerErr error
}

func (l *Lexer) Lex(lval *yySymType) int {
	token, literal := l.Scanner.Scan()

	if token == EOF {
		return 0

	} else if token == ILLEGAL {
		l.scannerErr = fmt.Errorf("invalid token: %v", literal)
		return 0
	}

	lval.token = Token{token: int(token), literal: literal}
	return int(token)
}

func (l *Lexer) Error(e string) {
	if l.scannerErr == nil {
		l.parserErr = fmt.Errorf("%v: %v\n", e, l.Scanner.Consumed())
	}
}

func (l *Lexer) Parse() error {
	yyParse(l)

	if l.parserErr != nil {
		return l.parserErr
	} else {
		return l.scannerErr
	}
}
