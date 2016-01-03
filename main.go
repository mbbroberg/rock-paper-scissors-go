package main

import (
  "fmt"
  "bufio"
  "os"
)

func robotPlay() string {
  return "Rock"
}

func main() {
  fmt.Println("Game on.")
  fmt.("Enter your move:")
  scanner := bufio.NewScanner(os.Stdin)

  // play forever within this loop
  for scanner.Scan() {
    fmt.Println("Rock..", "Paper..", "Scissor..", "Shoot!")
    userPlayed := scanner.Text()
    roboPlayed := robotPlay()

    fmt.Println("You played:", userPlayed)
    fmt.Println("Robot played:", roboPlayed)
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "issue with input:", err)
  }
}
