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

    fmt.Println("--------\n[MINIJEU] Bienvenu sur Pierre, Feuille, Ciseaux !\n//--------//\nChoisissez entre : \n- pierre\n- feuille\n- ciseaux")
	user := getUserChoice("")
    computer := getPcChoice()
    fmt.Println("Vous avez choisis : ", user, "\nL'Ordinateur a choisi : ",computer)

    winner(user,computer)

    rematch := playAgain("")

    if (rematch == "oui" || rematch == "o" || rematch == "yes") {
        main()
    } else {

    }
}
