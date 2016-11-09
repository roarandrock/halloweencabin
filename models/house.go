package models

//Room defines the rooms
type Room struct {
	RoomID int
	Name   string
	Adj    []int //Adjacent rooms array
	Desc   string
}

/*RoomIDs
Living Rooom 1
kitchen 2
hallway 3
Master bedroom 4
Guest bedroom 5
Bathroom 6
Basement (locked?) 7
*/

var (
	livingroom = Room{1, "Living Room", []int{2, 3}, "is so cozy you could just die."}
	kitchen    = Room{2, "Kitchen", []int{1}, "smells like butchered meat. Old butchered meat."}
	hallway    = Room{3, "Hallway", []int{1, 4, 5, 6}, "is poorly lit and the connected rooms could contain anything."}
	mbedroom   = Room{4, "Master Bedroom", []int{3, 5}, "is full of other people's memories and stains on the sheets."}
	gbedroom   = Room{5, "Guest Bedroom", []int{3, 4}, "is cramped and musty. Like a bad memory of an old relative."}
	bathroom   = Room{6, "Bathroom", []int{3}, "has a mirror.\nIn the mirror you cannot see yourself, just blurry shapes shifting in the flickering light."}
	//add basement with zombies later, need a room updated similar to monsters/items, or can add a door item to be checked
)

var roommap = map[int]Room{
	livingroom.RoomID: livingroom,
	kitchen.RoomID:    kitchen,
	hallway.RoomID:    hallway,
	mbedroom.RoomID:   mbedroom,
	gbedroom.RoomID:   gbedroom,
	bathroom.RoomID:   bathroom,
}

//RoomGet grabs current item by number
func RoomGet(r int) Room {
	i := roommap[r]
	return i
}
