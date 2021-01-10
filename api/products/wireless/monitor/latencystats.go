package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type APLatencyStats struct {
	Serial       string `json:"serial"`
	LatencyStats struct {
		BackgroundTraffic struct {
			RawDistribution struct {
				Num0    string `json:"0"`
				Num1    string `json:"1"`
				Num2    string `json:"2"`
				Num4    string `json:"4"`
				Num8    string `json:"8"`
				Num16   string `json:"16"`
				Num32   string `json:"32"`
				Num64   string `json:"64"`
				Num128  string `json:"128"`
				Num256  string `json:"256"`
				Num512  string `json:"512"`
				Num1024 string `json:"1024"`
				Num2048 string `json:"2048"`
			} `json:"rawDistribution"`
			Avg float64 `json:"avg"`
		} `json:"backgroundTraffic"`
		BestEffortTraffic string `json:"bestEffortTraffic"`
		VideoTraffic      string `json:"videoTraffic"`
		VoiceTraffic      string `json:"voiceTraffic"`
	} `json:"latencyStats"`
}

type APLatencyStat struct {
	BackgroundTraffic struct {
		RawDistribution struct {
			Num0    string `json:"0"`
			Num1    string `json:"1"`
			Num2    string `json:"2"`
			Num4    string `json:"4"`
			Num8    string `json:"8"`
			Num16   string `json:"16"`
			Num32   string `json:"32"`
			Num64   string `json:"64"`
			Num128  string `json:"128"`
			Num256  string `json:"256"`
			Num512  string `json:"512"`
			Num1024 string `json:"1024"`
			Num2048 string `json:"2048"`
		} `json:"rawDistribution"`
		Avg float64 `json:"avg"`
	} `json:"backgroundTraffic"`
	BestEffortTraffic string `json:"bestEffortTraffic"`
	VideoTraffic      string `json:"videoTraffic"`
	VoiceTraffic      string `json:"voiceTraffic"`
}

func GetAPLatencyStats(serial, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/wireless/clients/latencyStats",
		 serial)
	var datamodel = APLatencyStats{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"t1":       t1,
		"timespan": timespan,
		"band":     band,
		"ssid":     ssid,
		"vlan":     vlan,
		"apTag":    apTag,
		"fields":   fields}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetAPLatencyStat(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/clients/latencyStats",
		 networkId)
	var datamodel = APLatencyStat{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"t1":       t1,
		"timespan": timespan,
		"band":     band,
		"ssid":     ssid,
		"vlan":     vlan,
		"apTag":    apTag,
		"fields":   fields}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
