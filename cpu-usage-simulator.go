package main

import (
	"fmt"
	"os"
	"os/exec"
	"net/http"
)

var cpu = "1"
var load = "3"
var time = "5"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
	input_cmd := "nohup stress-ng -c " + cpu + " -l " + load + " -t " + time
	cmd := exec.Command("/bin/sh", "-c", input_cmd)
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
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
