package sys

import (
	"bytes"
	"testing"
)

func TestCompress(t *testing.T) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data"),
	}

	compressedData, err := data.Compress()
	if err != nil {
		t.Fatalf("Failed to compress data: %v", err)
	}

	if !compressedData.IsCompress {
		t.Errorf("Expected IsCompress to be true, got false")
	}

	if bytes.Equal(data.Data, compressedData.Data) {
		t.Errorf("Expected compressed data to be different from original data")
	}
}

func TestDecompress(t *testing.T) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data"),
	}
	compressedData, err := data.Compress()
	if err != nil {
		t.Fatalf("Failed to compress data: %v", err)
	}

	decompressedData, err := compressedData.Decompress()
	if err != nil {
		t.Fatalf("Failed to decompress data: %v", err)
	}

	if decompressedData.IsCompress {
		t.Errorf("Expected IsCompress to be false, got true")
	}

	if !bytes.Equal(decompressedData.Data, data.Data) {
		t.Errorf("Decompressed data does not match the expected value")
	}
}

func TestChunks(t *testing.T) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("this is a test of chunks"),
	}

	chunks, err := data.Chunks()
	if err != nil {
		t.Fatalf("Failed to get chunks: %v", err)
	}

	if len(chunks) == 0 {
		t.Errorf("Expected non-zero number of chunks")
	}

	// Validate each chunk
	for i, chunk := range chunks {
		if chunk.IsCompress != data.IsCompress {
			t.Errorf("Chunk %d has incorrect IsCompress value", i)
		}
		if len(chunk.Data) == 0 {
			t.Errorf("Chunk %d has empty data", i)
		}
	}
}

func BenchmarkCompress(b *testing.B) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data for compression"),
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := data.Compress()
		if err != nil {
			b.Fatalf("Failed to compress data: %v", err)
		}
	}
	b.StopTimer()
}

func BenchmarkDecompress(b *testing.B) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data"),
	}
	compressedData, err := data.Compress()
	if err != nil {
		b.Fatalf("Failed to compress data: %v", err)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, err := compressedData.Decompress()
		if err != nil {
			b.Fatalf("Failed to decompress data: %v", err)
		}
	}
	b.StopTimer()
}

func BenchmarkChunks(b *testing.B) {
	data := &DataObj{
		IsCompress: false,
		Data:       []byte("test data to split into chunks"),
	}

	for i := 0; i < b.N; i++ {
		_, err := data.Chunks()
		if err != nil {
			b.Fatalf("Failed to get chunks: %v", err)
		}
	}
}
