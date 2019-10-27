package recrypt

/*
#cgo LDFLAGS: -L./lib -lhello
#include "./lib/hello.h"
*/
import "C"

import "bytes"
import "fmt"
import "unsafe"

// type PublicKey bytes.Buffer
type PrivateKey []byte
type PublicSigningKey []byte
type PrivateSigningKey []byte
type Plaintext []byte

type PublicKey struct {
  x []byte
  y []byte
}

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
	keypair_pt := C.generate_keypair()

  publicKey := PublicKey {
      x: C.GoBytes(unsafe.Pointer(&keypair_pt.public_key.x), 32),
      y: C.GoBytes(unsafe.Pointer(&keypair_pt.public_key.y), 32),
  }

  privateKey := C.GoBytes(unsafe.Pointer(&keypair_pt.private_key), 32)

	return KeyPair {
		publicKey,
		privateKey,
	}
}

func (a Api256) generateEd25519KeyPair() SigningKeyPair {
  keypair_pt := C.generate_ed25519_keypair()

  publicKey := C.GoBytes(unsafe.Pointer(&keypair_pt.public_key), 32)
  privateKey := C.GoBytes(unsafe.Pointer(&keypair_pt.private_key), 64)

  return SigningKeyPair{
    publicKey,
    privateKey,
  }
}

func (a Api256) generatePlaintext() Plaintext {
  plaintext_pt := C.generate_plaintext()
  plaintext := C.GoBytes(unsafe.Pointer(&plaintext_pt.bytes), 384)
	return plaintext
}

func (a Api256) encrypt(plaintext Plaintext, toPublicKey PublicKey, privateSigningKey PrivateSigningKey) EncryptedValue {
	return EncryptedValue{}
}

func (a Api256) decrypt(encryptedValue EncryptedValue, privateKey PrivateKey) EncryptedValue {
	return EncryptedValue{}
}

// -----------------------------------------------------------------------------

func doThatRustThing() {
	C.hello(C.CString("Gopher"))
}

func proofOfConcept() {
	val, _ := C.get_proof_of_concept_string()
	fmt.Println(C.GoString(val))
}
