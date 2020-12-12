package user

import "testing"

func TestService_GetInfo(t *testing.T) {
	user, err := s.GetInfo(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", user)
}
