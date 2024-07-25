package crypt

import (
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
	"hash"
	"io"
)

//###########################################################//

func retString(h hash.Hash, input io.Reader) (string, error) {
	if _, err := io.Copy(h, input); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func genKey(key string) ([]byte, error) {
	h, err := blake2b.New256(nil)
	if err != nil {
		return nil, err
	}

	_, err = h.Write([]byte(key))
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

////

func hash32(input io.Reader, key []byte) (string, error) {
	h, err := blake2b.New256(key)
	if err != nil {
		return "", err
	}

	return retString(h, input)
}

func hash48(input io.Reader, key []byte) (string, error) {
	h, err := blake2b.New384(key)
	if err != nil {
		return "", err
	}

	return retString(h, input)
}

func hash64(input io.Reader, key []byte) (string, error) {
	h, err := blake2b.New512(key)
	if err != nil {
		return "", err
	}

	return retString(h, input)
}

////

func Hash32(input io.Reader) (string, error) {
	return hash32(input, nil)
}

func Hash48(input io.Reader) (string, error) {
	return hash48(input, nil)
}

func Hash64(input io.Reader) (string, error) {
	return hash64(input, nil)
}

////

func Hash32Key(input io.Reader, key string) (string, error) {
	k, err := genKey(key)
	if err != nil {
		return "", err
	}

	return hash32(input, k)
}

func Hash48Key(input io.Reader, key string) (string, error) {
	k, err := genKey(key)
	if err != nil {
		return "", err
	}

	return hash48(input, k)
}

func Hash64Key(input io.Reader, key string) (string, error) {
	k, err := genKey(key)
	if err != nil {
		return "", err
	}

	return hash64(input, k)
}
