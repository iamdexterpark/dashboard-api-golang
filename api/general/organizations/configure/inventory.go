package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type InventoryDevices []struct {
	InventoryDevice
}

type InventoryDevice struct {
	Mac                   string `json:"mac"`
	Serial                string `json:"serial"`
	Name                  string `json:"name"`
	Model                 string `json:"model"`
	NetworkID             string `json:"networkId"`
	OrderNumber           string `json:"orderNumber"`
	ClaimedAt             time.Time `json:"claimedAt"`
	LicenseExpirationDate time.Time `json:"licenseExpirationDate"`
}

// Return the device inventory for an organization
func GetInventoryDevices(organizationId, perPage, startingAfter, endingBefore, usedState, search string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/inventoryDevices",
		organizationId)
	var datamodel = InventoryDevices{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":          perPage,
		"startingAfter":    startingAfter,
		"endingBefore":		endingBefore,
		"usedState": 		usedState,
		"search":			search,
	}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return a single device from the inventory of an organization
func GetInventoryDevice(organizationId, serial string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/inventoryDevices/%s",
		organizationId, serial)
	var datamodel = InventoryDevice{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
