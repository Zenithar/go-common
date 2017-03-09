package argon2

import "github.com/lhecker/argon2"

// Hash password using Argon 2 algorithm
func Hash(pwd []byte) ([]byte, error) {
	// Initialize Argon 2 default config
	cfg := argon2.DefaultConfig()

	// Return hash encoded argon2i
	return cfg.HashEncoded([]byte(pwd))
}

// Verify hash password
func Verify(encoded, password []byte) (bool, error) {

	// Decode encoded hash structure
	raw, err := argon2.Decode(encoded)
	if err != nil {
		return false, err
	}

	// The Raw struct then allows us to verify it against a unencrypted password.
	return raw.Verify(password)
}
