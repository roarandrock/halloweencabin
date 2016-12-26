package models

//Item defines item stats
type Item struct {
	Name      string   //name of the item
	Loc       int      //RoomID, on player, on monster
	Carriable bool     //True for action1, false for action2 time
	Action    []string //
	Used      bool     //true for used, false for not used. Can remove from game when used
	Toggle    bool     //true initially, can be switched. Trying to us for umbrella opening, closing
}

/*RoomIDs
Living Rooom 1
kitchen 2
hallway 3
Master bedroom 4
Guest bedroom 5
Bathroom 6
Basement (locked?) 7 - not added yet
Items to be added by events 10
20 player
for cooked items:
30 in the burger
40 used and done with
*/

//items to be found and carried. Also fixed room items. Defaults

var imap = map[string]Item{}

//Iset sets defaults, needs to update outside of the function, otherwise, does not take
func Iset(np Player) {
	//defaults
	phone := Item{"phone", 1, true, []string{"Call 911 on", "Call family on", "Call friends on"}, false, true}
	umbrella := Item{"umbrella", 1, true, []string{"Open/close", "Dance with"}, false, true}
	frontdoor := Item{"front door", 1, false, []string{"Open/close", "Go outside via"}, false, true}
	destfrontdoor := Item{"ruins of front door", 10, false, []string{"Go outside via", "Collect scrap wood from"}, false, true}
	scrapwood := Item{"scrap wood", 10, true, []string{"Build a fire with"}, false, true}
	skillet := Item{"skillet", 2, true, []string{"Cook with"}, false, true}
	fridge := Item{"fridge", 2, false, []string{"Open"}, false, true}
	vodka := Item{"vodka bottle", 10, true, []string{"Drink", "Smash"}, false, true}
	glass := Item{"glass shards", 10, true, []string{"Eat"}, false, true}
	burger := Item{"burger meat", 10, true, []string{"Cook", "Eat"}, false, true}
	axe := Item{"axe", 4, true, []string{"Chop wood with"}, false, true} //need something to chop?
	trap := Item{"bear trap", 4, false, []string{"Set"}, false, true}
	bed := Item{"bed", 4, false, []string{"Sleep in", "Hide under"}, false, true} //surprise the psycho?
	cross := Item{"silver cross", 5, true, []string{"Eat"}, false, true}
	plunger := Item{"red plunger", 6, true, []string{"Plunge with"}, false, true}
	razor := Item{"rusty straight razor", 6, true, []string{"Shave with"}, false, true}
	bell := Item{"silver bell", 4, false, []string{"Ring"}, false, true}
	present := Item{"present", 10, false, []string{"Open", "Set on fire"}, false, true}

	//for creating initial map
	Itemupdate(phone)
	Itemupdate(umbrella)
	Itemupdate(frontdoor)
	Itemupdate(destfrontdoor)
	Itemupdate(scrapwood)
	Itemupdate(skillet)
	Itemupdate(fridge)
	Itemupdate(vodka)
	Itemupdate(glass)
	Itemupdate(burger)
	Itemupdate(axe)
	Itemupdate(trap)
	Itemupdate(bed)
	Itemupdate(cross)
	Itemupdate(plunger)
	Itemupdate(razor)
	Itemupdate(bell)
	Itemupdate(present)
}

//Inventory returns item map
func Inventory() map[string]Item {
	return imap
}

//Itemlocation takes current inventory and returns locations in a slice
func Itemlocation() []int {
	m := Inventory()
	a := [10]int{}
	i := 0
	for _, v := range m {
		a[i] = v.Loc
		i++
	}
	iloc := a[0:i]
	return iloc
}

//ItemGet grabs current item by name, can simplify
func ItemGet(s string) Item {
	m := Inventory()
	i := m[s]
	return i
}

//Itemupdate updates items in the  inventory map
func Itemupdate(upitem Item) {
	m := Inventory()
	m[upitem.Name] = upitem
}
