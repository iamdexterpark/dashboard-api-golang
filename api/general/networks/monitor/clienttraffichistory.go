package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)


type TrafficHistory []struct {
	Ts            time.Time `json:"ts"`
	Application   string    `json:"application"`
	Destination   string    `json:"destination"`
	Protocol      string    `json:"protocol"`
	Port          int       `json:"port"`
	Recv          int       `json:"recv"`
	Sent          int       `json:"sent"`
	NumFlows      int       `json:"numFlows"`
	ActiveSeconds int       `json:"activeSeconds"`
}


type ClientUsageHistory []struct {
	Sent     int       `json:"sent"`
	Received int       `json:"received"`
	Ts       time.Time `json:"ts"`
}




// Return the client's network traffic data over time
func GetClientTrafficHistory(networkId, clientId, perPage,
	startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/trafficHistory",  networkId, clientId)
	var datamodel = TrafficHistory{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":        perPage,
		"startingAfter":  startingAfter,
		"endingBefore":   endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}

// Return the client's daily usage history
func GetClientUsageHistory(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/usageHistory",  networkId, clientId)
	var datamodel = ClientUsageHistory{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


