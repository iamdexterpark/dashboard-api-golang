package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type SubnetPool struct {
	Cidr    string `json:"cidr"`
	Mask    string `json:"mask"`
	Subnets []struct {
		Serial      string `json:"serial"`
		Name        string `json:"name"`
		ApplianceIP string `json:"applianceIp"`
		Subnet      string `json:"subnet"`
	} `json:"subnets"`
}

// Return the subnet pool and mask configured for MGs in the network.
func GetSubnetPool(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/subnetPool",
		api.BaseUrl(), networkId)
	var datamodel = SubnetPool{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSubnetPool(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/subnetPool",
		api.BaseUrl(), networkId)
	var datamodel = SubnetPool{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}