package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type ManagementInterface struct {
	DdnsHostnames struct {
		ActiveDdnsHostname string `json:"activeDdnsHostname"`
		DdnsHostnameWan1   string `json:"ddnsHostnameWan1"`
		DdnsHostnameWan2   string `json:"ddnsHostnameWan2"`
	} `json:"ddnsHostnames"`
	Wan1 struct {
		WanEnabled       string   `json:"wanEnabled"`
		UsingStaticIP    string `json:"usingStaticIp"`
		StaticIP         string   `json:"staticIp"`
		StaticSubnetMask string   `json:"staticSubnetMask"`
		StaticGatewayIP  string   `json:"staticGatewayIp"`
		StaticDNS        []string `json:"staticDns"`
		Vlan             string      `json:"vlan"`
	} `json:"wan1"`
	Wan2 struct {
		WanEnabled    string `json:"wanEnabled"`
		UsingStaticIP string `json:"usingStaticIp"`
		Vlan          string `json:"vlan"`
	} `json:"wan2"`
}

type Device struct {
	Name           string   `json:"name"`
	Lat            float64  `json:"lat"`
	Lng            float64  `json:"lng"`
	Serial         string   `json:"serial"`
	Mac            string   `json:"mac"`
	Model          string   `json:"model"`
	Address        string   `json:"address"`
	Notes          string   `json:"notes"`
	LanIP          string   `json:"lanIp"`
	Tags           []string `json:"tags"`
	NetworkID      string   `json:"networkId"`
	BeaconIDParams struct {
		UUID  string `json:"uuid"`
		Major string `json:"major"`
		Minor string `json:"minor"`
	} `json:"beaconIdParams"`
	Firmware    string `json:"firmware"`
	FloorPlanID string `json:"floorPlanId"`
}


// Return the management interface settings for a device
func GetManagementInterface(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/managementInterface",  serial)
	var datamodel = ManagementInterface{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions

}


// Update the management interface settings for a device
func PutManagementInterface(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/managementInterface",  serial)
	var datamodel = ManagementInterface{}
	payload := user_agent.MarshalJSON(data)

	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return a single device
func GetDevice(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s",  serial)
	var datamodel = Device{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions

}

// Update the attributes of a device
func PutDevice(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s",  serial)
	payload := user_agent.MarshalJSON(data)
	var datamodel = Device{}
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}