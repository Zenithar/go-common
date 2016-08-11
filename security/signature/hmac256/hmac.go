package hmac256

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
)

// GetSignature returns an HMAC-SHA256 signature encoded in base64
func GetSignature(secret []byte, messages ...[]byte) string {
	h := hmac.New(sha256.New, secret)
	for _, m := range messages {
		h.Write(m)
	}
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

// CompareSignatures is used to compare given and reference signature using time constant algorithm
func CompareSignatures(given string, reference string) (bool, error) {
	h1, err := base64.URLEncoding.DecodeString(given)
	if err != nil {
		return false, err
	}

	h2, err := base64.URLEncoding.DecodeString(reference)
	if err != nil {
		return false, err
	}

	return subtle.ConstantTimeCompare(h1, h2) == 1, nil
}
