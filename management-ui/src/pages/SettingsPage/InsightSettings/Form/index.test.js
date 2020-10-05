// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import { act, fireEvent, renderWithProviders } from '../../../../test-utils'
import Form from './index'

jest.mock('../../../../components/FormikFocusError', () => () => <></>)

test('Form', async () => {
  const onSubmitHandlerSpy = jest.fn()

  const { getByLabelText, findByTestId } = renderWithProviders(
    <Form onSubmitHandler={onSubmitHandlerSpy} />,
  )

  const formElement = await findByTestId('form')

  await act(async () => {
    fireEvent.submit(formElement)
  })

  expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
    irmaServerURL: '',
    insightAPIURL: '',
  })

  await act(async () => {
    fireEvent.change(getByLabelText('IRMA server URL'), {
      target: { value: 'irma-server-url' },
    })

    fireEvent.change(getByLabelText('Insight API URL'), {
      target: { value: 'insight-api-url' },
    })

    fireEvent.submit(formElement)
  })

  expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
    irmaServerURL: 'irma-server-url',
    insightAPIURL: 'insight-api-url',
  })
})
