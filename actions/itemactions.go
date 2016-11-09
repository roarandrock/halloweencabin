package actions

import (
	"fmt"
	"halloweencabin/inputs"
	"halloweencabin/models"
)

//Useitem attempt at using items
func Useitem(item1 models.Item, action int, cp models.Player) (models.Item, models.Player) {
	//monster := models.Chosenmonsterget()
	//mid := monster.Number //number for chosen monster
	switch {
	case item1.Name == "phone":
		switch action {
		case 0:
			fmt.Println("You attempt to call and only hear dial tone. There is no reception.")
		case 1:
			fmt.Println("You attempt to call and only hear dial tone. There is no reception.")
		case 2:
			fmt.Println("You attempt to call and only hear dial tone. There is no reception.")
		default:
			fmt.Println("OK")
		}
	case item1.Name == "umbrella":
		switch action {
		case 0:
			switch item1.Toggle {
			case true:
				fmt.Println("You open the umbrella. And now you are protected. From indoor rainstorms.")
				item1.Toggle = false
			case false:
				fmt.Println("You close it. And now you are protected from bad luck.")
				item1.Toggle = true
			}
		case 1:
			fmt.Println("You dance with your ella, ella, ella.")
		default:
			fmt.Println("OK")
		}
	case item1.Name == "front door":
		switch action {
		case 0:
			if item1.Toggle == true {
				fmt.Println("You open the front door. Get a whiff of that nature.")
				item1.Toggle = false
			}
			if item1.Toggle == false {
				fmt.Println("You close the front door. Generates a false sense of security.")
			}
		case 1:
			switch item1.Toggle {
			case true:
				fmt.Println("But the door is the shut.")
			case false:
				fmt.Println("You go outside.")
				cp = outside(cp)
			}
		default:
			fmt.Println("OK")
		}
	case item1.Name == "ruins of front door":
		switch action {
		case 0:
			fmt.Println("You step over the splintered remains and run out into the woods.")
			cp = outside(cp)
		case 1:
			fmt.Println("You grab a short length of scrap wood from the previous door.")
			sw := models.ItemGet("scrap wood")
			sw.Loc = 20
			models.Itemupdate(sw)
		default:
			fmt.Println("OK")
		}
	case item1.Name == "scrap wood":
		switch action {
		case 0:
			fmt.Println("With this conveniently splintered wood you build a fire.")
			item1.Used = true
			cp = fire(cp)
		default:
			fmt.Println("OK")
		}
	case item1.Name == "skillet":
		switch action {
		case 0:
			if cp.Position == 2 {
				//call cook function
				cp, item1 = cook(cp, item1)
			} else {
				fmt.Println("You cannot cook here. You need heat.")
			}
		default:
			fmt.Println("OK")
		}

	case item1.Name == "fridge":
		switch action {
		case 0:
			fmt.Println("Inside you see a massive supply of burgers, ketchup, and old mayo. No beer.")
			bi := models.ItemGet("burger meat")
			if bi.Loc == 10 {
				fmt.Println("Take some burger meat?")
				fmt.Println("1. Yes\n2. No")
				r1 := inputs.Basicinput("?")
				if r1 == 1 {
					fmt.Println("You take a handful of cold meat.")
					bi := models.ItemGet("burger meat")
					bi.Loc = 20
					models.Itemupdate(bi)
				}
			}
			if item1.Toggle == true {
				fmt.Println("Look in the freezer?")
				fmt.Println("1. Yes\n2. No")
				r2 := inputs.Basicinput("?")
				switch r2 {
				case 1:
					fmt.Println("This needs to be defrosted. But buried in the ice you see a bottle of vodka.\nYou take the bottle.\nNo point in leaving it for the next occupant.")
					item1.Toggle = false
					vi := models.ItemGet("vodka bottle")
					vi.Loc = 20
					models.Itemupdate(vi)
				default:
					fmt.Println("You close the fridge. To save electricity and keep the ketchup cold.")
				}
			}
		default:
			fmt.Println("OK")
		}
	case item1.Name == "vodka bottle":
		switch action {
		case 0:
			fmt.Println("You take a short pull. It does nothing for your health but you feel better.")
			cp.Charisma = cp.Charisma + 10
		case 1:
			fmt.Println("You smash the bottle. Now there are glass shards and wasted vodka all over.") //can place in the burgers
			item1.Used = true
			gsi := models.ItemGet("glass shards")
			gsi.Loc = cp.Position
			models.Itemupdate(gsi)
		default:
			fmt.Println("OK")
		}
	case item1.Name == "glass shards":
		switch action {
		case 0:
			fmt.Println("Crunchy. And painful. Stupid idea.")
			cp.Health = cp.Health - 25
		default:
			fmt.Println("OK")
		}
	case item1.Name == "burger meat":
		switch action {
		case 0:
			if cp.Position == 2 {
				cp, item1 = cook(cp, item1)
			} else {
				fmt.Println("Cannot cook here. Where's the heat for the meat?")
			}
		case 1:
			switch item1.Toggle {
			case true:
				fmt.Println("No, it's raw. Gross.")
			case false: // cooked
				gsi := models.ItemGet("glass shards")
				cri := models.ItemGet("silver cross")
				if gsi.Loc == 30 && gsi.Used == false { //glass is in burger and not used yet
					fmt.Println("You feel the glass puncture holes in your mouth, throat and stomach. Terrible idea.")
					cp.Health = cp.Health - 60
					item1.Toggle = true
					item1.Loc = 10
					gsi.Used = true
					gsi.Loc = 40
					models.Itemupdate(gsi)
				} else if cri.Loc == 30 && cri.Used == false {
					fmt.Println("You barely notice the silver taste.")
					cp.Health = cp.Health + 15
					item1.Toggle = true
					item1.Loc = 10
					cri.Used = true
					cri.Loc = 40
					models.Itemupdate(gsi)
				} else {
					fmt.Println("That was a tasty burger. You feel a bit better. And nothing attacked you this whole time.")
					cp.Health = cp.Health + 15
					item1.Toggle = true
					item1.Loc = 10
				}
			}
		default:
			fmt.Println("OK")
		}
	case item1.Name == "axe":
		switch action {
		case 0:
			fd := models.ItemGet("front door")
			if cp.Position == 1 {
				switch fd.Used {
				case true:
					fmt.Println("There is nothing here to chop. If you need scrap wood, something already destroyed the front door.")
				case false:
					fmt.Println("You destroy the front door with the axe. Do you suffer from cabin fever?")
					fd.Used = true
					dfd := models.ItemGet("ruins of front door")
					dfd.Loc = 1
					models.Itemupdate(fd)
					models.Itemupdate(dfd)
				}
			} else {
				fmt.Println("There is nothing here to chop. Except the cabin itself. And that would be crazy.")
			}
		default:
			fmt.Println("OK")
		}
	case item1.Name == "bear trap":
		switch action {
		case 0:
			fmt.Println("Carefully you pry apart the metal teeth. You have a rough idea how this works. You've seen it before in cartoons." +
				"\nYou try to move that center bit with your foot.\nThe jaws slam shut around your leg.\nOuch.")
			cp.Health = cp.Health - 25
			item1.Used = true
		default:
			fmt.Println("OK")
		}
	case item1.Name == "bed":
		switch action {
		case 0:
			fmt.Println("You get under the stained covers. It's surprisingly comfortable. Soon you are in dreamland.")
			cm := models.Spawnmonsterget()
			if cm.Spawn == false {
				fmt.Println("You awake after a few hours of nightmare-ridden sleep.")
			} else {
				fmt.Println("You do not awake when the", cm.Name, "other enters the room.\nYou awake when they are standing next to your bed. And you are vulnerable as a babe.")
				cp.Health = 0
			}
		case 1:
			fmt.Println("You hide under the bed. It's just you, the spiders and broken toys.")
			cm := models.Spawnmonsterget()
			if cm.Spawn == false {
				fmt.Println("This is boring. You leave.")
				cp.Charisma = cp.Charisma - 15
			} else {
				//make and call hide function, can be useful elswhere
				cp, item1 = hide(cp, cm, item1) //only works once.
			}
		default:
			fmt.Println("OK")
		}
	case item1.Name == "silver cross":
		switch action {
		case 0:
			fmt.Println("It's an old cross on a necklace. You eat it. It's gone forever.")
			item1.Used = true
		default:
			fmt.Println("OK")
		}
	case item1.Name == "red plunger":
		switch action {
		case 0:
			if item1.Toggle == true {
				if cp.Position == 6 {
					fmt.Println("You look at the toilet. It is clogged. And really gross. Like ten sick babies used it. Sure you want to do this?")
					fmt.Println("1.Yes\n2.No")
					r1 := inputs.Basicinput("?")
					if r1 == 1 {
						fmt.Println("You plunge away. Filth and water spill over and onto your shoes. Keep plunging?")
						fmt.Println("1.Yes\n2.No")
						r2 := inputs.Basicinput("?")
						if r2 == 1 {
							fmt.Println("You keep going. The smell is terrible. You vomit in the tub multiple times." +
								"Eventually you unclog it. It's a small bottle with a note in it. Read the note?")
							fmt.Println("1.Yes\n2.No")
							r3 := inputs.Basicinput("?")
							if r3 == 1 {
								fmt.Println("The note says: \"Kill Yourself. Use the razor and kill yourself.\"")
								item1.Toggle = false
							}
						}
					}
				} else {
					fmt.Println("Nothing here to plunge.")
				}
			} else {
				fmt.Println("No. That was depressing.")
			}
		default:
			fmt.Println("OK")
		}
	case item1.Name == "rusty straight razor":
		switch action {
		case 0:
			fmt.Println("You cut yourself. It's rusty and you're no barber.")
			cp.Health = cp.Health - 15
			cm := models.Spawnmonsterget()
			if cm.Number == 2 && cm.Spawn == true {
				fmt.Println("Your blood drips on the floor.")
				cm.Position = cp.Position
				models.Monsterupdate(cm)
			}
		default:
			fmt.Println("OK")
		}
		/*
			case item1.Name == :
				switch action{
				case 0:
				default:
					fmt.Println("OK")
				}
		*/
	default:
		fmt.Println("Surprisingly, it does not do anything")
	}
	return item1, cp
}

