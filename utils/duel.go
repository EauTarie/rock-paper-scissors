package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

func getPcChoice() string {
	pcChoice := ""
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(3)

	switch randNum {
	case 0:
		pcChoice = "pierre"

	case 1:
		pcChoice = "feuille"

	case 2:
		pcChoice = "ciseaux"
	}

	return pcChoice
}

func winner(user string, computer string) {
	if (user == "pierre" && computer == "ciseaux") || (user == "papier" && computer == "pierre") || (user == "ciseaux" && computer == "papier") {
		fmt.Printf("Vous avez gagné(e) !\n")
	} else if user == computer {
		fmt.Printf("Égalité !\n")
	} else {
		fmt.Printf("Vous avez perdu(e) !\n")
	}
}
