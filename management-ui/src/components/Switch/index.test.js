// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { render } from '@testing-library/react'

import Switch from './index'

test('renders expected component', () => {
  const { getByText, queryByText } = render(
    <Switch test="case2">
      <Switch.Case value="case1">{() => 'case1'}</Switch.Case>
      <Switch.Case value="case2">{() => 'case2'}</Switch.Case>
      <Switch.Default>{() => 'default'}</Switch.Default>
    </Switch>,
  )

  expect(queryByText('case1')).not.toBeInTheDocument()
  expect(getByText('case2')).toBeInTheDocument()
  expect(queryByText('default')).not.toBeInTheDocument()
})

test('renders expected component when matching one of multiple values', () => {
  const { getByText, queryByText } = render(
    <Switch test="case3">
      <Switch.Case value={['case1', 'case3']}>{() => 'case1or3'}</Switch.Case>
      <Switch.Case value="case2">{() => 'case2'}</Switch.Case>
      <Switch.Default>{() => 'default'}</Switch.Default>
    </Switch>,
  )

  expect(getByText('case1or3')).toBeInTheDocument()
  expect(queryByText('case2')).not.toBeInTheDocument()
  expect(queryByText('default')).not.toBeInTheDocument()
})

test('renders the default component', () => {
  const { getByText, queryByText } = render(
    <Switch test={42}>
      <Switch.Case value={['case1', 'case3']}>{() => 'case1or3'}</Switch.Case>
      <Switch.Case value="case2">{() => 'case2'}</Switch.Case>
      <Switch.Default>{() => 'default'}</Switch.Default>
    </Switch>,
  )

  expect(queryByText('case1')).not.toBeInTheDocument()
  expect(queryByText('case2')).not.toBeInTheDocument()
  expect(getByText('default')).toBeInTheDocument()
})

test('shows error when no cases match and no default given', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})

  render(
    <Switch test={42}>
      <Switch.Case value="case1">{() => 'case1'}</Switch.Case>
      <Switch.Case value="case2">{() => 'case2'}</Switch.Case>
    </Switch>,
  )

  expect(errorSpy).toHaveBeenCalled()
  errorSpy.mockRestore()
})
