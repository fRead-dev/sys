package crypt

import (
	"bytes"
	"testing"
)

func TestAES(t *testing.T) {
	originalData := []byte("This is some test data.")
	key := "testkey1234567890"

	var encBuf, decBuf bytes.Buffer

	err := Encode(bytes.NewReader(originalData), &encBuf, key)
	if err != nil {
		t.Fatalf("Encode() error = %v", err)
	}

	if bytes.Equal(originalData, encBuf.Bytes()) {
		t.Errorf("Encode() produced data identical to original data")
	}

	err = Decode(&encBuf, &decBuf, key)
	if err != nil {
		t.Fatalf("Decode() error = %v", err)
	}

	if !bytes.Equal(originalData, decBuf.Bytes()) {
		t.Errorf("Decode() = %v, want %v", decBuf.Bytes(), originalData)
	}
}

func BenchmarkEncode(b *testing.B) {
	data := []byte("Benchmarking the encoding function with some data.")
	key := "benchmarkkey1234"
	var encBuf bytes.Buffer

	for i := 0; i < b.N; i++ {
		encBuf.Reset()
		err := Encode(bytes.NewReader(data), &encBuf, key)
		if err != nil {
			b.Fatalf("Encode() error = %v", err)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	data := []byte("Benchmarking the decoding function with some data.")
	key := "benchmarkkey1234"
	var encBuf, decBuf bytes.Buffer

	err := Encode(bytes.NewReader(data), &encBuf, key)
	if err != nil {
		b.Fatalf("Encode() error = %v", err)
	}

	for i := 0; i < b.N; i++ {
		decBuf.Reset()
		err := Decode(&encBuf, &decBuf, key)
		if err != nil {
			b.Fatalf("Decode() error = %v", err)
		}
	}
}
