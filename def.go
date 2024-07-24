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

/*####*/

// Hash retrieves the checksum. Fixed length of 16 characters.
func Hash(data *[]byte) (hash string) {
	hasher, err := blake2b.New(16, nil)
	if err != nil {
		return
	}

	hasher.Write(*data)
	sum := hasher.Sum(nil)

	return hex.EncodeToString(sum)
}

// Compression packs the received data for buffering. Brotli is used at maximum compression.
func Сompression(data *[]byte) (compressData []byte, err error) {
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
func Decompression(data *[]byte) (decompressData []byte, err error) {
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
func ThisDomain(url string) (isValid bool) { return }

//

// GetDomain retrieves the domain variable.
func GetDomain() (domain string) { return }

// GetVersion retrieves the module version.
func GetVersion() (version string) { return "0.0.0" }

// GetRegExp retrieves the list of regular expressions used to check domain compliance.
func GetRegExp() (regExpArr []string) { return }

//

// GetWork retrieves the work ID from the URL.
func GetWork(url string) (idWork string, err error) {
	err = defError("GetWork")
	return
}

// GetAuthor retrieves the author ID from the URL.
func GetAuthor(url string) (idWork string, idAutor string, err error) {
	err = defError("GeAuthor")
	return
}

// GetChapter retrieves the chapter ID from the URL.
func GetChapter(url string) (idWork string, idChapter string, err error) {
	err = defError("GetChapter")
	return
}

//

// UrlWork generates the URL to the work based
func UrlWork(idWork string) (url string, err error) {
	err = defError("UrlWork")
	return
}

// UrlAuthor generates the URL to the author based on their ID.
func UrlAuthor(idWork string, idAutor string) (url string, err error) {
	err = defError("UrlAuthor")
	return
}

// UrlChapter generates the URL to the chapter based on its ID.
func UrlChapter(idWork string, idChapter string) (url string, err error) {
	err = defError("UrlChapter")
	return
}

/*##*/

// Ping checks the availability of the domain.
func Ping() (err error) {
	err = defError("Ping")
	return
}

// PingWork checks the availability of the work.
func PingWork(idWork string) (err error) {
	err = defError("PingWork")
	return
}

// PingChapter checks the availability of the chapter.
func PingChapter(idWork string, idChapter string) (err error) {
	err = defError("PingChapter")
	return
}

//

// LoadWork loads the entire work by its ID.
func LoadWork(idWork string) (globalObj GlobalObj, err error) {
	err = defError("LoadWork")
	return
}

// LoadWorkFromFile loads the entire work from a file.
func LoadWorkFromFile(fileType string, htmlText io.Reader) (globalObj GlobalObj, err error) {
	err = defError("LoadWorkFromFile")
	return
}

// LoadInfo loads information about the work by its ID.
func LoadInfo(idWork string) (WorkObj WorkObj, err error) {
	err = defError("LoadInfo")
	return
}

// LoadChapter loads the chapter by its ID.
func LoadChapter(idWork string, idChapter string) (ChapterObj ChapterObj, err error) {
	err = defError("LoadChapter")
	return
}
