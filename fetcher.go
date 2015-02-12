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

	m := Measurement{} //ToStore
	m.Location= "TxState"
	m.Time = time.Now().Local()		
	for _,e := range mx.KeyPairs {
		kp := KeyPair{Nk: e.Nk, Tk: e.Tk, Data: e.Data,}
		m.Registers = append(m.Registers, kp)
		//fmt.Println(m.Registers)
	}
	ms := Measurements{m}
	//fmt.Println(ms)
	
	db.SetMeasurements(ms)			
	
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())	
	}
}
