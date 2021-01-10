package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type ConnectionStats struct {
	Mac             string `json:"mac"`
	ConnectionStats struct {
		Assoc   string `json:"assoc"`
		Auth    string `json:"auth"`
		Dhcp    string `json:"dhcp"`
		DNS     string `json:"dns"`
		Success string `json:"success"`
	} `json:"connectionStats"`
}

type ConnectionStat struct {
	Assoc   string `json:"assoc"`
	Auth    string `json:"auth"`
	Dhcp    string `json:"dhcp"`
	DNS     string `json:"dns"`
	Success string `json:"success"`
}


func GetConnectionStats(devices, serial, t0, t1, timespan,
	band, ssid, vlan, apTag string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/wireless/clients/%s/connectionStats",
		 devices, serial)
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
	baseurl := fmt.Sprintf("/networks/%s/wireless/clients/%s/connectionStats",
		 networkId, clientId)
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