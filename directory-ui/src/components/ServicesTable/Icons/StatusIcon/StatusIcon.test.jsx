import React from 'react'
import { shallow } from 'enzyme'
import StatusIcon from './StatusIcon';

it('should support the offline status', () => {
  expect(shallow(<StatusIcon status="online"/>)).toMatchSnapshot()
})

it('should support the online status', () => {
  expect(shallow(<StatusIcon status="offline"/>)).toMatchSnapshot()
})
