package buffer

import (
	"bytes"
	"io"
	"sync"
)

//###########################################################//

func New() *BufferObj {
	return &BufferObj{
		data: new(bytes.Buffer),
		mu:   new(sync.Mutex),
	}
}

func (obj *BufferObj) Close() {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	obj.data.Reset()
}

/////////////////////////////////////

func (obj *BufferObj) Write(input io.Reader) (n int64, err error) {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	return obj.data.ReadFrom(input)
}

func (obj *BufferObj) Read(output io.Writer) (n int64, err error) {
	obj.mu.Lock()
	defer obj.mu.Unlock()
	return obj.data.WriteTo(output)
}
