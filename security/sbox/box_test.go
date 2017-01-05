package sbox

import "testing"

func TestBox(t *testing.T) {

	m, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("m = %q : %q", mustPublicKey(t, m), m)

	a, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("a = %q : %q", mustPublicKey(t, a), a)

	b, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("b = %q : %q", mustPublicKey(t, b), b)

	c, err := GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("c = %q : %q", mustPublicKey(t, c), c)

	box, err := Seal([]byte("Hello there"), m, []string{
		mustPublicKey(t, a),
		mustPublicKey(t, b),
		mustPublicKey(t, c),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("len(box) = %d", len(box))
	t.Logf("box =\n%s", box)

	msg, err := Open(box, a, mustPublicKey(t, m))
	if err != nil {
		t.Fatal(err)
	}
	if string(msg) != "Hello there" {
		t.Fatalf("failed to decode box: %q", msg)
	}

	msg, err = Open(box, b, mustPublicKey(t, m))
	if err != nil {
		t.Fatal(err)
	}
	if string(msg) != "Hello there" {
		t.Fatalf("failed to decode box: %q", msg)
	}

	msg, err = Open(box, c, mustPublicKey(t, m))
	if err != nil {
		t.Fatal(err)
	}
	if string(msg) != "Hello there" {
		t.Fatalf("failed to decode box: %q", msg)
	}
}

func mustPublicKey(t *testing.T, s string) string {
	p, err := PublicKey(s)
	if err != nil {
		t.Fatal(err)
	}
	return p
}
