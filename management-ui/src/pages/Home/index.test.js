// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import Home from './index'

describe('Home', () => {
    const wrapper = shallow(<Home />)

    it('should exist', () => {
        expect(wrapper.exists()).toBe(true)
    })
})
