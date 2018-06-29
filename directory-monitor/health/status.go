// Package health is used by inway and healthceck to communicate
// health/status in a standardized json format.
package health

// Status is a structure used to indicate health of a service
// provided by an inway.
type Status struct {
	Healthy bool `json:"healthy"`
}
