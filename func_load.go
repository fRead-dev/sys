package sys

//###########################################################//

func Upload() error {
	if !IsInit() {
		return ErrGlobalNotInit
	}

	if IsAuthor() {
		data, err := Methods.Net.Author(protectParser.domain, protectParser.author.UID)
		if err != nil {
			return err
		}
		protectParser.load = &data
		return nil
	}

	if IsWork() {
		data, err := Methods.Net.Work(protectParser.domain, protectParser.work.UID)
		if err != nil {
			return err
		}
		protectParser.load = &data
		return nil
	}

	return ErrGlobalNil
}

////////

func ReadWorksFromAuthor() (*[]PersonalityObj, error) {
	if !IsInit() {
		return nil, ErrGlobalNotInit
	}
	if protectParser.load == nil {
		return nil, ErrGlobalNotLoad
	}
	if !IsAuthor() {
		return nil, ErrGlobalIsNotAuthor
	}

	arr, err := Methods.Parse.WorkChapters(protectParser.load)
	return &arr, err
}

func ReadChaptersFromWork() (*[]PersonalityObj, error) {
	if !IsInit() {
		return nil, ErrGlobalNotInit
	}
	if protectParser.load == nil {
		return nil, ErrGlobalNotLoad
	}
	if !IsWork() {
		return nil, ErrGlobalIsNotWork
	}

	arr, err := Methods.Parse.WorkChapters(protectParser.load)
	return &arr, err
}

////////

func ReadAuthor() (*PersonalityObj, error) {
	if !IsInit() {
		return nil, ErrGlobalNotInit
	}
	if protectParser.load == nil {
		return nil, ErrGlobalNotLoad
	}
	if !IsAuthor() {
		return nil, ErrGlobalIsNotAuthor
	}

	data, err := Methods.Parse.Author(protectParser.load)
	return &data, err
}

func ReadWork() (*WorkObj, error) {
	if !IsInit() {
		return nil, ErrGlobalNotInit
	}
	if protectParser.load == nil {
		return nil, ErrGlobalNotLoad
	}
	if !IsWork() {
		return nil, ErrGlobalIsNotWork
	}

	data, err := Methods.Parse.Work(protectParser.load)
	return &data, err
}

//###########################################################//
