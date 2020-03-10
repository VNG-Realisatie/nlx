package outway

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
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
