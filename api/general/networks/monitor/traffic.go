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

// Return the client's network traffic data over time. Usage data is in kilobytes.
// This endpoint requires detailed traffic analysis to be enabled on the Network-wide > General
// page. Clients can be identified by a client key or either the MAC or IP depending on whether
// the network uses Track-by-IP.
func GetTraffic(networkId, t0, timespan, deviceType string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/traffic",  networkId)
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
