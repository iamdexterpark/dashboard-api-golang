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
	Enabled                 bool     `json:"enabled"`
	PoeEnabled              bool     `json:"poeEnabled"`
	Type                    string   `json:"type"`
	Vlan                    int      `json:"vlan"`
	VoiceVlan               int      `json:"voiceVlan"`
	IsolationEnabled        bool     `json:"isolationEnabled"`
	RstpEnabled             bool     `json:"rstpEnabled"`
	StpGuard                string   `json:"stpGuard"`
	LinkNegotiation         string   `json:"linkNegotiation"`
	PortScheduleID          string   `json:"portScheduleId"`
	Udld                    string   `json:"udld"`
	AccessPolicyType        string   `json:"accessPolicyType"`
	StickyMacAllowList      []string `json:"stickyMacAllowList"`
	StickyMacAllowListLimit int      `json:"stickyMacAllowListLimit"`
	StormControlEnabled     bool     `json:"stormControlEnabled"`
}

type ConfigTemplatesProfiles []struct {
	SwitchProfileID string `json:"switchProfileId"`
	Name            string `json:"name"`
	Model           string `json:"model"`
}

func GetSwitchPortProfiles(organizationId, configTemplateId, profileId  string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles/%s/ports",
		api.BaseUrl(), organizationId, configTemplateId, profileId)
	var datamodel = SwitchProfile{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSwitchPortProfile(organizationId, configTemplateId, profileId, portId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles/%s/ports/%s",
		api.BaseUrl(), organizationId, configTemplateId, profileId, portId)
	var datamodel = SwitchProfile{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSwitchPortProfile(organizationId, configTemplateId, profileId, portId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles/%s/ports/%s",
		api.BaseUrl(), organizationId, configTemplateId, profileId, portId)
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
	baseurl := fmt.Sprintf("%s/organizations/%s/configTemplates/%s/switch/profiles",
		api.BaseUrl(), organizationId, configTemplateId)
	var datamodel = ConfigTemplatesProfiles{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
