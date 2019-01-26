import React from 'react'
import { shallow } from 'enzyme'

import Home from './Home'

let component

beforeAll(() => {
    component = shallow(<Home />)
})

it('renders Home component', () => {
    expect(component).toBeTruthy()
})
