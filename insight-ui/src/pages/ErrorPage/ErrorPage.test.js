import React from 'react'
import { shallow } from 'enzyme'

import ErrorPage from './ErrorPage'

it('renders the ErrorPage', () => {
    const component = shallow(<ErrorPage />)
    expect(component).toBeTruthy()
})

it('should match the snapshot', () => {
    const component = shallow(<ErrorPage />)
    expect(component).toMatchSnapshot()
})
