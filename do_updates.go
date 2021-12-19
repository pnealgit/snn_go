package main

import (
	"fmt"
)

func do_updates(rovers []Rover) [NUM_ROVERS][2]int {

	//var err error
	fmt.Println("IN DO UPDATES")
	var sensor_data string
	var min_index int
	var new_angle_index int
	var team_positions [NUM_ROVERS][2]int

	//var sensor_data []bytes
	for ir := 0; ir < NUM_ROVERS; ir++ {
		sensor_data = get_sensor_data(rovers[ir])
		min_index = think(rovers[ir].brain, sensor_data)

		new_angle_index = rovers[ir].Angle_index
		if min_index == 0 {
			new_angle_index = new_angle_index + 1
		}
		if new_angle_index > NUM_ANGLES-1 {
			new_angle_index = 0
		}
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
		//position := strconv.Itoa(rovers[ir].Xpos) + "," + strconv.Itoa(rovers[ir].Ypos)
		//fmt.Println(ir,position)
		team_positions[ir][0] = rovers[ir].Xpos
		team_positions[ir][1] = rovers[ir].Ypos

	} //end of loop on rovers
	return team_positions
} //end of do_update
