package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type BGP struct {
	Enabled       bool `json:"enabled"`
	AsNumber      int  `json:"asNumber"`
	IbgpHoldTimer int  `json:"ibgpHoldTimer"`
	Neighbors     []struct {
		IP             string `json:"ip"`
		RemoteAsNumber int    `json:"remoteAsNumber"`
		ReceiveLimit   int    `json:"receiveLimit"`
		AllowTransit   bool   `json:"allowTransit"`
		EbgpHoldTimer  int    `json:"ebgpHoldTimer"`
		EbgpMultihop   int    `json:"ebgpMultihop"`
	} `json:"neighbors"`
}

type SiteToSiteVPN struct {
	Mode string `json:"mode"`
	Hubs []struct {
		HubID           string `json:"hubId"`
		UseDefaultRoute bool   `json:"useDefaultRoute"`
	} `json:"hubs"`
	Subnets []struct {
		LocalSubnet string `json:"localSubnet"`
		UseVpn      bool   `json:"useVpn"`
	} `json:"subnets"`
}

type ThirdPartyVPNPeers struct {
	Peers []struct {
		Name           string   `json:"name"`
		PublicIP       string   `json:"publicIp"`
		PrivateSubnets []string `json:"privateSubnets"`
		Secret         string   `json:"secret"`
		IkeVersion     string   `json:"ikeVersion"`
		IpsecPolicies  struct {
			IkeCipherAlgo         []string `json:"ikeCipherAlgo"`
			IkeAuthAlgo           []string `json:"ikeAuthAlgo"`
			IkePrfAlgo            []string `json:"ikePrfAlgo"`
			IkeDiffieHellmanGroup []string `json:"ikeDiffieHellmanGroup"`
			IkeLifetime           int      `json:"ikeLifetime"`
			ChildCipherAlgo       []string `json:"childCipherAlgo"`
			ChildAuthAlgo         []string `json:"childAuthAlgo"`
			ChildPfsGroup         []string `json:"childPfsGroup"`
			ChildLifetime         int      `json:"childLifetime"`
		} `json:"ipsecPolicies,omitempty"`
		NetworkTags         []string `json:"networkTags"`
		RemoteID            string   `json:"remoteId,omitempty"`
		IpsecPoliciesPreset string   `json:"ipsecPoliciesPreset,omitempty"`
	} `json:"peers"`
}

type VpnFirewallRules struct {
	Rules []struct {
		Comment       string `json:"comment"`
		Policy        string `json:"policy"`
		Protocol      string `json:"protocol"`
		DestPort      int    `json:"destPort"`
		DestCidr      string `json:"destCidr"`
		SrcPort       string `json:"srcPort"`
		SrcCidr       string `json:"srcCidr"`
		SyslogEnabled bool   `json:"syslogEnabled"`
	} `json:"rules"`
}

func GetBGP(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/vpn/bgp",  networkId)
	var datamodel = BGP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutBGP(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/vpn/bgp",  networkId)
	var datamodel = BGP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSiteToSiteVPN(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/vpn/siteToSiteVpn",  networkId)
	var datamodel = SiteToSiteVPN{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSiteToSiteVPN(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/vpn/siteToSiteVpn",  networkId)
	var datamodel = SiteToSiteVPN{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetThirdPartyVPNPeers(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/vpn/thirdPartyVPNPeers",  organizationId)
	var datamodel = ThirdPartyVPNPeers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutThirdPartyVPNPeers(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/vpn/thirdPartyVPNPeers",  organizationId)
	var datamodel = ThirdPartyVPNPeers{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetVpnFirewallRules(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/vpn/vpnFirewallRules",  organizationId)
	var datamodel = VpnFirewallRules{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutVpnFirewallRules(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/vpn/vpnFirewallRules",  organizationId)
	var datamodel = VpnFirewallRules{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}