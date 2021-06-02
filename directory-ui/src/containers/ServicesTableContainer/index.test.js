// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import ServicesTableContainer from './index'

test('renders without crashing', () => {
  expect(() => render(<ServicesTableContainer />)).not.toThrow()
})
