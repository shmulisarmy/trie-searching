package main

func main() {

	words := []string{"hello", "hell", "heaven", "heavy", "hero", "puppy"}
	search_term := "h"
	search := new_Fuzzy_Search_list(words, search_term)

	for char := range getInput() {
		if char == 127 {
			search.remove_char()
		} else {
			search.search_term_extend(string(char))
		}
	}
}
