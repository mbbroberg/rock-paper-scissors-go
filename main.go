package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
  "strings"
)

var validPlay = map[int]string{
	0: "Rock",
	1: "Paper",
	2: "Scissors",
}

func robotPlay() string {
	playCalc := rand.Intn(len(validPlay))
	roboPlayed := strings.Title(validPlay[playCalc])
	return roboPlayed
}

func winningRules(player1 string, player2 string) string {
	switch {
	case player1 == player2:
		return "same"
	case player1 == "Rock":
		if player2 == "Paper" {
			return player2
		} else {
			return player1
		}
	case player1 == "Paper":
		if player2 == "Scissors" {
			return player2
		} else {
			return player1
		}
	case player1 == "Scissors":
		if player2 == "Rock" {
			return player2
		} else {
			return player1
		}
	}
  return "I got nothing here"
}

func main() {
	fmt.Println("Game on.")
	fmt.Println("-----------------------")
	fmt.Printf("Enter your move: ")
	scanner := bufio.NewScanner(os.Stdin)

	// play forever within this loop
	for scanner.Scan() {
		fmt.Println("Rock..", "Paper..", "Scissor..", "Shoot!")
		userPlayedRaw := scanner.Text()

    if userPlayedRaw == strings.ToLower("quit") {
      fmt.Println("Quitting time! Thanks for playing.")
      break
    }
        
    userPlayed := strings.Title(userPlayedRaw)
		roboPlayed := robotPlay()

		fmt.Println("You played:", userPlayed)
		fmt.Println("Robot played:", roboPlayed)

		winner := winningRules(userPlayed, roboPlayed)

		if winner == "same" {
			fmt.Println("It's a tie. Go again!\n")
		} else if winner == roboPlayed {
			fmt.Println("You lose!\n")
		} else if winner == userPlayed {
			fmt.Println("You won!\n")
		} else {
			fmt.Println("Uh oh. We had an issue. Play again!\n")
		}

		fmt.Println("-----------------------")
		fmt.Printf("Enter your move: ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "issue with input:", err)
	}
}
