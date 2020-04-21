package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Car struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Model       string        `bson:"model" json:"model"`
	Color       string        `bson:"color" json:"color"`
	MakeYear    string        `bson:"make_year" json:"make_year"`
	FuelType    string        `bson:"fuel_type" json:"fuel_type"`
}
