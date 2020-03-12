// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import App from './App'

test('exists', () => {
  const wrapper = shallow(<App />)
  expect(wrapper.exists()).toBe(true)
})
