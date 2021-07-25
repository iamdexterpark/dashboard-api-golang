package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type Networks []struct {
	Network
}

type Network struct {
	ID               string   `json:"id"`
	OrganizationID   string   `json:"organizationId"`
	Name             string   `json:"name"`
	TimeZone         string   `json:"timeZone"`
	Tags             []string `json:"tags"`
	ProductTypes     []string `json:"productTypes"`
	EnrollmentString string   `json:"enrollmentString"`
}

type CreateNetwork struct {
	Name             string   `json:"name"`
	ProductTypes     []string `json:"productTypes"`
	Tags             []string `json:"tags"`
	TimeZone         string   `json:"timeZone"`
	CopyFromNetworkId   string   `json:"copyFromNetworkId"`
	Notes   string   `json:"notes"`
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

type CombineNetworks struct {
	ResultingNetwork struct {
		ID               string   `json:"id"`
		OrganizationID   string   `json:"organizationId"`
		Name             string   `json:"name"`
		TimeZone         string   `json:"timeZone"`
		Tags             []string `json:"tags"`
		ProductTypes     []string `json:"productTypes"`
		EnrollmentString string   `json:"enrollmentString"`
	} `json:"resultingNetwork"`
}

// List the networks that the user has privileges on in an organization
func GetNetworks(organizationId, configTemplateId, tags, tagsFilterType, perPage,
	startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/networks",
		organizationId)

	// Parameters for Request URL
	var parameters = map[string]string{
		"configTemplateId": configTemplateId,
		"tags":             tags,
		"tagsFilterType":   tagsFilterType,
		"perPage":          perPage,
		"startingAfter":    startingAfter,
		"endingBefore":     endingBefore}

	var datamodel = Networks{}
	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Create a network
func PostNetworks(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/networks",
		organizationId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = CreateNetwork{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Combine multiple networks into a single network
func PostCombineNetworks(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/networks/combine",
		organizationId)
	payload := user_agent.MarshalJSON(data)
	var datamodel = Networks{}
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List the clients that have used this network in the timespan
func GetClients(networkId, t0, t1, timespan,
	perPage, startingAfter, endingBefore string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/clients",  networkId)
	var datamodel = Clients{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":            t0,
		"t1":            t1,
		"timespan":      timespan,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	return sessions
}