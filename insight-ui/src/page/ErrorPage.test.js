import React from 'react'
import { shallow } from 'enzyme'

import ErrorPage from './ErrorPage'

let component

beforeAll(() => {
    component = shallow(<ErrorPage />)
})

it('renders ErrorPage component', () => {
    expect(component).toBeTruthy()
})
