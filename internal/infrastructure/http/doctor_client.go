package http

import (
	"fmt"
	"net/http"
	"time"
)

type DoctorHTTPClient struct {
	baseURL string
	client  *http.Client
}

func NewDoctorHTTPClient(baseURL string) *DoctorHTTPClient {
	return &DoctorHTTPClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 2 * time.Second},
	}
}

func (c *DoctorHTTPClient) DoctorExists(id string) (bool, error) {
	url := fmt.Sprintf("%s/doctors/%s", c.baseURL, id)
	resp, err := c.client.Get(url)
	if err != nil {
		return false, fmt.Errorf("doctor service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	return false, fmt.Errorf("doctor service error")
}