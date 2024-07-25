package buffer

import (
	"bytes"
	"sync"
)

type BufferObj struct {
	mu *sync.Mutex

	bufData    *bytes.Buffer
	uploadData *[]byte
}
