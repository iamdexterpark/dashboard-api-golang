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
	MerakiAuthUserID string    `json:"merakiAuthUserId"`
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	CreatedAt        time.Time `json:"createdAt"`
	AccountType      string    `json:"accountType"`
	Authorizations   []struct {
		SsidNumber        int       `json:"ssidNumber"`
		AuthorizedZone    string    `json:"authorizedZone"`
		ExpiresAt         time.Time `json:"expiresAt"`
		AuthorizedByName  string    `json:"authorizedByName"`
		AuthorizedByEmail string    `json:"authorizedByEmail"`
	} `json:"authorizations"`
}

// List The Users Configured Under Meraki Authentication For A Network Splash Guest Or RADIUS Users For A Wireless Network Or Client VPN Users For A Wired Network
func GetMerakiAuthUsers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/merakiAuthUsers", api.BaseUrl(), networkId)
	var datamodel = MerakiAuthUsers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetMerakiAuthUser(networkId, merakiAuthUserId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/merakiAuthUsers/%s", api.BaseUrl(), networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelMerakiAuthUser(networkId, merakiAuthUserId string) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/merakiAuthUsers/%s", api.BaseUrl(), networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutMerakiAuthUser(networkId, merakiAuthUserId, name, password, emailPasswordToUser, authorizations string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/merakiAuthUsers/%s", api.BaseUrl(), networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"password": password,
		"emailPasswordToUser": emailPasswordToUser,
		"authorizations": authorizations}

	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostMerakiAuthUser(networkId, name, password, accountType, emailPasswordToUser, authorizations string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/networks/%s/merakiAuthUsers", api.BaseUrl(), networkId)
	var datamodel = MerakiAuthUser{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"name": name,
		"password": password,
		"accountType": accountType,
		"emailPasswordToUser": emailPasswordToUser,
		"authorizations": authorizations}

	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)

	if err != nil {
		log.Fatal(err)
	}
	return sessions
}