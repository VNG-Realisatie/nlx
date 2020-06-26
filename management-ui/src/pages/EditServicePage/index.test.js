// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, fireEvent } from '@testing-library/react'
import { Route, StaticRouter, Router } from 'react-router-dom'

import { createMemoryHistory } from 'history'
import UserContext from '../../user-context'
import { renderWithProviders } from '../../test-utils'
import EditServicePage from './index'

jest.mock('../../components/ServiceForm', () => ({ onSubmitHandler }) => (
  <form onSubmit={() => onSubmitHandler({ foo: 'bar' })} data-testid="form">
    <button type="submit" />
  </form>
))

describe('the EditServicePage', () => {
  afterEach(() => {
    jest.resetModules()
  })

  it('before the service has been loaded', async () => {
    jest.useFakeTimers()
    const getServiceByName = jest.fn().mockResolvedValue({
      name: 'mock-service',
    })
    const userContext = { user: { id: '42' } }
    const { findByRole, getByLabelText } = renderWithProviders(
      <StaticRouter location="/services/mock-service/edit-service">
        <Route path="/services/:name/edit-service">
          <UserContext.Provider value={userContext}>
            <EditServicePage getServiceByName={getServiceByName} />
          </UserContext.Provider>
        </Route>
      </StaticRouter>,
    )

    expect(await findByRole('progressbar')).toBeTruthy()
    const linkBack = getByLabelText(/Back/)
    expect(linkBack.getAttribute('href')).toBe('/services/mock-service')
  })

  it('when fetching the services fails', async () => {
    const getServiceByName = jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error'))
    const userContext = { user: { id: '42' } }
    const { findByRole, queryByRole } = renderWithProviders(
      <StaticRouter>
        <UserContext.Provider value={userContext}>
          <EditServicePage getServiceByName={getServiceByName} />
        </UserContext.Provider>
      </StaticRouter>,
    )

    expect(await findByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe('Failed to load the service.')
  })

  it('after the service has been fetched', async () => {
    const getServiceByName = jest.fn().mockResolvedValue({
      name: 'mock-service',
    })
    const userContext = { user: { id: '42' } }
    const { findByTestId } = renderWithProviders(
      <StaticRouter>
        <UserContext.Provider value={userContext}>
          <EditServicePage getServiceByName={getServiceByName} />
        </UserContext.Provider>
      </StaticRouter>,
    )

    expect(await findByTestId('form')).toBeTruthy()
  })

  it('successfully submitting the form', async () => {
    const history = createMemoryHistory()
    const getServiceByNameSpy = jest.fn().mockResolvedValue({
      name: 'mock-service',
    })
    const updateHandler = jest.fn().mockResolvedValue({
      name: 'mock-service',
    })
    const { findByTestId } = renderWithProviders(
      <Router history={history}>
        <EditServicePage
          updateHandler={updateHandler}
          getServiceByName={getServiceByNameSpy}
        />
      </Router>,
    )

    const editServiceForm = await findByTestId('form')
    await act(async () => {
      fireEvent.submit(editServiceForm)
    })

    expect(history.location.pathname).toEqual('/services/mock-service')
    expect(history.location.search).toEqual('?edited=true')
  })

  it('re-submitting the form when the previous submission went wrong', async () => {
    const history = createMemoryHistory()
    const getServiceByNameSpy = jest.fn().mockResolvedValue({
      name: 'mock-service',
    })
    const updateHandler = jest
      .fn()
      .mockResolvedValue({
        name: 'mock-service',
      })
      .mockRejectedValueOnce(new Error('arbitrary error'))

    const { findByTestId, queryByRole } = renderWithProviders(
      <Router history={history}>
        <EditServicePage
          updateHandler={updateHandler}
          getServiceByName={getServiceByNameSpy}
        />
      </Router>,
    )

    const editServiceForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(editServiceForm)
    })

    expect(updateHandler).toHaveBeenCalledTimes(1)
    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe(
      'Failed to update the service.arbitrary error',
    )

    await act(async () => {
      await fireEvent.submit(editServiceForm)
    })

    expect(updateHandler).toHaveBeenCalledTimes(2)

    expect(history.location.pathname).toEqual('/services/mock-service')
    expect(history.location.search).toEqual('?edited=true')
  })

  it('submitting when the HTTP response is not ok', async () => {
    const getServiceByNameSpy = jest.fn().mockResolvedValue({
      name: 'mock-service',
    })
    const updateHandler = jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error'))

    const { findByTestId, queryByRole } = renderWithProviders(
      <StaticRouter>
        <EditServicePage
          updateHandler={updateHandler}
          getServiceByName={getServiceByNameSpy}
        />
      </StaticRouter>,
    )

    const editServiceForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(editServiceForm)
    })

    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe(
      'Failed to update the service.arbitrary error',
    )
  })
})
