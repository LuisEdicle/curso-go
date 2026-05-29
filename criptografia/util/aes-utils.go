package util

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"math/rand"
	"os"
)

func EncryptLargFiles(infile, outfile string, key []byte) error {
	buf := make([]byte, 4096)
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer out.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Read, iv); err != nil {
		return err
	}

	stream := cipher.NewCTR(block, iv)
	for {
		n, err := in.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf, buf[:n])
			out.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	out.Write(iv)
	return nil
}
