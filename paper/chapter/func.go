package chapter

import (
	"bytes"
	fReadCompress "github.com/fRead-dev/sys/data/compress"
)

//###########################################################//

func (chapter *ChapterObj) Chunks() (*ChapterChunksObj, error) {
	obj := &ChapterChunksObj{}
	data := new(bytes.Buffer)

	obj.Time = chapter.Time
	obj.Size = chapter.Size
	obj.Check = chapter.Check

	if chapter.Chapter.IsCompress {
		compressedReader := bytes.NewReader(chapter.Chapter.Data)
		err := fReadCompress.Decode(compressedReader, data)
		if err != nil {
			return nil, err
		}
	} else {
		data.Write(chapter.Chapter.Data)
	}

	chunks, err := chapter.Chapter.Chunks()
	if err != nil {
		return nil, err
	}

	obj.Chunks = chunks
	return obj, nil
}

func (chunks *ChapterChunksObj) Chapter() *ChapterObj {
	obj := &ChapterObj{}

	return obj
}
