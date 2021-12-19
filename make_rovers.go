package main
import (
	"time"
	"math/rand"
//	"fmt"
)

func make_rovers()  {
	rand.Seed(time.Now().UTC().UnixNano())
	var rover Rover	
	for i := 0; i < NUM_ROVERS; i++ {
		brain_seed  := (time.Now().UnixNano())
		rover.brain = make_brain(brain_seed)
		rover.Fitness = 0
		rover.Xpos = getRandomInt(0,arena.Width)
		rover.Ypos = getRandomInt(0,arena.Height)
		rover.Angle_index = getRandomInt(0,9)

		rovers = append(rovers, rover)
	} //end of for loop on num_rovers
	//return rovers
} //end of make_rovers

func make_brain(b_seed int64) Brain {
	//fmt.Println("in make_brain")
	rand.Seed(b_seed)
	var brain Brain
	brain.sign = byte(getRandomInt(0,2))
	var iconn []byte
	//fmt.Println("BEFORE ICONN NUM_NEURONS",NUM_NEURONS)
	for i:=0;i<NUM_NEURONS;i++ {
		iconn = append(iconn,1)   //fully connected on input
	}
	brain.iconn = iconn
	var nconn  [][]byte
	//fmt.Println("doing nconn")
	var junk []byte
	for ix := 0;ix<NUM_NEURONS;ix++ {
		for iy := 0;iy<NUM_NEURONS;iy++ {
			junk = append(junk,byte(getRandomInt(0,2)))
		}
		nconn = append(nconn,junk)
	}
	//fmt.Println("past nconn")
	brain.nconn = nconn
	return brain
}
