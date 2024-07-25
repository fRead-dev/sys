package buffer

import (
	"bytes"
	"errors"
	"io"
	"regexp"
	"sync"
)

//###########################################################//

func New() *BufferObj {
	return &BufferObj{
		bufData: new(bytes.Buffer),
		mu:      new(sync.Mutex),
	}
}

func (obj *BufferObj) Close() {
	obj.mu.Lock()
	defer obj.mu.Unlock()

	data := obj.bufData.Bytes()
	obj.uploadData = &data

	obj.bufData.Reset()
	obj.bufData = nil
}

/////////////////////////////////////

func (obj *BufferObj) Write(input io.Reader) (int64, error) {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	return obj.bufData.ReadFrom(input)
}

func (obj *BufferObj) Read() ([]byte, error) {
	if obj.bufData != nil {
		return nil, errors.New("write to buffer is not closed")
	}
	if obj.uploadData == nil {
		return nil, errors.New("buffer data is nil")
	}

	return *obj.uploadData, nil
}

////

func (obj *BufferObj) Bytes() uint64 {
	if obj.bufData != nil {
		return 0
	}
	if obj.uploadData == nil {
		return 0
	}

	return uint64(len(*obj.uploadData))
}

func (obj *BufferObj) Letters() uint64 {
	if obj.bufData != nil {
		return 0
	}
	if obj.uploadData == nil {
		return 0
	}

	buf := bytes.NewBuffer(*obj.uploadData)
	re := regexp.MustCompile(`^[a-zA-Z]$`)
	var count uint64

	for {
		char, _, err := buf.ReadRune()
		if err != nil {
			break
		}

		if re.MatchString(string(char)) {
			count++
		}
	}

	return count
}
