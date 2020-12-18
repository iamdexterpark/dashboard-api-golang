package api

import (
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

// Set Environmental Variables
func EnvironmentalVariables()  {

	// read value ENV variable
	viper.AutomaticEnv()

	// Ensure API Token is set correctly
	_, ok := viper.Get("MERAKI_DASHBOARD_API_KEY").(string)
	if !ok {
		log.Fatal("Please set your MERAKI_DASHBOARD_API_KEY environmental variable to use this tool. ")
	}

	// Set default values

	// API Version
	viper.SetDefault("MERAKI_DASHBOARD_API_VERSION", "v1")

	// Base URL
	viper.SetDefault("MERAKI_DASHBOARD_API_URL", "https://api-mp.meraki.com/api/")



}


// Create a struct that allows us to pass data through a single interface
type ResponseData struct {
	Pagination []user_agent.Link
}

// request of API Call
type Request struct {
	method              string
	baseurl             string
	parameters          string
	payload             io.ReadSeeker
	DashboardAPIVersion string
}

// HTTP Protocol
type Response struct {
	Status     string      // e.g. "200 OK"
	StatusCode int         // e.g. 200
	Proto      string      // e.g. "HTTP/1.0"
	ProtoMajor int         // e.g. 1
	ProtoMinor int         // e.g. 0
	Header     http.Header // response headers
	// Body  io.ReadCloser  // response body includes entire cert chain...
}

// result of API Call
type Results struct {
	Request
	Response
	Pagination []user_agent.Link
	Payload    interface{}
}

// Session initializes a communication channel with the Meraki Dashboard API
func Session(apikey, baseurl, method string, payload io.ReadSeeker,
	parameter map[string]string, datamodel interface{}) (Results, error) {
	restSession := user_agent.MerakiClient()
	// response variable
	var session *http.Response
	var err error

	apiVersion := viper.GetString("MERAKI_DASHBOARD_API_VERSION")
	// Determine HTTP Method
	switch method {

	// Get-based API Calls
	case "GET":
		session, err = restSession.Get(baseurl, apiVersion, apikey)
		if err != nil {
			log.Fatal(err)
		}

	// Create-based API Calls
	case "POST":
		session, err = restSession.Post(baseurl, apiVersion, apikey, payload)
		if err != nil {
			log.Fatal(err)
		}

	// Update-based API Calls
	case "PUT":
		session, err = restSession.Put(baseurl, apiVersion, apikey, payload)
		if err != nil {
			log.Fatal(err)
		}

	// Delete-based API Calls
	case "DELETE":
		session, err = restSession.Delete(baseurl, apiVersion, apikey)
		if err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal("Unable to determine HTTP Method: ", method)
	}

	// append parameters to url
	parameters := session.Request.URL.Query()

	for param, value := range parameter {
		parameters.Add(param, value)
	}
	session.Request.URL.RawQuery = parameters.Encode()

	// marshal data into model
	user_agent.UnMarshalJSON(session.Body, &datamodel)

	dashboardRequest := Request{
		method:              method,
		baseurl:             baseurl,
		parameters:          session.Request.URL.RawQuery,
		payload:             payload,
		DashboardAPIVersion: viper.GetString("MERAKI_DASHBOARD_API_VERSION"),
	}

	dashboardResponse := Response{
		Status:     session.Status,
		StatusCode: session.StatusCode,
		Proto:      session.Proto,
		ProtoMajor: session.ProtoMajor,
		ProtoMinor: session.ProtoMinor,
		Header:     session.Header,
		// Body:       session.Body,
	}

	metadata := Results{

		Request:    dashboardRequest,
		Response:   dashboardResponse,
		Pagination: user_agent.ParseHeader(session.Header),
		Payload:    datamodel,
	}

	return metadata, err
}

func Sessions(resource string, method string, payload io.ReadSeeker,
	parameters map[string]string, datamodel interface{}) ([]Results, error) {

	// Set Environmental Variables
	EnvironmentalVariables()

	// Declare var
	apiKey := viper.GetString("MERAKI_DASHBOARD_API_KEY")
	baseUrl :=  viper.GetString("MERAKI_DASHBOARD_API_URL")
	version :=  viper.GetString("MERAKI_DASHBOARD_API_VERSION")

	var sessions []Results

	// Assemble url
	url := baseUrl+version+resource
	session, err := Session(apiKey, url, method, payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}

	sessions = append(sessions, session)

	for _, page := range session.Pagination {
		// Update parameters
		parameters["perPage"] = page.PerPage
		parameters["startingAfter"] = page.StartingAfter
		parameters["endingBefore"] = page.EndingBefore

		if page.Rel == "next" || page.Rel == "last" {
			paginatedSession, err := Session(apiKey, page.URI, "GET", payload, parameters, datamodel)
			if err != nil {
				log.Fatal(err)
			}
			sessions = append(sessions, paginatedSession)
		}

	}
	return sessions, err
}
