package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "time"
)

type NextToArrive struct {
    OrigTrain   string `json:"orig_train"`
    OrigLine    string `json:"orig_line"`
    OrigDepTime string `json:"orig_departure_time"`
    OrigArrTime string `json:"orig_arrival_time"`
    OrigDelay   string `json:"orig_delay"`
    IsDirect    string `json:"isdirect"`
}

// get information for the next trains to arrive
func GETTRAINS(url string, stationStart string, stationEnd string) (timeMessage string){
	// pull the webpage based on the url
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer response.Body.Close()

	// PARSE JSON DATA
	var availTrains []NextToArrive
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&availTrains); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// printing info for each train
	// curTime := time.Now()
	// timeResp = curTime.Format("020106150405")
	// train := availTrains[0]
	train1 := availTrains[0]
	train2 := availTrains[1]
	if train1.OrigDelay != "On time" {
		train1.OrigDelay = fmt.Sprintf("+%s", train1.OrigDelay)
	} 
	if train2.OrigDelay != "On time" {
		train2.OrigDelay = fmt.Sprintf("+%s", train2.OrigDelay)
	} 
	// create message for if train is running late and if train is on time
	// if train.OrigDelay != "On time" {
	// 	timeMessage := fmt.Sprintf("Next train is running %s late. Now leaving at %s.", train.OrigDelay, train.OrigDepTime)
	// 	// fmt.Println(timeMessage)
	// 	return availTrains, timeMessage, timeResp
	// } else {
	// 	timeMessage := fmt.Sprintf("Next train is on time, leaving at %s.", train.OrigDepTime)
	// 	// fmt.Println(timeMessage)
	// 	return availTrains, timeMessage, timeResp
	// }
	timeMessage = fmt.Sprintf("Next 2 trains from %s to %s are at %s (%s) and %s (%s).", stationStart, stationEnd, train1.OrigDepTime, train1.OrigDelay, train2.OrigDepTime, train2.OrigDelay)
	
	return timeMessage
	
}

// send message
func SENDMESSAGE(timeMessage string) {
	
}
