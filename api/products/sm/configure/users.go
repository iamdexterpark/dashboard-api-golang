package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type SmUserProfiles []struct {
	DeviceID    string `json:"deviceId"`
	ID          string `json:"id"`
	IsEncrypted bool   `json:"isEncrypted"`
	IsManaged   bool   `json:"isManaged"`
	ProfileData struct {
	} `json:"profileData"`
	ProfileIdentifier string `json:"profileIdentifier"`
	Name              string `json:"name"`
	Version           string `json:"version"`
}

type SmUserSoftware []struct {
	AppID             string      `json:"appId"`
	BundleSize        interface{} `json:"bundleSize"`
	CreatedAt         time.Time   `json:"createdAt"`
	DeviceID          string      `json:"deviceId"`
	DynamicSize       interface{} `json:"dynamicSize"`
	ID                string      `json:"id"`
	Identifier        string      `json:"identifier"`
	InstalledAt       time.Time   `json:"installedAt"`
	ToInstall         interface{} `json:"toInstall"`
	IosRedemptionCode interface{} `json:"iosRedemptionCode"`
	IsManaged         bool        `json:"isManaged"`
	ItunesID          interface{} `json:"itunesId"`
	LicenseKey        interface{} `json:"licenseKey"`
	Name              string      `json:"name"`
	Path              string      `json:"path"`
	RedemptionCode    interface{} `json:"redemptionCode"`
	ShortVersion      interface{} `json:"shortVersion"`
	Status            interface{} `json:"status"`
	ToUninstall       bool        `json:"toUninstall"`
	UninstalledAt     interface{} `json:"uninstalledAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
	Vendor            string      `json:"vendor"`
	Version           string      `json:"version"`
}

type SmUsers []struct {
	ID                     string        `json:"id"`
	Email                  string        `json:"email"`
	FullName               string        `json:"fullName"`
	Username               string        `json:"username"`
	HasPassword            bool          `json:"hasPassword"`
	Tags                   []string      `json:"tags"`
	AdGroups               []interface{} `json:"adGroups"`
	AsmGroups              []interface{} `json:"asmGroups"`
	IsExternal             bool          `json:"isExternal"`
	DisplayName            string        `json:"displayName"`
	HasIdentityCertificate bool          `json:"hasIdentityCertificate"`
	UserThumbnail          string        `json:"userThumbnail"`
}

func GetSmUserProfiles(networkId, userId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/users/%s/deviceProfiles",
		 networkId, userId)
	var datamodel = SmUserProfiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSmUserSoftware(networkId, userId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/users/%s/softwares",
		 networkId, userId)
	var datamodel = SmUserSoftware{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSmUsers(networkId, ids, usernames, emails, scope string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/users",
		 networkId)
	var datamodel = SmUsers{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"ids":       ids,
		"usernames": usernames,
		"emails":    emails,
		"scope":     scope}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
