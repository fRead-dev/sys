![Fork GitHub Release](https://img.shields.io/github/v/release/fRead-dev/ParserInterface)
[![Go Report Card](https://goreportcard.com/badge/github.com/fRead-dev/ParserInterface)](https://goreportcard.com/report/github.com/fRead-dev/ParserInterface)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/fRead-dev/ParserInterface?color=orange)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/fRead-dev/ParserInterface?color=green)
![GitHub repo size](https://img.shields.io/github/repo-size/fRead-dev/ParserInterface)

## ParserInterface

Интерфейс для создания парсеров
- В любом созданном парсере ожидается минимальная реализация методов описанных в interface
- Нежелательна модификация структур
- Ожидается что теги будут браться с модуля

## Использование

В модуле описаны интерфейсы всех обязательных методов.

Для удобства использования можно наследовать шаблон структуры:
```go
import "github.com/fRead-dev/ParserInterface"

type ParserStruct struct {
	ParserInterface.StructDef
}
```
И просто переопределить только нужные методы:
```go
func (obj *ParserStruct) GetVersion()  {
	return ConstVersion
}
```

---

В шаблоне предопределены все обязательные методы. Некторые имеют готовый функционал:
- `Hash(data *[]byte) (hash string)` Используется алгоритм **blake2b** с получением хеша длинной 16 символов.
- `Сompression(data *[]byte) (compressData []byte, err error)` Используется алгоритм **brotli** с максимальным сжатием.
- `Decompression(data *[]byte) (decompressData []byte, err error)` Используется алгоритм **brotli**