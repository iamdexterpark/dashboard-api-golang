package main

import (
	apiclient "github.com/ddexterpark/dashboard-api-golang/client"
	"github.com/ddexterpark/dashboard-api-golang/client/organizations"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/kr/pretty"
	"log"
	"os"
)

func main() {
	apikey := os.Getenv("MERAKI_DASHBOARD_API_KEY")
	host := "api.meraki.com"
	path := "/api/v1"
	orgid := "762234236932456587"

	transport := httptransport.New(host, path, []string{"https"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", apikey)
	c := apiclient.New(transport, nil)

	params := organizations.NewDeleteOrganizationParams()
	params.OrganizationID = orgid

	prod, err := c.Organizations.DeleteOrganization(params, transport.DefaultAuthentication)
	if err != nil {
		log.Fatal(err)
	}
	// Returns Data
	pretty.Println(prod.Error())
}
