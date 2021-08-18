// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { ThemeProvider } from 'styled-components'
import { renderWithProviders } from '../../test-utils'
import theme from '../../styling/theme'
import SearchSummary from './index'

test('shows correct copy for number of services', () => {
  const { container, rerender } = renderWithProviders(
    <ThemeProvider theme={theme}>
      <SearchSummary totalServices={1} totalFilteredServices={1} />
    </ThemeProvider>,
  )
  expect(container).toHaveTextContent('1 BESCHIKBARE SERVICE')

  rerender(<SearchSummary totalServices={2} totalFilteredServices={2} />)
  expect(container).toHaveTextContent('2 BESCHIKBARE SERVICES')

  rerender(<SearchSummary totalServices={2} totalFilteredServices={1} />)
  expect(container).toHaveTextContent('1 VAN 2 BESCHIKBARE SERVICES')
})
