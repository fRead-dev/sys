package ParserInterface

import "io"

type BasicMethods interface {
	Hash(data *[]byte) (hash string)                    // Hash Получение контрольной суммы. Реализация может различаться между парсерами
	Decompression(data *[]byte) (decompressData []byte) // Decompression Распаковка сжатых данных. Реализация может различаться между парсерами

	ThisDomain(url string) (isValid bool) // ThisDomain Проверка соответствует ли URL данному домену

	GetDomain() (domain string)      // GetDomain Получение переменной домена
	GetVersion() (version string)    // GetVersion Получение версии модуля
	GetRegExp() (regExpArr []string) // GetRegExp Получение списка регулярных выражений которыми проверятся соответствие домену

	GetWork(url string) (err error, idWork string)                      // GetWork Получение ID работы из url
	GeAuthor(url string) (err error, idWork string, idAutor string)     // GeAuthor Получение ID автора из url
	GetChapter(url string) (err error, idWork string, idChapter string) // GetChapter Получение ID главы из url

	UrlWork(idWork string) (url string)                      // UrlWork Генерация URL к произведению по ID
	UrlAuthor(idWork string, idAutor string) (url string)    // UrlAuthor Генерация URL к автору по ID
	UrlChapter(idWork string, idChapter string) (url string) // UrlChapter Генерация URL к главе по ID
}

type LoadMethods interface {
	Ping() (err error)                                       // Ping Проверка доступности домена
	PingWork(idWork string) (err error)                      // PingWork Проверка доступности работы
	PingChapter(idWork string, idChapter string) (err error) // PingChapter Проверка доступности главы

	LoadWork(idWork string) (err error, globalObj GlobalObj)                               // LoadWork Загрузка работы полностью по ID
	LoadWorkFromFile(fileType string, htmlText io.Reader) (err error, globalObj GlobalObj) // LoadWorkFromFile Загрузка работы полностью из файла. Если домен не поддерживает скачивание файлов то должна быть заглушка с ошибкой

	LoadInfo(idWork string) (err error, infoObj InfoObj)                      // LoadInfo Загрузка информация о работе по ID
	LoadChapter(idWork string, idChapter string) (err error, pageObj PageObj) // LoadChapter Загрузка главы по ID
}
