package main

import (
	"time"
	"encoding/json"
)

type KeyPair struct {
        Nk      string          `json:"name"`  //CT1 L1F L1V
        Tk      string          `json:"type"`  //I,V,F
        Data    float64         `json:"data"`
}

type Measurement struct {
	Time 		time.Time	`json:"time"`
	Location 	string		`json:"location"`
	Registers	[]KeyPair	`json:"registers"`
}

func (m Measurement) RegistersToJson() ([]byte, error) {
	jsonString, err := json.Marshal(m.Registers)
	return jsonString, err
}

//How to refence self
func (m *Measurement) ParseRegisters(j []byte) (error) {
	err := json.Unmarshal(j, &m.Registers)
	return err
}

type Measurements []Measurement


type LocationInfo struct {
	LocationAbbrv	string		`json:"location"`
 	Serials		[]string	`json:"serials"`
}

type LocationsInfos []LocationInfo


//XML Structs
type KeyPairx struct {
        Nk      string          `xml:"n,attr"`  //CT1 L1F L1V
        Tk      string          `xml:"t,attr"`  //I,V
        Data    float64         `xml:"i"`
}

type Measurementx struct {
	KeyPairs	[]KeyPairx	`xml:"r"`	
	Time		int64		`xml:"ts"`
}

