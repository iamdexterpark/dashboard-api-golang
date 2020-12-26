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
		UsingStaticIP    bool     `json:"usingStaticIp"`
		StaticIP         string   `json:"staticIp"`
		StaticSubnetMask string   `json:"staticSubnetMask"`
		StaticGatewayIP  string   `json:"staticGatewayIp"`
		StaticDNS        []string `json:"staticDns"`
		Vlan             int      `json:"vlan"`
	} `json:"wan1"`
	Wan2 struct {
		WanEnabled    string `json:"wanEnabled"`
		UsingStaticIP bool   `json:"usingStaticIp"`
		Vlan          int    `json:"vlan"`
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
		Major int    `json:"major"`
		Minor int    `json:"minor"`
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

/*
Update the management interface settings for a device
 #### Body Parameters
wan1 | object | WAN 1 settings

wan2 | object | WAN 2 settings (only for MX devices)
*/
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

/*
Update the attributes of a device
 #### Body Parameters
name | string | The name of a device

tags | array | The list of tags of a device

lat | number | The latitude of a device

lng | number | The longitude of a device

address | string | The address of a device

notes | string | The notes for the device. String. Limited to 255 characters.

moveMapMarker | boolean | Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.

switchProfileId | string | The ID of a switch profile to bind to the device (for available switch profiles, see the 'Switch Profiles' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch profile, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.

floorPlanId | string | The floor plan to associate to this device. null disassociates the device from the floorplan.
*/
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