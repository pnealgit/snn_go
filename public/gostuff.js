function Rover(xy) {
	console.log("XY IN ROVER: ",xy)
	//start in the middle
    this.r = 10;
    this.sensor_data = xy;
    //console.log("SENSOR DATA IN NEW: ",this.sensor_data)

    this.draw = function() {
	    //console.log("DRAW SENSOR DATA XY: ",this.sensor_data)
	    //console.log("NUM SENSOrS: ",NUM_SENSORS)
	x = this.sensor_data.shift()
	y = this.sensor_data.shift()
	//console.log("center x,y:",x,y);
        ctx = myGameArea.context;
        ctx.beginPath();
        ctx.arc(x,y, this.r, 0, 2 * Math.PI);
        ctx.fillStyle = "red";
        ctx.fill();
        ctx.beginPath();
        ctx.strokeStyle = '#000000';
        ctx.stroke();
        ctx.closePath();

	      for(i =0 ;i<NUM_SENSORS;i++) {
		xp = this.sensor_data.shift()
		yp = this.sensor_data.shift()
                ctx.beginPath()
                ctx.strokeStyle = '#000000';
                ctx.moveTo(x,y);
                ctx.lineTo(xp,yp);
                ctx.stroke();
                ctx.closePath();
        }       

    } //end of rover draw
}
//end of Rover function

function draw_rovers() {
	//console.log("IN DrAW roVerS")
    for (var i = 0; i < ROVERS.length; i++) {
        ROVERS[i].draw();
    } //end of loop on rovers
}
//end of function 

