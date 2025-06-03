package encryption

import (
	"fmt"
	"testing"
)

func TestDict(t *testing.T) {

	s := SaltBcrypt("123")
	fmt.Printf("TestDict result:%v", s)
}
