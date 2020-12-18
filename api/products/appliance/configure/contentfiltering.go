package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type ContentFilteringCategories struct {
	Categories []interface{} `json:"categories"`
}

type ContentFiltering struct {
	AllowedURLPatterns   []string `json:"allowedUrlPatterns"`
	BlockedURLPatterns   []string `json:"blockedUrlPatterns"`
	BlockedURLCategories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"blockedUrlCategories"`
	URLCategoryListSize string `json:"urlCategoryListSize"`
}

func GetContentFilteringCategories(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/contentFiltering/categories",  networkId)
	var datamodel = ContentFilteringCategories{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func GetContentFiltering(networkId string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/contentFiltering",  networkId)
	var datamodel = ContentFiltering{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutContentFiltering(networkId string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/appliance/contentFiltering",  networkId)
	var datamodel = ContentFiltering{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
