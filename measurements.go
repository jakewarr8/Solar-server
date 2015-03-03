package main

import (
	"time"
)

type Point struct {
	Time	time.Time	`json:"time"`
	Value	float64		`json:"value"`
}

type Measurement struct {
	Location	string		`json:"location"`
	Serial		string		`json:"serial"`
	Register	string		`json:"register"`
	Type		string		`json:"type"`
	Data		[]Point		`json:"data"`
}

type Measurements []Measurement

//***LocationsInfo***
type LocationInfo struct {
	LocationAbbrv	string		`json:"location"`
 	Serials		[]string	`json:"serials"`
}

type LocationsInfos []LocationInfo

//***RegistersInfo***
type Register struct {
	Name		string		`json:"name"`
	Type		string		`json:"type"`
}

type RegistersInfos []Register

//***XML Structs***
type KeyPairx struct {
        Nk      string          `xml:"n,attr"`  //CT1 L1F L1V
        Tk      string          `xml:"t,attr"`  //I,V
        Data    float64         `xml:"i"`
}

type Measurementx struct {
	KeyPairs	[]KeyPairx	`xml:"r"`	
	Time		int64		`xml:"ts"`
	Location	string
	Serial		string
	TimeS		time.Time
}

