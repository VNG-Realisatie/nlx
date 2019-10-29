// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"path/filepath"
)

func parseRSAPrivateKeyFile(keyFile string) (*rsa.PrivateKey, error) {
	pemBlock, err := parseFileToPem(keyFile)
	if err != nil {
		return nil, err
	}

	rsaKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	return rsaKey, nil
}

func parseRSAPublicKeyFile(keyFile string) (*rsa.PublicKey, error) {
	pemBlock, err := parseFileToPem(keyFile)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}

	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not a RSA key")
	}

	return rsaKey, nil
}

func parseFileToPem(keyFile string) (*pem.Block, error) {
	pemBlock, err := ioutil.ReadFile(filepath.Clean(keyFile))
	if err != nil {
		return nil, err
	}

	var block *pem.Block
	if block, _ = pem.Decode(pemBlock); block == nil {
		return nil, errors.New("key is not in PEM format")
	}

	return block, nil
}
