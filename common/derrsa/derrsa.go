package derrsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
)

// DecodeDEREncodedRSAPrivateKey reads from the io.Reader until EOF.
// The resulting bytes are decoded as parsed as an rsa private key.
func DecodeDEREncodedRSAPrivateKey(r io.Reader) (*rsa.PrivateKey, error) {
	decoder := base64.NewDecoder(base64.StdEncoding, r)
	bts, err := ioutil.ReadAll(decoder)
	if err != nil {
		return nil, err
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(bts); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(bts); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
		return nil, errors.New("key is not an rsa private key")
	}

	return pkey, nil
}

// DecodeDEREncodedRSAPublicKey reads from the io.Reader until EOF.
// The resulting bytes are decoded as parsed as an rsa public key.
func DecodeDEREncodedRSAPublicKey(r io.Reader) (*rsa.PublicKey, error) {
	decoder := base64.NewDecoder(base64.StdEncoding, r)
	bts, err := ioutil.ReadAll(decoder)
	if err != nil {
		return nil, err
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PublicKey(bts); err != nil {
		if parsedKey, err = x509.ParsePKIXPublicKey(bts); err != nil {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, errors.New("key is not an rsa public key")
	}

	return pkey, nil
}
