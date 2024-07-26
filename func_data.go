package sys

import (
	"bytes"
	fReadCompress "github.com/fRead-dev/sys/data/compress"
	fReadCrypt "github.com/fRead-dev/sys/data/crypt"
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
		IsCrypt:    data.IsCrypt,
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
		IsCrypt:    data.IsCrypt,
		Data:       decompressedBuffer.Bytes(),
	}, nil
}

////

func (data *DataObj) Encrypt(key string) (*DataObj, error) {
	if data.IsCrypt {
		return data, nil
	}

	var bufData DataObj

	if data.IsCompress {
		buf, err := data.Decompress()
		if err != nil {
			return nil, err
		}
		bufData = *buf
	} else {
		bufData = *data
	}

	deBuffer := new(bytes.Buffer)
	coReader := bytes.NewReader(bufData.Data)
	err := fReadCrypt.Encode(coReader, deBuffer, key)
	if err != nil {
		return nil, err
	}

	encodeData := DataObj{
		IsCrypt: true,
		Data:    deBuffer.Bytes(),
	}

	if data.IsCompress {
		buf, err := encodeData.Compress()
		return buf, err
	}

	return &encodeData, nil
}
func (data *DataObj) Decrypt(key string) (*DataObj, error) {
	if !data.IsCrypt {
		return data, nil
	}

	var bufData DataObj

	if data.IsCompress {
		buf, err := data.Decompress()
		if err != nil {
			return nil, err
		}
		bufData = *buf
	} else {
		bufData = *data
	}

	deBuffer := new(bytes.Buffer)
	coReader := bytes.NewReader(bufData.Data)
	err := fReadCrypt.Decode(coReader, deBuffer, key)
	if err != nil {
		return nil, err
	}

	decodeData := DataObj{
		IsCrypt: false,
		Data:    deBuffer.Bytes(),
	}

	if data.IsCompress {
		buf, err := decodeData.Compress()
		return buf, err
	}

	return &decodeData, nil
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

func ChunksToData(chunks []DataObj) (*DataObj, error) {
	var dataBuf bytes.Buffer
	isCompress := false

	for _, chunk := range chunks {
		if chunk.IsCompress {
			isCompress = true
			data, err := chunk.Decompress()
			if err != nil {
				return nil, err
			}
			dataBuf.Write(data.Data)
		} else {
			dataBuf.Write(chunk.Data)
		}
	}

	obj := &DataObj{Data: dataBuf.Bytes()}
	if isCompress {
		data, err := obj.Decompress()
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return obj, nil
}
