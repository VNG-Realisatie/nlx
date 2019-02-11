import React from 'react'
import { shallow } from 'enzyme'
import LinkIcon from '../LinkIcon/LinkIcon';

it('should support the offline status', () => {
  expect(shallow(<LinkIcon status="online"/>)).toMatchSnapshot()
})

it('should support the online status', () => {
  expect(shallow(<LinkIcon status="offline"/>)).toMatchSnapshot()
})
