package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type ChannelUtilizationHistory []struct {
	StartTs             time.Time `json:"startTs"`
	EndTs               time.Time `json:"endTs"`
	UtilizationTotal    float64   `json:"utilizationTotal"`
	Utilization80211    float64   `json:"utilization80211"`
	UtilizationNon80211 float64   `json:"utilizationNon80211"`
}

func GetChannelUtilizationHistory(networkId, t0, t1, timespan,
	resolution, autoResolution, clientId, deviceSerial, apTag,
	band string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/channelUtilizationHistory",  networkId)
	var datamodel = ChannelUtilizationHistory{}

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
		"band":           band}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
