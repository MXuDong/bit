package bit

import (
	"reflect"
	"testing"
)

func TestBoolToBit(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want Bit
	}{
		{name: "Bool true to One", args: args{v: true}, want: One},
		{name: "Bool true to Zero", args: args{v: false}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolToBit(tt.args.v); got != tt.want {
				t.Errorf("BoolToBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteToBits(t *testing.T) {
	type args struct {
		byt byte
	}
	tests := []struct {
		name string
		args args
		want []Bit
	}{
		{name: "Common byte to bits", args: args{byt: 0b10101010}, want: []Bit{One, Zero, One, Zero, One, Zero, One, Zero}},
		{name: "Common byte to bits(all zero)", args: args{byt: 0b00000000}, want: []Bit{Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero}},
		{name: "Common byte to bits(all one)", args: args{byt: 0b11111111}, want: []Bit{One, One, One, One, One, One, One, One}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteToBits(tt.args.byt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteToBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToBits(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want []Bit
	}{
		{
			name: "Common []byte to []Bit",
			args: args{
				bytes: []byte{
					0b10101010,
					0b00000000,
					0b11111111,
				},
			},
			want: []Bit{
				One, Zero, One, Zero, One, Zero, One, Zero,
				Zero, Zero, Zero, Zero, Zero, Zero, Zero, Zero,
				One, One, One, One, One, One, One, One,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToBits(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
