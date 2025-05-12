# Trie Search Library

A powerful and efficient Go library for implementing trie-based search functionality with fuzzy matching capabilities.

## Overview

This library provides a robust implementation of a trie data structure with advanced search features. It's designed to handle efficient string operations, prefix matching, and fuzzy searching with customizable error tolerance.

### Key Features

- **Trie Data Structure**: Efficient prefix tree implementation for fast string operations
- **Fuzzy Search**: Support for approximate string matching with configurable error tolerance
- **Prefix Matching**: Quick prefix-based search operations
- **Interactive Search**: Real-time search with highlighted matches
- **Memory Efficient**: Optimized for memory usage while maintaining performance

## Installation

1. Ensure you have Go 1.23.2 or later installed:
   ```bash
   go version
   ```

2. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/trie-search.git
   cd trie-search
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

## Usage

### Basic Trie Operations

```go
// Create a new trie
t := NewTrie()

// Insert words
t.insert("apple")
t.insert("application")
t.insert("app")

// Check if a word exists
exists := t.contains_word("apple") // returns true

// Check if a prefix exists
hasPrefix := t.contains_prefix("app") // returns true

// Search for words with a prefix
results := t.search("app") // returns ["app", "apple", "application"]
```

### Fuzzy Search

```go
// Create a fuzzy search list
words := []string{"apple", "application", "app"}
searchList := new_Fuzzy_Search_list(words, "")

// Add a word to search
searchList.add_word("banana")

// Extend search term
searchList.search_term_extend("ap")

// Remove last character from search term
searchList.remove_char()
```

## API Reference

### Trie Structure

#### `NewTrie() *Trie`
Creates and returns a new empty trie.

#### `(t *Trie) insert(word string)`
Inserts a word into the trie.

#### `(t *Trie) contains_word(word string) bool`
Checks if a word exists in the trie.

#### `(t *Trie) contains_prefix(word string) bool`
Checks if any word in the trie starts with the given prefix.

#### `(t *Trie) search(word string) []string`
Returns all words that start with the given prefix.

#### `(t *Trie) human_search(word string, max_mistakes int) []string`
Performs fuzzy search with a maximum number of allowed mistakes.

### Fuzzy Search Structure

#### `new_Fuzzy_Search_list(words []string, search_term string) *Fuzzy_Search_list`
Creates a new fuzzy search list with the given words and initial search term.

#### `(f *Fuzzy_Search_list) add_word(word string)`
Adds a word to the search list.

#### `(f *Fuzzy_Search_list) search_term_extend(add_on_to_search_term string)`
Extends the current search term.

#### `(f *Fuzzy_Search_list) remove_char()`
Removes the last character from the search term.

## Configuration

The library provides several configurable options:

- **Max Mistakes**: Control the tolerance for fuzzy matching
- **Highlight Colors**: Customize the colors used for highlighting matches
- **Search Behavior**: Configure how the search handles special characters and case sensitivity

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a new branch for your feature
3. Make your changes
4. Run tests: `go test ./...`
5. Submit a pull request

### Coding Standards

- Follow Go's official [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Write tests for new functionality
- Update documentation for any changes
- Keep code simple and maintainable

## Troubleshooting

### Common Issues

1. **No matches found**
   - Check if the search term is correct
   - Verify that words are properly inserted into the trie
   - Adjust the max_mistakes parameter for fuzzy search

2. **Performance issues**
   - Ensure you're not inserting duplicate words
   - Consider using a larger max_mistakes value for fuzzy search
   - Check memory usage with large datasets

### Solutions

- Use the `contains_word` method to verify word insertion
- Monitor memory usage with large datasets
- Adjust search parameters based on your use case

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support, please open an issue in the GitHub repository or contact the maintainers.
