package compress

import (
	"github.com/andybalholm/brotli"
	"io"
)

//###########################################################//

func Encode(input io.Reader, output io.Writer) error {
	encoder := brotli.NewWriterLevel(output, brotli.BestCompression)

	_, err := io.Copy(encoder, input)
	if err != nil {
		return err
	}

	err = encoder.Close()
	if err != nil {
		return err
	}

	return nil
}

func Decode(input io.Reader, output io.Writer) error {
	decoder := brotli.NewReader(input)

	_, err := io.Copy(output, decoder)
	if err != nil {
		return err
	}

	return nil
}
