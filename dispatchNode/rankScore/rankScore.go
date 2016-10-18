package rankScore

import (
	"OD3N/dispatchNode/objects"
	"math"

	"github.com/kellydunn/golang-geo"
)

//RankScore calculation
func RankScore(delivery objects.Delivery, carrier objects.Carrier, trafficStatus objects.TrafficStatus) float64 {
	var ETArrvial, ETAvailable float64
	var distance, deliveryRemain float64

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
