package delivery

import (
	"encoding/json"
	"fmt"
	"gitlab.com/samkomarov/locator-svc.git/internal/core"
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"net/http"
	"net/url"
	"strconv"
)

func DecodeLocationRequest(r *http.Request) (*service.GeoPoint, error) {
	urlQuery := r.URL.Query()
	latitude, err := parseFloatQuery(urlQuery, "latitude")
	if err != nil {
		return nil, err
	}
	longitude, err := parseFloatQuery(urlQuery, "longitude")
	if err != nil {
		return nil, err
	}
	return &service.GeoPoint{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}

func parseFloatQuery(query url.Values, name string) (float64, error) {
	value, err := strconv.ParseFloat(query.Get(name), 64)
	if err != nil {
		return 0, &core.ClientError{
			DisplayMessage: fmt.Sprintf("couldn't parse '%s' url parameter as float", name),
			HTTPCode:       http.StatusBadRequest,
		}
	}
	return value, nil
}

type MissingPersonResponse struct {
	PhotoURL string `json:"photo_url"`
}

func EncodeMissingPeopleResponse(w http.ResponseWriter, missing []service.MissingPerson) {
	resp := make([]MissingPersonResponse, len(missing))
	for i, m := range missing {
		resp[i] = MissingPersonResponse{PhotoURL: m.VerticalURL}
	}
	json.NewEncoder(w).Encode(resp)
}
