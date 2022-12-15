package matches

import (
	"math/rand"
	"time"
)

func GenerateRandomChoices() (choice string) {
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