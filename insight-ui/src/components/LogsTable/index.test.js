import React from 'react'
import { shallow } from 'enzyme'
import LogsTable from './LogsTable'

it('renders without crashing', () => {
  shallow(<LogsTable/>)
})
