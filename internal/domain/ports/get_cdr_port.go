package ports

import (
	"github.com/google/uuid"
	"my-first-golang-program/internal/domain"
)

type GetCdrPort interface {
	GetCdr(id uuid.UUID) (domain.Cdr, error)
}
