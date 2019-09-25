package recrypt

import (
	"bytes"
	"fmt"
	wasm "github.com/wasmerio/go-ext-wasm/wasmer"
	"reflect"
)

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

func doop() {
	// Reads the WebAssembly module as bytes.
	bytes, _ := wasm.ReadBytes("recrypt_wasm_binding.wasm")

	// Instantiates the WebAssembly module.
	instance, _ := wasm.NewInstance(bytes)
	defer instance.Close()

	// for k := range instance.Exports {
	// 	// fmt.Printf("key[%s] value[%s]\n", k, v)
	// 	fmt.Printf("key[%s]", k)
	// }

	fmt.Println(reflect.ValueOf(instance.Exports).MapKeys())

	// fmt.Printf("%v", instance)

	// result, err := instance.Exports["Api256"]
	//
	// fmt.Println(err)
	// fmt.Println(result)

	// api, _ := Api256()

	// return api

	// fmt.Println(err)
	// fmt.Println(api)

	// fmt.Printf("%T", Api256)
	// fmt.Printf("%T", api)

	// Api256 := instance.Exports["Api256"]
	// api, _ := Api256()
	//
	// fmt.Println(api)

	// // Gets the `sum` exported function from the WebAssembly instance.
	// sum := instance.Exports["sum"]
	//
	// // Calls that exported function with Go standard values. The WebAssembly
	// // types are inferred and values are casted automatically.
	// result, _ := sum(5, 37)
	//
	// fmt.Println(result) // 42!
}

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
