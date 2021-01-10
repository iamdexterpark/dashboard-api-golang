package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type AirMarshal[]struct {
	Ssid   string `json:"ssid"`
	Bssids []struct {
		Bssid      string `json:"bssid"`
		Contained  string `json:"contained"`
		DetectedBy []struct {
			Device string `json:"device"`
			Rssi   string `json:"rssi"`
		} `json:"detectedBy"`
	} `json:"bssids"`
	Channels      []string `json:"channels"`
	FirstSeen     string      `json:"firstSeen"`
	LastSeen      string      `json:"lastSeen"`
	WiredMacs     []string `json:"wiredMacs"`
	WiredVlans    []string `json:"wiredVlans"`
	WiredLastSeen string      `json:"wiredLastSeen"`
}

func GetAirMarshal(serial, t0, timespan string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/airMarshal",
		 serial)
	var datamodel = AirMarshal{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"timespan": timespan}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
