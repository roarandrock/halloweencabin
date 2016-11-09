package models

import "fmt"

//Monster defines stats
type Monster struct {
	Name     string
	Health   int
	Charisma int
	Position int
	Spawn    bool //exists in the house
	Chosen   bool //exists in the game
	Number   int  //unique number for selecting chosen
	Intro    string
	Outrom   string //for monster death
	Outrop   string //for player death
	Met      bool   //for after player meets
	Found    string //text when player encounters the monster
}

var (
	//psycho
	i1 = "He stands in the wreckage of your cabin door. A bearded man in red flannel." +
		"\nA fever in his eyes. And a massive chainsaw in his arms.\nHe doesn't say a word, just strides towards you."
	om1 = "He is down. He doesn't get up. He was only a man and you defeated him.\nHis lust for killing cannot overcome his mortality."
	op1 = "The chainsaw rips through your flesh. Bones shred beneath the spinning blades. You can feel it emerging again on the other side." +
		"\nThere is no recovering from this. No more fight. He has won."
	f1 = "You see the man. He sees you. The chainsaw roars to life."
	//werewolf
	i2 = "You see a creature. Its red eyes stare out from coarse gray hair. Saliva drips from long teeth.\nStands on two legs but keeps its head low." +
		"\nIts heavy panting breath breaks for a long howl.\nHuman and wolf mashed into one unholy monstrosity."
	om2 = "The creature lunges at you desperately. You dodge it easily and turn ready to continue the fight." +
		"\nThe creature is laying on the floor. Blood pooling beneath it. Its heavy panting breath slows then stops."
	op2 = "You try to hold the monster back and you fail. Its teeth sink into your shoulder. Mercifully you lose conciousness as it starts to eat your face."
	f2  = "There in the shadows, the wolf. The howl is close, loud and terrifying."
)

var (
	psychopath = Monster{"dude with a chainsaw", 100, 75, 1, false, false, 1, i1, om1, op1, false, f1} //items do different damage, so don't mess with health now
	werewolf   = Monster{"werewolf", 100, 10, 1, false, false, 2, i2, om2, op2, false, f2}
	//tiny respawning aliens - critters
	//demon, hellspawn,zombie
	chosenmonster = Monster{}
)

var mstart = map[string]Monster{
	psychopath.Name: psychopath,
	werewolf.Name:   werewolf,
	//chosenmonster.Name: chosenmonster,
}

var mmap = map[string]Monster{
	psychopath.Name: psychopath,
	werewolf.Name:   werewolf,
	//chosenmonster.Name: chosenmonster,
}

//Monstermap returns map
func Monstermap() map[string]Monster {
	return mmap
}

//Monsterget grabs monster by name
func Monsterget(name string) Monster {
	cm := chosenmonster
	if name != cm.Name {
		fmt.Println("error in Monsterget")
	}
	return cm
}

//Monsterdefault grabs monster by name
func Monsterdefault(name string) Monster {
	m := Monstermap()
	cm := m[name]
	return cm
}

//Spawnmonsterget grabs current spawned monster
func Spawnmonsterget() Monster {
	/*m := Monstermap()
	a := [1]string{} //only for one monster
	i := 0
	cm := Monster{}
	for _, v := range m {
		if v.Spawn == true {
			a[i] = v.Name
			i++
		}
	}
	cm = m[a[0]]*/
	cm := chosenmonster
	return cm
}

//Chosenmonsterget grabs monster chosen for game
func Chosenmonsterget() Monster {
	/*m := Monstermap()
	a := [1]string{} //only for one monster
	i := 0
	cm := Monster{}
	for _, v := range m {
		if v.Chosen == true {
			a[i] = v.Name
			i++
		}
	}
	cm = m[a[0]]*/
	cm := chosenmonster
	return cm
}

//Monsterupdate updates items in the  inventory map
func Monsterupdate(upm Monster) {
	//m := Monstermap()
	//m[upm.Name] = upm
	chosenmonster = upm
}

//Monsterhealth prints strings detailing the monster's state
func Monsterhealth(cm Monster) {
	//fmt.Println(cp.Health) for testing
	h := cm.Health
	switch cm.Number {
	case 1:
		switch {
		case h < 15:
			fmt.Println("Limping and blood soaked, he is a sliver from death. Nothing less is going to stop this monster.")
		case h < 40:
			fmt.Println("He is seriously damaged. Anyone else would have retreated. But he does not stop.")
		case h < 70:
			fmt.Println("He is wounded and frustrated by your resilience.")
		case h < 100:
			fmt.Println("He has a few small wounds. Nothing serious.")
		case h == 100:
			fmt.Println("Past the chainsaw, you notice he looks insanely strong and uninjured. ")
		}
	case 2:
		switch {
		case h < 30:
			fmt.Println("Its fur is soaked in blood. Its movements are no longer graceful and powerful. The beast is wounded.")
		case h < 70:
			fmt.Println("The beast moves quickly and attacks fiercely. But you see its bleeding and hear its breath is ragged.")
		default:
			fmt.Println("The beast stalks the room. It can attack at any moment.")
		}
	default:
		fmt.Println("this monster is inscrutable.")
	}
}

//Mset sets the defaults for the monster, need to only set up stats, location after spawn
func Mset(np Player) {
	Mcreate(np)
	mpos := Mpos()       //random position in house
	mchoose := Mchoice() // returns int for randomly chosen monster
	var cname string
	for _, v := range mstart {
		if v.Number == mchoose {
			cname = v.Name
		}
	}
	newm := Monsterdefault(cname)
	newm.Position = mpos
	newm.Chosen = true
	Monsterupdate(newm)
}
