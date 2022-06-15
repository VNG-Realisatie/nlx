// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom'
import { fireEvent, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../../../../../../test-utils'
import AccessProofModel from '../../../../../../../stores/models/AccessProofModel'
import AccessProofs from './index'

const createComponent = (accessProofs) => {
  const accessProofModels = accessProofs.map((accessProof) => {
    return new AccessProofModel({
      accessProofData: accessProof,
    })
  })

  const history = createMemoryHistory({
    initialEntries: ['/'],
  })

  renderWithProviders(
    <HistoryRouter history={history}>
      <AccessProofs accessProofs={accessProofModels} />
    </HistoryRouter>,
  )

  return {
    history,
  }
}

test('no access proofs available', async () => {
  createComponent([])

  fireEvent.click(screen.getByText('Requestable services'))

  expect(
    await screen.findByText('No services have been connected'),
  ).toBeInTheDocument()
})

test('listing the access proofs', async () => {
  const { history } = createComponent([
    {
      serviceName: 'My Service',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'Organization One',
      },
      publicKeyFingerprint: 'public-key-fingerprint',
    },
  ])

  fireEvent.click(screen.getByText('Requestable services'))

  const service = await screen.findByText('My Service')
  expect(service).toBeInTheDocument()
  expect(screen.getByText('Organization One')).toBeInTheDocument()
  expect(screen.getByText('public-key-fingerprint')).toBeInTheDocument()

  fireEvent.click(service)

  expect(history.location.pathname).toEqual(
    '/directory/00000000000000000001/My Service',
  )
})
