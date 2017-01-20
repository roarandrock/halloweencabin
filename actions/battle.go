package actions

import (
	"fmt"
	"halloweencabin/inputs"
	"halloweencabin/models"
)

var ( //for tweaking battles, light medium and heavy damage
	d1 = 15
	d2 = 30
	d3 = 49
)

//Checkposition checks current position in the environment against the map
func battle(cp models.Player, cm models.Monster) (models.Player, models.Monster) {
	stuff, q := Itonperson() //array of items (by name) and total amount, on the player only. Not in the room
	fmt.Println("1.Use\n2.Talk\n3.Run")
	s1 := "?"
	switch r1 := inputs.Basicinput(s1); {
	case r1 == 1:
		if q == 0 {
			fmt.Println("You are not carrying anything you can use.")
			return cp, cm
		}
		fmt.Println("What do you want to use?")
		for i := 0; i < q; i++ {
			fmt.Println(i+1, stuff[i])
		}
		r2 := inputs.Basicinput(s1)
		iname := stuff[r2-1]
		item1 := models.ItemGet(iname)
		cp, cm = Useitembattle(item1, cp, cm)
	case r1 == 2:
		fmt.Println("1.Hello\n2.Who are you?\n3.What are you doing here?\n4.Why?")
		r2 := inputs.Basicinput(s1)
		switch cm.Number {
		case 1: //for pyscho, he doesn't care to talk
			fmt.Println("He lets his chainsaw do the talking. It says \"Fuck You\" and cuts your arm.")
			cp.Health = cp.Health - d1
		case 2:
			fmt.Println("The wolf cocks its head to the side. Like it's listening.")
			fmt.Println("Cautiously extend your hand?\n1.Palm up\n2.Palm down\n3.No, do not do this")
			r3 := inputs.Basicinput("?")
			if r3 == 1 || r3 == 2 {
				fmt.Println("You place your hand out. The wolf sniffs it. Licks it. Eats it. You have one less hand now.")
				cp.Health = cp.Health - d3
			} else {
				fmt.Println("The wolf licks its lips and stares at you.")
			}
		case 3:
			switch r2 {
			case 1:
				fmt.Println("\"Greetings\"")
			case 2:
				fmt.Println("\"Ho ho ho\"")
			case 3:
				fmt.Println("\"Tis the season.\"")
			case 4:
				fmt.Println("\"You know why.\"")
			default:
				fmt.Println("\"Ho ho ho\"")
			}
		default:
			fmt.Println("This monster has nothing to say. He looks at you. Oddly.")
		}
	case r1 == 3:
		fmt.Println("You run away!")
		cp = Move(cp)
		switch cm.Number {
		default:
			fmt.Println("As you flee the room the", cm.Name, "cuts you down the back.")
			cp.Health = cp.Health - d1
			if cp.Health <= 0 {
				return cp, cm
			}
			fmt.Println("You are hurt but he does not follow you into the adjacent room.\nAt first you are relieved.\nIn the silence, you realize that you cannot see or hear your adversary.\nHe could be anywhere.")
			cm = Mrun(cp, cm)
		}
	}
	return cp, cm
}

