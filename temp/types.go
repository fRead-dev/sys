package temp

type TimestampObj struct {
	Create uint64 `json:"create"` //Время создания. Секунды от эпохи Unix
	Update uint64 `json:"update"` //Время последнего обновления. Секунды от эпохи Unix
}

type SizeObj struct {
	Bytes   uint64 `json:"bytes"`   //Размер в байтах
	Letters uint64 `json:"letters"` //Количество символов без учета разметки
}

type PersonalityObj struct {
	ID   string `json:"id"`   // Уникальный указатель в рамках домена
	Name string `json:"name"` // Имя\название в рамках уникального указателя
}

//###########################################################//

// InfoObj Информация о книге
type InfoObj struct {
	Work   PersonalityObj `json:"work"`   // Указатель произведения
	Author PersonalityObj `json:"author"` // Указатель автора

	Language  string       `json:"language"`  // Язык произведения
	Timestamp TimestampObj `json:"timestamp"` // Временные метки произведения
	Size      SizeObj      `json:"size"`      // Размер произведения
	Hash      string       `json:"hash"`      // Контрольная сумма произведения глобально

	Tags       []string `json:"tags"`       // Указаные теги
	Fandoms    []string `json:"fandoms"`    // Указанные фендомы
	Characters []string `json:"characters"` // Указанные персонажи

	Status StatusTag `json:"status"` // Статус произведения
	Rating RatingTag `json:"rating"` // Рейтинг произведения
	Focus  FocusTag  `json:"focus"`  // Направление произведения

	Description []byte `json:"description" compress:"true"` // Описание к произведению
}

// PageObj Содержание страницы
type PageObj struct {
	Chapter PersonalityObj `json:"chapter"` // Указатель на главу

	Timestamp TimestampObj `json:"timestamp"` // Временные метки главы
	Size      SizeObj      `json:"size"`      // Размер главы
	Hash      string       `json:"hash"`      // Контрольная сумма главы

	Data []byte `json:"data" compress:"true"` // Текст главы
}

//###########################################################//

type GlobalObj struct {
	Info  InfoObj   `json:"info"`
	Pages []PageObj `json:"pages"`
}
