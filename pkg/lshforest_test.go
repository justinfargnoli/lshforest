package lshforest

import "testing"

func TestNewDefault(t *testing.T) {
	_ = NewDefault(300, Cosine)
	_ = NewDefault(500, Cosine)
	_ = NewDefault(100, Cosine)
}

func TestNewDefaultPanicJaccard(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.FailNow()
		}
	}()
	_ = NewDefault(1, 100)
}

func TestInsert(t *testing.T) {
	lshforest := NewDefault(3, Cosine)

	if lshforest.Insert(&[]float64{0, 0, 0}, nil) == nil {
		t.FailNow()
	}
	if lshforest.Insert(&[]float64{0, 0, 1, 3}, nil) == nil {
		t.FailNow()
	}

	if err := lshforest.Insert(&[]float64{1, 2, 3}, nil); err != nil {
		t.Fatal(err)
	}
	if err := lshforest.Insert(&[]float64{1.1, 4, -3}, nil); err != nil {
		t.Fatal(err)
	}
	if err := lshforest.Insert(&[]float64{1, -2, 3}, nil); err != nil {
		t.Fatal(err)
	}
	if err := lshforest.Insert(&[]float64{1, 1, 1}, nil); err != nil {
		t.Fatal(err)
	}
	if err := lshforest.Insert(&[]float64{-1, 2, 3}, nil); err != nil {
		t.Fatal(err)
	}
	if err := lshforest.Insert(&[]float64{0.1, 0.2, 0.3}, nil); err != nil {
		t.Fatal(err)
	}
}

func TestInsertAll(t *testing.T) {
	vectors := [][]float64{
		{1, 2, 3},
		{1.1, 4, -3},
		{1, -2, 3},
		{1, 1, 1},
		{-1, 2, 3},
		{0.1, 0.2, 0.3},
	}
	values := []interface{}{nil, nil, nil, nil, nil, nil}
	lshforest := NewDefault(3, Cosine)
	if err := lshforest.InsertAll(&vectors, &values); err != nil {
		t.Fatal(err)
	}
}

func insertAll(t *testing.T) *LSHForest {
	vectors := [][]float64{
		{1, 2, 3},
		{1.1, 4, -3},
		{1, -2, 3},
		{1, 1, 1},
		{-1, 2, 3},
		{0.1, 0.2, 0.4},
	}
	values := []interface{}{0, 1, 2, 3, 4, 5}
	lshforest := NewDefault(3, Cosine)
	if err := lshforest.InsertAll(&vectors, &values); err != nil {
		t.Fatal(err)
	}
	return lshforest
}

func TestQuery(t *testing.T) {
	lshforest := insertAll(t)

	if _, err := lshforest.Query(&[]float64{1}, 5); err != ErrEqDim {
		t.Fatal(err)
	}
	if _, err := lshforest.Query(&[]float64{0, 0, 0}, 5); err != ErrNonZero {
		t.Fatal(err)
	}

	value, err := lshforest.Query(&[]float64{1, 1, 1}, 6)
	if err != nil {
		t.Fatal(err)
	}
	if (*value)[0].(int) != 3 {
		t.Fatalf("Expected \"a\" | got (%v)", (*value)[0].(string))
	}
	value, err = lshforest.Query(&[]float64{0.1, 0.2, 0.4}, 6)
	if err != nil {
		t.Fatal(err)
	}
	if (*value)[0].(int) != 5 {
		t.Fatalf("Expected \"a\" | got (%v)", (*value)[0].(string))
	}
	value, err = lshforest.Query(&[]float64{1, 2, 3}, 1)
	if err != nil {
		t.Fatal(err)
	}
	if (*value)[0].(int) != 0 {
		t.Fatalf("Expected \"a\" | got (%v)", (*value)[0].(string))
	}
}
