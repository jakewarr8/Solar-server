package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"time"
	"io/ioutil"
	"strings"
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
			"HomeIndex",
			"GET",
			"/home/{path}",
			HomeIndex,
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
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
    	}

    return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

type Page struct {
	Title string
	Body  []byte
}
 
func loadPage(title string) (*Page, error) {
    filename := title
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
  
func HomeIndex(w http.ResponseWriter, r *http.Request) {	
	vars := mux.Vars(r)
	title := vars["path"]
	p, err := loadPage(title)
	if err != nil {
		fmt.Fprintln(w,"Error Loading Page ^_^. Try Again Later.")
		return
	}
	
	bs := string(p.Body[:])
	
	if strings.Contains(title, "css") {
		log.Println("CSS")
		w.Header().Set("Content-Type", "text/css; charset=UTF-8")
	} else if strings.Contains(title, "js") {
		log.Println("JS")
		w.Header().Set("Content-Type", "text/javascript; charset=UTF-8")
	}
	w.WriteHeader(http.StatusOK)
	
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


