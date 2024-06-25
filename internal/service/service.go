package service

import (
	"fmt"
	"sort"
	"sync"
	"time"
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

const Limit = 20

// GetRelevantMissing returns missing people near `aroundPoint` obtained from LocatorRepo ordered by relevancy from highest to lowest limited by Limit
func (l *LocatorService) GetRelevantMissing(aroundPoint GeoPoint) ([]MissingPerson, error) {
	allMissingModels, err := l.repo.GetAllMissing()
	if err != nil {
		return nil, fmt.Errorf("while getting all missing from repo: %w", err)
	}

	allMissing := make([]MissingPerson, len(allMissingModels))
	var wg sync.WaitGroup
	wg.Add(len(allMissingModels))
	for i, m := range allMissingModels {
		go func(i int, m LizaAlertPerson) {
			allMissing[i] = MissingPerson{
				PhotoURL:   m.VerticalURL,
				DateOfLoss: time.Time(m.DateOfLoss),
				Locations:  m.GeoPoints,
				Relevance:  personRelevance(aroundPoint, m.GeoPoints, time.Time(m.DateOfLoss)),
			}
			wg.Done()
		}(i, m)
	}
	wg.Wait()

	sort.Slice(allMissing, func(i, j int) bool {
		return allMissing[i].Relevance > allMissing[j].Relevance
	})

	return allMissing[:min(Limit, len(allMissing))], nil
}
