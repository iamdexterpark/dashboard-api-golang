package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type Traffic []struct {
	Application string      `json:"application"`
	Destination interface{} `json:"destination"`
	Protocol    string      `json:"protocol"`
	Port        int         `json:"port"`
	Sent        int         `json:"sent"`
	Recv        int         `json:"recv"`
	NumClients  int         `json:"numClients"`
	ActiveTime  int         `json:"activeTime"`
	Flows       int         `json:"flows"`
}

// Return the traffic analysis data for this network
func GetTraffic(networkId, t0, timespan, deviceType string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/traffic", api.BaseUrl(), networkId)
	var datamodel = Traffic{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":         t0,
		"timespan":   timespan,
		"deviceType": deviceType}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}
