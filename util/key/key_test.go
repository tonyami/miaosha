package key

import "testing"

func TestCreate(t *testing.T) {
	token := create(All, 64)
	if len(token) != 64 {
		t.Fatal()
	}
}
