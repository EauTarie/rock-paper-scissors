package main

import "github.com/AlecAivazis/survey"

func winner()





func getUserChoice() userChoice string {
	
	selectPrompt := &survey.Select{
		Message: "Make a selection:",
		Options: []string{"rock", "paper", "scissors"},
	}
	survey.AskOne(selectPrompt, &choice)

	return
}

func getPcChoice() pcChoice string {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(3)

	switch randNum {
		case 1:
			choice = "rock"

		case 2:
			choice = "paper"

		case 3:
			choice ="scissor"
	}
	
	return
}


func duelVerdict(pcChoice, userChoice) winner string {
	if(pcChoice == userChoice) {
		winner = "TIE !"
		color.Yellow("Both picked %v ! It's a tie !", pcChoice)
	} else if((pcChoice == 'rock' && userChoice == 'papper') || (pcChoice == 'papper' && userChoice == 'scissor') || (pcChoice == 'scissor' && userChoice'rock')) {
		winner = "User ! "
		color.Green("You win !\nComputer picked %v and you picked %v", pcChoice, userChoice)
	} else {
		winner = "Computer ! "
		color.Red("You loose !\nComputer picked %v and you picked %v", pcChoice, userChoice)
	}
}