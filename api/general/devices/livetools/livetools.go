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

/*
Blink the LEDs on a device
 #### Body Parameters
duration | integer | The duration in seconds. Must be between 5 and 120. Default is 20 seconds
period | integer | The period in milliseconds. Must be between 100 and 1000. Default is 160 milliseconds
duty | integer | The duty cycle as the percent active. Must be between 10 and 90. Default is 50.
 */
func PostBlinkLeds(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/blinkLeds",  serial)
	var datamodel BlinkLeds

	sessions, err := api.Sessions(baseurl, "POST", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PostReboot(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/blinkLeds",  serial)
	var datamodel Reboot
	data := Reboot{true}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
