package configure

import (
	"fmt"
	"github.com/ddexterpark/dashboard-api-golang/api"
	user_agent "github.com/ddexterpark/dashboard-api-golang/user-agent"
	"log"
)

type LoginSecurity struct {
	EnforcePasswordExpiration bool     `json:"enforcePasswordExpiration"`
	PasswordExpirationDays    int      `json:"passwordExpirationDays"`
	EnforceDifferentPasswords bool     `json:"enforceDifferentPasswords"`
	NumDifferentPasswords     int      `json:"numDifferentPasswords"`
	EnforceStrongPasswords    bool     `json:"enforceStrongPasswords"`
	EnforceAccountLockout     bool     `json:"enforceAccountLockout"`
	AccountLockoutAttempts    int      `json:"accountLockoutAttempts"`
	EnforceIdleTimeout        bool     `json:"enforceIdleTimeout"`
	IdleTimeoutMinutes        int      `json:"idleTimeoutMinutes"`
	EnforceTwoFactorAuth      bool     `json:"enforceTwoFactorAuth"`
	EnforceLoginIPRanges      bool     `json:"enforceLoginIpRanges"`
	LoginIPRanges             []string `json:"loginIpRanges"`
}


func GetLoginSecurity(organizationId string) []api.Results {
	baseurl := fmt.Sprintf("%s/organizations/%s/loginSecurity", api.BaseUrl(),
		organizationId)

	var datamodel = LoginSecurity{}
	sessions, err := api.Sessions(baseurl, "GET", nil, nil, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}

func PutLoginSecurity(organizationId, enforcePasswordExpiration, passwordExpirationDays,
	enforceDifferentPasswords, numDifferentPasswords, enforceStrongPasswords,
	enforceAccountLockout, accountLockoutAttempts, enforceIdleTimeout,
	idleTimeoutMinutes, enforceTwoFactorAuth, enforceLoginIpRanges,
	loginIpRanges string, data interface{}) []api.Results{
	baseurl := fmt.Sprintf("%s/organizations/%s/loginSecurity", api.BaseUrl(),
		organizationId)
	var datamodel = LoginSecurity{}
	payload := user_agent.MarshalJSON(data)
	// Parameters for Request URL
	var parameters = map[string]string{
		"cenforcePasswordExpiration": enforcePasswordExpiration,
		"passwordExpirationDays": passwordExpirationDays,
		"enforceDifferentPasswords": enforceDifferentPasswords,
		"numDifferentPasswords": numDifferentPasswords,
		"enforceStrongPasswords": enforceStrongPasswords,
		"enforceAccountLockout": enforceAccountLockout,
		"accountLockoutAttempts": accountLockoutAttempts,
		"enforceIdleTimeout": enforceIdleTimeout,
		"idleTimeoutMinutes": idleTimeoutMinutes,
		"enforceTwoFactorAuth": enforceTwoFactorAuth,
		"enforceLoginIpRanges": enforceLoginIpRanges,
		"loginIpRanges": loginIpRanges}
	sessions, err := api.Sessions(baseurl, "PUT", payload, parameters, datamodel)
	if err != nil {
		log.Fatal(err)
	}
	return sessions
}
