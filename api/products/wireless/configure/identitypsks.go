package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type IdentityPSKs []struct {
	IdentityPSK
}
type IdentityPSK struct {
	IdentityPskID string `json:"identityPskId"`
	Name          string `json:"name"`
	Passphrase    string `json:"passphrase"`
	GroupPolicyID string `json:"groupPolicyId"`
}

func GetIdentityPSKs(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks",
		api.BaseUrl(), networkId, number)
	var datamodel IdentityPSKs
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetIdentityPSK(networkId, number, identityPskId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks/%s",
		api.BaseUrl(), networkId, number, identityPskId)
	var datamodel IdentityPSK
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelIdentityPSK(networkId, number, identityPskId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks/%s",
		api.BaseUrl(), networkId, number, identityPskId)
	var datamodel IdentityPSK
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutIdentityPSK(networkId, number, identityPskId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks/%s",
		api.BaseUrl(), networkId, number, identityPskId)
	var datamodel IdentityPSK
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostIdentityPSK(networkId, number string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/wireless/ssids/%s/identityPsks",
		api.BaseUrl(), networkId, number)
	var datamodel IdentityPSK
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}