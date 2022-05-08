//My First Program in GO!
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//create secret number
	secret := getRandomNumber()
	//fmt.Println(secret)
	for matching := false; !matching; {
		//get user input
		guess := getUserInput()
		//fmt.Println(secret, guess)

		//Make comparison (secret vs guess)
		matching = isMatching(secret, guess)
		//fmt.Println(matching)
	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your guess is too big!")
		return false
	} else if guess < secret {
		fmt.Println("Your guess is too small, sorry")
		return false
	} else {
		fmt.Println("GOOD JOB, YOU GOT IT!")
		return true
	}

}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11 // in order to get numbers from 0 to 10
}

func getUserInput() int {
	fmt.Print("Please input your guess number: ")

	var input int
	_, err := fmt.Scan(&input) //the address location of the variable
	if err != nil {
		fmt.Println("Failed to parse your input")
	} else {
		fmt.Println("You guessed:", input)
	}
	return input
}
