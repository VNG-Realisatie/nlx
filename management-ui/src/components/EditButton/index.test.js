// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../test-utils'
import EditButton from './index'

test('renders an icon and text', () => {
  const { container } = renderWithProviders(<EditButton />)
  expect(container).toHaveTextContent('pencil.svg' + 'Edit') // eslint-disable-line no-useless-concat
})
