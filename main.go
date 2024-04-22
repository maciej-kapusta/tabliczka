package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("Jak masz na imię?")
	prompt()
	var name string
	fmt.Scanln(&name)
	fmt.Println("Witaj", name)
	fmt.Println("Pomóż nam uratować Kraków")
	for i := 0; i < 2; i++ {

		first := rand.Int63n(10)
		second := rand.Int63n(10)
		actualResult := first * second
		fmt.Println("Ile jest", first, "razy", second, "?")
		prompt()
		var answer string
		fmt.Scanln(&answer)
		result, err := strconv.ParseInt(answer, 10, 32)
		if err != nil {
			fmt.Println("Nie rozumiem")
		} else {
			if result == actualResult {
				fmt.Println("Dobrze!")
			} else {
				fmt.Println("Prawidłowa odpowiedź to", actualResult)
			}
		}

	}
}

func prompt() {
	fmt.Print("> ")
}
