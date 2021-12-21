package main
import (
	//"time"
	//"math/rand"
//	"fmt"
)

func make_rovers()  {
	//rand.Seed(time.Now().UTC().UnixNano())
	var rover Rover	
	for i := 0; i < NUM_ROVERS; i++ {
		rover.brain = make_brain()
		rover.Fitness = 0
		rover.Dead = false
		//rover.Xpos = getRandomInt(0,arena.Width)
		//rover.Ypos = getRandomInt(0,arena.Height)
		rover.Xpos = getRandomInt(20,arena.Width-20)
		rover.Ypos = getRandomInt(20,arena.Height-20)
		rover.Angle_index = getRandomInt(0,8)
		//array or slice
		//rovers = append(rovers, rover)
		rovers[i] =  rover
	} //end of for loop on num_rovers
	//return rovers
} //end of make_rovers

func make_brain() Brain {
	var brain Brain
	for i:=0;i<NUM_NEURONS;i++ {
		brain.sign[i] = byte(getRandomInt(0,2))
	}
	var iconn [NUM_NEURONS]byte
	//fmt.Println("BEFORE ICONN NUM_NEURONS",NUM_NEURONS)
	//fully connected on inputs
	for i:=0;i<NUM_NEURONS;i++ {
		iconn[i] = byte(i)
	}
	brain.iconn = iconn
	var nconn  [NUM_NEURONS][NUM_NEURONS]byte
	//fmt.Println("doing nconn")
	//var junk []byte
	for ix := 0;ix<NUM_NEURONS;ix++ {
		for iy := 0;iy<NUM_NEURONS;iy++ {
			//junk = append(junk,byte(getRandomInt(0,2)))
			nconn[ix][iy] = byte(getRandomInt(0,2))
		}
		//nconn = append(nconn,junk)
	}
	//fmt.Println("past nconn")
	brain.nconn = nconn
	return brain
}
