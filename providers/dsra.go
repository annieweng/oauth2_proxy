package providers

import (
	"log"
	"net/http"
	"net/url"
	"github.com/annieweng/oauth2_proxy/api"
)

type DsraProvider struct {
	*ProviderData
}

func NewDsraProvider(p *ProviderData) *DsraProvider {
	p.ProviderName = "Dsra"
	if p.LoginURL == nil || p.LoginURL.String() == "" {
		p.LoginURL = &url.URL{
			Scheme: "https",
			Host:   "xdataproxy.com",
			Path:   "/oauth/authorize",
		}
	}
	if p.RedeemURL == nil || p.RedeemURL.String() == "" {
		p.RedeemURL = &url.URL{
			Scheme: "https",
			Host:   "xdataproxy.com",
			Path:   "/oauth/token",
		}
	}
	if p.ValidateURL == nil || p.ValidateURL.String() == "" {
		p.ValidateURL = &url.URL{
			Scheme: "https",
			Host:   "xdataproxy.com",
			Path:   "/oauth/api/me",
		}
	}
	if p.Scope == "" {
		p.Scope = "api"
	}
	return &DsraProvider{ProviderData: p}
}

func (p *DsraProvider) GetEmailAddress(s *SessionState) (string, error) {

	req, err := http.NewRequest("GET",
		p.ValidateURL.String()+"?access_token="+s.AccessToken, nil)
	if err != nil {
		log.Printf("failed building request %s", err)
		return "", err
	}
	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed making request %s", err)
		return "", err
	}
	return json.Get("username").String()
}


