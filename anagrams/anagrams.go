package anagrams

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func isAnagramMultiWord(w1 string, words []string) bool {
	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(w)
	}
	w2 := sb.String()
	return isAnagram(w1, w2)
}

func isTwoWordAnagram(w1 string, pair Pair) bool {
	var sb strings.Builder
	sb.WriteString(pair.w1)
	sb.WriteString(pair.w2)
	w2 := sb.String()
	return isAnagram(w1, w2)
}

func isAnagram(w1, w2 string) bool {
	//check if words are of same length
	if len(w1) == len(w2) {
		//check if maps of letter counts are equal
		m1 := string2CountMap(strings.ToUpper(w1))
		m2 := string2CountMap(strings.ToUpper(w2))
		eq := reflect.DeepEqual(m1, m2) //TODO check out how this alg is implemented
		return eq
	} else {
		return false
	}
}

//NOTE this assumes no spaces in words
func string2CountMap(w string) map[string]int {
	m := make(map[string]int)
	for _, c := range w { //_ because we dont care about the index
		//Getting non-existant element in map will return a zero
		key := string(c)
		m[key] = m[key] + 1
	}
	return m
}

//start just walking through the list - O(n)
func FindAllAnagramsInList(w string, words []string) []string {
	anagrams := []string{}
	for _, word := range words { //readable vs creating fewer structs
		if isAnagram(w, word) && (w != word) { // is anagram && not the same word
			anagrams = append(anagrams, word)
		}
	}
	return anagrams
}

func GetWordsFromFile(filename string, candidateWord string, filterByWord func(string, string) bool) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return make([]string, 0)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// This is our buffer now
	var words []string
	var filteredWords []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, word := range words {
		if filterByWord(candidateWord, word) {
			filteredWords = append(filteredWords, word)
		}
	}
	return filteredWords

}

type Pair struct {
	w1 string
	w2 string
}

//ew
func (p Pair) GetW1() string {
	return p.w1
}
func (p Pair) GetW2() string {
	return p.w2
}

func GetWordCombinations(words []string) []Pair {
	var pairs []Pair

	for pair := range GenerateWordCombinations(words) {
		pairs = append(pairs, pair)
	}
	return pairs

}

func GenerateWordCombinations(words []string) <-chan Pair {
	c := make(chan Pair)
	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan Pair) {
		defer close(c)
		CreatePair(c, words)
	}(c)
	return c // Return the channel to the calling function
}

func CreatePair(c chan Pair, words []string) {
	for i, v := range words {
		for i2, v2 := range words {
			if i != i2 {
				c <- Pair{string(v), string(v2)}
			}
		}
	}

}

func FilterOnLetterMembership(targetWord string, w2 string) bool {
	member := true
	for _, letter := range w2 {
		member = strings.Contains(targetWord, string(letter)) && member
	}
	//check if maps of letter counts are equal
	return member

}

//ASSUME UPPER CASE WORDS IN FILE
func FindAnagramsInFile(filename string, targetWord string) []Pair {
	target := strings.ToUpper(targetWord)
	words := GetWordsFromFile(filename, target, FilterOnLetterMembership)

	fmt.Printf("Found %v words in file with letters subset of %v \n", len(words), target)

	var anagrams []Pair
	for pair := range GenerateWordCombinations(words) {
		if isTwoWordAnagram(target, pair) {
			anagrams = append(anagrams, pair)
		}
	}
	return anagrams
}

func FindAnagramsInFileChan(filename string, targetWord string) <-chan Pair {
	target := strings.ToUpper(targetWord)
	words := GetWordsFromFile(filename, target, FilterOnLetterMembership)

	fmt.Printf("Found %v words in file with letters subset of %v \n", len(words), target)

	c := make(chan Pair)
	// Starting a separate goroutine that will create all the combinations,
	// feeding them to the channel c
	go func(c chan Pair) {
		defer close(c)
		for pair := range GenerateWordCombinations(words) {
			if isTwoWordAnagram(target, pair) {
				c <- pair
			}
		}

	}(c)
	return c // Return the channel to the calling function

}
