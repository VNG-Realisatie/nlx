// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import Sidebar from './index'

describe('Sidebar', () => {
    const wrapper = shallow(<Sidebar />)

    it('should exist', () => {
        expect(wrapper.exists()).toBe(true)
    })
})
