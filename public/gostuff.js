function Rover(xy) {
	//console.log("XY IN ROVER: ",xy)
	//start in the middle
    this.r = 10;
    this.sensor_data = xy;
    //console.log("SENSOR DATA IN NEW: ",this.sensor_data)

    this.draw = function() {
//	    console.log("SENSOR DATA XY: ",this.sensor_data)
	x = this.sensor_data.shift()
	y = this.sensor_data.shift()
	    //console.log("x,y:",x,y);
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

    for (var i = 0; i < ROVERS.length; i++) {
        //ROVERS[i].move();
        ROVERS[i].draw();
    }
    //end of loop on rovers
}
//end of function 

function reset_rover_positions(rovers) {
    sum = 0;
    best = -9999;
    worst = 9999;

    for (var nn = 0; nn < NUM_ROVERS; nn++) {
        rovers[nn].reset_position();

        r = rovers[nn].reward;
        if (r > best) {
            best = r;
        }
        if (r < worst) {
            worst = r;
        }
        sum += r
        rovers[nn].reward = 0;
    }
    //end of loop
    console.log("SUM: \t", sum, "\tBEST:\t", best, "\tWORST:\t", worst);
}

Rover.prototype.reset_position = function() {
    this.x = width / 2 + getRandomInt(-5, 5);
    this.y = height / 2 + getRandomInt(-5, 5);
    junk = getRandomInt(0, 4)
    this.angle = junk * Math.PI / 2;
    this.velocity = 2.0;
}
