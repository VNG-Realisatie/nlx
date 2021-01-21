// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../../../test-utils'
import CollapsibleHeader from './index'

test('the CollapsibleHeader component', async () => {
  const { rerender, getByText } = renderWithProviders(
    <CollapsibleHeader counter={0} />,
  )

  const toggler = getByText(/Organizations with access/i)

  expect(toggler).toHaveTextContent(
    'checkbox-multiple.svg' + 'Organizations with access' + '0', // eslint-disable-line no-useless-concat
  )

  rerender(<CollapsibleHeader counter={1} />)

  expect(toggler).toHaveTextContent(
    'checkbox-multiple.svg' + 'Organizations with access' + '1', // eslint-disable-line no-useless-concat
  )
})
