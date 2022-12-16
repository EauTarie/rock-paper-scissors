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
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(3)
	
	pcChoice := ""
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


// func duelVerdict(pcChoice, userChoice) (winner string) {
// 	if(pcChoice == userChoice) {
// 		winner = "TIE !\n"
// 		color.Yellow("Both picked %v ! It's a tie !", pcChoice)
// 	} else if ((pcChoice == "rock" && userChoice == "papper") || (pcChoice == "papper" && userChoice == "scissors") || (pcChoice == "scissors" && userChoice == "rock")) {
// 		winner = "User !"
// 		color.Green("You win !\nComputer picked %v and you picked %v", pcChoice, userChoice)
// 	} else {
// 		winner = "Computer !"
// 		color.Red("You loose !\nComputer picked %v and you picked %v", pcChoice, userChoice)
// 	}
// }