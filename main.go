package main

import (
	r "bufio"
	f "fmt"
	"os"
	scvt "strconv"
	ss "strings"
)

func main() {

	max := getMax()

	words := getWords()

	// For each number from 1 to max print the number or set of keywords
	for i := 1; i <= max; i++ {

		f.Println(fizzBuzz(words, i))

	}

}

// Format this final output based on the given array of keywords
func fizzBuzz(words map[string]int, i int) string {

	var out string

	// Anonymous function serves as a tenary operation
	// If i is divisible by v return the word else return nothing
	q := func(k string, v int, i int) string {
		if i%v == 0 {
			return k
		}
		return ""
	}

	// Iterate over the map of words
	for k, v := range words {
		out += q(k, v, i)
	}

	// I want to be sure that the number is returned in all other cases
	// I also want to tag the number in brackets otherwise (easier to read)
	if out == "" {
		out = scvt.Itoa(i)
	} else {
		out += "(" + scvt.Itoa(i) + ")"
	}

	return out
}

// Ask the user for a max value or return the default
func getMax() int {

	// These are arbitrary numbers
	def := 100
	min := 1

	reader := r.NewReader(os.Stdin)

	for {
		f.Printf("Please give a max value (>= %v) and press ENTER\n", min)
		f.Printf("Enter for default (Default = %v)\n", def)

		text, _ := reader.ReadString('\n')
		text = ss.Replace(text, "\n", "", -1)

		if text == "" {
			return def
		}

		iTxt, err := scvt.Atoi(text)
		if iTxt < min {
			f.Printf("%v < %v, try again\n", iTxt, min)
		} else if err == nil {
			return iTxt
		}

		f.Println("Please give a valid integer or no value for default")

	}
}

// Gets user input to populate a map of keywords and ints
func getWords() map[string]int {

	reader := r.NewReader(os.Stdin)

	words := make(map[string]int)

	for {
		f.Println("-----")
		f.Println("Current words are:")
		if len(words) > 0 {
			for k, v := range words {
				f.Printf("%q = %v\n", k, v)
			}
		} else {
			f.Println("No words")
		}
		f.Println()
		f.Println("Enter - Finish updating words")
		f.Println("Add - Add a word")
		f.Println("Mod - Change a word value")
		f.Println("Del - Remove a word")

		text, _ := reader.ReadString('\n')
		cmd := ss.ToLower(ss.Replace(text, "\n", "", -1))

		switch cmd {
		case "":
			return words
		case "add":
			addWord(words)
		case "mod":
			modWords(words)
		case "del":
			delWords(words)
		default:
			continue
		}

	}
}

// Function to add a new custom key/val to the map
func addWord(words map[string]int) {

	reader := r.NewReader(os.Stdin)

	// Get a kewyowrd
	f.Println("Give a word:")
	f.Println("Warning, giving a prexisting word will update the value")
	text, err := reader.ReadString('\n')

	// Hate this repetitious if err return
	if err != nil {
		f.Println(err)
		return
	}

	key := ss.Replace(text, "\n", "", -1)
	f.Printf("[%q] recieved...\n", key)

	// Get a value
	val := getVal(words)
	if val == 0 {
		return
	}

	// Maps are always passed by reference conveniently
	words[key] = val

}

// Fucntion to modify the value of an existing key
func modWords(words map[string]int) {

	// Display word options
	ok := displayWords(words)
	if !ok {
		return
	}

	// Wait for a word
	f.Printf("-> ")
	reader := r.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		f.Println(err)
		return
	}

	// Escape if the user gave the wrong word
	word := ss.Replace(text, "\n", "", -1)
	_, ok = words[word]
	if !ok {
		f.Println("No word found, please try again")
		return
	}

	// Update the value
	val := getVal(words)
	if val == 0 {
		return
	}

	words[word] = val
}

func delWords(words map[string]int) {

	// Display word options
	ok := displayWords(words)
	if !ok {
		return
	}

	reader := r.NewReader(os.Stdin)
	f.Println("Type a word to remove:")
	text, err := reader.ReadString('\n')

	// Hate this repetitious if err return
	if err != nil {
		f.Println(err)
		return
	}

	key := ss.Replace(text, "\n", "", -1)
	f.Printf("[%q] recieved...\n", key)

	_, ok = words[key]
	if !ok {
		f.Println("Word not found")
		return
	}

	delete(words, key)
}

// Takes user input, converts to int and checks it is a valid choice
func getVal(words map[string]int) int {
	f.Println("Give a unique positive integer value > 0:")

	reader := r.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		f.Println(err)
		return 0
	}

	val, err := scvt.Atoi(ss.Replace(text, "\n", "", -1))
	if err != nil {
		f.Println(err)
		return 0
	}

	if val <= 0 {
		f.Println("Please pick a number above 0")
		return 0
	}

	for k, v := range words {
		if val == v {
			f.Printf("%v already found with key %q\n", v, k)
			return 0
		}
	}
	return val
}

func displayWords(words map[string]int) bool {
	if len(words) > 0 {
		// Display current words in the map
		f.Println("Pick a word:")
		for k, v := range words {
			f.Printf("%q = %v\n", k, v)
		}
	} else {
		f.Println("No words, please add a word.")
		return false
	}
	return true
}
