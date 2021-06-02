// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import ServicesTable from './index'

test('renders without crashing', () => {
  expect(() => render(<ServicesTable />)).not.toThrow()
})
