package hash

import (
	"fmt"
	"testing"
)

func TestHasher(t *testing.T) {
	var hash Hasher
	hash = NewOnline(1, 1)
	_ = fmt.Sprintf("(%v)", hash)
}
