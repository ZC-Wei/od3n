package dispatch

import (
	"OD3N/dispatchNode/objects"
	"OD3N/dispatchNode/rankScore"
	"time"

	"github.com/satori/go.uuid"
)

//Dispatch method return an assignment
func Dispatch(delivery objects.Delivery, carrierPool []objects.Carrier, trafficStatus objects.TrafficStatus) objects.Assignment {
	rankMap := make(map[string]float64)
	for _, value1 := range carrierPool {
		rankMap[value1.ID] = rankScore.RankScore(delivery, value1, trafficStatus)
	}
	var smallestScoreCarrier string
	var smallestScore float64 = 99999999
	for key, value2 := range rankMap {
		if value2 < smallestScore {
			smallestScoreCarrier = key
			smallestScore = value2
		}
	}
	assignmentID := uuid.NewV4()
	assignment := objects.Assignment{ID: assignmentID.String(), DeliveryID: delivery.ID, CarrierID: smallestScoreCarrier, TimeStamp: time.Now()}
	return assignment
}