func hide(cp models.Player, cm models.Monster, item1 models.Item) (models.Player, models.Item) {
	//needs content, only called when monster already exists
	fmt.Println("You feel safe here. You wait.")
	switch cm.Number {
	case 1:
		fmt.Println("You see a pair of boots. They enter the room, approach the bed." +
			"\nHe throws back the bed cover. Nothing. You can feel his disappointment.")
		switch item1.Toggle {
		case true:
			fmt.Println("He turns and starts walking away.")
			fmt.Println("1. Ambush the monster\n2. Stay hidden")
			r1 := inputs.Basicinput("?")
			switch r1 {
			case 1:
				fmt.Println("Like a ninja, you emerge and sneak up behind your foe.")
				axei := models.ItemGet("axe")
				rzi := models.ItemGet("rusty straight razor")
				switch {
				case rzi.Loc == 20:
					fmt.Println("You try to slit his throat. The blood spills hot and fast onto your hand." +
						"\nHe screams and frantically strikes behind him. He manages to push you off." +
						"\nHe flees the room. One hand on his chainsaw, the other holding his bleeding throat.")
					cm.Health = cm.Health - 60
					cm = Mrun(cp, cm)
				case axei.Loc == 20:
					fmt.Println("You swing the axe into this back. He roars in pain." +
						"\nHe spins, throwing you back. The axe remains planted in his back.")
					axei.Used = true
					models.Itemupdate(axei)
					cm.Health = cm.Health - 40
					cm.Position = cp.Position
				default:
					fmt.Println("If only you had something sharp. You settle for punching him in the back of the head." +
						"\nHe yells and turns around, ready to fight.")
					cm.Health = cm.Health - 15
					cm.Position = cp.Position
				}
				item1.Toggle = false
			default:
				fmt.Println("You stay hidden and safe until the monster leaves. Then emerge.")
			}
		case false:
			fmt.Println("The chainsaw roars to life. He drives it through the bed frame and into your back.")
			cp.Health = 0
		}
	case 2: //wolf finds you. bites your leg.
		fmt.Println("You see the wolf enter the room. Its dripping wet nose sniffs the air." +
			"\nIt walks around slowly. It is still in the room but out of your sight. You hold your breath." +
			"\nThe bed jumps. Twisting, you see the wolf is coming under the bed behind you." +
			"\nYou crawl to get away. Your halfway out when the teeth sink into your leg." +
			"\nIt hurts. With your free leg, you kick frantically. A lucky hit and you are free. You run into the hallway.")
		cp.Position = 3
		cp.Health = cp.Health - 30
		cm.Health = cm.Health - 10
		cm = Mrun(cp, cm)
	}
	models.Monsterupdate(cm)
	return cp, item1
}

