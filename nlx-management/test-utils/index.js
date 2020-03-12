import React from 'react'
import { render } from '@testing-library/react'
import { I18nextProvider } from 'react-i18next'

import i18n from '../src/i18n'

const AllTheProviders = ({ children }) => (
  <I18nextProvider i18n={i18n}>
    {children}
  </I18nextProvider>
)

const renderWithProviders = (ui, options) =>
  render(ui, { wrapper: AllTheProviders, ...options })

// re-export everything
export * from '@testing-library/react'

// override render method
export { renderWithProviders as render }
