package sbox

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io"
	"sort"

	"golang.org/x/crypto/nacl/box"
)

// Seal a container
func Seal(data []byte, masterKey string, receiverKeys []string) ([]byte, error) {
	_, masterPrivateKey, err := parseKey(masterKey)
	if err != nil {
		return nil, err
	}

	sort.Strings(receiverKeys)
	receiverPublicKeys := make([]*[32]byte, 0, len(receiverKeys))
	for _, rk := range receiverKeys {
		pub, err2 := parsePublicKey(rk)
		if err2 != nil {
			return nil, err2
		}

		receiverPublicKeys = append(receiverPublicKeys, pub)
	}

	b, err := sealBox(data, receiverPublicKeys, masterPrivateKey)
	if err != nil {
		return nil, err
	}

	encdata, err := b.Marshal()
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = pem.Encode(&buf, &pem.Block{
		Type:  "SEC BOX",
		Bytes: encdata,
	})
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Open secure box
func Open(data []byte, receiverKey, masterKey string) ([]byte, error) {
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "SEC BOX" {
		return nil, errors.New("invalid pem block")
	}

	masterPublicKey, err := parsePublicKey(masterKey)
	if err != nil {
		return nil, err
	}

	receiverPublicKey, receiverPrivateKey, err := parseKey(receiverKey)
	if err != nil {
		return nil, err
	}

	var b Box
	err = b.Unmarshal(block.Bytes)
	if err != nil {
		return nil, err
	}

	decdata, ok := openBox(&b, masterPublicKey, receiverPublicKey, receiverPrivateKey)
	if !ok {
		return nil, errors.New("failed to decrypt box")
	}

	return decdata, nil
}

// GenerateKey for box
func GenerateKey() (string, error) {
	pub, prv, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return "", errInvalidRandom
	}

	var kp [64]byte
	copy(kp[:32], pub[:])
	copy(kp[32:], prv[:])

	return base64.RawURLEncoding.EncodeToString(kp[:]), nil
}

func PublicKey(privateKey string) (string, error) {
	pub, _, err := parseKey(privateKey)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(pub[:]), nil
}

// -----------------------------------------------------------------------------

func parsePublicKey(s string) (pub *[32]byte, err error) {
	kp, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil || len(kp) != 32 {
		return nil, errInvalidKey
	}

	var (
		pubKey = new([32]byte)
	)

	copy(pubKey[:], kp)

	return pubKey, nil
}

func parseKey(s string) (pub, prv *[32]byte, err error) {
	kp, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil || len(kp) != 64 {
		return nil, nil, errInvalidKey
	}

	var (
		pubKey = new([32]byte)
		prvKey = new([32]byte)
	)

	copy(pubKey[:], kp[:32])
	copy(prvKey[:], kp[32:])

	return pubKey, prvKey, nil
}

func sealBox(message []byte, receiverPublicKeys []*[32]byte, masterPrivateKey *[32]byte) (*Box, error) {
	boxPublicKey, boxPrivateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, errInvalidRandom
	}

	var b Box

	payload, err := sealPayload(message, boxPublicKey, masterPrivateKey)
	if err != nil {
		return nil, err
	}
	b.Payload = payload

	for _, receiverPublicKey := range receiverPublicKeys {
		key, err := sealBoxKey(boxPrivateKey, receiverPublicKey, masterPrivateKey)
		if err != nil {
			return nil, err
		}

		b.Keys = append(b.Keys, key)
	}

	return &b, nil
}

func openBox(b *Box, masterPublicKey, receiverPublicKey, receiverPrivateKey *[32]byte) ([]byte, bool) {
	var boxPrivateKey [32]byte

	for _, k := range b.Keys {
		if !bytes.Equal(k.PublicKey, receiverPublicKey[:]) {
			continue
		}

		if !openBoxKey(&boxPrivateKey, k, masterPublicKey, receiverPrivateKey) {
			continue
		}

		return openPayload(b.Payload, masterPublicKey, &boxPrivateKey)
	}

	return nil, false
}

func sealPayload(message []byte, boxPublicKey, masterPrivateKey *[32]byte) (*Payload, error) {
	var (
		payload Payload
		nonce   = new([24]byte)
	)

	err := readRand(nonce[:])
	if err != nil {
		return nil, err
	}

	payload.Nonce = nonce[:]
	payload.Data = box.Seal(nil, message, nonce, boxPublicKey, masterPrivateKey)
	return &payload, nil
}

func openPayload(payload *Payload, masterPublicKey, boxPrivateKey *[32]byte) ([]byte, bool) {
	var (
		nonce [24]byte
	)

	copy(nonce[:], payload.Nonce)
	return box.Open(nil, payload.Data, &nonce, masterPublicKey, boxPrivateKey)
}

func sealBoxKey(boxPrivateKey, receiverPublicKey, masterPrivateKey *[32]byte) (*Key, error) {
	var (
		key   Key
		nonce = new([24]byte)
	)

	err := readRand(nonce[:])
	if err != nil {
		return nil, err
	}

	key.Nonce = nonce[:]
	key.PublicKey = append([]byte(nil), receiverPublicKey[:]...)
	key.BoxKey = box.Seal(nil, boxPrivateKey[:], nonce, receiverPublicKey, masterPrivateKey)
	return &key, nil
}

func openBoxKey(out *[32]byte, key *Key, masterPublicKey, receiverPrivateKey *[32]byte) bool {
	var (
		nonce [24]byte
	)

	copy(nonce[:], key.Nonce)
	data, ok := box.Open(nil, key.BoxKey, &nonce, masterPublicKey, receiverPrivateKey)
	copy(out[:], data)
	return ok
}

var errInvalidKey = errors.New("invalid key")
var errInvalidRandom = errors.New("failed to read random bytes")

func readRand(d []byte) error {
	n, err := io.ReadFull(rand.Reader, d)
	if err != nil {
		return errInvalidRandom
	}
	if n != len(d) {
		return errInvalidRandom
	}
	return nil
}
