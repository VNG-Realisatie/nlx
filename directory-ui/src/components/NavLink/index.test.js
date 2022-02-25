// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { BrowserRouter as Router } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import NavLink from './index'

test('receives active classname', () => {
  window.history.pushState({}, 'Test page', '/page')

  const { getByText, rerender } = renderWithProviders(
    <NavLink to="/other-page">home</NavLink>,
    { wrapper: Router },
  )
  expect(getByText('home')).not.toHaveClass('active')

  rerender(<NavLink to="/page">page</NavLink>)
  expect(getByText('page')).toHaveClass('active')
})

test('works with external links', () => {
  renderWithProviders(
    <Router>
      <NavLink to="https://nlx.io">nlx.io</NavLink>
    </Router>,
  )

  const link = screen.getByText('nlx.io')

  expect(link.getAttribute('href')).toBe('https://nlx.io')
})
