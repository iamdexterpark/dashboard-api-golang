package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type LoginSecurity struct {
	EnforcePasswordExpiration string `json:"enforcePasswordExpiration"`
	PasswordExpirationDays    string      `json:"passwordExpirationDays"`
	EnforceDifferentPasswords string `json:"enforceDifferentPasswords"`
	NumDifferentPasswords     string      `json:"numDifferentPasswords"`
	EnforceStrongPasswords    string `json:"enforceStrongPasswords"`
	EnforceAccountLockout     string `json:"enforceAccountLockout"`
	AccountLockoutAttempts    string      `json:"accountLockoutAttempts"`
	EnforceIdleTimeout        string `json:"enforceIdleTimeout"`
	IdleTimeoutMinutes        string      `json:"idleTimeoutMinutes"`
	EnforceTwoFactorAuth      string `json:"enforceTwoFactorAuth"`
	EnforceLoginIPRanges      string `json:"enforceLoginIpRanges"`
	LoginIPRanges             []string `json:"loginIpRanges"`
}

// Returns the login security settings for an organization.
func GetLoginSecurity(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("/organizations/%s/loginSecurity",
		organizationId)

	var datamodel = LoginSecurity{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

// Update the login security settings for an organization
func PutLoginSecurity(organizationId string, data interface{}) []api.Results{
	baseurl := fmt.Sprintf("/organizations/%s/loginSecurity",
		organizationId)
	var datamodel = LoginSecurity{}
	payload := user_agent.MarshalJSON(data)
	sessions, err := api.Sessions(baseurl, "PUT", payload, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
