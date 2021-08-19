// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, act } from '@testing-library/react'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import GeneralSettings from './index'

describe('the General settings section', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })

  it('on initialization', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementGetSettings = jest.fn().mockResolvedValue({
      inway: { name: 'inway1' },
    })

    const store = new RootStore({ managementApiClient })

    const { findByTestId, queryByTestId } = renderWithAllProviders(
      <StoreProvider rootStore={store}>
        <GeneralSettings />
      </StoreProvider>,
    )

    const formElement = await findByTestId('form')

    expect(formElement).toBeTruthy()
    expect(queryByTestId('error-message')).toBeNull()
  })

  it('successfully submits the form', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementGetSettings = jest.fn().mockResolvedValue({
      organizationInway: 'inway1',
    })
    managementApiClient.managementUpdateSettings = jest.fn().mockResolvedValue()

    const store = new RootStore({ managementApiClient })

    jest
      .spyOn(store.applicationStore, 'updateOrganizationInway')
      .mockResolvedValue({
        isOrganizationInwaySet: true,
      })

    const { findByTestId, getByRole } = renderWithAllProviders(
      <StoreProvider rootStore={store}>
        <GeneralSettings />
      </StoreProvider>,
    )

    const settingsForm = await findByTestId('form')
    await act(async () => {
      fireEvent.submit(settingsForm)
    })

    expect(store.applicationStore.updateOrganizationInway).toHaveBeenCalledWith(
      {
        isOrganizationInwaySet: true,
      },
    )

    expect(getByRole('alert').textContent).toBe(
      'Successfully updated the settings',
    )
  })

  it('should re-submit the form when the previous submission went wrong', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementGetSettings = jest
      .fn()
      .mockResolvedValue({ organizationInway: 'inway1' })

    managementApiClient.managementUpdateSettings = jest
      .fn()
      .mockRejectedValueOnce(new Error('arbitrary error'))
      .mockResolvedValueOnce([])

    const store = new RootStore({ managementApiClient })

    const { findByTestId, getAllByRole } = renderWithAllProviders(
      <StoreProvider rootStore={store}>
        <GeneralSettings />
      </StoreProvider>,
    )

    const settingsForm = await findByTestId('form')
    await act(async () => {
      fireEvent.submit(settingsForm)
    })

    expect(getAllByRole('alert')[0]).toBeTruthy()
    expect(getAllByRole('alert')[0].textContent).toBe(
      'Failed to update the settings',
    )

    await act(async () => {
      fireEvent.submit(settingsForm)
    })

    expect(managementApiClient.managementUpdateSettings).toHaveBeenCalledTimes(
      2,
    )
    expect(getAllByRole('alert')[1].textContent).toBe(
      'Successfully updated the settings',
    )
  })
})
