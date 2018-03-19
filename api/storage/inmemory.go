package storage

import (
	"sync"

	"github.com/feckmore/surgeon-a/api/surgeon"
)

// surgeonInMemoryRepo satisfies the Repositor interface
type surgeonInMemoryRepo struct {
	mtx      sync.RWMutex
	surgeons map[surgeon.ID]*surgeon.Surgeon
}

// NewSurgeonInMemoryRepository returns object that implements the Repositor interface
func NewSurgeonInMemoryRepository() surgeon.Repositor {
	return &surgeonInMemoryRepo{
		// surgeons: make(map[surgeon.ID]*surgeon.Surgeon),
		surgeons: map[surgeon.ID]*surgeon.Surgeon{
			1: {
				ID:        1,
				Firstname: "Dilbert",
				Lastname:  "Adams",
				Email:     "dilbert@adams.com",
			},
		},
	}
}
func (ur *surgeonInMemoryRepo) Store(surgeon *surgeon.Surgeon) error {
	ur.mtx.Lock()
	defer ur.mtx.Unlock()
	ur.surgeons[surgeon.ID] = surgeon
	return nil
}

func (ur *surgeonInMemoryRepo) FindByID(id surgeon.ID) (*surgeon.Surgeon, error) {
	ur.mtx.RLock()
	defer ur.mtx.RUnlock()
	if val, ok := ur.surgeons[id]; ok {
		return val, nil
	}
	return nil, surgeon.ErrUnknown
}

func (ur *surgeonInMemoryRepo) FindAll() []*surgeon.Surgeon {
	ur.mtx.RLock()
	defer ur.mtx.RUnlock()
	u := make([]*surgeon.Surgeon, 0, len(ur.surgeons))
	for _, val := range ur.surgeons {
		u = append(u, val)
	}

	return u
}

func (ur *surgeonInMemoryRepo) Close() {
	ur.surgeons = nil // not sure if
}
