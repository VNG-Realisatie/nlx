// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, act, screen } from '@testing-library/react'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementServiceApi } from '../../../api'
import GeneralSettings from './index'

describe('the General settings section', () => {
  afterEach(() => {
    jest.clearAllMocks()
  })

  it('on initialization', async () => {
    const managementApiClient = new ManagementServiceApi()
    managementApiClient.managementServiceGetSettings = jest
      .fn()
      .mockResolvedValue({
        settings: { inway: { name: 'inway1' } },
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
    const managementApiClient = new ManagementServiceApi()

    managementApiClient.managementServiceGetSettings = jest
      .fn()
      .mockResolvedValue({
        settings: { organizationInway: 'inway1' },
      })

    managementApiClient.managementServiceUpdateSettings = jest.fn()

    const store = new RootStore({ managementApiClient })

    jest.spyOn(store.applicationStore, 'updateOrganizationInway')

    jest
      .spyOn(store.applicationStore, 'updateGeneralSettings')
      .mockRejectedValueOnce({ response: { status: 403 } })
      .mockResolvedValue({})

    renderWithAllProviders(
      <StoreProvider rootStore={store}>
        <GeneralSettings />
      </StoreProvider>,
    )

    const settingsForm = await screen.findByTestId('form')

    await act(async () => {
      fireEvent.submit(settingsForm)
    })

    expect(store.applicationStore.updateGeneralSettings).toHaveBeenCalledWith({
      organizationInway: 'inway1',
      organizationEmailAddress: '',
    })

    expect(
      await screen.findByText('Failed to update the settings'),
    ).toBeInTheDocument()

    await act(async () => {
      fireEvent.submit(settingsForm)
    })

    expect(
      await screen.findByText('Successfully updated the settings'),
    ).toBeInTheDocument()

    expect(store.applicationStore.updateOrganizationInway).toHaveBeenCalledWith(
      {
        isOrganizationInwaySet: true,
      },
    )
  })

  it('should re-submit the form when the previous submission went wrong', async () => {
    const managementApiClient = new ManagementServiceApi()
    managementApiClient.managementServiceGetSettings = jest
      .fn()
      .mockResolvedValue({ settings: { organizationInway: 'inway1' } })

    managementApiClient.managementServiceUpdateSettings = jest
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

    expect(
      managementApiClient.managementServiceUpdateSettings,
    ).toHaveBeenCalledTimes(2)
    expect(getAllByRole('alert')[1].textContent).toBe(
      'Successfully updated the settings',
    )
  })
})
