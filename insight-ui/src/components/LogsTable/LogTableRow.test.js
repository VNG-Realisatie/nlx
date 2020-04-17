// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { shallow } from 'enzyme'
import LogTableRow from './LogTableRow'

test('renders without crashing', () => {
  const wrapper = shallow(
    <LogTableRow
      subjects={['a', 'b']}
      requestedBy="foo"
      requestedAt="bar"
      reason="baz"
      date={new Date()}
    />,
  )
  expect(wrapper).toBeTruthy()
})
