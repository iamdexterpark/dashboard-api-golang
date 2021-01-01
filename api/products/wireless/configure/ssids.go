package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)



type SSID struct {
	Number              int    `json:"number"`
	Name                string `json:"name"`
	Enabled             bool   `json:"enabled"`
	SplashPage          string `json:"splashPage"`
	SsidAdminAccessible bool   `json:"ssidAdminAccessible"`
	AuthMode            string `json:"authMode"`
	EncryptionMode      string `json:"encryptionMode"`
	WpaEncryptionMode   string `json:"wpaEncryptionMode"`
	RadiusServers       []struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Secret string `json:"secret"`
	} `json:"radiusServers"`
	RadiusAccountingEnabled         bool     `json:"radiusAccountingEnabled"`
	RadiusEnabled                   bool     `json:"radiusEnabled"`
	RadiusAttributeForGroupPolicies string   `json:"radiusAttributeForGroupPolicies"`
	RadiusFailoverPolicy            string   `json:"radiusFailoverPolicy"`
	RadiusLoadBalancingPolicy       string   `json:"radiusLoadBalancingPolicy"`
	IPAssignmentMode                string   `json:"ipAssignmentMode"`
	AdminSplashURL                  string   `json:"adminSplashUrl"`
	SplashTimeout                   string   `json:"splashTimeout"`
	WalledGardenEnabled             bool     `json:"walledGardenEnabled"`
	WalledGardenRanges              []string `json:"walledGardenRanges"`
	MinBitrate                      int      `json:"minBitrate"`
	BandSelection                   string   `json:"bandSelection"`
	PerClientBandwidthLimitUp       int      `json:"perClientBandwidthLimitUp"`
	PerClientBandwidthLimitDown     int      `json:"perClientBandwidthLimitDown"`
	Visible                         bool     `json:"visible"`
	AvailableOnAllAps               bool     `json:"availableOnAllAps"`
	AvailabilityTags                []string `json:"availabilityTags"`
	PerSsidBandwidthLimitUp         int      `json:"perSsidBandwidthLimitUp"`
	PerSsidBandwidthLimitDown       int      `json:"perSsidBandwidthLimitDown"`
	MandatoryDhcpEnabled            bool     `json:"mandatoryDhcpEnabled"`
}

type SSIDS []struct {
	SSID
}


func GetSSIDS(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids",  networkId)
	var datamodel SSIDS
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSSID(networkId, ssidNumber string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s",  networkId, ssidNumber)
	var datamodel SSID
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// payload := user_agent.FormatPayload(data, datamodel)
func PutSSID(networkId, ssidNumber string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s",  networkId, ssidNumber)

	payload := user_agent.MarshalJSON(data)
	var datamodel SSID
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
