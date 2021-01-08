package key

import (
	"testing"
)

func TestToken(t *testing.T) {
	if token := Token(); len(token) != 64 {
		t.Fatal()
	}
}
