// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../test-utils'
import GlobalAlert from './index'

test('renders as expected', () => {
  const { container } = renderWithProviders(
    <GlobalAlert>
      text <button>button</button>
    </GlobalAlert>,
  )

  // eslint-disable-next-line no-useless-concat
  expect(container).toHaveTextContent('warning.svg' + 'text button')
})
