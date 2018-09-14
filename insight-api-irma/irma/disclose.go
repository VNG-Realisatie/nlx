package irma

// DiscloseRequest contains the data for a disclose request
type DiscloseRequest struct {
	Content []DiscloseRequestContent `json:"content"`
}

// DiscloseRequestContent contains information about a required attribute(set) in a disclose request
type DiscloseRequestContent struct {
	Label      string      `json:"label"`
	Attributes []Attribute `json:"attributes"`
}
