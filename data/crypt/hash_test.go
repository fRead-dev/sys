package crypt

import (
	"bytes"
	"testing"
)

func TestHash32(t *testing.T) {
	input := bytes.NewReader([]byte("test input"))
	expected := "14f9a1375564fd6e1753125374374be21fb72319d989381a443e5a11de6687e4"

	result, err := Hash32(input)
	if err != nil {
		t.Fatalf("Hash32() error = %v", err)
	}
	if result != expected {
		t.Errorf("Hash32() = %v, want %v", result, expected)
	}
}

func TestHash48(t *testing.T) {
	input := bytes.NewReader([]byte("test input"))
	expected := "9a6157d785c72d1de26309ddc3a6dc998bdf71c2b82f5682ac50db7067f2606e4dae31591276838253ce87ed2ae1de26"

	result, err := Hash48(input)
	if err != nil {
		t.Fatalf("Hash48() error = %v", err)
	}
	if result != expected {
		t.Errorf("Hash48() = %v, want %v", result, expected)
	}
}

func TestHash64(t *testing.T) {
	input := bytes.NewReader([]byte("test input"))
	expected := "399522f300cea9b23a73adb37acba28fa528789f3931c38a55fca68e12e5e6a0de8b94cebbdbc62713cd5a8e7f35eb219b02574455d35c6ca4b8b2888acc0a51"

	result, err := Hash64(input)
	if err != nil {
		t.Fatalf("Hash64() error = %v", err)
	}
	if result != expected {
		t.Errorf("Hash64() = %v, want %v", result, expected)
	}
}

func TestHash32Key(t *testing.T) {
	input := bytes.NewReader([]byte("test input"))
	key := "testkey"
	expected := "f95da50f27138f43e7490aae301a89868f14e7d3728212fff015f91221459f27"

	result, err := Hash32Key(input, key)
	if err != nil {
		t.Fatalf("Hash32Key() error = %v", err)
	}
	if result != expected {
		t.Errorf("Hash32Key() = %v, want %v", result, expected)
	}
}

func TestHash48Key(t *testing.T) {
	input := bytes.NewReader([]byte("test input"))
	key := "testkey"
	expected := "7844c9e4ae68f2668a92ec907b3c234a4a0d9eaa2113060632580e9de0ad6cee39b543f4bf3412b4e7303649e60038cb"

	result, err := Hash48Key(input, key)
	if err != nil {
		t.Fatalf("Hash48Key() error = %v", err)
	}
	if result != expected {
		t.Errorf("Hash48Key() = %v, want %v", result, expected)
	}
}

func TestHash64Key(t *testing.T) {
	input := bytes.NewReader([]byte("test input"))
	key := "testkey"
	expected := "6e4996915e79c740ee9d5ffa5aa31c96edc9856072fa31dff57d340c8703e0f79d27150d8c9e598528d96aaf3ba6208c02a3846ee8485dccd7cac713f53e77d1"

	result, err := Hash64Key(input, key)
	if err != nil {
		t.Fatalf("Hash64Key() error = %v", err)
	}
	if result != expected {
		t.Errorf("Hash64Key() = %v, want %v", result, expected)
	}
}

func BenchmarkHash32(b *testing.B) {
	input := bytes.NewReader([]byte("test input"))
	for i := 0; i < b.N; i++ {
		_, _ = Hash32(input)
	}
}

func BenchmarkHash48(b *testing.B) {
	input := bytes.NewReader([]byte("test input"))
	for i := 0; i < b.N; i++ {
		_, _ = Hash48(input)
	}
}

func BenchmarkHash64(b *testing.B) {
	input := bytes.NewReader([]byte("test input"))
	for i := 0; i < b.N; i++ {
		_, _ = Hash64(input)
	}
}

func BenchmarkHash32Key(b *testing.B) {
	input := bytes.NewReader([]byte("test input"))
	key := "testkey"
	for i := 0; i < b.N; i++ {
		_, _ = Hash32Key(input, key)
	}
}

func BenchmarkHash48Key(b *testing.B) {
	input := bytes.NewReader([]byte("test input"))
	key := "testkey"
	for i := 0; i < b.N; i++ {
		_, _ = Hash48Key(input, key)
	}
}

func BenchmarkHash64Key(b *testing.B) {
	input := bytes.NewReader([]byte("test input"))
	key := "testkey"
	for i := 0; i < b.N; i++ {
		_, _ = Hash64Key(input, key)
	}
}
