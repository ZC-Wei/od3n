package rankScore

import (
	"math"
	"od3n/dispatchNode/objects"

	"github.com/kellydunn/golang-geo"
)

//RankScore calculation
func RankScore(delivery objects.Delivery, carrier objects.Carrier, trafficStatus objects.TrafficStatus) float64 {
	/*
		RankScore = ETArrvial + ETAvailable
		Expected Time Arrvial and Expected Time Available
		ETArrvial is carrier pending duration for next delivery
		ETAvailable is carrier travel duration for new delivery departure
	*/
	var ETArrvial, ETAvailable float64
	var distance, deliveryRemain float64

	//Calculation real speed by vechicle theory speed and city traffic congestion degree
	speed := carrier.Vechicle.Speed * math.Pow(trafficStatus.CongestionDegree, carrier.Vechicle.VechicleType)

	carrierPoint := geo.NewPoint(carrier.Location.Latitude, carrier.Location.Longitude)
	departurePoint := geo.NewPoint(delivery.Departure.Latitude, delivery.Departure.Longitude)
	distance = carrierPoint.GreatCircleDistance(departurePoint)
	ETArrvial = distance / speed

	if carrier.CurrentJob.ID != "NULL" {
		lastDestinationPoint := geo.NewPoint(carrier.CurrentJob.Destination.Latitude, carrier.CurrentJob.Destination.Longitude)
		deliveryRemain = carrierPoint.GreatCircleDistance(lastDestinationPoint)
	} else {
		deliveryRemain = 0
	}
	ETAvailable = deliveryRemain / speed

	return ETArrvial + ETAvailable
}
