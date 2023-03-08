package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Cdr struct {
	Id     uuid.UUID
	Field2 int
	Field3 bool
}

func CreateNewCdr() (Cdr, error) {
	cdrId, err := uuid.NewUUID()

	if err != nil {
		return Cdr{}, errors.New("unable to generate an id for the cdr")
	}

	return Cdr{
		Id:     cdrId,
		Field2: 100,
		Field3: true,
	}, nil
}
