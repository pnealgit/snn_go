package main

import (
		"fmt"
)

func think(ir int, sensor_data_string string) {
	fmt.Println("\nIR,DATA IN: ",ir,sensor_data_string)

	var brain Brain
	brain = rovers[ir].brain
	var sensor_data [NUM_NEURONS]byte

	sensor_data = convert_sensor_data(sensor_data_string)  

	var temp_outps [NUM_NEURONS]byte
	var memb [NUM_NEURONS]int //because memb can go negative
	var outps [NUM_NEURONS]byte
	var fire_knt [NUM_NEURONS]int
	var inps [NUM_NEURONS]byte

	/*for k := 0; k < NUM_NEURONS; k++ {
		temp_outps[k] = 0
		outps[k] = 0
		fire_knt[k] = 0
		memb[k] = 0
	}
	*/

	// do I need to do this ?
	// it is array to array copy
	inps = sensor_data
	fmt.Println("INPS: ",inps)

	sign := brain.sign
	iconn := brain.iconn
	outps = inps //just to start things off

	for epoch := 0; epoch < SETTLING_TIME; epoch++ {
		for k:=0;k<NUM_NEURONS;k++ {
			outps[k] = temp_outps[k]
			temp_outps[k] = 0
		}
		fmt.Println("OUTPS: ",outps)

		for nindex := 0; nindex < NUM_NEURONS; nindex++ {
			memb[nindex] = 0
			if outps[nindex] == 0 {
				//not in refactory state
				//do input to membrane
				memb[nindex] = int(input_membrane(inps,iconn))
				//fmt.Println("AFTER INPUT MEMBRANE MEMB: ",memb)

				nconn := brain.nconn[nindex]
				memb[nindex] += int(n_n_membrane(nconn,sign,outps ))

			} //end of not refactory
		} //just save membrane and compute firings after

	fmt.Println("MEMB: ",memb)
	temp_outps = get_output_state(memb)
	//fmt.Println("AFtEr GET TEMP_OUTPS: ",temp_outps)

		//fire_knt is used to choose what sensor to go with
		for k := 0; k < NUM_NEURONS; k++ {
			fire_knt[k] += int(temp_outps[k])
		}
	} //end of settling_time loop (epochs)


//		fmt.Println("\nIR,OUTPS: ",ir,temp_outps)
//		fmt.Println("IR,INPS : ",ir,inps)

		dx:= temp_outps[0] + temp_outps[NUM_NEURONS-1]
		dy:= temp_outps[1] + temp_outps[NUM_NEURONS-2]
		if dx == 0 {
			rovers[ir].Accel_x = 0
		}
		if dx == 1 {
			rovers[ir].Accel_x = 1
		}
		if dx == 2 {
			rovers[ir].Accel_x = -1
		}
		if dy == 0 {
			rovers[ir].Accel_y = 0
		}
		if dy == 1 {
			rovers[ir].Accel_y = 1
		}
		if dy == 2 {
			rovers[ir].Accel_y = -1
		}


} //end of think
func convert_sensor_data(sensor_data_string string ) [NUM_NEURONS]byte {
	        //could have called Atoi on this, but meh
	var sig byte
	sig = 0
	var sensor_data [NUM_NEURONS]byte

        for i := 0; i < len(sensor_data_string); i++ {
                if sensor_data_string[i] == '0' {
                        sig = 0
                } else {
                        sig = 1
                }
                sensor_data[i] = sig
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

	return sensor_data
}
func input_membrane(inps [NUM_NEURONS]byte,iconn [NUM_NEURONS]byte) int {

    junk := 0
    
    for il := 0; il < NUM_NEURONS; il++ {
        junk+= int(inps[il] * iconn[il])
    }
    return junk
}

func n_n_membrane(nconn [NUM_NEURONS]byte,sign [NUM_NEURONS]byte,outps [NUM_NEURONS]byte) int {

junk := 0
//fmt.Println("IN N_N - nconn: ",nconn)
//fmt.Println("IN N_N - outps: ",outps)

for il := 0; il < NUM_NEURONS; il++ {
    if sign[il] == 1 {
             junk += int(outps[il] * nconn[il])
    } else {
             junk -= int(outps[il] * nconn[il])
    }
                 
}
return junk
}

func get_output_state(membrane [NUM_NEURONS]int) [NUM_NEURONS]byte {
	var junk  [NUM_NEURONS]byte

	//fmt.Println("OUTPUT STATE - MEMBRANE: ",membrane)
       for nindex:=0;nindex<NUM_NEURONS;nindex++ {
           junk[nindex] = 0
           if membrane[nindex] < 0 {
                 membrane[nindex] = 0
           }
           if membrane[nindex] >= LEAKING_CONSTANT {
                 membrane[nindex] -= LEAKING_CONSTANT
           }

           r := getRandomInt(-2, 3)
           if membrane[nindex] >= (THRES + r) {
                      //fmt.Println("\nFIRED NEURON : ",nindex)
                       //fmt.Println("MEMB ",membrane)
                       //fmt.Println("THRESHOLD WAS ",THRES+r)
                       junk[nindex] = 1
           }
       }
	return junk
}
