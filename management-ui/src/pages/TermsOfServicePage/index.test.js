// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { act, fireEvent, screen } from '@testing-library/react'
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { configure } from 'mobx'
import { renderWithAllProviders } from '../../test-utils'
import { ToSContextProvider } from '../../tos-context'
import { RootStore, StoreProvider } from '../../stores'
import TermsOfServicePage from './index'

test('TermsOfService page', async () => {
  configure({ safeDescriptors: false })

  const rootStore = new RootStore({})

  rootStore.applicationStore.acceptTermsOfService = jest
    .fn()
    .mockRejectedValueOnce({ response: { status: 403 } })
    .mockResolvedValue({})

  const history = createMemoryHistory()

  const tos = { enabled: true, url: 'https://example.com', accepted: false }

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <HistoryRouter history={history}>
        <ToSContextProvider tos={tos}>
          <TermsOfServicePage />
        </ToSContextProvider>
      </HistoryRouter>
    </StoreProvider>,
  )

  expect(await screen.getByRole('link')).toHaveAttribute('href', tos.url)

  const confirmButton = screen.getByText('Confirm agreement')

  await act(async () => {
    fireEvent.click(confirmButton)
  })

  expect(rootStore.applicationStore.acceptTermsOfService).toHaveBeenCalledTimes(
    1,
  )

  expect(screen.queryByRole('alert')).toHaveTextContent(
    "Failed to accept Terms of ServiceYou don't have the required permission.",
  )

  await act(async () => {
    fireEvent.click(confirmButton)
  })

  expect(rootStore.applicationStore.acceptTermsOfService).toHaveBeenCalledTimes(
    2,
  )
  expect(history.location.pathname).toEqual('/')
})
