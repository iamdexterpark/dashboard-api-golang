package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type AccessControlLists struct {
	Rules []struct {
		Comment   string `json:"comment"`
		Policy    string `json:"policy"`
		IPVersion string `json:"ipVersion"`
		Protocol  string `json:"protocol"`
		SrcCidr   string `json:"srcCidr"`
		SrcPort   string `json:"srcPort"`
		DstCidr   string `json:"dstCidr"`
		DstPort   string `json:"dstPort"`
		Vlan      string `json:"vlan"`
	} `json:"rules"`
}

func GetAccessControlLists(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/accessControlLists",
		 networkId)
	var datamodel = AccessControlLists{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutAccessControlLists(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/accessControlLists",
		 networkId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = AccessControlLists{}
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
