import React from 'react'
import { shallow } from 'enzyme'

import ErrorMessage from './ErrorMessage'

it('should match the snapshot', () => {
    expect(shallow(<ErrorMessage />)).toMatchSnapshot()
})
