package verwerkingenlogging

import "time"

type GetTokenRequest struct {
	Scopes []string `json:"scopes"`
}

type GetTokenResponse struct {
	Refresh string `json:"refresh"`
	Access  string `json:"access"`
}

type WriteLogRequest struct {
	ActieNaam                       string            `json:"actieNaam,omitempty"`
	HandelingNaam                   string            `json:"handelingNaam,omitempty"`
	VerwerkingNaam                  string            `json:"verwerkingNaam,omitempty"`
	VerwerkingID                    string            `json:"verwerkingId,omitempty"`
	VerwerkingsactiviteitID         string            `json:"verwerkingsactiviteitId,omitempty"`
	VerwerkingsactiviteitURL        string            `json:"verwerkingsactiviteitUrl,omitempty"`
	Vertrouwelijkheid               string            `json:"vertrouwelijkheid,omitempty"`
	Bewaartermijn                   string            `json:"bewaartermijn,omitempty"`
	Uitvoerder                      string            `json:"uitvoerder,omitempty"`
	Systeem                         string            `json:"systeem,omitempty"`
	Gebruiker                       string            `json:"gebruiker,omitempty"`
	Gegevensbron                    string            `json:"gegevensbron,omitempty"`
	SoortAfnemerID                  string            `json:"soortAfnemerId,omitempty"`
	AfnemerID                       string            `json:"afnemerId,omitempty"`
	VerwerkingsactiviteitIDAfnemer  string            `json:"verwerkingsactiviteitIdAfnemer,omitempty"`
	VerwerkingsactiviteitURLAfnemer string            `json:"verwerkingsactiviteitUrlAfnemer,omitempty"`
	VerwerkingIDAfnemer             string            `json:"verwerkingIdAfnemer,omitempty"`
	Tijdstip                        time.Time         `json:"tijdstip,omitempty"`
	VerwerkteObjecten               []*VerwerktObject `json:"verwerkteObjecten,omitempty"`
}

type VerwerktObject struct {
	Objecttype               string                     `json:"objecttype"`
	SoortObjectID            string                     `json:"soortObjectId"`
	ObjectID                 string                     `json:"objectId"`
	Betrokkenheid            string                     `json:"betrokkenheid"`
	VerwerkteSoortenGegevens []*VerwerktSoortGegegevens `json:"verwerkteSoortenGegevens"`
}

type VerwerktSoortGegegevens struct {
	SoortGegeven string `json:"soortGegeven"`
}
