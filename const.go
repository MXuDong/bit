package bit

// Bit is a 'bit' value. The bit is only two values: 1(One), 0(Zero).
// Use true and false of bool to represent it.
type Bit bool

const (
	One  = true  // the bit value : 1
	Zero = false // the bit value : 0

	ByteMaxBit = 8 // the count of bit in a byte max is 8
	ByteMinBit = 0 // the count of bit in a byte min is 0
)
