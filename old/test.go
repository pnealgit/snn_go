package main
import (
	"fmt"
	"sort"
)
func main() {

	type S1 struct {
		id	int
		First	int
		Stuff	[4]int
	}


	var s1 S1
	for i:=0;i<4;i++ {
		s1.Stuff[i] = i
	}
	s1.First = 11
	s1.id = 1

	var s2 S1

	for i:=0;i<4;i++ {
		s2.Stuff[i] = i*2
	}
	s2.First = 222
	s2.id = 2

	var s3 S1
	for i:=0;i<4;i++ {
		s3.Stuff[i] = i*3
	}
	s3.First = 333 
	s3.id = 3

	var bigs []S1

	bigs = append(bigs,s1)
	bigs = append(bigs,s2)
	bigs = append(bigs,s3)

	for i:=0;i<len(bigs);i++ {
		fmt.Println(bigs[i])
	}


	fmt.Println("S2: alone")
	s2.Stuff[3] = 9
	for i:=0;i<len(bigs);i++ {
		fmt.Println(bigs[i])
	}

	sort.Slice(bigs,func(i,j int) bool {
		return bigs[i].First < bigs[j].First
	})
	fmt.Println("\nbigs as slice sorted")
	for i:=0;i<len(bigs);i++ {
		fmt.Println(bigs[i])
	}

	sort.Slice(bigs,func(i,j int) bool {
		return bigs[i].First > bigs[j].First
	})
	
	fmt.Println("\nbigs as slice reverse sorted")
	for i:=0;i<len(bigs);i++ {
		fmt.Println(bigs[i])
	}

	bigs[2] = bigs[0]
	bigs[2].First = 41
	bigs[2].Stuff[3] = 99
	bigs[0].Stuff[1] = 12434

	sort.Slice(bigs,func(i,j int) bool {
		return bigs[i].First > bigs[j].First
	})

	fmt.Println("\nbigs as slice reverse bigs2 = bigs2 bs2.f = 41 bs2 sff3=99 sorted")
	for i:=0;i<len(bigs);i++ {
		fmt.Println(bigs[i])
	}

}



