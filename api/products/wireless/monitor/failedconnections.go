package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type FailedConnections []struct {
	SsidNumber  int       `json:"ssidNumber"`
	Vlan        int       `json:"vlan"`
	ClientMac   string    `json:"clientMac"`
	Serial      string    `json:"serial,omitempty"`
	FailureStep string    `json:"failureStep"`
	Type        string    `json:"type"`
	Ts          time.Time `json:"ts"`
	NodeID      string    `json:"nodeId,omitempty"`
}

func GetFailedConnections(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/failedConnections",
		 networkId)
	var datamodel = FailedConnections{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"t1":       t1,
		"timespan": timespan,
		"band":     band,
		"ssid":     ssid,
		"vlan":     vlan,
		"apTag":    apTag}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
