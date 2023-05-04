package table

import (
	"github.com/srrrs-7/micro_service.git/template/src/pkg/domain"
)

type tableRepository struct {
}

func NewTableRepository() *domain.DomainEntity {
	return &domain.DomainEntity{}
}

func (r *tableRepository) GetTable() ([]*domain.DomainEntity, error) {
	return nil, nil
}
