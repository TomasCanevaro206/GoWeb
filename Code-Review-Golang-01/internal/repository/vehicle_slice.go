package repository

import (
	"app/internal"
)

// NewVehicleSlice returns a new instance of a vehicle repository in an slice.
func NewVehicleSlice(db []internal.Vehicle, lastId int) *VehicleSlice {
	return &VehicleSlice{
		db:     db,
		lastId: lastId,
	}
}

// VehicleSlice is an struct that represents a vehicle repository in an slice.
type VehicleSlice struct {
	// db is the database of vehicles.
	db 	   []internal.Vehicle
	// lastId is the last id of the database.
	lastId int
}

// FindAll returns all vehicles
func (s *VehicleSlice) FindAll() (v []internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	// make a copy of the database
	v = make([]internal.Vehicle, len(s.db))
	copy(v, s.db)
	return
}

func (s *VehicleSlice) FindByMarcaAños(brand string, minYear int, maxYear int) (v []internal.Vehicle, err error){
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	var foundVehicles []internal.Vehicle

	for _, ve := range s.db {
		if ve.Attributes.Brand == brand && ve.Attributes.Year > minYear && ve.Attributes.Year < maxYear {
			foundVehicles = append(foundVehicles, ve)
		}
	}
	if len(foundVehicles) == 0 {
		return nil, internal.ErrRepositoryVehicleNotFound
	}
	return foundVehicles, nil
}