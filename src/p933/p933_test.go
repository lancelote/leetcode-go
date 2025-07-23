package p933

import (
	"testing"
)

func Test_Ping(t *testing.T) {
	rc := Constructor()

	if pong1 := rc.Ping(1); pong1 != 1 {
		t.Fatalf("pong1: want 1, got %d", pong1)
	}

	if pong2 := rc.Ping(100); pong2 != 2 {
		t.Fatalf("pong2: want 2, got %d", pong2)
	}

	if pong3 := rc.Ping(3001); pong3 != 3 {
		t.Fatalf("pong3: want 3, got %d", pong3)
	}

	if pong4 := rc.Ping(3002); pong4 != 3 {
		t.Fatalf("pong4: want 3, got %d", pong4)
	}
}
