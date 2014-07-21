package main

import (
	"crypto/sha1"
	"testing"
)

var trueBytes = []byte{'(', 't', 'r', 'u', 'e', ')'}

func BenchmarkByteCasts(b *testing.B) {
	checksum := sha1.New()
	for i := 0; i < b.N; i++ {
		checksum.Write([]byte("(true)"))
	}
}

func BenchmarkByteConstants(b *testing.B) {
	checksum := sha1.New()
	for i := 0; i < b.N; i++ {
		checksum.Write(trueBytes)
	}
}
