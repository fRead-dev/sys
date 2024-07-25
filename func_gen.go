package sys

//###########################################################//

func genUrlDomain(domain DomainType) string {
	return "https://" + string(domain)
}

func genUrlAuthor(domain DomainType, author UIDType) string {
	return "https://" + string(domain) + "/" + string(author)
}

func genUrlWork(domain DomainType, work UIDType) string {
	return "https://" + string(domain) + "/" + string(work)
}

func genUrlWorkChapter(domain DomainType, work UIDType, chapter UIDType) string {
	return "https://" + string(domain) + "/" + string(work) + "/" + string(chapter)
}

////

func Chapter(data LoadWebDataType) (*ChapterObj, error) {
	personality, err := Methods.Parse.Field.Chapter.Personality(&data)
	if err != nil {
		return nil, err
	}

	timestamp, err := Methods.Parse.Field.Chapter.Timestamp(&data)
	if err != nil {
		return nil, err
	}

	text, err := Methods.Parse.Field.Chapter.Data(&data)
	if err != nil {
		return nil, err
	}

	letters, err := Methods.Gen.Text.Letters(&text)
	if err != nil {
		return nil, err
	}

	hash := Methods.SYS.Hash([]byte(text))
	Bytes := Methods.Gen.Text.Bytes(&text)

	obj := ChapterObj{}
	obj.Chapter = personality
	obj.Timestamp = timestamp
	obj.Hash = hash
	obj.Data = text
	obj.Size = SizeObj{
		Letters: letters,
		Bytes:   Bytes,
	}

	return &obj, nil
}

func Work(data LoadWebDataType) (*WorkObj, error) {
	personality, err := Methods.Parse.Field.Work.Personality(&data)
	if err != nil {
		return nil, err
	}

	timestamp, err := Methods.Parse.Field.Work.Timestamp(&data)
	if err != nil {
		return nil, err
	}

	text, err := Methods.Parse.Field.Work.Description(&data)
	if err != nil {
		return nil, err
	}

	lang, err := Methods.Parse.Field.Work.Language(&data)
	if err != nil {
		return nil, err
	}

	Tags, err := Methods.Parse.Field.Work.Tags(&data)
	if err != nil {
		return nil, err
	}

	Fandoms, err := Methods.Parse.Field.Work.Fandoms(&data)
	if err != nil {
		return nil, err
	}

	Personages, err := Methods.Parse.Field.Work.Personages(&data)
	if err != nil {
		return nil, err
	}

	Status, err := Methods.Parse.Field.Work.Status(&data)
	if err != nil {
		return nil, err
	}

	Rating, err := Methods.Parse.Field.Work.Rating(&data)
	if err != nil {
		return nil, err
	}

	Focus, err := Methods.Parse.Field.Work.Focus(&data)
	if err != nil {
		return nil, err
	}

	obj := WorkObj{}
	obj.Work = personality
	obj.Language = lang
	obj.Tags = Tags
	obj.Fandoms = Fandoms
	obj.Personages = Personages
	obj.Status = Status
	obj.Rating = Rating
	obj.Focus = Focus
	obj.Description = text
	obj.Timestamp = timestamp

	buf := []byte(text)
	buf = append(buf, byte(Status))
	buf = append(buf, byte(Rating))
	buf = append(buf, byte(Focus))
	buf = append(buf, []byte(lang)...)
	buf = append(buf, []byte(timestamp.Create.String())...)
	buf = append(buf, []byte(timestamp.Update.String())...)
	for _, txt := range Tags {
		buf = append(buf, []byte(txt)...)
	}
	for _, txt := range Fandoms {
		buf = append(buf, []byte(txt)...)
	}
	for _, txt := range Personages {
		buf = append(buf, []byte(txt)...)
	}
	buf = append(buf, []byte(text)...)

	obj.Hash = Methods.SYS.Hash(buf)

	return &obj, nil
}

////////

func genNormalisedText(data LoadWebDataType) (NormalisedTextType, error) {
	return []byte{}, ErrGenTextLettersNil
}

func genLettersSum(data *NormalisedTextType) (uint64, error) {
	return 0, ErrGenTextNormalisedNil
}

func genBytesSum(data *NormalisedTextType) uint64 {
	return uint64(len(*data))
}
