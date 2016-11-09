package narrative

import (
	"fmt"
	"halloweencabin/actions"
	"halloweencabin/inputs"
	"halloweencabin/models"
)

//Midgame is the guts of the game, trigger encounters
func Midgame(cp models.Player) models.Player {

	//For testing things
	actions.Test1()

	//check current position and return options

	cr := models.RoomGet(cp.Position)

	fmt.Println("You are in the ", cr.Name)
	fmt.Println("What would you like to do?")
	//add other actions? When to move monster?
	fmt.Println("1.Look around the room\n2.Move to an adjacent room\n3.Use something in the room")
	s1 := "?"
	r1 := inputs.Basicinput(s1)
	if r1 == 1 {
		cp = actions.Lookr(cp)
	}
	if r1 == 2 {
		cp = actions.Move(cp)
	}
	if r1 == 3 {
		cp = actions.Usecheck(cp)
	}
	actions.Spawncheck(cp) //check to spawn monster, this way can be spawned after looking, moving or using
	//actions.Mroam(cp) turning it off for now
	return cp
}
