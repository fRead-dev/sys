package sys

import (
	"bytes"
	fReadCompress "github.com/fRead-dev/sys/data/compress"
)

//###########################################################//

func (data *DataObj) Compress() (*DataObj, error) {
	if data.IsCompress {
		return data, nil
	}

	inputReader := bytes.NewReader(data.Data)
	compressedBuffer := new(bytes.Buffer)

	err := fReadCompress.Encode(inputReader, compressedBuffer)
	if err != nil {
		return data, err
	}

	return &DataObj{
		IsCompress: true,
		Data:       compressedBuffer.Bytes(),
	}, nil
}

func (data *DataObj) Decompress() (*DataObj, error) {
	if !data.IsCompress {
		return data, nil
	}

	decompressedBuffer := new(bytes.Buffer)
	compressedReader := bytes.NewReader(data.Data)
	err := fReadCompress.Decode(compressedReader, decompressedBuffer)

	if err != nil {
		return data, err
	}

	return &DataObj{
		IsCompress: false,
		Data:       decompressedBuffer.Bytes(),
	}, nil
}

//////////////

func (data *DataObj) Chunks() ([]DataObj, error) {
	isCompress := false

	if data.IsCompress {
		bufData, err := data.Decompress()
		data = bufData
		isCompress = true

		if err != nil {
			return nil, err
		}
	}

	inputReader := bytes.NewReader(data.Data)
	chunkSize, numChunks := ChunkSize(inputReader.Len())
	arr := make([]DataObj, numChunks)

	pos := 0
	chunk := make([]byte, chunkSize)
	for {
		_, err := inputReader.Read(chunk)
		if err != nil {
			break
		}

		arr[pos].IsCompress = isCompress
		arr[pos].Data = chunk

		if isCompress {
			buf, _ := arr[pos].Compress()
			arr[pos] = *buf
		}
	}

	return arr, nil
}
