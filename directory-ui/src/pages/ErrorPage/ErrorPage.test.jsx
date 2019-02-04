import React from 'react'
import { shallow } from 'enzyme'

import ErrorPage from './ErrorPage'

it('should match the snapshot', () => {
    expect(shallow(<ErrorPage />)).toMatchSnapshot()
})
