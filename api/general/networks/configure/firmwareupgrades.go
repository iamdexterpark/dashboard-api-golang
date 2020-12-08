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
	baseurl := fmt.Sprintf("%s/networks/%s/firmwareUpgrades", api.BaseUrl(), networkId)
	var datamodel = FirmwareUpgrades{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutFirmwareUpgrades(networkId, upgradeWindow string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/firmwareUpgrades", api.BaseUrl(), networkId)
	var datamodel = FirmwareUpgrades{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"upgradeWindow": upgradeWindow}
	sessions, err := api.Sessions(baseurl, "PUT", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}