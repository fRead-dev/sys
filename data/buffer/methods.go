package buffer

import (
	"bytes"
	"io"
	"sync"
)

//###########################################################//

func New() *BufferObj {
	obj := BufferObj{}

	var data bytes.Buffer
	var mu sync.Mutex

	obj.data = &data
	obj.mu = &mu

	return &obj
}

/////////////////////////////////////

func (obj *BufferObj) Close() {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	obj.data.Reset()
}

func (obj *BufferObj) Write(r io.Reader) (n int64, err error) {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	return obj.data.ReadFrom(r)
}

func (obj *BufferObj) Read(w io.Writer) (n int64, err error) {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	return obj.data.WriteTo(w)
}
