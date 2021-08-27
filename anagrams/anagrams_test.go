package anagrams

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		w1        string
		w2        string
		isAnagram bool
	}{
		{"word", "drow", true},
		{"word", "word", true}, //same word check
		{"god", "dog", true},
		{"act", "cat", true},
		{"balms", "lambs", true},
		{"wolf", "flow", true},
		{"looped", "poodle", true},
		{"lion", "loin", true},
		{"not", "anagram", false}, //diff length
		{"note", "made", false},   // same length, not anagram
	}
	for _, test := range tests {
		assert.Equal(t, isAnagram(test.w1, test.w2), test.isAnagram)
	}
}

func TestIsAnagramMultiword(t *testing.T) {
	var tests = []struct {
		w1        string
		w2        []string
		isAnagram bool
	}{
		{"documenting", []string{"document", "gin"}, true},
		{"network", []string{"net", "work"}, true}, //same word check
		{"network", []string{"not", "krew"}, true},
		{"looped", []string{"poo", "led"}, true},
		{"lion", []string{"lo", "in"}, true},
		{"not", []string{"anagram"}, false}, //diff length
		{"note", []string{"made"}, false},   // same length, not anagram
	}
	for _, test := range tests {
		assert.Equal(t, isAnagramMultiWord(test.w1, test.w2), test.isAnagram)
	}
}
func TestIsAnagramPair(t *testing.T) {
	var tests = []struct {
		w1        string
		w2        Pair
		isAnagram bool
	}{
		{"documenting", Pair{"document", "gin"}, true},
		{"documenting", Pair{"gin", "document"}, true}, //order
		{"network", Pair{"net", "work"}, true},         //same word check
		{"network", Pair{"not", "krew"}, true},
		{"looped", Pair{"poo", "led"}, true},
		{"lion", Pair{"lo", "in"}, true},
		{"not", Pair{"anagram", ""}, false}, //diff length
		{"note", Pair{"made", ""}, false},   // same length, not anagram
	}
	for _, test := range tests {
		assert.Equal(t, isTwoWordAnagram(test.w1, test.w2), test.isAnagram)
	}
}

func TestFindAllAnagramsInList(t *testing.T) {
	var words = []string{"word", "drow", "god", "dog", "act", "cat", "balms", "lambs", "wolf", "flow", "looped", "poodle", "lion", "loin", "not", "anagram", "note", "made"}

	var tests = []struct {
		input string
		found []string
	}{
		{"word", []string{"drow"}},
		{"god", []string{"dog"}},
		{"act", []string{"cat"}},
		{"balms", []string{"lambs"}},
		{"wolf", []string{"flow"}},
		{"looped", []string{"poodle"}},
		{"lion", []string{"loin"}},
		{"not", []string{}},
		{"note", []string{}},
	}
	for _, test := range tests {
		assert.Equal(t, reflect.DeepEqual(FindAllAnagramsInList(test.input, words), test.found), true) //TODO find sleeker impl
	}
}

func TestReadWordsFromFile(t *testing.T) {
	filterFunc := func(w string, w2 string) bool {

		return true
	}
	var words = GetWordsFromFile("wordlist.txt", "testWord", filterFunc)
	assert.Len(t, words, 1633)
}

func TestAnagramsInFile(t *testing.T) {

	var words = GetWordsFromFile("wordlist.txt", "bolster", isAnagram)
	assert.Contains(t, words, "lobster")
}

func TestReadWordsFromFileFileNotFound(t *testing.T) {
	filterFunc := func(w string, w2 string) bool {
		return true
	}
	GetWordsFromFile("does_not_exist.txt", "testWord", filterFunc)
	assert.Error(t, assert.AnError)
}

// find two word combintation
func TestFind2WordCombos(t *testing.T) {
	test := []string{"a", "b", "c", "d"}
	pairs := GetWordCombinations(test)
	fmt.Println(pairs)
	assert.Len(t, pairs, 12)
}

func TestFilterOnMembership(t *testing.T) {
	member := FilterOnLetterMembership("test", "set")
	member = FilterOnLetterMembership("bolster", "settttt") && member
	assert.True(t, member)

	var words = GetWordsFromFile("wordlist.txt", "bolster", FilterOnLetterMembership)
	fmt.Println(words)
	assert.Contains(t, words, "lobster")

}

func TestFindAnagrams(t *testing.T) {
	anagrams := FindAnagramsInFile("dictionary.txt", "DOCUMENTING")
	fmt.Println(anagrams)
	assert.Len(t, anagrams, 88)
}
