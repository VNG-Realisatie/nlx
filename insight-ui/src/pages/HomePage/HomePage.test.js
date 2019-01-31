import React from 'react'
import { shallow } from 'enzyme'

import HomePage from './Home'

let component

beforeAll(() => {
    component = shallow(<HomePage />)
})

it('renders Home component', () => {
    expect(component).toBeTruthy()
})
