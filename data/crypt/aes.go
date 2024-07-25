package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

//###########################################################//

func encode(input io.Reader, output io.Writer, key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(input); err != nil {
		return err
	}

	if buf.Len() == 0 {
		return errors.New("no data to encode")
	}

	ciphertext := make([]byte, aes.BlockSize+buf.Len())
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], buf.Bytes())
	if _, err = output.Write(ciphertext); err != nil {
		return err
	}

	return nil
}

func decode(input io.Reader, output io.Writer, key []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(input); err != nil {
		return err
	}

	if buf.Len() < aes.BlockSize {
		return errors.New("ciphertext too short")
	}

	iv := buf.Next(aes.BlockSize)
	ciphertext := buf.Bytes()
	stream := cipher.NewCFBDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)
	if _, err = output.Write(plaintext); err != nil {
		return err
	}

	return nil
}

////

func Encode(input io.Reader, output io.Writer, key string) error {
	k, err := genKey(key)
	if err != nil {
		return err
	}
	return encode(input, output, k)
}

func Decode(input io.Reader, output io.Writer, key string) error {
	k, err := genKey(key)
	if err != nil {
		return err
	}
	return decode(input, output, k)
}
