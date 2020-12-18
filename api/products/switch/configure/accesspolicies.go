package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type AccessPolicies []struct {
	AccessPolicy
}

type AccessPolicy struct {
	AccessPolicyNumber string `json:"accessPolicyNumber"`
	Name               string `json:"name"`
	RadiusServers      []struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Secret string `json:"secret"`
	} `json:"radiusServers"`
	RadiusTestingEnabled    bool `json:"radiusTestingEnabled"`
	RadiusCoaSupportEnabled bool `json:"radiusCoaSupportEnabled"`
	RadiusAccountingEnabled bool `json:"radiusAccountingEnabled"`
	RadiusAccountingServers []struct {
		Host   string `json:"host"`
		Port   int    `json:"port"`
		Secret string `json:"secret"`
	} `json:"radiusAccountingServers"`
	HostMode                       string `json:"hostMode"`
	AccessPolicyType               string `json:"accessPolicyType"`
	IncreaseAccessSpeed            bool   `json:"increaseAccessSpeed"`
	GuestVlanID                    int    `json:"guestVlanId"`
	VoiceVlanClients               bool   `json:"voiceVlanClients"`
	URLRedirectWalledGardenEnabled bool   `json:"urlRedirectWalledGardenEnabled"`
	URLRedirectWalledGardenRanges  string `json:"urlRedirectWalledGardenRanges"`
}

func GetAccessPolicies(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/accessPolicies",
		 networkId)
	var datamodel = AccessPolicies{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetAccessPolicy(networkId, accessPolicyNumber string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/accessPolicies/%s",
		 networkId, accessPolicyNumber)
	var datamodel = AccessPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelAccessPolicy(networkId, accessPolicyNumber string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/accessPolicies/%s",
		 networkId, accessPolicyNumber)
	var datamodel = AccessPolicy{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutAccessPolicy(networkId, accessPolicyNumber string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/accessPolicies/%s",
		 networkId, accessPolicyNumber)
	var datamodel = AccessPolicy{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostAccessPolicy(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/accessPolicies",
		 networkId)
	var datamodel = AccessPolicy{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}