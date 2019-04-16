import React from 'react'
import { shallow } from 'enzyme'
import LogsTable from './index'

it('renders without crashing', () => {
  shallow(<LogsTable/>)
})
