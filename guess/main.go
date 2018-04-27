package main

import (
  "bufio"
  "fmt"
  "math/rand"
  "os"
  "strconv"
  "time"
)

func main() {
  rand.Seed(time.Now().Unix())
  fmt.Println("Welcome to the Guessing Game")

  secretNumber := rand.Intn(10)
  scanner := bufio.NewScanner(os.Stdin)

  for {
    scanner.Scan()

    input, err := scanner.Text(), scanner.Err()
    if err != nil {
      fmt.Println("What??")
      continue
    }

    num, err := strconv.Atoi(input)
    if err != nil {
      fmt.Println("Must be a number")
      continue
    }

    if num == secretNumber {
      fmt.Println("YAY.  You win!")
      break
    }

    fmt.Println("Nope.  Try again.")
  }

  fmt.Println("GAME OVER")
}
