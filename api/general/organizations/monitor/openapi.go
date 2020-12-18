package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
)

type OpenapiSpec struct {
	Swagger string `json:"swagger"`
	Info    struct {
		Version     string `json:"version"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"info"`
	Paths struct {
		Organizations struct {
			Get struct {
				Description string `json:"description"`
				OperationID string `json:"operationId"`
				Responses   struct {
					Num200 struct {
						Description string `json:"description"`
						Examples    struct {
							ApplicationJSON []struct {
								ID   string `json:"id"`
								Name string `json:"name"`
							} `json:"application/json"`
						} `json:"examples"`
					} `json:"200"`
				} `json:"responses"`
			} `json:"get"`
		} `json:"/organizations"`
	} `json:"paths"`
}

// Return the OpenAPI 2.0 Specification of the organization's API documentation in JSON
func GetOpenapiSpec(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/openapiSpec",  organizationId)
	var datamodel = OpenapiSpec{}

	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
