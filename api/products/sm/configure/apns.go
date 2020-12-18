package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type APNSCertificate struct {
	Certificate string `json:"certificate"`
}

func GetAPNSCertificate(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/sm/apnsCert",  organizationId)
	var datamodel = APNSCertificate{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
