package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {

	//Logger
	// var handler http.Handler
	// handler = route.HandlerFunc
	// handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

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
        "MeasurementShow",
        "GET",
        "/measurements/{measurementId}",
        db.MeasurementShow,
    },
    Route{
        "MeasurementPut",
        "PUT",
        "/measurements/insert",
        db.MeasurementPut,
    },
}
