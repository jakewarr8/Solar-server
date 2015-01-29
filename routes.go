package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"time"
	"io/ioutil"
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
type Page struct {
	Title string
	Body  []byte
}
 
func loadPage(title string) (*Page, error) {
    filename := title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
  
func MeasurementsIndex(w http.ResponseWriter, r *http.Request) {	
	p, err := loadPage("measurements")
	if err != nil {
		fmt.Fprintln(w,"Error Loading Page ^_^. Try Again Later.") 
		return
	}
	bs := string(p.Body[:])
	fmt.Fprintf(w,bs)
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


