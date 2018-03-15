package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)

var arr []string = []string{"string1"}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var output, _ = json.Marshal(arr)
		//fmt.Println(output)
		//fmt.Fprintf(w, string(output))
		w.Write(output)
	case "POST":
		var data, err = ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request Body", http.StatusInternalServerError)
		}
		arr = append(arr, string(data))
	}
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handle)

	fmt.Println("Server started on port 9001...")
	fmt.Println(http.ListenAndServe(":9001", nil))
}
