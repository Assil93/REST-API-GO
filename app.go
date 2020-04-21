package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/Assil/Go_Training/config"
	. "github.com/Assil/Go_Training/dao"
	. "github.com/Assil/Go_Training/models"
)

var config = Config{}
var dao = CarsDAO{}

// GET list of cars
func AllCarsEndPoint(w http.ResponseWriter, r *http.Request) {
	cars, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, cars)
}

// GET a car by its ID
func FindCarEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	car, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, car)
}

// POST a new car
func CreateCarEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	car.ID = bson.NewObjectId()
	if err := dao.Insert(car); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, car)
}

// PUT update an existing car
func UpdateCarEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(car); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing car
func DeleteCarEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(car); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	fmt.Println("Starting Server")
	r.HandleFunc("/cars", AllCarsEndPoint).Methods("GET")
	r.HandleFunc("/cars", CreateCarEndPoint).Methods("POST")
	r.HandleFunc("/cars", UpdateCarEndPoint).Methods("PUT")
	r.HandleFunc("/cars", DeleteCarEndPoint).Methods("DELETE")
	r.HandleFunc("/cars/{id}", FindCarEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3300", r); err != nil {
		log.Fatal(err)
	}
}