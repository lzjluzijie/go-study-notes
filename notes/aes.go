package notes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func init() {
	addCommand(cli.Command{
		Name:    "aes",
		Aliases: []string{"aes"},
		Usage:   "AES.",
		Subcommands: []cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Usage:   "Encrypt a file.",
				Action: func(c *cli.Context) error {
					if c.NArg() >= 2 {
						encrypt(c.Args().Get(0), c.Args().Get(1))
					} else {
						log.Println("Not enough arguments.")
					}
					return nil
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Usage:   "Decrypt a file.",
				Action: func(c *cli.Context) error {
					if c.NArg() >= 2 {
						decrypt(c.Args().Get(0), c.Args().Get(1))
					} else {
						log.Println("Not enough arguments.")
					}
					return nil
				},
			},
		},
	})
}

func encrypt(in, out string) (err error) {
	key := []byte("91f84c7bb89d02bd30812f3ddae2b048")

	inFile, err := os.Open(in)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer inFile.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	outFile, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer outFile.Close()

	writer := &cipher.StreamWriter{S: stream, W: outFile}
	if _, err := io.Copy(writer, inFile); err != nil {
		log.Fatalln(err.Error())
	}

	return
}

func decrypt(in, out string) (err error) {
	key := []byte("91f84c7bb89d02bd30812f3ddae2b048")

	inFile, err := os.Open(in)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer inFile.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	outFile, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer outFile.Close()

	reader := &cipher.StreamReader{S: stream, R: inFile}
	if _, err := io.Copy(outFile, reader); err != nil {
		log.Fatalln(err.Error())
	}
	return
}

func encryptSlice(in, out string) (err error) {
	key := []byte("91f84c7bb89d02bd30812f3ddae2b048")

	data, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalln(err.Error())
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	ioutil.WriteFile(out, ciphertext, 0644)
	return
}

func decryptSlice(in, out string) (err error) {
	key := []byte("91f84c7bb89d02bd30812f3ddae2b048")

	ciphertext, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	if len(ciphertext) < aes.BlockSize {
		err = errors.New("Ciphertext too short.")
		return
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	err = ioutil.WriteFile(out, ciphertext, 0644)

	return
}
