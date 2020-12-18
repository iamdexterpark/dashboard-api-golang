package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type DHCP struct {
	DhcpMode             string   `json:"dhcpMode"`
	DhcpLeaseTime        string   `json:"dhcpLeaseTime"`
	DNSNameserversOption string   `json:"dnsNameserversOption"`
	DNSCustomNameservers []string `json:"dnsCustomNameservers"`
	BootOptionsEnabled   bool     `json:"bootOptionsEnabled"`
	BootNextServer       string   `json:"bootNextServer"`
	BootFileName         string   `json:"bootFileName"`
	DhcpOptions          []struct {
		Code  string `json:"code"`
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"dhcpOptions"`
	ReservedIPRanges []struct {
		Start   string `json:"start"`
		End     string `json:"end"`
		Comment string `json:"comment"`
	} `json:"reservedIpRanges"`
	FixedIPAssignments []struct {
		Mac  string `json:"mac"`
		Name string `json:"name"`
		IP   string `json:"ip"`
	} `json:"fixedIpAssignments"`
}

type Interfaces []struct {
	Interface
}

type Interface struct {
	InterfaceID      string `json:"interfaceId"`
	Name             string `json:"name"`
	Subnet           string `json:"subnet"`
	InterfaceIP      string `json:"interfaceIp"`
	MulticastRouting string `json:"multicastRouting"`
	VlanID           int    `json:"vlanId"`
	DefaultGateway   string `json:"defaultGateway"`
	OspfSettings     struct {
		Area             string `json:"area"`
		Cost             int    `json:"cost"`
		IsPassiveEnabled bool   `json:"isPassiveEnabled"`
	} `json:"ospfSettings"`
}

type RendezvousPoints [][]struct {
	RendezvousPoint
}

type RendezvousPoint []struct {
	RendezvousPointID string `json:"rendezvousPointId"`
	Serial            string `json:"serial,omitempty"`
	InterfaceName     string `json:"interfaceName,omitempty"`
	InterfaceIP       string `json:"interfaceIp"`
	MulticastGroup    string `json:"multicastGroup"`
	SwitchStackID     string `json:"switchStackId,omitempty"`
}

type Multicast struct {
	DefaultSettings struct {
		IgmpSnoopingEnabled                 bool `json:"igmpSnoopingEnabled"`
		FloodUnknownMulticastTrafficEnabled bool `json:"floodUnknownMulticastTrafficEnabled"`
	} `json:"defaultSettings"`
	Overrides []struct {
		Switches                            []string `json:"switches,omitempty"`
		IgmpSnoopingEnabled                 bool     `json:"igmpSnoopingEnabled"`
		FloodUnknownMulticastTrafficEnabled bool     `json:"floodUnknownMulticastTrafficEnabled"`
		Stacks                              []string `json:"stacks,omitempty"`
	} `json:"overrides"`
}

func GetDHCP(serial, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/routing/interfaces/%s/dhcp",
		 serial, interfaceId)
	var datamodel = DHCP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutDHCP(serial, interfaceId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/routing/interfaces/%s/dhcp",
		 serial, interfaceId)
	var datamodel = DHCP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "GET", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetInterfaces(serial string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/routing/interfaces",
		 serial)
	var datamodel = Interfaces{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetInterface(serial, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/routing/interfaces/%s",
		 serial, interfaceId)
	var datamodel = Interface{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelInterface(serial, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/routing/interfaces/%s",
		 serial, interfaceId)
	var datamodel = Interface{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutInterface(serial, interfaceId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/routing/interfaces/%s",
		 serial, interfaceId)
	var datamodel = Interface{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostInterface(serial string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/devices/%s/switch/routing/interfaces",
		 serial)
	var datamodel = Interface{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetRendezvousPoints(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/routing/multicast/rendezvousPoints",
		 networkId)
	var datamodel = RendezvousPoints{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetRendezvousPoint(networkId, rendezvousPointId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/routing/multicast/rendezvousPoints/%s",
		 networkId, rendezvousPointId)
	var datamodel = RendezvousPoint{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func delRendezvousPoint(networkId, rendezvousPointId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/routing/multicast/rendezvousPoints/%s",
		 networkId, rendezvousPointId)
	var datamodel = RendezvousPoint{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutRendezvousPoint(networkId, rendezvousPointId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/routing/multicast/rendezvousPoints/%s",
		 networkId, rendezvousPointId)
	var datamodel = RendezvousPoint{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostRendezvousPoint(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/routing/multicast/rendezvousPoints",
		 networkId)
	var datamodel = RendezvousPoint{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetMulticast(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/routing/multicast",
		 networkId)
	var datamodel = Multicast{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutMulticast(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/routing/multicast",
		 networkId)
	var datamodel = Multicast{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}