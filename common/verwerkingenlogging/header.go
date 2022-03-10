package verwerkingenlogging

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	HeaderVertrouwelijkheid       string = "Vertrouwelijkheid"
	HeadersActieNaam              string = "Actienaam"
	HeaderBewaarTermijn           string = "Bewaartermijn"
	HeaderHandelingNaam           string = "Handelingnaam"
	HeaderVerwerkingNaam          string = "Verwerkingnaam"
	HeaderSysteem                 string = "Systeem"
	HeaderGebruiker               string = "Gebruiker"
	HeaderUitvoerder              string = "Uitvoerder"
	HeaderVerwerkingID            string = "Verwerkingid"
	HeaderVerwerkingsActiviteitID string = "Verwerkingsactiviteitid"
	HeaderVerwerkteObjecten       string = "Verwerkteobjecten"
)

var ErrHeadersDoNotContainVerwerkingenLoggingData = fmt.Errorf("no verwerkingenlogging data in headers")

func getHTTPHeaders() map[string]string {
	return map[string]string{
		HeadersActieNaam:              HeadersActieNaam,
		HeaderBewaarTermijn:           HeaderBewaarTermijn,
		HeaderHandelingNaam:           HeaderHandelingNaam,
		HeaderSysteem:                 HeaderSysteem,
		HeaderGebruiker:               HeaderGebruiker,
		HeaderUitvoerder:              HeaderUitvoerder,
		HeaderVertrouwelijkheid:       HeaderVertrouwelijkheid,
		HeaderVerwerkingID:            HeaderVerwerkingID,
		HeaderVerwerkingNaam:          HeaderVerwerkingNaam,
		HeaderVerwerkingsActiviteitID: HeaderVerwerkingsActiviteitID,
		HeaderVerwerkteObjecten:       HeaderVerwerkteObjecten,
	}
}

func extractHTTPHeaders(headers http.Header) http.Header {
	validHeaderMap := getHTTPHeaders()
	validVerwerkingenLoggingHeaders := http.Header{}

	for headerName, headerValues := range headers {
		if _, ok := validHeaderMap[headerName]; !ok {
			continue
		}

		if len(headerValues) > 0 {
			validVerwerkingenLoggingHeaders[headerName] = headerValues
		}
	}

	return validVerwerkingenLoggingHeaders
}

//nolint gocyclo: allot of headers
func BuildLogRequestFromHeaders(headers http.Header) (*WriteLogRequest, error) {
	headerData := extractHTTPHeaders(headers)

	if len(headerData) == 0 {
		return nil, ErrHeadersDoNotContainVerwerkingenLoggingData
	}

	request := &WriteLogRequest{}

	for header, values := range headerData {
		switch header {
		case HeadersActieNaam:
			request.ActieNaam = values[0]
		case HeaderBewaarTermijn:
			request.Bewaartermijn = values[0]
		case HeaderGebruiker:
			request.Gebruiker = values[0]
		case HeaderHandelingNaam:
			request.HandelingNaam = values[0]
		case HeaderSysteem:
			request.Systeem = values[0]
		case HeaderUitvoerder:
			request.Uitvoerder = values[0]
		case HeaderVertrouwelijkheid:
			request.Vertrouwelijkheid = values[0]
		case HeaderVerwerkingID:
			request.VerwerkingID = values[0]
		case HeaderVerwerkingNaam:
			request.VerwerkingNaam = values[0]
		case HeaderVerwerkingsActiviteitID:
			request.VerwerkingsactiviteitID = values[0]
		case HeaderVerwerkteObjecten:
			vewerkteObjectenJSONString := values[0]
			verwerktenObjecten := &[]*VerwerktObject{}
			err := json.Unmarshal([]byte(vewerkteObjectenJSONString), verwerktenObjecten)
			if err != nil {
				return nil, fmt.Errorf("unable to parse the verwerktenobject: %s", err)
			}

			request.VerwerkteObjecten = *verwerktenObjecten
		}
	}

	request.Tijdstip = time.Now()

	return request, nil
}
