package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type Device []struct {
	Name      string   `json:"name"`
	Lat       float64  `json:"lat"`
	Lng       float64  `json:"lng"`
	Address   string   `json:"address"`
	Notes     string   `json:"notes"`
	Tags      []string `json:"tags"`
	NetworkID string   `json:"networkId"`
	Serial    string   `json:"serial"`
	Model     string   `json:"model"`
	Mac       string   `json:"mac"`
	LanIP     string   `json:"lanIp"`
	Firmware  string   `json:"firmware"`
}

// Devices - List the devices in an organization
type Devices []struct {
	Name      string `json:"name"`
	Serial    string `json:"serial"`
	Mac       string `json:"mac"`
	Status    string `json:"status"`
	LanIP     string `json:"lanIp"`
	PublicIP  string `json:"publicIp"`
	NetworkID string `json:"networkId"`
}

// GetOrganizationDevices - List the devices in an organization
func GetOrganizationDevices(organizationId, perPage, startingAfter,
	configurationUpdatedAfter string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/devices", api.BaseUrl(), organizationId)
	var datamodel = Device{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"configurationUpdatedAfter": configurationUpdatedAfter,
		"perPage":                   perPage,
		"startingAfter":             startingAfter}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}
