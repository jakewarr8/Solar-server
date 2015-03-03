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

func (d DB) GetMeasurements(l string, s string, r string, st time.Time, et time.Time) (m Measurement, err error) {
	log.Println("GetMS",l,s,r,st,et) 	
	var query = "SELECT time, data FROM measurements WHERE location=? AND serial=? AND register=? AND time>=? AND time<?;"
	rows, err := d.Query(query, l, s, r, st, et)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	m.Location = l
	m.Serial = s
	m.Register = r
	for rows.Next() {
		p := Point{}
		var t string
		err = rows.Scan(&t, &p.Value)
		if err != nil {
			log.Println(err)
		} else {
			p.Time, err = time.Parse("2006-01-02 15:04:05", t)
			m.Data = append(m.Data, p)
		}
	}


	return m,err
}

func (d DB) SetMeasurements(m Measurementx) (err error) {
	
	// Create tx
	tx, err := d.Begin()	
	if err != nil {
		log.Println(err)
		return
	}

	var query = "INSERT INTO measurements (location,serial,time,register,type,data) VALUES (?, ?, ?, ?, ?, ?);"
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
	
	for _,r := range m.KeyPairs {
		_, err = stmt.Exec(m.Location, m.Serial, m.TimeS, r.Nk, r.Tk, r.Data)
		if err != nil {
			log.Println(err)
			return
		}
	}
	
	return 
}

func (d DB) GetLocationsClusters ()(locInfos LocationsInfos,  err error) {
	log.Println("GetLocationsClusters")
	var query = "SELECT location, serial FROM measurements GROUP BY location, serial;"
	rows, err := d.Query(query)
	
	if err != nil {
                log.Println(err)
                return
        }
        defer rows.Close()

	x := make(map[string]LocationInfo)	
	
        for rows.Next() {        
		var l string 
		var serial string
                
		err = rows.Scan(&l, &serial)
                //log.Println(l,clusterid)
		

		if err != nil {
			log.Println(err)
			return
		} else {
			locinfo, ok := x[l]
			if ok {
				locinfo.Serials = append(locinfo.Serials,serial)
				x[l] = locinfo
				//log.Println(x)
			} else {
				locinfo.LocationAbbrv = l
				locinfo.Serials = append(locinfo.Serials,serial)
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

func (d DB) GetRegisters(loc string, ser string) (regs RegistersInfos, err error){
	log.Println("GetRegisters ", loc, ser)
	var query = "SELECT register, type FROM measurements WHERE location=? AND serial=? GROUP BY register, type;"
	rows, err := d.Query(query,loc,ser)
	
	if err != nil {
		log.Println(err)
		return
	}	
	defer rows.Close()

	for rows.Next() {
		r := Register{}
		err = rows.Scan(&r.Name, &r.Type)
		if err != nil {
			log.Println(err)
		} else {
			regs = append(regs, r)
		}
	}

	return
}
