package key

import (
	"testing"
)

func TestToken(t *testing.T) {
	t.Log(Token())
	t.Log(Token())
	t.Log(Token())
}

func TestSmsCode(t *testing.T) {
	t.Log(SmsCode())
	t.Log(SmsCode())
	t.Log(SmsCode())
}
