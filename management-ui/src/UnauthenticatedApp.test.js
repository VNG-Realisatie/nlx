// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import AuthenticatedApp from './AuthenticatedApp'

it('exists', () => {
    const wrapper = shallow(<AuthenticatedApp />)
    expect(wrapper.exists()).toBe(true)
})
