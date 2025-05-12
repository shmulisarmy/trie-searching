package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func string_to_rune_list(word string) []rune {
	runes := []rune{}
	for _, char := range word {
		runes = append(runes, char)
	}
	return runes
}
func display_rune_list(runes []rune) {
	for _, r := range runes {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}

// func display_string_with_highlighted_letters(word string, index_of_letters_to_highlight []int) {
// 	for i, char := range word {
// 		if contains(index_of_letters_to_highlight, i) {
// 			fmt.Printf("%s%c%s", GREEN, char, RESET)
// 		} else {
// 			fmt.Printf("%c", char)
// 		}
// 	}
// 	fmt.Println()
// }

func last_in_list(list []int) int {
	if len(list) == 0 {
		panic("list is empty")
	}
	return list[len(list)-1]
}

func assert(condition bool) {
	if !condition {
		panic("assertion failed")
	}

}
func getInput() <-chan byte {
	inputChan := make(chan byte)

	go func() {
		defer close(inputChan)

		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}
		defer term.Restore(int(os.Stdin.Fd()), oldState)

		// Clear the screen and move cursor to top-left at the beginning
		fmt.Print("\033[2J\033[H")

		buf := make([]byte, 1)

		for {
			_, err := os.Stdin.Read(buf)
			if err != nil {
				panic(err)
			}

			inputChan <- buf[0]

			if buf[0] == 3 { // Ctrl+C
				return
			}
		}
	}()

	return inputChan
}

func sort[T any](slice []T, lessThan func(a, b T) bool) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if lessThan(slice[j], slice[i]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}
