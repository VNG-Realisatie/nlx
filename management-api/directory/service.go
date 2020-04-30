// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory

type Service struct {
	Name                 string   `json:"service_name"`
	OrganizationName     string   `json:"organization_name"`
	APISpecificationType string   `json:"api_specification_type"`
	Inways               []*Inway `json:"inways"`
}

type servicesRoot map[string][]*Service

func (d *Client) ListServices() ([]*Service, error) {
	req, err := d.newRequest("GET", "list-services")
	if err != nil {
		return nil, err
	}

	var services servicesRoot

	_, err = d.sendRequest(req, &services)
	if err != nil {
		return nil, err
	}

	for _, s := range services {
		return s, nil
	}

	return nil, nil
}
