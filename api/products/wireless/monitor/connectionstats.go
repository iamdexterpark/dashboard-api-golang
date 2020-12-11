package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type ConnectionStats struct {
	Mac             string `json:"mac"`
	ConnectionStats struct {
		Assoc   int `json:"assoc"`
		Auth    int `json:"auth"`
		Dhcp    int `json:"dhcp"`
		DNS     int `json:"dns"`
		Success int `json:"success"`
	} `json:"connectionStats"`
}

type ConnectionStat struct {
	Assoc   int `json:"assoc"`
	Auth    int `json:"auth"`
	Dhcp    int `json:"dhcp"`
	DNS     int `json:"dns"`
	Success int `json:"success"`
}


func GetConnectionStats(devices, serial, t0, t1, timespan,
	band, ssid, vlan, apTag string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/wireless/clients/%s/connectionStats",
		api.BaseUrl(), devices, serial)
	var datamodel = ConnectionStats{}

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

func GetConnectionStat(networkId, clientId, t0, t1, timespan,
	band, ssid, vlan, apTag string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/clients/%s/connectionStats",
		api.BaseUrl(), networkId, clientId)
	var datamodel = ConnectionStat{}

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