package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type ChannelUtilizations []struct {
	ChannelUtilization
}

type ChannelUtilization struct {
	Serial string `json:"serial"`
	Model  string `json:"model"`
	Tags   string `json:"tags"`
	Wifi0  []struct {
		StartTime           time.Time `json:"startTime"`
		EndTime             time.Time `json:"endTime"`
		UtilizationTotal    float64   `json:"utilizationTotal"`
		Utilization80211    int       `json:"utilization80211"`
		UtilizationNon80211 float64   `json:"utilizationNon80211"`
	} `json:"wifi0"`
	Wifi1 []struct {
		StartTime           time.Time `json:"startTime"`
		EndTime             time.Time `json:"endTime"`
		UtilizationTotal    float64   `json:"utilizationTotal"`
		Utilization80211    int       `json:"utilization80211"`
		UtilizationNon80211 float64   `json:"utilizationNon80211"`
	} `json:"wifi1"`
}

func GetChannelUtilization(networkId, t0, t1, timespan, resolution, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/networkHealth/channelUtilization",  networkId)
	var datamodel = ChannelUtilizations{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":            t0,
		"t1":            t1,
		"timespan":      timespan,
		"resolution":    resolution,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
