package repository

import (
	"encoding/json"
	"fmt"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"gitlab.com/samkomarov/locator-svc.git/internal/setup"
	"io"
	"net/http"
)

type ExternalAPILocatorRepo struct {
	cfg setup.ExternalAPILocatorConfig
}

func NewExternalAPILocatorRepo(cfg setup.ExternalAPILocatorConfig) *ExternalAPILocatorRepo {
	return &ExternalAPILocatorRepo{cfg}
}

func (e *ExternalAPILocatorRepo) GetAllMissing() ([]service.MissingPerson, error) {
	resp, err := http.Get(e.cfg.EndpointURL)
	if err != nil {
		return nil, fmt.Errorf("while making request to external api: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("making request to external api returned %d: %s", resp.StatusCode, body)
	}
	var body struct {
		Result []service.MissingPerson `json:"result"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, fmt.Errorf("while decoding external api's response body: %v", err)
	}
	return body.Result, nil
}
