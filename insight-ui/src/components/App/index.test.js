import React from 'react'
import { shallow } from 'enzyme'
import App from './index'

describe('App component', () => {
  it('should render without errors', () => {
    const wrapper = shallow(<App />)
    expect(wrapper).toBeTruthy()
  })
})
