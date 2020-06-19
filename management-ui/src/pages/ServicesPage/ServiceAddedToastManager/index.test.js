// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders, act } from '../../../test-utils'
import ServiceAddedToastManager from './index'

test('navigating to the new service when it has just been added', async () => {
  const history = createMemoryHistory()

  const { queryByRole } = renderWithProviders(
    <Router history={history}>
      <ServiceAddedToastManager />
    </Router>,
  )

  expect(queryByRole('alert')).toBeNull()

  act(() => {
    history.push('/services/my-new-service?new=true')
  })

  expect(queryByRole('alert')).toBeTruthy()
  expect(history.location.pathname).toEqual('/services/my-new-service')
})
