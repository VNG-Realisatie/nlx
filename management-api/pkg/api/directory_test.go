// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package api

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/pkg/directory"
)

var directoryServiceStatusTests = []struct {
	ExpectedStatus DirectoryServiceStatus
	Inways         []*directory.Inway
}{
	{
		DirectoryServiceStatusUnknown,
		nil,
	},
	{
		DirectoryServiceStatusUnknown,
		[]*directory.Inway{
			{State: directory.InwayStateUnknown},
		},
	},
	{
		DirectoryServiceStatusUp,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
		},
	},
	{
		DirectoryServiceStatusUp,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
			{State: directory.InwayStateUp},
			{State: directory.InwayStateUp},
		},
	},
	{
		DirectoryServiceStatusDown,
		[]*directory.Inway{
			{State: directory.InwayStateDown},
		},
	},
	{
		DirectoryServiceStatusDown,
		[]*directory.Inway{
			{State: directory.InwayStateDown},
			{State: directory.InwayStateDown},
		},
	},
	{
		DirectoryServiceStatusDegraded,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
			{State: directory.InwayStateDown},
		},
	},
	{
		DirectoryServiceStatusDegraded,
		[]*directory.Inway{
			{State: directory.InwayStateDown},
			{State: directory.InwayStateUnknown},
		},
	},
}

func TestDirectoryServiceStatus(t *testing.T) {
	for i, test := range directoryServiceStatusTests {
		name := strconv.Itoa(i + 1)
		test := test

		t.Run(name, func(t *testing.T) {
			status := DetermineDirectoryServiceStatus(test.Inways)
			assert.Equal(t, test.ExpectedStatus, status)
		})
	}
}
