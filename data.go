package main 

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"log"
)

type DB struct {
	*sql.DB
}

func NewOpen(dt string, c string) (DB, error) {
        db, err := sql.Open(dt,c)
        return DB{db}, err
}

func (d DB) GetMeasurement(l string, id int64, t time.Time) (Measurement, error) {
	var (
		query = "SELECT V, I, humidity, temp, angleTheta, angleAlpha, spTemp FROM measurements WHERE location=? AND clusterID=? AND time=?;"
		m = Measurement{ClusterID: id, Time: t, Location: l}
		
	)
	err := d.QueryRow(query, l, id, t).Scan(&m.Voltage, &m.Ampere, &m.Humidity, &m.Temp, &m.AngleTheta, &m.AngleAlpha, &m.SpTemp)
	if err != nil {
		return Measurement{}, err
	}
	return m, nil
}

func (d DB) GetMeasurements(l string, id int64, st time.Time, et time.Time) (ms Measurements, err error) {
	var (
                query = "SELECT time, V, I, humidity, temp, angleTheta, angleAlpha, spTemp FROM measurements WHERE location=? AND clusterID=? AND time>=? AND time<?;"
	)
	//var rows sql.Rows
	rows, err := d.Query(query, l, id, st, et)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	
	for rows.Next() {
		m := Measurement{ClusterID: id, Location: l}
		var t string//[]uint8
		err = rows.Scan(&t, &m.Voltage, &m.Ampere, &m.Humidity, &m.Temp, &m.AngleTheta, &m.AngleAlpha, &m.SpTemp)	
		log.Println(t)
		if err != nil {
                	log.Println(err)
		} else {
			m.Time, err = time.Parse("2006-01-02 15:04:05", t)
			ms = append(ms, m)
		}
	}
	return ms, rows.Err()
}

func (d DB) SetMeasurements(ms Measurements) (err error) {

	// Create tx
	tx, err := d.Begin()	
	if err != nil {
		log.Println(err)
		return
	}

	var query = "INSERT INTO measurements (time,location,clusterID,V,I,humidity,temp,angleTheta,angleAlpha,spTemp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
        stmt, err := tx.Prepare(query)	
	if err != nil {
		log.Println(err)
		return
	}
	
	// Defer
        defer func () {
                if err == nil {
                        log.Println("Commit")
                        tx.Commit()
                } else {
                        log.Println("RollBack")
                        tx.Rollback()
                }
		stmt.Close()
        }()
	
	for _,m := range ms {
			
		_, err = stmt.Exec(m.Time, m.Location, m.ClusterID, m.Voltage, m.Ampere, m.Humidity, m.Temp, m.AngleTheta, m.AngleAlpha, m.SpTemp)
		if err != nil {
			log.Println(err)
			return
		}
	}
	
	return 
}

func (d DB) GetLocationsClusters ()(locInfos LocationsInfos,  err error) {
	var query = "SELECT location, clusterID FROM measurements GROUP BY location, clusterID;"
	rows, err := d.Query(query)
	
	if err != nil {
                log.Println(err)
                return
        }
        defer rows.Close()

	x := make(map[string]LocationInfo)	
	
        for rows.Next() {        
		var l string 
		var clusterid int64
                
		err = rows.Scan(&l, &clusterid)
                //log.Println(l,clusterid)
		

		if err != nil {
			log.Println(err)
			return
		} else {
			locinfo, ok := x[l]
			if ok {
				locinfo.ClusterIDs = append(locinfo.ClusterIDs,clusterid)
				x[l] = locinfo
				log.Println(x)
			} else {
				locinfo.LocationAbbrv = l
				locinfo.ClusterIDs = append(locinfo.ClusterIDs,clusterid)
				x[l] = locinfo
				//log.Println(x)
			}
		}
        }

	for _, value := range x {
		locInfos = append(locInfos, value)
	}	
        return locInfos,rows.Err()
}

