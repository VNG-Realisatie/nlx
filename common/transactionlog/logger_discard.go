// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

type DiscardTransactionLogger struct{}

func NewDiscardTransactionLogger() TransactionLogger {
	return &DiscardTransactionLogger{}
}

func (txl *DiscardTransactionLogger) AddRecord(rec *Record) error {
	return nil
}

func (txl *DiscardTransactionLogger) Close() error {
	return nil
}
