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
	ErrParseDomainNil  = errors.New("method `Methods.Parse.Domain` not init")
	ErrParseAuthorNil  = errors.New("method Methods.Parse.Author` not init")
	ErrParseWorkNil    = errors.New("method Methods.Parse.Work` not init")
	ErrParseChapterNil = errors.New("method Methods.Parse.Chapter` not init")

	ErrNetPingNil    = errors.New("method `Methods.Net.Ping` not init")
	ErrNetAuthorNil  = errors.New("method `Methods.Net.Author` not init")
	ErrNetWorkNil    = errors.New("method `Methods.Net.Work` not init")
	ErrNetChapterNil = errors.New("method `Methods.Net.Chapter` not init")
)
