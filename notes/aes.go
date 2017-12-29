package notes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"time"

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
				Action:  encrypt,
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Usage:   "Decrypt a file.",
				Action:  decrypt,
			},
		},
	})
}

func encrypt(c *cli.Context) (err error) {
	key := []byte("1234567890123456")

	t := time.Now()

	data, err := ioutil.ReadFile("data")
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

	ioutil.WriteFile("encrypted", ciphertext, 0644)

	log.Println(time.Since(t))

	return
}

func decrypt(c *cli.Context) (err error) {
	key := []byte("1234567890123456")

	t := time.Now()

	ciphertext, err := ioutil.ReadFile("encrypted")
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

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	err = ioutil.WriteFile("decrypted", ciphertext, 0644)

	log.Println(time.Since(t))
	return
}
