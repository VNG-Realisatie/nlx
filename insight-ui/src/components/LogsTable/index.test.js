// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import LogsTable from './index'

test('renders without crashing', () => {
  shallow(<LogsTable />)
})
