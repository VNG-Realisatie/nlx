// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { shallow } from 'enzyme'
import ServicesTable from './ServicesTable'

test('renders without crashing', () => {
  expect(() => {
    shallow(<ServicesTable />)
  }).not.toThrow()
})
