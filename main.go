package main

import (
	"log"	
	"net/http"
)

func main() {
	var err error
	db, err := NewOpen("mysql","jake:password@/solar")

        if err != nil {
                log.Fatal(err)
       	}

        defer db.Close(); //????

	//fe := FrontEnd{DataHandler: db}
	router := NewRouter(db)
	log.Fatal(http.ListenAndServe(":8080", router))
}
