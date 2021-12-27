package bit

import "io"

// ReadBit return bit from special offset of byte.
// If offset outside byt length(0 - 7), return zero.
// Byte 0b10000000 get by offset - 0 return One, other is Zero
func ReadBit(byt byte, offset int) Bit {
	if offset < ByteMinBit || offset >= ByteMaxBit {
		return Zero
	}
	return (byt >> (ByteMaxBit - 1 - offset) & 1) == 1
}

// WriteBit return new byte with be written bit.
// If offset outside byt length(0 - 7), return old byte.
// Byte 0b10000000 write One by offset 2 return 0b10100000
func WriteBit(byt byte, bit Bit, offset int) byte {
	if offset < ByteMinBit || offset >= ByteMaxBit {
		return byt
	}

	if bit {
		return byt | 1<<(ByteMaxBit-1-offset)
	} else {
		return byt & (^(1 << (ByteMaxBit - 1 - offset)))
	}
}

// ReverseByte return the reverse bits in special byte.
// If input is 11001100, the return is 00110011
func ReverseByte(byt byte) byte {
	nb := byte(0)
	old := byt
	for i := ByteMinBit; i < ByteMaxBit; i++ {
		nb <<= 1
		nb |= old % 2
		old >>= 1
	}
	return nb
}

type BitStream struct {
	stream      []byte
	endOffset   int8 // the index of be written next bit
	startOffset int8 // the index of be read next bit
}

// ByteLength return the length of BitStream in byte
func (bs *BitStream) ByteLength() int {
	return len(bs.stream)
}

// BitLength return the length of BitStream in bit
func (bs *BitStream) BitLength() int {
	return (len(bs.stream)-1)*8 - int(bs.startOffset) + int(bs.endOffset)
}

// WriteBit write a bit to BitStream.
// WriteBit will append special bit to end of BitStream. And WriteBit is not thread-safe.
func (bs *BitStream) WriteBit(bit Bit) {
	if len(bs.stream) == 0 {
		bs.stream = append(bs.stream, 0)
		bs.endOffset = 0
	}
	if bs.endOffset == ByteMaxBit {
		bs.endOffset = 0
		bs.stream = append(bs.stream, 0)
	}

	currentByteIndex := len(bs.stream) - 1
	// because default bit is zero, so if write zero, can skip it
	if bit {
		// 01-23-45-67 (index)
		//-00-00-00-00, write pos = 0, write bit 1:
		// 00-00-00-00 | 10-00-00-00, 	(1 << (7-0=)7) => 10-00-00-00
		//-10-00-00-00, write pos = 3, write bit 1:
		// 10-00-00-00 | 1-00-00, 		(1 << (7-3=)4) => 10-01-00-00
		//-10-01-00-00, write pos = 7, write bit 1:
		// 10-01-00-00 | 1, 			(1 << (7-7=)0) => 10-01-00-01
		bs.stream[currentByteIndex] |= 1 << (ByteMaxBit - 1 - bs.endOffset)
	}

	bs.endOffset++
}

// ReadBit return a bit from head of BitStream.
// ReadBit see BitStream as a queue. The ReadBit is not thread-safe. When BitStream.BitLength() is zero, return
// Zero and io.EOF.
func (bs *BitStream) ReadBit() (Bit, error) {
	if bs.BitLength() == 0 {
		return Zero, io.EOF
	}
	if bs.startOffset == ByteMaxBit {
		bs.stream = bs.stream[1:]
		bs.startOffset = ByteMinBit
	}
	// 01234567 (index)
	// 01000000 1 : >> 6(=7-1) 01 		& 1 => 1
	// 11000000 1 : >> 6(=7-1) 11 		& 1 => 1
	// 11010100 5 : >> 2(=7-5) 110101 	& 1 => 1
	// 11010100 6 : >> 1(=7-6) 1101010  & 1 => 0
	d := bs.stream[0] >> (ByteMaxBit - 1 - bs.startOffset)
	bs.startOffset++
	return d&0b1 == 1, nil
}

// WriteByte always return nil, and write byt to end of BitStream.
// WriteByte is not thread-safe. WriteByte always see byt as 8 bits, it writes 8 bit to stream.
func (bs *BitStream) WriteByte(byt byte) error {
	if len(bs.stream) == 0 || bs.endOffset == ByteMaxBit {
		bs.stream = append(bs.stream, 0)
		bs.endOffset = 0
	}
	writeIndex := len(bs.stream) - 1

	bs.stream[writeIndex] |= byt >> (bs.endOffset)
	bs.stream = append(bs.stream, 0)
	writeIndex++
	bs.stream[writeIndex] = byt << (ByteMaxBit - bs.endOffset)
	return nil
}

// ReadByte return the byt from head of BitStream.
// ReadByte see BitStream as a queue, and is not thread-safe. If BitStream.BitLength < 8, it always returns 0 and
//	io.EOF. ReadByte will try to return 8 bits as a byte.
func (bs *BitStream) ReadByte() (byte, error) {
	if bs.ByteLength() == 0 {
		return 0, io.EOF
	}
	if bs.startOffset == ByteMaxBit {
		bs.stream = bs.stream[1:]
		if bs.ByteLength() == 0 {
			return 0, io.EOF
		}
		bs.startOffset = 0
	}

	if bs.startOffset == 0 {
		return bs.stream[0], nil
	}

	byt := bs.stream[0] << bs.startOffset
	bs.stream = bs.stream[1:]

	if len(bs.stream) == 0 {
		return 0, io.EOF
	}

	byt |= bs.stream[0] >> (ByteMaxBit - bs.startOffset)

	return byt, nil
}

func (bs *BitStream) PopBit() (Bit, error) {
	if bs.BitLength() == 0 {
		return Zero, io.EOF
	}
	if bs.endOffset == ByteMinBit {
		bs.stream = bs.stream[:len(bs.stream)-1]
		bs.endOffset = ByteMaxBit
	}

	bs.endOffset--
	return bs.stream[bs.endOffset]&1 == 1, nil
}
