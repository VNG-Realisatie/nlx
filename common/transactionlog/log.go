package transactionlog

// Record encompases the data stored in the transactionlog for a single recorded transaction.
type Record struct {
	SrcOrganization  string
	DestOrganization string
	ServiceName      string
	RequestPath      string
	Data             interface{}
}
