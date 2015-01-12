package main

import (
	//"database/sql"
	"time"
	"encoding/xml"
)

type Measurement struct {
	ClusterID 	int64		`json:"clusterid"`
	Time 		time.Time	`json:"time"`
	Location 	string		`json:"location"`
	Voltage		float64 	`json:"voltage"`
	Ampere		float64		`json:"ampere"`
	Humidity	float32		`json:"humidity"`
	Temp		float32		`json:"temp"`
	AngleTheta	float32		`json:"theta"`
	AngleAlpha	float32		`json:"alpha"`
	SpTemp		float32		`json:"sptemp"`
}

type Measurements []Measurement


type LocationInfo struct {
	LocationAbbrv	string		`json:"location"`
 	ClusterIDs	[]int64		`json:"clusterids"`
}

type LocationsInfos []LocationInfo



type Measurementx struct {
	LVs		[]KeyPair	`xml:"r"`	
	ts		int64		`xml:"ts"`
	XMLName		xml.Name	`xml:"data"`

}

type KeyPair struct {
	nk	string		`xml:"n,attr"`	//CT1
	tk 	string		`xml:"t,attr"`	//I,V
	data	float64		`xml:"v>i"`	
}
