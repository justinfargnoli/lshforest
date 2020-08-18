package hash

import "testing"

func TestOnlineSimhash(t *testing.T) {
	vectors := [][]float64{
		{0.0, 2.0, 3.3, -4.2},
		{0.0, -2.0, 3.3, -4.2},
		{0.0, 2.0, -3.3, -4.2},
	}
	online := NewOnline(300, uint(len(vectors[0])))
	for _, v := range vectors {
		online.Hash(v)
	}
}

func TestOffline(t *testing.T) {
	vectors := [][]float64{
		{0.0, 2.0, 3.3, -4.2},
		{0.0, -2.0, 3.3, -4.2},
		{0.0, 2.0, -3.3, -4.2},
	}
	Offline(vectors, 300)
}
