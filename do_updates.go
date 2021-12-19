package main

import (
	"fmt"
)

func do_updates(rovers []Rover) [NUM_ROVERS][8]int {

	//var err error
	fmt.Println("IN DO UPDATES")
	var binary_sensor_data string
	var min_index int
	var new_angle_index int
	var team_positions [NUM_ROVERS][8]int

	//var sensor_data []bytes
	for ir := 0; ir < NUM_ROVERS; ir++ {
		rovers[ir]  = get_sensor_data(rovers[ir])
		fmt.Println("SENSOR DATA: ",rovers[ir].Sensor_data)

		binary_sensor_data = make_binary_sensor_data(rovers[ir])

		min_index = think(rovers[ir].brain, binary_sensor_data)
		//start east go counterclockwise
		//sensors go from left to right 0-2
		new_angle_index = rovers[ir].Angle_index
		if min_index == 0 {
			new_angle_index = new_angle_index + 1
		}
		if new_angle_index > NUM_ANGLES-1 {
			new_angle_index = 0
		}
		//if mindex = 1 just go in direction already heading

		if min_index == 2 {
			if new_angle_index > 0 {
				new_angle_index = new_angle_index - 1
			} else {
				new_angle_index = NUM_ANGLES - 1
			}
		}

		fmt.Println(" angle index ", new_angle_index)
		//why is this done.. reward for going straight
		if new_angle_index == rovers[ir].Angle_index {
			rovers[ir].Fitness += 1.0
		}

		rovers[ir].Angle_index = new_angle_index

		rovers[ir].Xpos += ANGLES_DX[new_angle_index]
		rovers[ir].Ypos += ANGLES_DY[new_angle_index]
		//for dumping back to javascript
		team_positions[ir][0] = rovers[ir].Xpos
		team_positions[ir][1] = rovers[ir].Ypos
		//fmt.Println("rover: ",rovers[ir]);
		knt := 2
		fmt.Println("TEAM POSITION ROVER SENSOR_DATA:",rovers[ir].Sensor_data)

		for isensor:=0;isensor < NUM_SENSORS;isensor++ {
			for j :=0;j<2;j++ {
				
			    team_positions[ir][knt] = rovers[ir].Sensor_data[isensor][j]
			    knt++
		    }
	    }
	} //end of loop on rovers
	fmt.Println("TEAM POS: ",team_positions)
	return team_positions
} //end of do_update
