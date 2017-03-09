package argon2

import "testing"

var password = []byte("password")

func TestArgon2Hash(t *testing.T) {
	encoded, err := Hash(password)
	if err != nil {
		t.Fail()
	}

	ok, err := Verify(encoded, password)
	if err != nil {
		t.Fail()
	}

	if !ok {
		t.Fail()
	}
}

func TestBadArgon2Hash(t *testing.T) {
	encoded := []byte("")

	ok, err := Verify(encoded, password)
	if err == nil {
		t.Fail()
	}

	if ok {
		t.Fail()
	}
}
