package p2336

import (
	"testing"
)

func TestSmallestInfiniteSet(t *testing.T) {
	set := Constructor()
	set.AddBack(2)

	if x := set.PopSmallest(); x != 1 {
		t.Errorf("want 1, got %d", x)
	}

	if x := set.PopSmallest(); x != 2 {
		t.Errorf("want 2, got %d", x)
	}

	if x := set.PopSmallest(); x != 3 {
		t.Errorf("want 3, got %d", x)
	}

	set.AddBack(1)

	if x := set.PopSmallest(); x != 1 {
		t.Errorf("want 1, got %d", x)
	}

	if x := set.PopSmallest(); x != 4 {
		t.Errorf("want 4, got %d", x)
	}

	if x := set.PopSmallest(); x != 5 {
		t.Errorf("want 5, got %d", x)
	}
}

func Test_addSame(t *testing.T) {
	set := Constructor()

	set.PopSmallest()
	set.PopSmallest()
	set.PopSmallest()

	set.AddBack(1)
	set.AddBack(1)

	if x := set.PopSmallest(); x != 1 {
		t.Errorf("want 1, got %d", x)
	}

	if x := set.PopSmallest(); x != 4 {
		t.Errorf("want 4, got %d", x)
	}
}
