package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
	version        = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)
	address := Base58Encode(fullHash)

	// fmt.Printf("pub key: %x\n", w.PublicKey)
	// fmt.Printf("pub hash: %x\n", pubHash)
	// fmt.Printf("address: %x\n", address)

	return address

}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256() //output will be 256 bytes

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	//elliptic curve multiplication to create public key
	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

// to create address ,
// take pub key run through sha256 hashing algo
// then take through ripemd160 hashing algo
// these steps generate the public key hash
// now take pkh and pass through sha256 hash algo twice
// take off the first 4 bytes of the resulting hash and
// that will be the checksum
// then take the checksum, original pkh, and version value
// and concatenate all these values and pass through a
// base 58 encoder algo to create the address

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}
