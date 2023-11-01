package repository

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/patrickmn/go-cache"
	"gitlab.com/samkomarov/locator-svc.git/internal/config"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"io"
	"net/http"
	"time"
)

const cacheExpirationInterval = 10 * time.Minute
const missingCacheKey = "missing"

type ExternalAPILocatorRepo struct {
	cfg    config.ExternalAPILocatorConfig
	client *retryablehttp.Client
	cache  *cache.Cache
}

func NewExternalAPILocatorRepo(cfg config.ExternalAPILocatorConfig) *ExternalAPILocatorRepo {
	client := retryablehttp.NewClient()
	client.RetryWaitMax = 30 * time.Second

	cacheClient := cache.New(cacheExpirationInterval, cacheExpirationInterval/2)
	return &ExternalAPILocatorRepo{cfg, client, cacheClient}
}

func (e *ExternalAPILocatorRepo) GetAllMissing() ([]service.MissingPerson, error) {
	missing, ok := e.cache.Get(missingCacheKey)
	if !ok {
		var err error
		missing, err = e.getAllMissingFromAPI()
		if err != nil {
			return nil, fmt.Errorf("while getting from api: %v", err)
		}
		e.cache.Set(missingCacheKey, missing, cache.DefaultExpiration)
	}
	return missing.([]service.MissingPerson), nil
}

func (e *ExternalAPILocatorRepo) getAllMissingFromAPI() ([]service.MissingPerson, error) {
	resp, err := e.client.Get(e.cfg.EndpointURL)
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
		return nil, fmt.Errorf("while decoding external api's response body: %w", err)
	}
	return body.Result, nil
}
