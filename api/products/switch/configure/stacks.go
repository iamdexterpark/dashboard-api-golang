package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type StackDHCP struct {
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

type StackInterfaces []struct {
	StackInterface
}

type StackInterface struct {
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

type StackStaticRoutes []struct {
	StackStaticRoute
}

type StackStaticRoute struct {
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

type Stacks []struct {
	Stack
}

type Stack struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Serials []string `json:"serials"`
}

func GetStackDHCP(networkId, switchStackId, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/routing/interfaces/%s/dhcp",
		 networkId, switchStackId, interfaceId)
	var datamodel = StackDHCP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutStackDHCP(networkId, switchStackId, interfaceId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/routing/interfaces/%s/dhcp",
		 networkId, switchStackId, interfaceId)
	var datamodel = StackDHCP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStackInterfaces(networkId, switchStackId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/routing/interfaces",
		 networkId, switchStackId)
	var datamodel = StackInterfaces{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStackInterface(networkId, switchStackId, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/routing/interfaces/%s",
		 networkId, switchStackId, interfaceId)
	var datamodel = StackInterface{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelStackInterface(networkId, switchStackId, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/routing/interfaces/%s",
		 networkId, switchStackId, interfaceId)
	var datamodel = StackInterface{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutStackInterface(networkId, switchStackId, interfaceId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/routing/interfaces/%s",
		 networkId, switchStackId, interfaceId)
	var datamodel = StackInterface{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostStackInterface(networkId, switchStackId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/routing/interfaces",
		 networkId, switchStackId)
	var datamodel = StackInterface{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStackStaticRoutes(networkId, switchStackId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%srouting/interfaces",
		 networkId, switchStackId)
	var datamodel = StackStaticRoutes{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStackStaticRoute(networkId, switchStackId, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%srouting/interfaces/%s",
		 networkId, switchStackId, interfaceId)
	var datamodel = StackStaticRoute{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelStackStaticRoute(networkId, switchStackId, interfaceId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%srouting/interfaces/%s",
		 networkId, switchStackId, interfaceId)
	var datamodel = StackStaticRoute{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutStackStaticRoute(networkId, switchStackId, interfaceId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%srouting/interfaces/%s",
		 networkId, switchStackId, interfaceId)
	var datamodel = StackStaticRoute{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostStackStaticRoute(networkId, switchStackId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%srouting/interfaces",
		 networkId, switchStackId)
	var datamodel = StackStaticRoute{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStacks(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks",
		 networkId)
	var datamodel = Stacks{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetStack(networkId, switchStackId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s",
		 networkId, switchStackId)
	var datamodel = Stack{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelStack(networkId, switchStackId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s",
		 networkId, switchStackId)
	var datamodel = Stack{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostStack(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks",
		 networkId)
	var datamodel = Stack{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostRemoveFromStack(networkId, switchStackId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/remove",
		 networkId, switchStackId)
	var datamodel = Stack{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostAddToStack(networkId, switchStackId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/switch/stacks/%s/add",
		 networkId, switchStackId)
	var datamodel = Stack{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}