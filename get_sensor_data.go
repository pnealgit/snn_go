package main

import (
	"fmt"
	"math"
)

func get_sensor_data(ir int) {
	//the return is a vector of 1s and 0s
//		var my_rover  Rover
//		rovers[ir] = rovers[ir]
	//gotta do it this way to avoid jumping over obstacles
		wall := 0
		var Xpos int
		var Ypos int
		var sensor_angle_index int
		var dist int
	for isensor := 0; isensor < NUM_SENSORS; isensor++ {
		sensor_angle_index = get_sensor_angle_index(isensor, rovers[ir].Angle_index)
		deltax := ANGLES_DX[sensor_angle_index]
		deltay := ANGLES_DY[sensor_angle_index]
		Xpos = rovers[ir].Xpos
		Ypos = rovers[ir].Ypos
		var zorro int
		zorro = check_food_position(Xpos, Ypos)
		if zorro == 7 {
			rovers[ir].Fitness += 1
		}
		zorro = check_wall_position(Xpos, Ypos)
		if zorro > 0 {
			rovers[ir].Dead = true
			fmt.Println("IN GET_SENSOR ROVER IR IS DEAD ",ir)

		}
		wall = 0
		for step := 0; step < SENSOR_LENGTH; step++ {
			Xpos += deltax
			Ypos += deltay
			fdx := float64(Xpos - rovers[ir].Xpos)
			fdy := float64(Ypos - rovers[ir].Ypos)
			dist = int(math.Hypot(fdx, fdy))
			if dist >= SENSOR_LENGTH {
				dist = SENSOR_LENGTH
				break
			}
			wall = check_wall_position(Xpos, Ypos)
			if wall > 0 {
				break
			}

			wall = check_food_position(Xpos, Ypos)
			if wall > 0 {
				//rovers[ir].Fitness +=2; //get some for food
				break
			}
		} //end of step loop
			rovers[ir].Sensor_data[isensor][0] = Xpos
			rovers[ir].Sensor_data[isensor][1] = Ypos
			rovers[ir].Sensor_data[isensor][2] = wall
			rovers[ir].Sensor_data[isensor][3] = dist
	} //end of isensor loop

	rovers[ir] = rovers[ir]
}

func make_binary_sensor_data(ir int) string {
	//fmt.Println("in make binary : ",rovers[ir].Sensor_data)
	
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
		otype := rovers[ir].Sensor_data[i][2]
		if otype == 0 {
			bsd = bsd + nothing
		}
		if otype > 0 && otype < 7 {
			bsd = bsd + no_go
		}
		if otype == 7 {
			bsd = bsd + eats
		}

		dist := rovers[ir].Sensor_data[i][3]
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
