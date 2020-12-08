package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

// Data Structure
type BrandingPolicyPriorities struct {
	BrandingPolicyIds []string `json:"brandingPolicyIds"`
}

type BrandingPolicies []struct {
	BrandingPolicy
}

type BrandingPolicy struct {
	BrandingPolicyID string `json:"brandingPolicyId"`
	Name             string `json:"name"`
	Enabled          bool   `json:"enabled"`
	AdminSettings    struct {
		AppliesTo string   `json:"appliesTo"`
		Values    []string `json:"values"`
	} `json:"adminSettings"`
	HelpSettings struct {
		HelpTab                            string `json:"helpTab"`
		GetHelpSubtab                      string `json:"getHelpSubtab"`
		CommunitySubtab                    string `json:"communitySubtab"`
		CasesSubtab                        string `json:"casesSubtab"`
		DataProtectionRequestsSubtab       string `json:"dataProtectionRequestsSubtab"`
		GetHelpSubtabKnowledgeBaseSearch   string `json:"getHelpSubtabKnowledgeBaseSearch"`
		UniversalSearchKnowledgeBaseSearch string `json:"universalSearchKnowledgeBaseSearch"`
		CiscoMerakiProductDocumentation    string `json:"ciscoMerakiProductDocumentation"`
		SupportContactInfo                 string `json:"supportContactInfo"`
		NewFeaturesSubtab                  string `json:"newFeaturesSubtab"`
		FirewallInfoSubtab                 string `json:"firewallInfoSubtab"`
		APIDocsSubtab                      string `json:"apiDocsSubtab"`
		HardwareReplacementsSubtab         string `json:"hardwareReplacementsSubtab"`
		SmForums                           string `json:"smForums"`
	} `json:"helpSettings"`
}


func GetBrandingPolicyPriorities(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/priorities", api.BaseUrl(),
		organizationId)

	var datamodel = BrandingPolicyPriorities{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutBrandingPolicyPriorities(organizationId, brandingPolicyIds string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/priorities", api.BaseUrl(),
		organizationId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = BrandingPolicyPriorities{}
	// Parameters for Request URL
	var parameters = map[string]string{
		"brandingPolicyIds": brandingPolicyIds}
	sessions, err := api.Sessions(baseurl, "GET", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetBrandingPolicies(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies", api.BaseUrl(),
		organizationId)

	var datamodel = BrandingPolicies{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetBrandingPolicy(organizationId, brandingPolicyId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/%s", api.BaseUrl(),
		organizationId, brandingPolicyId)
	var datamodel = BrandingPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelBrandingPolicy(organizationId, brandingPolicyId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/%s", api.BaseUrl(),
		organizationId, brandingPolicyId)
	var datamodel = BrandingPolicy{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutBrandingPolicy(organizationId, brandingPolicyId, name, enabled, adminSettings, helpSettings string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies/%s", api.BaseUrl(),
		organizationId, brandingPolicyId)
	var datamodel = BrandingPolicy{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"enabled": enabled,
		"adminSettings": adminSettings,
		"helpSettings": helpSettings}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostBrandingPolicy(organizationId, name, enabled, adminSettings, helpSettings string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/brandingPolicies", api.BaseUrl(),
		organizationId)
	var datamodel = BrandingPolicy{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"enabled": enabled,
		"adminSettings": adminSettings,
		"helpSettings": helpSettings}
	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}