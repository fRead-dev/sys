package ParserInterface

import (
	"bytes"
	"encoding/hex"
	"github.com/andybalholm/brotli"
	"golang.org/x/crypto/blake2b"
	"io"
)

//////////////////////////////////////////////////////////////////////////////////////

func (sys *MethodsSysObject) Hash(src io.Reader) HashType {
	var reader bytes.Buffer

	_, err := reader.ReadFrom(src)
	if err != nil {
		return ""
	}

	hasher, err := blake2b.New(32, nil)
	if err != nil {
		return ""
	}

	hasher.Write(reader.Bytes())
	sum := hasher.Sum(nil)
	hash := hex.EncodeToString(sum)

	return HashType(hash)
}

func (sys *MethodsSysObject) Compress(src io.Reader) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	var reader bytes.Buffer

	_, err := reader.ReadFrom(src)
	if err != nil {
		return nil, err
	}

	writer := brotli.NewWriterLevel(&buf, brotli.BestCompression)

	_, err = writer.Write(reader.Bytes())
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func (sys *MethodsSysObject) Decompress(src io.Reader) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	bro := brotli.NewReader(src)

	_, err := io.Copy(&buf, bro)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
