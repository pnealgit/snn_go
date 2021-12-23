package main

import (
	"fmt"
)

func do_update() [NUM_ROVERS][2]int {

	//var err error
	var binary_sensor_data string
	//var max_index int
	var team_positions [NUM_ROVERS][2]int

	for ir := 0; ir < NUM_ROVERS; ir++ {
		if rovers[ir].Dead 		{
			continue
		}
		get_sensor_data(ir)
		binary_sensor_data = make_binary_sensor_data(ir)
		think(ir, binary_sensor_data)
		
		rovers[ir].Vel_x += rovers[ir].Accel_x
		rovers[ir].Vel_y += rovers[ir].Accel_y

		rovers[ir].Xpos += rovers[ir].Vel_x
		rovers[ir].Ypos += rovers[ir].Vel_y

		//for dumping back to javascript
		team_positions[ir][0] = rovers[ir].Xpos
		team_positions[ir][1] = rovers[ir].Ypos
	} //end of loop on rovers
	return team_positions
} //end of do_update

func change_direction(nai int) int  {
	//dx := ANGLES_DX[nai]
	//dy := ANGLES_DY[nai]

	
       if nai == 0 {
	       return 4
       }
       if nai == 1 {
	       return 3
       }
       if nai == 2 {
	       return 6
       }
       if nai == 3 {
	       return 5
       }
       if nai == 4 {
	       return 0
       }
       if nai == 5 {
	       return 7 
       }
       if nai == 6 {
	       return 2
       }
       if nai == 7 {
	       return 1
       }
       fmt.Println("DIRECTION FELL THROUGH")
       return 0
}


