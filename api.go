package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

const API_URL_BASE = "https://home.sensibo.com/api"
const API_URL_SCHEDULE_TIMER = "v1/pods"

var API_FIELDS_GET_PODS = []string{"room", "timer", "acState", "id"}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func apiUrlGetPods(apiKey string) string {
	path := "v2/users/me/pods"

	fieldsJoined := strings.Join(API_FIELDS_GET_PODS, ",")
	url := fmt.Sprintf("%s/%s?apiKey=%s&fields=%s", API_URL_BASE, path, apiKey, fieldsJoined)
	return url
}

func apiUrlScheduleTimer(apiKey, deviceId string) string {
	path := fmt.Sprintf("v1/pods/%s/timer", deviceId)
	url := fmt.Sprintf("%s/%s?apiKey=%s", API_URL_BASE, path, apiKey)
	return url
}

func sendGedPodsRequest(apiKey string) (*SensiboPodsResponse, error) {
	sensiboResponse := &SensiboPodsResponse{}
	url := apiUrlGetPods(apiKey)
	if err := getJson(httpClient, url, &sensiboResponse); err != nil {
		return nil, fmt.Errorf("Error getting Sensibo pods: %v", err)
	}
	return sensiboResponse, nil
}

func sendScheduleTimerForPodRequest(apiKey string, pod SensiboPod, minutesFromNow int) error {
	url := apiUrlScheduleTimer(apiKey, pod.Id)

	request := &SensiboScheduleTimerRequest{
		MinutesFromNow: minutesFromNow,
		AcState: AcState{
			On: false,
		},
	}

	response := &SensiboOperationResponse{}
	if err := sendJson(httpClient, http.MethodPut, url, &request, &response); err != nil {
		return fmt.Errorf("Error setting timer state: %v", err)
	}
	if response.Status != "success" {
		return fmt.Errorf("Bad response sent: %v", response.Status)
	}
	return nil
}
