package mdf

import fReadText "github.com/fRead-dev/sys/data/text"

//###########################################################//

var format = map[fReadText.FormatType][2]string{
	fReadText.FormatBold:          {"**", "**"},
	fReadText.FormatItalic:        {"_", "_"},
	fReadText.FormatUnderline:     {"__", "__"},
	fReadText.FormatStrikethrough: {"~~", "~~"},
	fReadText.FormatNewLine:       {"\n", ""},

	fReadText.FormatLargerText: {"[#:1](", ")"},
	fReadText.FormatSmallText:  {"[#:2](", ")"},

	fReadText.FormatSuperscript: {"[#:3](", ")"},
	fReadText.FormatSubscript:   {"[#:4](", ")"},
	fReadText.FormatBlockquote:  {"`", "`"},

	fReadText.FormatLeftText:   {"[#:7](", ")"},
	fReadText.FormatRightText:  {"[#:8](", ")"},
	fReadText.FormatCenterText: {"[#:9](", ")"},
}

func Format() map[fReadText.FormatType][2]string {
	return format
}
