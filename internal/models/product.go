package models

type Product struct {
	ID              int
	Name            string
	BasicRack       string
	AdditionalRacks []string
}
