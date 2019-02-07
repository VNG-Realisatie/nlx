import React from 'react'
import { shallow } from 'enzyme'
import HomePage from './HomePage'

it('renders the HomePage', () => {
    const component = shallow(<HomePage />)
    expect(component).toBeTruthy()
})

it('should match the snapshot', () => {
    const component = shallow(<HomePage />)
    expect(component).toMatchSnapshot()
})
