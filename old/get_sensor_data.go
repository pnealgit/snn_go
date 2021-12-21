package main

import (
	"fmt"
	"math"
)


func get_sensor_data(my_rover Rover) []byte {
	//the return is a vector of 1s and 0s 
	
	fmt.Println("IN GET SENSOR DATA")
	fmt.Println("NUM_SENSORS: ",NUM_SENSORS)

	//gotta do it this way to avoid jumping over obstacles
	for isensor:=0;isensor<NUM_SENSORS;isensor++ {
		wall := 0
		Xpos := my_rover.Xpos
		Ypos := my_rover.Ypos
		var sensor_angle_index int
		sensor_angle_index = get_sensor_angle_index(isensor,my_rover.Angle_index)
		deltax := ANGLES_DX[sensor_angle_index]
		deltay := ANGLES_DY[sensor_angle_index]
		for step:=0;step<SENSOR_LENGTH;step++{
			Xpos += deltax
			Ypos += deltay
			fdx := float64(Xpos - my_rover.Xpos)
			fdy := float64(Ypos - my_rover.Ypos)
			dist := int(math.Hypot(fdx,fdy))
			if dist >= SENSOR_LENGTH {
				dist = SENSOR_LENGTH
				break
			}
			wall  = check_wall_position(Xpos,Ypos)
			if wall > 0 {
				my_rover.Sensor_data[isensor][0] = Xpos
				my_rover.Sensor_data[isensor][1] = Ypos
				my_rover.Sensor_data[isensor][2] = wall
				my_rover.Sensor_data[isensor][3] = dist
				break
			}

			wall  = check_food_position(Xpos,Ypos)
			if wall > 0 {
				my_rover.Sensor_data[isensor][0] = Xpos
				my_rover.Sensor_data[isensor][1] = Ypos
				my_rover.Sensor_data[isensor][2] = wall
				my_rover.Sensor_data[isensor][3] = dist
				break
			}

		} //end of step loop
	} //end of isensor loop

	//ok. Got data convert it to a binary string
	var sensor_data []byte
	sensor_data = make_binary_sensor_data(my_rover.Sensor_data)
	return sensor_data
}

func make_binary_sensor_data(sensor_data [NUM_SENSORS][4]int) []byte {
	
	fmt.Println("MAKE BINARY SENSOR DATA")
	knt := 0
	//obstacles
	nothing:= [4]byte{0,0,0,0}
	no_go  := [4]byte{0,0,0,1}
	eats   := [4]byte{0,0,1,1}

	//distances
	far  := [4]byte{1,0,0,0}
	soso := [4]byte{1,0,0,1}
	clos := [4]byte{1,0,1,1} //close is a reserved word
	alert:= [4]byte{1,0,1,1}

	var bsd []byte
	//string contains 3 groups containing a type code and a distance code
	for i:=0;i<NUM_SENSORS;i++ {
		otype := sensor_data[i][2]
		if otype == 0 {
			bsd = append(bsd,no_go)
		}
		if otype > 0 && otype < 7 {
			bsd = append(bsd,no_go)
		}
		if otype == 7 {
			bsd = append(bsd,eats)
		}

		dist := sensor_data[i][3]
		junkf := float64(dist)
		slf := float64(SENSOR_LENGTH)
		if junkf > .8 * slf {
			bsd = append(bsd,far)
			break
		}
		if junkf > .5 * slf {
			bsd = append(bsd,soso)
			break
		}
		if junkf > .15 * slf {
			bsd = append(bsd,clos)
			break
		}
		bsd = append(bsd,alert)

	}
	
	return bsd
}

func check_wall_position(xp int,yp int) int {
	/*
	if yp < 0 {
		return 1
	}

	if yp > HEIGHT  {
		return 2
	}

	if xp < 0 {
		return 3

	if xp > WIDTH  {
		return 4
	}
	*/
	return 0

}

func check_food_position(xp int,yp int) int {
	status := 0
	flen := len(arena.Food)
	fmt.Println("FOOD LENGTH: ",flen)

	for i:=0;i<flen;i++{
		f := arena.food[i]
		dx := float64(f[0]-xp)
		dy := float64(f[1]-yp)
		dist = int(math.Hypot((dx),(dy)))
		if dist <= FOOD_RADIUS {
			status = 7
			break
		}
	}
	return status
}

func get_sensor_angle_index(rover_angle_index int,isensor int) int {
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

	if ai > NUM_ANGLES - 1 {
		ai = ai % NUM_ANGLES
	}

	if ai < 0 {
		ai = NUM_ANGLES -1
	}

	return ai
}
