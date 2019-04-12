import React from 'react'
import { shallow } from 'enzyme'
import LogTableRow from './LogTableRow'

it('renders without crashing', () => {
  shallow(<LogTableRow subjects={['a', 'b']}
                       requestedBy="foo"
                       requestedAt="bar"
                       reason="baz"
                       date={new Date()}
  />)
})
