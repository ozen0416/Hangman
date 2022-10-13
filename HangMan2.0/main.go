package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	content, err := os.Open("words.txt")

	if err != nil {
		log.Fatal(err)
		fmt.Println("An error has occur")
	}
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanWords)
	tmp := make([]string, 0)
	for scanner.Scan() {
		tmp = append(tmp, scanner.Text())
	}

	found := []string{}
	for i := 0; i < len(tmp); i++ {
		found = append(found, "_")
	}
	Guesses := 10
	for Guesses > 0 {
		fmt.Println("Hello! All the words you need to find are in english :)")
		fmt.Println("You have", Guesses, "remaining guesses")

	}

	rand.Seed(time.Now().Unix())
	result := rand.Intn(len(tmp) - 1)

	fmt.Println(tmp[result])
}

func GetLetter(found string) (string, error) {
	letter := ""
	var p rune
	if p >= 65 || p <= 90 && p >= 97 || p <= 122 {
		for _, a := range found {
			tmp := a
			if a >= 97 && a <= 122 {
				tmp = a - 32
			}
			letter += string(tmp)
		}

	} else {
		fmt.Println("This is not a letter !")
	}
	return letter, nil
}

func GetWord(guess string) {
	var word []string
	letter := ""

}

func Join(strings []string, separator string) string {
	if len(strings) == 0 {
		return ""
	}
	s := "_"
	index := len(strings) - 1
	for _, v := range strings[:index] {
		s += string(v) + separator
	}
	return s + strings[index]
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
