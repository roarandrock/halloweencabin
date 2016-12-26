package actions

import (
	"fmt"
	"halloweencabin/models"
)

//Monstercheck checks the new player position and starts battle loop
func Monstercheck(cp models.Player) models.Player {
	cm := models.Spawnmonsterget()
	if cp.Position == cm.Position {
		bcheck := 1 //1 means fight, 0 means stop
		//plays recurring text if player already met monster. Otherwise intro text and updates that
		if cm.Met == true {
			fmt.Println(cm.Found)
		}
		if cm.Met == false {
			//fmt.Println("You see a", cm.Name)
			fmt.Println(cm.Intro)
			cm.Met = true
			models.Monsterupdate(cm)
		}
		for bcheck == 1 { //need to check if player is alive, monster is alive, and they are in the same room
			cp, cm = battle(cp, cm)
			if cm.Position == cp.Position {
				if cm.Health > 0 && cp.Health > 0 {
					models.Monsterhealth(cm)
				}
			}
			bcheck = 0
			if cm.Health > 0 && cp.Health > 0 {
				if cp.Position == cm.Position {
					bcheck = 1
				}
			}
		}
		if cm.Health <= 0 {
			cp.Continue = false
			fmt.Println(cm.Outrom)
		}
		if cp.Health <= 0 && cm.Health > 0 { //player wins ties
			cp.Continue = false
			//fmt.Println(cm.Outrop) moved to ending.go
		}
		models.Monsterupdate(cm)
		return cp
	}
	return cp
}

//Monsterspawn gets chosen monster and spawn it
func Monsterspawn() {
	cm := models.Chosenmonsterget()
	cm.Spawn = true
	models.Monsterupdate(cm)
}

//Spawncheck checks if monster should spawn. Allows for monster specific spawning
func Spawncheck(np models.Player) {
	cm := models.Chosenmonsterget()
	if cm.Spawn == false {
		switch cm.Number {
		case 1:
			if np.Position != 1 {
				Monsterspawn() //currently spawns when player is not in the living room
				fmt.Println("You hear the roar of a chainsaw. The crash and tearing of metal chewing wood.\nIt's coming from the living room.")
				fd := models.ItemGet("front door")
				fd.Used = true
				nd := models.ItemGet("ruins of front door")
				nd.Loc = 1
				models.Itemupdate(fd)
				models.Itemupdate(nd)
				cm = models.Chosenmonsterget() //to get fresh copy
				cm.Position = 1
			}
		case 2:
			skillet := models.ItemGet("skillet")
			if skillet.Toggle == false && np.Position != 1 {
				Monsterspawn() //currently spawns when player uses skillet
				fmt.Println("You hear a loud crack and crash. It sounds like something just tore down your front door.")
				fd := models.ItemGet("front door")
				fd.Used = true
				nd := models.ItemGet("ruins of front door")
				nd.Loc = 1
				models.Itemupdate(fd)
				models.Itemupdate(nd)
				cm = models.Chosenmonsterget() //to get fresh copy
				cm.Position = 1
			} else {
				fmt.Println("Staying in this cabin alone is hungry work. Your stomach is growling.")
			}
		case 3:
			bell := models.ItemGet("silver bell")
			if bell.Toggle == false {
				Monsterspawn()
				cm = models.Chosenmonsterget()
				cm.Position = 3
			} else {
				fmt.Println("With the snow falling outside, you hope it will be a white Christmas this year.")
			}
		default:
			fmt.Println("Nothing evil lurks in these woods.")
		}
	}
	models.Monsterupdate(cm)
}

//Mroam moves the monster in the house when not next to the player.
func Mroam(cp models.Player) models.Player {
	cm := models.Spawnmonsterget()
	if cm.Spawn == false { //in case monster not spawned
		return cp
	}
	if cp.Position == cm.Position { //in case already in the same place
		return cp
	}
	//for making the monster adjacent
	cr := models.RoomGet(cp.Position)
	l := len(cr.Adj)
	if l == 0 {
		return cp
	}
	adjr := models.Madj(l)
	cm.Position = cr.Adj[adjr]
	models.Monsterupdate(cm)
	return cp
}
