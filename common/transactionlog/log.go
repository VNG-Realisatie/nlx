// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

// Record encompasses the data stored in the transactionlog for a single recorded transaction.
type Record struct {
	SrcOrganization  string                 `json:"source_organization"`
	DestOrganization string                 `json:"destination_organization"`
	ServiceName      string                 `json:"service_name"`
	LogrecordID      string                 `json:"logrecord-id"`
	Data             map[string]interface{} `json:"data"`
	DataSubjects     map[string]string      `json:"dataSubjects`
}
