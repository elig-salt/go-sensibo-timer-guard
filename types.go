package main

import "time"

type SensiboPodsResponse struct {
	Status string       `json:"status"`
	Result []SensiboPod `json:"result"`
}

type SensiboPod struct {
	Id   string `json:"id"`
	Room struct {
		UID  string `json:"uid"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	} `json:"room"`
	AcState struct {
		Timestamp struct {
			Time       time.Time `json:"time"`
			SecondsAgo int       `json:"secondsAgo"`
		} `json:"timestamp"`
		On                bool   `json:"on"`
		Mode              string `json:"mode"`
		TargetTemperature int    `json:"targetTemperature"`
		TemperatureUnit   string `json:"temperatureUnit"`
		FanLevel          string `json:"fanLevel"`
		Swing             string `json:"swing"`
	} `json:"acState"`
	Timer struct {
		ID        string `json:"id"`
		IsEnabled bool   `json:"isEnabled"`
		AcState   struct {
			On                bool   `json:"on"`
			TargetTemperature int    `json:"targetTemperature"`
			TemperatureUnit   string `json:"temperatureUnit"`
			Mode              string `json:"mode"`
			FanLevel          string `json:"fanLevel"`
			Swing             string `json:"swing"`
		} `json:"acState"`
		CausedBy struct {
			Username  string `json:"username"`
			Email     string `json:"email"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		} `json:"causedBy"`
		CreateTime               string `json:"createTime"`
		CreateTimeSecondsAgo     int    `json:"createTimeSecondsAgo"`
		TargetTime               string `json:"targetTime"`
		TargetTimeSecondsFromNow int    `json:"targetTimeSecondsFromNow"`
	} `json:"timer"`
}

type AcState struct {
	On                bool    `json:"on"`
	Mode              *string `json:"mode,omitempty"`
	FanLevel          *string `json:"fanLevel,omitempty"`
	TargetTemperature *int    `json:"targetTemperature,omitempty"`
	TemperatureUnit   *string `json:"temperatureUnit,omitempty"`
	Swing             *string `json:"swing,omitempty"`
}
type SensiboScheduleTimerRequest struct {
	MinutesFromNow int     `json:"minutesFromNow"`
	AcState        AcState `json:"acState"`
}

type SensiboOperationResponse struct {
	Status string `json:"status"`
}
