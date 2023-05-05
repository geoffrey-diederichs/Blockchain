package blockchain

import (
	"crypto"
	"crypto/sha256"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
)

type KeyPair struct {
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

func HashMsg(msg string) []byte {
	hash := sha256.New()
	_, err := hash.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

	return hash.Sum(nil)
}

func GenKeyPair() KeyPair {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	return KeyPair{ key, &key.PublicKey }
}

func SavePrivateKey(path string, key *rsa.PrivateKey) {
	SaveData(path, x509.MarshalPKCS1PrivateKey(key))
}

func OpenPrivateKey(path string) *rsa.PrivateKey {
	data := OpenData(path)
	key, _ := x509.ParsePKCS1PrivateKey(data)

	return key
}

func Sign(hash []byte, privKey *rsa.PrivateKey) []byte {
	signature, err := rsa.SignPSS(rand.Reader, privKey, crypto.SHA256, hash, nil)
	if err != nil {
		panic(err)
	}

	return signature
}

func VerifySignature(hash []byte, signature []byte, pubKey *rsa.PublicKey) int {
	err := rsa.VerifyPSS(pubKey, crypto.SHA256, hash, signature, nil)
	if err != nil {
		return 0
	}

	return 1
}