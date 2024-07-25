package text

import (
	"bufio"
	"io"
	"regexp"
)

//###########################################################//

func CountLetters(input io.Reader) uint64 {
	reader := bufio.NewReader(input)
	re := regexp.MustCompile(`^[a-zA-Z]$`)
	var count uint64

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}

		if re.MatchString(string(char)) {
			count++
		}
	}

	return count
}
