package sys

//###########################################################//

type PersonalityObj struct {
	UID  UIDType `json:"uid"`  // UID в рамках домена
	Name string  `json:"name"` // Название/описание для UID
}

//###########################################################//

type LanguageType string // Тим языка по ISO-639

type UIDType string
type HashType string
type DomainType string

type TagType string
type FandomType string
type PersonType string

type LoadWebDataType []byte
type NormalisedTextType []byte

//###########################################################//
