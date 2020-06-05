// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../test-utils'
import SpecList from './index'

test('renders one value with defaults without crashing', () => {
  let getByText

  expect(() => {
    const rendered = renderWithProviders(
      <SpecList>
        <SpecList.Item title="key" value="value" />
      </SpecList>,
    )

    getByText = rendered.getByText
  }).not.toThrow()

  expect(getByText('value')).not.toHaveStyle('text-align: right')
})

test('renders values with text aligned right', () => {
  const { getByText } = renderWithProviders(
    <SpecList alignValuesRight>
      <SpecList.Item title="key1" value="value1" />
      <SpecList.Item title="key2" value="value2" />
    </SpecList>,
  )

  expect(getByText('value1')).toHaveStyle('text-align: right')
  expect(getByText('value2')).toHaveStyle('text-align: right')
})

test('renders values with text aligned right', () => {
  const { getByText } = renderWithProviders(
    <SpecList alignValuesRight>
      <SpecList.Item title="key1" value="value1" />
      <SpecList.Item title="key2" value="value2" alignValue="left" />
      <SpecList.Item title="key3" value="value3" />
    </SpecList>,
  )

  expect(getByText('value1')).toHaveStyle('text-align: right')
  expect(getByText('value2')).toHaveStyle('text-align: left')
  expect(getByText('value2')).not.toHaveStyle('text-align: right')
  expect(getByText('value3')).toHaveStyle('text-align: right')
})

test('it allows addition classNames to be set (eg. wrapped with styled-components)', () => {
  const { container, getByText } = renderWithProviders(
    <SpecList className="uno dos" alignValuesRight>
      <SpecList.Item title="key" value="value" />
    </SpecList>,
  )

  const list = container.querySelector('.uno')
  expect(container).toContainElement(list)
  expect(list).toHaveClass('dos')
  expect(getByText('value')).toHaveStyle('text-align: right')
})
