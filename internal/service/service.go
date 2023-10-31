package service

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
func (l LocatorService) GetRelevantMissing(aroundPoint GeoPoint) ([]MissingPerson, error) {
	//TODO implement me
	panic("implement me")
}
