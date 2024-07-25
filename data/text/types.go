package text

//###########################################################//

type FormatType byte

const (
	FormatBold          FormatType = 1 //	Жирный
	FormatItalic        FormatType = 2 //	Курсив
	FormatUnderline     FormatType = 3 //	подчеркивания под текстом.
	FormatStrikethrough FormatType = 4 //	текст, перечеркнутый горизонтальной линией
	FormatNewLine       FormatType = 5 //	Переход на новую строку

	FormatLargerText FormatType = 51 //	 увеличенный размер шрифта для создания визуального акцента или для заголовков.
	FormatSmallText  FormatType = 52 //	уменьшенный размер шрифта, может использоваться для заметок или дисклеймеров.

	FormatSuperscript FormatType = 10 //	маленький текст, размещенный выше строки письма.
	FormatSubscript   FormatType = 20 //	 маленький текст, размещенный ниже строки письма.
	FormatBlockquote  FormatType = 30 //	для выделения цитат или значимых утверждений

	FormatLeftText   FormatType = 81 //	текст по левому краю
	FormatRightText  FormatType = 82 //	текст по правому краю
	FormatCenterText FormatType = 83 //	текст по центру
)
