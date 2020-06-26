// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { act, renderWithProviders } from '../../../test-utils'
import ServiceEditedToastManager from './index'

test('navigating to the new service when it has just been edited', async () => {
  const history = createMemoryHistory()

  const { queryByRole } = renderWithProviders(
    <Router history={history}>
      <ServiceEditedToastManager />
    </Router>,
  )

  expect(queryByRole('alert')).toBeNull()

  act(() => {
    history.push('/services/my-new-service?edited=true')
  })

  expect(queryByRole('alert')).toBeTruthy()
  expect(history.location.pathname).toEqual('/services/my-new-service')
})
