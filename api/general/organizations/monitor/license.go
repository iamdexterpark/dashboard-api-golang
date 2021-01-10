package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type License struct {
	Status               string `json:"status"`
	ExpirationDate       string `json:"expirationDate"`
	LicensedDeviceCounts struct {
		MS int `json:"MS"`
	} `json:"licensedDeviceCounts"`
}

// Return an overview of the license state for an organization
func GetLicenseOverview(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/licenses/overview",  organizationId)
	var datamodel = License{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
