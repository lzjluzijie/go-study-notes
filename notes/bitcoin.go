package notes

import (
	"crypto/sha256"
	"fmt"

	"github.com/decred/base58"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func init() {
	addCommand(cli.Command{
		Name:    "bitcoin",
		Aliases: []string{"btc"},
		Usage:   "Something about bitcoin.",
		Subcommands: []cli.Command{
			{
				Name:    "new",
				Aliases: []string{"new"},
				Usage:   "Generate a new bitcoin wallet.",
				Action:  newWallet,
			},
		},
	})
}

//NewWallet generates a new bitcoin wallet.
func newWallet(c *cli.Context) (err error) {
	s := []byte(c.Args().First())
	//Generate secp256k1 private key
	h1 := sha3.Sum256(s)
	h2 := sha3.Sum256(h1[:])
	private := h2[:]
	fmt.Printf("Secp256k1 private key in hexadecimal is : %X\n", h2)

	//WIF private key
	k := append([]byte{0x80}, private...)
	hk := sha2(k)
	key := append(k, hk[:4]...)
	wif := base58.Encode(key)
	fmt.Printf("WIF private key is : %s\n", wif)

	//Address
	_, pub := secp256k1.PrivKeyFromBytes(private)
	s256 := sha256.Sum256(pub.SerializeUncompressed())
	r := ripemd160.New()
	r.Write(s256[:])
	r160 := r.Sum(nil)
	addr := append([]byte{0x00}, r160...)
	ha := sha2(addr)
	address := base58.Encode(append(addr, ha[:4]...))
	fmt.Printf("Address is : %s\n", address)

	return nil
}

///SHA-256 twice
func sha2(data []byte) (hash []byte) {
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}
