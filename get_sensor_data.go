package main

import (
//	"fmt"
)
func get_sensor_data(my_rover Rover) [NUM_NEURONS]byte {
//	fmt.Println("IN GET SENSOR DATA")
	var sensor_data [NUM_NEURONS]byte
	//make this real later
	for i:=0;i<NUM_NEURONS;i++ {
	    //sensor_data[i] = append(sensor_data,byte(getRandomInt(0,2)))
	    sensor_data[i] = byte(getRandomInt(0,2))
    	}
	return sensor_data
}

