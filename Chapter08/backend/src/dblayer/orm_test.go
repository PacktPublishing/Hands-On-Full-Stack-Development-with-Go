package dblayer

import "testing"

func BenchmarkHashPassword(b *testing.B) {
	text := "A String to be Hashed"
	for i := 0; i < b.N; i++ {
		hashPassword(&text)
	}
}
