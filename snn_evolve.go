package main

import (
	//"fmt"
	"math/rand"
)

func getRandomFloat64(min float64, max float64) float64 {
	return 0.0 + (rand.Float64() * (max - min)) + min
}

func getRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}


func mutate_brains(rovers []Rover) {
	//I am not mutating sign here. Too drastic
	//fmt.Println("IN MUTATE BRAINS")
	var num_mutations int
	var nn  float64
	nn =float64(NUM_NEURONS)
	num_mutations = int(nn * nn/5.0)
	//fmt.Println("NUM MUTATIONS: ",num_mutations)
	for im := 4; im < NUM_ROVERS; im++ {
		for k:=0;k<num_mutations;k++ {
			ix := getRandomInt(0,NUM_NEURONS)
			iy := getRandomInt(0,NUM_NEURONS)
			rovers[im].brain.nconn[ix][iy] = byte(getRandomInt(0,2))
		}
	} //end of loop on num_rovers
} //end of mutate func

// FitnessSorter sorts rovers by score
type FitnessSorter []Rover

func (a FitnessSorter) Len() int           { return len(a) }
func (a FitnessSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FitnessSorter) Less(i, j int) bool { return a[i].Fitness > a[j].Fitness }