func outside(cp models.Player) models.Player {
	cm := models.Chosenmonsterget()
	switch cm.Spawn {
	case true:
		switch cm.Number {
		default:
			fmt.Println("You run for hours. In a dark forest. Alone and in panic. You stop to catch your breath." +
				"\nThe monster catches you.")
			cp.Health = 0
		}
	case false:
		fmt.Println("You wander around for hours in the forest at night.\nBasking in the alien beauty of nature.\nIn all that time, you never meet another soul.\nAnd you return to your cabin.")
		cp.Charisma += 20
	}
	return cp
}

func cook(cp models.Player, item1 models.Item) (models.Player, models.Item) {
	bi := models.ItemGet("burger meat")
	gsi := models.ItemGet("glass shards")
	ski := models.ItemGet("skillet")
	cri := models.ItemGet("silver cross")
	cm := models.Spawnmonsterget()
	switch {
	case bi.Loc == 20 && ski.Loc == 20 && ski.Used == false && bi.Toggle == true:
		switch cm.Number {
		case 2:
			fmt.Println("You have a skillet, range and raw meat to cook up some burgers.\nBefore you start cooking do you want to add anything?")
			fmt.Println("1. Yes\n2. No")
			r1 := inputs.Basicinput("?")
			if r1 == 1 {
				switch {
				case cri.Loc == 20:
					fmt.Println("Would you like to add the old silver cross you found?")
					fmt.Println("1. Yes\n2. No")
					r2 := inputs.Basicinput("?")
					if r2 == 1 {
						fmt.Println("You insert the old silver cross into the raw meat.")
						cri.Loc = 30 //placing it in the burger
					}
				case gsi.Loc == 20:
					fmt.Println("Would you like to add tasty shards of glass?")
					fmt.Println("1. Yes\n2. No")
					r2 := inputs.Basicinput("?")
					if r2 == 1 {
						fmt.Println("You insert shards of glass into the raw meat.")
						gsi.Loc = 30
					}
				default:
					fmt.Println("You are not carrying anything that goes in burgers.")
				}
			}
			switch cm.Spawn {
			case false:
				fmt.Println("The smell fills the room and you open the window to ventilate.")
				fmt.Println("As your teeth sink into the soft meat, you hear a long howl.")
				fmt.Println("You pause and set the burger down.")
				ski.Toggle = false //to trigger wolf
				bi.Toggle = false  //to indicate burger is cooked
				bi.Loc = 2         //to place burger down
			case true:
				cm.Position = 2
				fmt.Println("The smell fills the room. You hear something behind you.")
				models.Monsterupdate(cm)
				bi.Toggle = false
			}
			ski.Loc = 2

		default:
			fmt.Println("Combining skillet, range and raw meat you cook up some burgers. They're super bloody.")
			fmt.Println("That was a tasty burger. You feel a bit better. And nothing attacked you this whole time.")
			cp.Health = cp.Health + 15
			ski.Loc = 2
			bi.Loc = 10
		}
	case bi.Toggle == false:
		fmt.Println("This meat is already cooked.")
	case bi.Loc == 20:
		fmt.Println("You need something to cook with. Unless you want to hold the raw meat over the range in your hand?")
		fmt.Println("1. Yes\n2. No")
		r1 := inputs.Basicinput("?")
		if r1 == 1 {
			fmt.Println("Your hand is on fire. And the meat is still raw. Stupid idea.")
			cp.Health = cp.Health - 25
		}
	default:
		fmt.Println("You are not ready to cook. Unless you want to eat yourself?")
		fmt.Println("1. Yes\n2. No")
		r1 := inputs.Basicinput("?")
		if r1 == 1 {
			fmt.Println("Ouch. That hurts. Stupid idea.")
			cp.Health = cp.Health - 10
		}
	}
	models.Itemupdate(bi)
	models.Itemupdate(gsi)
	models.Itemupdate(ski)
	item1 = models.ItemGet(item1.Name)
	return cp, item1
}

