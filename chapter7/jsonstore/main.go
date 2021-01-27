package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/4L3X4NND3RR/chapter7/jsonstore/helper"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type DBclient struct {
	db *gorm.DB
}

type PackageResponse struct {
	Package helper.Package `json:"Package"`
}

// GetPackage fetches a package
func (driver *DBclient) GetPackage(rw http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	vars := mux.Vars(r)

	driver.db.First(&Package, vars["id"])
	var PackageData interface{}

	json.Unmarshal([]byte(Package.Data), &PackageData)
	var response = PackageResponse{Package: Package}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	respJson, _ := json.Marshal(response)
	rw.Write(respJson)
}

// GetPackagesbyWeight fetches all packages with given weight
func (driver *DBclient) GetPackagesbyWeight(rw http.ResponseWriter, r *http.Request) {
	var packages []helper.Package
	weight := r.FormValue("weight")
	// Handle response detail
	var query = "SELECT * FROM \"Package\" WHERE data->>'weight'=?"
	driver.db.Raw(query, weight).Scan(&packages)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(packages)
	rw.Write(respJSON)
}

// PostPackage saves the package information
func (driver *DBclient) PostPackage(rw http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	postBody, _ := ioutil.ReadAll(r.Body)
	Package.Data = string(postBody)
	driver.db.Save(&Package)
	responseMap := map[string]interface{}{"id": Package.ID}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	rw.Write(response)
}

func main() {
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBclient{db: db}
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/v1/package/{id:[a-zA-Z0-9]*}", dbclient.GetPackage).Methods("GET")
	r.HandleFunc("/v1/package", dbclient.PostPackage).Methods("POST")
	r.HandleFunc("/v1/package", dbclient.GetPackagesbyWeight).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
