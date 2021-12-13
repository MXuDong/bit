package bit

import "testing"

func TestReadBit(t *testing.T) {
	type args struct {
		byt    byte
		offset int
	}
	tests := []struct {
		name string
		args args
		want Bit
	}{
		{name: "test 10000000 - 0 == 1", args: args{byt: 0b10000000, offset: 0}, want: One},
		{name: "test 10000000 - 1 == 0", args: args{byt: 0b10000000, offset: 1}, want: Zero},
		{name: "test 00000001 - 7 == 1", args: args{byt: 0b00000001, offset: 7}, want: One},
		{name: "test 10100000 - 1 == 0", args: args{byt: 0b10100000, offset: 1}, want: Zero},
		{name: "test 10100000 - 2 == 1", args: args{byt: 0b10100000, offset: 2}, want: One},
		{name: "test 10100000 - 3 == 0", args: args{byt: 0b10100000, offset: 3}, want: Zero},
		{name: "test greater than 7", args: args{byt: 0, offset: 8}, want: Zero},
		{name: "test less than 0", args: args{byt: 0, offset: -1}, want: Zero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadBit(tt.args.byt, tt.args.offset); got != tt.want {
				t.Errorf("ReadBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteBit(t *testing.T) {
	type args struct {
		byt    byte
		offset int
		bit    Bit
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "Test 00000000 write 7 1 = 00000001", args: args{byt: 0b00000000, bit: One, offset: 7}, want: 0b00000001},
		{name: "Test 00000000 write 0 1 = 10000000", args: args{byt: 0b00000000, bit: One, offset: 0}, want: 0b10000000},
		{name: "Test 10000000 write 0 1 = 10000000", args: args{byt: 0b10000000, bit: One, offset: 0}, want: 0b10000000},
		{name: "Test 10000000 write 0 0 = 00000000", args: args{byt: 0b10000000, bit: Zero, offset: 0}, want: 0b00000000},
		{name: "Test 00000001 write 7 0 = 00000000", args: args{byt: 0b00000001, bit: Zero, offset: 7}, want: 0b00000000},
		{name: "Test 10100000 write 1 1 = 11100000", args: args{byt: 0b10100000, bit: One, offset: 1}, want: 0b11100000},
		{name: "Test 10100000 write 2 0 = 10000000", args: args{byt: 0b10100000, bit: Zero, offset: 2}, want: 0b10000000},
		{name: "Test less than 0", args: args{byt: 0, bit: One, offset: -1}, want: 0},
		{name: "Test greater than 7", args: args{byt: 0, bit: One, offset: 8}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WriteBit(tt.args.byt, tt.args.bit, tt.args.offset); got != tt.want {
				t.Errorf("WriteBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitStream_WriteByte(t *testing.T) {
	type fields struct {
		stream      []byte
		writeOffset int8
		readOffset  int8
	}
	type args struct {
		byt byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "Test 00000000 write 10000001 = 10000001", args: args{byt: 0b10000001}, fields: fields{stream: []byte{0}, writeOffset: 0, readOffset: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BitStream{
				stream:      tt.fields.stream,
				endOffset:   tt.fields.writeOffset,
				startOffset: tt.fields.readOffset,
			}
			_ = bs.WriteByte(tt.args.byt)
			byt, _ := bs.ReadByte()

			if byt != tt.args.byt {
				t.Errorf("ReadByte() got = %v, want %v", byt, tt.args.byt)
			}
		})
	}
}

func TestBitStream_ReadByte(t *testing.T) {
	type fields struct {
		stream      []byte
		writeOffset int8
		readOffset  int8
	}
	tests := []struct {
		name    string
		fields  fields
		want    byte
		wantErr bool
	}{
		{name: "Test common error", fields: fields{stream: []byte{}, writeOffset: 0, readOffset: 0}, want: 0, wantErr: true},
		{name: "Test read in one byte 0", fields: fields{stream: []byte{0}, writeOffset: 0, readOffset: 0}, want: 0, wantErr: false},
		{name: "Test read in one byte 1", fields: fields{stream: []byte{0b10001010}, writeOffset: 0, readOffset: 0}, want: 0b10001010, wantErr: false},
		{name: "Test read in two byte 2 with error", fields: fields{stream: []byte{0b10001010}, writeOffset: 0, readOffset: 1}, want: 0, wantErr: true},
		{name: "Test read in two byte 2 with error", fields: fields{stream: []byte{0b10001010, 0b10000001}, writeOffset: 0, readOffset: 1}, want: 0b00010101, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BitStream{
				stream:      tt.fields.stream,
				endOffset:   tt.fields.writeOffset,
				startOffset: tt.fields.readOffset,
			}
			got, err := bs.ReadByte()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadByte() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseByte(t *testing.T) {
	type args struct {
		byt byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "common", args: args{byt: 0b11001100}, want: 0b00110011},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseByte(tt.args.byt); got != tt.want {
				t.Errorf("ReverseByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
