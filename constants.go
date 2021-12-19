package main

//binary group size
var BGS = 4
const NUM_SENSORS = 3

//gotta have 1 neuron per binary digit
//4 digits for distance code
//4 digits for type code
//so for now NUM_NEURONS should equal 24
//const NUM_NEURONS = 2*NUM_SENSORS*BGS
const  NUM_NEURONS = 24
const NUM_ROVERS  = 10
var THRES = 3
var LEAKING_CONSTANT  = 1
var SETTLING_TIME = 20
//e,ne,n,nw,w,sw,s,se
var ANGLES_DX = [8]int{1, 1, 0, -1, -1, -1, 0, 1}
var ANGLES_DY = [8]int{0, 1, 1, 1, 0, -1, -1, -1}
var NUM_ANGLES = 8
var SENSOR_LENGTH = 30
var FOOD_RADIUS = 15
