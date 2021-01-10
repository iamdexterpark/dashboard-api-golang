package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type RadioSettings struct {
	Serial             string `json:"serial"`
	RfProfileID        string `json:"rfProfileId"`
	TwoFourGhzSettings struct {
		Channel     string `json:"channel"`
		TargetPower string `json:"targetPower"`
	} `json:"twoFourGhzSettings"`
	FiveGhzSettings struct {
		Channel      string `json:"channel"`
		ChannelWidth string `json:"channelWidth"`
		TargetPower  string `json:"targetPower"`
	} `json:"fiveGhzSettings"`
}

func GetRadioSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/wireless/radio/settings",
		 serial)
	var datamodel = RadioSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutRadioSettings(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/wireless/radio/settings",
		 serial)
	var datamodel = RadioSettings{}
	paylaod := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", paylaod, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}