// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import NotFound from './index'

describe('NotFound', () => {
    const wrapper = shallow(<NotFound />)

    it('should exist', () => {
        expect(wrapper.exists()).toBe(true)
    })
})
