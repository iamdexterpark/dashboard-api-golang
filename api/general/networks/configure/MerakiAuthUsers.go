package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
	"time"
)

type MerakiAuthUsers []struct {
	MerakiAuthUser
}
type MerakiAuthUser struct {
	MerakiAuthUserID string `json:"merakiAuthUserId"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	CreatedAt        time.Time `json:"createdAt"`
	AccountType      string `json:"accountType"`
	Authorizations   []struct {
		SsidNumber        string       `json:"ssidNumber"`
		AuthorizedZone    string `json:"authorizedZone"`
		ExpiresAt         time.Time `json:"expiresAt"`
		AuthorizedByName  string `json:"authorizedByName"`
		AuthorizedByEmail string `json:"authorizedByEmail"`
	} `json:"authorizations"`
}

// List the users configured under Meraki Authentication for
// a network (splash guest or RADIUS users for a wireless network,
// or client VPN users for a wired network)
func GetMerakiAuthUsers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/merakiAuthUsers",  networkId)
	var datamodel = MerakiAuthUsers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Create a user configured with Meraki Authentication for a network
//(currently supports 802.1X, splash guest, and client VPN users,
//and currently, organizations have a 50,000 user cap
func PostMerakiAuthUser(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/merakiAuthUsers",  networkId)
	var datamodel = MerakiAuthUser{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)

	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// List the users configured under Meraki Authentication for a network
// (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)
func GetMerakiAuthUser(networkId, merakiAuthUserId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/merakiAuthUsers/%s",  networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Delete a user configured with Meraki Authentication
//(currently, 802.1X RADIUS, splash guest,
//and client VPN users can be deleted)
func DelMerakiAuthUser(networkId, merakiAuthUserId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/merakiAuthUsers/%s",  networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update a user configured with Meraki Authentication
//(currently, 802.1X RADIUS, splash guest,
//and client VPN users can be deleted)
func PutMerakiAuthUser(networkId, merakiAuthUserId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/merakiAuthUsers/%s",  networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

