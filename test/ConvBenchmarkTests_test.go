package bit_test

import (
	"github.com/MXuDong/bit"
	"testing"
)

func Benchmark_BytesToBits_1length_test(b *testing.B) {
	bs := []byte{0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bit.BytesToBits(bs)
	}
}

func Benchmark_BytesToBits_100length_test(b *testing.B) {
	bs := []byte{}
	for i := 0; i < 100; i++ {
		bs = append(bs, 0)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bit.BytesToBits(bs)
	}
}

func Benchmark_BytesToBits_1000000length_test(b *testing.B) {
	bs := []byte{}
	for i := 0; i < 1000000; i++ {
		bs = append(bs, 0)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bit.BytesToBits(bs)
	}
}
