package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type LAN struct {
	DeviceName         string `json:"deviceName"`
	DeviceLanIP        string `json:"deviceLanIp"`
	DeviceSubnet       string `json:"deviceSubnet"`
	FixedIPAssignments []struct {
		Mac  string `json:"mac"`
		Name string `json:"name"`
		IP   string `json:"ip"`
	} `json:"fixedIpAssignments"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
}

// Show the LAN Settings of a MG
func GetLan(serial string) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/cellularGateway/lan",
		api.BaseUrl(), serial)
	var datamodel = LAN{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutLan(serial, reservedIpRanges, fixedIpAssignments string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/devices/%s/cellularGateway/lan",
		api.BaseUrl(), serial)
	var datamodel = LAN{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"reservedIpRanges": reservedIpRanges,
		"fixedIpAssignments": fixedIpAssignments}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}