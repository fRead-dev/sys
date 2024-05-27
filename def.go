package ParserInterface

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/andybalholm/brotli"
	"golang.org/x/crypto/blake2b"
	"io"
)

func defError(methodNAme string) error {
	return fmt.Errorf("Method `%s` not specified", methodNAme)
}

//////////////////////////////////////////////////////////////////////////////////////

type StructDef struct{}

// Hash retrieves the checksum. Fixed length of 16 characters.
func (obj *StructDef) Hash(data *[]byte) (hash string) {
	hasher, err := blake2b.New(16, nil)
	if err != nil {
		return
	}

	hasher.Write(*data)
	sum := hasher.Sum(nil)

	return hex.EncodeToString(sum)
}

// Compression packs the received data for buffering. Brotli is used at maximum compression.
func (obj *StructDef) Сompression(data *[]byte) (compressData []byte, err error) {
	var buf bytes.Buffer
	writer := brotli.NewWriterLevel(&buf, brotli.BestCompression)

	_, err = writer.Write(*data)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decompression unpacks the compressed data. Brotli is used.
func (obj *StructDef) Decompression(data *[]byte) (decompressData []byte, err error) {
	var buf bytes.Buffer
	reader := brotli.NewReader(bytes.NewReader(*data))

	_, err = io.Copy(&buf, reader)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

//

// ThisDomain checks if the URL matches the given domain.
func (obj *StructDef) ThisDomain(url string) (isValid bool) { return }

//

// GetDomain retrieves the domain variable.
func (obj *StructDef) GetDomain() (domain string) { return }

// GetVersion retrieves the module version.
func (obj *StructDef) GetVersion() (version string) { return "0.0.0" }

// GetRegExp retrieves the list of regular expressions used to check domain compliance.
func (obj *StructDef) GetRegExp() (regExpArr []string) { return }

//

// GetWork retrieves the work ID from the URL.
func (obj *StructDef) GetWork(url string) (idWork string, err error) {
	err = defError("GetWork")
	return
}

// GetAuthor retrieves the author ID from the URL.
func (obj *StructDef) GetAuthor(url string) (idWork string, idAutor string, err error) {
	err = defError("GeAuthor")
	return
}

// GetChapter retrieves the chapter ID from the URL.
func (obj *StructDef) GetChapter(url string) (idWork string, idChapter string, err error) {
	err = defError("GetChapter")
	return
}

//

// UrlWork generates the URL to the work based
func (obj *StructDef) UrlWork(idWork string) (url string, err error) {
	err = defError("UrlWork")
	return
}

// UrlAuthor generates the URL to the author based on their ID.
func (obj *StructDef) UrlAuthor(idWork string, idAutor string) (url string, err error) {
	err = defError("UrlAuthor")
	return
}

// UrlChapter generates the URL to the chapter based on its ID.
func (obj *StructDef) UrlChapter(idWork string, idChapter string) (url string, err error) {
	err = defError("UrlChapter")
	return
}

/*##*/

// Ping checks the availability of the domain.
func (obj *StructDef) Ping() (err error) {
	err = defError("Ping")
	return
}

// PingWork checks the availability of the work.
func (obj *StructDef) PingWork(idWork string) (err error) {
	err = defError("PingWork")
	return
}

// PingChapter checks the availability of the chapter.
func (obj *StructDef) PingChapter(idWork string, idChapter string) (err error) {
	err = defError("PingChapter")
	return
}

//

// LoadWork loads the entire work by its ID.
func (obj *StructDef) LoadWork(idWork string) (globalObj GlobalObj, err error) {
	err = defError("LoadWork")
	return
}

// LoadWorkFromFile loads the entire work from a file.
func (obj *StructDef) LoadWorkFromFile(fileType string, htmlText io.Reader) (globalObj GlobalObj, err error) {
	err = defError("LoadWorkFromFile")
	return
}

// LoadInfo loads information about the work by its ID.
func (obj *StructDef) LoadInfo(idWork string) (infoObj InfoObj, err error) {
	err = defError("LoadInfo")
	return
}

// LoadChapter loads the chapter by its ID.
func (obj *StructDef) LoadChapter(idWork string, idChapter string) (pageObj PageObj, err error) {
	err = defError("LoadChapter")
	return
}
