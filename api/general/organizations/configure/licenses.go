package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
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

type AssignSeats struct {
	ResultingLicenses []struct {
		ID                        string        `json:"id"`
		LicenseType               string        `json:"licenseType"`
		LicenseKey                string        `json:"licenseKey"`
		OrderNumber               string        `json:"orderNumber"`
		DeviceSerial              interface{}   `json:"deviceSerial"`
		NetworkID                 string        `json:"networkId"`
		State                     string        `json:"state"`
		SeatCount                 int           `json:"seatCount"`
		TotalDurationInDays       int           `json:"totalDurationInDays"`
		DurationInDays            int           `json:"durationInDays"`
		PermanentlyQueuedLicenses []interface{} `json:"permanentlyQueuedLicenses"`
		ClaimDate                 time.Time     `json:"claimDate"`
		ActivationDate            time.Time     `json:"activationDate"`
		ExpirationDate            time.Time     `json:"expirationDate"`
	} `json:"resultingLicenses"`
}

type MoveLicenses struct {
	DestOrganizationID string   `json:"destOrganizationId"`
	LicenseIds         []string `json:"licenseIds"`
}

type MoveSeats struct {
	DestOrganizationID string `json:"destOrganizationId"`
	LicenseID          string `json:"licenseId"`
	SeatCount          int    `json:"seatCount"`
}

type RenewSeats struct {
	ResultingLicenses []struct {
		ID                        string        `json:"id"`
		LicenseType               string        `json:"licenseType"`
		LicenseKey                string        `json:"licenseKey"`
		OrderNumber               string        `json:"orderNumber"`
		DeviceSerial              interface{}   `json:"deviceSerial"`
		NetworkID                 string        `json:"networkId"`
		State                     string        `json:"state"`
		SeatCount                 int           `json:"seatCount"`
		TotalDurationInDays       int           `json:"totalDurationInDays"`
		DurationInDays            int           `json:"durationInDays"`
		PermanentlyQueuedLicenses []interface{} `json:"permanentlyQueuedLicenses"`
		ClaimDate                 time.Time     `json:"claimDate"`
		ActivationDate            time.Time     `json:"activationDate"`
		ExpirationDate            time.Time     `json:"expirationDate"`
	} `json:"resultingLicenses"`
}


func GetLicenses(organizationId, perPage, startingAfter,
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


func GetLicense(organizationId, licenseId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/%s", api.BaseUrl(),
		organizationId, licenseId)
	var datamodel = License{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutLicense(organizationId, licenseId, deviceSerial string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/%s", api.BaseUrl(),
		organizationId, licenseId)
	var datamodel = License{}
	// Parameters for Request URL
	var parameters = map[string]string{
		"deviceSerial": deviceSerial}
	sessions, err := api.Sessions(baseurl, "PUT", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PostAssignSeats(organizationId, deviceSerial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/assignSeats", api.BaseUrl(),
		organizationId)
	var datamodel = AssignSeats{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostMoveLicenses(organizationId, deviceSerial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/move", api.BaseUrl(),
		organizationId)
	var datamodel = AssignSeats{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostMoveSeats(organizationId, deviceSerial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/moveSeats", api.BaseUrl(),
		organizationId)
	var datamodel = MoveSeats{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func PostRenewSeats(organizationId, deviceSerial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/licenses/renewSeats", api.BaseUrl(),
		organizationId)
	var datamodel = RenewSeats{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
