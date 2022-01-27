// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { act, render } from '@testing-library/react'
import { renderWithProviders } from '../test-utils'
import { RootStore, StoreProvider } from '../stores'
import { DirectoryApi } from '../api'
import { ToSContextProvider } from './index'
import ToSContext from './index'

describe('ToSContext', () => {
  describe('Provider', () => {
    it('should fetch the current terms of service', async () => {
      const directoryApiClient = new DirectoryApi()

      directoryApiClient.directoryGetTermsOfService = jest
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
        directoryApiClient.directoryGetTermsOfService,
      ).toHaveBeenCalledTimes(1)
    })

    it('should make the user available to the context consumers', () => {
      const { getByTestId } = render(
        <ToSContextProvider tos={{ url: 'https://example.com', enabled: true }}>
          <ToSContext.Consumer>
            {({ tos }) => <div data-testid="tos">{tos ? tos.url : ''}</div>}
          </ToSContext.Consumer>
        </ToSContextProvider>,
      )

      expect(getByTestId('tos')).toHaveTextContent('https://example.com')
    })
  })
})
