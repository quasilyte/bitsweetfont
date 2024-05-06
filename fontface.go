// Code generated by fontget, DO NOT EDIT

package bitsweetfont

import (
	_ "embed"
	"golang.org/x/image/font"
)

// runeAndIndex is a compact version for fonts with low max rune bound.
// Its size is 4 bytes instead of 8.
type runeAndIndex struct {
	r uint16
	i uint16
}

// New1 allocates a font of size=1.
//
// Allocating several instances of the same-sized font is unadvised
// as they would not share the resources.
func New1() font.Face {
	data := uncompress(size1_00data)
	img := newBitmapImage(data, 6, 12)
	f := newBitmapFont(0, img, 0, 10)
	f.MinRune = 32
	f.MaxRune = 10064
	f.XHeight = 6
	f.CapHeight = 9
	f.GlyphBitSize = 72
	f.RuneMapping = size1_00mapping[:]
	return f
}

// New1_3 allocates a font of size=1.3.
//
// Allocating several instances of the same-sized font is unadvised
// as they would not share the resources.
func New1_3() font.Face {
	data := uncompress(size1_30data)
	img := newBitmapImage(data, 8, 16)
	f := newBitmapFont(1, img, 0, 13)
	f.MinRune = 32
	f.MaxRune = 10064
	f.XHeight = 8
	f.CapHeight = 11
	f.GlyphBitSize = 128
	f.RuneMapping = size1_30mapping[:]
	return f
}

var (
	//go:embed 1_00.data.gz
	size1_00data []byte

	//go:embed 1_30.data.gz
	size1_30data []byte
)

const (
	onMissing = "stub"
)

func getStubImageIndex(id int) uint {
	switch id {
	case 0: // size=1_00
		return 154
	case 1: // size=1_30
		return 151
	default:
		panic("unreachable")
	}
}

