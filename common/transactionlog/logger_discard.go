// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

// DiscardTransactionLogger discards records it gets
type DiscardTransactionLogger struct{}

// NewDiscardTransactionLogger creates a new TransactionLogger that discards all records.
func NewDiscardTransactionLogger() TransactionLogger {
	return &DiscardTransactionLogger{}
}

// AddRecord implements TransactionLogger.AddRecord and discards any given record.
func (txl *DiscardTransactionLogger) AddRecord(rec *Record) error {
	return nil
}

// Close implements TransactionLogger.Close.
func (txl *DiscardTransactionLogger) Close() error {
	return nil
}
