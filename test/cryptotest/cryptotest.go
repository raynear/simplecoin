package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/md5"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/howeyc/gopass"
)

func err(msg string, e error) {
	if e != nil {
		fmt.Println(msg, e)
		log.Fatal(msg, e)
	}
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func encode(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func decode(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

func test() {
	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	publicKey := &privateKey.PublicKey

	encPriv, encPub := encode(privateKey, publicKey)

	fmt.Println(encPriv)
	fmt.Println(encPub)

	priv2, pub2 := decode(encPriv, encPub)

	if !reflect.DeepEqual(privateKey, priv2) {
		fmt.Println("Private keys do not match.")
	}
	if !reflect.DeepEqual(publicKey, pub2) {
		fmt.Println("Public keys do not match.")
	}
}

func main() {
	pubkeyCurve := elliptic.P256() // P256이 가장 효율적이라 함 from https://safecurves.cr.yp.to

	privateKey := new(ecdsa.PrivateKey)
	privateKey, e := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	err("Key generate Error", e)

	//	var publicKey ecdsa.PublicKey
	publicKey := &privateKey.PublicKey

	fmt.Println("Private Key ", privateKey)
	fmt.Println("Public Key ", publicKey)

	// sign
	//	var h hash.Hash
	h := md5.New()
	//	r := big.NewInt(0)
	//	s := big.NewInt(0)

	io.WriteString(h, "This is a message")
	signhash := h.Sum(nil)

	r, s, e := ecdsa.Sign(rand.Reader, privateKey, signhash)
	err("sign error", e)

	signature := r.Bytes()
	signature = append(signature, s.Bytes()...)
	fmt.Println("Signature", signature)

	// verify
	verifystatus := ecdsa.Verify(publicKey, signhash, r, s)
	fmt.Println("Verify", verifystatus)

	// pem encode & decode test
	encPriv, encPub := encode(privateKey, publicKey)

	fmt.Println(encPriv)
	fmt.Println(encPub)

	priv2, pub2 := decode(encPriv, encPub)

	if !reflect.DeepEqual(privateKey, priv2) {
		fmt.Println("Private keys do not match.")
	}
	if !reflect.DeepEqual(publicKey, pub2) {
		fmt.Println("Public keys do not match.")
	}

	// encrypt decrypt test

	origintext := "this is test text. i am raynear. i need more sleep."
	fmt.Println("Start Encrypt : ", origintext)
	fmt.Printf("Input Password : ")
	silentPassword, e := gopass.GetPasswdMasked()
	err("Input Password Error", e)
	ciphertext := encrypt([]byte(origintext), string(silentPassword))
	fmt.Println("Encrypt : ", string(ciphertext), ciphertext)
	plaintext := decrypt(ciphertext, string(silentPassword))
	fmt.Println("Decrypt : ", string(plaintext), plaintext)

	//	aPrivKey := new(ecdsa.PrivateKey)
	//	aPrivKey.Curve = privateKey.Curve

	//	privateKey2 := new(ecdsa.PrivateKey)
	//	a, _ := new(big.Int).SetString("111761706917907600925492748084378392556684259101016704467799728152493962787727", 10)
	//	privateKey2, e = ecdsa.GenerateKey(pubkeyCurve, a)

	// encrypt
	// decrypt

	// privatekey save to file
	// privatekey load from file
	// privatekey file password encrypt
}
