package service

type LocatorRepo interface {
}

type LocatorService struct {
	repo LocatorRepo
}

func NewLocatorService(repo LocatorRepo) *LocatorService {
	return &LocatorService{repo: repo}
}

func (l LocatorService) GetMissing() ([]MissingPerson, error) {
	//TODO implement me
	panic("implement me")
}
