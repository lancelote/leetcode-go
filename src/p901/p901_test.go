package p901

import (
	"testing"
)

func TestStockSpanner(t *testing.T) {
	s := Constructor()

	assertNext := func(p int, want int) {
		got := s.Next(p)
		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	}

	assertNext(100, 1)
	assertNext(80, 1)
	assertNext(60, 1)
	assertNext(70, 2)
	assertNext(60, 1)
	assertNext(75, 4)
	assertNext(85, 6)
}
