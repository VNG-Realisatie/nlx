// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, act } from '@testing-library/react'
import { MemoryRouter as Router } from 'react-router-dom'

import UserContext from '../../user-context'
import { renderWithProviders } from '../../test-utils'
import AddServicePage from './index'

jest.mock('./AddServiceForm', () => ({ onSubmitHandler }) => (
  <form onSubmit={() => onSubmitHandler({ foo: 'bar' })} data-testid="form">
    <button type="submit" />
  </form>
))

describe('the AddServicePage', () => {
  afterEach(() => {
    jest.resetModules()
  })

  it('on initialization', () => {
    const userContext = { user: { id: '42' } }
    const { getByTestId, queryByRole } = renderWithProviders(
      <Router>
        <UserContext.Provider value={userContext}>
          <AddServicePage createHandler={() => {}} />
        </UserContext.Provider>
      </Router>,
    )

    expect(getByTestId('form')).toBeTruthy()
    expect(queryByRole('dialog')).toBeNull()
  })

  it('successfully submitting the form', async () => {
    const createHandler = jest.fn().mockResolvedValue()
    const { findByTestId, queryByTestId, queryByRole } = renderWithProviders(
      <Router>
        <AddServicePage createHandler={createHandler} />
      </Router>,
    )

    const addComponentForm = await findByTestId('form')
    await act(async () => {
      fireEvent.submit(addComponentForm)
    })

    expect(queryByTestId('form')).toBeNull()

    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe('De service is toegevoegd.')
  })

  it('re-submitting the form when the previous submission went wrong', async () => {
    const createHandler = jest
      .fn()
      .mockResolvedValue({})
      .mockRejectedValueOnce(new Error('arbitrary error'))

    const { findByTestId, queryByRole } = renderWithProviders(
      <Router>
        <AddServicePage createHandler={createHandler} />
      </Router>,
    )

    const addComponentForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(addComponentForm)
    })

    expect(createHandler).toHaveBeenCalledTimes(1)
    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe(
      'Toevoegen misluktarbitrary error',
    )

    await act(async () => {
      await fireEvent.submit(addComponentForm)
    })

    expect(await queryByRole('alert')).toBeTruthy()

    expect(createHandler).toHaveBeenCalledTimes(2)
    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe('De service is toegevoegd.')
  })

  it('submitting when the HTTP response is not ok', async () => {
    const createHandler = jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error'))

    const { findByTestId, queryByRole } = renderWithProviders(
      <Router>
        <AddServicePage createHandler={createHandler} />
      </Router>,
    )

    const addComponentForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(addComponentForm)
    })

    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe(
      'Toevoegen misluktarbitrary error',
    )
  })
})
