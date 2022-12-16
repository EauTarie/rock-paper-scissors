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

    fmt.Println("Welcome to Rock, Paper Scissors !\nChoice between : rock, paper or scissors\n//--------//")
	z := getUserChoice("")

    fmt.Println("my big Conpyuteur choose : ",getPcChoice())
	if z == "rock" {
		fmt.Println("bonjour")
	}
	fmt.Printf("Hello world !")
    playAgain("")
    if (z== "valentin") {
        main()
    }
}
