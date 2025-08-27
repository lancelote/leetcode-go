package p208

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Fatalf("cannot find %q", "apple")
	}

	if trie.Search("app") {
		t.Fatalf("can find %q", "app")
	}

	if !trie.StartsWith("app") {
		t.Fatalf("cannot find prefix %q", "app")
	}

	trie.Insert("app")

	if !trie.Search("app") {
		t.Fatalf("cannot find %q", "app")
	}
}