func fire(cp models.Player) models.Player {
	cm := models.Spawnmonsterget()
	if cm.Spawn == false { //for scenario where player burns down the house before a monster appears
		fmt.Println("The fire spreads rapidly. The whole room is ablaze.\nYou have burned down a cabin.\nWhy? Unless...\nYou're the monster. Wasn't your family in there?")
		cp.Health = 1
		cp.Continue = false
		cp.Charisma = 999
		return cp
	}
	switch cp.Position {
	case 1:
		fmt.Println("The fire spreads rapidly. The whole room is ablaze. You can barely see the doorframe through the smoke.")
		fmt.Println("Run for the door?\n1. Yes\n2. No")
		r1 := inputs.Basicinput("?")
		switch r1 {
		case 1:
			fmt.Println("Running, you barely make it outside before the roof collapses. You cough to clear the smoke from your lungs. You watch the cabin burn.")
			switch cm.Number {
			case 1:
				fmt.Println("The fire transforms the cabin into a bonfire. There's no way the monster could survive it." +
					"\nYou start walking away. Plotting how to get home. Then you hear the terrible chainsaw." +
					"\nA figure wrapped in burning flannel emerges from the cabin. Chainsaw raised." +
					"\nHe collapses to the ground.")
				cm.Health = 0
				cp.Continue = false
				fmt.Println(cm.Outrom)
			case 2:
				fmt.Println("The fire transforms the cabin into a bonfire. There's no way the monster could survive it." +
					"\nYou start walking away. Plotting how to get home. Then you hear the terrible howl." +
					"\nThe source emerges from the smoke and flame. It runs. Not in panic, but fast, fierce and straight towards you." +
					"\nFire doesn't stop werewolves. It lunges. Teeth wrapped in smoke and burning hair.")
				//fmt.Println(cm.Outrop) Not needed. Need the monster death text but not the player
				cp.Health = 0
			default:
				fmt.Println("Nothing specific happens. Oddly. The fire does not kill the monster. Monster kills you. Game over!")
				cp.Health = 0
			}
		default:
			fmt.Println("A timber from the roof falls and traps you." +
				"\nThrough the smoke and flame, the monster appears.")
			cp.Health = 0
		}
	default:
		fmt.Println("The fire spreads. Soon the whole room is ablaze.\nAnd you are in it. A timber from the roof falls and traps your leg." +
			"\nThrough the smoke and flame, the monster appears.")
		cp.Health = 0

	}
	return cp
}
