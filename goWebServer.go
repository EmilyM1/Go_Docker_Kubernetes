package main

import (

	"fmt"
	"log"
	"net/http"

	)

func main(){

	fmt.Println("Hello")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It works")
	})


	log.Fatal(http.ListenAndServe(":8081", nil))

}
