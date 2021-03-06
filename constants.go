package main

const NUM_SENSORS = 8
var SENSOR_LENGTH = 200 

//gotta have 1 neuron per binary digit
//4 digits for distance code
//4 digits for type code
//so for now NUM_NEURONS should equal 24

//NUM_SENSORS * 4 * 4
const NUM_NEURONS = 16
const NUM_ROVERS = 20

var THRES = 8
var LEAKING_CONSTANT = 1
var SETTLING_TIME = 20

//e,ne,n,nw,w,sw,s,se
var ANGLES_DX = [8]int{1, 1,  0, -1, -1, -1, 0, 1}
var ANGLES_DY = [8]int{0,-1, -1, -1,  0,  1, 1, 1}
var NUM_ANGLES = 8
var FOOD_RADIUS = 15
var NUM_MAX_STEPS = 1200
