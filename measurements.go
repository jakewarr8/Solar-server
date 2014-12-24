package main

import (
	//"database/sql"
	"time"
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