//Useitembattle is a giant switch for using different items in battle. Defaults to no effect
func Useitembattle(item1 models.Item, cp models.Player, cm models.Monster) (models.Player, models.Monster) {

	switch {
	case item1.Name == "phone":
		fmt.Println("You throw the phone at the monster.")
		switch cm.Number {
		case 1:
			fmt.Println("The busted phone falls to the floor in pieces.\nHe is bruised and he takes his revenge with his chainsaw.\nYou are cut.")
			cm.Health = cm.Health - d1
			cp.Health = cp.Health - d2
			item1.Used = true
		case 2:
			fmt.Println("The wolf snatches it from the air with its jaws. It chews then spits the remains on the floor. The plastic never stood a chance.")
			item1.Used = true
		case 3:
			fmt.Println("He catches the phone and stuffs it into his sack.")
			item1.Used = true
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "umbrella":
		switch cm.Number {
		case 1:
			fmt.Println("Umbrella attack!\nYou poke him right in the eye.\nFrustrated, he runs away.")
			cm = Mrun(cp, cm)
		case 2:
			fmt.Println("The wolf claws at the umbrella, tearing it to shreds.\nYou have lost an umbrella. And gained some fresh wounds from the wolf's claws.")
			item1.Used = true
			cp.Health = cp.Health - d1
		case 3:
			fmt.Println("Umbrella attack!\nYou poke him right in the eye.\nFrustrated, he runs away.")
			cm = Mrun(cp, cm)
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "scrap wood":
		switch cm.Number {
		case 1:
			fmt.Println("You swing at him with your chunk of wood. Chainsaw is wood's natural enemy.\nThe chainsaw destroys your wood and ravages your arm.")
			item1.Loc = 10 //not used, can get more.
		case 2:
			fmt.Println("You try to stab the wolf with the wood. Wolves do not fear wood. You manage to cut the beast but are cut by its claws.")
			cp.Health = cp.Health - d1
			cm.Health = cm.Health - d1
		case 3:
			fmt.Println("You swing at him with your chunk of wood. It breaks across his face.\nHe grins, it didn't hurt him at all.")
			item1.Loc = 10 //not used, can get more.
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "skillet":
		switch cm.Number {
		case 1:
			fmt.Println("You wallop him good with the skillet." +
				"\nHe stumbles back, seriously hurt.\nThen in a rage he savagely swings the chainsaw.\nYou bring your skillet to block" +
				" and the old cookery is destroyed.")
			cm.Health = cm.Health - d2
			item1.Used = true
		case 2:
			fmt.Println("You swing your trusty skillet into action. The wolf's claws cannot get past your cookery defense." +
				"\nYou cannot hurt the beast but you frustrate it. It slinks off, disappearing into the shadows.")
			cm = Mrun(cp, cm)
		case 3:
			fmt.Println("You swing your trusty skillet into action. The butcher knife is deflected." +
				"\nYou cannot hurt the man but you frustrate him. He disappears with a twinkle of his eye.")
			cm = Mrun(cp, cm)
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "vodka bottle":
		switch cm.Number {
		case 1:
			fmt.Println("With deft manuevering you get behind him and smash the vodka bottle on his head." +
				"\nShards of glass scatter on the floor.") //create glass shards
			cm.Health = cm.Health - d1
			gsi := models.ItemGet("glass shards")
			gsi.Loc = cp.Position
			models.Itemupdate(gsi)
			item1.Used = true
		case 2:
			fmt.Println("You chuck the vodka bottle at the wolf. It snatches it out of the air with a powerful bite." +
				"\nThe bottle explodes. Shards of glass scatter on the floor. The wolf smells and licks the floor." +
				"It yelps in pain and runs away. Either wolves don't like vodka or eating shards of glass.") //create glass shards
			cm.Health = cm.Health - d1
			gsi := models.ItemGet("glass shards")
			gsi.Loc = cp.Position
			models.Itemupdate(gsi)
			cm = Mrun(cp, cm)
			item1.Used = true
		case 3:
			fmt.Println("You swing the bottle at him. Deftly he intercepts and takes the bottle from you." +
				"\nHe steps back, takes a big swig of vodka and then stuffs the bottle away in his sack." +
				"\n\"Ho, ho, burp, ho.\"")
			item1.Used = true
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "burger meat":
		switch cm.Number {
		case 1:
			fmt.Println("You throw the meat at the man. It splatters across his face.\nHe pulls a piece from his beard, plops it in his mouth and grins.")
			item1.Loc = 10
			item1.Toggle = true
		case 2:
			switch item1.Toggle {
			case true:
				fmt.Println("You throw the raw meat at the wolf. It ignores the cold burger meat as it splatters on the floor.\nEither it doesn't like beef or it prefers burgers cooked.")
				item1.Loc = 10
			case false:
				fmt.Println("You throw the cooked burger at the wolf. It snatches it out of the air and literally wolfs it down.")
				gsi := models.ItemGet("glass shards")
				cri := models.ItemGet("silver cross")
				if gsi.Loc == 30 {
					fmt.Println("The wolf begins to cough. It falls and rolls around the floor. Blood and glass drip from it's jaws." +
						"\nYou've seriously hurt the beast.")
					cm.Health = cm.Health - d3
					gsi.Used = true
					gsi.Loc = 40
					models.Itemupdate(gsi)
				}
				if cri.Loc == 30 {
					fmt.Println("The wolf stares at you. It stops moving entirely. A drop of blood falls from its lips." +
						"\nIt opens its mouth and a torrent of blood rushes out. It shakes uncontrollably, vomiting blood across the floor." +
						"\nSlowly it regains control and limps out of the room.")
					cm.Health = cm.Health - d3 - d2
					cri.Used = true
					cri.Loc = 40
					models.Itemupdate(cri)
					cm = Mrun(cp, cm)
				}
				item1.Loc = 10
				item1.Toggle = true
			}
		case 3:
			fmt.Println("You throw the meat at the man. It splatters across his face.\nHe pulls a piece from his beard, plops it in his mouth and grins.")
			item1.Loc = 10
			item1.Toggle = true
			cm.Health = cm.Health + 15
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "glass shards":
		switch cm.Number {
		case 1:
			fmt.Println("A shard of glass is a poor weapon. You cut yourself. The man laughs.")
			cp.Health = cp.Health - d1
		case 2:
			fmt.Println("A shard of glass is a poor weapon. You cut yourself." +
				"\nThe wolf licks it's massive jaws at the sight of warm blood.")
			cp.Health = cp.Health - d1
		case 3:
			fmt.Println("A shard of glass is a poor weapon. You cut yourself.")
			cp.Health = cp.Health - d1
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "axe":
		switch cm.Number {
		case 1:
			fmt.Println("Your axe is a fantastic tool for chopping wood. But it's been replaced by the chainsaw." +
				"\nIt's also been replaced in combat. You hold the axe menacingly and he keeps his distance." +
				"\nBut you cannot reach him now, only if you caught him unaware.")
		case 2:
			fmt.Println("You swing the axe and it catches the wolf in the arm. It hurts the beast. The wolf grabs the handle with its jaws." +
				"\nThe wood shatters and there is no more axe.")
			cm.Health = cm.Health - d1
			item1.Used = true
		case 3:
			fmt.Println("You swing the axe. It cuts the man." +
				"\nHe glares at you. Before you can react, he grabs and pulls the weapon away. It disappears into his sack.")
			cm.Health = cm.Health - d1
			item1.Used = true
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "silver cross":
		switch cm.Number {
		case 1:
			fmt.Println("You dangle the cross in the air. The man is stunned by the sudden religious icon." +
				"\nThen he snaps out of it and slashes you with his chainsaw.")
			cp.Health = cp.Health - d1
		case 2:
			fmt.Println("You dangle the silver cross in the air. The wolf is not afraid. It slashes at you." +
				"\nIt's paw connects with the cross and instantly starts burning. Like acid. The wolf howls and runs away." +
				"\nIt happened instantly, and you are not sure if it hurt the beast. Maybe if it had prolonged contact.")
			cm = Mrun(cp, cm)
		case 3:
			fmt.Println("You dangle the cross in the air. The man smiles at the sight of the icon." +
				"\n\"Wrong day.\"" +
				"\nHe cuts you with his butcher knife.")
			cp.Health = cp.Health - d1
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "red plunger":
		switch cm.Number {
		case 1:
			fmt.Println("Plop. You stick it right on his face. Hilarious." +
				"\nHe tears it off, taking skin with it, and breaks it over his knee.")
			cm.Health = cm.Health - d1
			cp.Charisma = cp.Charisma + 30
			item1.Used = true
		case 2:
			fmt.Println("Plop. You stick it right on the wolf's snout. Hilarious." +
				"\nThe wolf freaks out. Runs around the room trying to dislodge it. The smell must be terrible." +
				"\nEventually it frees itself. It gets revenge on the plunger and destroys it.")
			cm.Health = cm.Health - d1
			cp.Charisma = cp.Charisma + 30
			item1.Used = true
		case 3:
			fmt.Println("Plop. You stick it right on his face. Hilarious." +
				"\nHe tears it off, taking skin with it, and breaks it over his knee.")
			cm.Health = cm.Health - d1
			cp.Charisma = cp.Charisma + 30
			item1.Used = true
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "rusty straight razor":
		switch cm.Number {
		case 1:
			fmt.Println("You display the razor menacingly. He revs his chainsaw." +
				"\nYou try to cut him but it's impossible to get past the chainsaw. You almost lose an arm and luckily are only grazed.")
			cp.Health = cp.Health - d1
		case 2:
			fmt.Println("The beast could use a shave. But it's not letting you anywhere near it." +
				"\nIts claws draws blood from your arm.")
			cp.Health = cp.Health - d1
		case 3:
			fmt.Println("You display the razor menacingly. He displays a butcher knife." +
				"\nYou try to cut him but it's impossible to get past the knife. You almost lose an arm and luckily are only grazed.")
			cp.Health = cp.Health - d1
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
	case item1.Name == "heart":
		switch cm.Number {
		case 3:
			fmt.Println("You display the heart. The man pauses, unsure of what to do." +
				"\n\"Giving that to me?\"")
			fmt.Println("1. Yes\n2. No")
			rh := inputs.Basicinput("?")
			switch rh {
			case 1:
				fmt.Println("He takes the heart. Then plants his butcher knife in your skull." +
					"\n\"Presents are not returned. Rude.\"")
				cp.Health = cp.Health - d3
				item1.Used = true
			case 2:
				fmt.Println("You place the heart back in your pocket. Poor pocket.")
			default:
				fmt.Println("You place the heart away.")
			}
		default:
			fmt.Println(item1.Name, "Has no impact on this monster.")
		}
		/*
							case item1.Name == "":
							switch cm.Number{
						case 1:
						fmt.Println("")
					case 2:
					fmt.Println("")
				default:
				fmt.Println(item1.Name, "Has no impact on this monster.")
			}
		*/

	default:
		fmt.Println("It has no effect. Bad plan.")
	}
	models.Itemupdate(item1)
	return cp, cm
}

//Mrun moves the monster to any location that is not where the player is
func Mrun(cp models.Player, cm models.Monster) models.Monster {
	if cm.Health == 0 {
		fmt.Println("The monster does not get far. Overwhelmed he collapses.")
		return cm
	}
	cm.Position = models.Mrunner(cp)
	return cm
}
