package tls_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.nlx.io/nlx/common/tls"
)

func TestValideSerialNumber(t *testing.T) {
	t.Parallel()

	testUUID, err := uuid.New().MarshalBinary()
	require.NoError(t, err)

	type args struct {
		serialNumber string
	}

	tests := map[string]struct {
		args    args
		wantErr error
	}{
		"when_empty": {
			args: args{
				serialNumber: "",
			},
			wantErr: tls.ErrSerialNumberEmpty,
		},
		"when_too_long_uuid": {
			args: args{
				serialNumber: "123e4567-e89b-12d3-a456-426614174000",
			},
			wantErr: tls.ErrSerialNumberTooLong,
		},
		"when_bytes_too_long_but_string_length_correct": {
			args: args{
				serialNumber: "tøō_lõng_whīlè_is_20",
			},
			wantErr: tls.ErrSerialNumberTooLong,
		},
		"when_bytes_too_long_but_string_length_correct_emojis": {
			args: args{
				serialNumber: "1️⃣2️⃣3️⃣4️⃣5️⃣6️⃣",
			},
			wantErr: tls.ErrSerialNumberTooLong,
		},
		"when_too_long_21": {
			args: args{
				serialNumber: "000000000000000000001",
			},
			wantErr: tls.ErrSerialNumberTooLong,
		},
		"happy_flow_20": {
			args: args{
				serialNumber: "00000000000000000001",
			},
			wantErr: nil,
		},
		"happy_flow_1": {
			args: args{
				serialNumber: "1",
			},
			wantErr: nil,
		},
		"happy_flow_uuid": {
			args: args{
				serialNumber: string(testUUID),
			},
			wantErr: nil,
		},
		"happy_flow_string": {
			args: args{
				serialNumber: "serialnumber:5",
			},
			wantErr: nil,
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := tls.ValidateSerialNumber(tt.args.serialNumber)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
