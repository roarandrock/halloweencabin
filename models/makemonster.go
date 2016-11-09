package models

import (
	"math/rand"
	"time"
)

var m int        //monster location
var p int        //phone location
var r int        //numbers of room
var moptions int //number of monsters
var monster int  //monster chosen

//Mcreate sets the monster
func Mcreate(np Player) {
	//choosing the monster
	moptions = len(mmap)
	s3 := rand.NewSource(time.Now().UnixNano())
	r3 := rand.New(s3)
	monster = r3.Intn(moptions) + 1
	//choosing the monster location
	r = len(roommap)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	m = r1.Intn(r) + 1
	for m == np.Position {
		m = r1.Intn(r) + 1
	}
}

//Mrunner runs the monster to a new room
func Mrunner(np Player) int {
	//choosing the monster location
	r = len(roommap)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	m1 := r1.Intn(r) + 1
	for m1 == np.Position {
		m1 = r1.Intn(r) + 1
	}
	return m1
}

//Mpos returns roomid for monster
func Mpos() int {
	return m
}

//Pcheck checks if phone is in the room
func Pcheck() int {
	return p
}

//Mchoice returns chosen monster
func Mchoice() int {
	return monster
}

//Madj returns random number (0 to l) for adjacency check
func Madj(l int) int {
	s4 := rand.NewSource(time.Now().UnixNano())
	r4 := rand.New(s4)
	adjr := r4.Intn(l)
	return adjr
}
