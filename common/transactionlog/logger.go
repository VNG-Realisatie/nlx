// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

// TransactionLogger abstracts the writing of transactionlogs.
type TransactionLogger interface {
	AddRecord(rec *Record) error
	Close() error
}
