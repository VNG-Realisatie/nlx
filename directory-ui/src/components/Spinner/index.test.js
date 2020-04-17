// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import ReactDOM from 'react-dom'
import { shallow } from 'enzyme'

import Spinner from './index'

test('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Spinner />, div)
  }).not.toThrow()
})

test('should contain 8 dots', () => {
  const tree = shallow(<Spinner />)
  expect(tree.find('[data-test="bullet"]')).toHaveLength(8)
})
