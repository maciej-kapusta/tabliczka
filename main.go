package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	name := readValue("Jak masz na imię?")

	fmt.Println("Witaj", name)
	fmt.Println("Pomóż nam uratować Kraków")
	count := rand.Intn(10)
	fmt.Println("Zadam Ci", count, "pytań")
	good := 0
	bad := 0
	for i := 0; i < count; i++ {
		first := randomMultiplier()
		second := randomMultiplier()
		result := getAnswer(first, second)
		actualResult := first * second
		if result == actualResult {
			fmt.Println("Dobrze!")
			good++
		} else {
			fmt.Println("Prawidłowa odpowiedź to", actualResult)
			bad++
		}
	}

	fmt.Println("Udzielono", good, "odpowiedzi poprawnych, oraz", bad, "złych.")

}

func randomMultiplier() int64 {
	return rand.Int63n(9) + 1
}

func getAnswer(first int64, second int64) int64 {
	for {
		answer := readValue("Ile jest", first, "razy", second, "?")
		result, err := strconv.ParseInt(answer, 10, 32)
		if err != nil {
			fmt.Println("Nie rozumiem")
			continue
		}
		return result
	}
}

func readValue(question ...any) string {
	var answer string
	fmt.Println(question...)
	fmt.Print("> ")
	_, _ = fmt.Scanln(&answer)
	return answer
}
