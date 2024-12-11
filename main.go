package main

import (
	"fmt"
	"strings"
	"time"
	"os"
	"log"
)


func main() {
    // set up logs
	dirPath := "./logs"
	CREATE_LOGS(dirPath)
	// open log file at start of run
	curTime := time.Now()
	formattedTIme := curTime.Format("01-02-2006_15_04_05")
	os.Chdir(dirPath)
	logFileName := "Trains_Log_" + formattedTIme + ".log"
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatal("Failed to open log file:", err)
    }
    defer file.Close()
	log.SetOutput(file)
	LOG_START()

	// Define the URL of the web page you want to fetch
    ssMorn := "Overbrook"
	seMorn := "30th Street Station"
	ssEve := "30th Street Station"
	seEve := "Overbrook"
	stationStartMorn := strings.ReplaceAll(ssMorn, " ", "%20")
	stationEndMorn := strings.ReplaceAll(seMorn, " ", "%20")
	urlMorn := fmt.Sprintf("https://www3.septa.org/api/NextToArrive/index.php?req1=%v&req2=%v", stationStartMorn, stationEndMorn)
	stationStartEve := strings.ReplaceAll(ssEve, " ", "%20")
	stationEndEve := strings.ReplaceAll(seEve, " ", "%20")
	urlEve := fmt.Sprintf("https://www3.septa.org/api/NextToArrive/index.php?req1=%v&req2=%v", stationStartEve, stationEndEve)

	// deal with telegram stuff
	apiToken := TelegramAPIToken // using constant from config.go
	chatID := int64(UserChatID)
	
	// set start and end times
	startTimeMorn := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 07, 05, 0, 0, time.Local)
	endTimeMorn := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 07, 55, 0, 0, time.Local)
	startTimeEve := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 15, 05, 0, 0, time.Local)
	endTimeEve := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 16, 05, 0, 0, time.Local)
	
	// for testing:
	// startTimeMorn := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 00, 00, 0, 0, time.Local)
	// endTimeMorn := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 00, 0, 0, time.Local)
	// startTimeEve := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 01, 0, 0, time.Local)
	// endTimeEve := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 0, 0, time.Local)
	
	// Create a timer that fires every X minutes
    ticker := time.NewTicker(5 * time.Minute)
	// ticker := time.NewTicker(20 * time.Second)
    
	for {
		<-ticker.C 
		currentTime := time.Now()
		LOG_REQUEST_RUN()

		// check for weekday or weekend
		var isWeekday bool
		if currentTime.Weekday() != time.Saturday && currentTime.Weekday() != time.Sunday {
			isWeekday = true
			LOG_WEEKDAY(isWeekday, currentTime.Weekday())
		} else {
			isWeekday = false //should be false when not testing
			LOG_WEEKDAY(isWeekday, currentTime.Weekday())
		}

		// do the stuff if it's a weekday
		if isWeekday == true {
			if currentTime.After(startTimeMorn) && currentTime.Before(endTimeMorn) {
				timeMessage := GETTRAINS(urlMorn, ssMorn, seMorn)
				LOG_TIME_MESSAGE(timeMessage)
				fmt.Println(currentTime, timeMessage)
				fmt.Println("FIRST SECTION")
				sendStatus := SEND_MESSAGE(apiToken, chatID, timeMessage)
				LOG_MESSAGE_STATS(chatID, sendStatus)
			} else if currentTime.After(startTimeEve) && currentTime.Before(endTimeEve) {
				timeMessage := GETTRAINS(urlEve, ssEve, seEve)
				LOG_TIME_MESSAGE(timeMessage)
				fmt.Println(currentTime, timeMessage)
				fmt.Println("SECOND SECTION")
				sendStatus := SEND_MESSAGE(apiToken, chatID, timeMessage)
				LOG_MESSAGE_STATS(chatID, sendStatus)
			}
		} else {
			continue
		}
	} 
}