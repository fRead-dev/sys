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

// Hash Получение контрольной суммы.
func (obj *StructDef) Hash(data *[]byte) (hash string) {
	hasher, err := blake2b.New(16, nil)
	if err != nil {
		return
	}

	hasher.Write(*data)
	sum := hasher.Sum(nil)

	return hex.EncodeToString(sum)
}

// Сompression	Упаковка полученных данных для буферизации
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

// Decompression Распаковка сжатых данных
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

// ThisDomain Проверка соответствует ли URL данному домену
func (obj *StructDef) ThisDomain(url string) (isValid bool) { return }

//

// GetDomain Получение переменной домена
func (obj *StructDef) GetDomain() (domain string) { return }

// GetVersion Получение версии модуля
func (obj *StructDef) GetVersion() (version string) { return "0.0.0" }

// GetRegExp Получение списка регулярных выражений которыми проверятся соответствие домену
func (obj *StructDef) GetRegExp() (regExpArr []string) { return }

//

// GetWork Получение ID работы из url
func (obj *StructDef) GetWork(url string) (idWork string, err error) {
	err = defError("GetWork")
	return
}

// GeAuthor Получение ID автора из url
func (obj *StructDef) GeAuthor(url string) (idWork string, idAutor string, err error) {
	err = defError("GeAuthor")
	return
}

// GetChapter Получение ID главы из url
func (obj *StructDef) GetChapter(url string) (idWork string, idChapter string, err error) {
	err = defError("GetChapter")
	return
}

//

// UrlWork Генерация URL к произведению по ID
func (obj *StructDef) UrlWork(idWork string) (url string, err error) {
	err = defError("UrlWork")
	return
}

// UrlAuthor Генерация URL к автору по ID
func (obj *StructDef) UrlAuthor(idWork string, idAutor string) (url string, err error) {
	err = defError("UrlAuthor")
	return
}

// UrlChapter Генерация URL к главе по ID
func (obj *StructDef) UrlChapter(idWork string, idChapter string) (url string, err error) {
	err = defError("UrlChapter")
	return
}

/*##*/

// Ping Проверка доступности домена
func (obj *StructDef) Ping() (err error) {
	err = defError("Ping")
	return
}

// PingWork Проверка доступности работы
func (obj *StructDef) PingWork(idWork string) (err error) {
	err = defError("PingWork")
	return
}

// PingChapter Проверка доступности главы
func (obj *StructDef) PingChapter(idWork string, idChapter string) (err error) {
	err = defError("PingChapter")
	return
}

//

// LoadWork Загрузка работы полностью по ID
func (obj *StructDef) LoadWork(idWork string) (globalObj GlobalObj, err error) {
	err = defError("LoadWork")
	return
}

// LoadWorkFromFile Загрузка работы полностью из файла.
func (obj *StructDef) LoadWorkFromFile(fileType string, htmlText io.Reader) (globalObj GlobalObj, err error) {
	err = defError("LoadWorkFromFile")
	return
}

// LoadInfo Загрузка информация о работе по ID
func (obj *StructDef) LoadInfo(idWork string) (infoObj InfoObj, err error) {
	err = defError("LoadInfo")
	return
}

// LoadChapter Загрузка главы по ID
func (obj *StructDef) LoadChapter(idWork string, idChapter string) (pageObj PageObj, err error) {
	err = defError("LoadChapter")
	return
}
