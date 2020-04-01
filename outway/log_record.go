// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
package outway

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"go.nlx.io/nlx/common/transactionlog"
)

const (
	logRecordSize        = 16
	logRecordStringSizes = logRecordSize * 2
)

type LogRecordID [logRecordSize]byte

func (l *LogRecordID) String() string {
	b := make([]byte, logRecordStringSizes)
	hex.Encode(b, l[:])

	return string(b)
}

func NewLogRecordID() (*LogRecordID, error) {
	var l LogRecordID

	if _, err := io.ReadFull(rand.Reader, l[:]); err != nil {
		return nil, fmt.Errorf("failed to read random bytes: %v", err)
	}

	return &l, nil
}

func createRecordData(h http.Header, p string) map[string]interface{} {
	recordData := make(map[string]interface{})
	recordData["request-path"] = p

	if processID := h.Get("X-NLX-Request-Process-Id"); processID != "" {
		recordData["doelbinding-process-id"] = processID
	}

	if dataElements := h.Get("X-NLX-Request-Data-Elements"); dataElements != "" {
		recordData["doelbinding-data-elements"] = dataElements
	}

	if userData := h.Get("X-NLX-Requester-User"); userData != "" {
		recordData["doelbinding-user"] = userData
	}

	if claims := h.Get("X-NLX-Requester-Claims"); claims != "" {
		recordData["doelbinding-claims"] = claims
	}

	if userID := h.Get("X-NLX-Request-User-Id"); userID != "" {
		recordData["doelbinding-user-id"] = userID
	}

	if applicationID := h.Get("X-NLX-Request-Application-Id"); applicationID != "" {
		recordData["doelbinding-application-id"] = applicationID
	}

	if subjectIdentifier := h.Get("X-NLX-Request-Subject-Identifier"); subjectIdentifier != "" {
		recordData["doelbinding-subject-identifier"] = subjectIdentifier
	}

	return recordData
}

func addLogRecordIDToRequest(r *http.Request, logRecordID *LogRecordID) {
	r.Header.Set("X-NLX-Logrecord-Id", logRecordID.String())
}

func (o *Outway) createLogRecord(r *http.Request, destination *destination) (*LogRecordID, error) {
	logRecordID, err := NewLogRecordID()
	if err != nil {
		return nil, errors.Wrap(err, "could not get new request ID")
	}

	dataSubjects, err := parseDataSubjects(r)
	if err != nil {
		return nil, errors.Wrap(err, "invalid data subject header")
	}

	recordData := createRecordData(r.Header, destination.Path)

	err = o.txlogger.AddRecord(&transactionlog.Record{
		SrcOrganization:  o.organizationName,
		DestOrganization: destination.Organization,
		ServiceName:      destination.Service,
		LogrecordID:      logRecordID.String(),
		Data:             recordData,
		DataSubjects:     dataSubjects,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to add record to database")
	}

	return logRecordID, nil
}

func parseDataSubjects(r *http.Request) (map[string]string, error) {
	if dataSubjectsHeader := r.Header.Get("X-NLX-Request-Data-Subject"); dataSubjectsHeader != "" {
		return transactionlog.ParseDataSubjectHeader(dataSubjectsHeader)
	}

	return map[string]string{}, nil
}
