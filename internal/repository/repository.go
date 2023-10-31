package repository

import "gitlab.com/samkomarov/locator-svc.git/internal/setup"

type ExternalAPILocatorRepo struct {
	cfg setup.ExternalAPILocatorConfig
}

func NewExternalAPILocatorRepo(cfg setup.ExternalAPILocatorConfig) *ExternalAPILocatorRepo {
	return &ExternalAPILocatorRepo{cfg}
}
