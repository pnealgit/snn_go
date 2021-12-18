package main
import (
	"fmt"
	"sort"
)

func select_brains(rovers []Rover) {
	sum := 0
	fmt.Println("sum ", sum)
	for ir := 0; ir < NUM_ROVERS; ir++ {
		sum += rovers[ir].Fitness
	}
	fmt.Println("team sum score ", sum)

	sort.Sort(FitnessSorter(rovers))
	fmt.Println("BEST SCORE ", rovers[0].Fitness)
	fmt.Println("WRST SCORE ", rovers[NUM_ROVERS-1].Fitness)
	//zero out all the scores -- starting another epoch
	for ir := 0; ir < NUM_ROVERS; ir++ {
		rovers[ir].Fitness = 0
	}

	elite_cut := int(float64(NUM_ROVERS) * .3)
	//fmt.Println("elite cut is : ", elite_cut)

	for ib := 0;ib<NUM_ROVERS;ib++{
		old_idx := getRandomInt(0,elite_cut)
		rovers[ib].brain = make_brain(rovers[old_idx].brain.seed)
	}

} //end of select
