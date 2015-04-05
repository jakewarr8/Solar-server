package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
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
	LastMeasurement(l string, s string, r string) (p Point, err error)
	GetMeasurements(string, string, string, time.Time, time.Time) (Measurement, error)
	SetMeasurements(Measurementx) (error)
	GetLocationsClusters ()(LocationsInfoSets, error)
}

func NewRouter(db DataHandler) *mux.Router {
	
	fe := FrontEnd{DataHandler: db}

	var routes = Routes{
		Route{
			"MobileView",
			"POST",
			"/mobile",
			MobileView,
		},
		Route{
			"LastMeasurement",
			"GET",
			"/lastmeasurement/loc/{loc}/ser/{ser}/reg/{reg}",
			fe.LastMeasurement,
		},
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
	}

	router := mux.NewRouter().StrictSlash(true)
	
    	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
    	}
	
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))

	return router
}

func MobileView(w http.ResponseWriter, r *http.Request){
	log.Println("EHEHEHEHEHEHEHEHEHEH")
	var data LocationTables
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&data); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
	} else {
		//w.WriteHeader(http.StatusOK)
		json, err := json.Marshal(data)	
		if err != nil {
			log.Println(err)
		}
		pagedata := &DataPayload{Data:string(json[:])}
		RenderTemplate(w,"templates/mobile.tmpl",pagedata)
		log.Println(data)
	}
	
}

type DataPayload struct {
	Data string 
}

func RenderTemplate(w http.ResponseWriter, tmlp string, data *DataPayload) { //data interface{}
	if data != nil {
		log.Println(data)
	}
	t, err := template.ParseFiles(tmlp)
	if err != nil {
		log.Println(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type FrontEnd struct {
	DataHandler
}

func (fe FrontEnd) LastMeasurement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loc := vars["loc"]
	ser := vars["ser"]
	reg := vars["reg"]
	
	p, err := fe.DataHandler.LastMeasurement(loc, ser, reg) 
	if err != nil {
		fmt.Fprintln(w,"No Measurement Found.")
		
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(p); err != nil {
			log.Println(err)
			fmt.Fprintln(w,"No Measurement Found...")
		}
	}
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


