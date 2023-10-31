package repository

import (
	"gitlab.com/samkomarov/locator-svc.git/internal/service"
	"gitlab.com/samkomarov/locator-svc.git/internal/setup"
)

type ExternalAPILocatorRepo struct {
	cfg setup.ExternalAPILocatorConfig
}

func NewExternalAPILocatorRepo(cfg setup.ExternalAPILocatorConfig) *ExternalAPILocatorRepo {
	return &ExternalAPILocatorRepo{cfg}
}

func (e ExternalAPILocatorRepo) GetAllMissing() ([]service.MissingPerson, error) {
	//TODO implement me
	panic("implement me")
}
