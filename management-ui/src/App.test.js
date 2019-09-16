// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import App from './App'

it('exists', () => {
    const wrapper = shallow(<App />)
    expect(wrapper.exists()).toBe(true)
})
