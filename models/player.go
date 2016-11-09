package models

import (
	"fmt"
	"halloweencabin/inputs"
)

//Player defines player stats
type Player struct {
	Age      int
	Health   int
	Charisma int
	Points   int
	Continue bool
	Position int
}

//Pset sets defaults player stats
func Pset() Player {
	np := Player{}
	np.Health = 100
	np.Charisma = 50
	np.Points = 0
	np.Continue = true
	np.Position = 1
	//age
	s := "How old are you?"
	i := inputs.Basicinput(s)
	np.Age = i
	if i < 18 {
		fmt.Println("Sorry, you are too young for this experience.\nGo play Monopoly with your babysitter.")
		np.Continue = false
		return np
	}
	fmt.Println("Looking for peace, you have rented a cabin in the woods. Somewhere isolated and free.\nYou arrive just as the sun starts setting.")
	return np
}
