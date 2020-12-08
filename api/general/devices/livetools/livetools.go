package livetools

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type BlinkLeds struct {
	Duration int `json:"duration"`
	Period   int `json:"period"`
	Duty     int `json:"duty"`
}

type Reboot struct {
	Success bool `json:"success"`
}

func PostBlinkLeds(serial string, duration, period, duty int) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/blinkLeds", api.BaseUrl(), serial)
	var datamodel BlinkLeds
	data := BlinkLeds{
		Duration: duration,
		Period:   period,
		Duty:     duty,
	}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PostReboot(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/blinkLeds", api.BaseUrl(), serial)
	var datamodel Reboot
	data := Reboot{true}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}