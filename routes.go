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
	GetMeasurements(string, string, string, time.Time, time.Time) (Measurement, error)
	SetMeasurements(Measurementx) (error)
	GetRegisters(string, string) (RegistersInfos, error)
	GetLocationsClusters ()(LocationsInfos, error)
}

func NewRouter(db DataHandler) *mux.Router {
	
	fe := FrontEnd{DataHandler: db}

	var routes = Routes{
		Route{
                        "MeasurementsShow",
                        "GET",
                        "/measurements/location/{location}/serial/{serial}/reg/{reg}/start/{start}/end/{end}",
                        fe.MeasurementsShow,
                },
		Route{
                        "ShowLocationsClusters",
                        "GET",
                        "/locationsInfo",
                        fe.ShowLocationsClusters,
                },
		Route{
			"ShowRegistersInfo",
			"GET",
			"/registersInfo/location/{location}/serial/{serial}",
			fe.ShowRegistersInfo,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	
    	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
    	}
	
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))

    return router
}

type FrontEnd struct {
	DataHandler
}

func (fe FrontEnd) MeasurementsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
        l := vars["location"]
	ser := vars["serial"]
	reg := vars["reg"]

        start, err := time.Parse(time.RFC3339, vars["start"])
        if err != nil {
                log.Println(err)
        }

        end, err := time.Parse(time.RFC3339, vars["end"])
        if err != nil {
                log.Println(err)
        }
	
	var ms Measurement
	ms, err = fe.DataHandler.GetMeasurements(l,ser,reg,start,end)
	
	if err != nil {
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

func (fe FrontEnd) ShowRegistersInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loc := vars["location"]
	ser := vars["serial"]
	regs, err := fe.DataHandler.GetRegisters(loc,ser)
	if err != nil {
		log.Println(err)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	if err := json.NewEncoder(w).Encode(regs); err != nil {
		log.Println(err)
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


