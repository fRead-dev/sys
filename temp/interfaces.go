package temp

import "io"

type BasicMethods interface {
	Hash(data *[]byte) (hash string)      // Hash [MUST BE] Получение контрольной суммы. Реализация может различаться между парсерами
	ThisDomain(url string) (isValid bool) // ThisDomain [MUST BE] Проверка соответствует ли

	GetDomain() (domain string)   // GetDomain [MUST BE] Получение переменной домена
	GetVersion() (version string) // GetVersion [MUST BE] Получение версии модуля

	GetRegExp() (regExpArr []string)                                    // GetRegExp [MUST BE] Получение списка регулярных выражений которыми проверятся соответствие домену
	GetWork(url string) (isOK bool, idWork string)                      // GetWork [MUST BE] Получение ID работы из url
	GeAuthor(url string) (isOK bool, idWork string, idAutor string)     // GeAuthor [MUST BE] Получение ID автора из url
	GetChapter(url string) (isOK bool, idWork string, idChapter string) // GetChapter [MUST BE] Получение ID главы из url

	UrlWork(idWork string) (url string)                      // UrlWork [MUST BE] Генерация URL к произведению по ID
	UrlAuthor(idWork string, idAutor string) (url string)    // UrlAuthor [MUST BE] Генерация URL к автору по ID
	UrlChapter(idWork string, idChapter string) (url string) // UrlChapter [MUST BE] Генерация URL к главе по ID
}

type LoadMethods interface {
	Ping() (isAvailable bool)                                       // Ping [MUST BE] Проверка доступности домена
	PingWork(idWork string) (isAvailable bool)                      // PingWork [MUST BE] Проверка доступности работы
	PingChapter(idWork string, idChapter string) (isAvailable bool) // PingChapter [MUST BE] Проверка доступности главы

	LoadWork(idWork string) (globalObj GlobalObj)                  // LoadWork [MUST BE] Загрузка работы полностью по ID
	LoadInfo(idWork string) (infoObj InfoObj)                      // LoadInfo [MUST BE] Загрузка информация о работе по ID
	LoadChapter(idWork string, idChapter string) (pageObj PageObj) // LoadChapter [MUST BE] Загрузка главы по ID
}

type ParserMethods interface {
	ParseInfo(htmlText io.Reader) (infoObj InfoObj)    // ParseInfo [MUST BE]	Парсинг информации о работе из html
	ParseChapter(htmlText io.Reader) (pageObj PageObj) // ParseChapter [MUST BE] Парсинг главы из html
}
