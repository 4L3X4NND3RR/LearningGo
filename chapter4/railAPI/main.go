package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/4L3X4NND3RR/chapter4/dbutils"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	ID int
	DriverName string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID int
	Name string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID int
	TrainID int
	StationID int
	ArrivalTime time.Time
}

// Register adds paths and routes to a new service instance
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

// GET http://localhost:8000/v1/trains/1
func (t TrainResource) getTrain(rq *restful.Request, rp *restful.Response) {
	id := rq.PathParameter("train-id")
	err := DB.QueryRow("SELECT ID, DRIVER_NAME, OPERATING_STATUS FROM train WHERE id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		rp.AddHeader("Content-Type", "text/plain")
		rp.WriteErrorString(http.StatusNotFound, "Train could not be found")
	} else {
		rp.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
func (t TrainResource) createTrain(rq *restful.Request, rp *restful.Response) {
	log.Println(rq.Request.Body)
	decoder := json.NewDecoder(rq.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	// Error handling is obvious here. So omitting...
	statement, _ := DB.Prepare("INSERT INTO train (DRIVER_NAME, OPERATING_STATUS) VALUES (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		rp.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		rp.AddHeader("Content-Type", "text/plain")
		rp.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8000/v1/trains/1
func (t TrainResource) removeTrain(rq *restful.Request, rp *restful.Response) {
	id := rq.PathParameter("train-id")
	statement, _ := DB.Prepare("DELETE FROM train WHERE id=?")
	_, err := statement.Exec(id)
	if err == nil {
		rp.WriteHeader(http.StatusOK)
	} else {
		rp.AddHeader("Content-Type", "text/plain")
		rp.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
