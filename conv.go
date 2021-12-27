package bit

func BoolToBit(v bool) Bit {
	return Bit(v)
}

// ByteToBits return []Bit, length = 8.
func ByteToBits(byt byte) []Bit {
	bits := make([]Bit, ByteMaxBit)
	for index := 0; index < 8; index++ {
		bits[index] = ReadBit(byt, index)
	}
	return bits
}
func BytesToBits(bytes []byte) []Bit {
	if bytes == nil || len(bytes) == 0 {
		return []Bit{}
	}
	bits := make([]Bit, ByteMaxBit*len(bytes))

	for index := 0; index < len(bytes); index++ {
		byteBits := ByteToBits(bytes[index])
		for i, bit := range byteBits {
			bits[index*8+i] = bit
		}
	}

	return bits
}
