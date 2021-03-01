// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import selectEvent from 'react-select-event'
import { waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { renderWithProviders } from '../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../stores'
import { ManagementApi } from '../../../../api'
import Form from './index'

test('changing organization inway', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [{ name: 'inway-a' }],
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  const onSubmitHandlerSpy = jest.fn()

  const { findByLabelText, getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Form onSubmitHandler={onSubmitHandlerSpy} />
    </StoreProvider>,
  )

  const organizationInwayInput = await findByLabelText(/Organization inway/)
  await selectEvent.select(organizationInwayInput, /inway-a/)
  const submitButton = getByText('Save settings')
  userEvent.click(submitButton)

  await waitFor(() =>
    expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
      organizationInway: 'inway-a',
    }),
  )
})
