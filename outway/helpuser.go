package outway

import (
	"net/http"
	"sort"
	"strings"
)

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
		if i > 10 {
			break
		}
	}
	return suggestion
}

// HelpUserOrg suggest organizations for input
func (o *Outway) helpUserOrg(w http.ResponseWriter, organization string) {
	suggestion := make([]string, 0)
	if organization != "" {
		for k := range o.servicesDirectory {
			if strings.HasPrefix(k, organization) {
				matchorg := strings.Split(k, ".")
				suggestion = append(suggestion, matchorg[0])
			}
		}
	}
	if len(suggestion) == 0 {
		// list all organizations.
		for k := range o.servicesDirectory {
			matchorg := strings.Split(k, ".")
			suggestion = append(suggestion, matchorg[0])
		}
	}

	msg := createList(suggestion)
	http.Error(
		w,
		"nlx outway: invalid /organization/service/ url: valid organizations : ["+msg+"]",
		http.StatusBadRequest)
}

func (o *Outway) helpUserService(
	w http.ResponseWriter, organization, service string) {
	// check if organization exists.
	org := ""
	services := make([]string, 0)
	for k := range o.servicesDirectory {
		matchorg := strings.Split(k, ".")
		if matchorg[0] == organization {
			org = matchorg[0]
			// store services
			services = append(services, matchorg[1])
		}
	}
	// if no matching organization suggest organizations
	if org == "" {
		o.helpUserOrg(w, organization)
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
	msg := ""
	if len(serviceOptions) > 0 {
		// suggest matching services
		msg = createList(serviceOptions)
	} else {
		// suggest all services
		msg = createList(services)
	}
	http.Error(
		w,
		"nlx outway: invalid organization/service path: valid services : ["+msg+"]",
		http.StatusBadRequest)
}

func (o *Outway) helpUser(w http.ResponseWriter, msg string, dest *destination, urlPath string) {

	if urlPath == "" {
		http.Error(w, "nlx outway: missing urlpath"+msg, http.StatusBadRequest)
	}

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
		// users gave a 'complete' path, but still failing
		// do suggestions
	} else {
		o.helpUserService(w, dest.Organization, dest.Service)
	}
}
