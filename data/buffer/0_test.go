package buffer

import (
	"bytes"
	"testing"
)

func TestGlobal(t *testing.T) {
	data := New()
	text := "hello world"

	reader := bytes.NewBufferString(text)
	_, err := data.Write(reader)
	if err != nil {
		t.Error(err)
	}

	var buffer bytes.Buffer
	_, err = data.Read(&buffer)
	if err != nil {
		t.Error(err)
	}

	if buffer.String() != text {
		t.Log("{"+buffer.String()+"}", "{"+text+"}")
		t.Fatal("buffer is not equal to text")
	}
}
