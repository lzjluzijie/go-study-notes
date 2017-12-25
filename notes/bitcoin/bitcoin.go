package bitcoin

import (
	"crypto/sha256"
	"fmt"

	"github.com/decred/base58"
	"github.com/urfave/cli"
	"golang.org/x/crypto/sha3"
)

func NewWallet(c *cli.Context) error {
	//Generate secp256k1 private key
	s := []byte("Hello World!")
	h1 := sha3.Sum256(s)
	h2 := sha3.Sum256(h1[:])
	fmt.Printf("Secp256k1 private key in hexadecimal is : %x\n", h2)

	//WIF
	k := append([]byte{0x80}, h2[:]...)
	s1 := sha256.Sum256(k)
	s2 := sha256.Sum256(s1[:])
	key := append(k, s2[:4]...)
	b58 := base58.Encode(key)
	fmt.Printf("WIF private key is : %s\n", b58)

	return nil
}
