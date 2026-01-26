package morse

const (
	Dot  = 0b0
	Dash = 0b1
)

// ASCIIToMorse converts an ASCII byte to its Morse code representation.
func ASCIIToMorse(b byte) uint8 {
	if (b < ' ') || (b > 'Z') {
		return 0
	}

	return codes[b]
}

const (
	guardBit = 0b1
)

// codes contain the Morse character table as per ITU-R M.1677-1.
// The Morse code representation is stored as LSB first, using an additional bit as a guard bit.
var codes = []uint8{
	' ':  0b00,
	'!':  0b110101, // unsupported?
	'"':  0b1010010,
	'\'': 0b1011110,
	'(':  0b101101,
	')':  0b1101101,
	'+':  0b101010,
	',':  0b1110011,
	'-':  0b1100001,
	'.':  0b1101010,
	'/':  0b101001,
	'0':  0b111111,
	'1':  0b111110,
	'2':  0b111100,
	'3':  0b111000,
	'4':  0b110000,
	'5':  0b100000,
	'6':  0b100001,
	'7':  0b100011,
	'8':  0b100111,
	'9':  0b101111,
	':':  0b1000111,
	'=':  0b110001,
	'?':  0b1001100,
	'@':  0b1010110,
	'A':  0b110,
	'B':  0b10001,
	'C':  0b10101,
	'D':  0b1001,
	'E':  0b10,
	'F':  0b10100,
	'G':  0b1011,
	'H':  0b10000,
	'I':  0b100,
	'J':  0b11110,
	'K':  0b1101,
	'L':  0b10010,
	'M':  0b111,
	'N':  0b101,
	'O':  0b1111,
	'P':  0b10110,
	'Q':  0b11011,
	'R':  0b1010,
	'S':  0b1000,
	'T':  0b11,
	'U':  0b1100,
	'V':  0b11000,
	'W':  0b1110,
	'X':  0b11001,
	'Y':  0b11101,
	'Z':  0b10011,
}
