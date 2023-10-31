package internal

type LocatorRepo interface {
}

type LocatorService struct {
	repo LocatorRepo
}

func NewLocatorService(repo LocatorRepo) *LocatorService {
	return &LocatorService{repo: repo}
}
