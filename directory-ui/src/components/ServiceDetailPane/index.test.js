// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import ServiceDetailPane from './index'

describe('ServiceDetailPane component', () => {
  it('should render without errors', () => {
    const wrapper = shallow(<ServiceDetailPane />)
    expect(wrapper).toBeTruthy()
  })
})
