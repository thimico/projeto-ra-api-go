package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	UserAgent               = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36"
	PageViewsURL            = "/company/pageviews/%s"
	ShortNameURL            = "/company/shortname/%s"
	QueryURL                = "/query/%s/10/0"
	SearchURL               = "/companies/search/%s"
	TopBadStatusOf30DaysURL = "/company/rankings-full/topBadStatusOf30Days/10"
	ErrorCode				= "4000"
)

var (
	RAiositeURL = "https://iosite.reclameaqui.com.br/raichu-io-site-v1"
)

type ReclameAquiApiInterface interface {
	CountPageViewsRaExternalApi(ctx context.Context, complainName string) (int, error)
	IsInTopBad10RaExternalApi(ctx context.Context, complainName string) (bool, error)
}

type searchResponse struct {
	Results []complainResponse `json:"results"`
}

type PageviewsResponse struct {
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	PageViews int `json:"pageviews"`
}

type complainResponse struct {
	Name     string   `json:"name"`
	Reclaims []string `json:"reclaims"`
}

type TopBadResponse struct {
	Empresas []ScoreCompany `json:"data"`
}

type ScoreCompany struct {
	Score       float64 `json:"score"`
	CompanyName string  `json:"companyName"`
	Logo        string  `json:"logo"`
	ID          string  `json:"id"`
	Shortname   string  `json:"shortname"`
	Status      string  `json:"status"`
}

type ReclameAquiApi struct {
	APIURL string
}

func NewReclameAquiExternalApi() ReclameAquiApiInterface {
	return ReclameAquiApi{APIURL: RAiositeURL}
}

func (s ReclameAquiApi) CountPageViewsRaExternalApi(ctx context.Context, companyName string) (int, error) {

	u, _ := url.Parse(s.APIURL + fmt.Sprintf(PageViewsURL, strings.ToLower(companyName)))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var pageviewsResponse PageviewsResponse
	err = json.NewDecoder(resp.Body).Decode(&pageviewsResponse)
	if err != nil {
		return 0, err
	}

	if pageviewsResponse.Code == ErrorCode {
		return 0, nil
	}

	return pageviewsResponse.PageViews, nil
}


func (s ReclameAquiApi) IsInTopBad10RaExternalApi(ctx context.Context, companyName string) (bool, error) {

	u, _ := url.Parse(s.APIURL + TopBadStatusOf30DaysURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var topBadRAList TopBadResponse
	err = json.NewDecoder(resp.Body).Decode(&topBadRAList)
	if err != nil {
		return false, err
	}

	for _, p := range topBadRAList.Empresas {
		if strings.EqualFold(p.CompanyName, companyName) {
			return true, nil
		}
	}


	return false, nil
}