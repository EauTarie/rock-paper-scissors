package main

import (
	"fmt"
	"bufio"
    "os"
    "strings"
)

func playAgain(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+"//--------//\nDo you want to play again ? ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
	
}