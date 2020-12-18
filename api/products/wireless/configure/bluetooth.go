package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type BluetoothNetworkSettings struct {
	ScanningEnabled          bool   `json:"scanningEnabled"`
	AdvertisingEnabled       bool   `json:"advertisingEnabled"`
	UUID                     string `json:"uuid"`
	MajorMinorAssignmentMode string `json:"majorMinorAssignmentMode"`
	Major                    int    `json:"major"`
}

type BluetoothDeviceSettings struct {
	UUID  string `json:"uuid"`
	Major int    `json:"major"`
	Minor int    `json:"minor"`
}

func GetBluetoothNetworkSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/bluetooth/settings",
		 networkId)
	var datamodel = BluetoothNetworkSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutBluetoothNetworkSettings(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/bluetooth/settings",
		 networkId)
	var datamodel = BluetoothNetworkSettings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetBluetoothDeviceSettings(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/wireless/bluetooth/settings",
		 serial)
	var datamodel = BluetoothDeviceSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutBluetoothDeviceSettings(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/wireless/bluetooth/settings",
		 serial)
	var datamodel = BluetoothDeviceSettings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}