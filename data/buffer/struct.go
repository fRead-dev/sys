package buffer

import (
	"bytes"
	"sync"
)

type BufferObj struct {
	data *bytes.Buffer
	mu   *sync.Mutex
}
