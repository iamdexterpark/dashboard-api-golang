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

func PostBlinkLeds(serial, duration, period, duty string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/blinkLeds", api.BaseUrl(), serial)
	var datamodel BlinkLeds

	// Parameters for Request URL
	var parameters = map[string]string{
		"duration": duration,
		"period": period,
		"duty": duty}

	sessions, err := api.Sessions(baseurl, "POST", nil, parameters, datamodel)
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
