// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../../../../../../../test-utils'
import AccessDetails from './index'

test('Terminate Access Details', () => {
  const accessRequest = {
    serviceName: 'service',
    organization: {
      serialNumber: '00000000000000000001',
      name: 'organization',
    },
    publicKeyFingerprint: 'public-key-fingerprint',
  }

  const { container } = renderWithProviders(
    <AccessDetails
      subTitle="test"
      organization={accessRequest.organization}
      serviceName={accessRequest.serviceName}
      publicKeyFingerprint={accessRequest.publicKeyFingerprint}
    />,
  )

  expect(container).toHaveTextContent(/Organization/)
  expect(container).toHaveTextContent(/Public Key Fingerprint/)
  expect(container).toHaveTextContent(/Service/)
})
