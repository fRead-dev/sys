package ParserInterface

import "errors"

//###########################################################//

var (
	ErrNet202 = errors.New(`request accepted but not processed`)
	ErrNet204 = errors.New(`page loaded but no data`)

	ErrNet3xx = errors.New(`page has been moved to another address`)

	ErrNet400 = errors.New(`the server cannot process the request due to invalid syntax`)
	ErrNet401 = errors.New(`authentication required`)
	ErrNet403 = errors.New(`no rights to access the page`)
	ErrNet404 = errors.New(`page not found`)
	ErrNet405 = errors.New(`request method not allowed`)
	ErrNet408 = errors.New(`the server waits too long to complete the request`)
	ErrNet429 = errors.New(`too many requests in a certain time`)

	ErrNet5xx = errors.New(`server side error`)
)

//###########################################################//

var (
	ErrGlobalNil         = errors.New("undefined error")
	ErrGlobalNotInit     = errors.New("not init")
	ErrGlobalIsNotAuthor = errors.New("this is not author")
	ErrGlobalIsNotWork   = errors.New("this is not work")
	ErrGlobalNotLoad     = errors.New("data not loaded")
)

//###########################################################//

var (
	ErrParseUrlDomainNil  = errors.New("method `Methods.Parse.URL.Domain` not init")
	ErrParseUrlAuthorNil  = errors.New("method `Methods.Parse.URL.Author` not init")
	ErrParseUrlWorkNil    = errors.New("method `Methods.Parse.URL.Work` not init")
	ErrParseUrlChapterNil = errors.New("method `Methods.Parse.URL.Chapter` not init")

	ErrParseAuthorNil  = errors.New("method Methods.Parse.Author` not init")
	ErrParseWorkNil    = errors.New("method Methods.Parse.Work` not init")
	ErrParseChapterNil = errors.New("method Methods.Parse.Chapter` not init")

	ErrParseAuthorWorksNil  = errors.New("method Methods.Parse.AuthorWorks` not init")
	ErrParseWorkChaptersNil = errors.New("method Methods.Parse.WorkChapters` not init")

	ErrParseFieldChapterPersonalityNil = errors.New("method Methods.Parse.Field.Chapter.Personality ` not init")
	ErrParseFieldChapterTimestampNil   = errors.New("method Methods.Parse.Field.Chapter.Timestamp ` not init")
	ErrParseFieldChapterDataNil        = errors.New("method Methods.Parse.Field.Chapter.Data ` not init")

	ErrParseFieldWorkPersonalityNil = errors.New("method Methods.Parse.Field.Work.Personality ` not init")
	ErrParseFieldWorkTimestampNil   = errors.New("method Methods.Parse.Field.Work.Timestamp ` not init")
	ErrParseFieldWorkLanguageNil    = errors.New("method Methods.Parse.Field.Work.Language ` not init")
	ErrParseFieldWorkTagsNil        = errors.New("method Methods.Parse.Field.Work.Tags ` not init")
	ErrParseFieldWorkFandomsNil     = errors.New("method Methods.Parse.Field.Work.Fandoms ` not init")
	ErrParseFieldWorkPersonagesNil  = errors.New("method Methods.Parse.Field.Work.Personages ` not init")
	ErrParseFieldWorkStatusNil      = errors.New("method Methods.Parse.Field.Work.Status ` not init")
	ErrParseFieldWorkRatingNil      = errors.New("method Methods.Parse.Field.Work.Rating ` not init")
	ErrParseFieldWorkFocusNil       = errors.New("method Methods.Parse.Field.Work.Focus ` not init")
	ErrParseFieldWorkDescriptionNil = errors.New("method Methods.Parse.Field.Work.Description ` not init")
)

var (
	ErrNetPingNil    = errors.New("method `Methods.Net.Ping` not init")
	ErrNetAuthorNil  = errors.New("method `Methods.Net.Author` not init")
	ErrNetWorkNil    = errors.New("method `Methods.Net.Work` not init")
	ErrNetChapterNil = errors.New("method `Methods.Net.Chapter` not init")
)

var (
	ErrGenTextLettersNil    = errors.New("method `Methods.Gen.Text.Letters` not init")
	ErrGenTextNormalisedNil = errors.New("method `Methods.Gen.Text.Normalised` not init")
)
