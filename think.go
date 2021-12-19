package main

import (
	"fmt"
)

func think(brain Brain, sensor_data_string string) int {
	//fmt.Println("IN THINK sensor_data: ", sensor_data_string)
	//var new_angle_index byte
	//new_angle_index = 0

	//throw in some noise if the sensors sense nothing...Just vast
	//empty space....
	var sensor_data []byte
	var sig byte
	for i := 0; i < len(sensor_data_string); i++ {
		if sensor_data_string[i] == '0' {
			sig = 0
		} else {
			sig = 1
		}
		sensor_data = append(sensor_data, sig)
	}
	fmt.Println("SENSOR DATA BYTES: ", sensor_data)
	knt := 0
	for ik := 0; ik < len(sensor_data); ik++ {
		knt = knt + int(sensor_data[ik])
	}

	if knt <= 0 {
		sensor_data[getRandomInt(0, len(sensor_data))] = 1
	}

	var temp_outps [NUM_NEURONS]byte
	var memb [NUM_NEURONS]int //because memb can go negative
	var outps [NUM_NEURONS]byte
	var fire_knt [NUM_NEURONS]byte
	var inps []byte

	for k := 0; k < NUM_NEURONS; k++ {
		temp_outps[k] = 0
		memb[k] = 0
		outps[k] = 0
		fire_knt[k] = 0
	}

	inps = sensor_data // need deep copy here ?
	sign := brain.sign
	nconn := brain.nconn
	iconn := brain.iconn

	for epoch := 0; epoch < SETTLING_TIME; epoch++ {

		for nindex := 0; nindex < len(sensor_data); nindex++ {
			memb[nindex] = 0
			if outps[nindex] == 0 {
				//not in refactory state
				memb[nindex] += int(inps[nindex] * iconn[nindex])
				//count from other neurons with positive or negative
				var stuff int
				stuff = 0
				for n := 0; n < len(sensor_data); n++ {
					if sign == 1 {
						stuff = int(outps[n] * nconn[nindex][n])
					} else {
						stuff -= int(outps[n] * nconn[nindex][n])
					}
				}
				memb[nindex] += stuff
				if memb[nindex] < 0 {
					memb[nindex] = 0
				}
			} //end of not refactory
			//fire or not !
			r := getRandomInt(-2, 3)
			if memb[nindex] >= (THRES + r) {
				temp_outps[nindex] = 1
				memb[nindex] = 0
			} else {
				temp_outps[nindex] = 0
			}

			//leakage
			if memb[nindex] >= LEAKING_CONSTANT {
				memb[nindex] -= LEAKING_CONSTANT
			}
		}
		//end of pass through all neurons

		//fire_knt is used to choose what sensor to go with
		for k := 0; k < 4; k++ {
			fire_knt[0] += temp_outps[k]
		}
		for k := 4; k < 8; k++ {
			fire_knt[1] += temp_outps[k]
		}
		for k := 8; k < NUM_NEURONS; k++ {
			fire_knt[2] += temp_outps[k]
		}

		//#save outps for refactory
		for k := 0; k < NUM_NEURONS; k++ {
			outps[k] = temp_outps[k]
			temp_outps[k] = 0
		}
	} //end of settling_time loop (epochs)

	var min_index int
	min_index = 1 //go straight if nothing happens;
	var min_value int
	min_value = 99

	//choose a direction based on sensor.
	for i := 0; i < NUM_SENSORS; i++ {
		if fire_knt[i] <= byte(min_value) {
			min_value = int(fire_knt[i])
			min_index = i
		}
		fire_knt[i] = 0
	}

	return min_index

} //end of think
