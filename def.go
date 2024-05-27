package ParserInterface

import "io"

type StructDef struct{}

// Hash Получение контрольной суммы.
func (obj *StructDef) Hash(data *[]byte) (hash string) { return }

// Сompression	Упаковка полученных данных для буферизации
func (obj *StructDef) Сompression(data *[]byte) (compressData []byte) { return }

// Decompression Распаковка сжатых данных
func (obj *StructDef) Decompression(data *[]byte) (decompressData []byte) { return }

//

// ThisDomain Проверка соответствует ли URL данному домену
func (obj *StructDef) ThisDomain(url string) (isValid bool) { return }

//

// GetDomain Получение переменной домена
func (obj *StructDef) GetDomain() (domain string) { return }

// GetVersion Получение версии модуля
func (obj *StructDef) GetVersion() (version string) { return }

// GetRegExp Получение списка регулярных выражений которыми проверятся соответствие домену
func (obj *StructDef) GetRegExp() (regExpArr []string) { return }

//

// GetWork Получение ID работы из url
func (obj *StructDef) GetWork(url string) (err error, idWork string) { return }

// GeAuthor Получение ID автора из url
func (obj *StructDef) GeAuthor(url string) (err error, idWork string, idAutor string) { return }

// GetChapter Получение ID главы из url
func (obj *StructDef) GetChapter(url string) (err error, idWork string, idChapter string) { return }

//

// UrlWork Генерация URL к произведению по ID
func (obj *StructDef) UrlWork(idWork string) (url string) { return }

// UrlAuthor Генерация URL к автору по ID
func (obj *StructDef) UrlAuthor(idWork string, idAutor string) (url string) { return }

// UrlChapter Генерация URL к главе по ID
func (obj *StructDef) UrlChapter(idWork string, idChapter string) (url string) { return }

/*##*/

// Ping Проверка доступности домена
func (obj *StructDef) Ping() (err error) { return }

// PingWork Проверка доступности работы
func (obj *StructDef) PingWork(idWork string) (err error) { return }

// PingChapter Проверка доступности главы
func (obj *StructDef) PingChapter(idWork string, idChapter string) (err error) { return }

//

// LoadWork Загрузка работы полностью по ID
func (obj *StructDef) LoadWork(idWork string) (err error, globalObj GlobalObj) { return }

// LoadWorkFromFile Загрузка работы полностью из файла.
func (obj *StructDef) LoadWorkFromFile(fileType string, htmlText io.Reader) (err error, globalObj GlobalObj) {
	return
}

// LoadInfo Загрузка информация о работе по ID
func (obj *StructDef) LoadInfo(idWork string) (err error, infoObj InfoObj) { return }

// LoadChapter Загрузка главы по ID
func (obj *StructDef) LoadChapter(idWork string, idChapter string) (err error, pageObj PageObj) {
	return
}
