package monitor

import (
	"github.com/ddexterpark/dashboard-api-golang/api"
	"fmt"
	"log"
)

type UsageHistory struct {
	ID                   string      `json:"id"`
	Description          string      `json:"description"`
	Mac                  string      `json:"mac"`
	IP                   string      `json:"ip"`
	User                 string      `json:"user"`
	Vlan                 string      `json:"vlan"`
	Switchport           interface{} `json:"switchport"`
	IP6                  string      `json:"ip6"`
	FirstSeen            int         `json:"firstSeen"`
	LastSeen             int         `json:"lastSeen"`
	Manufacturer         string      `json:"manufacturer"`
	Os                   string      `json:"os"`
	Ssid                 string      `json:"ssid"`
	WirelessCapabilities string      `json:"wirelessCapabilities"`
	SmInstalled          bool        `json:"smInstalled"`
	RecentDeviceMac      string      `json:"recentDeviceMac"`
	ClientVpnConnections []struct {
		RemoteIP       string `json:"remoteIp"`
		ConnectedAt    int    `json:"connectedAt"`
		DisconnectedAt int    `json:"disconnectedAt"`
	} `json:"clientVpnConnections"`
	Lldp   [][]string  `json:"lldp"`
	Cdp    interface{} `json:"cdp"`
	Status string      `json:"status"`
}

type Clients []struct {
	Usage struct {
		Sent int `json:"sent"`
		Recv int `json:"recv"`
	} `json:"usage"`
	ID                 string      `json:"id"`
	Description        string      `json:"description"`
	Mac                string      `json:"mac"`
	IP                 string      `json:"ip"`
	User               string      `json:"user"`
	Vlan               int         `json:"vlan"`
	Switchport         interface{} `json:"switchport"`
	IP6                string      `json:"ip6"`
	FirstSeen          int         `json:"firstSeen"`
	LastSeen           int         `json:"lastSeen"`
	Manufacturer       string      `json:"manufacturer"`
	Os                 string      `json:"os"`
	RecentDeviceSerial string      `json:"recentDeviceSerial"`
	RecentDeviceName   string      `json:"recentDeviceName"`
	RecentDeviceMac    string      `json:"recentDeviceMac"`
	Ssid               string      `json:"ssid"`
	Status             string      `json:"status"`
	Notes              string      `json:"notes"`
	IP6Local           string      `json:"ip6Local"`
	SmInstalled        bool        `json:"smInstalled"`
	GroupPolicy8021X   string      `json:"groupPolicy8021x"`
}


type Client struct {
	ID                   string      `json:"id"`
	Description          string      `json:"description"`
	Mac                  string      `json:"mac"`
	IP                   string      `json:"ip"`
	User                 string      `json:"user"`
	Vlan                 string      `json:"vlan"`
	Switchport           interface{} `json:"switchport"`
	IP6                  string      `json:"ip6"`
	FirstSeen            int         `json:"firstSeen"`
	LastSeen             int         `json:"lastSeen"`
	Manufacturer         string      `json:"manufacturer"`
	Os                   string      `json:"os"`
	Ssid                 string      `json:"ssid"`
	WirelessCapabilities string      `json:"wirelessCapabilities"`
	SmInstalled          bool        `json:"smInstalled"`
	RecentDeviceMac      string      `json:"recentDeviceMac"`
	ClientVpnConnections []struct {
		RemoteIP       string `json:"remoteIp"`
		ConnectedAt    int    `json:"connectedAt"`
		DisconnectedAt int    `json:"disconnectedAt"`
	} `json:"clientVpnConnections"`
	Lldp   [][]string  `json:"lldp"`
	Cdp    interface{} `json:"cdp"`
	Status string      `json:"status"`
}

// Return the client's daily usage history. Usage data is in kilobytes.
// Clients can be identified by a client key or either the MAC or IP
// depending on whether the network uses Track-by-IP.
func GetUsageHistory(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s/usageHistory", networkId, clientId)
	var datamodel = UsageHistory{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


// List the clients that have used this network in the timespan
func GetClients(networkId, t0, timespan, perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients", networkId)
	var datamodel = Clients{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"perPage":                    perPage,
		"startingAfter":              startingAfter,
		"endingBefore":               endingBefore,
		"t0":                         t0,
		"timespan":                   timespan,
	}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return the client associated with the given identifier. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.
func GetClient(networkId, clientId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients/%s", networkId, clientId)
	var datamodel = Client{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}