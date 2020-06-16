// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package directory_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.nlx.io/nlx/management-api/pkg/directory"
)

var directoryServiceStatusTests = []struct {
	ExpectedStatus directory.DirectoryService_Status
	Inways         []*directory.Inway
}{
	{
		directory.DirectoryService_unknown,
		nil,
	},
	{
		directory.DirectoryService_unknown,
		[]*directory.Inway{
			{State: directory.InwayStateUnknown},
		},
	},
	{
		directory.DirectoryService_up,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
		},
	},
	{
		directory.DirectoryService_up,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
			{State: directory.InwayStateUp},
			{State: directory.InwayStateUp},
		},
	},
	{
		directory.DirectoryService_down,
		[]*directory.Inway{
			{State: directory.InwayStateDown},
		},
	},
	{
		directory.DirectoryService_down,
		[]*directory.Inway{
			{State: directory.InwayStateDown},
			{State: directory.InwayStateDown},
		},
	},
	{
		directory.DirectoryService_degraded,
		[]*directory.Inway{
			{State: directory.InwayStateUp},
			{State: directory.InwayStateDown},
		},
	},
	{
		directory.DirectoryService_degraded,
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
			status := directory.DetermineDirectoryServiceStatus(test.Inways)
			assert.Equal(t, test.ExpectedStatus, status)
		})
	}
}
