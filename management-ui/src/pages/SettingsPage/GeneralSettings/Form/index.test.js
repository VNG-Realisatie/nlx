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

test('Changing organization inway', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [{ name: 'inway-a' }],
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  const onSubmitHandlerSpy = jest.fn()

  const { getByText, getByLabelText, findByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Form onSubmitHandler={onSubmitHandlerSpy} />
    </StoreProvider>,
  )

  const submitButton = getByText('Save settings')

  userEvent.click(submitButton)

  // Shows modal
  expect(await findByText('Are you sure?')).toBeInTheDocument()
  userEvent.click(getByText('Save'))

  await waitFor(() =>
    expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
      organizationInway: '',
    }),
  )

  await selectEvent.select(getByLabelText(/Organization inway/), /inway-a/)
  userEvent.click(submitButton)

  await waitFor(() =>
    expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
      organizationInway: 'inway-a',
    }),
  )
})
