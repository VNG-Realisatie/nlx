import React from 'react'
import { shallow } from 'enzyme'
import DocsIcon from '../DocsIcon/DocsIcon';

it('should support the offline status', () => {
  expect(shallow(<DocsIcon status="online"/>)).toMatchSnapshot()
})

it('should support the online status', () => {
  expect(shallow(<DocsIcon status="offline"/>)).toMatchSnapshot()
})
