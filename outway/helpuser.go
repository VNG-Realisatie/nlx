// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package outway

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"go.nlx.io/nlx/common/httperrors"
	outway_http "go.nlx.io/nlx/outway/http"
	"go.nlx.io/nlx/outway/plugins"
)

const maxSuggestions = 10

// createList create max 10 unique sorted suggestions.
func createList(options []string) string {
	suggestion := ""
	unique := make(map[string]bool)

	sort.Strings(options)

	for i := range options {
		option := options[i]
		_, exists := unique[option]

		if !exists {
			unique[option] = true

			if suggestion == "" {
				suggestion = option
			} else {
				suggestion = suggestion + ", " + option
			}
		}
		// limit suggestions
		if i > maxSuggestions {
			break
		}
	}

	return suggestion
}

// HelpUserOrg suggest organizations for input
func (o *Outway) helpUserOrg(w http.ResponseWriter, serialNumber string) {
	suggestion := make([]string, 0)

	if serialNumber != "" {
		for k := range o.servicesDirectory {
			if strings.HasPrefix(k, serialNumber) {
				matchSerialNumber := strings.Split(k, ".")
				suggestion = append(suggestion, matchSerialNumber[0])
			}
		}
	}

	if len(suggestion) == 0 {
		// list all organizations.
		for k := range o.servicesDirectory {
			matchSerialNumber := strings.Split(k, ".")
			suggestion = append(suggestion, matchSerialNumber[0])
		}
	}

	msg := createList(suggestion)
	outway_http.WriteError(w, httperrors.C1, httperrors.InvalidURL(fmt.Sprintf("invalid /serialNumber/service/ url: valid organization serial numbers : [%s]", msg)))
}

func (o *Outway) helpUserService(
	w http.ResponseWriter, organizationSerialNumber, service string) {
	// check if organization exists.
	org := ""
	services := make([]string, 0)

	for k := range o.servicesDirectory {
		matchSerialNumber := strings.Split(k, ".")
		if matchSerialNumber[0] == organizationSerialNumber {
			org = matchSerialNumber[0]
			// store services
			services = append(services, matchSerialNumber[1])
		}
	}
	// if no matching organization suggest organizations
	if org == "" {
		o.helpUserOrg(w, organizationSerialNumber)
		return
	}

	// suggest service(s)
	serviceOptions := make([]string, 0)

	for i := range services {
		s := services[i]
		if strings.HasPrefix(s, service) {
			serviceOptions = append(serviceOptions, s)
		}
	}

	var msg string

	if len(serviceOptions) > 0 {
		// suggest matching services
		msg = createList(serviceOptions)
	} else {
		// suggest all services
		msg = createList(services)
	}

	outway_http.WriteError(w, httperrors.C1, httperrors.InvalidURL(fmt.Sprintf("invalid serialNumber/service path: valid services : [%s]", msg)))
}

func (o *Outway) helpUser(w http.ResponseWriter, msg string, dest *plugins.Destination, urlPath string) {
	// we did not get a complete 3 part url path. help user create one.
	if dest == nil {
		pathParts := strings.SplitN(strings.TrimPrefix(urlPath, "/"), "/", 3)

		if len(pathParts) == 0 {
			// suggest organization
			o.helpUserOrg(w, "")
			return
		}

		if len(pathParts) == 1 {
			// suggest organization
			o.helpUserOrg(w, pathParts[0])
			return
		}

		if len(pathParts) == 2 {
			// suggest service
			o.helpUserService(w, pathParts[0], pathParts[1])
			return
		}
	} else {
		// users gave a 'complete' path, but still failing
		// do suggestions
		o.helpUserService(w, dest.OrganizationSerialNumber, dest.Service)
		return
	}

	if urlPath == "" {
		outway_http.WriteError(w, httperrors.C1, httperrors.InvalidURL(fmt.Sprintf("missing urlpath for service %s", msg)))
		return
	}
}
