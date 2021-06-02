// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import CloseIcon from './index'

test('renders without crashing', () => {
  expect(() => render(<CloseIcon />)).not.toThrow()
})
