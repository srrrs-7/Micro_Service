package domain

type domainService interface {
}

type domainRepository struct {
}

func NewDomainService() domainService {
	return &domainRepository{}
}

func (d *domainRepository) DomainService() error {
	return nil
}
