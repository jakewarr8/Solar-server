package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"io/ioutil"
)

func NewFetcher() {
	r, _ := http.Get("http://egauge7055.egaug.es/cgi-bin/egauge?inst&v1")
	defer r.Body.Close()	
	
	contents, err := ioutil.ReadAll(r.Body)
	checkError(err)	

	ms := &Measurementx{}
	err = xml.Unmarshal(contents,ms)
	checkError(err)
	fmt.Printf("%s\n", string(contents))
	fmt.Println(ms)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	
	}
}
