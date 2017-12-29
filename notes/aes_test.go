package notes

import (
	"log"
	"testing"
)

func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := encrypt("halulu.png", "e")
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func BenchmarkEncryptSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := encryptSlice("halulu.png", "e")
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func BenchmarkDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := encrypt("e", "halulu.png")
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func BenchmarkDecryptSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := encryptSlice("e", "halulu.png")
		if err != nil {
			log.Println(err.Error())
		}
	}
}
