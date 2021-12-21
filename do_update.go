package main

import (
	"fmt"
)

func do_update() [NUM_ROVERS][2]int {

	//var err error
	var binary_sensor_data string
	var max_index int
	var new_angle_index int
	var team_positions [NUM_ROVERS][2]int

	for ir := 0; ir < NUM_ROVERS; ir++ {
		get_sensor_data(ir)
		binary_sensor_data = make_binary_sensor_data(ir)
		max_index = think(ir, binary_sensor_data)

		//start east go counterclockwise
		//sensors go from left to right 0-2
		new_angle_index = rovers[ir].Angle_index
		if max_index == 0 {
			new_angle_index = new_angle_index + 1
		}
		if new_angle_index > NUM_ANGLES-1 {
			new_angle_index = 0
		}
		if max_index == 1 {
			//do nothing. Just a straight ahead
		}

		if max_index == 2 {
			if new_angle_index > 0 {
				new_angle_index = new_angle_index - 1
			} else {
				new_angle_index = NUM_ANGLES - 1
			}
		}

		//why is this done.. reward for going straight
		if new_angle_index == rovers[ir].Angle_index && !rovers[ir].Dead 		{
			rovers[ir].Fitness += 1 //go straight
		}
		if !rovers[ir].Dead 		{
			rovers[ir].Fitness += 1 //for staying alive
		}
		deltax := ANGLES_DX[new_angle_index]
		deltay := ANGLES_DY[new_angle_index]

		x := rovers[ir].Xpos + deltax
		y := rovers[ir].Ypos + deltay

		if x<= 1||y<= 1|| x >= arena.Width-1 || y >= arena.Height-1 {
			//new_angle_index = change_direction(new_angle_index)
			//do nothing
		} else {

			rovers[ir].Angle_index = new_angle_index
			rovers[ir].Xpos += ANGLES_DX[new_angle_index]
			rovers[ir].Ypos += ANGLES_DY[new_angle_index]
		}
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


