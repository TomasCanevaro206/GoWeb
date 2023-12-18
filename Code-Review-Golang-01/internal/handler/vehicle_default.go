package handler

import (
	"app/internal"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// VehicleJSON is an struct that represents a vehicle in json format.
type VehicleJSON struct {
	ID				int		`json:"id"`
	Brand			string	`json:"brand"`
	Model			string	`json:"model"`
	Registration	string	`json:"registration"`
	Year			int		`json:"year"`
	Color			string	`json:"color"`
	MaxSpeed		int		`json:"max_speed"`
	FuelType		string	`json:"fuel_type"`
	Transmission	string	`json:"transmission"`
	Passengers 		int		`json:"passengers"`
	Height			float64	`json:"height"`
	Width			float64	`json:"width"`
	Weight			float64	`json:"weight"`
}

// NewVehicleDefault returns a new instance of a vehicle handler.
func NewVehicleDefault(sv internal.ServiceVehicle) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is an struct that contains handlers for vehicle.
type VehicleDefault struct {
	sv internal.ServiceVehicle
}

func (c *VehicleDefault) GetByMarcaAños(brand string, minYear int, maxYear int) (gin.HandlerFunc) {
	return func(ctx *gin.Context) {

		br := ctx.Param("brand")
		//in := ctx.Param("start_year")
		//max := ctx.Param("end_year")

		vehicles, err := c.sv.FindByMarcaAños(br, 0, 9999)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID: vehicle.ID,
				Brand: vehicle.Attributes.Brand,
				Model: vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year: vehicle.Attributes.Year,
				Color: vehicle.Attributes.Color,
				MaxSpeed: vehicle.Attributes.MaxSpeed,
				FuelType: vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers: vehicle.Attributes.Passengers,
				Height: vehicle.Attributes.Height,
				Width: vehicle.Attributes.Width,
				Weight: vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) GetAll() (gin.HandlerFunc) {
	return func(ctx *gin.Context) {
		// request
		// ...
		
		// process
		// - get all vehicles from the service
		vehicles, err := c.sv.FindAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID: vehicle.ID,
				Brand: vehicle.Attributes.Brand,
				Model: vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year: vehicle.Attributes.Year,
				Color: vehicle.Attributes.Color,
				MaxSpeed: vehicle.Attributes.MaxSpeed,
				FuelType: vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers: vehicle.Attributes.Passengers,
				Height: vehicle.Attributes.Height,
				Width: vehicle.Attributes.Width,
				Weight: vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}
