import React from 'react'
import { shallow } from 'enzyme'
import HomePage from './index'

describe('HomePage component', () => {
  it('should render without errors', () => {
    const wrapper = shallow(<HomePage />)
    expect(wrapper).toBeTruthy()
  })
})
