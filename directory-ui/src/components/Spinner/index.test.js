import React from 'react';
import ReactDOM from 'react-dom'
import { shallow } from 'enzyme'

import Spinner from './index'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Spinner />, div)
  }).not.toThrow()
})

it('should contain 8 dots', () => {
  const tree = shallow(<Spinner />);
  expect(tree.find('[data-test="bullet"]').length).toBe(8);
});
