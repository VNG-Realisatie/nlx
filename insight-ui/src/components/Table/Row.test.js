import React from 'react'
import Row from './Row'

xit('should render child elements', () => {
  const wrapper = shallow(<Row><th>Table head</th></Row>);
  expect(wrapper.contains(<th>Table head</th>)).toEqual(true)
});
