package convert

import (
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSQLToProtoTimestamp(t *testing.T) {
	timestamp := time.Now()

	tests := map[string]struct {
		input sql.NullTime
		want  *timestamppb.Timestamp
	}{
		"happy_flow_valid_timestamp": {
			input: sql.NullTime{
				Valid: true,
				Time:  timestamp,
			},
			want: timestamppb.New(timestamp),
		},
		"happy_flow_nil_timestamp": {
			input: sql.NullTime{},
			want:  nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			got := SQLToProtoTimestamp(tt.input)

			assert.Equal(t, got, tt.want)
		})
	}
}

func TestProtoToSQLTimestamp(t *testing.T) {
	timestamp := time.Now().UTC()

	tests := map[string]struct {
		input *timestamppb.Timestamp
		want  sql.NullTime
	}{
		"happy_flow_valid_timestamp": {
			input: timestamppb.New(timestamp),
			want: sql.NullTime{
				Valid: true,
				Time:  timestamp,
			},
		},
		"happy_flow_nil_timestamp": {
			input: nil,
			want:  sql.NullTime{},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			got := ProtoToSQLTimestamp(tt.input)

			assert.Equal(t, got, tt.want)
		})
	}
}
