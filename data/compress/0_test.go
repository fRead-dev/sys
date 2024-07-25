package compress

import (
	"bytes"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	inputData := []byte("This is test data to test the compression and decompression functions.")
	inputReader := bytes.NewReader(inputData)
	compressedBuffer := new(bytes.Buffer)

	if err := Encode(inputReader, compressedBuffer); err != nil {
		t.Fatalf("Error during compression:%v", err)
	}

	decompressedBuffer := new(bytes.Buffer)
	compressedReader := bytes.NewReader(compressedBuffer.Bytes())
	if err := Decode(compressedReader, decompressedBuffer); err != nil {
		t.Fatalf("Error while unpacking:%v", err)
	}

	if !bytes.Equal(decompressedBuffer.Bytes(), inputData) {
		t.Fatalf("The extracted data does not correspond to the original data")
	}
}

func BenchmarkEncode(b *testing.B) {
	inputData := []byte("This is test data to test the performance of the compression function.")
	for i := 0; i < b.N; i++ {
		inputReader := bytes.NewReader(inputData)
		compressedBuffer := new(bytes.Buffer)

		if err := Encode(inputReader, compressedBuffer); err != nil {
			b.Fatalf("Error during compression:%v", err)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	inputData := []byte("This is test data to test the performance of the compression function.")
	inputReader := bytes.NewReader(inputData)
	compressedBuffer := new(bytes.Buffer)

	if err := Encode(inputReader, compressedBuffer); err != nil {
		b.Fatalf("Error during compression: %v", err)
	}

	compressedData := compressedBuffer.Bytes()

	for i := 0; i < b.N; i++ {
		compressedReader := bytes.NewReader(compressedData)
		decompressedBuffer := new(bytes.Buffer)

		if err := Decode(compressedReader, decompressedBuffer); err != nil {
			b.Fatalf("Error while unpacking: %v", err)
		}
	}
}
