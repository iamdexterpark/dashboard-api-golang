package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

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

// Return the branding policy IDs of an organization in priority order.
// IDs are ordered in ascending order of priority
// (IDs later in the array have higher priority).
func GetBrandingPolicyPriorities(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/brandingPolicies/priorities",
		organizationId)

	var datamodel = BrandingPolicyPriorities{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the priority ordering of an organization's branding policies.
func PutBrandingPolicyPriorities(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/brandingPolicies/priorities",
		organizationId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = BrandingPolicyPriorities{}
	sessions, err := api.Sessions(baseurl, "GET", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List the branding policies of an organization
func GetBrandingPolicies(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/brandingPolicies",
		organizationId)

	var datamodel = BrandingPolicies{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return a branding policy
func GetBrandingPolicy(organizationId, brandingPolicyId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/brandingPolicies/%s",
		organizationId, brandingPolicyId)
	var datamodel = BrandingPolicy{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Delete a branding policy
func DelBrandingPolicy(organizationId, brandingPolicyId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/brandingPolicies/%s",
		organizationId, brandingPolicyId)
	var datamodel = BrandingPolicy{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update a branding policy
func PutBrandingPolicy(organizationId, brandingPolicyId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/brandingPolicies/%s",
		organizationId, brandingPolicyId)
	var datamodel = BrandingPolicy{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Add a new branding policy to an organization
func PostBrandingPolicy(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/brandingPolicies",
		organizationId)
	var datamodel = BrandingPolicy{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}