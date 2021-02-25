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

jest.mock('../../../../components/Modal')

const setupComponent = () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [{ name: 'inway-a' }],
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  const onSubmitHandlerSpy = jest.fn()

  const helpers = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Form onSubmitHandler={onSubmitHandlerSpy} />
    </StoreProvider>,
  )

  return {
    ...helpers,
    onSubmitHandlerSpy,
  }
}

test('changing organization inway', async () => {
  const { findByLabelText, getByText, onSubmitHandlerSpy } = setupComponent()

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

test('when not specifying an organization inway, a confirmation should be shown', async () => {
  const { findByText, getByText, onSubmitHandlerSpy } = setupComponent()

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
})
