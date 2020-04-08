// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { node } from 'prop-types'
import { render } from '@testing-library/react'
import { I18nextProvider } from 'react-i18next'

import { ThemeProvider } from 'styled-components'
import theme from '../theme'
import i18n from './i18nTestConfig'

// based on https://testing-library.com/docs/react-testing-library/setup#custom-render
const AllTheProviders = ({ children }) => (
  <ThemeProvider theme={theme}>
    <I18nextProvider i18n={i18n}>{children}</I18nextProvider>
  </ThemeProvider>
)
AllTheProviders.propTypes = {
  children: node,
}

const renderWithProviders = (ui, options) =>
  render(ui, { wrapper: AllTheProviders, ...options })

// re-export everything
export * from '@testing-library/react'

// override render method
export { renderWithProviders }
