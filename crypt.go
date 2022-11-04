package main

import (
	"encoding/base64"

	"github.com/jurgen-kluft/go-pass/crypter"
)

type Cryptor struct {
	key    []byte
	len    int
	tweak  []uint64
	cipher *crypter.Cipher
}

// New will instanciated a Cryptor with the given size in bits (256, 512 or 1024)
// e.g. New("pigs cannot fly and chickens are aliens", 1024)
func New(keyAsText string, bits int) *Cryptor {
	c := &Cryptor{}

	c.len = bits / 8
	c.key = make([]byte, c.len)
	c.key = []byte(keyAsText)
	copy(c.key, []byte(keyAsText))
	c.tweak = []uint64{0x0706050403020100, 0x0F0E0D0C0B0A0908}

	c.cipher, _ = crypter.New(c.key, c.tweak)
	return c
}

func (c *Cryptor) Encrypt(text string) string {
	plaintext := []byte(text)
	encodedtext := make([]byte, len(plaintext))
	c.cipher.Encrypt(encodedtext, plaintext)
	return base64.URLEncoding.EncodeToString(encodedtext)
}

func (c *Cryptor) Decrypt(encodedbase64text string) string {
	encodedbytes, _ := base64.URLEncoding.DecodeString(encodedbase64text)
	decodedbytes := make([]byte, len(encodedbytes))
	c.cipher.Decrypt(encodedbytes, decodedbytes)
	return string(decodedbytes)
}
