package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type FirmwareUpgrades struct {
	UpgradeWindow struct {
		HourOfDay string `json:"hourOfDay"`
		DayOfWeek string `json:"dayOfWeek"`
	} `json:"upgradeWindow"`
}


func GetFirmwareUpgrades(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/firmwareUpgrades",  networkId)
	var datamodel = FirmwareUpgrades{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutFirmwareUpgrades(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/firmwareUpgrades",  networkId)
	var datamodel = FirmwareUpgrades{}
	sessions, err := api.Sessions(baseurl, "PUT", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}