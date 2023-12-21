package main

import (
  "errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
  problems := ReadFile("problems.csv") 
  highscores := ReadFile("highscore.csv")
  highscore := [][]string{}
  vals := [][]string{}

  // Get all the problems
  for i := 0; i < len(problems); i += 1 {
    str := problems[i]
    question := strings.Split(str, ",")

    if len(question) < 2 {
      continue
    }

    vals = append(vals, question)
  }

  // Get them highscores
  for i := 0; i < len(highscores); i+=1 {
    line := highscores[i]
    cols := strings.Split(line, ",")

    if len(cols) >= 3 {
      highscore = append(highscore, cols)
    }
  }

  println("Welcome to math hell")
  println()

  // Game starts here
  gameloop:
  score := 0 

  println("how many games u wanna play")
  gameCount := GetGameCount()

  for ; gameCount > 0; gameCount-- {
    index := rand.Intn(len(vals))
    problem := vals[index]
    
    answer, err := strconv.ParseInt(strings.TrimSpace(string(problem[1])), 10, 64)
    if err != nil {
      panic(err)
    }

    print(string(problem[0]))

    var i int
    fmt.Scan(&i)

    if i == int(answer) {
      score += 1
      println("ur nod dumb")
    } else {
      fmt.Printf("guess u r dumb its %d", answer)
      println()
    }
  }

  // U got stuff
  fmt.Printf("u got %d", score)
  println()
  println("ur done, do u wanna stab pins in ur cock again")
  println("y/n")

  // Exit game loop at some point
  var yesorno string
  fmt.Scan(&yesorno)

  if string(yesorno) == "y" {
    goto gameloop
  } else {
    println("bye")
    os.Exit(0)
  }
}

func GetGameCount() (int) {
  var gameCount int
  fmt.Scan(&gameCount)

  if gameCount <= 0 {
    gameCount = 1
  }

  return gameCount
}

func RemoveFromArray(dest [][]string, rmIndex int) ([][]string) {
  var newArray [][]string

  for i := 0; i < len(dest); i += 1 {
    if i == rmIndex {
      continue
    }

    newArray = append(newArray, dest[i])
  }

  return newArray
}

func CheckFileExists(filename string) {
  if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
    os.Create(filename)
  }
}

func ReadFile(filename string) ([]string) {
  CheckFileExists(filename)
  handle, err := os.ReadFile(filename)
  if err != nil {
    panic("yknow i cba to debug this properly mane")
  }

  lines := strings.Split(string(handle), "\n")
  return lines
}


