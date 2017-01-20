/* Mini Horror Game
Author: roarandrock
Short term:

Added Santa. Need to implement battle with him now.


long term:
7. Monster outro playing when player kills themselves. Should change where it goes in the loop.
Need to pass back game over toggle and reason for ending through all actions to the main loop.
8. fixed monster resetting, can refactor the monster calls
9. add new basic input function for yes/no, autmoatically makes 1. Yes 2. No. And one for 1,2 statements maybe?
10.can go over 100 with health for player, need to cap
11. Encumberance and losing hands to the wolf
12. specific damage to body parts
add basement with zombies later, need a room updated similar to monsters/items, or can add a door item to be checked
make an attack function. Just handles damage and stuff. can use outside of battle loop then. Only use battle items for specific things
need item descriptions
test function to tally damage from battles per monster
*/

package main

import (
	"fmt"
	"halloweencabin/narrative"
	"log"
)

func main() {
	counter := 0 // how many plays
	c := 0       // continuation
	for c == 0 { //checks for retry
		player := narrative.Intro() //need to set/reset everything
		check := narrative.Failcheck(player)
		for check == true {
			player = narrative.Midgame(player)
			check = narrative.Failcheck(player)
		}
		narrative.Ending(player)
		c = narrative.Tryagain() //changes retry
		counter++
	}
	fmt.Println("Game Over at: ")
	log.Println(log.Ldate)
	fmt.Println(counter, " tries")
	fmt.Println("Happy Halloween")
}
