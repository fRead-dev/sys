package chapter

import (
	fReadBuffer "github.com/fRead-dev/sys/data/buffer"
	"io"

	fReadText "github.com/fRead-dev/sys/data/text"
)

//###########################################################//

func DataToChapter(input io.Reader) *ChapterObj {
	obj := ChapterObj{}

	obj.Size.Letters = fReadText.CountLetters(input)
	obj.Size.Bytes = fReadText.CountBytes(input)

	obj.Data = fReadBuffer.New()
	obj.Data.Write(input)

	return &obj
}

func DataToChunks(input io.Reader) (*ChapterChunksObj, error) {
	chapter, err := DataToChapter(input)
	if err != nil {
		return nil, err
	}

	return chapter.Chunks(), nil
}

/////////////////////

func (chapter *ChapterObj) Chunks() *ChapterChunksObj {
	return nil
}

func (chunks *ChapterChunksObj) Chapter() *ChapterObj {
	return nil
}
