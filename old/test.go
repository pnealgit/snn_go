package main
import (
	"fmt"
)
func main() {

	type S1 struct {
		First	int
		Stuff	[]int
	}


	var s1 S1
	for i:=0;i<4;i++ {
		s1.Stuff = append(s1.Stuff,i)
	}
	var s2 S1

	s2 = s1
	s2.Stuff[3] = 9
	fmt.Println("S1: ",s1)
	fmt.Println("S2: ",s2)

	s1.Stuff[2] = 8
	fmt.Println("S1: ",s1)
	fmt.Println("S2: ",s2)

	
	var b1  []int
	var b2  []int

	for i:=0;i<5;i++ {
		b1 = append(b1,i)
	}

	b2 = b1
	fmt.Println(b1)
	fmt.Println(b2)

	b2[3] = 999
	fmt.Println(b1)
	fmt.Println(b2)
	b1[2] = 222
	fmt.Println(b1)
	fmt.Println(b2)
	
}



