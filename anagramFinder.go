package main

import (
	"fmt"

	"kt/anagramFinder/anagrams"
)

func main() {
	// fmt.Println("Hello World")
	// foundAnagrams := anagrams.FindAnagramsInFile("anagrams/dictionary.txt", "DOCUMENTING")
	// fmt.Println(foundAnagrams)

	fmt.Println("Two word anagrams for DOCUMENTING:")
	anagramCounter := 0
	for pair := range anagrams.FindAnagramsInFileChan("anagrams/dictionary.txt", "DOCUMENTING") {
		fmt.Printf("%v - %v\n", pair.GetW1(), pair.GetW2())
		anagramCounter++
	}
	fmt.Printf("Found %v anagrams.\n", anagramCounter)

}
