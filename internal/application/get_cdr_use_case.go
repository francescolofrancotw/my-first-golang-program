package application

import (
	"errors"
	"github.com/google/uuid"
	"my-first-golang-program/internal/domain"
	"my-first-golang-program/internal/domain/ports"
)

type GetCdrUseCase struct {
	cdrRepository ports.CdrRepositoryPort
}

func NewGetCdrUseCase(cdrRepository ports.CdrRepositoryPort) *GetCdrUseCase {
	return &GetCdrUseCase{
		cdrRepository: cdrRepository,
	}
}

func (srv *GetCdrUseCase) GetCdr(id uuid.UUID) (domain.Cdr, error) {
	cdr, err := srv.cdrRepository.Get(id)

	if err != nil {
		return domain.Cdr{}, errors.New("cdr not found")
	}

	return cdr, nil
}
