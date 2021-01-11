// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import selectEvent from 'react-select-event'
import { act, fireEvent, renderWithProviders } from '../../../../test-utils'
import Form from './index'

test('Form', async () => {
  const onSubmitHandlerSpy = jest.fn()
  const getInwaysHandler = jest.fn().mockResolvedValue([{ name: 'inway-a' }])

  const { findByTestId, getByText, getByLabelText } = renderWithProviders(
    <Form getInways={getInwaysHandler} onSubmitHandler={onSubmitHandlerSpy} />,
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
