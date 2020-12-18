package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type SecurityIntrusion struct {
	Mode              string `json:"mode"`
	IdsRulesets       string `json:"idsRulesets"`
	ProtectedNetworks struct {
		UseDefault   bool     `json:"useDefault"`
		IncludedCidr []string `json:"includedCidr"`
		ExcludedCidr []string `json:"excludedCidr"`
	} `json:"protectedNetworks"`
}

type OrganizationSecurityIntrusion struct {
	AllowedRules []struct {
		RuleID  string `json:"ruleId"`
		Message string `json:"message"`
	} `json:"allowedRules"`
}

type MalwareSettings struct {
	Mode        string `json:"mode"`
	AllowedUrls []struct {
		URL     string `json:"url"`
		Comment string `json:"comment"`
	} `json:"allowedUrls"`
	AllowedFiles []struct {
		Sha256  string `json:"sha256"`
		Comment string `json:"comment"`
	} `json:"allowedFiles"`
}

func GetNetworkSecurityIntrusion(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/security/intrusion",  networkId)
	var datamodel = SecurityIntrusion{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutNetworkSecurityIntrusion(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/security/intrusion",  networkId)
	var datamodel = SecurityIntrusion{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetOrganizationSecurityIntrusion(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/appliance/security/intrusion",  organizationId)
	var datamodel = OrganizationSecurityIntrusion{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutOrganizationSecurityIntrusion(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/appliance/security/intrusion",  organizationId)
	var datamodel = OrganizationSecurityIntrusion{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetMalwareSettings(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/security/malware",  networkId)
	var datamodel = MalwareSettings{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutMalwareSettings(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/security/malware",  networkId)
	var datamodel = MalwareSettings{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
