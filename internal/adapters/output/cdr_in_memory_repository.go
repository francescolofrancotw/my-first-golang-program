package output

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"my-first-golang-program/internal/domain"
)

type CdrInMemoryRepository struct {
	storage map[string][]byte
}

func NewCdrInMemoryRepository() CdrInMemoryRepository {
	inMemoryRepository := CdrInMemoryRepository{
		storage: map[string][]byte{},
	}

	cdr, _ := domain.CreateNewCdr()

	cdrMarshalled, err := json.Marshal(cdr)

	if err != nil {
		panic(err.Error())
	}

	inMemoryRepository.storage[cdr.Id.String()] = cdrMarshalled

	return inMemoryRepository
}

func (repo CdrInMemoryRepository) Get(id uuid.UUID) (domain.Cdr, error) {

	if value, ok := repo.storage[id.String()]; ok {
		cdr := domain.Cdr{}
		err := json.Unmarshal(value, &cdr)

		if err != nil {
			return domain.Cdr{}, errors.New("error")
		}

		return cdr, nil
	}

	return domain.Cdr{}, errors.New("cdr not found in in-mem storage")
}

func (repo CdrInMemoryRepository) Save(cdr domain.Cdr) error {
	return errors.New("not implemented")
}
