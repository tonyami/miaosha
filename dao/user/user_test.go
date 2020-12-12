package user

import (
	"testing"
)

func TestDao_QueryById(t *testing.T) {
	user, err := d.QueryById(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", user)
}
