package narrative

import "halloweencabin/models"

//buildHouse sets up the basics and can reset everything once the game repeats
func buildHouse() models.Player {
	np := models.Pset() //Sets up new player with basic stat
	//models.Houseset(np)   //Sets up fresh house
	models.Iset(np) //Sets up items including phone
	models.Mset(np) //Sets up monster
	return np
}

// Intro launches the start of the game
func Intro() models.Player {
	np := buildHouse()
	return np
}
