package trie_searching

import "fmt"

type Node struct {
	children map[rune]*Node
	is_word  bool
	letter   rune
}

type Trie struct {
	root *Node
}

func (t *Trie) insert(word string) {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			current.children[char] = &Node{
				children: make(map[rune]*Node),
				letter:   char,
			}
		}
		current = current.children[char]
	}
	current.is_word = true
}

func (t *Trie) contains_word(word string) bool {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}
	return current.is_word
}

func display_rune_list(runes []rune) {
	for _, r := range runes {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}

func word_to_rune_list(word string) []rune {
	runes := []rune{}
	for _, char := range word {
		runes = append(runes, char)
	}
	return runes
}

func (current *Node) is_child_of(other_node *Node) bool {
	return other_node.children[current.letter] == current
}

func (current *Node) bfs(prefix string) []string {
	current_letters := word_to_rune_list(prefix)
	results := []string{}

	queue := []*Node{current}
	for len(queue) > 0 {
		next_one := queue[0]
		next_ones_letter := next_one.letter
		if !next_one.is_child_of(current) {
			current_letters = current_letters[:len(current_letters)-1]
		}
		current_letters = append(current_letters, next_ones_letter)

		current = next_one
		queue = queue[1:]
		if current.is_word {
			results = append(results, string(current_letters))
		}
		for _, child := range current.children {
			queue = append(queue, child)
		}
	}
	return results
}

func (t *Trie) contains_prefix(word string) bool {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return false
		}
		current = current.children[char]
	}
	return true
}

func (t *Trie) search(word string) []string {
	current := t.root
	for _, char := range word {
		if _, exists := current.children[char]; !exists {
			return []string{}
		}
		current = current.children[char]
	}
	return current.bfs(word)
}

func (t *Trie) human_search(word string, max_mistakes int) []string {
	mistakes_made_so_far := 0
	current := t.root
	used_prefix := []rune{}
	for _, char := range word {
		if _, exists := current.children[char]; exists {
			used_prefix = append(used_prefix, char)
			current = current.children[char]
		} else {
			mistakes_made_so_far++
			if mistakes_made_so_far > max_mistakes {
				return []string{}
			}
		}
	}
	return current.bfs(string(used_prefix))
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[rune]*Node),
		},
	}
}

func main() {
	t := NewTrie()
	t.insert("apple")
	t.insert("application")
	t.insert("app")
	fmt.Println(t.contains_word("apple"))
	fmt.Println(t.contains_word("application"))
	fmt.Println(t.contains_word("app"))
	fmt.Println(t.contains_prefix("app"))
	for _, word := range t.human_search("apfp", 1) {
		fmt.Println(word)
	}
}
