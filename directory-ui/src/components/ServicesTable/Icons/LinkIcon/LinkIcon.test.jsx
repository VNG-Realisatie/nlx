import React from 'react'
import { shallow } from 'enzyme'
import LinkIcon from '../LinkIcon/LinkIcon';

it('should support the grey color', () => {
  expect(shallow(<LinkIcon color="grey"/>)).toMatchSnapshot()
})

it('should support the blue color', () => {
  expect(shallow(<LinkIcon color="blue"/>)).toMatchSnapshot()
})
