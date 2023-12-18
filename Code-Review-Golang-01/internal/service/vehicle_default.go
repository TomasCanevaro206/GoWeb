package service

import (
	"app/internal"
	"errors"
	"fmt"
)

// NewDefault returns a new instance of a vehicle service.
func NewDefault(rp internal.RepositoryVehicle) *Default {
	return &Default{rp: rp}
}

// Default is an struct that represents a vehicle service.
type Default struct {
	rp internal.RepositoryVehicle
}

// FindAll returns all vehicles.
func (s *Default) FindAll() (v []internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.FindAll()
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) FindByMarcaAños(brand string, minYear int, maxYear int) (v []internal.Vehicle, err error) {
	v, err = s.rp.FindByMarcaAños(brand, minYear, maxYear)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}
	return
}