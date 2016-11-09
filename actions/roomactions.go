package actions

import (
	"fmt"
	"halloweencabin/inputs"
	"halloweencabin/models"
)

//Lookr looks around the room
func Lookr(cp models.Player) models.Player {

	cr := models.RoomGet(cp.Position)
	fmt.Println("You are in the ", cr.Name)
	fmt.Println("It", cr.Desc)

	stuff, q := itinroom(cp.Position)
	if q == 0 {
		fmt.Println("Nothing else interesting here.")
	}
	if q > 0 {
		fmt.Println("There are some random things here you can use:")
		for _, v := range stuff {
			fmt.Println(v)
		}
	}

	stuff2, q2 := Itonperson()
	if q2 > 0 {
		fmt.Println("You are carrying:")
		for _, v := range stuff2 {
			fmt.Println(v)
		}
	}

	//health check
	healthcheck(cp)
	return cp
}

//Itinroom checks for items in room and on person and returns array of items
func itinroom(rid int) ([]string, int) {
	q := 0 //number of items in room
	m := models.Inventory()
	vlist := [10]string{}
	i := 0
	for _, v := range m {
		if v.Used == false {
			if v.Loc == rid {
				vlist[i] = v.Name
				i++
				q++
			}
		}
	}
	if q == 0 {
		empty := []string{}
		return empty, q
	}
	stuff := vlist[0:i]
	return stuff, q
}

//Itonperson returns items carried by the player
func Itonperson() ([]string, int) {
	q := 0 //number of items in room
	m := models.Inventory()
	vlist := [10]string{}
	i := 0
	for _, v := range m {
		if v.Used == false {
			if v.Loc == 20 {
				vlist[i] = v.Name
				i++
				q++
			}
		}
	}
	if q == 0 {
		empty := []string{}
		return empty, q
	}
	stuff := vlist[0:i]
	return stuff, q
}

//Healthcheck returns strings describing player's health
func healthcheck(cp models.Player) {
	//fmt.Println(cp.Health) for testing
	h := cp.Health
	switch {
	case h < 15:
		fmt.Println("You can barely keep your eyes open. You're hanging on by a narrow thread.\nDeath is near.\nDon't give up.")
	case h < 30:
		fmt.Println("You look awful. Blood everywhere. You are in dire need of medical attention.\nBut there's a monster to be deal with first.")
	case h < 60:
		fmt.Println("You are bleeding and not looking great. Time to focus and get even.")
	case h < 100:
		fmt.Println("You are hurt but it's nothing serious. Assuming you survive the night.")
	case h == 100:
		fmt.Println("And you are looking completely healthy. Way to go.")
	}
}

//Move options
func Move(cp models.Player) models.Player {

	cr := models.RoomGet(cp.Position)
	l := len(cr.Adj)
	if l == 0 {
		fmt.Println("Nowhere to run.")
		return cp
	}
	fmt.Println("Where would you like to go?")
	adjr := models.Room{}
	for i := 0; i < l; i++ {
		k := cr.Adj[i] //can be simplified, broken up for testing
		adjr = models.RoomGet(k)
		fmt.Println(i+1, ":", adjr.Name)
	}
	s1 := "?"
	r1 := inputs.Basicinput(s1)
	if r1 > 0 && r1 <= l { //checks that room is valid
		cp.Position = cr.Adj[r1-1]
	}
	cm := models.Spawnmonsterget() //checks if spawned monster is met
	if cm.Spawn == true {
		cp = Monstercheck(cp)
	}
	return cp
}

//Usecheck items
func Usecheck(cp models.Player) models.Player {
	stuff, q := itinroom(cp.Position)
	stuff2, q2 := Itonperson()
	q3 := q2 + q
	if q3 == 0 {
		fmt.Println("There is a scarcity of things to use.")
		return cp
	}

	if q > 0 {
		fmt.Println("In this room you see:")
		for i, v := range stuff {
			fmt.Println(i+1, v)
		}
	}
	if q2 > 0 {
		a1 := "You"
		if q > 0 {
			a1 = "And you"
		}
		fmt.Println(a1, " are carrying:")
		for i2, v2 := range stuff2 {
			fmt.Println(i2+1+q, v2)
		}
	}
	s1 := "Which would you like to use?"
	r1 := inputs.Basicinput(s1)
	if r1 == 0 || r1 > q3 {
		fmt.Println("What? that's not here.")
		return cp
	}
	upitem := models.Item{}
	if r1 <= q {
		upitem = models.ItemGet(stuff[r1-1])
	} else {
		upitem = models.ItemGet(stuff2[r1-1-q])
	}

	if upitem.Carriable == true {
		if upitem.Loc != 20 {
			fmt.Println("Take the", upitem.Name)
			fmt.Println("1. Yes\n2. No")
			r3 := inputs.Basicinput("?")
			if r3 == 1 {
				upitem.Loc = 20
				fmt.Println("You are now carrying the", upitem.Name)
				models.Itemupdate(upitem)
			}
		}
	}
	fmt.Println("Would you like to:")
	i4 := 0 //for last option
	for i3, v3 := range upitem.Action {
		fmt.Println(i3+1, ":", v3, "the", upitem.Name)
		i4 = i3
	}
	fmt.Println(i4+2, ": None of that")
	r2 := inputs.Basicinput("?")
	upitem, cp = Useitem(upitem, r2-1, cp)
	models.Itemupdate(upitem)
	cm := models.Spawnmonsterget() //checks if spawned monster is met
	if cm.Spawn == true {
		cp = Monstercheck(cp)
	}
	return cp
}
