package recrypt

import "bytes"

type PublicKey bytes.Buffer
type PrivateKey bytes.Buffer
type PublicSigningKey bytes.Buffer
type PrivateSigningKey bytes.Buffer
type Plaintext bytes.Buffer

type KeyPair struct {
	publicKey  PublicKey
	privateKey PrivateKey
}

type SigningKeyPair struct {
	publicKey  PublicSigningKey
	privateKey PrivateSigningKey
}

type TransformBlock struct {
	publicKey                       PublicKey
	encryptedTempKey                bytes.Buffer
	randomTransformPublicKey        PublicKey
	randomTransformEncryptedTempKey bytes.Buffer
}

type EncryptedValue struct {
	ephemeralPublicKey PublicKey
	encryptedMessage   bytes.Buffer
	authHash           bytes.Buffer
	transformBlocks    []TransformBlock
	publicSigningKey   PublicSigningKey
	signature          bytes.Buffer
}

type Api256 struct{}

// -----------------------------------------------------------------------------

func NewApi256() Api256 {
	return Api256{}
}

func (a Api256) generateKeyPair() KeyPair {
	return KeyPair{}
}

func (a Api256) generateEd25519KeyPair() SigningKeyPair {
	return SigningKeyPair{}
}

func (a Api256) generatePlaintext() Plaintext {
	return Plaintext{}
}

func (a Api256) encrypt(plaintext Plaintext, toPublicKey PublicKey, privateSigningKey PrivateSigningKey) EncryptedValue {
	return EncryptedValue{}
}

func (a Api256) decrypt(encryptedValue EncryptedValue, privateKey PrivateKey) EncryptedValue {
	return EncryptedValue{}
}
