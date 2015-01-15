package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"time"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

type DataHandler interface {
	GetMeasurements(l string, s string, st time.Time, et time.Time) (ms Measurements, err error)
	SetMeasurements(Measurements) (error)
	GetLocationsClusters ()(locInfos LocationsInfos,  err error)
}

func NewRouter(db DataHandler) *mux.Router {
	
	fe := FrontEnd{DataHandler: db}

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index,
	    	},
		Route{
			"MeasurementsIndex",
			"GET",
			"/measurements",
			MeasurementsIndex,
		},
		Route{
                        "MeasurementsShow",
                        "GET",
                        "/measurements/location/{location}/serial/{serial}/start/{start}/end/{end}",
                        fe.MeasurementsShow,
                },
		Route{
                        "ShowLocationsClusters",
                        "GET",
                        "/locationsInfo",
                        fe.ShowLocationsClusters,
                },
	}

	router := mux.NewRouter().StrictSlash(true)

    	for _, route := range routes {
		//Logger
		// var handler http.Handler
		// handler = route.HandlerFunc
		// handler = Logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
    	}

    return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi!")
}
//2014-12-12 14:07:00
func MeasurementsIndex(w http.ResponseWriter, r *http.Request) {
	/*
	measurements := Measurements{
		Measurement{ClusterID: 1, 
			    Time: time.Date(2014, time.December, 12, 14, 07, 00, 0, time.UTC),
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
	*/
		
	//Added StatusCode for JSON	
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    	//w.WriteHeader(http.StatusOK)	
	
	//if err := json.NewEncoder(w).Encode(measurements); err != nil {
	//	panic(err)
	//}
	
}

type FrontEnd struct {
	DataHandler
}

func (fe FrontEnd) MeasurementsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
        l := vars["location"]

        start, err := time.Parse(time.RFC3339, vars["start"])
        if err != nil {
                log.Println(err)
        }

        end, err := time.Parse(time.RFC3339, vars["end"])
        if err != nil {
                log.Println(err)
        }

        //fmt.Fprintln(w, "Todo show:", measurementId)

        ms, err := fe.DataHandler.GetMeasurements(l,vars["serial"],start,end)
        if err != nil || ms == nil{
                fmt.Fprintln(w,"No Measurement Found.")
        } else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        	w.WriteHeader(http.StatusOK)

        	if err := json.NewEncoder(w).Encode(ms); err != nil {
                	log.Println(err)
			fmt.Fprintln(w,"No Measurement Found...")
        	}
	}
}

func (fe FrontEnd) ShowLocationsClusters(w http.ResponseWriter, r *http.Request) {

	locInfos, err := fe.DataHandler.GetLocationsClusters() 
	if err != nil {
		log.Println(err)
	}

        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)

        if err := json.NewEncoder(w).Encode(locInfos); err != nil {
                log.Println(err)
        }

}


