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

	transport := httptransport.New(host, path, []string{"https"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("X-Cisco-Meraki-API-Key", "header", apikey)
	c := apiclient.New(transport, nil)

	params := organizations.NewCreateOrganizationParams()
	name := "testOrg"
	params.CreateOrganization.Name = &name

	prod, err := c.Organizations.CreateOrganization(params, transport.DefaultAuthentication)
	if err != nil {
		log.Fatal(err)
	}
	// Returns Data
	pretty.Println(prod.GetPayload().ID)
	pretty.Println(prod.GetPayload().Licensing)
	pretty.Println(prod.GetPayload().API)
	pretty.Println(prod.GetPayload().URL)
	pretty.Println(prod.GetPayload().Name)
	pretty.Println(prod.GetPayload().Cloud.Region.Name)

}
