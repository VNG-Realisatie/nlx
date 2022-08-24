// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package transactionlog

type TransactionLogger interface {
	AddRecord(rec *Record) error
	Close() error
}
