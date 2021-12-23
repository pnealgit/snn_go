package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func select_brains() {
	sum := 0
	fmt.Println("\nFITNESS\n")

	for ir := 0; ir < NUM_ROVERS; ir++ {
		sum += rovers[ir].Fitness
	}
	fmt.Println("team sum score ", sum)

	//gotta sort an array with a slice sort
	sort.Slice(rovers[:], func(i, j int) bool {
		return rovers[i].Fitness > rovers[j].Fitness
	})
	fmt.Println("\n after sort")
	for ir := 0; ir < NUM_ROVERS; ir++ {
		fmt.Println(ir, rovers[ir].Fitness)
	}

	fmt.Println("\nBEST SCORE ", rovers[0].Fitness)
	fmt.Println("WRST SCORE ", rovers[NUM_ROVERS-1].Fitness)
	//zero out all the scores -- starting another epoch
	for ir := 0; ir < NUM_ROVERS; ir++ {
		rovers[ir].Fitness = 0
		rovers[ir].Dead = false
		//rovers[ir].Xpos = getRandomInt(20, arena.Width-20)
		//rovers[ir].Ypos = getRandomInt(20, arena.Height-20)
		rovers[ir].Xpos = arena.Width/2
		rovers[ir].Ypos = arena.Height/2
		rovers[ir].Accel_x = getRandomInt(-1,2)
		rovers[ir].Accel_y = getRandomInt(-1,2)
		rovers[ir].Vel_x = getRandomInt(-1,2)
		rovers[ir].Vel_y = getRandomInt(-1,2)

	}

	//HOPE ! HAHAHaaa
	//I hope this works. If brain is an array, than no problem
	//If brain is a slice, got problems. Because the brain copy
	//is the same as the brain it was copied from (*pointer stuff)
	//go doesn't have a deep copy function

	//var test_brain Brain
	//test_brain = rovers[NUM_ROVERS-1].brain
	elite_cut := int(float64(NUM_ROVERS) * .2)
	//elite_cut = 1
	for ib := elite_cut; ib < NUM_ROVERS; ib++ {
		//old_idx := getRandomInt(0,elite_cut)
		//rovers[ib].brain = rovers[old_idx].brain
		bam := getRandomInt(0,elite_cut)
		c := [NUM_NEURONS][NUM_NEURONS]byte{}
		c = rovers[bam].brain.nconn
		rovers[ib].brain.nconn = c

		sig := [NUM_NEURONS]byte{}
		sig = rovers[bam].brain.sign
		rovers[ib].brain.sign = sig
	}

	mutate_brains(elite_cut)
	//fmt.Println("TEST BRAIN: ",test_brain.nconn[1])
	//fmt.Println("NEW  BRAIN: ",rovers[NUM_ROVERS-1].brain.nconn[1])
	//fmt.Println("BEST BRAIN: ",rovers[0].brain.nconn[1])

} //end of select

func getRandomFloat64(min float64, max float64) float64 {
	return 0.0 + (rand.Float64() * (max - min)) + min
}

func getRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func mutate_brains(elite_cut int) {
	//I am not mutating sign here. Too drastic
	//fmt.Println("IN MUTATE BRAINS")
	var num_mutations int
	var nn float64
	nn = float64(NUM_NEURONS)
	num_mutations = int(nn * nn / 5.0)
	//fmt.Println("NUM MUTATIONS: ",num_mutations)
	for im := elite_cut; im < NUM_ROVERS; im++ {
		for k := 0; k < num_mutations; k++ {
			ix := getRandomInt(0, NUM_NEURONS)
			iy := getRandomInt(0, NUM_NEURONS)
			if rovers[im].brain.nconn[ix][iy] == 1 {
				rovers[im].brain.nconn[ix][iy] = 0
			} else {
				rovers[im].brain.nconn[ix][iy] = 1
			}
		}
	} //end of loop on num_rovers
} //end of mutate func
