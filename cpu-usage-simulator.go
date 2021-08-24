package main

import (
	"fmt"
	"os"
	"os/exec"
	"net/http"
)

// stress environment variable of service
var cpu  = "1"
var load = "3"
var time = "5"

// handle http requests
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
	input_cmd := "stress-ng -c " + cpu + " -l " + load + " -t " + time
	cmd       := exec.Command("/bin/sh", "-c", input_cmd)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}

func main() {
    if (os.Args[2]!="") {
		cpu = string(os.Args[2])
	}
	if (os.Args[3]!="") {
		load = string(os.Args[3])
	}
	if (os.Args[4]!="") {
		time = string(os.Args[4])
	}

	// simulate service listen on port 9090
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}