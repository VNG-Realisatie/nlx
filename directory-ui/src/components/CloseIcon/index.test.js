// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { shallow } from 'enzyme'
import CloseIcon from './index'

describe('CloseIcon component', () => {
  it('should render without errors', () => {
    const wrapper = shallow(<CloseIcon />)
    expect(wrapper).toBeTruthy()
  })
})
