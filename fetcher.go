package main

import (
	"encoding/xml"
	"log"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"time"
)

func NewFetcher(db DataHandler) {
	


	ss, err := db.GetSerials()
	if err != nil {
		log.Println(err)
		return
	}
	
	for _, s := range ss {
		
		var buffer bytes.Buffer		
		buffer.WriteString(fmt.Sprint("http://egauge", s.Name, ".egaug.es/cgi-bin/egauge?inst&v1"))
		url := buffer.String()
		
		r, err := http.Get(url)
		if err != nil || r.StatusCode != 200  {
			log.Println("Egauge Appers to be down!!!")
		}
		defer r.Body.Close()
			
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Fatal error ", err.Error())
		}
		
		mx := Measurementx{}
		err = xml.Unmarshal(contents,&mx)
		if err != nil {
			log.Println("Fatal error ", err.Error())
		}

		u, err := db.GetUserWithId(s.User_Id)
		mx.Location = u.UserName
		mx.Serial = s.Name
		mx.TimeS = time.Now()
	
		if err == nil {
			db.SetMeasurements(mx)
		}

	}	
	

	
}

