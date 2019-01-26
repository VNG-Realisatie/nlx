import React from 'react'
import { shallow } from 'enzyme'

import Logo from './Logo'

let component

beforeAll(() => {
    component = shallow(<Logo />)
})

it('renders Logo component', () => {
    expect(component).toBeTruthy()
})
