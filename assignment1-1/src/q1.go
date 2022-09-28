package cos418_hw1_1

import (
	"fmt"
	i "io/ioutil"
	l "log"
	r "regexp"
	"sort"
	st "strings"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// read file and catch an error if it exists
	content, err := i.ReadFile(path)
	if err != nil {
		l.Fatal(err)
	}
	// convert content []byte to string and format it
	mainText := st.ToLower(string(content))
	// delete spaces and convert to array
	totalText := st.Fields(mainText)
	// panic when not needed characters (mask)
	b := r.MustCompile("[^0-9a-zA-Z]+")
	//create map with string key and int value - count of words
	wordsMap := make(map[string]int)
	//range - for slice or map - return for each iteration [index, element]
	for _, word := range totalText {
		//remove all not letters
		word := b.ReplaceAllString(string(word), "")
		if len(word) >= charThreshold {
			//if word in wordsMap, then exists is true
			if _, exists := wordsMap[word]; exists {
				wordsMap[word] ++
			} else {
				wordsMap[word] = 1
			}
		}
	}

	counter := 0
	final := make([] WordCount, len(wordsMap))
	for word, k := range wordsMap {
		final[counter] = WordCount{word, k}
		counter ++
	}
	sortWordCounts(final)

	if (numWords < len(final)){
		return final[:numWords]
	} else {
		return final
	}



	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"
	//return nil
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
