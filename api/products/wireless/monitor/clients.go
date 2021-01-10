package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)



type ConnectivityEvents []struct {
	OccurredAt   string     `json:"occurredAt"`
	DeviceSerial string  `json:"deviceSerial"`
	Band         string     `json:"band"`
	SsidNumber   string     `json:"ssidNumber"`
	Type         string  `json:"type"`
	Subtype      string  `json:"subtype"`
	Severity     string  `json:"severity"`
	DurationMs   float64 `json:"durationMs"`
	Channel      string     `json:"channel"`
	Rssi         string     `json:"rssi"`
	EventData    struct {
		ClientIP string `json:"clientIp"`
	} `json:"eventData"`
}

type ClientLatencyTraffic struct {
	Zero5   string `json:"0.5"`
	One1    string `json:"1.1"`
	Two1    string `json:"2.1"`
	Four1   string `json:"4.1"`
	Eight1  string `json:"8.1"`
	One61   string `json:"16.1"`
	Three21 string `json:"32.1"`
	Six41   string `json:"64.1"`
	One281  string `json:"128.1"`
	Two561  string `json:"256.1"`
	Five121 string `json:"512.1"`
	One0241 string `json:"1024.1"`
	Two0481 string `json:"2048.1"`
}
type ClientLatencyHistory []struct {
	T0                    string `json:"t0"`
	T1                    string `json:"t1"`
	LatencyBinsByCategory struct {
		BackgroundTraffic struct {
			ClientLatencyTraffic
		} `json:"backgroundTraffic"`
		BestEffortTraffic struct {
			ClientLatencyTraffic
		} `json:"bestEffortTraffic"`
		VideoTraffic struct {
			ClientLatencyTraffic
		} `json:"videoTraffic"`
		VoiceTraffic struct {
			ClientLatencyTraffic
		} `json:"voiceTraffic"`
	} `json:"latencyBinsByCategory"`
}

type AggregatedLatencies struct {
	AggregatedLatency
}

type AggregatedLatency struct {
	Mac          string `json:"mac"`
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

func GetConnectivityEvents(networkId, clientId, perPage, startingAfter,
	endingBefore, t0, t1, timespan, types, includedSeverities, band, ssidNumber,
	deviceSerial string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/clients/%s/connectivityEvents",
		 networkId, clientId)
	var datamodel = ConnectivityEvents{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":            perPage,
		"startingAfter":      startingAfter,
		"endingBefore":       endingBefore,
		"t0":                 t0,
		"t1":                 t1,
		"timespan":           timespan,
		"types":              types,
		"includedSeverities": includedSeverities,
		"band":               band,
		"ssidNumber":         ssidNumber,
		"deviceSerial":       deviceSerial}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetClientLatencyHistory(networkId, clientId, t0, t1, timespan,
	resolution string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/clients/%s/latencyHistory",
		 networkId, clientId)
	var datamodel = ClientLatencyHistory{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":         t0,
		"t1":         t1,
		"timespan":   timespan,
		"resolution": resolution}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetAggregatedLatencies(networkId, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/clients/latencyStats",
		 networkId)
	var datamodel = AggregatedLatencies{}

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

func GetAggregatedLatency(networkId, clientId, t0, t1, timespan,
	band, ssid, vlan, apTag, fields string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/clients/%s/latencyStats",
		 networkId, clientId)
	var datamodel = AggregatedLatency{}

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
