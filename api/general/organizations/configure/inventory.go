package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type DeviceInventory []struct {
	Inventory
}

type Inventory struct {
	Mac                   string    `json:"mac"`
	Serial                string    `json:"serial"`
	Name                  string    `json:"name"`
	Model                 string    `json:"model"`
	NetworkID             string    `json:"networkId"`
	OrderNumber           string    `json:"orderNumber"`
	ClaimedAt             time.Time `json:"claimedAt"`
	LicenseExpirationDate time.Time `json:"licenseExpirationDate"`
}

// Return The Device Inventory For An Organization
func GetInventoryDevices(organizationId, perPage, startingAfter, endingBefore, usedState, search string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/inventoryDevices", api.BaseUrl(),
		organizationId)
	var datamodel = DeviceInventory{}

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

// Return A Single Device From The Inventory Of An Organization
func GetInventoryDevice(organizationId, serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/inventoryDevices/%s", api.BaseUrl(),
		organizationId, serial)

	var datamodel = Inventory{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
