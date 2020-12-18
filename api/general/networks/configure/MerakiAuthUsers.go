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

func GetMerakiAuthUsers(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/merakiAuthUsers",  networkId)
	var datamodel = MerakiAuthUsers{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetMerakiAuthUser(networkId, merakiAuthUserId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/merakiAuthUsers/%s",  networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelMerakiAuthUser(networkId, merakiAuthUserId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/merakiAuthUsers/%s",  networkId, merakiAuthUserId)
	var datamodel = MerakiAuthUser{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

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