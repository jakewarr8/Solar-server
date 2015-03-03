package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

func NewFetcher(db DataHandler) {
	
	r, err := http.Get("http://egauge7055.egaug.es/cgi-bin/egauge?inst&v1")
	if err != nil {
		fmt.Println("Egauge Appers to be down!!!")
		return;
	}
	defer r.Body.Close()	
	
	contents, err := ioutil.ReadAll(r.Body)
	checkError(err)	

	mx := Measurementx{}
	err = xml.Unmarshal(contents,&mx)
	checkError(err)

//	fmt.Printf("%s\n", string(contents))
//	fmt.Println(mx)
	
	fmt.Println(time.Now().Format(time.RFC850))

	mx.Location = "TxState"
	mx.Serial = "0001"
	mx.TimeS = time.Now()		
	
	db.SetMeasurements(mx)			
	
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())	
	}
}
