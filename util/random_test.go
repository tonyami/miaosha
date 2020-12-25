package util

import "testing"

func TestCreateToken(t *testing.T) {
	t.Logf("%s\n", CreateToken())
}

func TestCreateSalt(t *testing.T) {
	t.Logf("%s\n", CreateSalt())
}
