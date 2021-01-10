package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
	"time"
)

type DeviceCerts []struct {
	Name           string `json:"name"`
	NotValidAfter  time.Time `json:"notValidAfter"`
	NotValidBefore time.Time `json:"notValidBefore"`
	CertPem        string `json:"certPem"`
	DeviceID       string `json:"deviceId"`
	Issuer         string `json:"issuer"`
	Subject        string `json:"subject"`
	ID             string `json:"id"`
}

type DeviceProfiles []struct {
	DeviceID    string `json:"deviceId"`
	ID          string `json:"id"`
	IsEncrypted string `json:"isEncrypted"`
	IsManaged   string `json:"isManaged"`
	ProfileData struct {
	} `json:"profileData"`
	ProfileIdentifier string `json:"profileIdentifier"`
	Name              string `json:"name"`
	Version           string `json:"version"`
}

type NetworkAdapters []struct {
	DhcpServer string `json:"dhcpServer"`
	DNSServer  string `json:"dnsServer"`
	Gateway    string `json:"gateway"`
	ID         string `json:"id"`
	IP         string `json:"ip"`
	Mac        string `json:"mac"`
	Name       string `json:"name"`
	Subnet     string `json:"subnet"`
}

type DeviceRestrictions []struct {
	Profile      string `json:"profile"`
	Restrictions struct {
		MyRestriction struct {
			Value string `json:"value"`
		} `json:"myRestriction"`
	} `json:"restrictions"`
}

type DeviceSecurityCenters []struct {
	IsRooted             string `json:"isRooted"`
	HasAntiVirus         string `json:"hasAntiVirus"`
	AntiVirusName        string `json:"antiVirusName"`
	IsFireWallEnabled    string `json:"isFireWallEnabled"`
	HasFireWallInstalled string `json:"hasFireWallInstalled"`
	FireWallName         string `json:"fireWallName"`
	IsDiskEncrypted      string `json:"isDiskEncrypted"`
	IsAutoLoginDisabled  string `json:"isAutoLoginDisabled"`
	ID                   string `json:"id"`
	RunningProcs         string `json:"runningProcs"`
}

type DeviceSoftwares []struct {
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
	IsManaged         string        `json:"isManaged"`
	ItunesID          interface{} `json:"itunesId"`
	LicenseKey        interface{} `json:"licenseKey"`
	Name              string      `json:"name"`
	Path              string      `json:"path"`
	RedemptionCode    interface{} `json:"redemptionCode"`
	ShortVersion      interface{} `json:"shortVersion"`
	Status            interface{} `json:"status"`
	ToUninstall       string        `json:"toUninstall"`
	UninstalledAt     interface{} `json:"uninstalledAt"`
	UpdatedAt         time.Time   `json:"updatedAt"`
	Vendor            string      `json:"vendor"`
	Version           string      `json:"version"`
}

type DeviceWlanLists []struct {
	CreatedAt time.Time `json:"createdAt"`
	ID        string `json:"id"`
	XML       string `json:"xml"`
}

type Fields struct {
	ID      string `json:"id"`
	Serial  string `json:"serial"`
	WifiMac string `json:"wifiMac"`
	Name    string `json:"name"`
	Notes   string `json:"notes"`
}


type ModifyTags []struct {
	ID      string   `json:"id"`
	Serial  string   `json:"serial"`
	WifiMac string   `json:"wifiMac"`
	Tags    []string `json:"tags"`
}

type SmIds struct {
	Ids        string `json:"ids"`
	NewNetwork string `json:"newNetwork"`
}

type SMDevices struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Tags         []string `json:"tags"`
	Ssid         string   `json:"ssid"`
	WifiMac      string   `json:"wifiMac"`
	OsName       string   `json:"osName"`
	SystemModel  string   `json:"systemModel"`
	UUID         string   `json:"uuid"`
	SerialNumber string   `json:"serialNumber"`
	Serial       string   `json:"serial"`
	IP           string   `json:"ip"`
	Notes        string   `json:"notes"`
}

type Unenroll struct {
	Success string `json:"success"`
}

func GetDeviceCerts(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/certs",
		 networkId, deviceId)
	var datamodel = DeviceCerts{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetDeviceProfiles(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/deviceProfiles",
		 networkId, deviceId)
	var datamodel = DeviceProfiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutFields(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/fields",
		 networkId)
	var datamodel = Fields{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetNetworkAdapters(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/networkAdapters",
		 networkId, deviceId)
	var datamodel = NetworkAdapters{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetDeviceRestrictions(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/restrictions",
		 networkId, deviceId)
	var datamodel = DeviceRestrictions{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetDeviceSecurityCenters(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/securityCenters",
		 networkId, deviceId)
	var datamodel = DeviceSecurityCenters{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetDeviceSoftwares(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/softwares",
		 networkId, deviceId)

	var datamodel = DeviceSoftwares{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetDeviceWlanLists(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/wlanLists",
		 networkId, deviceId)

	var datamodel = DeviceWlanLists{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostModifyTags(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/modifyTags",
		 networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = ModifyTags{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostCheckin(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/checkin",
		 networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = SmIds{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSMDevices(networkId, fields, wifiMacs, serials, ids, scope, perPage,
	startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices",  networkId)
	var datamodel = SMDevices{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"fields":        fields,
		"wifiMacs":      wifiMacs,
		"serials":       serials,
		"ids":           ids,
		"scope":         scope,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostLock(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/lock",
		 networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = SmIds{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostMove(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/move",
		 networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = SmIds{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostRefreshDetails(networkId, deviceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/refreshDetails",
		 networkId, deviceId)
	var datamodel interface{}
	sessions, err := api.Sessions(baseurl, "POST", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostUnenroll(networkId, deviceId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/%s/unenroll",
		 networkId, deviceId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = Unenroll{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostWipeDevice(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/sm/devices/wipe",
		 networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = SmIds{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}