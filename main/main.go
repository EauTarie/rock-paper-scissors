package main

import (
	"fmt"
	"bufio"
    "os"
    "strings"
)

func StringPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}

func main() {

    fmt.Println("--------\n[MINIGAME] Welcome to Rock, Paper, Scissors !\n//--------//\nChoose between : \n-> rock\n-> paper\n-> scissors\n --------")
	user := getUserChoice("")
    computer := getPcChoice()
    fmt.Println("you have chosen : ", user, "\nComputer has choose : ",computer)

    winner(user,computer)

    rematch := playAgain("")

    if (rematch == "yes" || rematch == "y" || rematch == "oui") {
        main()
    } else {

    }
}
