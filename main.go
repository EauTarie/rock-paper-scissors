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

    fmt.Println("--------\n[MINI-JEU] Bienvenu dans le jeu du Pierre Papier Ciseaux  !\n//--------//\nChoisissez entre : \n- pierre\n- papier\n- ciseaux")
	user := getUserChoice("")
    computer := getPcChoice()
    fmt.Println("Vous avez choisis : ", user, "\nL'ordinateur a choisis : ",computer)

    winner(user,computer)

    rematch := playAgain("")

    if (rematch == "oui" || rematch == "o" || rematch == "Oui") {
        main()
    } else {

    }
}
