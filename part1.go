//  to run this file change name to part1.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address  := r.FormValue("address")
	// fmt.Fprintf is equivalent to printing to the screen
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
	


func helloHandler(w http.ResponseWriter, r *http.Request){
	// w and r are pointer

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w,"Get method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

// func main(){
// 	fileserver := http.FileServer(http.Dir("./static"))
// 	http.Handle("/", fileserver)
// 	http.HandleFunc("/form", formHandler) // shows the form
// 	http.HandleFunc("/hello",helloHandler) // runs the hello function

// 	fmt.Printf("Starting server at port 8080\n")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

