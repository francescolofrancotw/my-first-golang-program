package ports

import (
	"github.com/google/uuid"
	"my-first-golang-program/internal/domain"
)

type CdrRepositoryPort interface {
	Get(id uuid.UUID) (domain.Cdr, error)
	Save(cdr domain.Cdr) error
}
