package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type IDPS []struct {
	IDP
}

type IDP struct {
	IdpID                   string `json:"idpId"`
	ConsumerURL             string `json:"consumerUrl"`
	X509CertSha1Fingerprint string `json:"x509certSha1Fingerprint"`
	SloLogoutURL            string `json:"sloLogoutUrl"`
}

type SAML struct {
	Enabled bool `json:"enabled"`
}

func GetIDPS(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/saml/idps",
		organizationId)
	var datamodel = IDPS{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetIDP(organizationId, idpId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/saml/idp/%s",
		organizationId, idpId)

	var datamodel = IDP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelIDP(organizationId, idpId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/saml/idp/%s",
		organizationId, idpId)

	var datamodel = IDP{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutIDP(organizationId, idpId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/saml/idp/%s",
		organizationId, idpId)
	var datamodel = IDP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostIDP(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/saml/idp",
		organizationId)
	var datamodel = IDP{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "POST", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSAML(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/saml",
		organizationId)

	var datamodel = SAML{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSAML(organizationId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/saml",
		organizationId)
	var datamodel = SAML{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
