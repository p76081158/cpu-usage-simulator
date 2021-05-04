package main

import (
	"fmt"
	"os/exec"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL.Path)
	input_cmd := "nohup stress-ng -c 1 -l 3 -t 5"
	cmd := exec.Command("/bin/sh", "-c", input_cmd)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
