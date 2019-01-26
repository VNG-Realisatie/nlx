import React from 'react'
import { shallow } from 'enzyme'

import Table from './Table'

let component

beforeAll(() => {
    component = shallow(<Table />)
})

it('renders Table component', () => {
    expect(component).toBeTruthy()
})
