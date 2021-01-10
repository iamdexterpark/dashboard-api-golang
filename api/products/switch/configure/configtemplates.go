package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type SwitchProfile struct {
	PortID                  string   `json:"portId"`
	Name                    string   `json:"name"`
	Tags                    []string `json:"tags"`
	Enabled                 string `json:"enabled"`
	PoeEnabled              string `json:"poeEnabled"`
	Type                    string   `json:"type"`
	Vlan                    string      `json:"vlan"`
	VoiceVlan               string      `json:"voiceVlan"`
	IsolationEnabled        string `json:"isolationEnabled"`
	RstpEnabled             string `json:"rstpEnabled"`
	StpGuard                string   `json:"stpGuard"`
	LinkNegotiation         string   `json:"linkNegotiation"`
	PortScheduleID          string   `json:"portScheduleId"`
	Udld                    string   `json:"udld"`
	AccessPolicyType        string   `json:"accessPolicyType"`
	StickyMacAllowList      []string `json:"stickyMacAllowList"`
	StickyMacAllowListLimit string      `json:"stickyMacAllowListLimit"`
	StormControlEnabled     string `json:"stormControlEnabled"`
}

type ConfigTemplatesProfiles []struct {
	SwitchProfileID string `json:"switchProfileId"`
	Name            string `json:"name"`
	Model           string `json:"model"`
}

func GetSwitchPortProfiles(organizationId, configTemplateId, profileId  string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates/%s/switch/profiles/%s/ports",
		 organizationId, configTemplateId, profileId)
	var datamodel = SwitchProfile{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSwitchPortProfile(organizationId, configTemplateId, profileId, portId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates/%s/switch/profiles/%s/ports/%s",
		 organizationId, configTemplateId, profileId, portId)
	var datamodel = SwitchProfile{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSwitchPortProfile(organizationId, configTemplateId, profileId, portId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates/%s/switch/profiles/%s/ports/%s",
		 organizationId, configTemplateId, profileId, portId)
	var datamodel = SwitchProfile{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetConfigTemplatesProfiles(organizationId,
	configTemplateId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/configTemplates/%s/switch/profiles",
		 organizationId, configTemplateId)
	var datamodel = ConfigTemplatesProfiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
