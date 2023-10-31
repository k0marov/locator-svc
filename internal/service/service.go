package service

import (
	"fmt"
	"sort"
)

type LocatorRepo interface {
	GetAllMissing() ([]MissingPerson, error)
}

type LocatorService struct {
	repo LocatorRepo
}

func NewLocatorService(repo LocatorRepo) *LocatorService {
	return &LocatorService{repo: repo}
}

// GetRelevantMissing returns missing people near `aroundPoint` obtained from LocatorRepo ordered by relevancy from highest to lowest
func (l *LocatorService) GetRelevantMissing(aroundPoint GeoPoint) ([]MissingPerson, error) {
	allMissing, err := l.repo.GetAllMissing()
	if err != nil {
		return nil, fmt.Errorf("while getting all missing from repo: %w", err)
	}
	sort.Slice(allMissing, func(i, j int) bool {
		return personRelevance(aroundPoint, allMissing[i]) > personRelevance(aroundPoint, allMissing[j])
	})
	return allMissing, nil
}
