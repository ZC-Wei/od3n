package objects

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Point object
type Point struct {
	Latitude  float64 `bson:"Latitude"`
	Longitude float64 `bson:"Longitude"`
}

//Delivery object
type Delivery struct {
	ID          string    `bson:"ID"`
	Departure   Point     `bson:"Departure"`
	Destination Point     `bson:"Destination"`
	TimeStamp   time.Time `bson:"TimeStamp"`
}

//Vechicle object
type Vechicle struct {
	//VechicleType Bike = 1 Car = 1.2 Express = 1.5
	VechicleType float64 `bson:"VechicleType"`
	Speed        float64 `bson:"Speed"`
}

//Courier object
type Courier struct {
	_id        bson.ObjectId `bson:"_id,omitempty"`
	ID         string        `bson:"ID"`
	Location   Point         `bson:"Location"`
	Vechicle   Vechicle      `bson:"Vechicle"`
	CurrentJob Delivery      `bson:"Delivery"`
}

//TrafficStatus object
type TrafficStatus struct {
	City             string  `bson:"City"`
	CongestionDegree float64 `bson:"CongestionDegree"`
}

//Assignment object
type Assignment struct {
	ID         string    `bson:"ID"`
	DeliveryID string    `bson:"DeliveryID"`
	CourierID  string    `bson:"CourierID"`
	TimeStamp  time.Time `bson:"TimeStamp"`
}
