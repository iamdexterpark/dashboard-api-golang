package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type SplashPageSettings struct {
	SsidNumber     string `json:"ssidNumber"`
	SplashPage     string `json:"splashPage"`
	SplashURL      string `json:"splashUrl"`
	UseSplashURL   string `json:"useSplashUrl"`
	RedirectURL    string `json:"redirectUrl"`
	UseRedirectURL string `json:"useRedirectUrl"`
	WelcomeMessage string `json:"welcomeMessage"`
	SplashLogo     struct {
		Md5       string `json:"md5"`
		Extension string `json:"extension"`
	} `json:"splashLogo"`
	SplashImage struct {
		Md5 interface{} `json:"md5"`
	} `json:"splashImage"`
	SplashPrepaidFront struct {
		Md5 interface{} `json:"md5"`
	} `json:"splashPrepaidFront"`
}

func GetSplashPageSettings(networkId, number string) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s/splash/settings",
		 networkId, number)
	var datamodel SplashPageSettings
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutSplashPageSettings(networkId, number string, data interface{}) []api.Results {
	baseurl := fmt.Sprintf("/networks/%s/wireless/ssids/%s/splash/settings",
		 networkId, number)
	var datamodel SplashPageSettings
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}