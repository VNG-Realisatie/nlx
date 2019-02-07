import React from 'react'
import { shallow } from 'enzyme'
import TableRow from './TableRow'

it('should match snapshot', () => {
  const wrapper = shallow(<TableRow/>)
  expect(wrapper).toMatchSnapshot()
})

it('should render child elements', () => {
  const wrapper = shallow(<TableRow><th>Table head</th></TableRow>);
  expect(wrapper.contains(<th>Table head</th>)).toEqual(true)
})