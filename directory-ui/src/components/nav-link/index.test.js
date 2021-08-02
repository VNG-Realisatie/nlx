// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { render } from '@testing-library/react'
import { ThemeProvider } from 'styled-components'
import theme from '../../styling/theme'
import NavLink from './index'

// https://github.com/vercel/next.js/issues/7479#issuecomment-587145429
const useRouter = jest.spyOn(require('next/router'), 'useRouter')

afterEach(() => {
  jest.clearAllMocks()
})

test('receives active classname', () => {
  useRouter.mockImplementation(() => ({
    pathname: '/page',
  }))

  const { getByText, rerender } = render(
    <ThemeProvider theme={theme}>
      <NavLink to="/">home</NavLink>
    </ThemeProvider>,
  )
  expect(getByText('home')).not.toHaveClass('active')

  rerender(
    <ThemeProvider theme={theme}>
      <NavLink to="/page">page</NavLink>
    </ThemeProvider>,
  )
  expect(getByText('page')).toHaveClass('active')
})

test('prepends basepath', () => {
  useRouter.mockImplementation(() => ({
    basePath: '/site',
  }))

  const { getByText } = render(
    <ThemeProvider theme={theme}>
      <NavLink to="/page">page</NavLink>
    </ThemeProvider>,
  )
  expect(getByText('page')).toHaveAttribute(
    'href',
    expect.stringContaining('/site/page'),
  )
})

test('leaves external links untouched', () => {
  const { getByText, rerender } = render(
    <ThemeProvider theme={theme}>
      <NavLink to="http://commonground.nl">page</NavLink>
    </ThemeProvider>,
  )
  expect(getByText('page')).toHaveAttribute('href', 'http://commonground.nl')

  rerender(
    <ThemeProvider theme={theme}>
      <NavLink to="https://commonground.nl">page</NavLink>
    </ThemeProvider>,
  )
  expect(getByText('page')).toHaveAttribute('href', 'https://commonground.nl')
})
