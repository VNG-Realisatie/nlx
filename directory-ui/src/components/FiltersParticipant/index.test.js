// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import userEvent from '@testing-library/user-event'
import { renderWithProviders } from '../../test-utils'
import Filters from './index'

test('should call the onQueryChanged handler with the query', () => {
  const onQueryChangedSpy = jest.fn()

  const { getByPlaceholderText } = renderWithProviders(
    <Filters onQueryChanged={onQueryChangedSpy} />,
  )

  const input = getByPlaceholderText('Zoeken…')

  userEvent.clear(input)
  userEvent.type(input, 'abc')

  expect(onQueryChangedSpy).toHaveBeenCalledWith('abc')
})

test('should show the correct default value when rendering the dropdown', () => {
  const { location } = window
  delete window.location
  window.location = { hostname: 'directory.demo.nlx' }

  const { container } = renderWithProviders(<Filters />)

  expect(container).toHaveTextContent(/Demo/)

  window.location = location
})
