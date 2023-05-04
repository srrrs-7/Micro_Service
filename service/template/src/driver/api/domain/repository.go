package domain

import "github.com/srrrs-7/micro_service.git/template/src/pkg/domain"

type server interface {
}

type domainRepository struct {
}

func NewTableRepository() domainInterface {
	return &domainRepository{}
}

func GetDomain() ([]*domain.DomainEntity, error) {
	return nil
}
