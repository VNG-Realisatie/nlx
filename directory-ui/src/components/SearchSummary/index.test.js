// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { ThemeProvider } from 'styled-components'
import { renderWithProviders } from '../../test-utils'
import theme from '../../theme'
import SearchSummary from './index'

test('shows correct copy for number of services', () => {
  const { container, rerender } = renderWithProviders(
    <ThemeProvider theme={theme}>
      <SearchSummary
        totalItems={1}
        totalFilteredItems={1}
        itemDescription="beschikbare service"
        itemPluralDescription="beschikbare services"
      />
    </ThemeProvider>,
  )
  expect(container).toHaveTextContent('1 BESCHIKBARE SERVICE')

  rerender(
    <SearchSummary
      totalItems={2}
      totalFilteredItems={2}
      itemDescription="beschikbare service"
      itemPluralDescription="beschikbare services"
    />,
  )
  expect(container).toHaveTextContent('2 BESCHIKBARE SERVICES')

  rerender(
    <SearchSummary
      totalItems={2}
      totalFilteredItems={1}
      itemDescription="beschikbare service"
      itemPluralDescription="beschikbare services"
    />,
  )
  expect(container).toHaveTextContent('1 VAN 2 BESCHIKBARE SERVICES')
})
