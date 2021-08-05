// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import SectionGroup from './index'

test('rendering one child', () => {
  const { getByTestId } = renderWithProviders(
    <SectionGroup>
      <span data-testid="one" />
    </SectionGroup>,
  )
  expect(getByTestId('one')).toBeTruthy()
})

test('rendering multiple children', () => {
  const { getByTestId, container } = renderWithProviders(
    <SectionGroup>
      <span data-testid="one" />
      <span data-testid="two" />
    </SectionGroup>,
  )
  expect(getByTestId('one')).toBeInTheDocument()
  expect(getByTestId('two')).toBeInTheDocument()
  expect(container.querySelectorAll('span')).toHaveLength(2)
})

test('nulls are filtered out', () => {
  const showOptionalComponent = false
  const { getByTestId, queryByTestId, container } = renderWithProviders(
    <SectionGroup>
      <span data-testid="one" />
      {showOptionalComponent ? <span data-testid="two" /> : null}
      <span data-testid="three" />
    </SectionGroup>,
  )

  expect(getByTestId('one')).toBeInTheDocument()
  expect(queryByTestId('two')).not.toBeInTheDocument()
  expect(getByTestId('three')).toBeInTheDocument()
  expect(container.querySelectorAll('span')).toHaveLength(2)
})
