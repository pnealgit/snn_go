package main

import (
	//"fmt"
)

func think(ir int, sensor_data_string string) int {
//	fmt.Println("THINK SDATA STR: ",sensor_data_string)
	var brain Brain
	brain = rovers[ir].brain
	var sensor_data []byte
	var sig byte
	//could have called Atoi on this, but meh
	for i := 0; i < len(sensor_data_string); i++ {
		if sensor_data_string[i] == '0' {
			sig = 0
		} else {
			sig = 1
		}
		sensor_data = append(sensor_data, sig)
	}

	//throw in some noise if the sensors sense nothing...Just vast
	//empty space....
	knt := 0
	for ik := 0; ik < len(sensor_data); ik++ {
		knt = knt + int(sensor_data[ik])
	}

	if knt <= 0 {
		sensor_data[getRandomInt(0, len(sensor_data))] = 1
	}

	var temp_outps [NUM_NEURONS]byte
	var memb int //because memb can go negative
	var outps [NUM_NEURONS]byte
	var fire_knt [NUM_NEURONS]int
	var inps []byte

	for k := 0; k < NUM_NEURONS; k++ {
		temp_outps[k] = 0
		memb = 0
		outps[k] = 0
		fire_knt[k] = 0
	}

	// do I need to do this ?
	// it is array to array copy
	inps = sensor_data
	sign := brain.sign
	nconn := brain.nconn
	iconn := brain.iconn

	for epoch := 0; epoch < SETTLING_TIME; epoch++ {
		for i := 0; i < len(temp_outps); i++ {
			outps[i] = temp_outps[i]
			//save what was done from last epoch
			temp_outps[i] = 0
		}
		for nindex := 0; nindex < len(sensor_data); nindex++ {
			memb = 0
			if outps[nindex] == 0 {
				//not in refactory state
				memb += int(inps[nindex] * iconn[nindex])
				//count from other neurons with positive or negative
				for n := 0; n < len(sensor_data); n++ {
					if sign[n] == 1 {
						memb += int(outps[n] * nconn[nindex][n])
					} else {
						memb -= int(outps[n] * nconn[nindex][n])
					}
				}
				if memb < 0 {
					memb = 0
				}
			} //end of not refactory
			//fire or not !
			//leakage
			if memb >= LEAKING_CONSTANT {
				memb -= LEAKING_CONSTANT
			}
			r := getRandomInt(-2, 3)
			if memb >= (THRES + r) {
				//	fmt.Println("\nFIRED NEURON : ",nindex)
				//	fmt.Println("MEMB ",memb)
				//	fmt.Println("THRESHOLD WAS ",THRES+r)

				temp_outps[nindex] = 1
			} else {
				temp_outps[nindex] = 0
			}
		}
		//end of pass through all neurons

		//fire_knt is used to choose what sensor to go with
		for k := 0; k < NUM_NEURONS; k++ {
			fire_knt[k] += int(temp_outps[k])
		}
		//#save outps for refactory
		for k := 0; k < NUM_NEURONS; k++ {
			outps[k] = temp_outps[k]
			temp_outps[k] = 0
		}
		//fmt.Println("\nAT END")
		//fmt.Println("OUTPS     : ",outps)
		//fmt.Println("TEMP_OUTPS: ",temp_outps)

	} //end of settling_time loop (epochs)

	var max_index int
	max_index = 1 //go straight if nothing happens;
	var max_value int
	max_value = 0
	sum := 0
	//choose a direction based on fireknt
	for i := 0; i < 8; i++ {
		sum += fire_knt[i]
	}
	if sum > max_value {
		max_value = sum
		max_index = 0
	}
	sum = 0
	for i := 8; i < 16; i++ {
		sum += fire_knt[i]
	}
	if sum > max_value {
		max_value = sum
		max_index = 1
	}
	sum = 0
	for i := 16; i < 24; i++ {
		sum += fire_knt[i]
	}
	if sum > max_value {
		max_value = sum
		max_index = 2
	}

	for i := 24; i < 32; i++ {
		sum += fire_knt[i]
	}
	if sum > max_value {
		max_value = sum
		max_index = 3
	}
	
	for i := 32; i < 40; i++ {
		sum += fire_knt[i]
	}
	if sum > max_value {
		max_value = sum
		max_index = 4
	}
	
	for i := 40; i < 48; i++ {
		sum += fire_knt[i]
	}
	if sum > max_value {
		max_value = sum
		max_index = 5
	}
	
	for i := 48; i < 56; i++ {
		sum += fire_knt[i]
	}
	if sum > max_value {
		max_value = sum
		max_index = 6
	}
	
	for i := 56; i < 64; i++ {
		sum += fire_knt[i]
	}
	if sum > max_value {
		max_value = sum
		max_index = 7
	}



	return max_index

} //end of think
