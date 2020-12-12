package code

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	for k, v := range codes {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
}
