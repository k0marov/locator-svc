package service

import (
	"fmt"
	"sort"
)

type LocatorRepo interface {
	GetAllMissing() ([]LizaAlertPerson, error)
}

type LocatorService struct {
	repo LocatorRepo
}

func NewLocatorService(repo LocatorRepo) *LocatorService {
	return &LocatorService{repo: repo}
}

const Limit = 10

// GetRelevantMissing returns missing people near `aroundPoint` obtained from LocatorRepo ordered by relevancy from highest to lowest limited by Limit
func (l *LocatorService) GetRelevantMissing(aroundPoint GeoPoint) ([]MissingPerson, error) {
	allMissingModels, err := l.repo.GetAllMissing()
	if err != nil {
		return nil, fmt.Errorf("while getting all missing from repo: %w", err)
	}
	allMissing := make([]MissingPerson, len(allMissingModels))
	for i, m := range allMissingModels {
		allMissing[i] = MissingPerson{
			VerticalURL: m.VerticalURL,
			DateOfLoss:  m.DateOfLoss,
			GeoPoints:   m.GeoPoints,
		}
	}

	sort.Slice(allMissing, func(i, j int) bool {
		return personRelevance(aroundPoint, allMissing[i]) > personRelevance(aroundPoint, allMissing[j])
	})
	return allMissing[:min(Limit, len(allMissing))], nil
}
