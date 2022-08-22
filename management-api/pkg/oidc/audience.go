// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package oidc

import "encoding/json"

type audience []string

func (a *audience) UnmarshalJSON(b []byte) error {
	var s string
	if json.Unmarshal(b, &s) == nil {
		*a = audience{s}
		return nil
	}

	var auds []string

	if err := json.Unmarshal(b, &auds); err != nil {
		return err
	}

	*a = auds

	return nil
}
