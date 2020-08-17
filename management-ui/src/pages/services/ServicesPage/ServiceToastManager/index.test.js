// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders, act } from '../../../../test-utils'
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
    history.push('/services/my-new-service?lastAction=added')
  })

  const alert = queryByRole('alert')

  expect(alert).toBeTruthy()
  expect(alert).toHaveTextContent('The service has been added')
  expect(history.location.pathname).toEqual('/services/my-new-service')
})

test('navigating to the new service when it has just been edited', async () => {
  const history = createMemoryHistory()
  const { queryByRole } = renderWithProviders(
    <Router history={history}>
      <ServiceAddedToastManager />
    </Router>,
  )

  expect(queryByRole('alert')).toBeNull()

  act(() => {
    history.push('/services/my-new-service?lastAction=edited')
  })

  const alert = queryByRole('alert')

  expect(alert).toBeTruthy()
  expect(alert).toHaveTextContent('The service has been updated')
  expect(history.location.pathname).toEqual('/services/my-new-service')
})

test('navigating to the new service when it has just been removed', async () => {
  const history = createMemoryHistory()
  const { queryByRole } = renderWithProviders(
    <Router history={history}>
      <ServiceAddedToastManager />
    </Router>,
  )

  expect(queryByRole('alert')).toBeNull()

  act(() => {
    history.push('/services/my-new-service?lastAction=removed')
  })

  const alert = queryByRole('alert')

  expect(alert).toBeTruthy()
  expect(alert).toHaveTextContent('The service has been removed')
  expect(history.location.pathname).toEqual('/services')
})
