// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../test-utils'
import SpecTable from './index'

test('renders with defaults without crashing', () => {
  let getByText

  expect(
    () =>
      (getByText = renderWithProviders(
        <SpecTable>
          <tbody>
            <SpecTable.Tr>
              <SpecTable.Td>key</SpecTable.Td>
              <SpecTable.Td>value</SpecTable.Td>
            </SpecTable.Tr>
          </tbody>
        </SpecTable>,
      ).getByText),
  ).not.toThrow()

  expect(getByText('value')).not.toHaveStyle('text-align: right')
})

test('renders value text aligned right', () => {
  const { getByText } = renderWithProviders(
    <SpecTable valueAlignRight>
      <tbody>
        <SpecTable.Tr>
          <SpecTable.Td>key</SpecTable.Td>
          <SpecTable.Td>value</SpecTable.Td>
        </SpecTable.Tr>
      </tbody>
    </SpecTable>,
  )

  expect(getByText('value')).toHaveStyle('text-align: right')
})

test('shows a console warning when using more then two Td elements', () => {
  const warnSpy = jest
    .spyOn(console, 'warn')
    .mockImplementation(() => undefined)

  renderWithProviders(
    <SpecTable valueAlignRight>
      <tbody>
        <SpecTable.Tr>
          <SpecTable.Td>key</SpecTable.Td>
          <SpecTable.Td>value1</SpecTable.Td>
          <SpecTable.Td>value2</SpecTable.Td>
        </SpecTable.Tr>
      </tbody>
    </SpecTable>,
  )

  expect(warnSpy).toHaveBeenCalled()
  console.warn.mockRestore()
})
