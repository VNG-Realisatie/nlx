// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import selectEvent from 'react-select-event'
import { act, fireEvent, renderWithProviders } from '../../../../test-utils'
import Form from './index'

test('Form', async () => {
  global.confirm = jest.fn(() => true)

  const onSubmitHandlerSpy = jest.fn()
  const getInwaysHandler = jest.fn().mockResolvedValue([{ name: 'inway-a' }])

  const { getByLabelText, findByTestId } = renderWithProviders(
    <Form getInways={getInwaysHandler} onSubmitHandler={onSubmitHandlerSpy} />,
  )

  const formElement = await findByTestId('form')

  await act(async () => {
    fireEvent.submit(formElement)
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
