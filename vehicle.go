package ves

import (
	"github.com/jfoster/go-ves/internal/date"
)

type Vehicle struct {
	RegistrationNumber           string         `json:"registrationNumber"`
	TaxStatus                    TaxStatus      `json:"taxStatus,omitempty"`
	TaxDueDate                   *date.YYYYMMDD `json:"taxDueDate,omitempty"`
	ArtEndDate                   *date.YYYYMMDD `json:"artEndDate,omitempty"`
	MotStatus                    MOTStatus      `json:"motStatus,omitempty"`
	MotExpiryDate                *date.YYYYMMDD `json:"motExpiryDate,omitempty"`
	Make                         string         `json:"make,omitempty"`
	MonthOfFirstDvlaRegistration *date.YYYYMM   `json:"monthOfFirstDvlaRegistration,omitempty"`
	MonthOfFirstRegistration     *date.YYYYMM   `json:"monthOfFirstRegistration,omitempty"`
	YearOfManufacture            date.YYYY      `json:"yearOfManufacture,omitempty"`
	EngineCapacity               int32          `json:"engineCapacity,omitempty"`
	Co2Emissions                 int32          `json:"co2Emissions,omitempty"`
	FuelType                     FuelType       `json:"fuelType,omitempty"`
	MarkedForExport              bool           `json:"markedForExport,omitempty"`
	Colour                       string         `json:"colour,omitempty"`
	TypeApproval                 string         `json:"typeApproval,omitempty"`
	Wheelplan                    string         `json:"wheelplan,omitempty"`
	RevenueWeight                int32          `json:"revenueWeight,omitempty"`
	RealDrivingEmissions         string         `json:"realDrivingEmissions,omitempty"`
	DateOfLastV5CIssued          *date.YYYYMMDD `json:"dateOfLastV5CIssued,omitempty"`
	EuroStatus                   EuroStatus     `json:"euroStatus,omitempty"`
}

type TaxStatus string

const (
	TaxStatusNotTaxedFordRoadUse TaxStatus = "Not Taxed for on Road Use"
	TaxStatusSORN                TaxStatus = "SORN"
	TaxStatusTaxed               TaxStatus = "Taxed"
	TaxStatusUntaxed             TaxStatus = "Untaxed"
)

type MOTStatus string

const (
	MOTStatusNoStatus  MOTStatus = "No details held by DVLA"
	MOTStatusNoResults MOTStatus = "No results returned"
	MOTStatusNotValid  MOTStatus = "Not valid"
	MOTStatusValid     MOTStatus = "Valid"
)

type EuroStatus string

const (
	EuroStatus6 EuroStatus = "Euro 6"
	EuroStatus5 EuroStatus = "Euro 5"
	EuroStatus4 EuroStatus = "Euro 4"
	EuroStatus3 EuroStatus = "Euro 3"
	EuroStatus2 EuroStatus = "Euro 2"
	EuroStatus1 EuroStatus = "Euro 1"
)

type FuelType string

const (
	FuelTypePetrol   FuelType = "PETROL"
	FuelTypeDiesel   FuelType = "DIESEL"
	FuelTypeElectric FuelType = "ELECTRICITY"
)

type VehicleRequest struct {
	RegistrationNumber string `json:"registrationNumber"`
}
