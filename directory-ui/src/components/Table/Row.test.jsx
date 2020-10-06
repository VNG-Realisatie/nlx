// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { shallow } from 'enzyme'
import Row from './Row'

test('should render child elements', () => {
  const wrapper = shallow(
    <Row>
      <th>Table head</th>
    </Row>,
  )
  expect(wrapper.contains(<th>Table head</th>)).toEqual(true)
})
