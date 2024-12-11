package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func CREATE_LOGS(dirPath string) {
	// get list of directories and find logs folder if it exists. if not, create it
	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Directory doesn't exist, create it
		err := os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating logs directory:", err)
			return
		}
	} else if err != nil {
		// Some error occurred while checking the directory
		fmt.Println("Error:", err)
		return
	}
}

// log start of file
func LOG_START() {
	log.SetPrefix("Start of File: ")
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("----START----")
}

// log weekday info
func LOG_WEEKDAY(isWeekday bool, day time.Weekday) {
	log.SetPrefix("Weekday: ")
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(day, ",", isWeekday)
}

// log time message
func LOG_TIME_MESSAGE(timeMessage string) {
	log.SetPrefix("Message: ")
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(timeMessage)
}

// log message send status
func LOG_MESSAGE_STATS(chatID int64, sendStatus int) {
	log.SetPrefix("Send Status: ")
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(chatID, ",", sendStatus)
}

// log when new request is started
func LOG_REQUEST_RUN(){
	log.SetPrefix("New Request Started: ")
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println("----REQUEST----")
}