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
	s := NewState()
	for i := 0; i < count; i++ {
		first := randomMultiplier()
		second := randomMultiplier()
		result := getAnswer(first, second)
		actualResult := first * second
		if result == actualResult {
			fmt.Println("Dobrze!")
			s.markGood()
		} else {
			fmt.Println("Prawidłowa odpowiedź to", actualResult)
			s.markBad()
		}
	}

	fmt.Println("Udzielono", s.good, "odpowiedzi poprawnych, oraz", s.bad, "złych.")
	if s.passing() {
		fmt.Println("Dobrze!", name, "Udało Ci się uratować miasto!!")
	} else {
		fmt.Println("Niestety", name, "nie udało się uratować miasta :(")
	}

}

func NewState() *State {
	return &State{0, 0}
}

type State struct {
	good int
	bad  int
}

func (s *State) markGood() {
	s.good++
}

func (s *State) markBad() {
	s.bad++
}

func (s *State) passing() bool {
	return float32(s.bad)/float32(s.good) <= 0.2
}

func randomMultiplier() int64 {
	return rand.Int63n(8) + 2
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
