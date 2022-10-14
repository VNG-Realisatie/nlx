// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//

import React from 'react'
import { act, render, screen } from '@testing-library/react'
import { renderWithProviders } from '../test-utils'
import { RootStore, StoreProvider } from '../stores'
import { DirectoryServiceApi, ManagementServiceApi } from '../api'
import { ToSContextProvider } from './index'
import ToSContext from './index'

describe('ToSContext', () => {
  describe('Provider', () => {
    it('should fetch the current terms of service', async () => {
      const directoryApiClient = new DirectoryServiceApi()

      directoryApiClient.directoryServiceGetTermsOfService = jest
        .fn()
        .mockResolvedValue({
          url: '',
          enabled: false,
        })

      const store = new RootStore({
        directoryApiClient,
      })

      await act(async () =>
        renderWithProviders(
          <StoreProvider rootStore={store}>
            <ToSContextProvider />
          </StoreProvider>,
        ),
      )
      expect(
        directoryApiClient.directoryServiceGetTermsOfService,
      ).toHaveBeenCalledTimes(1)
    })

    it('should fetch the terms of service status if enabled', async () => {
      const directoryApiClient = new DirectoryServiceApi()

      directoryApiClient.directoryServiceGetTermsOfService = jest
        .fn()
        .mockResolvedValue({
          url: 'https://example.com',
          enabled: true,
        })

      const managementApiClient = new ManagementServiceApi()

      managementApiClient.managementServiceGetTermsOfServiceStatus = jest
        .fn()
        .mockResolvedValue({
          accepted: true,
        })

      const store = new RootStore({
        directoryApiClient,
        managementApiClient,
      })

      await act(async () =>
        renderWithProviders(
          <StoreProvider rootStore={store}>
            <ToSContextProvider />
          </StoreProvider>,
        ),
      )
      expect(
        directoryApiClient.directoryServiceGetTermsOfService,
      ).toHaveBeenCalledTimes(1)
      expect(
        managementApiClient.managementServiceGetTermsOfServiceStatus,
      ).toHaveBeenCalledTimes(1)
    })

    it('should disable ToS if directory is not available', async () => {
      const directoryApiClient = new DirectoryServiceApi()

      directoryApiClient.directoryServiceGetTermsOfService = jest
        .fn()
        .mockRejectedValue(new Error('arbitrary error'))

      const managementApiClient = new ManagementServiceApi()

      managementApiClient.managementServiceGetTermsOfServiceStatus = jest.fn()

      const store = new RootStore({
        directoryApiClient,
        managementApiClient,
      })

      await act(async () =>
        renderWithProviders(
          <StoreProvider rootStore={store}>
            <ToSContextProvider>
              <ToSContext.Consumer>
                {({ tos }) => (
                  <div data-testid="tos-enabled">
                    {tos ? JSON.stringify(tos.enabled) : ''}
                  </div>
                )}
              </ToSContext.Consumer>
            </ToSContextProvider>
          </StoreProvider>,
        ),
      )
      expect(
        directoryApiClient.directoryServiceGetTermsOfService,
      ).toHaveBeenCalledTimes(1)
      expect(
        managementApiClient.managementServiceGetTermsOfServiceStatus,
      ).toHaveBeenCalledTimes(0)

      expect(screen.getByTestId('tos-enabled')).toHaveTextContent('false')
    })

    it('should make the ToS available to the context consumers', () => {
      render(
        <ToSContextProvider tos={{ url: 'https://example.com', enabled: true }}>
          <ToSContext.Consumer>
            {({ tos }) => <div data-testid="tos">{tos ? tos.url : ''}</div>}
          </ToSContext.Consumer>
        </ToSContextProvider>,
      )

      expect(screen.getByTestId('tos')).toHaveTextContent('https://example.com')
    })
  })
})
