package monitor

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	"log"
	"time"
)

type APIRequestsSummary struct {
	ResponseCodeCounts struct {
		Num200 string `json:"200"`
		Num201 string `json:"201"`
		Num204 string `json:"204"`
		Num400 string `json:"400"`
		Num404 string `json:"404"`
		Num429 string `json:"429"`
	} `json:"responseCodeCounts"`
}

type APIRequests []struct {
	APIRequest
}
type APIRequest struct {
	AdminID      string `json:"adminId"`
	Method       string `json:"method"`
	Host         string `json:"host"`
	Path         string `json:"path"`
	QueryString  string `json:"queryString"`
	UserAgent    string `json:"userAgent"`
	Ts           time.Time `json:"ts"`
	ResponseCode string       `json:"responseCode"`
	SourceIP     string `json:"sourceIp"`
}

// Return an aggregated overview of API requests data
func GetAPIRequestSummary(organizationId, t0, t1, timespan string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/apiRequests/overview",  organizationId)
	var datamodel = APIRequestsSummary{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":       t0,
		"t1":       t1,
		"timespan": timespan}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Return an aggregated overview of API requests data
func GetAPIRequests(organizationId, t0, t1, timespan, perPage, startingAfter, endingBefore,
	adminId, path, method, responseCode, sourceIp string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/apiRequests",  organizationId)
	var datamodel = APIRequests{}

	// Parameters for Request URL
	var parameters = map[string]string{
		"t0":            t0,
		"t1":            t1,
		"timespan":      timespan,
		"perPage":       perPage,
		"startingAfter": startingAfter,
		"endingBefore":  endingBefore,
		"adminId":       adminId,
		"path":          path,
		"method":        method,
		"responseCode":  responseCode,
		"sourceIp":      sourceIp,
	}

	sessions, err := api.Sessions(baseurl, "GET", nil, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
