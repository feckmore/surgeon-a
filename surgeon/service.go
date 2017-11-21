package surgeon

// domain level validation & logic

import (
	"context"
	"errors"
	"log"
)

// Servicer provides access to the
type Servicer interface {
	//TODO: define & implement methods for getting all surgeons and saving a surgeon
	GetSurgeonByID(context.Context, ID) (*Surgeon, error)
}

// Repositor provides access to a surgeon store.
// from kit/examples/shipping/cargo/cargo.go
type Repositor interface {
	Store(surgeon *Surgeon) error
	FindByID(id ID) (*Surgeon, error)
	FindAll() []*Surgeon
	Close()
}

// surgeonService satisfies the SurgeonServicer interface
type surgeonService struct {
	surgeons Repositor
}

// NewSurgeonService allows access to instance of object that implements the Servicer interface
func NewSurgeonService(repo Repositor) Servicer {
	log.Println("NewSurgeonService:", repo)
	service := new(surgeonService)
	service.surgeons = repo

	return service
}

// GetSurgeonByID returns a surgeon.
func (us surgeonService) GetSurgeonByID(ctx context.Context, id ID) (*Surgeon, error) {
	log.Println("surgeonService.GetSurgeonByID:", id)

	return us.surgeons.FindByID(id)
}

// ErrUnknown is used when a surgeon could not be found.
var ErrUnknown = errors.New("unknown surgeon")
