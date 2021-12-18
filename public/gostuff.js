var Team = function(num_rovers,num_inputs,num_hidden,num_outputs) {
    this.team_name = "make_team";
    this.num_rovers = num_rovers;
    this.num_inputs = num_inputs;
    this.num_hidden = num_hidden;
    this.num_outputs = num_outputs;
}
//end of function

function make_rovers() {
    ROVERS = [];
    for (var ri = 0; ri < NUM_ROVERS; ri++) {
        ROVERS[ri] = new Rover(ri);
    }
    return rovers;
}
//end of function make_rovers

function Rover(xy) {
	//start in the middle
    this.x = xy[0]
    this.y = xy[1];
    this.r = 10;

    this.draw = function() {
        ctx = myGameArea.context;
        ctx.beginPath();
        ctx.arc(this.x, this.y, this.r, 0, 2 * Math.PI);
        ctx.fillStyle = "red";
        ctx.fill();

        ctx.beginPath();
        ctx.strokeStyle = '#000000';
        ctx.stroke();
        ctx.closePath();
    } //end of rover draw
}
//end of Rover function

function update_rovers() {

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

