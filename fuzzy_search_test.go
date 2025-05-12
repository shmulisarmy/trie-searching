package main

import (
	"fmt"
	"testing"
)

func TestFuzzySearchListBasic(t *testing.T) {
	words := []string{"hello", "world", "help"}
	searchList := new_Fuzzy_Search_list(words, "")

	// Test initial state
	if len(searchList.words) != len(words) {
		t.Errorf("Initial state check failed:\n"+
			"Expected %d words in search list\n"+
			"Got %d words\n"+
			"Words in list: %v",
			len(words), len(searchList.words), getWordList(searchList))
	}
	if searchList.search_term != "" {
		t.Errorf("Initial search term check failed:\n"+
			"Expected empty search term\n"+
			"Got: '%s'",
			searchList.search_term)
	}

	// Test adding a word
	searchList.add_word("hello")
	expectedCount := len(words) + 1
	if len(searchList.words) != expectedCount {
		t.Errorf("Word addition check failed:\n"+
			"Expected %d words after adding 'hello'\n"+
			"Got %d words\n"+
			"Words in list: %v",
			expectedCount, len(searchList.words), getWordList(searchList))
	}

	// Test search term extension
	searchList.search_term_extend("hel")
	for _, word := range searchList.words {
		if word.word == "hello" || word.word == "help" {
			if word.upto_in_search_term != 3 {
				t.Errorf("Search term extension check failed for word '%s':\n"+
					"Expected upto_in_search_term to be 3\n"+
					"Got %d\n"+
					"Current match indexes: %v",
					word.word, word.upto_in_search_term, word.match_indexes)
			}
		}
	}

	// Test remove_char
	searchList.remove_char()
	if searchList.search_term != "he" {
		t.Errorf("Character removal check failed:\n"+
			"Expected search term to be 'he'\n"+
			"Got '%s'\n"+
			"Current word states:\n%s",
			searchList.search_term, getWordStates(searchList))
	}
}

func TestFuzzySearchListEdgeCases(t *testing.T) {
	// Test with empty word list
	searchList := new_Fuzzy_Search_list([]string{}, "")
	if len(searchList.words) != 0 {
		t.Errorf("Empty word list check failed:\n"+
			"Expected 0 words in search list\n"+
			"Got %d words\n"+
			"Words in list: %v",
			len(searchList.words), getWordList(searchList))
	}

	// Test with empty strings
	searchList = new_Fuzzy_Search_list([]string{""}, "")
	searchList.search_term_extend("a")
	if searchList.words[0].upto_in_search_term != 0 {
		t.Errorf("Empty string matching check failed:\n"+
			"Expected 0 matches for empty string\n"+
			"Got %d matches\n"+
			"Current word state: %s",
			searchList.words[0].upto_in_search_term, getWordState(searchList.words[0]))
	}

	// Test with special characters
	words := []string{"hello-world", "hello_world", "hello.world"}
	searchList = new_Fuzzy_Search_list(words, "")
	searchList.search_term_extend("hello-")
	if searchList.words[0].upto_in_search_term != 6 {
		t.Errorf("Special character matching check failed:\n"+
			"Expected 6 matches for 'hello-world'\n"+
			"Got %d matches\n"+
			"Match indexes: %v\n"+
			"Current word state: %s",
			searchList.words[0].upto_in_search_term, searchList.words[0].match_indexes,
			getWordState(searchList.words[0]))
	}
}

// Helper functions for better error messages
func getWordList(list *Fuzzy_Search_list) []string {
	words := make([]string, len(list.words))
	for i, w := range list.words {
		words[i] = w.word
	}
	return words
}

func getWordState(word *Word_search_state) string {
	return fmt.Sprintf("Word: '%s', Upto: %d, Matches: %v",
		word.word, word.upto_in_search_term, word.match_indexes)
}

func getWordStates(list *Fuzzy_Search_list) string {
	var states []string
	for _, word := range list.words {
		states = append(states, getWordState(word))
	}
	return fmt.Sprintf("%v", states)
}
