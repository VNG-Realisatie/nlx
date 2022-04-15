// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import userEvent from '@testing-library/user-event'
import { renderWithProviders } from '../../test-utils'
import Filters from './index'

test('should call the onQueryChanged handler with the query', async () => {
  const onQueryChangedSpy = jest.fn()

  const { getByPlaceholderText } = renderWithProviders(
    <Filters onQueryChanged={onQueryChangedSpy} />,
  )

  const input = getByPlaceholderText('Zoeken…')

  await userEvent.clear(input)
  await userEvent.type(input, 'abc')

  expect(onQueryChangedSpy).toHaveBeenCalledWith('abc')
})

test('should call the onStatusFilterChanged handler with the checked state', async () => {
  const onStatusFilterChangedSpy = jest.fn()

  const { getByLabelText } = renderWithProviders(
    <Filters onStatusFilterChanged={onStatusFilterChangedSpy} />,
  )

  const input = getByLabelText('Toon offline services')

  await userEvent.click(input)

  expect(onStatusFilterChangedSpy).toHaveBeenCalledWith(false)
})

test('should show the correct default value when rendering the dropdown', () => {
  const { location } = window
  delete window.location
  window.location = { hostname: 'directory.demo.nlx' }

  const { container } = renderWithProviders(<Filters />)

  expect(container).toHaveTextContent(/Demo/)

  window.location = location
})
