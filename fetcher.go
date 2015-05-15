package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

func NewFetcher(db DataHandler) {
	
	r, err := http.Get("http://egauge7056.egaug.es/cgi-bin/egauge?inst&v1")
	if err != nil {
		fmt.Println("Egauge Appers to be down!!!")
		return;
	}
	defer r.Body.Close()	
	
	contents, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}

	mx := Measurementx{}
	err = xml.Unmarshal(contents,&mx)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}	
	
/*
//	fmt.Printf("%s\n", string(contents))
//	fmt.Println(mx)	
//	fmt.Println(time.Now().Format(time.RFC850))	
//	fmt.Println(time.Now().Local())
*/

	mx.Location = "TxState"
	mx.TimeS = time.Now()		
	
	db.SetMeasurements(mx)			
	
}

