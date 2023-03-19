package covid

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Data is the response from the COVID API.
type Data struct {
	Records []Record `json:"Data"`
}

// Record is the data of each COVID case.
// The data is untrust and unclean, so we need to use pointer to handle the null value.
// For example, the Age field is null for some records.
type Record struct {
	ConfirmDate    *string `json:"ConfirmDate"`
	No             *int    `json:"No"`
	Age            *int    `json:"Age"`
	Gender         *string `json:"Gender"`
	GenderEn       *string `json:"GenderEn"`
	Nation         *string `json:"Nation"`
	NationEn       *string `json:"NationEn"`
	Province       *string `json:"Province"`
	ProvinceID     *int    `json:"ProvinceId"`
	District       *string `json:"District"`
	ProvinceEn     *string `json:"ProvinceEn"`
	StatQuarantine *int    `json:"StatQuarantine"`
}

// Fetcher provides the ability to fetch data from the COVID API.
type Fetcher interface {
	Fetch() (*Data, error)
}

type client struct {
	http *http.Client
	url  string
}

// NewClient creates a new client.
func NewClient(url string) *client {
	return &client{
		http: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
		url: url,
	}
}

// Fetch fetches data from the COVID API.
func (c *client) Fetch() (*Data, error) {
	req, err := http.NewRequest(http.MethodGet, c.url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, string(b))
	}

	var data Data
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
