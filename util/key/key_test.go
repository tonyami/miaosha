package key

import (
	"testing"
)

func TestCreateToken(t *testing.T) {
	if token := CreateToken(); len(token) != 64 {
		t.Fatal()
	}
}
