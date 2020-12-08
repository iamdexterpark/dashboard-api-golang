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
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idps", api.BaseUrl(),
		organizationId)
	var datamodel = IDPS{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}


func GetIDP(organizationId, idpId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idp/%s", api.BaseUrl(),
		organizationId, idpId)

	var datamodel = IDP{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func DelIDP(organizationId, idpId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idp/%s", api.BaseUrl(),
		organizationId, idpId)

	var datamodel = IDP{}
	sessions, err := api.Sessions(baseurl, "DELETE", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutIDP(organizationId, idpId, x509certSha1Fingerprint, sloLogoutUrl string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idp/%s", api.BaseUrl(),
		organizationId, idpId)
	var datamodel = IDP{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"x509certSha1Fingerprint": x509certSha1Fingerprint,
		"sloLogoutUrl": sloLogoutUrl}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PostIDP(organizationId, x509certSha1Fingerprint, sloLogoutUrl string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml/idp", api.BaseUrl(),
		organizationId)
	var datamodel = IDP{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"x509certSha1Fingerprint": x509certSha1Fingerprint,
		"sloLogoutUrl": sloLogoutUrl}
	sessions, err := api.Sessions(baseurl, "POST", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetSAML(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml", api.BaseUrl(),
		organizationId)

	var datamodel = SAML{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSAML(organizationId, enabled string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/saml", api.BaseUrl(),
		organizationId)
	var datamodel = SAML{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"enabled": enabled}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
