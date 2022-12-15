package matches

import "fmt"

func determinateWinner(pcChoices, userChoices string) (winner string) {
	if userChoices && pcChoices == "rock" || userChoices && pcChoices == "paper" || userChoices && pcChoices == "scissor" {
		winner = "Tie !"

	} else if userChoices == "rock" {
		if pcChoices == "paper" {
			fmt.Println("Rock being covered by Paper !\nYou loose !")
		} else {
			fmt.Println("Rock crushes Sissor !\nYou win !")
		}
	} else if userChoices == "paper" {
		if pcChoices == "scissor" {
			fmt.Println("Paper get cut by Scissor !\nYou loose !")
		} else {
			fmt.Println("Paper covers Rock !\nYou win !")
		}
	} else {
		if pcChoices == "rock" {
			fmt.Println("Scissor gets crushed by Rock!\nYou loose !")
		} else {
			fmt.Println("Scissor cuts Papper !\nYou win !")
		}
	}

	return
}