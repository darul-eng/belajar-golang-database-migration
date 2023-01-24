package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{
		Error: true,
	}
}

type SimpleService struct {
	SimpleRepository *SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed created service")
	} else {
		return &SimpleService{SimpleRepository: repository}, nil
	}
}
