package main

import (
	"encoding/json"
	"fmt"
	"globals"
	"net/http"
	"os"
	"pq"
)

// JSONIn :  incoming outer JSON
type JSONIn struct {
	List []JSONInInner `json:"inList"`
}

// JSONInInner : incoming inner JSON
type JSONInInner struct {
	Cmd string `json:"cmd"`
	Key string `json:"name"`
	Pri int    `json:"pri"`
}

// JSONOut : JSON to return if successful
type JSONOut struct {
	List []string `json:"outList"`
}

// JSONError : JSON to return if error
type JSONError struct {
	Msg string `json:"message"`
}

// main entry point of program
func main() {

	// AWS EBS Stuff
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// testing endpoints
	http.HandleFunc("/priorityqueue", func(w http.ResponseWriter, r *http.Request) {

		// POST request
		if r.Method == "POST" {
			w.Header().Set("Content-Type", "application/json")

			// handle incoming JSON, try and decode it
			var data *JSONIn
			decoder := json.NewDecoder(r.Body)
			e := decoder.Decode(&data)

			// this will check that data types are proper, and upper level keys match
			if e != nil || data.List == nil {
				emsg := JSONError{Msg: "Malformed JSON."}
				b, _ := json.Marshal(emsg)
				fmt.Fprint(w, string(b))
				return
			}

			// create new pq
			queue := new(pq.PriorityQueue)

			// for each item in the incoming list, ensure that the command is either enqueue or dequeue, other than that
			// we dont really care what everything else is (already validated that pri is an int, name is a string)
			for i := 0; i < len(data.List); i++ {
				if data.List[i].Cmd != "enqueue" && data.List[i].Cmd != "dequeue" {
					emsg := JSONError{Msg: "Malformed JSON."}
					b, _ := json.Marshal(emsg)
					fmt.Fprint(w, string(b))
					return
				}

				// cmds are only enqueue and dequeue at this point, do the appropriate action
				if data.List[i].Cmd == "enqueue" {
					queue.Enqueue(globals.Item{Pri: data.List[i].Pri, Key: data.List[i].Key})
				} else {
					queue.Dequeue()
				}
			}

			// for each remaining item in the queue, dequeue it and add it to the outJSON
			outJSON := JSONOut{}
			for !queue.IsEmpty() {
				outJSON.List = append(outJSON.List, queue.Dequeue().Key)
			}

			// return empty string instead of Null
			if outJSON.List == nil {
				outJSON.List = append(outJSON.List, "")
			}

			// return dequeued items to the handler
			fmt.Printf("%v\n", data.List)
			b, _ := json.Marshal(outJSON)
			fmt.Fprint(w, string(b))

			// GET request
		} else {

			b, _ := json.Marshal("Hello! Please send a post request.")
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(b))
		}
	})

	// serve requests
	http.ListenAndServe(":"+port, nil)
}
