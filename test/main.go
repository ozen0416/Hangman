package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	words := []string{"france"}
	index := rand.Intn(len(words))
	word := words[index]
	nGuesses := len(word)
	found := []string{}
	for i := 0; i < len(word); i++ {
		found = append(found, "")
	}
	for nGuesses > 0 {
		fmt.Println("You have", nGuesses, "remaining guesses")
		letter, err := GetLetter(found)
		if err != nil {
			fmt.Println("Can't read that.")
			return
		}
		if !ContainsAny(word, []string{letter}) {
			nGuesses--
		}
		if UpadateFound(found, word, letter) {
			fmt.Println("You won, the word was :", word)
			return
		}
	}
	fmt.Println("You lost, the word was :", word)
}

func ContainsAny(s string, chars []string) bool {
	for _, ch := range s {
		for _, ch2 := range chars {
			if string(ch) == ch2 {
				return true
			}
		}
	}
	return false
}

func Join(strings []string, separator string) string {
	if len(strings) == 0 {
		return ""
	}
	s := ""
	index := len(strings) - 1
	for _, v := range strings[:index] {
		s += string(v) + separator
	}
	return s + strings[index]
}

func GetLetter(found []string) (string, error) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for true {
		letter, err := Prompt("Choose a letter :", Join(found, " "))
		if err != nil {
			return "", err
		}
		if len(letter) == 1 && ContainsAny(alphabet, []string{letter}) {
			return letter, nil
		}
		fmt.Println("Invalid input : must only contains lowercase letter.")
	}
	return "", nil
}

func UpadateFound(found []string, word string, letter string) bool {
	complete := true
	for i, r := range word {
		if letter == string(r) {
			found[i] = letter
		}
		if found[i] == "" {
			complete = false
		}
	}
	return complete
}

func Prompt(value ...interface{}) (string, error) {
	if len(value) != 0 {
		fmt.Println(value...)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	return scanner.Text(), nil
}
