package dao

import (
	"log"

	. "github.com/Assil/Go_Training/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CarsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "cars"
)

// Establish a connection to database
func (c *CarsDAO) Connect() {
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(c.Database)
}

// Find list of cars
func (c *CarsDAO) FindAll() ([]Car, error) {
	var cars []Car
	err := db.C(COLLECTION).Find(bson.M{}).All(&cars)
	return cars, err
}

// Find a car by its id
func (c *CarsDAO) FindById(id string) (Car, error) {
	var car Car
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&car)
	return car, err
}

// Insert a car into database
func (c *CarsDAO) Insert(car Car) error {
	err := db.C(COLLECTION).Insert(&car)
	return err
}

// Delete an existing car
func (c *CarsDAO) Delete(car Car) error {
	err := db.C(COLLECTION).Remove(&car)
	return err
}

// Update an existing car
func (c *CarsDAO) Update(car Car) error {
	err := db.C(COLLECTION).UpdateId(car.ID, &car)
	return err
}
