package p374

import (
	"testing"
)

func guess(num int) int {
	if num == 6 {
		return 0
	} else if num > 6 {
		return -1
	} else {
		return +1
	}
}

func Test_guessNumber(t *testing.T) {
	if got := guessNumber(10); got != 6 {
		t.Errorf("want 6, got %d", got)
	}
}
