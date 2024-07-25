package buffer

import (
	"strings"
	"testing"
)

func TestBufferObj_WriteRead(t *testing.T) {
	obj := New()

	inputData := "This is test data"
	inputReader := strings.NewReader(inputData)

	n, err := obj.Write(inputReader)
	if err != nil {
		t.Fatalf("Error while recording: %v", err)
	}
	if n != int64(len(inputData)) {
		t.Fatalf("The number of bytes written (%d) is not as expected(%d)", n, len(inputData))
	}

	obj.Close()
	bb, err := obj.Read()
	if err != nil {
		t.Fatalf("Error while reading:%v", err)
	}
	if len(bb) != len(inputData) {
		t.Fatalf("The number of bytes read (%d) is not as expected (%d)", len(bb), len(inputData))
	}
	if string(bb) != inputData {
		t.Fatalf("Data read (%s) is not as expected (%s)", string(bb), inputData)
	}
}

func BenchmarkBufferObj_Write(b *testing.B) {
	inputData := "This is test data to test the performance of the recording function."
	for i := 0; i < b.N; i++ {
		obj := New()
		inputReader := strings.NewReader(inputData)
		_, err := obj.Write(inputReader)
		if err != nil {
			b.Fatalf("Error while recording: %v", err)
		}
	}
}

func BenchmarkBufferObj_Read(b *testing.B) {
	inputData := "This is test data to test the performance of the recording function."
	obj := New()
	inputReader := strings.NewReader(inputData)
	_, err := obj.Write(inputReader)
	obj.Close()
	if err != nil {
		b.Fatalf("Error while recording: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_, err = obj.Read()
		if err != nil {
			b.Fatalf("Error while reading: %v", err)
		}
	}
}
