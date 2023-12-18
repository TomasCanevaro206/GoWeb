package internal

import (
	"errors"
)

var (
	// ErrServiceVehicleNotFound is returned when no vehicle is found.
	ErrServiceVehicleNotFound = errors.New("service: vehicle not found")
)

// ServiceVehicle is the interface that wraps the basic methods for a vehicle service.
// - conections with external apis
// - business logic
type ServiceVehicle interface {
	// FindAll returns all vehicles
	FindAll() (v []Vehicle, err error)
	FindByMarcaAños(brand string, minYear int, maxYear int) (v []Vehicle, err error)
}