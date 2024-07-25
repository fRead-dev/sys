package sys

import "time"

//###########################################################//

type TimesObj struct {
	Create time.Time
	Update time.Time
}

type SizesObj struct {
	Bytes   uint64
	Letters uint64
}

type IntegrityChecksObj struct {
	Sign string // Подпись объекта, формируются из хешей и домена
	Hash string // Хеш-сумма объекта для проверки целостности
}

type DataObj struct {
	IsCompress bool
	Data       []byte
}
