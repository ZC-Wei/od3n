package main

import (
	"OD3N/dispatchNode/dispatch"
	"OD3N/dispatchNode/objects"
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

//MongoDB container on AWS EC2 Docker
const serverURL = "ec2-52-49-174-40.eu-west-1.compute.amazonaws.com:27017"

func main() {
	//Initalize a delivery request with fake data
	departure := objects.Point{Latitude: 53.320060, Longitude: -6.279071}
	destination := objects.Point{Latitude: 53.305786, Longitude: -6.291633}
	delivery := objects.Delivery{ID: "d000001", Departure: departure, Destination: destination, TimeStamp: time.Now()}
	//Initalize a traffic parameter
	traffic := objects.TrafficStatus{City: "Dublin", CongestionDegree: 1}

	//Connect to MongoDB
	session, err := mgo.Dial(serverURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//Get a collection of carriers
	collection := session.DB("OD3N").C("carriers")
	var carriers []objects.Carrier
	err = collection.Find(nil).All(&carriers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(carriers)

	//Dispatch delivery to best score carrier
	assignment := dispatch.Dispatch(delivery, carriers, traffic)

	fmt.Println(assignment)
}

//iris go web framework
/*
import "github.com/kataras/iris"

func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello, %s", "World!")
	})

	iris.Get("/myjson", func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{
			"Name":    "Iris",
			"Release": "13 March 2016",
			"Stars":   5525,
		})
	})

	iris.Listen(":8080")
}
*/