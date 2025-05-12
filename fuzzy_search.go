package main

import (
	"fmt"
)

type Color string

const (
	RED   Color = "\033[31m"
	GREEN Color = "\033[32m"
	RESET Color = "\033[0m"
)

type Word_search_state struct {
	upto_in_search_term int
	match_indexes       []int
	word                string
}

type Fuzzy_Search_list struct {
	words       []*Word_search_state
	search_term string
}

func (word *Word_search_state) SearchAndHighlight(full_search_term string, previous_letter_amount int) {
	for letter_index, char := range word.word[previous_letter_amount:] {
		if word.upto_in_search_term < len(full_search_term) && full_search_term[word.upto_in_search_term] == byte(char) {
			word.match_indexes = append(word.match_indexes, letter_index+previous_letter_amount)
			word.upto_in_search_term++
		}
	}
	display_string_with_highlighted_letters(word.word, word.match_indexes)
}

func (word *Word_search_state) on_remove_char(letter_removed byte) {
	if len(word.match_indexes) == 0 {
		assert(word.upto_in_search_term == 0)
		return
	}
	// fmt.Printf("last_in_list(word.match_indexes): %d\n", last_in_list(word.match_indexes))
	// fmt.Printf("letter_removed: %d\n", letter_removed)

	if letter_removed == word.word[last_in_list(word.match_indexes)] {
		word.match_indexes = word.match_indexes[:len(word.match_indexes)-1]
		word.upto_in_search_term--
	}
}

func (word *Word_search_state) on_search_term_extended(full_search_term string) {
	fmt.Print("word.match_indexes\r\n")
	fmt.Printf("%v \r\n", word.match_indexes)
	word.SearchAndHighlight(full_search_term, len(word.match_indexes))
}

func (fuzzy_Search_list *Fuzzy_Search_list) search_term_extend(add_on_to_search_term string) {
	fuzzy_Search_list.search_term += add_on_to_search_term
	fmt.Println("on_search_term_extended on Fuzzy_Search_list called with add_on_to_search_term:", add_on_to_search_term)
	for _, word := range fuzzy_Search_list.words {
		word.on_search_term_extended(fuzzy_Search_list.search_term)
	}
}

func (s *Fuzzy_Search_list) SearchAndHighlight() {
	for _, word := range s.words {
		word.SearchAndHighlight(s.search_term, 0)
	}
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (fuzzy_Search_list *Fuzzy_Search_list) add_word(word string) {
	word_state := &Word_search_state{
		upto_in_search_term: 0,
		match_indexes:       []int{},
		word:                word,
	}
	word_state.SearchAndHighlight(fuzzy_Search_list.search_term, 0)
	fuzzy_Search_list.words = append(fuzzy_Search_list.words, word_state)
}

func new_Fuzzy_Search_list(words []string, search_term string) *Fuzzy_Search_list {
	result := &Fuzzy_Search_list{
		words:       []*Word_search_state{},
		search_term: search_term,
	}
	for _, word := range words {
		word_state := &Word_search_state{
			upto_in_search_term: 0,
			match_indexes:       []int{},
			word:                word,
		}
		result.words = append(result.words, word_state)
	}
	result.SearchAndHighlight()
	return result
}

func display_string_with_highlighted_letters(word string, match_indexes []int) {
	for i, char := range word {
		if contains(match_indexes, i) {
			fmt.Printf("%s%c%s", GREEN, char, RESET)
		} else {
			fmt.Printf("%c", char)
		}
	}
	fmt.Print("\r\n")
}

func (fuzzy_Search_list *Fuzzy_Search_list) remove_char() {
	last_index := len(fuzzy_Search_list.search_term)
	for _, word := range fuzzy_Search_list.words {
		word.on_remove_char(fuzzy_Search_list.search_term[last_index-1])
	}
	fuzzy_Search_list.search_term = fuzzy_Search_list.search_term[:len(fuzzy_Search_list.search_term)-1]
	sort(fuzzy_Search_list.words, func(a, b *Word_search_state) bool {
		return a.upto_in_search_term < b.upto_in_search_term
	})
	for _, word := range fuzzy_Search_list.words {
		word.on_search_term_extended(fuzzy_Search_list.search_term)
	}
}