var (
	// len=168 sizeApprox=672
	size1_00mapping = [...]runeAndIndex{
		{r: 32, i: 0},      // ' '
		{r: 33, i: 1},      // '!'
		{r: 34, i: 2},      // '"'
		{r: 35, i: 3},      // '#'
		{r: 36, i: 4},      // '$'
		{r: 37, i: 5},      // '%'
		{r: 38, i: 6},      // '&'
		{r: 39, i: 7},      // '\''
		{r: 40, i: 8},      // '('
		{r: 41, i: 9},      // ')'
		{r: 42, i: 10},     // '*'
		{r: 43, i: 11},     // '+'
		{r: 44, i: 12},     // ','
		{r: 45, i: 13},     // '-'
		{r: 46, i: 14},     // '.'
		{r: 47, i: 15},     // '/'
		{r: 48, i: 16},     // '0'
		{r: 49, i: 17},     // '1'
		{r: 50, i: 18},     // '2'
		{r: 51, i: 19},     // '3'
		{r: 52, i: 20},     // '4'
		{r: 53, i: 21},     // '5'
		{r: 54, i: 22},     // '6'
		{r: 55, i: 23},     // '7'
		{r: 56, i: 24},     // '8'
		{r: 57, i: 25},     // '9'
		{r: 58, i: 26},     // ':'
		{r: 60, i: 27},     // '<'
		{r: 61, i: 28},     // '='
		{r: 62, i: 29},     // '>'
		{r: 63, i: 30},     // '?'
		{r: 64, i: 31},     // '@'
		{r: 65, i: 32},     // 'A'
		{r: 66, i: 33},     // 'B'
		{r: 67, i: 34},     // 'C'
		{r: 68, i: 35},     // 'D'
		{r: 69, i: 36},     // 'E'
		{r: 70, i: 37},     // 'F'
		{r: 71, i: 38},     // 'G'
		{r: 72, i: 39},     // 'H'
		{r: 73, i: 40},     // 'I'
		{r: 74, i: 41},     // 'J'
		{r: 75, i: 42},     // 'K'
		{r: 76, i: 43},     // 'L'
		{r: 77, i: 44},     // 'M'
		{r: 78, i: 45},     // 'N'
		{r: 79, i: 46},     // 'O'
		{r: 80, i: 47},     // 'P'
		{r: 81, i: 48},     // 'Q'
		{r: 82, i: 49},     // 'R'
		{r: 83, i: 50},     // 'S'
		{r: 84, i: 51},     // 'T'
		{r: 85, i: 52},     // 'U'
		{r: 86, i: 53},     // 'V'
		{r: 87, i: 54},     // 'W'
		{r: 88, i: 55},     // 'X'
		{r: 89, i: 56},     // 'Y'
		{r: 90, i: 57},     // 'Z'
		{r: 91, i: 58},     // '['
		{r: 92, i: 59},     // '\\'
		{r: 93, i: 60},     // ']'
		{r: 94, i: 61},     // '^'
		{r: 95, i: 62},     // '_'
		{r: 96, i: 63},     // '`'
		{r: 97, i: 64},     // 'a'
		{r: 98, i: 65},     // 'b'
		{r: 99, i: 66},     // 'c'
		{r: 100, i: 67},    // 'd'
		{r: 101, i: 68},    // 'e'
		{r: 102, i: 69},    // 'f'
		{r: 103, i: 70},    // 'g'
		{r: 104, i: 71},    // 'h'
		{r: 105, i: 72},    // 'i'
		{r: 106, i: 73},    // 'j'
		{r: 107, i: 74},    // 'k'
		{r: 108, i: 75},    // 'l'
		{r: 109, i: 76},    // 'm'
		{r: 110, i: 77},    // 'n'
		{r: 111, i: 78},    // 'o'
		{r: 112, i: 79},    // 'p'
		{r: 113, i: 80},    // 'q'
		{r: 114, i: 81},    // 'r'
		{r: 115, i: 82},    // 's'
		{r: 116, i: 83},    // 't'
		{r: 117, i: 84},    // 'u'
		{r: 118, i: 85},    // 'v'
		{r: 119, i: 86},    // 'w'
		{r: 120, i: 87},    // 'x'
		{r: 121, i: 88},    // 'y'
		{r: 122, i: 89},    // 'z'
		{r: 123, i: 90},    // '{'
		{r: 124, i: 91},    // '|'
		{r: 125, i: 92},    // '}'
		{r: 126, i: 93},    // '~'
		{r: 215, i: 94},    // '×'
		{r: 1025, i: 95},   // 'Ё'
		{r: 1040, i: 32},   // 'А'
		{r: 1041, i: 96},   // 'Б'
		{r: 1042, i: 33},   // 'В'
		{r: 1043, i: 97},   // 'Г'
		{r: 1044, i: 98},   // 'Д'
		{r: 1045, i: 36},   // 'Е'
		{r: 1046, i: 99},   // 'Ж'
		{r: 1047, i: 100},  // 'З'
		{r: 1048, i: 101},  // 'И'
		{r: 1049, i: 102},  // 'Й'
		{r: 1050, i: 42},   // 'К'
		{r: 1051, i: 103},  // 'Л'
		{r: 1052, i: 104},  // 'М'
		{r: 1053, i: 39},   // 'Н'
		{r: 1054, i: 46},   // 'О'
		{r: 1055, i: 105},  // 'П'
		{r: 1056, i: 106},  // 'Р'
		{r: 1057, i: 34},   // 'С'
		{r: 1058, i: 107},  // 'Т'
		{r: 1059, i: 108},  // 'У'
		{r: 1060, i: 109},  // 'Ф'
		{r: 1061, i: 55},   // 'Х'
		{r: 1062, i: 110},  // 'Ц'
		{r: 1063, i: 111},  // 'Ч'
		{r: 1064, i: 112},  // 'Ш'
		{r: 1065, i: 113},  // 'Щ'
		{r: 1066, i: 114},  // 'Ъ'
		{r: 1067, i: 115},  // 'Ы'
		{r: 1068, i: 116},  // 'Ь'
		{r: 1069, i: 117},  // 'Э'
		{r: 1070, i: 118},  // 'Ю'
		{r: 1071, i: 119},  // 'Я'
		{r: 1072, i: 64},   // 'а'
		{r: 1073, i: 120},  // 'б'
		{r: 1074, i: 121},  // 'в'
		{r: 1075, i: 122},  // 'г'
		{r: 1076, i: 123},  // 'д'
		{r: 1077, i: 68},   // 'е'
		{r: 1078, i: 124},  // 'ж'
		{r: 1079, i: 125},  // 'з'
		{r: 1080, i: 126},  // 'и'
		{r: 1081, i: 127},  // 'й'
		{r: 1082, i: 128},  // 'к'
		{r: 1083, i: 129},  // 'л'
		{r: 1084, i: 130},  // 'м'
		{r: 1085, i: 131},  // 'н'
		{r: 1086, i: 78},   // 'о'
		{r: 1087, i: 132},  // 'п'
		{r: 1088, i: 79},   // 'р'
		{r: 1089, i: 66},   // 'с'
		{r: 1090, i: 133},  // 'т'
		{r: 1091, i: 88},   // 'у'
		{r: 1092, i: 134},  // 'ф'
		{r: 1093, i: 135},  // 'х'
		{r: 1094, i: 136},  // 'ц'
		{r: 1095, i: 137},  // 'ч'
		{r: 1096, i: 138},  // 'ш'
		{r: 1097, i: 139},  // 'щ'
		{r: 1098, i: 140},  // 'ъ'
		{r: 1099, i: 141},  // 'ы'
		{r: 1100, i: 142},  // 'ь'
		{r: 1101, i: 143},  // 'э'
		{r: 1102, i: 144},  // 'ю'
		{r: 1103, i: 145},  // 'я'
		{r: 1105, i: 146},  // 'ё'
		{r: 9633, i: 147},  // '□'
		{r: 9651, i: 148},  // '△'
		{r: 9675, i: 149},  // '○'
		{r: 9676, i: 150},  // '◌'
		{r: 9679, i: 151},  // '●'
		{r: 9776, i: 152},  // '☰'
		{r: 10064, i: 153}, // '❐'
	}

	// len=168 sizeApprox=672
	size1_30mapping = [...]runeAndIndex{
		{r: 32, i: 0},      // ' '
		{r: 33, i: 1},      // '!'
		{r: 34, i: 2},      // '"'
		{r: 35, i: 3},      // '#'
		{r: 36, i: 4},      // '$'
		{r: 37, i: 5},      // '%'
		{r: 38, i: 6},      // '&'
		{r: 39, i: 7},      // '\''
		{r: 40, i: 8},      // '('
		{r: 41, i: 9},      // ')'
		{r: 42, i: 10},     // '*'
		{r: 43, i: 11},     // '+'
		{r: 44, i: 12},     // ','
		{r: 45, i: 13},     // '-'
		{r: 46, i: 14},     // '.'
		{r: 47, i: 15},     // '/'
		{r: 48, i: 16},     // '0'
		{r: 49, i: 17},     // '1'
		{r: 50, i: 18},     // '2'
		{r: 51, i: 19},     // '3'
		{r: 52, i: 20},     // '4'
		{r: 53, i: 21},     // '5'
		{r: 54, i: 22},     // '6'
		{r: 55, i: 23},     // '7'
		{r: 56, i: 24},     // '8'
		{r: 57, i: 25},     // '9'
		{r: 58, i: 26},     // ':'
		{r: 60, i: 27},     // '<'
		{r: 61, i: 28},     // '='
		{r: 62, i: 29},     // '>'
		{r: 63, i: 30},     // '?'
		{r: 64, i: 31},     // '@'
		{r: 65, i: 32},     // 'A'
		{r: 66, i: 33},     // 'B'
		{r: 67, i: 34},     // 'C'
		{r: 68, i: 35},     // 'D'
		{r: 69, i: 36},     // 'E'
		{r: 70, i: 37},     // 'F'
		{r: 71, i: 38},     // 'G'
		{r: 72, i: 39},     // 'H'
		{r: 73, i: 40},     // 'I'
		{r: 74, i: 41},     // 'J'
		{r: 75, i: 42},     // 'K'
		{r: 76, i: 43},     // 'L'
		{r: 77, i: 44},     // 'M'
		{r: 78, i: 45},     // 'N'
		{r: 79, i: 46},     // 'O'
		{r: 80, i: 47},     // 'P'
		{r: 81, i: 48},     // 'Q'
		{r: 82, i: 49},     // 'R'
		{r: 83, i: 50},     // 'S'
		{r: 84, i: 51},     // 'T'
		{r: 85, i: 52},     // 'U'
		{r: 86, i: 53},     // 'V'
		{r: 87, i: 54},     // 'W'
		{r: 88, i: 55},     // 'X'
		{r: 89, i: 56},     // 'Y'
		{r: 90, i: 57},     // 'Z'
		{r: 91, i: 58},     // '['
		{r: 92, i: 59},     // '\\'
		{r: 93, i: 60},     // ']'
		{r: 94, i: 61},     // '^'
		{r: 95, i: 62},     // '_'
		{r: 96, i: 63},     // '`'
		{r: 97, i: 64},     // 'a'
		{r: 98, i: 65},     // 'b'
		{r: 99, i: 66},     // 'c'
		{r: 100, i: 67},    // 'd'
		{r: 101, i: 68},    // 'e'
		{r: 102, i: 69},    // 'f'
		{r: 103, i: 70},    // 'g'
		{r: 104, i: 71},    // 'h'
		{r: 105, i: 72},    // 'i'
		{r: 106, i: 73},    // 'j'
		{r: 107, i: 74},    // 'k'
		{r: 108, i: 75},    // 'l'
		{r: 109, i: 76},    // 'm'
		{r: 110, i: 77},    // 'n'
		{r: 111, i: 78},    // 'o'
		{r: 112, i: 79},    // 'p'
		{r: 113, i: 80},    // 'q'
		{r: 114, i: 81},    // 'r'
		{r: 115, i: 82},    // 's'
		{r: 116, i: 83},    // 't'
		{r: 117, i: 84},    // 'u'
		{r: 118, i: 85},    // 'v'
		{r: 119, i: 86},    // 'w'
		{r: 120, i: 87},    // 'x'
		{r: 121, i: 88},    // 'y'
		{r: 122, i: 89},    // 'z'
		{r: 123, i: 90},    // '{'
		{r: 124, i: 91},    // '|'
		{r: 125, i: 92},    // '}'
		{r: 126, i: 93},    // '~'
		{r: 215, i: 94},    // '×'
		{r: 1025, i: 95},   // 'Ё'
		{r: 1040, i: 32},   // 'А'
		{r: 1041, i: 96},   // 'Б'
		{r: 1042, i: 33},   // 'В'
		{r: 1043, i: 97},   // 'Г'
		{r: 1044, i: 98},   // 'Д'
		{r: 1045, i: 36},   // 'Е'
		{r: 1046, i: 99},   // 'Ж'
		{r: 1047, i: 100},  // 'З'
		{r: 1048, i: 101},  // 'И'
		{r: 1049, i: 102},  // 'Й'
		{r: 1050, i: 42},   // 'К'
		{r: 1051, i: 103},  // 'Л'
		{r: 1052, i: 44},   // 'М'
		{r: 1053, i: 39},   // 'Н'
		{r: 1054, i: 46},   // 'О'
		{r: 1055, i: 104},  // 'П'
		{r: 1056, i: 47},   // 'Р'
		{r: 1057, i: 34},   // 'С'
		{r: 1058, i: 105},  // 'Т'
		{r: 1059, i: 106},  // 'У'
		{r: 1060, i: 107},  // 'Ф'
		{r: 1061, i: 55},   // 'Х'
		{r: 1062, i: 108},  // 'Ц'
		{r: 1063, i: 109},  // 'Ч'
		{r: 1064, i: 110},  // 'Ш'
		{r: 1065, i: 111},  // 'Щ'
		{r: 1066, i: 112},  // 'Ъ'
		{r: 1067, i: 113},  // 'Ы'
		{r: 1068, i: 114},  // 'Ь'
		{r: 1069, i: 115},  // 'Э'
		{r: 1070, i: 116},  // 'Ю'
		{r: 1071, i: 117},  // 'Я'
		{r: 1072, i: 64},   // 'а'
		{r: 1073, i: 118},  // 'б'
		{r: 1074, i: 119},  // 'в'
		{r: 1075, i: 120},  // 'г'
		{r: 1076, i: 121},  // 'д'
		{r: 1077, i: 68},   // 'е'
		{r: 1078, i: 122},  // 'ж'
		{r: 1079, i: 123},  // 'з'
		{r: 1080, i: 124},  // 'и'
		{r: 1081, i: 125},  // 'й'
		{r: 1082, i: 126},  // 'к'
		{r: 1083, i: 127},  // 'л'
		{r: 1084, i: 128},  // 'м'
		{r: 1085, i: 129},  // 'н'
		{r: 1086, i: 78},   // 'о'
		{r: 1087, i: 130},  // 'п'
		{r: 1088, i: 79},   // 'р'
		{r: 1089, i: 66},   // 'с'
		{r: 1090, i: 131},  // 'т'
		{r: 1091, i: 88},   // 'у'
		{r: 1092, i: 132},  // 'ф'
		{r: 1093, i: 87},   // 'х'
		{r: 1094, i: 133},  // 'ц'
		{r: 1095, i: 134},  // 'ч'
		{r: 1096, i: 135},  // 'ш'
		{r: 1097, i: 136},  // 'щ'
		{r: 1098, i: 137},  // 'ъ'
		{r: 1099, i: 138},  // 'ы'
		{r: 1100, i: 139},  // 'ь'
		{r: 1101, i: 140},  // 'э'
		{r: 1102, i: 141},  // 'ю'
		{r: 1103, i: 142},  // 'я'
		{r: 1105, i: 143},  // 'ё'
		{r: 9633, i: 144},  // '□'
		{r: 9651, i: 145},  // '△'
		{r: 9675, i: 146},  // '○'
		{r: 9676, i: 147},  // '◌'
		{r: 9679, i: 148},  // '●'
		{r: 9776, i: 149},  // '☰'
		{r: 10064, i: 150}, // '❐'
	}
)
