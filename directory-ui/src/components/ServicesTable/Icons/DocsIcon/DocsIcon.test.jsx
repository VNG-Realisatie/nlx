import React from 'react'
import { shallow } from 'enzyme'
import DocsIcon from '../DocsIcon/DocsIcon';

it('should support the grey color', () => {
  expect(shallow(<DocsIcon color="grey"/>)).toMatchSnapshot()
})

it('should support the blue color', () => {
  expect(shallow(<DocsIcon color="blue"/>)).toMatchSnapshot()
})
