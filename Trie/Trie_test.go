package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	// Test Insert
	trie.Insert("hello")
	trie.Insert("world")
	trie.Insert("hey")
	trie.Insert("hi")
	trie.Insert("howdy")

	// Test Search
	if !trie.Search("hello") {
		t.Errorf("Expected 'hello' to be found in Trie")
	}
	if !trie.Search("world") {
		t.Errorf("Expected 'world' to be found in Trie")
	}
	if !trie.Search("hey") {
		t.Errorf("Expected 'hey' to be found in Trie")
	}
	if !trie.Search("hi") {
		t.Errorf("Expected 'hi' to be found in Trie")
	}
	if !trie.Search("howdy") {
		t.Errorf("Expected 'howdy' to be found in Trie")
	}
	if trie.Search("notfound") {
		t.Errorf("Expected 'notfound' to not be found in Trie")
	}

	// Test Delete
	trie.Delete("hello")
	if trie.Search("hello") {
		t.Errorf("Expected 'hello' to be deleted from Trie")
	}

	// Test PrefixSearch
	prefixResults := trie.PrefixSearch("h")
	expectedPrefixResults := []string{"hey", "hi", "howdy"}
	if len(prefixResults) != len(expectedPrefixResults) {
		t.Errorf("Expected %d results, but got %d", len(expectedPrefixResults), len(prefixResults))
	}
	for i, word := range expectedPrefixResults {
		if prefixResults[i] != word {
			t.Errorf("Expected '%s', but got '%s'", word, prefixResults[i])
		}
	}
}
