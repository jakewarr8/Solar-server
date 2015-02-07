package main

import (
	"log"	
	"net/http"
	"time"
)

func main() {
	var err error
	db, err := NewOpen("mysql","josh:password@/solar")

        if err != nil {
                log.Fatal(err)
       	}

        defer db.Close(); //????

	ticker := time.NewTicker(time.Minute)
	go func() {
        	for _ = range ticker.C {
			NewFetcher(db)
       		}
	}()
	
	router := NewRouter(db)
	log.Fatal(http.ListenAndServe(":8080", router))
}
