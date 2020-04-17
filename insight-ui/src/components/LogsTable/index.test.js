// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { shallow } from 'enzyme'
import LogsTable from './index'

test('renders without crashing', () => {
  const wrapper = shallow(<LogsTable />)
  expect(wrapper).toBeTruthy()
})
