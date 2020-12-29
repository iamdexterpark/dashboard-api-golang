package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type FirmwareUpgrades struct {
	UpgradeWindow struct {
		HourOfDay string `json:"hourOfDay"`
		DayOfWeek string `json:"dayOfWeek"`
	} `json:"upgradeWindow"`
}

// Get current maintenance window for a network
func GetFirmwareUpgrades(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/firmwareUpgrades",  networkId)
	var datamodel = FirmwareUpgrades{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update current maintenance window for a network
func PutFirmwareUpgrades(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/firmwareUpgrades",  networkId)
	var datamodel = FirmwareUpgrades{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}