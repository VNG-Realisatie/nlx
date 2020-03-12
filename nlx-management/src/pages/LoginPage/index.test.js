import React from 'react'
import { I18nextProvider } from 'react-i18next'
import { render } from '@testing-library/react'
import i18n from '../../i18n'
import LoginPage from './index'

test('renders a welcome message', () => {
  const { getByText } = render(
    <I18nextProvider i18n={i18n}>
      <LoginPage />
    </I18nextProvider>,
  )
  expect(getByText(/^Welkom$/)).toBeInTheDocument()
})
