package company

import (
	"context"
)

type Company struct {
	ID string
}

type CompanyStore interface {
	getCompany(ctx context.Context, id string) (Company, error)
}

type CompanyService struct {
	repository CompanyStore
}

func NewCompanyService(repository CompanyStore) *CompanyService {
	return &CompanyService{
		repository: repository,
	}
}

func (s *CompanyService) getCompany(ctx context.Context, id string) (Company, error) {
	company, err := s.repository.getCompany(ctx, id)

	if err != nil {
		return Company{}, err
	}

	return company, nil
}

func (s *CompanyService) createCompany(ctx context.Context, company *Company) error {
	return nil
}

func (s *CompanyService) updateCompany(ctx context.Context, company *Company) error {
	return nil
}

func (s *CompanyService) deleteCompany(ctx context.Context, id string) error {
	return nil
}
