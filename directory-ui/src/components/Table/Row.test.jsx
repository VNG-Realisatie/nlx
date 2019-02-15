import React from 'react'
import { shallow } from 'enzyme'
import Row from './Row'

it('should match snapshot', () => {
  const wrapper = shallow(<Row/>)
  expect(wrapper).toMatchSnapshot()
})

it('should render child elements', () => {
  const wrapper = shallow(<Row><th>Table head</th></Row>);
  expect(wrapper.contains(<th>Table head</th>)).toEqual(true)
})