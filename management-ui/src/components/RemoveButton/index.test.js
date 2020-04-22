// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../test-utils'
import RemoveButton from './index'

test('renders an icon and text', () => {
  const { container } = renderWithProviders(<RemoveButton />)
  expect(container).toHaveTextContent('bin.svg' + 'Remove') // eslint-disable-line no-useless-concat
})
