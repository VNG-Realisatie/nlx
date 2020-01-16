// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import UnauthenticatedApp from './UnauthenticatedApp'

it('exists', () => {
    const wrapper = shallow(<UnauthenticatedApp />)
    expect(wrapper.exists()).toBe(true)
})
