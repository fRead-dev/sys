package chapter

import (
	fReadSYS "github.com/fRead-dev/sys"
	fReadBuffer "github.com/fRead-dev/sys/data/buffer"
)

//###########################################################//

type ChapterObj struct {
	Time  fReadSYS.TimesObj
	Size  fReadSYS.SizesObj
	Check fReadSYS.IntegrityChecksObj

	Data *fReadBuffer.BufferObj
}

type ChapterChunksObj struct {
	Time  fReadSYS.TimesObj
	Size  fReadSYS.SizesObj
	Check fReadSYS.IntegrityChecksObj

	Chunks []*fReadBuffer.BufferObj
}
