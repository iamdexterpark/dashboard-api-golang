package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type Status struct {
	BasicServiceSets []struct {
		SsidName     string `json:"ssidName"`
		SsidNumber   string `json:"ssidNumber"`
		Enabled      string `json:"enabled"`
		Band         string `json:"band"`
		Bssid        string `json:"bssid"`
		Channel      string `json:"channel"`
		ChannelWidth string `json:"channelWidth"`
		Power        string `json:"power"`
		Visible      string `json:"visible"`
		Broadcasting string `json:"broadcasting"`
	} `json:"basicServiceSets"`
}

func GetStatus(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/wireless/status",
		 serial)
	var datamodel = Status{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
