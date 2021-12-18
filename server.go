package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
)

//remember kids, Marshal only converts members of a struct if
//the name is Capitalized.... 55 minutes on that problem...
type Mess struct {
	Msg_type  string
	Positions [NUM_ROVERS][2]int
}

type Wall struct {
	xy [2]int
}

type Arena struct {
	Width  int
	Height int
	Food   [][2]int
	Epochs int
}

type Team struct {
	Num_rovers int
	Rovers     []Rover
}

type Brain struct {
	seed  int64
	sign  byte
	iconn []byte
	nconn [][]byte
}

type Rover struct {
	brain       Brain
	Xpos        int
	Ypos        int
	Sensor_data []byte
	Fitness     int
	Angle_index int
}

//var message_type int

var rovers []Rover
var arena Arena
var num_episodes int
var addr = flag.String("addr", "localhost:8081", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func talk(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	//big loop... maybe
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		} //end of if on message err

		fmt.Println("MESSAGE TYPE: ", mt)
		junk := string(message)
		fmt.Println("ARENA DATA: ", junk)
		if strings.Contains(junk, "make_arena") {
			fmt.Println("MAKE ARENA!!")
			jerr := json.Unmarshal(message, &arena)
			if jerr != nil {
				fmt.Println("error on arena unmarshal")
				os.Exit(3)
			} //end of if on jerr
			fmt.Println("ARENA WIDTH: ", arena.Width)
			fmt.Println("ARENA HEIGHT: ", arena.Height)
			fmt.Println("ARENA FOOD: ", arena.Food)
			make_rovers()
		} //end of if on arena

		//ok now we just spew data to web
		for {
		var draw_message []byte
		var draw_positions [NUM_ROVERS][2]int
		draw_positions = do_updates(rovers)
		fmt.Println("PAST DO UPDATES")
		var mmm Mess
		mmm.Msg_type = "positions"
		mmm.Positions = draw_positions
		draw_message, err = json.Marshal(mmm)
		if err != nil {
			fmt.Println("bad angles Marshal")
			os.Exit(7)
		}

		err = c.WriteMessage(mt, draw_message)
		if err != nil {
			log.Println("BAD DRAW MESSAGE:", err)
			os.Exit(4)
		} //end of if on write err
		num_episodes += 1
		if num_episodes > 100 {
			select_brains(rovers)
			mutate_brains(rovers)
			num_episodes = 0
		}
	} //end of for loop around do_updates
	} //end of for loop to receive and send messages

} //end of talk

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/talk", talk)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	fmt.Println("listening on 8081")
	log.Fatal(http.ListenAndServe(*addr, nil))
} //end of main
