package chapter

import fReadSYS "github.com/fRead-dev/sys"

//###########################################################//

func (chapter *ChapterObj) Chunks() (*ChapterChunksObj, error) {
	obj := &ChapterChunksObj{}

	obj.Time = chapter.Time
	obj.Size = chapter.Size
	obj.Check = chapter.Check

	chunks, err := chapter.Chapter.Chunks()
	if err != nil {
		return nil, err
	}

	obj.Chunks = chunks
	return obj, nil
}

func (chunks *ChapterChunksObj) Chapter() (*ChapterObj, error) {
	obj := &ChapterObj{}

	obj.Time = chunks.Time
	obj.Size = chunks.Size
	obj.Check = chunks.Check

	chapter, err := fReadSYS.ChunksToData(chunks.Chunks)
	if err != nil {
		return nil, err
	}

	obj.Chapter = *chapter
	return obj, nil
}
