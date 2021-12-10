package main

import (
	"log"
	"os"
	"strconv"
)

func init() {
	log.SetFlags(0)
}

const TIMER_SCHEDULE_MINS_DEFAULT = 60

var mins int = TIMER_SCHEDULE_MINS_DEFAULT

func main() {
	apiKey := os.Getenv("SENSIBO_API_KEY")
	if apiKey == "" {
		log.Fatal("API key missing. Please provide SENSIBO_API_KEY env var")
	}

	minsEnv := os.Getenv("SENSIBO_TIMER_SCHEDULE_MINS")
	if minsEnv != "" {
		minsParsed, err := strconv.ParseInt(minsEnv, 10, 0)
		if err != nil {
			log.Fatal("Could not parse SENSIBO_TIMER_SCHEDULE_MINS to an integer")
		}
		mins = int(minsParsed)
	}

	sensiboResponse, err := sendGedPodsRequest(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	roomsToScheduleTimer := getUnscheduledRooms(sensiboResponse)
	scheduleTimerForRooms(apiKey, roomsToScheduleTimer, mins)
}
