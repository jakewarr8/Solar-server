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
	Data		[][]interface{}	`json:"data"`
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
	Serial		string		`xml:"serial,attr"`
	TimeS		time.Time
}

//***Mobile***
type Table struct {
	Serial		string		`json:"serial"`
	Registers	[]string	`json:"regs"`
}

type LocationTable struct {
	Location	string		`json:"location"`
	Tables		[]Table		`json:"tables"`	
}

type LocationTables []LocationTable

//***LocationInfoV2***
type Serial struct {
	Name		string		`json:"serial"`
	Registers	[]Register	`json:"regs"`	
}

type Location struct {
	Name		string		`json:"location"`
	Serials		[]Serial	`json:"serials"`
}

type LocationsInfoSets []Location

