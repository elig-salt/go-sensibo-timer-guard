package main

import (
	"fmt"
	"sync"
)

func getUnscheduledRooms(sensiboResponse *SensiboPodsResponse) []SensiboPod {
	roomIds := []SensiboPod{}
	for i := 0; i < len(sensiboResponse.Result); i++ {
		pod := sensiboResponse.Result[i]
		if pod.AcState.On && !pod.Timer.IsEnabled {
			fmt.Printf("Room %s is without a timer !\n", pod.Room.Name)
			roomIds = append(roomIds, pod)
		}
	}
	return roomIds
}

func scheduleTimerForRooms(apiKey string, pods []SensiboPod, mins int) error {
	wg := &sync.WaitGroup{}
	wg.Add(len(pods))

	for i := 0; i < len(pods); i++ {
		go func(index int) {
			fmt.Printf("Setting Timer for pod: %s\n", pods[index].Room.Name)
			if err := sendScheduleTimerForPodRequest(apiKey, pods[index], mins); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	return nil
}
