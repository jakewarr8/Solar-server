package main

import (
	"log"	
	"time"
	"net/http"
)

func main() {
	var err error
	db, err := NewOpen("mysql","jake:password@/solar")

        if err != nil {
                log.Fatal(err)
       	}

        defer db.Close();

	ticker := time.NewTicker(time.Minute)
	go func() {
        	for _ = range ticker.C {
			NewFetcher(db)
       		}
	}()
	
	router := NewRouter(db)
	log.Fatal(http.ListenAndServe(":8080", router))
}
