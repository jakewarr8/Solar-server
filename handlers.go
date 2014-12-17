package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"	
	"log"
	"time"
)

type DB struct {
	*sql.DB
}

func NewOpen(dt string, c string) (DB, error) {
        db, err := sql.Open(dt,c)
	return DB{db}, err
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func MeasurementsIndex(w http.ResponseWriter, r *http.Request) {
	measurements := Measurements{
		Measurement{ClusterID: 1, 
			    Time: time.Date(2014, time.December, 15, 23, 12, 53, 0, time.UTC),
        		    Location: "SM",
			    Voltage: 3.77432,
			    Ampere: 1.83551,
			    Humidity: 0.18,
			    Temp: 18.3,
			    AngleTheta: 0.00,
			    AngleAlpha: 45.0,
			    SpTemp: 25.6,
			    },
		Measurement{ClusterID: 2,
                            Time: time.Date(2014, time.December, 16, 12, 53, 0, 0, time.UTC),
                            Location: "SA",
                            Voltage: 3.77432,
                            Ampere: 1.83551,
                            Humidity: 0.18,
                            Temp: 18.3,
                            AngleTheta: 0.00,
                            AngleAlpha: 45.0,
                            SpTemp: 25.6,
                            },
	}
	
		
	//Added StatusCode for JSON	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    	w.WriteHeader(http.StatusOK)	
	
	if err := json.NewEncoder(w).Encode(measurements); err != nil {
		panic(err)
	}
}

func (db DB) MeasurementShow(w http.ResponseWriter, r *http.Request) {
   	vars := mux.Vars(r)
    	measurementId := vars["measurementId"]
    	fmt.Fprintln(w, "Todo show:", measurementId)

	id , err := strconv.Atoi(measurementId)
	if err != nil {
		log.Println(err)
	} else { 

		var (
        	        query = "SELECT time, location FROM measurements WHERE id=?;"
			time time.Time
               		location string
        	)
        	err = db.QueryRow(query, id).Scan(&time, &location)
        	if err != nil {
                	log.Fatal(err)
        	} else {
                	fmt.Fprintln(w, "TIME:", time, " LOC:", location)
        	}
	}

}

func (db DB) MeasurementPut(w http.ResponseWriter, r *http.Request) {
	var m Measurement
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
	} else {
		//db.Query()
		fmt.Fprintln(w,m)
	}
}
