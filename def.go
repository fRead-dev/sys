package ParserInterface

import "io"

type StructDef struct{}

func (obj *StructDef) Hash(data *[]byte) (hash string)                    { return }
func (obj *StructDef) Decompression(data *[]byte) (decompressData []byte) { return }

func (obj *StructDef) ThisDomain(url string) (isValid bool) { return }

func (obj *StructDef) GetDomain() (domain string)      { return }
func (obj *StructDef) GetVersion() (version string)    { return }
func (obj *StructDef) GetRegExp() (regExpArr []string) { return }

func (obj *StructDef) GetWork(url string) (err error, idWork string)                      { return }
func (obj *StructDef) GeAuthor(url string) (err error, idWork string, idAutor string)     { return }
func (obj *StructDef) GetChapter(url string) (err error, idWork string, idChapter string) { return }

func (obj *StructDef) UrlWork(idWork string) (url string)                      { return }
func (obj *StructDef) UrlAuthor(idWork string, idAutor string) (url string)    { return }
func (obj *StructDef) UrlChapter(idWork string, idChapter string) (url string) { return }

////

func (obj *StructDef) Ping() (err error)                                       { return }
func (obj *StructDef) PingWork(idWork string) (err error)                      { return }
func (obj *StructDef) PingChapter(idWork string, idChapter string) (err error) { return }

func (obj *StructDef) LoadWork(idWork string) (err error, globalObj GlobalObj) { return }
func (obj *StructDef) LoadWorkFromFile(fileType string, htmlText io.Reader) (err error, globalObj GlobalObj) {
	return
}
func (obj *StructDef) LoadInfo(idWork string) (err error, infoObj InfoObj) { return }
func (obj *StructDef) LoadChapter(idWork string, idChapter string) (err error, pageObj PageObj) {
	return
}
