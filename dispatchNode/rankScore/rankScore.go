package rankScore

import (
	"math"
	"od3n/dispatchNode/objects"

	"github.com/kellydunn/golang-geo"
)

//RankScore calculation
func RankScore(delivery objects.Delivery, courier objects.Courier, trafficStatus objects.TrafficStatus) float64 {
	/*
		RankScore = ETArrvial + ETAvailable
		Expected Time Arrvial and Expected Time Available
		ETArrvial is carrier pending duration for next delivery
		ETAvailable is carrier travel duration for new delivery departure
	*/
	var ETArrvial, ETAvailable float64
	var distance, deliveryRemain float64

	//Calculation real speed by vechicle theory speed and city traffic congestion degree
	speed := courier.Vechicle.Speed * math.Pow(trafficStatus.CongestionDegree, courier.Vechicle.VechicleType)

	courierPoint := geo.NewPoint(courier.Location.Latitude, courier.Location.Longitude)
	departurePoint := geo.NewPoint(delivery.Departure.Latitude, delivery.Departure.Longitude)
	distance = courierPoint.GreatCircleDistance(departurePoint)
	ETArrvial = distance / speed

	if courier.CurrentJob.ID != "NULL" {
		lastDestinationPoint := geo.NewPoint(courier.CurrentJob.Destination.Latitude, courier.CurrentJob.Destination.Longitude)
		deliveryRemain = courierPoint.GreatCircleDistance(lastDestinationPoint)
	} else {
		deliveryRemain = 0
	}
	ETAvailable = deliveryRemain / speed

	return ETArrvial + ETAvailable
}
