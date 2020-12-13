package conf

import "testing"

func TestInit(t *testing.T) {
	if err := Init("test"); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", conf.Server)
	t.Logf("%#v", conf.Mysql)
}
