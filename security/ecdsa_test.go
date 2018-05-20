package security

import (
	"testing"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"encoding/base64"
)

func TestEcdsa(t *testing.T) {

	// Private key
	c := elliptic.P521()
	privateKey, err := ecdsa.GenerateKey(c,rand.Reader)
	if err != nil {
		t.Error(err)
		return
	}
	dataPrivateKey,err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		t.Error(err)
		return
	}
	strPrivateKey := base64.RawStdEncoding.EncodeToString(dataPrivateKey)
	fmt.Println("Private Key : ",strPrivateKey)

	// Public key
	publicKey := privateKey.Public()
	dataPublicKey,err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		t.Error(err)
		return
	}
	strPublicKey := base64.RawStdEncoding.EncodeToString(dataPublicKey)
	fmt.Println("Public Key : ",strPublicKey)


	// TODO message:
}