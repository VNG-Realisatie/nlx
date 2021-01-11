// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import selectEvent from 'react-select-event'
import { act, fireEvent, renderWithProviders } from '../../../../test-utils'
import { RootStore, StoreProvider } from '../../../../stores'
import { ManagementApi } from '../../../../api'
import Form from './index'

test('Form', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [{ name: 'inway-a' }],
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  const onSubmitHandlerSpy = jest.fn()

  const { findByTestId, getByText, getByLabelText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Form onSubmitHandler={onSubmitHandlerSpy} />
    </StoreProvider>,
  )

  const formElement = await findByTestId('form')

  await act(async () => {
    fireEvent.submit(formElement)
  })

  await act(async () => {
    fireEvent.click(getByText('Save'))
  })

  expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
    organizationInway: '',
  })

  await selectEvent.select(getByLabelText(/Organization inway/), /inway-a/)

  await act(async () => {
    fireEvent.submit(formElement)
  })

  expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
    organizationInway: 'inway-a',
  })
})
