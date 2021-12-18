var ws
var pause = false
FOODXY = [] 

function WebsocketStart() {
	console.log("WEBSOCKET START")
      	setup();
    ws = new WebSocket("ws://localhost:8081/talk")
    ws.onopen = function(evt) {
	junk = {}
	junk['Width'] = width;
	junk['Height'] = height;
	junk['Epochs'] = 100;
	junk['Food'] = FOODXY;
	junk['msg_type'] = "make_arena";	
        senddata(junk);
      	myGameArea.start(); 
    }
    ws.onclose = function(evt) {
      console.log('WEBSOCKET CLOSE');
      myGameArea.stop();
      //ws = null;
    }

    ws.onmessage = function(e) {
      	n = e.data.indexOf("position");
      	if (n != -1 ) {
	 var response = JSON.parse(e.data)
         positions = response.Positions
	      if (ROVERS.length <= 0) {
         	for (var pos=0;pos < response.Positions.length;pos++) {
        		ROVERS[pos] = new Rover(response.Positions[pos]);
    		}
	      } //end of if on length
         for (var pos=0;pos < response.Positions.length;pos++) {
        	ROVERS[pos].x = response.Positions[pos][0];
        	ROVERS[pos].y = response.Positions[pos][1];
    	}
	
      } //end of if found 'positions'
    } //endo of onmessage


    ws.onerror = function(evt) {
        console.log('onerror ',evt.data);
    }

} //end of WebsocketStart

senddata = function(data) {
    if (pause) {
      return;
    }
    if (!ws) {
        console.log('cannot send data -- no ws');
        return false;
    }
    stuff = JSON.stringify(data);
    ws.send(stuff);
} //end of function senddata

function setup() {
    make_foods();
    reset_food_positions();

} //end of setup
    
function updateGameArea() {
    if (pause) {
       return
    }
       var mydata = {};
       //reset_rover_positions();
       //mydata['num_episodes'] =  num_episodes;
       //senddata(mydata);
       reset_food_positions();

    myGameArea.clear();
    update_rovers();
    update_foods();
} //end of updateGameArea

myGameArea = {
    canvas : document.createElement("canvas"),
    start : function() {
        this.millis = 75;  //game intervale milliseconds
        this.canvas.width = width;
        this.canvas.height = height;
        this.context = this.canvas.getContext("2d");
        document.body.insertBefore(this.canvas, document.body.childNodes[0]);
        pause = false;
        this.interval = setInterval(updateGameArea,this.millis);
    },  
    stop : function() {
        pause = true; 
        console.log("STOP !!! ");
        clearInterval(this.interval);
        //ws.close();
    },  
    clear : function() {
        this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
        this.context.fillStyle = "rgba(255,255,255,255)";
        this.context.fillRect(0,0,this.canvas.width,this.canvas.height);
    } 
}    //end of gamearea


