// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory

type InwayState string

const (
	InwayStateUnknown InwayState = "UNKNOWN"
	InwayStateUp      InwayState = "UP"
	InwayStateDown    InwayState = "DOWN"
)

type Inway struct {
	Address string     `json:"address"`
	State   InwayState `json:"state"`
}
