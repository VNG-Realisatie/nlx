// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import PageTemplate from './index'

describe('PageTemplate', () => {
    it('should exist', () => {
        const wrapper = shallow(<PageTemplate />)
        expect(wrapper.exists()).toBe(true)
    })
})
