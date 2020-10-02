// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import { act, fireEvent, renderWithProviders } from '../../../test-utils'
import SettingsForm from './index'

jest.mock('../../../components/FormikFocusError', () => () => <></>)

test('SettingsForm', async () => {
  const onSubmitHandlerSpy = jest.fn()
  const getInwaysHandler = jest.fn().mockResolvedValue([{ name: 'inway-a' }])

  const { getByLabelText, findByTestId } = renderWithProviders(
    <SettingsForm
      getInways={getInwaysHandler}
      onSubmitHandler={onSubmitHandlerSpy}
    />,
  )

  const formElement = await findByTestId('form')

  const inwayField = getByLabelText('Organization inway')

  await act(async () => {
    fireEvent.submit(formElement)
  })

  expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
    organizationInway: '',
  })

  fireEvent.change(inwayField, {
    target: { value: 'inway-a' },
  })

  await act(async () => {
    fireEvent.submit(formElement)
  })

  expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
    organizationInway: 'inway-a',
  })
})
