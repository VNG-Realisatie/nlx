// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { configure } from 'mobx'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../test-utils'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'
import {
  SERVICE_STATE_DEGRADED,
  SERVICE_STATE_UP,
} from '../../../../../components/StateIndicator'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../../../../stores/models/OutgoingAccessRequestModel'
import { RootStore } from '../../../../../stores'
import AccessProofModel from '../../../../../stores/models/AccessProofModel'
import DirectoryServiceRow from './index'

test('display service information', async () => {
  configure({ safeDescriptors: false })

  const rootStore = new RootStore()

  const service = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    serviceData: {
      id: 'my-service',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Test Organization',
      },
      serviceName: 'Test Service',
      state: SERVICE_STATE_DEGRADED,
      apiSpecificationType: 'API',
    },
  })

  const { container } = renderWithProviders(
    <MemoryRouter>
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>
    </MemoryRouter>,
  )

  expect(container).toHaveTextContent('Test Organization')
  expect(container).toHaveTextContent('Test Service')
  expect(container).toHaveTextContent('state-degraded.svg')
  expect(container).toHaveTextContent('API')

  service.update({
    serviceData: {
      state: SERVICE_STATE_UP,
    },
    accessStates: [
      {
        accessRequest: new OutgoingAccessRequestModel({
          accessRequestData: {
            publicKeyFingerprint: 'public-key-fingerprint',
            state: ACCESS_REQUEST_STATES.FAILED,
            errorDetails: {
              cause: 'cause of failed access request',
            },
          },
          outgoingAccessRequestStore: null,
        }),
      },
    ],
  })

  expect(
    await screen.findByText('Request could not be sent'),
  ).toBeInTheDocument()
  expect(screen.getByTestId('directory-service-row')).toHaveTextContent(
    'state-up.svg',
  )
})

test('display warning if there is a sync error', async () => {
  const rootStore = new RootStore()

  const service = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    serviceData: {
      id: 'my-service',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Test Organization',
      },
      serviceName: 'Test Service',
      state: SERVICE_STATE_DEGRADED,
      apiSpecificationType: 'API',
    },
    accessStates: [
      {
        accessRequest: new OutgoingAccessRequestModel({
          outgoingAccessRequestStore: {},
          accessRequestData: {
            state: ACCESS_REQUEST_STATES.APPROVED,
            publicKeyFingerprint: 'public-key-fingerprint',
          },
        }),
        accessProof: new AccessProofModel({
          accessProofData: {
            id: '42',
          },
        }),
      },
    ],
  })

  renderWithProviders(
    <MemoryRouter>
      <table>
        <tbody>
          <DirectoryServiceRow service={service} />
        </tbody>
      </table>
    </MemoryRouter>,
  )

  rootStore.outgoingAccessRequestSyncErrorStore.loadFromSyncResponse(
    '00000000000000000001',
    'Test Service',
    {
      message: 'service_provider_no_organization_inway_specified',
    },
  )

  expect(
    await screen.findByTitle(
      'The organization has not specified an organization Inway. We are unable to retrieve the current state of your access requests.',
    ),
  ).toBeInTheDocument()
})
