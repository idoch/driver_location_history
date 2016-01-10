package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"fmt"

	_ "github.com/avct/prestgo"
)

type PointEvent struct {
	DriverID     int
	Lat          float64
	Lon          float64
	DriverStatus string
	OrderID      int
	GpsAt        time.Time
	Bearing      float64
	GpsDate      string
	Country      string
}

func alive(res http.ResponseWriter, req *http.Request) {
	renderJSON(res, http.StatusOK, map[string]interface{}{"alive": true})
}

func getDriverLocation(res http.ResponseWriter, req *http.Request) {
	env := req.URL.Query().Get(":env")
	driverID := req.URL.Query().Get(":ID")
	startDate := req.URL.Query().Get("startDate")
	endDate := req.URL.Query().Get("endDate")
	db, err := sql.Open("prestgo", Settings["database"].GetString(environments+".host"))
	if err != nil {
		log.Fatalf("failed to connect to presto: %v", err)
	}
	rows, err := db.Query(fmt.Sprintf(" SELECT * "+
		" FROM drivers_locations_s "+
		" WHERE gps_date >= '%s' and  gps_date < '%s' and country = '%s' AND drivergk = %s",
		startDate, endDate, env, driverID))
	if err != nil {
		log.Fatalf("failed to run query: %v", err)
	}

	defer rows.Close()
	points := []PointEvent{}
	// point := PointEvent{
	// 	Lat:          33.749990,
	// 	Lon:          -117.942114,
	// 	DriverID:     1,
	// 	DriverStatus: "Free",
	// 	OrderID:      2,
	// 	Bearing:      1,
	// 	Country:      "US",
	// }
	// points = append(points, point)
	for rows.Next() {
		var point PointEvent
		err := rows.Scan(
			&point.DriverID,
			&point.Lat,
			&point.Lon,
			&point.DriverStatus,
			&point.OrderID,
			&point.GpsAt,
			&point.Bearing,
			&point.GpsDate,
			&point.Country)
		if err != nil {
			log.Println(err.Error())
		}
		points = append(points, point)
	}

	if err := rows.Err(); err != nil {
		log.Println(err.Error())
	}

	renderJSON(res, http.StatusOK, points)
}

func renderJSON(r http.ResponseWriter, status int, v interface{}) {
	var result []byte
	var err error
	result, err = json.Marshal(v)
	if err != nil {
		http.Error(r, err.Error(), 500)
		return
	}
	// json rendered fine, write out the result
	r.Header().Set("Content-Type", "application/json")
	r.WriteHeader(status)
	r.Write(result)
}
