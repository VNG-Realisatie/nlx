// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithAllProviders, act } from '../../../../test-utils'
import ServiceAddedToastManager from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

afterEach(() => {
  jest.runOnlyPendingTimers()
  jest.useRealTimers()
})

test('navigating to the new service when it has just been added', async () => {
  const history = createMemoryHistory()
  const { queryByRole } = renderWithAllProviders(
    <Router history={history}>
      <ServiceAddedToastManager />
    </Router>,
  )

  expect(queryByRole('alert')).toBeNull()

  act(() => {
    history.push('/services/my-new-service?lastAction=added')
    jest.advanceTimersByTime(250)
  })

  const alert = queryByRole('alert')

  expect(alert).toBeTruthy()
  expect(alert).toHaveTextContent('The service has been added')
  expect(history.location.pathname).toEqual('/services/my-new-service')
})

test('navigating to the new service when it has just been edited', async () => {
  const history = createMemoryHistory()
  const { queryByRole } = renderWithAllProviders(
    <Router history={history}>
      <ServiceAddedToastManager />
    </Router>,
  )

  expect(queryByRole('alert')).toBeNull()

  act(() => {
    history.push('/services/my-new-service?lastAction=edited')
    jest.advanceTimersByTime(250)
  })

  const alert = queryByRole('alert')

  expect(alert).toBeTruthy()
  expect(alert).toHaveTextContent('The service has been updated')
  expect(history.location.pathname).toEqual('/services/my-new-service')
})

test('navigating to the new service when it has just been removed', async () => {
  const history = createMemoryHistory()
  const { queryByRole } = renderWithAllProviders(
    <Router history={history}>
      <ServiceAddedToastManager />
    </Router>,
  )

  expect(queryByRole('alert')).toBeNull()

  act(() => {
    history.push('/services/my-new-service?lastAction=removed')
    jest.advanceTimersByTime(250)
  })

  const alert = queryByRole('alert')

  expect(alert).toBeTruthy()
  expect(alert).toHaveTextContent('The service has been removed')
  expect(history.location.pathname).toEqual('/services')
})
