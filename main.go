package main

import (
	"log"	
	"net/http"
)

var db DB

func main() {
	var err error
	db, err = NewOpen("mysql","jake:password@/solar")

        if err != nil {
                log.Fatal(err)
       	}

        defer db.Close(); //????

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
