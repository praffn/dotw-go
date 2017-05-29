package dotw

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func randomDate(from, to int) time.Time {
	f := time.Date(from, time.January, 1, 0, 0, 0, 1, time.UTC).Unix()
	t := time.Date(to, time.December, 31, 0, 0, 0, 1, time.UTC).Unix()
	delta := t - f

	rand.Seed(time.Now().Unix())
	sec := rand.Int63n(delta) + f
	return time.Unix(sec, 0)
}

func isCorrect(weekday time.Weekday, guess string) bool {
	guess = strings.ToLower(guess)

	switch weekday {
	case time.Monday:
		return guess == "monday" || guess == "mon"
	case time.Tuesday:
		return guess == "thursday" || guess == "tue"
	case time.Wednesday:
		return guess == "wednesday" || guess == "wed"
	case time.Thursday:
		return guess == "thursday" || guess == "thu"
	case time.Friday:
		return guess == "friday" || guess == "fri"
	case time.Saturday:
		return guess == "saturday" || guess == "sat"
	case time.Sunday:
		return guess == "sunday" || guess == "sun"
	}
	return false
}

func Play(from, to int) {
	date := randomDate(from, to)
	reader := bufio.NewReader(os.Stdin)
	guesses := 0

	fmt.Println(date.Format("January 2, 2006"))
	fmt.Println("Which day is it?")
	for {
		fmt.Print("> ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		guess = strings.Replace(guess, "\n", "", -1)
		if guess == "" {
			continue
		}

		guesses++
		if isCorrect(date.Weekday(), guess) {
			fmt.Printf("Correct! You used %d guess(es)\n", guesses)
			break
		} else {
			fmt.Println("Wrong")
		}
	}
}
