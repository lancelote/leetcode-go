package p380

import (
	"math/rand"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	rand.Seed(42)

	rs := Constructor()

	if rs.Insert(1) != true {
		t.Fatal("expect true on inserting first 1")
	}

	if rs.Remove(2) != false {
		t.Fatal("expect false on removing non-existing 2")
	}

	if rs.Insert(2) != true {
		t.Fatal("expect true on inserting first 2")
	}

	val := rs.GetRandom()
	if val != 1 && val != 2 {
		t.Fatal("first random element is not 1 or 2")
	}

	if rs.Remove(1) != true {
		t.Fatal("expect true on removing 1")
	}

	if rs.Insert(2) != false {
		t.Fatal("expect false on inserting second 2")
	}

	if rs.GetRandom() != 2 {
		t.Fatal("second randome element is not 2")
	}
}

func TestRemoveLast(t *testing.T) {
	rs := Constructor()
	rs.Insert(1)
	rs.Remove(1)

	if rs.Insert(1) != true {
		t.Fatal("expect true on re-inserting 1")
	}
}
