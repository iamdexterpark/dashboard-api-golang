package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type UsageHistory []struct {
	StartTs      time.Time `json:"startTs"`
	EndTs        time.Time `json:"endTs"`
	TotalKbps    string       `json:"totalKbps"`
	SentKbps     string       `json:"sentKbps"`
	ReceivedKbps string       `json:"receivedKbps"`
}

func GetUsageHistory(networkId, t0, t1, timespan,
	resolution, autoResolution, clientId, deviceSerial, apTag,
	band, ssid string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/usageHistory",  networkId)
	var datamodel = UsageHistory{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":             t0,
		"t1":             t1,
		"timespan":       timespan,
		"resolution":     resolution,
		"autoResolution": autoResolution,
		"clientId":       clientId,
		"deviceSerial":   deviceSerial,
		"apTag":          apTag,
		"band":           band,
		"ssid":           ssid}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
