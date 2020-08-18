package hash

import (
	"testing"
	"fmt"
)


func TestHasher(t *testing.T) {
	var hash Hasher
	hash = NewOnline(1, 1)	
	_ = fmt.Sprintf("(%v)", hash)
}