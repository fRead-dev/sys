package chapter

import (
	fReadSYS "github.com/fRead-dev/sys"
)

//###########################################################//

type ChapterObj struct {
	Time  fReadSYS.TimesObj
	Size  fReadSYS.SizesObj
	Check fReadSYS.IntegrityChecksObj

	Chapter fReadSYS.DataObj
}

type ChapterChunksObj struct {
	Time  fReadSYS.TimesObj
	Size  fReadSYS.SizesObj
	Check fReadSYS.IntegrityChecksObj

	Chunks []fReadSYS.DataObj
}
