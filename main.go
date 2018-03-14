package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	var arr []string
	arr = make([]string, 0)
	arr = append(arr, "String0")
	arr = append(arr, "String1")
	arr = append(arr, "String2")
	var output, _ = json.Marshal(arr)
	//fmt.Println(output)
	//fmt.Fprintf(w, string(output))
	w.Write(output)
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handle)

	fmt.Println("Server started on port 8080...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
