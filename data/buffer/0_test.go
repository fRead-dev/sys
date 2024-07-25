package buffer

import (
	"bytes"
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

	outputBuffer := &bytes.Buffer{}
	n, err = obj.Read(outputBuffer)
	if err != nil {
		t.Fatalf("Error while reading:%v", err)
	}
	if n != int64(len(inputData)) {
		t.Fatalf("The number of bytes read (%d) is not as expected (%d)", n, len(inputData))
	}
	if outputBuffer.String() != inputData {
		t.Fatalf("Data read (%s) is not as expected (%s)", outputBuffer.String(), inputData)
	}
}

func TestBufferObj_Close(t *testing.T) {
	obj := New()

	inputData := "This is test data"
	inputReader := strings.NewReader(inputData)

	_, err := obj.Write(inputReader)
	if err != nil {
		t.Fatalf("Error while recording: %v", err)
	}

	obj.Close()

	if obj.data.Len() != 0 {
		t.Fatalf("The buffer was not cleared after closing")
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
	if err != nil {
		b.Fatalf("Error while recording: %v", err)
	}

	for i := 0; i < b.N; i++ {
		outputBuffer := &bytes.Buffer{}
		_, err := obj.Read(outputBuffer)
		if err != nil {
			b.Fatalf("Error while reading: %v", err)
		}
	}
}
