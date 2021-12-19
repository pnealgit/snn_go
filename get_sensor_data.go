package main

import (
	"fmt"
	"math"
)

func get_sensor_data(my_rover Rover) Rover {
	fmt.Println("SENSOR LENGTH: ",SENSOR_LENGTH);
	//the return is a vector of 1s and 0s

	//gotta do it this way to avoid jumping over obstacles
		wall := 0
		var Xpos int
		var Ypos int
		var sensor_angle_index int
		var dist int
	for isensor := 0; isensor < NUM_SENSORS; isensor++ {
		sensor_angle_index = get_sensor_angle_index(isensor, my_rover.Angle_index)
		fmt.Println("SENSOR ROVER ANGLE INDEX: ",sensor_angle_index,my_rover.Angle_index)
		deltax := ANGLES_DX[sensor_angle_index]
		deltay := ANGLES_DY[sensor_angle_index]
		fmt.Println("DELTAX,Y: ",deltax,deltay)
		Xpos = my_rover.Xpos
		Ypos = my_rover.Ypos
		wall = 0
		for step := 0; step < SENSOR_LENGTH; step++ {
			Xpos += deltax
			Ypos += deltay
			fdx := float64(Xpos - my_rover.Xpos)
			fdy := float64(Ypos - my_rover.Ypos)
			dist = int(math.Hypot(fdx, fdy))
			if dist >= SENSOR_LENGTH {
				fmt.Println("DIST: ",dist)
				dist = SENSOR_LENGTH
				break
			}
			wall = check_wall_position(Xpos, Ypos)
			if wall > 0 {
				break
			}

			wall = check_food_position(Xpos, Ypos)
			if wall > 0 {
				break
			}
		} //end of step loop
			my_rover.Sensor_data[isensor][0] = Xpos
			my_rover.Sensor_data[isensor][1] = Ypos
			my_rover.Sensor_data[isensor][2] = wall
			my_rover.Sensor_data[isensor][3] = dist
	} //end of isensor loop
	fmt.Println("return get_sensor_data rover SENSOR_DATA: ",my_rover.Sensor_data)

	return my_rover
}

func make_binary_sensor_data(my_rover Rover) string {

	fmt.Println("in make binary : ",my_rover.Sensor_data)

	//knt := 0
	//obstacles
	nothing := "0000"
	no_go := "0001"
	eats := "0011"

	//distances
	far := "1000"
	soso := "1001"
	clos := "1011" //close is a reserved word
	alert := "1011"

	var bsd string
	bsd = ""

	//string contains 3 groups containing a type code and a distance code
	for i := 0; i < NUM_SENSORS; i++ {
		otype := my_rover.Sensor_data[i][2]
		if otype == 0 {
			bsd = bsd + nothing
		}
		if otype > 0 && otype < 7 {
			bsd = bsd + no_go
		}
		if otype == 7 {
			bsd = bsd + eats
		}

		dist := my_rover.Sensor_data[i][3]
		junkf := float64(dist)
		slf := float64(SENSOR_LENGTH)
		if junkf > .8*slf {
			if otype > 0{
				bsd = bsd + far
			} else {
				bsd = bsd + nothing
			}
			break
		}
		if junkf > .5*slf {
			bsd = bsd + soso
			break
		}
		if junkf > .15*slf {
			bsd = bsd + clos
			break
		}
		bsd = bsd + alert
	}

	return bsd
}

func check_wall_position(xp int, yp int) int {

	if yp < 0 {
		return 1
	}

	if yp > arena.Height {
		return 2
	}

	if xp < 0 {
		return 3
	}
	if xp > arena.Width {
		return 4
	}
	return 0

}

func check_food_position(xp int, yp int) int {
	status := 0
	flen := len(arena.Food)

	for i := 0; i < flen; i++ {
		f := arena.Food[i]
		dx := float64(f[0] - xp)
		dy := float64(f[1] - yp)
		dist := int(math.Hypot((dx), (dy)))
		if dist <= FOOD_RADIUS {
			status = 7
			break
		}
	}
	return status
}

func get_sensor_angle_index(isensor int,rover_angle_index int) int {
	ai := 99
	//first sensor
	if isensor == 0 {
		ai = rover_angle_index + 1
	}

	//middle sensor points in direction of movement
	if isensor == 1 {
		ai = rover_angle_index
	}

	if isensor == 2 {
		ai = rover_angle_index - 1
	}

	if ai > NUM_ANGLES-1 {
		ai = ai % NUM_ANGLES
	}

	if ai < 0 {
		ai = NUM_ANGLES - 1
	}

	return ai
}
