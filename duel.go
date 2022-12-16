package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
    "os"
    "strings"
)

func getUserChoice(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+">")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}

func getPcChoice()string {
	pcChoice := ""
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(3)
	
	switch randNum {
		case 0:
			pcChoice = "rock"

		case 1:
			pcChoice = "paper"

		case 2:
			pcChoice ="scissors"
	}
	
	return pcChoice
}

func winner(user string, computer string) {
	if ((user == "rock" && computer == "scissors") || (user == "paper" && computer == "rock") || (user == "scissors" && computer == "paper")) {
		fmt.Printf("You won !\n")
	} else if (user == computer) {
		fmt.Printf("Tie !\n")
	} else {
		fmt.Printf("You lose !\n")
	}
}