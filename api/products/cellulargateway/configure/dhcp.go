package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type DHCP struct {
	DhcpLeaseTime        string   `json:"dhcpLeaseTime"`
	DNSNameservers       string   `json:"dnsNameservers"`
	DNSCustomNameservers []string `json:"dnsCustomNameservers"`
}

// List common DHCP settings of MGs
func GetDhcp(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/dhcp",
		api.BaseUrl(), networkId)
	var datamodel = DHCP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutDhcp(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/cellularGateway/dhcp",
		api.BaseUrl(), networkId)
	var datamodel = DHCP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}