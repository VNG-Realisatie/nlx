// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package convert

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func SQLToProtoTimestamp(in sql.NullTime) *timestamppb.Timestamp {
	if in.Valid {
		return timestamppb.New(in.Time)
	}

	return nil
}

func ProtoToSQLTimestamp(in *timestamppb.Timestamp) sql.NullTime {
	out := sql.NullTime{}
	if in != nil {
		out.Time = in.AsTime()
		out.Valid = true
	}

	return out
}
