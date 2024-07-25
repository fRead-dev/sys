package md

import fReadText "github.com/fRead-dev/sys/data/text"

//###########################################################//

var format = map[fReadText.FormatType][2]string{
	fReadText.FormatBold:          {"**", "**"},
	fReadText.FormatItalic:        {"*", "*"},
	fReadText.FormatUnderline:     {"~~", "~~"},
	fReadText.FormatStrikethrough: {"`", "`"},
	fReadText.FormatNewLine:       {"\n", ""},
}

func Format() map[fReadText.FormatType][2]string {
	return format
}
