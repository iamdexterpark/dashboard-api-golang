package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type Licenses []struct {
	License
}

type License struct {
	ID                        string      `json:"id"`
	LicenseType               string      `json:"licenseType"`
	LicenseKey                string      `json:"licenseKey"`
	OrderNumber               string      `json:"orderNumber"`
	DeviceSerial              string      `json:"deviceSerial"`
	NetworkID                 string      `json:"networkId"`
	State                     string      `json:"state"`
	SeatCount                 interface{} `json:"seatCount"`
	TotalDurationInDays       int         `json:"totalDurationInDays"`
	DurationInDays            int         `json:"durationInDays"`
	PermanentlyQueuedLicenses []struct {
		ID             string `json:"id"`
		LicenseType    string `json:"licenseType"`
		LicenseKey     string `json:"licenseKey"`
		OrderNumber    string `json:"orderNumber"`
		DurationInDays int    `json:"durationInDays"`
	} `json:"permanentlyQueuedLicenses"`
	ClaimDate      time.Time `json:"claimDate"`
	ActivationDate time.Time `json:"activationDate"`
	ExpirationDate time.Time `json:"expirationDate"`
}

// List The Licenses For An Organization
func GetOrganizationLicenses(organizationId, perPage, startingAfter,
	endingBefore, deviceSerial, networkId, state string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses", api.BaseUrl(),
		organizationId)
	var datamodel = Licenses{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":          perPage,
		"startingAfter":    startingAfter,
		"endingBefore":		endingBefore,
		"deviceSerial": 	deviceSerial,
		"networkId":		networkId,
		"state":			state,
	}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List A Single License For An Organization
func GetOrganizationLicense(organizationId, licenseId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/%s", api.BaseUrl(),
		organizationId, licenseId)

	var datamodel = License{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
