package ParserInterface

import "time"

//###########################################################//

type TimestampObj struct {
	Create time.Time `json:"create"` // Время создания
	Update time.Time `json:"update"` // Время последнего обновления
}

type SizeObj struct {
	Bytes   uint64 `json:"bytes"`   // размер в байтах
	Letters uint64 `json:"letters"` // размер в символах без учета разметки
}

type PersonalityObj struct {
	UID  UIDType `json:"uid"`  // UID в рамках домена
	Name string  `json:"name"` // Название/описание для UID
}

//###########################################################//

type StatusTag byte
type RatingTag byte
type FocusTag byte

type LanguageType string // Тим языка по ISO-639

type UIDType string
type HashType string
type DomainType string

type TagType string
type FandomType string
type PersonType string

type LoadWebDataType []byte

//###########################################################//
