// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { node } from 'prop-types'
import { render } from '@testing-library/react'
import { ThemeProvider } from 'styled-components'
import theme from '../theme'

// based on https://testing-library.com/docs/react-testing-library/setup#custom-render
const AllTheProviders = ({ children }) => (
  <ThemeProvider theme={theme}>{children}</ThemeProvider>
)

AllTheProviders.propTypes = {
  children: node,
}

const renderWithProviders = (ui, options) => {
  const reactRoot = document.createElement('div')
  reactRoot.setAttribute('id', 'root')
  document.body.appendChild(reactRoot)
  return render(ui, { wrapper: AllTheProviders, ...options })
}

// re-export everything
export * from '@testing-library/react'

// override render method
export { renderWithProviders }
