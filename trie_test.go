package main

import (
	"testing"
)

func TestTrieBasicOperations(t *testing.T) {
	trie := NewTrie()
	word := "apple"

	// Test empty trie
	if trie.contains_word(word) {
		t.Errorf("Empty trie should not contain word %s as it wasn't inserted yet", word)
	}

	// Test insertion and word existence
	trie.insert(word)
	if !trie.contains_word(word) {
		t.Errorf("Trie should contain word %s after insertion", word)
	}

	// Test all prefixes
	prefixes := []string{"a", "ap", "app", "appl", "apple"}
	for _, prefix := range prefixes {
		if !trie.contains_prefix(prefix) {
			t.Errorf("Trie should contain prefix %s for word %s", prefix, word)
		}
	}

}

func TestTrieMultipleWords(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "application", "app"}

	// Insert multiple words
	for _, word := range words {
		trie.insert(word)
	}

	// Test each word exists
	for _, word := range words {
		if !trie.contains_word(word) {
			t.Errorf("Trie should contain word %s", word)
		}
	}

	// Test common prefix
	if !trie.contains_prefix("app") {
		t.Errorf("Trie should contain prefix 'app'")
	}

	// Test search functionality
	results := trie.search("app")
	expectedCount := 3
	if len(results) != expectedCount {
		t.Errorf("Search for 'app' should return %d results, got %d", expectedCount, len(results))
	}

	// Verify all expected words are in results
	expectedWords := map[string]bool{
		"app":         true,
		"apple":       true,
		"application": true,
	}
	for _, result := range results {
		if !expectedWords[result] {
			t.Errorf("Unexpected word in search results: %s", result)
		}
	}
}

func TestTrieFuzzySearch(t *testing.T) {
	trie := NewTrie()
	words := []string{"apple", "application", "app"}

	// Insert words
	for _, word := range words {
		trie.insert(word)
	}

	// Test fuzzy search with one mistake
	results := trie.search_while_discarding_extra_letters("apfple", 1)
	if len(results) == 0 {
		t.Errorf("search_while_discarding_extra_letters with one extra letter should find 'apple'")
	}

	if len(trie.search_while_discarding_extra_letters("apffple", 1)) != 0 {
		t.Errorf("search_while_discarding_extra_letters with 2 extra letter should not find 'apple'")
	}

	// Test fuzzy search with no matches
	results = trie.search_while_discarding_extra_letters("xyz", 1)
	if len(results) != 0 {
		t.Errorf("Fuzzy search should not find matches for 'xyz'")
	}

	// Test fuzzy search with multiple possible matches
	results = trie.search_while_discarding_extra_letters("app", 1)
	if len(results) != 3 {
		t.Errorf("Fuzzy search for 'app' should return 3 results, got %d", len(results))
	}
}

func TestTrieEmptyAndEdgeCases(t *testing.T) {
	trie := NewTrie()

	// Test empty string
	trie.insert("")
	if !trie.contains_word("") {
		t.Errorf("Trie should contain empty string after insertion")
	}

	// Test search with empty string
	results := trie.search("")
	if len(results) != 1 {
		t.Errorf("Search with empty string should return 1 result, got %d", len(results))
	}

	// Test contains_prefix with empty string
	if !trie.contains_prefix("") {
		t.Errorf("Trie should contain empty string as prefix")
	}

	// Test human_search with empty string
	results = trie.search_while_discarding_extra_letters("", 1)
	if len(results) != 1 {
		t.Errorf("Human search with empty string should return 1 result, got %d", len(results))
	}
}
