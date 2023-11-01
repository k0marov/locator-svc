package repository

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/patrickmn/go-cache"
	"gitlab.com/samkomarov/locator-svc.git/internal/config"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"io"
	"log"
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
	missingCached, ok := e.cache.Get(missingCacheKey)
	if ok {
		log.Printf("got missing people info from cache")
		return missingCached.([]service.MissingPerson), nil
	}
	log.Printf("missing people cache expired, fetching from API")
	missing, err := e.getAllMissingFromAPI()
	if err != nil {
		return nil, fmt.Errorf("while getting from api: %v", err)
	}
	e.cache.Set(missingCacheKey, missing, cache.DefaultExpiration)
	return missing, nil
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
	log.Printf("got %v missing people data from API", len(body.Result))
	return body.Result, nil
}
